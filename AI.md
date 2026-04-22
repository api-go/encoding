# AI Coding Context: @go/encoding

本索引供 AI 模型理解 `@go/encoding` 的设计规范，以生成符合本项目“语义优先、安全第一”哲学的代码。

## 🤖 AI 行为准则

1.  **区分数据与文本**：二进制数据 (Hex/Base64/Int) 必须基于 `[]byte`；文本格式 (URL/HTML) 优先使用 `string`。
2.  **默认开启静默模式**：在业务逻辑中，优先推荐使用 `MustUnXxx` 系列 API。一旦解码失败，它们返回零值，无需额外 `if err != nil`，保持逻辑流平滑。
3.  **IntEncoder 集成**：所有整数进制编码、填充、混淆、哈希操作均通过 `IntEncoder` 实例或其导出的封装函数进行。
4.  **安全链路**：优先使用基于字节切片的 API，减少不必要的内存分配与 string 转换。

## 🛠 关键 API 设计约定

| 类型 | 语义 API | 静默 API (Quiet) |
| :--- | :--- | :--- |
| **Hex** | `Hex/UnHex` | `MustUnHex/MustUnHexFromString` |
| **Base64** | `Base64/UnBase64` | `MustUnBase64/MustUnBase64FromString` |
| **IntEncoder** | `EncodeInt/AppendInt/FillInt/ExchangeIntInt/DecodeInt/HashInt` | - |
| **Web** | `UrlEncode/HtmlEscape` | `MustUnUrlEncode/MustUnHtmlEscape` |

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
