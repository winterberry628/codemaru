package codemaru

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func WebFileDownload(urlAddress string, dir string, fileName string, referer string, showMessage bool) {
	parsedUrl, err := url.Parse(urlAddress)
	if err != nil {
		fmt.Println("URL 파싱 불가.")
	}

	urlPath := parsedUrl.Path
	segments := strings.Split(urlPath, "/")
	if fileName == "" { // 파일명이 없으면 웹 파일명으로
		fileName = segments[len(segments)-1]
	}

	// 디렉토리 생성
	os.MkdirAll(dir, 0755)

	// 파일 생성
	file, err := os.Create(dir + string(os.PathSeparator) + fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("파일 생성 불가.")
	}

	data := url.Values{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlAddress, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("URL 리퀘스트 실패.")
	}

	req.Header.Add("User-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36")
	req.Header.Add("User-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36")
	req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Add("Accept", "image/webp,image/*,*/*;q=0.8")
	req.Header.Add("Accept-Language", "ko-KR,ko;q=0.8,en-US;q=0.6,en;q=0.4")
	req.Header.Add("Referer", referer)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(urlAddress, "다운로드 실패!")
	}

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(urlAddress, "파일쓰기 실패!")
	}

	if showMessage == true {
		fmt.Println(urlAddress, "다운로드 완료!", size, "Byte")
	}

	return
}
