# Ruriweb_crawler

## Stack
|           |     Crawler      |
|:---------:|:---------:|
| Developer | 제정민 | 
| Develop Language | GO |  
| Develop Tool     | Visual Studio Code|

## 🧺 동기화 Synchronization
    - mutex.Lock(), mutex.Unlock() vs sync.WaitGroup
    - mutex 사용 시 고루틴이 하나씩 실행되게 끔 동기화 가능
    - sync.WaitGroup을 사용해 mutex보다 조금 더 동기화에 신경쓰지 않을 수 있는 거 같음.
```go
//동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup
```

## 📃 웹 메인페이지에서 원하는 URL파싱 후 반환하는 함수 생성 Function will use return important URL from Web mainpage
만약 웹 브라우저 HTML 코드에 a태그와 부모태그가 모두 있는 태그라면 class가 row인지 분석 후 반환
```go
//첫 번째 방문(메인페이지) 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}
```


## 👩‍💻 웹 브라우저 코드 갖고와서 분석 Analysis HTML code
```go
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
```

## 🧵 urlList 객체 분석 후 파싱해오고자 하는 사이트들 고루틴(쓰레드) 통해 접속 Connect to desirous site with Goroutine(Thread)
    - 쓰레드 하나 당 하나씩 작업 대기열 추가(동기화, wg.Add(1))
```go
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
```

## 📁 파일 생성 Create File
    - 지정한 경로에 파일이 없다면 생성/있다면 수정 
    - 퍼미션은 777로 줌.
    - defer함수로 파일 안 닫으면 대참사 발생(주의)
```go
//파일 스크림 생성(열기) -> 파일명, 옵션, 권한
    file, err := os.OpenFile("/Users/jejeongmin/documents/go/src/Web_crawler/scrape/"+fn+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
//에러체크
    errCheck(err)
//메소드 종료 시 파일 닫기
	defer file.Close()
```

## 🗒️ .txt파일에 파싱내용 넣기 Input to txt
    - scrape.FindAll 사용해서 원하는 내용만 긁어오기
    - w.Flush로 버퍼 비워주면서 한 번에 내용 다 넣기
```go
    //쓰기 버퍼 선언
	w := bufio.NewWriter(file)

	//parseSubNodes 함수를 사용해서 원하는 노드 순회(Iterator)하면서 출력
	for _, g := range scrape.FindAll(root, parseSubNodes) {
		//Url 및 해당 데이터 출력
		fmt.Println("result : ", scrape.Text(g))
		//파싱 데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}
    w.Flush()
```

<img src="./Screenshot/Scrape .png" width="1000">

## 😞 삽질했던 부분 Difficult point
    - 어디서 에러가 발생했는지 한참 헤매다가 찾음
    - 어찌저찌 해결했긴 했으나 아직 좀 의문이다.

### 🙅‍♂️ 에러코드
<img src="./Screenshot/It's error .png" width="1000">

### 🙆‍♂️수정한 코드
<img src="./Screenshot/It's correct .png" width="1000">






