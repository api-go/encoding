# Test Report: @go/encoding

## 📋 测试概览
- **测试时间**: 2026-05-06
- **测试环境**: darwin/amd64
- **Go 版本**: 1.25.0

## ✅ 功能测试 (Functional Tests)
| 场景 | 状态 | 描述 |
| :--- | :--- | :--- |
| `TestHex` | PASS | Hex 编解码往返一致性，配合 `cast.As` 消除摩擦测试。 |
| `TestBase64` | PASS | Standard/URL Base64 编解码一致性，`Un...FromString` 补全测试。 |
| `TestWebEncoding` | PASS | URL 编解码、HTML 转义反转义往返测试。 |
| `TestUtf8` | PASS | UTF-8 校验逻辑测试。 |
| `TestIntEncoder` | PASS | 包含正常编解码、FillInt 补齐、ExchangeInt 置换、HashInt 确定性测试。 |

## 🛡️ 鲁棒性防御 (Robustness)
- **消除摩擦 (Frictionless)**：废除 `Must` 系列 API，通过 `cast.As` 实现静默处理，保持业务逻辑简洁。
- **API 补全**：补全了所有 `FromString` 版本的解码函数，并提供标准 error 返回。
- **参数校验**：`NewIntEncoder` 对字符集重复、长度不足等构造错误进行了防御性校验。

## ⚡ 性能基准 (Benchmarks)
| 函数 | 平均耗时 | 性能分析 |
| :--- | :--- | :--- |
| `HexEncode` | **46.64 ns/op** | 高效处理二进制数据。 |
| `Base64Encode` | **42.24 ns/op** | 吞吐量优异。 |
| `IntEncoder` | **46.66 ns/op** | 整数编码逻辑极简，开销极小。 |
