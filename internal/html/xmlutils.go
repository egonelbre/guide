package html

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func XMLText(decoder *xml.Decoder) (string, error) {
	r := ""
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return r, io.EOF
		}
		if err != nil {
			return r, err
		}

		switch token := token.(type) {
		case xml.EndElement:
			return r, nil
		case xml.CharData:
			r += string(token)
		case xml.StartElement:
			sub, err := XMLText(decoder)
			r += sub
			if err != nil {
				return r, err
			}
		case xml.Comment: // ignore
		case xml.ProcInst: // ignore
		case xml.Directive: // ignore
		default:
			panic(fmt.Sprintf("unknown token %T=%v", token, token))
		}
	}
}

func XMLStripTags(xmlcontent string) (string, error) {
	if xmlcontent == "" {
		return "", nil
	}
	return XMLText(xml.NewDecoder(strings.NewReader(xmlcontent)))
}
