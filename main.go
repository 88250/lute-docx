// Lute DOCX - 一款将 Markdown 文本转换为 Word 文档 (.docx) 的小工具
// Copyright (c) 2020-present, b3log.org
//
// LianDi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/88250/gulu"
	"github.com/88250/lute/parse"
)

var logger *gulu.Logger

func init() {
	logger = gulu.Log.NewLogger(os.Stdout)
}

func main() {
	argMdPath := flag.String("mdPath", "D:/88250/lute-docx/sample.md", "待转换的 Markdown 文件路径")
	argSavePath := flag.String("savePath", "D:/88250/lute-docx/sample.docx", "转换后 DOCX 的保存路径")

	argRegularFontPath := flag.String("regularFontPath", "D:/88250/lute-docx/fonts/msyh.ttf", "正常字体文件路径")
	argBoldFontPath := flag.String("boldFontPath", "D:/88250/lute-docx/fonts/msyhb.ttf", "粗体字体文件路径")
	argItalicFontPath := flag.String("italicFontPath", "D:/88250/lute-docx/fonts/msyhl.ttf", "斜体字体文件路径")

	argCoverTitle := flag.String("coverTitle", "Lute DOCX - Markdown 生成 DOCX", "封面 - 标题")
	argCoverAuthor := flag.String("coverAuthor", "88250", "封面 - 作者")
	argCoverAuthorLink := flag.String("coverAuthorLink", "https://hacpai.com/member/88250", "封面 - 作者链接")
	argCoverLink := flag.String("coverLink", "https://github.com/88250/lute-docx", "封面 - 原文链接")
	argCoverSource := flag.String("coverSource", "GitHub", "封面 - 来源网站")
	argCoverSourceLink := flag.String("coverSourceLink", "https://github.com", "封面 - 来源网站链接")
	argCoverLicense := flag.String("coverLicense", "署名-相同方式共享 4.0 国际 (CC BY-SA 4.0)", "封面 - 文档许可协议")
	argCoverLicenseLink := flag.String("coverLicenseLink", "https://creativecommons.org/licenses/by-sa/4.0/", "封面 - 文档许可协议链接")
	argCoverLogoLink := flag.String("coverLogoLink", "https://static.b3log.org/images/brand/b3log-128.png", "封面 - 图标链接")
	argCoverLogoTitle := flag.String("coverLogoTitle", "B3log 开源", "封面 - 图标标题")
	argCoverLogoTitleLink := flag.String("coverLogoTitleLink", "https://b3log.org", "封面 - 图标标题链接")

	flag.Parse()

	mdPath := trimQuote(*argMdPath)
	savePath := trimQuote(*argSavePath)

	regularFontPath := trimQuote(*argRegularFontPath)
	boldFontPath := trimQuote(*argBoldFontPath)
	italicFontPath := trimQuote(*argItalicFontPath)

	coverTitle := trimQuote(*argCoverTitle)
	coverAuthorLabel := "　　作者："
	coverAuthor := trimQuote(*argCoverAuthor)
	coverAuthorLink := trimQuote(*argCoverAuthorLink)
	coverLinkLabel := "原文链接："
	coverLink := trimQuote(*argCoverLink)
	coverSourceLabel := "来源网站："
	coverSource := trimQuote(*argCoverSource)
	coverSourceLink := trimQuote(*argCoverSourceLink)
	coverLicenseLabel := "许可协议："
	coverLicense := trimQuote(*argCoverLicense)
	coverLicenseLink := trimQuote(*argCoverLicenseLink)
	coverLogoLink := trimQuote(*argCoverLogoLink)
	coverLogoTitle := trimQuote(*argCoverLogoTitle)
	coverLogoTitleLink := trimQuote(*argCoverLogoTitleLink)

	options := &parse.Options{
		GFMTable:            true,
		GFMTaskListItem:     true,
		GFMStrikethrough:    true,
		GFMAutoLink:         true,
		SoftBreak2HardBreak: true,
		Emoji:               true,
		Footnotes:           true,
	}
	options.AliasEmoji, options.EmojiAlias = parse.NewEmojis()

	markdown, err := ioutil.ReadFile(mdPath)
	if nil != err {
		logger.Fatal(err)
	}

	markdown = bytes.ReplaceAll(markdown, []byte("\t"), []byte("    "))
	for emojiUnicode, emojiAlias := range options.EmojiAlias {
		markdown = bytes.ReplaceAll(markdown, []byte(emojiUnicode), []byte(":"+emojiAlias+":"))
	}

	tree := parse.Parse("", markdown, options)
	renderer := NewDocxRenderer(tree, regularFontPath, boldFontPath, italicFontPath)
	renderer.Cover = &DocxCover{
		Title:         coverTitle,
		AuthorLabel:   coverAuthorLabel,
		Author:        coverAuthor,
		AuthorLink:    coverAuthorLink,
		LinkLabel:     coverLinkLabel,
		Link:          coverLink,
		SourceLabel:   coverSourceLabel,
		Source:        coverSource,
		SourceLink:    coverSourceLink,
		LicenseLabel:  coverLicenseLabel,
		License:       coverLicense,
		LicenseLink:   coverLicenseLink,
		LogoLink:      coverLogoLink,
		LogoTitle:     coverLogoTitle,
		LogoTitleLink: coverLogoTitleLink,
	}
	renderer.RenderCover()

	renderer.Render()
	renderer.Save(savePath)

	logger.Info("completed")
}

func trimQuote(str string) string {
	return strings.Trim(str, "\"'")
}
