package encoding

import "apigo.cc/go/jsmod"

func init() {
	jsmod.Register("encoding", map[string]any{
		"Base64":       Base64,
		"Base64Raw":    Base64Raw,
		"UnBase64":     jsUnBase64,
		"URLBase64":    URLBase64,
		"URLBase64Raw": URLBase64Raw,
		"UnURLBase64":  jsUnURLBase64,
		"Hex":          Hex,
		"UnHex":        jsUnHex,
		"URLEncode":    URLEncode,
		"UnURLEncode":  jsUnURLEncode,
		"HTMLEscape":   HTMLEscape,
		"HTMLUnescape": HTMLUnescape,
		"UTF8Valid":    UTF8Valid,
		"SortJoin":     SortJoin,
		"EncodeInt":    jsEncodeInt,
		"DecodeInt":    DecodeInt,
		"FillInt":      FillInt,
		"ExchangeInt":  ExchangeInt,
		"HashInt":      HashInt,
	})
}

func jsUnBase64(data any) ([]byte, error) {
	res, err := UnBase64(data)
	if err != nil {
		return nil, jsmod.MakeError(err)
	}
	return res, nil
}

func jsUnURLBase64(data any) ([]byte, error) {
	res, err := UnURLBase64(data)
	if err != nil {
		return nil, jsmod.MakeError(err)
	}
	return res, nil
}

func jsUnHex(data any) ([]byte, error) {
	res, err := UnHex(data)
	if err != nil {
		return nil, jsmod.MakeError(err)
	}
	return res, nil
}

func jsUnURLEncode(data any) ([]byte, error) {
	res, err := UnURLEncode(data)
	if err != nil {
		return nil, jsmod.MakeError(err)
	}
	return res, nil
}

func jsEncodeInt(u uint64) string {
	return string(EncodeInt(u))
}
