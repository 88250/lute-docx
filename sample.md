## 简介

Lute DOCX 是一款将 Markdown 文本转换为 Word 文档 (.docx) 的小工具。通过 [Lute](https://github.com/88250/lute) 解析 Markdown 然后再通过 [unioffice](https://github.com/unidoc/unioffice) 生成 DOCX。

## 特性

* 几乎支持所有 Markdown 语法元素
* 图片会通过地址自动拉取并渲染
* 支持封面配置

## 用法

命令行参数说明：

* `--mdPath`：待转换的 Markdown 文件路径
* `--savePath`：转换后 DOCX 的保存路径
* `--coverTitle`：封面 - 标题
* `--coverAuthor`：封面 - 作者
* `--coverAuthorLink`：封面 - 作者链接
* `--coverLink`：封面 - 原文链接
* `--coverSource`：封面 - 来源网站
* `--coverSourceLink`：封面 - 来源网站链接
* `--coverLicense`：封面 - 文档许可协议
* `--coverLicenseLink`：封面 - 文档许可协议链接
* `--coverLogoLink`：封面 - 图标链接
* `--coverLogoTitle`：封面 - 图标标题
* `--coverLogoTitleLink`：封面 - 图标标题链接

## 已知问题

* 没有代码高亮，代码块统一使用绿色渲染
* 没有渲染 Emoji
* 表格没有边框
* 表格单元格折行计算有问题

## 社区

* [讨论区](https://hacpai.com/tag/lute)
* [报告问题](https://github.com/88250/lute-docx/issues/new)
* 欢迎关注 B3log 开源社区微信公众号 `B3log开源`

  ![b3logos.png](https://img.hacpai.com/file/2019/10/image-d3c00d78.png)

## 开源协议

Lute DOCX 使用 [木兰宽松许可证, 第2版](http://license.coscl.org.cn/MulanPSL2) 开源协议。

## 鸣谢

* [对中文语境优化的 Markdown 引擎 Lute](https://hacpai.com/article/1567047822949)
* [Go 实现的 Office 文档操作工具库](https://github.com/unidoc/unioffice)
