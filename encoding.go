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

// UnHex 将 Hex 编码的字节切片解码
func UnHex(data []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(data)))
	n, err := hex.Decode(dst, data)
	return dst[:n], err
}

// UnHexFromString 将 Hex 编码的字符串解码
func UnHexFromString(data string) ([]byte, error) {
	return hex.DecodeString(data)
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

// UnBase64 将 Base64 编码的字节切片解码
func UnBase64(data []byte) ([]byte, error) {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dbuf, data)
	return dbuf[:n], err
}

// UnBase64FromString 将 Base64 编码的字符串解码
func UnBase64FromString(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// UnUrlBase64 将 URL 安全的 Base64 编码的字节切片解码
func UnUrlBase64(data []byte) ([]byte, error) {
	dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(data)))
	n, err := base64.URLEncoding.Decode(dbuf, data)
	return dbuf[:n], err
}

// UnUrlBase64FromString 将 URL 安全的 Base64 编码的字符串解码
func UnUrlBase64FromString(data string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(data)
}

// UrlEncode 对数据进行 URL 编码
func UrlEncode(data []byte) string {
	return url.QueryEscape(string(data))
}

// UnUrlEncode 对字符串进行 URL 解码
func UnUrlEncode(data string) ([]byte, error) {
	res, err := url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}
	return []byte(res), nil
}

// HtmlEscape 对数据进行 HTML 转义
func HtmlEscape(data []byte) string {
	return html.EscapeString(string(data))
}

// HtmlUnescape 对 HTML 字符串进行反转义
func HtmlUnescape(data string) string {
	return html.UnescapeString(data)
}

// Utf8Valid 检查字节切片是否为有效的 UTF-8 编码
func Utf8Valid(data []byte) bool {
	return utf8.Valid(data)
}
