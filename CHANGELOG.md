# Changelog: @go/encoding

## [v1.3.2] - 2026-05-30

### Added
- **JSMOD 注册**：将核心编解码能力注册到 `jsmod`，支持在 JS 环境中调用 Base64, Hex, URL, HTML 及 SortJoin 等能力。

## [v1.1.1] - 2026-05-08

### Added
- **文档补全**：在 README 中补充了 `SortJoin` 接口的详细说明，该接口用于 Map 或 Struct 的排序拼接（签名场景）。

## [v1.1.0] - 2026-05-06

### Added
- **Base64 无填充支持**：新增 `Base64Raw`、`Base64RawToString`、`UrlBase64Raw`、`UrlBase64RawToString` 接口，支持生成不带填充符（`=`）的编码。

### Changed
- **智能解码升级**：升级了 `UnBase64` 与 `UnUrlBase64` 系列函数，通过 O(1) 零分配检测自动兼容“带填充”与“无填充”的输入数据，无需额外调用 Raw 解码接口。

## [v1.0.6] - 2026-05-06

### Changed
- **设计哲学对齐**：全面废除 `Must` 前缀函数，改为配合 `go/cast` 的 `As` 函数消除摩擦，提升代码语义化。
- **API 重命名**：`MustUnHtmlEscape` 重命名为 `HtmlUnescape`，使其符合标准 API 习惯且保持 frictionless。

### Added
- **API 补全**：新增 `UnHexFromString`、`UnBase64FromString`、`UnUrlBase64FromString` 与 `UnUrlEncode` 异步错误返回版本。
- **依赖对齐**：引入 `apigo.cc/go/cast` 依赖以支持消除摩擦能力。

## [v1.0.5] - 2026-05-01
- (同步版本号)

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
