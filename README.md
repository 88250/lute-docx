## 💡 简介

Lute DOCX 是一款将 Markdown 文本转换为 Word 文档 (.docx) 的小工具。通过 [Lute](https://github.com/88250/lute) 解析 Markdown 然后再通过 [unioffice](https://github.com/unidoc/unioffice) 生成 DOCX。

## ✨  特性

* 几乎支持所有 Markdown 语法元素
* 图片会通过地址自动拉取并渲染
* 支持封面配置

## 📸 截图

![sample](https://user-images.githubusercontent.com/873584/79592318-69a17100-810c-11ea-8c26-a6168e681325.png)

## ⚗ 用法

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

## 🐛 已知问题

* 没有代码高亮，代码块统一使用绿色渲染
* 没有渲染 Emoji
* 表格没有边框
* 表格单元格折行计算有问题
* 粗体、斜体需要字体本身支持

## 🏘️ 社区

* [讨论区](https://ld246.com/tag/lute)
* [报告问题](https://github.com/88250/lute-docx/issues/new)
* 欢迎关注 B3log 开源社区微信公众号 `B3log开源`  
  ![b3logos.png](https://b3logfile.com/file/2019/10/image-d3c00d78.png)

## 📄 开源协议

Lute DOCX 使用 [AGPLv3](https://www.gnu.org/licenses/agpl-3.0.txt) 开源协议。

## 🙏 鸣谢

* [对中文语境优化的 Markdown 引擎 Lute](https://ld246.com/article/1567047822949)
* [Go 实现的 Office 文档操作工具库](https://github.com/unidoc/unioffice)
