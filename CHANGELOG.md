# Changelog: @go/encoding

## [v1.0.4] - 2026-05-01

### Changed
- **性能优化**：优化了 Hex 编解码函数，通过直接使用 `hex.Decode` 避免了不必要的 `string` 转换和内存分配。
- **鲁棒性增强**：重构了 `NewIntEncoder`，强化了对字符集的校验，明确支持单字节字符集并提升了初始化安全性。
- **文档修正**：修正了 `FillInt` 的描述（从随机填充改为循环序列填充），并同步更新了 README 中的返回类型描述。

### Added
- **API 补全**：新增包级别的 `FillInt` 函数，使 `IntEncoder` 的所有核心能力均可通过包级 API 访问。

## [v1.0.0] - 2026-04-22

### Added
- **基础编解码引擎**：提供基于 `[]byte` 的 Hex 和 Base64 (Standard/URL) 编解码支持。
- **静默 API (Must Series)**：新增 `MustUnXxx` 系列 API，自动屏蔽解码错误，简化业务处理逻辑。
- **Web 协议支持**：新增 URL 编解码与 HTML 转义/反转义接口。
- **高级整数编码**：移植并重构 `IntEncoder`，支持自定义进制、补齐填充 (`FillInt`)、位置置换 (`ExchangeInt`) 与 HMAC-SHA512 哈希校验。
- **健壮性校验**：新增 UTF-8 有效性检测。
