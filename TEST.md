# Test Report: @go/encoding

## 📋 测试概览
- **测试时间**: 2026-06-22
- **测试环境**: darwin/amd64
- **Go 版本**: 1.26.1

## ✅ 功能测试 (Functional Tests)
| 场景 | 状态 | 描述 |
| :--- | :--- | :--- |
| `TestHex` | PASS | Hex 编解码往返一致性，配合 `cast.As` 消除摩擦测试。 |
| `TestBase64` | PASS | Standard/URL Base64 编解码一致性，`Un...FromString` 补全测试。 |
| `TestWebEncoding` | PASS | URL 编解码、HTML 转义反转义往返测试。 |
| `TestSortJoin` | PASS | Map 与 Struct 的排序拼接测试，验证签名场景。 |
| `TestUtf8` | PASS | UTF-8 校验逻辑测试。 |
| `TestIntEncoder` | PASS | 包含正常编解码、FillInt 补齐、ExchangeInt 原地置换、HashInt 非线性单向散列与 FoldInt 模加折叠测试。 |
| `TestIntEncoderInstance` | PASS | 自定义 IntEncoder 实例的高性能编解码与原地置换测试。 |

## 🛡️ 鲁棒性防御 (Robustness)
- **消除摩擦 (Frictionless)**：废除 `Must` 系列 API，通过 `cast.As` 实现静默处理，保持业务逻辑简洁。
- **API 补全**：补全了所有 `FromString` 版本的解码函数，并提供 standard error 返回。
- **参数校验**：`NewIntEncoder` 对字符集重复、长度不足等构造错误进行了防御性校验。
- **折叠截断安全**：`FoldInt` 通过同余模加法在 Base62 字符集内折叠信息，彻底解决物理截断引起的局部哈希碰撞问题。

## ⚡ 性能基准 (Benchmarks)
| 函数 | 平均耗时 | 性能分析 |
| :--- | :--- | :--- |
| `HexEncode` | **209.7 ns/op** | 稳定且低开销。 |
| `Base64Encode` | **185.5 ns/op** | 吞吐量优异。 |
| `IntEncoder` | **47.04 ns/op** | 栈分配极致优化，零 GC 负担。 |
