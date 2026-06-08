package encoding

import (
	"encoding/base64"
	"encoding/hex"
	"html"
	"net/url"
	"sort"
	"strings"
	"unicode/utf8"

	"apigo.cc/go/cast"
)

// Hex 将数据转换为 Hex 编码的字符串
func Hex(data any) string {
	b := cast.To[[]byte](data)
	return hex.EncodeToString(b)
}

// UnHex 将 Hex 编码的数据解码为字节切片
func UnHex(data any) ([]byte, error) {
	s := cast.String(data)
	return hex.DecodeString(s)
}

// Base64 将数据转换为 Base64 编码的字符串
func Base64(data any) string {
	b := cast.To[[]byte](data)
	return base64.StdEncoding.EncodeToString(b)
}

// UnBase64 将 Base64 编码的数据解码为字节切片（自动兼容有无填充）
func UnBase64(data any) ([]byte, error) {
	s := cast.String(data)
	if len(s) > 0 && s[len(s)-1] == '=' {
		return base64.StdEncoding.DecodeString(s)
	}
	return base64.RawStdEncoding.DecodeString(s)
}

// Base64Raw 将数据转换为无填充的 Base64 编码的字符串
func Base64Raw(data any) string {
	b := cast.To[[]byte](data)
	return base64.RawStdEncoding.EncodeToString(b)
}

// URLBase64 将数据转换为 URL 安全的 Base64 编码的字符串
func URLBase64(data any) string {
	b := cast.To[[]byte](data)
	return base64.URLEncoding.EncodeToString(b)
}

// UnURLBase64 将 URL 安全的 Base64 编码的数据解码为字节切片（自动兼容有无填充）
func UnURLBase64(data any) ([]byte, error) {
	s := cast.String(data)
	if len(s) > 0 && s[len(s)-1] == '=' {
		return base64.URLEncoding.DecodeString(s)
	}
	return base64.RawURLEncoding.DecodeString(s)
}

// URLBase64Raw 将数据转换为 URL 安全且无填充的 Base64 编码的字符串
func URLBase64Raw(data any) string {
	b := cast.To[[]byte](data)
	return base64.RawURLEncoding.EncodeToString(b)
}

// URLEncode 对数据进行 URL 编码
func URLEncode(data any) string {
	return url.QueryEscape(cast.String(data))
}

// UnURLEncode 对字符串进行 URL 解码
func UnURLEncode(data any) ([]byte, error) {
	res, err := url.QueryUnescape(cast.String(data))
	if err != nil {
		return nil, err
	}
	return []byte(res), nil
}

// HTMLEscape 对数据进行 HTML 转义
func HTMLEscape(data any) string {
	return html.EscapeString(cast.String(data))
}

// HTMLUnescape 对 HTML 字符串进行反转义
func HTMLUnescape(data any) string {
	return html.UnescapeString(cast.String(data))
}

// UTF8Valid 检查数据是否为有效的 UTF-8 编码
func UTF8Valid(data any) bool {
	return utf8.Valid(cast.To[[]byte](data))
}

// SortJoin 将 Map 或 Struct 转换为排序并拼接后的字符串 (常用于签名)
func SortJoin(v any, separator, connector string, urlEncode bool) string {
	var m map[string]any
	if vm, ok := v.(map[string]any); ok {
		m = vm
	} else if v != nil {
		cast.Convert(&m, v)
	}

	if len(m) == 0 {
		return ""
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for i, k := range keys {
		if i > 0 {
			builder.WriteString(separator)
		}
		val := cast.String(m[k])
		if urlEncode {
			builder.WriteString(url.QueryEscape(k))
			builder.WriteString(connector)
			builder.WriteString(url.QueryEscape(val))
		} else {
			builder.WriteString(k)
			builder.WriteString(connector)
			builder.WriteString(val)
		}
	}
	return builder.String()
}
