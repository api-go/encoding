package encoding

import (
	"encoding/base64"
	"encoding/hex"
	"html"
	"net/url"
	"unicode/utf8"
)

// Hex 将数据转换为 Hex 编码的字节切片
func Hex(data []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	return dst
}

// HexToString 将数据转换为 Hex 编码的字符串
func HexToString(data []byte) string {
	return hex.EncodeToString(data)
}

// MustUnHex 将 Hex 编码的字节切片解码，出错时返回空字节切片
func MustUnHex(data []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(data)))
	n, err := hex.Decode(dst, data)
	if err != nil {
		return []byte{}
	}
	return dst[:n]
}

// MustUnHexFromString 将 Hex 编码的字符串解码，出错时返回空字节切片
func MustUnHexFromString(data string) []byte {
	dst := make([]byte, hex.DecodedLen(len(data)))
	n, err := hex.Decode(dst, []byte(data))
	if err != nil {
		return []byte{}
	}
	return dst[:n]
}

// UnHex 将 Hex 编码的字节切片解码
func UnHex(data []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(data)))
	n, err := hex.Decode(dst, data)
	return dst[:n], err
}

// Base64 将数据转换为 Base64 编码的字节切片
func Base64(data []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(buf, data)
	return buf
}

// Base64ToString 将数据转换为 Base64 编码的字符串
func Base64ToString(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// UrlBase64 将数据转换为 URL 安全的 Base64 编码的字节切片
func UrlBase64(data []byte) []byte {
	buf := make([]byte, base64.URLEncoding.EncodedLen(len(data)))
	base64.URLEncoding.Encode(buf, data)
	return buf
}

// UrlBase64ToString 将数据转换为 URL 安全的 Base64 编码的字符串
func UrlBase64ToString(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// MustUnBase64 将 Base64 编码的字节切片解码，出错时返回空字节切片
func MustUnBase64(data []byte) []byte {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dbuf, data)
	if err != nil {
		return []byte{}
	}
	return dbuf[:n]
}

// MustUnBase64FromString 将 Base64 编码的字符串解码，出错时返回空字节切片
func MustUnBase64FromString(data string) []byte {
	dbuf, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return []byte{}
	}
	return dbuf
}

// MustUnUrlBase64 将 URL 安全的 Base64 编码的字节切片解码，出错时返回空字节切片
func MustUnUrlBase64(data []byte) []byte {
	dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(data)))
	n, err := base64.URLEncoding.Decode(dbuf, data)
	if err != nil {
		return []byte{}
	}
	return dbuf[:n]
}

// MustUnUrlBase64FromString 将 URL 安全的 Base64 编码的字符串解码，出错时返回空字节切片
func MustUnUrlBase64FromString(data string) []byte {
	dbuf, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return []byte{}
	}
	return dbuf
}

// UnBase64 将 Base64 编码的字节切片解码
func UnBase64(data []byte) ([]byte, error) {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dbuf, data)
	return dbuf[:n], err
}

// UnUrlBase64 将 URL 安全的 Base64 编码的字节切片解码
func UnUrlBase64(data []byte) ([]byte, error) {
	dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(data)))
	n, err := base64.URLEncoding.Decode(dbuf, data)
	return dbuf[:n], err
}

// UrlEncode 对数据进行 URL 编码
func UrlEncode(data []byte) string {
	return url.QueryEscape(string(data))
}

// MustUnUrlEncode 对字符串进行 URL 解码，出错时返回空字节切片
func MustUnUrlEncode(data string) []byte {
	res, err := url.QueryUnescape(data)
	if err != nil {
		return []byte{}
	}
	return []byte(res)
}

// HtmlEscape 对数据进行 HTML 转义
func HtmlEscape(data []byte) string {
	return html.EscapeString(string(data))
}

// MustUnHtmlEscape 对 HTML 字符串进行反转义
func MustUnHtmlEscape(data string) string {
	return html.UnescapeString(data)
}

// Utf8Valid 检查字节切片是否为有效的 UTF-8 编码
func Utf8Valid(data []byte) bool {
	return utf8.Valid(data)
}
