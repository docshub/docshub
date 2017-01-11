package md

import (
	"testing"
)

func Test_RenderMd(t *testing.T) {
	RenderMd()
}
func Test_GetToc(t *testing.T) {

	html := RenderMd()
	GetToc(html)
}
