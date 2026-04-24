# AI Coding Context: @go/encoding

本索引供 AI 模型理解 `@go/encoding` 的设计规范，以生成符合本项目“语义优先、安全第一”哲学的代码。

## 🤖 AI 行为准则

1.  **区分数据与文本**：二进制数据 (Hex/Base64/Int) 必须基于 `[]byte`；文本格式 (URL/HTML) 优先使用 `string`。
2.  **默认开启静默模式**：在业务逻辑中，优先推荐使用 `MustUnXxx` 系列 API。一旦解码失败，它们返回零值，无需额外 `if err != nil`，保持逻辑流平滑。
3.  **IntEncoder 集成**：所有整数进制编码、填充、混淆、哈希操作均通过 `IntEncoder` 实例或其导出的封装函数进行。
4.  **安全链路**：优先使用基于字节切片的 API，减少不必要的内存分配与 string 转换。

## 🛠 API Reference

### 基础编解码 (Hex/Base64)
- `func Hex(data []byte) []byte`：将数据转换为 Hex 编码的字节切片。
- `func HexToString(data []byte) string`：将数据转换为 Hex 编码的字符串。
- `func MustUnHex(data []byte) []byte`：解码 Hex，出错时返回空字节切片。
- `func MustUnHexFromString(data string) []byte`：解码 Hex 字符串，出错时返回空字节切片。
- `func UnHex(data []byte) ([]byte, error)`：解码 Hex。
- `func Base64(data []byte) []byte`：将数据转换为 Base64 编码的字节切片。
- `func Base64ToString(data []byte) string`：将数据转换为 Base64 编码的字符串。
- `func UrlBase64(data []byte) []byte`：将数据转换为 URL 安全的 Base64 编码的字节切片。
- `func UrlBase64ToString(data []byte) string`：将数据转换为 URL 安全的 Base64 编码的字符串。
- `func MustUnBase64(data []byte) []byte`：解码 Base64，出错时返回空字节切片。
- `func MustUnBase64FromString(data string) []byte`：解码 Base64 字符串，出错时返回空字节切片。
- `func MustUnUrlBase64(data []byte) []byte`：解码 URL Base64，出错时返回空字节切片。
- `func MustUnUrlBase64FromString(data string) []byte`：解码 URL Base64 字符串，出错时返回空字节切片。
- `func UnBase64(data []byte) ([]byte, error)`：解码 Base64。
- `func UnUrlBase64(data []byte) ([]byte, error)`：解码 URL Base64。

### Web 编码 (URL/HTML)
- `func UrlEncode(data []byte) string`：URL 编码。
- `func MustUnUrlEncode(data string) []byte`：URL 解码，出错时返回空字节切片。
- `func HtmlEscape(data []byte) string`：HTML 转义。
- `func MustUnHtmlEscape(data string) string`：HTML 反转义。
- `func Utf8Valid(data []byte) bool`：校验 UTF-8。

### 整数编码 (IntEncoder)
- `func NewIntEncoder(digits string, radix uint8) (*IntEncoder, error)`：创建自定义编码器。
- `func EncodeInt(u uint64) []byte`：默认编码整数。
- `func AppendInt(buf []byte, u uint64) []byte`：默认追加整数。
- `func DecodeInt(buf []byte) uint64`：默认解码整数。
- `func ExchangeInt(buf []byte) []byte`：默认置换混淆。
- `func HashInt(data []byte, key []byte) []byte`：默认 HMAC-SHA512 哈希。
- `func (enc *IntEncoder) EncodeInt(u uint64) []byte`
- `func (enc *IntEncoder) AppendInt(buf []byte, u uint64) []byte`
- `func (enc *IntEncoder) FillInt(buf []byte, length int) []byte`：填充至指定长度。
- `func (enc *IntEncoder) ExchangeInt(buf []byte) []byte`
- `func (enc *IntEncoder) HashInt(data []byte, key []byte) []byte`
- `func (enc *IntEncoder) DecodeInt(buf []byte) uint64`

## 🧩 典型模式 (Best Practices)

*   **✅ 推荐：业务逻辑流 (Silence Mode)**:
    ```go
    // 解码失败即视为空，业务无需中断
    raw := encoding.MustUnHexFromString(config.Data)
    process(raw)
    ```

*   **✅ 推荐：整数编解码 (IntEncoder)**:
    ```go
    // 使用默认 62 进制编码器，支持追加与混淆
    buf := encoding.EncodeInt(userID)
    buf = encoding.DefaultIntEncoder.ExchangeInt(buf)
    ```
