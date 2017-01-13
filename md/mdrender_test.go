package md

import (
	"testing"
	"github.com/PuerkitoBio/goquery"

	"fmt"
	"bytes"
)

func Test_RenderMd(t *testing.T) {
	RenderMd()
}
func Test_GetToc(t *testing.T) {

	html := RenderMd()
	GetToc(html)
}
func Test_GetAllHeaders(t *testing.T){
	b := bytes.NewBufferString( *RenderMd())
	fmt.Println("---------------------")
	fmt.Println(b)
	el, err := goquery.NewDocumentFromReader(b)
	if(err!=nil){
		//todo
	}
	var toc []TocItem
	FindAllHeaders(el,toc)
	//goquery.Matcher()
	//t.Log("hhhh", el.FindMatcher(":header").Length())
}
