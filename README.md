# 关于本项目

本项目完全由 AI 维护。代码源自 github.com/ssgo/u 的重构。

# @go/encoding

`@go/encoding` 是一个为“极简、安全、无摩擦”业务开发设计的编解码工具库。它统一了二进制数据与文本处理的 API 语义，通过与 `go/cast` 结合，极大降低了业务逻辑中的错误处理心智负担。

## 🎯 设计哲学

*   **API 原生直觉**：二进制操作基于 `[]byte`，文本表现类操作（如 HTML/URL）基于 `string`，与 Go 原生习惯保持高度一致。
*   **消除摩擦 (Frictionless)**：废除 `Must` 前缀函数，推荐配合 `cast.As` 使用以实现更优雅的静默处理，降低业务代码中无效的错误检查噪声。
*   **极致纯粹**：废除所有冗余的封装，强制数据链路层以 `[]byte` 形式流转，确保底层安全性与性能。

## 🛠 API Reference

### 基础编解码 (Hex/Base64)
- `func Hex(data []byte) []byte` / `func HexToString(data []byte) string`
- `func UnHex(data []byte) ([]byte, error)` / `func UnHexFromString(data string) ([]byte, error)`
- `func Base64(data []byte) []byte` / `func Base64ToString(data []byte) string`
- `func Base64Raw(data []byte) []byte` / `func Base64RawToString(data []byte) string` (无填充版本)
- `func UnBase64(data []byte) ([]byte, error)` / `func UnBase64FromString(data string) ([]byte, error)` (智能兼容填充与无填充)
- `func UrlBase64(data []byte) []byte` / `func UrlBase64ToString(data []byte) string`
- `func UrlBase64Raw(data []byte) []byte` / `func UrlBase64RawToString(data []byte) string` (无填充版本)
- `func UnUrlBase64(data []byte) ([]byte, error)` / `func UnUrlBase64FromString(data string) ([]byte, error)` (智能兼容填充与无填充)

### Web 编码 (URL/HTML)
- `func UrlEncode(data []byte) string`
- `func UnUrlEncode(data string) ([]byte, error)`
- `func HtmlEscape(data []byte) string`
- `func HtmlUnescape(data string) string`
- `func Utf8Valid(data []byte) bool`

### 签名工具 (Signature Utils)
- `func SortJoin(v any, separator, connector string, urlEncode bool) string`
  将 Map 或 Struct 转换为排序并拼接后的字符串。常用于生成签名原串。支持自动 URL 编码。

### 整数与自定义进制 (IntEncoder)
- `func EncodeInt(u uint64) []byte`
- `func AppendInt(buf []byte, u uint64) []byte`
- `func DecodeInt(buf []byte) uint64`
- `func FillInt(buf []byte, length int) []byte`
- `func ExchangeInt(buf []byte) []byte`
- `func HashInt(buf []byte) []byte`
- `func FoldInt(buf []byte, length int) []byte`

## 📦 安装

```bash
go get apigo.cc/go/encoding
```

## 💡 示例 (配合 cast.As 消除摩擦)

```go
import (
    "apigo.cc/go/encoding"
    "apigo.cc/go/cast"
)

// 配合 cast.As 替代原有的 MustUnHex 系列
data := cast.As(encoding.UnHexFromString("68656c6c6f"))
```
