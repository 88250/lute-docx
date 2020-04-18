// Lute DOCX - 一款将 Markdown 文本转换为 Word 文档 (.docx) 的小工具
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
	renderer := NewDocxRenderer(tree)
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
