package biancheng

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

func Start() {
	url := "http://c.biancheng.net/view/%v.html"

	for i := 4871; i < 10000; i++ {
		url_f := fmt.Sprintf(url, i)
		fmt.Printf("Start -> url: %v\n", url_f)
		GetHTML(url_f, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func GetHTML(url string, i int) {
	c := colly.NewCollector()
	c.SetProxy("http://152.136.62.181:9999")
	c.OnHTML("title", func(h *colly.HTMLElement) {
		title := h.DOM.Filter("title").Text()
		fmt.Printf("title: %v\n", title)
		data := string(h.Response.Body)
		filePath := "/Users/lucas/projects/phalanx/source/biancheng/bc/" + strconv.Itoa(i) + "-" + title + ".html"
		f, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		writer := bufio.NewWriter(f)
		writer.WriteString(data)
		writer.Flush()
	})

	c.Visit(url)
}
