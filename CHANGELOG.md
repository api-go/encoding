# Changelog: @go/encoding

## [v1.0.0] - 2026-04-22

### Added
- **基础编解码引擎**：提供基于 `[]byte` 的 Hex 和 Base64 (Standard/URL) 编解码支持。
- **静默 API (Must Series)**：新增 `MustUnXxx` 系列 API，自动屏蔽解码错误，简化业务处理逻辑。
- **Web 协议支持**：新增 URL 编解码与 HTML 转义/反转义接口。
- **高级整数编码**：移植并重构 `IntEncoder`，支持自定义进制、补齐填充 (`FillInt`)、位置置换 (`ExchangeInt`) 与 HMAC-SHA512 哈希校验。
- **健壮性校验**：新增 UTF-8 有效性检测。
