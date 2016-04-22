package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/loov/guide/internal/html"
)

type renaming struct {
	tag   string
	class string
}

var rename = map[string]renaming{
	"title":     renaming{"h", ""},
	"premise":   renaming{"div", "premise"},
	"therefore": renaming{"div", "therefore"},

	"link": renaming{"span", "link"}, //TODO: fix

	"todo":     renaming{"span", "todo"},
	"todo-big": renaming{"div", "todo"},
}

type Context struct {
	FileSystem
	Path string

	SectionDepth int

	Output  *bytes.Buffer
	Decoder *xml.Decoder
	Encoder *html.Encoder

	Errors []error
}

func NewContext(fs FileSystem, path string) (*Context, error) {
	context := &Context{
		FileSystem: fs,
		Path:       path,

		Output: bytes.NewBuffer(nil),
	}

	data, _, err := context.FileSystem.ReadFile(path)
	context.check(err)

	context.Decoder = xml.NewDecoder(bytes.NewReader(data))
	context.Encoder = html.NewEncoder(context.Output)

	return context, err
}

func (context *Context) check(err error) {
	if err != nil {
		context.Errors = append(context.Errors, fmt.Errorf("%v: %v", context.Path, err))
	}
}

func (context *Context) HandleQuote(start xml.StartElement) error {
	context.Encoder.WriteStart("div", attr("class", "quote"))

	context.Encoder.WriteStart("div", attr("class", "text"))
	err := context.Recurse()
	context.Encoder.WriteEnd("div")

	href := getattr(&start, "href")
	by := getattr(&start, "by")

	if href != "" || by != "" {
		context.Encoder.WriteStart("div", attr("class", "by"))

		if href != "" {
			context.Encoder.WriteStart("a", attr("href", href))
		}
		context.Encoder.WriteRaw(html.EscapeString(by))
		if href != "" {
			context.Encoder.WriteEnd("a")
		}

		context.Encoder.WriteEnd("div")
	}

	context.Encoder.WriteEnd("div")
	return err
}

func (context *Context) HandleInclude(start xml.StartElement) error {
	context.Decoder.Skip()

	include := path.Join(path.Dir(context.Path), getattr(&start, "src"))

	child, err := NewContext(context.FileSystem, include)
	context.check(err)

	if err != nil {
		context.Encoder.WriteStart("div", attr("class", "invalid-include"))
		context.Encoder.WriteRaw(html.EscapeString(context.Path) + " with " + html.EscapeString(getattr(&start, "src")))
		context.Encoder.WriteEnd("div")
		return nil
	}

	child.Output = nil
	child.SectionDepth = context.SectionDepth
	child.Encoder = context.Encoder

	child.Recurse()

	context.Errors = append(context.Errors, child.Errors...)

	return nil
}

func (context *Context) HandleImage(start xml.StartElement) error {
	context.Encoder.WriteStart("figure")

	start.Name.Local = "img"
	if src := getattr(&start, "src"); src != "" {
		full := path.Join(path.Dir(context.Path), src)
		abs := path.Clean(full)
		setattr(&start, "src", abs)
	}

	if err := context.Encoder.Encode(start); err != nil {
		return err
	}
	context.check(context.Encoder.Encode(xml.EndElement{start.Name}))

	err := context.Recurse()
	context.Encoder.WriteEnd("figure")
	return err
}

func (context *Context) HandleGroup(start xml.StartElement) error {
	title := getattr(&start, "title")
	setattr(&start, "title", "")

	if err := context.Encoder.Encode(start); err != nil {
		return err
	}

	if title != "" {
		context.Encoder.WriteStart("div", attr("class", "separator"))
		context.Encoder.WriteRaw(html.EscapeString(title))
		context.Encoder.WriteEnd("div")
	}

	err := context.Recurse()
	context.check(context.Encoder.Encode(xml.EndElement{start.Name}))
	return err
}

func (context *Context) HandleTitle(start xml.StartElement) error {
	tag := fmt.Sprintf("h%d", context.SectionDepth-1)

	context.Encoder.WriteStart(tag)
	err := context.Recurse()
	context.Encoder.WriteEnd(tag)

	return err
}

func (context *Context) HandleCode(start xml.StartElement) error {
	//lang := getattr(&start, "language")
	file := getattr(&start, "file")

	setattr(&start, "language", "")
	setattr(&start, "file", "")

	context.Encoder.WriteStart("figure")

	start.Name.Local = "pre"
	if err := context.Encoder.Encode(start); err != nil {
		return err
	}

	content, err := html.XMLText(context.Decoder)
	context.check(err)
	context.Encoder.WriteRaw(html.EscapeString(StripIdentation(content)))

	context.check(context.Encoder.Encode(xml.EndElement{start.Name}))

	if file != "" {
		context.Encoder.WriteStart("div", attr("class", "caption"))
		context.Encoder.WriteRaw(html.EscapeString(file))
		context.Encoder.WriteEnd("div")
	}

	context.Encoder.WriteEnd("figure")
	return err
}

func (context *Context) Recurse() error {
	for {
		token, err := context.Decoder.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if _, ended := token.(xml.EndElement); ended {
			return nil
		}

		if err := context.Handle(token); err != nil {
			return err
		}
	}
}

