package main

import (
	"flag"
	"fmt"
	"github.com/766b/mobi"
	"log"
	"os"
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

	log.SetFlags(log.Lshortfile)
	//name="童年的终结"
	//url="http://reader.epubee.com/books/mobile/bb/bbca714c0c8f3140af2e8279b01ab22c/text00000.html"
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
	if cover != "" {
		m.AddCover("data/cover.jpg", "data/cover.jpg")
	}

	// Meta data
	m.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	m.NewExthRecord(mobi.EXTH_AUTHOR, "Book Author Name")
	m.NewExthRecord(mobi.EXTH_TITLE, name)
	// See exth.go for additional EXTH record IDs
	if len(book.chap) == 0 {
		fmt.Printf("%+v\n", "获取长度为0")
		return
	}
	for _, v := range book.chap {
		m.NewChapter(v.title, v.Content)
	}

	//m.NewChapter("asdasd",[]byte("asdasd"))
	// Output MOBI File
	m.Write()
	fmt.Printf("%+v\n", "完成")
}
