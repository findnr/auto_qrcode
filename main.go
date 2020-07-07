package main

import (
	r "auto_qrcode/types"
	"flag"
)

func main() {
	var url string
	var action int
	var key string
	var types string
	argument := make(map[string]interface{})
	flag.StringVar(&url, "u", "http://www.gzwea.com", "生成的url地址固定部，默认为：http://www.gzwea.com")
	flag.IntVar(&action, "a", 0, "生成二给码时，名称的组合方式，0是按自增序号，1是按可变数据")
	flag.StringVar(&key, "k", "", "生成的密钥")
	flag.StringVar(&types, "t", "default", "加密的方式,默认为md5")
	flag.Parse()
	argument["url"] = url
	argument["action"] = action
	argument["key"] = key
	argument["types"] = types
	switch types {
	case "md5":
		obj := r.Md5Controller{}
		obj.Create(argument)
	case "default":
		obj := r.DefaultController{}
		obj.Create(argument)
	}
}
