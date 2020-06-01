package transformations

import (
	"bytes"
	"unicode"
	"unicode/utf8"

	"github.com/asalih/guardian/helpers"
)

func init() {
	TransformationMaps.funcMap["cssDecode"] = func(variableData interface{}) interface{} {
		dec := decodeCSS([]byte(variableData.(string)))
		return string(dec)
	}
}

func decodeCSS(s []byte) []byte {
	i := bytes.IndexByte(s, '\\')
	if i == -1 {
		return s
	}
	// The UTF-8 sequence for a codepoint is never longer than 1 + the
	// number hex digits need to represent that codepoint, so len(s) is an
	// upper bound on the output length.
	b := make([]byte, 0, len(s))
	for len(s) != 0 {
		i := bytes.IndexByte(s, '\\')
		if i == -1 {
			i = len(s)
		}
		b, s = append(b, s[:i]...), s[i:]
		if len(s) < 2 {
			break
		}

		// http://www.w3.org/TR/css3-syntax/#SUBTOK-escape
		// escape ::= unicode | '\' [#x20-#x7E#x80-#xD7FF#xE000-#xFFFD#x10000-#x10FFFF]
		if helpers.IsHex(s[1]) {
			// http://www.w3.org/TR/css3-syntax/#SUBTOK-unicode
			//   unicode ::= '\' [0-9a-fA-F]{1,6} wc?
			j := 2
			for j < len(s) && j < 7 && helpers.IsHex(s[j]) {
				j++
			}
			r := helpers.HexDecode(s[1:j])
			if r > unicode.MaxRune {
				r, j = r/16, j-1
			}
			n := utf8.EncodeRune(b[len(b):cap(b)], r)
			// The optional space at the end allows a hex
			// sequence to be followed by a literal hex.
			// string(decodeCSS([]byte(`\A B`))) == "\nB"
			b, s = b[:len(b)+n], skipCSSSpace(s[j:])
		} else {
			// `\\` decodes to `\` and `\"` to `"`.
			_, n := utf8.DecodeRune(s[1:])
			b, s = append(b, s[1:1+n]...), s[1+n:]
		}
	}
	return b
}

// skipCSSSpace returns a suffix of c, skipping over a single space.
func skipCSSSpace(c []byte) []byte {
	if len(c) == 0 {
		return c
	}
	// wc ::= #x9 | #xA | #xC | #xD | #x20
	switch c[0] {
	case '\t', '\n', '\f', ' ':
		return c[1:]
	case '\r':
		// This differs from CSS3's wc production because it contains a
		// probable spec error whereby wc contains all the single byte
		// sequences in nl (newline) but not CRLF.
		if len(c) >= 2 && c[1] == '\n' {
			return c[2:]
		}
		return c[1:]
	}
	return c
}
