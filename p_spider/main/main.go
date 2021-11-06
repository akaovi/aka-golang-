// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"go_code/p_spider/util"
// )

// func Backtime() string {
// 	tem, _ := time.ParseDuration("-24h")
// 	yesterday := time.Now().Add(tem)
// 	year := yesterday.Format("2006")
// 	month := yesterday.Format("01")
// 	day := yesterday.Format("02")
// 	return year + month + day
// }

// func Download(url string, path string) {
// 	resp, _ := http.Get(url)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	name := path + "/" + fmt.Sprint(time.Now().UnixMicro()) + ".jpg"
// 	out, _ := os.Create(name)
// 	io.Copy(out, bytes.NewReader(body))
// }

// func Resp(path string) {
// 	arr1 := [...]string{"0", "30", "60", "90", "120"}
// 	time1 := Backtime()

// 	for _, v := range arr1 {
// 		url := "https://www.vilipix.com/api/illust?mode=daily&date=" + time1 + "&limit=30&offset=" + v

// 		path2 := path + "/" + v

// 		fmt.Printf("正在下载:%s\n", path2)

// 		err2 := os.MkdirAll(path2, os.ModePerm)
// 		if err2 != nil {
// 			fmt.Println(err2)
// 			return
// 		}

// 		client := &http.Client{}
// 		resq, _ := http.NewRequest("GET", url, nil)
// 		resq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36 Edg/95.0.1020.40")
// 		resp, err := client.Do(resq)
// 		if err != nil {
// 			fmt.Println("request err：", err)
// 			return
// 		}
// 		if resp.StatusCode != 200 {
// 			fmt.Println("Http status code:", resp.StatusCode)
// 			return
// 		}
// 		defer resp.Body.Close()
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println("Read error", err)
// 			return
// 		}

// 		var a map[string]interface{}
// 		encode := json.Unmarshal([]byte(string(body)), &a)
// 		if encode != nil {
// 			fmt.Println("Unmarshal err: ", encode)
// 			return
// 		}

// 		rows := a["rows"]

// 		var bar util.Bar
// 		bar.NewOption(0, 30)
// 		for i, v := range rows.([]interface{}) {
// 			regular_url := v.(map[string]interface{})["regular_url"]
// 			Download(regular_url.(string), path2)
// 			bar.Play(int64(i + 1))
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	// 返回当前运行的地址
// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	path2 := dir + "/每日一爬/" + Backtime()

// 	// 创建存储目录
// 	err2 := os.MkdirAll(path2, os.ModePerm)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		return
// 	}
// 	Resp(path2)
// 	fmt.Println("按任意键退出程序...")
// 	fmt.Scanln()
// }

// // 没用 协程 怕被封


package main

import (
	"fmt"
	"go_code/fetch2/bar"
	"go_code/fetch2/util"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 返回当前运行的地址
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	path := dir + "/每日一爬/" + util.Backtime()

	// 创建存储目录
	err2 := os.MkdirAll(path, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println("文件初始化完毕...")
	fmt.Println("加载下载组件...")

	arr1 := [...]string{"0", "30", "60", "90", "120"}
	exitchan := make(chan bool, 5)

	go func() {
		for i, v := range arr1 {
			fmt.Printf("协程%d开启...\n", i+1)
			path2 := path + "/" + v
			err2 := os.MkdirAll(path2, os.ModePerm)
			if err2 != nil {
				fmt.Println(err2)
				return
			}
			go util.Parsel(util.Backtime(), v, path2, exitchan)
		}
	}()

	var bar bar.Bar
	bar.NewOption(0, 5)
	for i := 0; i < 5; i++ {
		<-exitchan
		bar.Play(int64(i + 1))
	}
	fmt.Println()
	close(exitchan)

	fmt.Println("按任意键退出...")
	fmt.Scanln()
}


