# Web_crawler 웹크롤러 만들기

## 사용한 패키지 Need to Package 
```go
import (
	"bufio"
	"fmt"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)
```

## 배운 점  What I learned and other thing, etc
    - Got used to use other's package
    - Got used to use golang official package
    - learned how to use other's package
    - I Could learn how to scrape code from web browser
    - Next time, I'd like to make a lot of crawler
    - Could you tell me if you have any ad vice and feedback?

## 웹 사이트 주소 상수 선언 Declare of URL site with const 
```go
//스크래핑 대상 URL
const (
	urlRoot = "http://ruliweb.com"
)
```

<img src="./screenshot/ruliweb.png" width="300">

## 웹 메인페이지에서 원하는 URL파싱 후 반환하는 함수 생성 Function will use return important URL from Web mainpage
```go
//첫 번째 방문(메인페이지) 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}
```

<img src="./screenshot/ruliweb.png" width="300">




