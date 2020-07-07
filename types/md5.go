package types

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

	"auto_qrcode/common/cfunc"

	qrcode "github.com/skip2/go-qrcode"
)

type Md5Controller struct {
}

var md5Wg sync.WaitGroup

func (this *Md5Controller) createQrcode(url string, address string) {
	// fmt.Println(address)
	// var ss string
	defer md5Wg.Done()
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

func (this *Md5Controller) Create(a map[string]interface{}) {
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
	key, ok := a["key"].(string)
	if !ok {
		fmt.Println("非法操作#key")
		return
	}
	if action == 0 {
		i := 1
		for _, v := range arr {
			md5Wg.Add(1)
			go this.createQrcode(url+v+"&secret="+cfunc.MD5(v+key), strconv.Itoa(i))
			i++
		}
	} else {
		for _, v := range arr {
			// fmt.Printf("%T\n", 1)
			md5Wg.Add(1)
			go this.createQrcode(url+v+"&secret="+cfunc.MD5(v+key), v)
		}
	}
	md5Wg.Wait()
	fmt.Println("md5模式生成成功")
	fmt.Printf("%#v", a)
}
