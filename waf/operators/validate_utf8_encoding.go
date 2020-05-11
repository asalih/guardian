package operators

//UnicodeErrorCharactersMissing ...
var UnicodeErrorCharactersMissing int = -1

//UnicodeErrorInvalidEncoding ...
var UnicodeErrorInvalidEncoding int = -2

//UnicodeErrorOverlongCharacter ...
var UnicodeErrorOverlongCharacter int = -3

//UnicodeErrorRestrictedCharacter ...
var UnicodeErrorRestrictedCharacter int = -4

//UnicodeErrorDecodingError ...
var UnicodeErrorDecodingError int = -5

func init() {
	OperatorMaps.funcMap["validateUtf8Encoding"] = func(expression interface{}, variableData interface{}) bool {
		data := variableData.(string)

		i := 0
		bytesLeft := len(data)

		for i < bytesLeft {
			rc := detectUtf8Character(int(data[i]), bytesLeft)

			if rc <= 0 {
				return true
			}

			i += rc
			bytesLeft -= rc
		}

		return false
	}
}

func detectUtf8Character(pRead int, length int) int {
	unicodeLen := 0
	d := 0
	c := 0

	if pRead == 0 {
		return UnicodeErrorDecodingError
	}
	c = pRead

	/* If first byte begins with binary 0 it is single byte encoding */
	if (c & 0x80) == 0 {
		/* single byte unicode (7 bit ASCII equivalent) has no validation */
		return 1
	} else if (c & 0xE0) == 0xC0 {
		/* If first byte begins with binary 110 it is two byte encoding*/
		/* check we have at least two bytes */
		if length < 2 {
			unicodeLen = UnicodeErrorCharactersMissing
		} else if ((pRead + 1) & 0xC0) != 0x80 {
			/* check second byte starts with binary 10 */
			unicodeLen = UnicodeErrorInvalidEncoding
		} else {
			unicodeLen = 2
			/* compute character number */
			d = ((c & 0x1F) << 6) | ((pRead + 1) & 0x3F)
		}
	} else if (c & 0xF0) == 0xE0 {
		/* If first byte begins with binary 1110 it is three byte encoding */
		/* check we have at least three bytes */
		if length < 3 {
			unicodeLen = UnicodeErrorCharactersMissing
		} else if ((pRead + 1) & 0xC0) != 0x80 {
			/* check second byte starts with binary 10 */
			unicodeLen = UnicodeErrorInvalidEncoding
		} else if ((pRead + 2) & 0xC0) != 0x80 {
			/* check third byte starts with binary 10 */
			unicodeLen = UnicodeErrorInvalidEncoding
		} else {
			unicodeLen = 3
			/* compute character number */
			d = ((c & 0x0F) << 12) | (((pRead + 1) & 0x3F) << 6) | ((pRead + 2) & 0x3F)
		}
	} else if (c & 0xF8) == 0xF0 {
		/* If first byte begins with binary 11110 it is four byte encoding */
		/* restrict characters to UTF-8 range (U+0000 - U+10FFFF)*/
		if c >= 0xF5 {
			return UnicodeErrorRestrictedCharacter
		}
		/* check we have at least four bytes */
		if length < 4 {
			unicodeLen = UnicodeErrorCharactersMissing
		} else if ((pRead + 1) & 0xC0) != 0x80 {
			unicodeLen = UnicodeErrorInvalidEncoding
		} else if ((pRead + 2) & 0xC0) != 0x80 {
			unicodeLen = UnicodeErrorInvalidEncoding
		} else if ((pRead + 3) & 0xC0) != 0x80 {
			unicodeLen = UnicodeErrorInvalidEncoding
		} else {
			unicodeLen = 4
			/* compute character number */
			d = ((c & 0x07) << 18) | (((pRead + 1) & 0x3F) << 12) | (((pRead + 2) & 0x3F) << 6) | ((pRead + 3) & 0x3F)
		}
	} else {
		/* any other first byte is invalid (RFC 3629) */
		return UnicodeErrorInvalidEncoding
	}

	/* invalid UTF-8 character number range (RFC 3629) */
	if (d >= 0xD800) && (d <= 0xDFFF) {
		return UnicodeErrorRestrictedCharacter
	}

	/* check for overlong */
	if (unicodeLen == 4) && (d < 0x010000) {
		/* four byte could be represented with less bytes */
		return UnicodeErrorOverlongCharacter
	} else if (unicodeLen == 3) && (d < 0x0800) {
		/* three byte could be represented with less bytes */
		return UnicodeErrorOverlongCharacter
	} else if (unicodeLen == 2) && (d < 0x80) {
		/* two byte could be represented with less bytes */
		return UnicodeErrorOverlongCharacter
	}

	return unicodeLen
}
