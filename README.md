# Web_crawler 웹크롤러 만들기

## Need to Package 사용한 패키지
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

## What I learned and other thing, etc... 배운 점 
    - Got used to use other's package
    - Got used to use golang official package
    - learned how to use other's package
    - I Could learn how to scrape code from web browser
    - Next time, I'd like to make a lot of crawler
    - Could you tell me if you have any ad vice and feedback?

## First ) Declare of URL site with const 웹 사이트 주소 상수 선언
```go
//스크래핑 대상 URL
const (
	urlRoot = "http://ruliweb.com"
)
```
<img src="./screenshot/ruliweb.png" width="100">


