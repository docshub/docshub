package md

import (
	"fmt"

	"bytes"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

//DocTocItem toc children
type TocAttr struct {
	Name string
	href string
}

//
type TocItem struct {
	TocAttr
	children TocAttr
}

//RenderMd render markdown content to html
func RenderMd() *string {
	md := `
[toc]
### 云油四海 主工程
##### 在线参考文档
###### xxxx
- [交互原型](http://devkits.changhong.io/axure/yysh/)
- [APP API DOCS](./../../blob/master/docs/app-api.md)

### 工程模块
* yysh-portal: 云油四海企业门户web
`
	unsafe := blackfriday.MarkdownCommon([]byte(md))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	htmlStr := string(html)
	fmt.Println(htmlStr)
	return &htmlStr
}

//GetToc get md toc
func GetToc(htmlStr *string) {
	// node, err := html.Parse(strings.NewReader(*htmlStr))

	// fmt.Println(node.)
	b := bytes.NewBufferString(*htmlStr)
	fmt.Println("---------------------")
	fmt.Println(b)
	el, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		//todo

	}
	fmt.Printf("header length: %d\n", el.Find(":header").Length())
	// first check h1
	fmt.Println(el.Find("h1").Length())
	fmt.Println(el.Find("h2").Length())

	fmt.Println(el.Find("h3").Length())
	fmt.Println(el.Find("h8").Length())

	if el.Find("h1") == nil {
		fmt.Println("no h1")
	}

	selector := el.Find("h3")
	selector.AddClass(".cursor")
	//get tag name and add id
	fmt.Println(selector.First().Next().Get(0).Data)
	fmt.Println(selector.Length())
	fmt.Println("title" + selector.Text())

}
//FetchToc fecth Toc from html node
func FetchToc(toc *TocItem, dom *goquery.Document, parent *goquery.Selection) {

	//判断父节点是否包含子节点
	parentNodeName := parent.Get(0).Data

	//选择游标所在节点
	currentNode := dom.Find(".cursor")
	//判断游标所在标签是否为父节点
	if (currentNode.Get(0).Data == parentNodeName) {
		//todo toc item 操作

	}

	//游标处理
	//判断是否有子标题节点 标记游标位置


	//否则游标指向兄弟标题节点
	//children
	parent.Has(parentNodeName)
	selection := dom.Find(".cursor")
	fmt.Printf("toc name is: %s\n", selection.Get(0).Data)
	selection.RemoveClass(".cursor")
}
func FindAllHeaders(dom *goquery.Document) (headers []*goquery.Selection){

}
