package encoding

import "apigo.cc/go/jsmod"

func init() {
	jsmod.Register("encoding", map[string]any{
		"base64":                Base64,
		"base64ToString":        Base64ToString,
		"base64Raw":             Base64Raw,
		"base64RawToString":     Base64RawToString,
		"unBase64":              UnBase64,
		"unBase64FromString":    UnBase64FromString,
		"urlBase64":             UrlBase64,
		"urlBase64ToString":     UrlBase64ToString,
		"urlBase64Raw":          UrlBase64Raw,
		"urlBase64RawToString":  UrlBase64RawToString,
		"unUrlBase64":           UnUrlBase64,
		"unUrlBase64FromString": UnUrlBase64FromString,
		"hex":                   Hex,
		"hexToString":           HexToString,
		"unHex":                 UnHex,
		"unHexFromString":       UnHexFromString,
		"urlEncode":             UrlEncode,
		"unUrlEncode":           UnUrlEncode,
		"htmlEscape":            HtmlEscape,
		"htmlUnescape":          HtmlUnescape,
		"utf8Valid":             Utf8Valid,
		"sortJoin":              SortJoin,
		"encodeInt":             EncodeInt,
		"decodeInt":             DecodeInt,
		"fillInt":               FillInt,
		"exchangeInt":           ExchangeInt,
		"hashInt":               HashInt,
	})
}
