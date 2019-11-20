//대상 사이트 : 루리웹(www.ruriweb.com)
//대한민국 최고 규모 게임 관련 인터넷 커뮤니티

package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
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

func parseSubNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "deco"
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

//Url 대상이 되는 페이지(서브페이지) 대상으로 원하는 내용을 파싱 후 반환
func scrapContents(href, fn string) {
	//작업 종료 알림
	defer wg.Done()
	//Get 방식 요청
	response, err := http.Get(href)
	//에러체크
	errCheck(err)
	//코드 읽어왔으면 닫기
	defer response.Body.Close()
	//html 바디 코드 전부다 root에 넣기
	root, err := html.Parse(response.Body)
	//에러체크
	errCheck(err)

	//파일 스크림 생성(열기) -> 파일명, 옵션, 권한
	file, err := os.OpenFile("/Users/jejeongmin/documents/go/src/Scrape/"+fn+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	//에러체크
	errCheck(err)

	//메소드 종료 시 파일 닫기
	defer file.Close()

	//쓰기 버퍼 선언
	w := bufio.NewWriter(file)

	//matchNode 메소드를 사용해서 원하는 노드 순회(Iterator)하면서 출력
	for _, g := range scrape.FindAll(root, parseSubNodes) {
		//Url 및 해당 데이터 출력
		fmt.Println()
	}
}

func main() {
	//메인 페이지 Get 방식 요청
	response, err := http.Get(urlRoot) //response(응답), request(요청)
	errCheck(err)

	//요청 Body 닫기
	defer response.Body.Close()

	//응답 데이터(HTML)
	root, err := html.Parse(response.Body) //root에 루리웹사이트의 바디코드를 전부 다 넣음.
	errCheck(err)

	//대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)

	//class = row인 태그들 싹 다 긁어와서 for range로 순회
	for _, link := range urlList {
		//대상 Url 1차 출력
		// fmt.Println("Main Link : ", link, idx)
		// fmt.Println("TargetUrl : ", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)
		fmt.Println("fileName : ", fileName)

		//작업 대기열에 추가
		wg.Add(1) //Done 개수와 일치
		//고루틴 시작 -> 작업 대기열 개수와 같아야 함.
		go scrapContents(scrape.Attr(link, "href"), fileName) //href 값으로 각 링크에 들어가는 거, 텍스트 파일 만들 숫자만 갖고오는 거
	}
	wg.Wait()
}
