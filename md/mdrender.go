package md

import (
	"fmt"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"strconv"
	"strings"
)

//DocTocItem toc children
type TocAttr struct {
	Name string
	href string
}

//
type TocItem struct {
	Name     string
	url      string
	Id       string
	ParentId string
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
func FindAllHeaders(dom *goquery.Document, pToc []TocItem) {

	var startHeader *goquery.Selection
	for i := 1; i < 7; i++ {
		tagName := "h" + strconv.Itoa(i)
		fmt.Println("tag name: ", tagName)
		header := dom.Find(tagName)
		fmt.Println(" has " + tagName + ": " + strconv.Itoa(header.Length()))
		if (header.Length() > 0) {
			if (startHeader == nil) {
				startHeader = header.First()
				fmt.Println("header start " + goquery.NodeName(startHeader))
			}
			//counter += header.Length()
		}
	}
	//parentId := 0
	//for i := 0; i < counter; i++ {
	//	//var headerNode *html.Node
	//	if i == 0 {
	//		pToc = append(pToc, &TocItem{Name:goquery.NodeName(startHeader)})
	//		//headerNode = startHeader.Get(0).Data
	//		headerSelection = append(headerSelection, startHeader)
	//
	//	} else {
	//		//headerNode=headerSelection[i-1].Get(0).Data
	//		tagName := goquery.NodeName(headerSelection[i - 1].Next())
	//		fmt.Println("tagName: " + tagName)
	//		if (strings.HasPrefix(tagName, "h")) {
	//			headerSelection = append(headerSelection, headerSelection[i - 1].Next())
	//			pToc = append(pToc, &TocItem{Name:goquery.NodeName(headerSelection[i - 1].Next())})
	//		}
	//
	//	}
	//	//fmt.Printf("hs: %v\n", headerSelection[i].Next().Get(0).Data)
	//}
	var headerSelection []goquery.Selection
	cursor := startHeader
	//headerSelection = append(headerSelection, *startHeader)
	//pToc = append(pToc, TocItem{ Name:startHeader.Get(0).Data})
	for cursor.Length()>0 {
		fmt.Println("==============" + goquery.NodeName(cursor)+ "================")
		tagName := goquery.NodeName(cursor)
		if (strings.HasPrefix(tagName, "h")) {
			headerSelection = append(headerSelection, *cursor)
			pToc = append(pToc, TocItem{Name:tagName})
		fmt.Println("==============" + strconv.Itoa(cursor.Length())+ "================")
		}
		cursor = cursor.Next()
		fmt.Println("==============" + strconv.Itoa(cursor.Length())+ "================")

	}
	fmt.Printf("headers %+v\n", headerSelection)
	fmt.Printf("%+v\n", pToc)
	fmt.Println(pToc)
	//startToc := &TocItem{Name:startHeader.Get(0).Data, Id:parentId}
	return
}
