package main

import (
	"fmt"
	"testing"
)

func TestNewFereBook(t *testing.T) {
	url := "http://reader.epubee.com/books/mobile/ba/baddff340b4defb93fe29d15889bab10/text00000.html"
	cover := "http://files.epubee.com/getCover.ashx?fpath=c6/c6c17b6f11f12e20cdcf42d6078036d5_s.jpg"
	book := NewFereBook(url, cover)
	book.Run()
	fmt.Printf("%+v\n", "start")
}

func TestName(t *testing.T) {

}
