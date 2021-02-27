// 从 ferebook 下载的注意事项

// 1. 封面的下载
//	curl 'http://files.epubee.com/getCover.ashx?fpath=c6/c6c17b6f11f12e20cdcf42d6078036d5_s.jpg' \
//  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8' \
//  -o 1.jpg

// 必须带上 Accept_language 才能下载

// 推送下载 awz3 的版本,格式较好
// 内容下載格式為 : http://reader.epubee.com/books/mobile/ba/baddff340b4defb93fe29d15889bab10/text00000.html
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Content struct {
	Url     string
	title   string
	Content []byte
}

type Ferebook struct {
	baseUrl string

	// 章节
	chap []*Content

	// 封面 url
	cover string

	client *http.Client
	parse  parse
}

func NewFereBook(url string, cover string) *Ferebook {
	url = strings.TrimSpace(url)

	f := &Ferebook{}
	if strings.HasSuffix(url, ".html") {
		f.baseUrl = url[:len(url)-14]
	} else {
		f.baseUrl = url
	}

	f.cover = cover

	f.client = &http.Client{}
	f.parse = NewParseHtml()

	return f
}

func (f *Ferebook) Run() {

	if f.cover != "" {
		go f.getCover()
	}

	f.readUrl()

}

func (f *Ferebook) readUrl() {

	for i := 0; i < 200; i++ {
		url := fmt.Sprintf("%vtext%05d.html", f.baseUrl, i)
		fmt.Printf("获取 %s 内容\n ", url)

		resp, err := f.client.Get(url)
		fmt.Printf("%+v\n", resp)
		if err != nil {
			log.Print(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			log.Print("状态码不正确")
			_ = resp.Body.Close()
			return
		}

		htmlCon, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		if len(htmlCon) == 0 {
			fmt.Printf("%+v\n", "获取内容为0")
			_ = resp.Body.Close()
			return
		}

		c := &Content{
			Url: url,
		}

		f.parse.Run(bytes.NewReader(htmlCon))
		title := f.parse.ParseTitle()
		content := f.parse.ParseContent()

		c.title = title

		c.Content = []byte(content)

		f.chap = append(f.chap, c)
	}

}

func (f *Ferebook) getCover() {
	fmt.Printf("开启获取封面信息\n")
	request, err := http.NewRequest("GET", f.cover, nil)
	if err != nil {
		panic(err)
	}
	request.Header = make(http.Header)
	request.Header["Accept-Language"] = []string{"zh-CN,zh;q=0.9,en;q=0.8"}

	resp, err := f.client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if !checkDir("data") {
		err := os.Mkdir("data", os.ModeDir)
		if err != nil {
			panic(err)
		}
	}

	img, err := os.Create("data/cover.jpg")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(img, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("成功获取封面信息\n")
}

func checkDir(dir string) bool {
	s, err := os.Stat(dir)
	if err != nil {

		return false
	}

	return s.IsDir()
}
