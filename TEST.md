# Test Report: @go/encoding

## 📋 测试概览
- **测试时间**: 2026-04-22
- **测试环境**: darwin/amd64
- **Go 版本**: 1.25.0

## ✅ 功能测试 (Functional Tests)
| 场景 | 状态 | 描述 |
| :--- | :--- | :--- |
| `TestHex` | PASS | Hex 编解码往返一致性，Must 系列静默处理测试。 |
| `TestBase64` | PASS | Standard/URL Base64 编解码一致性，静默 API 测试。 |
| `TestWebEncoding` | PASS | URL 编解码、HTML 转义反转义往返测试。 |
| `TestUtf8` | PASS | UTF-8 校验逻辑测试。 |
| `TestIntEncoder` | PASS | 包含正常编解码、FillInt 补齐、ExchangeInt 置换、HashInt 确定性测试。 |

## 🛡️ 鲁棒性防御 (Robustness)
- **静默处理 (Quiet Mode)**：所有 `MustUnXxx` API 对非法数据均返回空切片 `[]byte{}` 或空字符串，有效防止业务逻辑中的非预期中断或 Panic。
- **参数校验**：`NewIntEncoder` 对字符集重复、长度不足等构造错误进行了防御性校验。

## ⚡ 性能基准 (Benchmarks)
| 函数 | 平均耗时 | 性能分析 |
| :--- | :--- | :--- |
| `HexEncode` | **44.85 ns/op** | 高效处理二进制数据。 |
| `Base64Encode` | **40.41 ns/op** | 吞吐量优异。 |
| `IntEncoder` | **44.18 ns/op** | 整数编码逻辑极简，开销极小。 |
