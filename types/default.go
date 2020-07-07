package types

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

	qrcode "github.com/skip2/go-qrcode"
)

var defaultWg sync.WaitGroup

type DefaultController struct{}

func (this *DefaultController) createQrcode(url string, address string) {
	// fmt.Println(address)
	// var ss string
	defer defaultWg.Done()
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

func (this *DefaultController) Create(a map[string]interface{}) {
	strFile, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("data.txt文件读取错误123")
		return
	}
	arr := strings.Split(strings.Replace(string(strFile), "\r", "", -1), "\n")
	action, ok := a["action"].(int)
	if !ok {
		fmt.Println("非法操作#action")
		return
	}
	url, ok := a["url"].(string)
	if !ok {
		fmt.Println("非法操作#url")
		return
	}
	if action == 0 {
		i := 1
		for _, v := range arr {
			defaultWg.Add(1)
			go this.createQrcode(url+v, strconv.Itoa(i))
			i++
		}
	} else {
		for _, v := range arr {
			// fmt.Printf("%T\n", 1)
			defaultWg.Add(1)
			go this.createQrcode(url+v, v)
		}
	}
	defaultWg.Wait()
	fmt.Println("defaule模式生成成功")
}
