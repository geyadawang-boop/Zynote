# Zynote

课堂学习桌面应用：智云课堂字幕、教师重点标注与课后练习。macOS 优先的 Go 后端重写版本。

## 项目沿革

Zynote 由两个独立项目合并、重构而来：

| 阶段 | 内容 |
| --- | --- |
| 起点 A | **BiliSum** —— [@geyadawang-boop](https://github.com/geyadawang-boop) 开发的 B 站视频字幕与内容总结工具；功能完整版本未开源 |
| 起点 B | **meiji-listening-guide** 0.1.85 —— 另一个独立项目，智云课堂字幕与课中标注工具（TypeScript/Electron），见 [mikudeshuo/meiji-listening-guide](https://github.com/mikudeshuo/meiji-listening-guide) |
| 合并 | 两个项目合并为一体化桌面应用，但合并阶段遗留了构建报错与若干兼容问题 |
| 重构 | [@EveSunMaple](https://github.com/EveSunMaple) 接手重构：业务核心改为 Go（`backend/`），原生助手改为 Rust（`native/`），形成 Zynote 并在其仓库公开发布 |
| 后续 | 由 [@geyadawang-boop](https://github.com/geyadawang-boop)、yslasd、zhongmian、mikudeshuo 等通过 PR 持续修复与改进 |

> 本仓库是 Zynote 的独立发布副本，遵循上游 AGPL-3.0-only 许可。版权归各阶段作者及贡献者所有，详见 `NOTICE`。

## 架构

| 层 | 技术 | 职责 |
| --- | --- | --- |
| 渲染层 | React 19 + HeroUI v3 + Tailwind 4 | 界面，不变 |
| Electron 壳 | TypeScript | 窗口、对话框、session、系统密钥库桥、自动更新 |
| 业务核心 | Go（`backend/`） | 数据库、课程、媒体、知识库、IM 编排、任务 |
| 原生助手 | Rust（`native/`） | PDF、ASR/向量、YouTube、微信本地读取 |

进程间协议：stdio 上的 JSON-RPC 2.0（换行分隔），stdout 只走协议。

## 开发

需要 Node.js 24+、pnpm、Go 1.25+、Rust（cargo）。

```sh
pnpm install
pnpm backend:build
pnpm dev
```

## 检查

```sh
pnpm verify        # tsc + biome
pnpm backend:test  # go test
pnpm native:test   # cargo test
pnpm smoke:go-core # Go core 冒烟（不启动 Electron）
pnpm update:smoke --app <lite app> # 本地更新源检查（macOS lite）
pnpm e2e:packaged --app <app>     # 打包版核心 IPC/DB E2E
pnpm perf:smoke --app <app>       # 冷启动/进程内存基线
pnpm migration:drill --app <app>  # 旧库备份/路径迁移/备份回读（macOS）
pnpm build
```

## 发布与自动更新

- `pnpm release lite --platform=mac --prepared` 生成安装包、`latest-mac.yml`、blockmap；lite 使用 `latest` 通道。
- Windows `full` 使用 `full` 通道；portable 目录版不生成更新元数据，也不检查更新。
- GitHub Actions 的 `release.yml` 默认只上传构建工件；手动勾选 `publish` 后会创建/更新 `v<version>` GitHub Release，并上传安装包与更新元数据。
- 生产自动更新必须使用 Developer ID 签名并公证；未签名的本地构建只能用于开发验证。

## 发布文档

见 `docs/RELEASE.md`；Windows IM 迁移见 `docs/WINDOWS_IM_HANDOFF.md`。

## 贡献者

- [@geyadawang-boop](https://github.com/geyadawang-boop)：BiliSum 作者，本仓库维护者
- mikudeshuo：meiji-listening-guide 作者
- [@EveSunMaple](https://github.com/EveSunMaple)（SunMaple）：Zynote 重构与上游维护者，拥有本仓库写权限
- yslasd、zhongmian：PR 贡献

## 许可

AGPL-3.0-only；第三方归属见 `NOTICE`、`docs/legal/THIRD_PARTY_NOTICES.md`。
