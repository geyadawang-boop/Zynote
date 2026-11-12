# 贡献者

按功能资产口径核算（14 项资产；每项 = 需求 0.35 + 契约 0.25 + 实现 0.40），快照 `main@8773dbb`。

| 贡献者 | 角色 | 需求 | 契约 | 实现 | 综合占比 |
| --- | --- | ---: | ---: | ---: | ---: |
| [@EveSunMaple](https://github.com/EveSunMaple) | 重构者：Go / Rust 后端、协议契约、发布 | 3.3 | 12.0 | 19.6 | **34.9%** |
| [@geyadawang-boop](https://github.com/geyadawang-boop) | BiliSum 作者、本仓库维护者 | 15.0 | 6.2 | 11.8 | **33.0%** |
| [@mikudeshuo](https://github.com/mikudeshuo) | meiji-listening-guide 0.1.85 作者 | 13.0 | 4.8 | 5.4 | **23.2%** |
| [@zhongmian085182019](https://github.com/zhongmian085182019) | 课表截图导入、教材 PDF 批注 | 2.9 | 1.5 | 1.8 | **6.2%** |
| [@yslasd](https://github.com/yslasd) | 研究模块：引用核对、检索渠道、PDFium | 0.8 | 0.5 | 1.4 | **2.7%** |

分项四舍五入，尾差 ±0.1。

## 项目沿革

1. **BiliSum**（geyadawang-boop）：B 站视频字幕与内容总结工具。功能完整版本未开源，本仓库只包含参与合并的开源部分。
2. **meiji-listening-guide 0.1.85**（mikudeshuo）：智云课堂字幕与课中标注，TypeScript/Electron。
3. 两个应用合并为一个桌面应用，合并阶段遗留构建报错与兼容问题。
4. EveSunMaple 接手重构：业务核心改为 Go（`backend/`），原生助手改为 Rust（`native/`），形成 Zynote。
5. 此后由上述成员通过 PR 持续修复与改进。

## 计分口径

```
每个功能资产 = 1 份资产
  需求与流程定义  0.35
  接口契约        0.25
  实现            0.40   ← 按继承系数 α 在原实现者与重写者之间分配
```

资产权重 = 终版有效代码量 45% + 难度 20% + 采纳度 15% + PR/合并历史 12% + git 工时 8%。

采用资产口径而非行数口径的原因：TypeScript 版本与 Go/Rust 版本是同一功能的两次实现，按行数会重复计分；早期工作（CI、三平台构建、测试、平台适配）行数少但耗时并不少。

## 口径敏感性

| 贡献者 | 主口径 | 14 项等权 | 仅实现 | 仅需求+契约 |
| --- | ---: | ---: | ---: | ---: |
| EveSunMaple | 34.9% | 34.7% | 49.0% | 25.5% |
| geyadawang-boop | 33.0% | 33.2% | 29.5% | 35.3% |
| mikudeshuo | 23.2% | 22.5% | 13.5% | 29.7% |
| zhongmian085182019 | 6.2% | 6.7% | 4.5% | 7.3% |
| yslasd | 2.7% | 2.9% | 3.5% | 2.2% |

- 只看实现：重构工作量集中在 EveSunMaple。
- 只看设计侧：两个原始应用定义了产品骨架，geyadawang-boop 与 mikudeshuo 合计约 65%。
- 等权与主口径基本一致，结果不依赖权重设计。

## 说明

- 这是自定义折算，不是 GitHub 官方统计。
- 资产归并与难度评分含人工判断，逐条证据与分数矩阵见 `docs/contributors-report.html`。
- 已合并 PR 共 18 个，逐条记录在 `pr-archive/`。
