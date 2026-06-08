package encoding

import "apigo.cc/go/jsmod"

func init() {
	jsmod.Register("encoding", map[string]any{
		"Base64":                Base64,
		"Base64Raw":             Base64Raw,
		"UnBase64":              UnBase64,
		"URLBase64":             URLBase64,
		"URLBase64Raw":          URLBase64Raw,
		"UnURLBase64":           UnURLBase64,
		"Hex":                   Hex,
		"UnHex":                 UnHex,
		"URLEncode":             URLEncode,
		"UnURLEncode":           UnURLEncode,
		"HTMLEscape":            HTMLEscape,
		"HTMLUnescape":          HTMLUnescape,
		"UTF8Valid":             UTF8Valid,
		"SortJoin":              SortJoin,
		"EncodeInt": func(u uint64) string {
			return string(EncodeInt(u))
		},
		"DecodeInt":             DecodeInt,
		"FillInt":               FillInt,
		"ExchangeInt":           ExchangeInt,
		"HashInt":               HashInt,
	})
}
