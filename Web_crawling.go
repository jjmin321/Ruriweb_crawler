//대상 사이트 : 루리웹(www.ruriweb.com)
//대한민국 최고 규모 게임 관련 인터넷 커뮤니티

package main

import (
	_ "bufio"
	"fmt"
	"net/http"
	_ "os"
	_ "strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// https://github.com/yhat/scrape : 사용하기 어렵지만, 코드 학습 위해서 사용
// https://go-colly.org/docs/ : goquery기반 굉장히 강력하고 쉬운 패키지(가장 많이 사용)
// https://github.com/PuerkitoBio/goquery : 쉬운 HTML Parsing 지원

//스크래핑 대상 URL
const (
	urlRoot = "http://ruliweb.com"
)

//첫 번째 방문(메인페이지) 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}

//에러 체크 공통 함수
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

//동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

func main() {
	//메인 페이지 Get 방식 요청
	response, err := http.Get(urlRoot) //response(응답), request(요청)
	errCheck(err)

	//요청 Body 닫기
	defer response.Body.Close()

	//응답 데이터(HTML)
	root, err := html.Parse(response.Body) //root에 루리웹사이트의 바디코드를 전부 다 넣음.
	errCheck(err)

	//define a matcher
	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "athing"
		}
		return false
	}
	//대상 URL 추출
	urlList := scrape.FindAll(root, parse)
	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}
