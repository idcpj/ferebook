package main

import (
	"flag"
	"github.com/766b/mobi"
	"os"
	"os/signal"
)

var (
	name  string
	url   string
	cover string
)

func init() {
	flag.StringVar(&name, "name", "", "")
	flag.StringVar(&url, "url", "", "")
	flag.StringVar(&cover, "cover", "", "")
}

func main() {
	//name="希腊古典神话"
	//url="http://reader.epubee.com/books/mobile/ba/baddff340b4defb93fe29d15889bab10/text00000.html"
	//cover="http://files.epubee.com/getCover.ashx?fpath=c6/c6c17b6f11f12e20cdcf42d6078036d5_s.jpg"
	//name="童年的终结"
	//url="http://reader.epubee.com/books/mobile/ed/edf00b57d6425a73bee490de44942a7e/"
	//cover="http://files.epubee.com/getCover.ashx?fpath=ea/eaf6561235ff5333b4f70f8cbf93781b_s.jpg"

	flag.Parse()

	if name == "" ||
		url == "" {

		flag.PrintDefaults()
		os.Exit(0)
	}

	book := NewFereBook(url, cover)
	book.Run()

	m, err := mobi.NewWriter(name + ".mobi")
	if err != nil {
		panic(err)
	}

	m.Title(name)
	m.Compression(mobi.CompressionNone)

	// Add cover image
	m.AddCover("data/cover.jpg", "data/cover.jpg")

	// Meta data
	m.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	m.NewExthRecord(mobi.EXTH_AUTHOR, "Book Author Name")
	m.NewExthRecord(mobi.EXTH_TITLE, name)
	// See exth.go for additional EXTH record IDs

	for _, v := range book.chap {
		m.NewChapter(v.title, v.Content)
	}

	// Output MOBI File
	m.Write()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
	<-s
}