func (context *Context) Handle(token xml.Token) error {
	// if shouldskip(token){
	// 	dec.Skip()
	// 	return nil
	// }

	startdepth := context.Encoder.Depth()
	startstack := append([]string{}, context.Encoder.Stack()...)
	defer func() {
		if startdepth != context.Encoder.Depth() {
			fmt.Println(startstack, " => ", context.Encoder.Stack())
			panic("mismatched start and end tag in html output")
		}
	}()

	if start, isStart := token.(xml.StartElement); isStart {
		switch start.Name.Local {
		case "code":
			return context.HandleCode(start)
		case "quote":
			return context.HandleQuote(start)
		case "image":
			return context.HandleImage(start)
		case "include":
			return context.HandleInclude(start)
		case "group":
			return context.HandleGroup(start)
		case "title":
			return context.HandleTitle(start)
		case "section":
			context.SectionDepth++
			defer func() { context.SectionDepth-- }()
		}

		// is it custom already before naming
		// if process, isCustom := context.Rules.Custom[start.Name.Local]; isCustom {
		// 	return process(context, start)
		// }

		if renaming, ok := rename[start.Name.Local]; ok {
			start.Name.Local = renaming.tag
			if renaming.class != "" {
				prev := getattr(&start, "class")
				if prev == "" {
					setattr(&start, "class", renaming.class)
				} else {
					setattr(&start, "class", prev+" "+renaming.class)
				}
			}
		}

		// if process, custom := customs[start.Name.Local]; custom {
		//
		// }

		return context.EmitWithChildren(start)
	}

	return context.Encoder.Encode(token)
}

func (context *Context) EmitWithChildren(start xml.StartElement) error {
	if err := context.Encoder.Encode(start); err != nil {
		return err
	}
	err := context.Recurse()
	context.check(context.Encoder.Encode(xml.EndElement{start.Name}))
	return err
}

func compile(input, output string) {
	context, err := NewContext(Dir("."), input)
	if err != nil {
		fmt.Printf("\n\n= %s\n", input)
		fmt.Println(err)
		return
	}

	context.Encoder.WriteStart("html")

	context.Encoder.WriteStart("head")
	context.Encoder.WriteRaw("<title>A Guide to ...</title>")
	context.Encoder.WriteRaw("<link rel='stylesheet' href='assets/main.css'>")
	context.Encoder.WriteEnd("head")

	context.Encoder.WriteStart("body")

	err = context.Recurse()
	if len(context.Errors) > 0 || err != nil {
		fmt.Printf("\n\n= %s\n", input)
		for _, err := range context.Errors {
			fmt.Println(err)
		}
		fmt.Println(err)
	}

	context.Encoder.WriteEnd("body")
	context.Encoder.WriteEnd("html")

	context.Encoder.Flush()

	ioutil.WriteFile(output, context.Output.Bytes(), 0755)
}

func main() {
	compile("programming/_index.xml", "programming.html")
	compile("patterns/_index.xml", "patterns.html")
	compile("coding/_index.xml", "coding.html")
	compile("algorithms/_index.xml", "algorithms.html")
	compile("software/_index.xml", "software.html")
}

type FileSystem interface {
	ReadFile(path string) (data []byte, modified time.Time, err error)
}

type Dir string

func (dir Dir) fullpath(name string) string {
	return filepath.FromSlash(path.Join(string(dir), name))
}

func (dir Dir) ReadFile(name string) (data []byte, modified time.Time, err error) {
	var file *os.File

	file, err = os.Open(dir.fullpath(name))
	if err != nil {
		return
	}

	var stat os.FileInfo
	stat, err = file.Stat()
	if err != nil {
		return
	}
	modified = stat.ModTime()

	data, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}

	return
}

func StripIdentation(content string) string {
	skip := []rune{}
	content = strings.Trim(content, "\n\r")
	for _, r := range content {
		if unicode.IsSpace(r) || r == '\t' || r == ' ' {
			skip = append(skip, r)
		} else {
			break
		}
	}

	var buf bytes.Buffer

	skipper := 0
	for _, r := range content {
		if r == '\r' {
			continue
		}

		if r == '\n' {
			skipper = 0
			buf.WriteRune(r)
			continue
		}

		if skipper < len(skip) {
			if skip[skipper] == r {
				skipper++
				continue
			}
			skipper = len(skip)
		}

		buf.WriteRune(r)
	}

	return buf.String()
}

func attr(name, value string) xml.Attr {
	return xml.Attr{xml.Name{"", name}, value}
}

func getattr(n *xml.StartElement, key string) (val string) {
	for _, attr := range n.Attr {
		if attr.Name.Local == key {
			return attr.Value
		}
	}
	return ""
}

func setattr(n *xml.StartElement, key, val string) {
	n.Attr = append([]xml.Attr{}, n.Attr...)

	for i := range n.Attr {
		attr := &n.Attr[i]
		if attr.Name.Local == key {
			if val == "" {
				n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
			} else {
				attr.Value = val
			}
			return
		}
	}

	if val == "" {
		return
	}

	n.Attr = append(n.Attr, xml.Attr{
		Name:  xml.Name{Local: key},
		Value: val,
	})
	sort.Sort(attrByName(n.Attr))
}

type attrByName []xml.Attr

func (xs attrByName) Len() int           { return len(xs) }
func (xs attrByName) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs attrByName) Less(i, j int) bool { return xs[i].Name.Local < xs[j].Name.Local }
