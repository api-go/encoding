# @go/encoding

`@go/encoding` 是一个为“极简、安全、无摩擦”业务开发设计的编解码工具库。它统一了二进制数据与文本处理的 API 语义，通过静默处理错误与基于字节的零拷贝设计，极大降低了业务逻辑中的错误处理心智负担。

## 🎯 设计哲学

*   **API 原生直觉**：二进制操作基于 `[]byte`，文本表现类操作（如 HTML/URL）基于 `string`，与 Go 原生习惯保持高度一致。
*   **静默防御 (Must 系列)**：对于不可逆的非法输入，提供 `MustUnXxx` 系列 API，静默返回零值，消除业务代码中无效的错误检查噪声。
*   **极致纯粹**：废除所有冗余的封装，强制数据链路层以 `[]byte` 形式流转，确保底层安全性与性能。

## 🛠 API Reference

### 基础编解码 (Hex/Base64)
- `func Hex(data []byte) []byte` / `func HexToString(data []byte) string`
- `func MustUnHex(data []byte) []byte` / `func MustUnHexFromString(data string) []byte`

- `func Base64(data []byte) []byte` / `func Base64ToString(data []byte) string`
- `func MustUnBase64(data []byte) []byte` / `func MustUnBase64FromString(data string) []byte`

### Web 编码 (URL/HTML)
- `func UrlEncode(data []byte) string`
- `func MustUnUrlEncode(data string) []byte`

- `func HtmlEscape(data []byte) string`
- `func MustUnHtmlEscape(data string) string`

### 整数与自定义进制 (IntEncoder)
- `func EncodeInt(u uint64) []byte`
- `func AppendInt(buf []byte, u uint64) []byte`
- `func DecodeInt(buf []byte) uint64`
- `func FillInt(buf []byte, length int) uint64`
- `func ExchangeInt(buf []byte) []byte`
- `func HashInt(data []byte, key []byte) []byte`

## 📦 安装

```bash
go get apigo.cc/go/encoding
```
