package core

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yin-zt/cobra-cli/utils"
	"os"
	"strings"
)

// Pq 使用方法是 cli pq -m html -f xxx.html
// 目的是查询html文件中所有href属性的值，并将每个html中的href组件的链接和内容以字典方式进入列表中返回
func (this *Common) Pq(module string, action string) {
	var (
		err   error
		dom   *goquery.Document
		title string
		href  string
		ok    bool
		text  string
		html  string
	)
	fmt.Println(ok)
	u := ""
	s := "html"
	m := "text"
	f := ""
	c := ""
	a := "link"

	argv := utils.GetArgsMap()
	if v, ok := argv["f"]; ok {
		f = v
	}
	if v, ok := argv["s"]; ok {
		s = v
	}
	if v, ok := argv["m"]; ok {
		m = v
	}

	if v, ok := argv["a"]; ok {
		a = v
	}

	if f != "" {

		c = utils.ReadFile(f)

	}

	if c == "" {
		var lines []string
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			lines = append(lines, input.Text())
		}
		c = strings.Join(lines, "")
	}

	_ = c
	_ = m
	_ = s
	_ = u
	_ = err

	dom, err = goquery.NewDocumentFromReader(strings.NewReader(c))

	var result []interface{}
	dom.Find(s).Each(func(i int, selection *goquery.Selection) {

		if a == "link" {
			href, ok = selection.Attr("href")
			title = selection.Text()
			item := make(map[string]string)
			item["href"] = strings.TrimSpace(href)
			item["title"] = strings.TrimSpace(title)
			result = append(result, item)
		} else if a == "table" {
			var rows []string
			selection.Find("table tr").Each(func(i int, selection *goquery.Selection) {
				var row []string
				selection.Find("td").Each(func(i int, selection *goquery.Selection) {
					row = append(row, strings.TrimSpace(selection.Text()))
				})
				rows = append(rows, strings.Join(row, "######"))
			})

			result = append(result, strings.Join(rows, "$$$$$"))

		} else {

			if m == "text" {
				text = selection.Text()
				result = append(result, text)
			}
			if m == "html" {
				fmt.Println(123)
				html, err = selection.Html()
				result = append(result, html)
			}
		}
	})

	fmt.Println(utils.JsonEncodePretty(result))

}
