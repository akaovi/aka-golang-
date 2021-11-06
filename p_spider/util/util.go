// package util

// import (
// 	"fmt"
// )

// type Bar struct {
// 	percent int64  
// 	cur     int64  
// 	total   int64  
// 	rate    string 
// 	graph   string 
// }

// func (b *Bar) NewOption(start, total int64) {
// 	b.cur = start
// 	b.total = total
// 	if b.graph == "" {
// 		b.graph = "█"
// 	}
// 	b.percent = b.GetPercent()
// }

// func (b *Bar) GetPercent() int64 {
// 	return int64(float32(b.cur) / float32(b.total) * 100)
// }

// func (b *Bar) Play(cur int64) {
// 	b.cur = cur
// 	last := b.percent
// 	b.percent = b.GetPercent()
// 	if b.percent != last && b.percent%2 == 0 {
// 		b.rate += b.graph
// 	}
// 	fmt.Printf("\r|%s] 已下载%d%%    当前：%d/%d", b.rate, b.percent, b.cur, b.total)
// }


package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Backtime() string {
	tem, _ := time.ParseDuration("-24h")
	yesterday := time.Now().Add(tem)
	year := yesterday.Format("2006")
	month := yesterday.Format("01")
	day := yesterday.Format("02")
	return year + month + day
}

func Download(url string, path string, boolchan chan bool) {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	name := path + "/" + fmt.Sprint(time.Now().UnixMicro()) + ".jpg"
	out, _ := os.Create(name)
	io.Copy(out, bytes.NewReader(body))
	boolchan <- true
}

func Parsel(time1, v, path string, exitchan chan bool) {
	url := "https://www.vilipix.com/api/illust?mode=daily&date=" + time1 + "&limit=30&offset=" + v
	client := &http.Client{}
	resq, _ := http.NewRequest("GET", url, nil)

	resq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36 Edg/95.0.1020.40")

	resp, err := client.Do(resq)

	if err != nil {
		fmt.Println("request err：", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return
	}

	var a map[string]interface{}
	encode := json.Unmarshal([]byte(string(body)), &a)
	if encode != nil {
		fmt.Println("Unmarshal err: ", encode)
		return
	}

	rows := a["rows"]

	// 每个pasel有30个协程
	urlchan := make(chan string, 30)
	boolchan := make(chan bool, 30)

	go func() {
		for _, v := range rows.([]interface{}) {
			urlchan <- v.(map[string]interface{})["regular_url"].(string)
		}
		close(urlchan)
	}()

	go func() {
		for i := 0; i < 30; i++ {
			oneurl := <-urlchan
			go Download(oneurl, path, boolchan)
		}
	}()

	for i := 0; i < 30; i++ {
		<-boolchan
	}
	close(boolchan)
	exitchan <- true

}



