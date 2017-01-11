package service

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

//RenderMd render markdown content to html
func RenderMd() *string {
	md := `
[toc]
### 云油四海 主工程
##### 在线参考文档
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
