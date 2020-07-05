package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

	qrcode "github.com/skip2/go-qrcode"
)

var wg sync.WaitGroup

func CreateQrcode(url string, address string) {
	// fmt.Println(address)
	// var ss string
	defer wg.Done()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	err := qrcode.WriteFile(url, qrcode.Medium, 256, "qrcode/"+address+".png")
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	var url string
	var action int
	flag.StringVar(&url, "u", "http://www.gzwea.com", "生成的url地址固定部，默认为：http://www.gzwea.com")
	flag.IntVar(&action, "a", 0, "生成二给码时，名称的组合方式，0是按自增序号，1是按可变数据")
	flag.Parse()
	strFile, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("data.txt文件读取错误")
		return
	}
	arr := strings.Split(strings.Replace(string(strFile), "\r", "", -1), "\n")

	if action == 0 {
		i := 1
		for _, v := range arr {
			wg.Add(1)
			go CreateQrcode(url+v, strconv.Itoa(i))
			i++
		}
	} else {
		for _, v := range arr {
			// fmt.Printf("%T\n", str)
			wg.Add(1)
			go CreateQrcode(url+v, v)
		}
	}
	wg.Wait()
	fmt.Println("创建完成")
}
