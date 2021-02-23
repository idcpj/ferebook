package main

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"io"
	"strings"
)

type parse interface {
	Run(io.Reader)
	ParseTitle() string
	ParseContent() string
}

type ParseHtml struct {
	title   string
	Content string

	doc *goquery.Document
}

func NewParseHtml() *ParseHtml {
	return &ParseHtml{
		doc: &goquery.Document{},
	}
}

func (p *ParseHtml) Run(reader io.Reader) {
	node, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	p.doc = goquery.NewDocumentFromNode(node)
}

func (p *ParseHtml) ParseTitle() string {
	ret, err := p.doc.Find("title").Html()
	if err != nil {
		panic(err)
	}
	return ret

}

func (p *ParseHtml) ParseContent() string {
	text, err := p.doc.Find(".readercontent-inner").Html()
	if err != nil {
		panic(err)
	}
	replacer := strings.NewReplacer("上一章", "下一章")
	return replacer.Replace(text)
}
