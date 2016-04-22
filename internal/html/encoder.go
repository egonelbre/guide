package html

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

type Encoder struct {
	RewriteID string

	buf bytes.Buffer
	w   io.Writer

	stack  []string
	invoid bool
}

func NewEncoder(out io.Writer) *Encoder {
	return &Encoder{
		buf:   bytes.Buffer{},
		w:     out,
		stack: []string{},
	}
}

func (enc *Encoder) Depth() int      { return len(enc.stack) }
func (enc *Encoder) Stack() []string { return enc.stack }

func (enc *Encoder) WriteXMLStart(token *xml.StartElement) error {
	return enc.WriteStart(token.Name.Local, token.Attr...)
}

func (enc *Encoder) WriteStart(tag string, attrs ...xml.Attr) error {
	enc.stack = append(enc.stack, tag)
	enc.invoid = voidElements[tag]

	enc.buf.WriteByte('<')
	enc.buf.WriteString(tag)

	for _, attr := range attrs {
		if attr.Name.Local == "" {
			continue
		}

		enc.buf.WriteByte(' ')
		if attr.Name.Local == "id" && enc.RewriteID != "" {
			enc.buf.WriteString(enc.RewriteID)
		} else {
			if attr.Name.Space != "" {
				enc.buf.WriteString(attr.Name.Space + ":" + attr.Name.Local)
			} else {
				enc.buf.WriteString(attr.Name.Local)
			}
		}
		enc.buf.WriteString(`="`)
		enc.buf.WriteString(EscapeAttribute(attr.Value))
		enc.buf.WriteByte('"')
	}
	enc.buf.WriteByte('>')

	return enc.flush()
}

func (enc *Encoder) WriteXMLEnd(token *xml.EndElement) error {
	return enc.WriteEnd(token.Name.Local)
}

func (enc *Encoder) WriteEnd(tag string) error {
	if len(enc.stack) == 0 {
		return fmt.Errorf("no unclosed tags")
	}

	var current string
	n := len(enc.stack) - 1
	current, enc.stack = enc.stack[n], enc.stack[:n]
	if current != tag {
		return fmt.Errorf("writing end tag %v expected %v", tag, current)
	}

	enc.invoid = (len(enc.stack) > 0) && voidElements[enc.stack[len(enc.stack)-1]]

	// void elements have only a single tag
	if voidElements[tag] {
		return nil
	}

	enc.buf.WriteString("</")
	enc.buf.WriteString(tag)
	enc.buf.WriteByte('>')

	return enc.flush()
}

func (enc *Encoder) WriteRaw(data string) error {
	_, err := enc.buf.WriteString(data)
	return err
}

func (enc *Encoder) voiderror() error {
	return fmt.Errorf("content not allowed inside void tag %s", enc.stack[len(enc.stack)-1])
}

func (enc *Encoder) Encode(token xml.Token) error {
	switch token := token.(type) {
	case xml.StartElement:
		return enc.WriteXMLStart(&token)
	case xml.EndElement:
		return enc.WriteXMLEnd(&token)
	case xml.CharData:
		if enc.invoid {
			return enc.voiderror()
		}
		enc.buf.Write([]byte(EscapeCharData(string(token))))
		return enc.flush()
	case xml.Comment:
		if enc.invoid {
			return enc.voiderror()
		}
		enc.buf.WriteString("<!--")
		enc.buf.Write([]byte(EscapeCharData(string(token))))
		enc.buf.WriteString("-->")
		return enc.flush()
	case xml.ProcInst:
		if enc.invoid {
			return enc.voiderror()
		}
		// skip processing instructions
		return nil
	case xml.Directive:
		if enc.invoid {
			return enc.voiderror()
		}
		// skip directives
		return nil
	default:
		panic("invalid token")
	}
}

func (enc *Encoder) flush() error {
	if enc.buf.Len() > 1<<8 {
		return enc.Flush()
	}
	return nil
}

func (enc *Encoder) Flush() error {
	_, err := enc.buf.WriteTo(enc.w)
	enc.buf.Reset()
	return err
}

// Section 12.1.2, "Elements", gives this list of void elements. Void elements
// are those that can't have any contents.
var voidElements = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}
