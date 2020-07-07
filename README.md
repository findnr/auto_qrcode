# auto_qrcode
命令批量二维码生成
# 使用说明
- 安装
```shell
git clone https://github.com/findnr/auto_qrcode
cd auto_qrcode
go build -o main main.go
```
#查看帮助
- linux
```shell
./main -h
```
- windows
```shell
.\main.exe -h
```
-a int
    生成二给码时，名称的组合方式，0是按自增序号，1是按可变数据
-k string
    生成的密钥
-t string
    加密的方式,默认为default (其中包括md5)
-u string
    生成的url地址固定部，默认为：http://www.gzwea.com (default "http://www.gzwea.com")
# 数据说明
数据来自data.txt按行来生成（有几行就生成几个二维码）
# 二维码生成的目录
qrcode/
# 例
```shell
./main -u http://check.gzwea.com/jiaoyiyuan/check?n= -a 1 -t md5 -k 123456
```
生成出来的二维码地址：http://check.gzwea.com/jiaoyiyuan/check?n=522101198909134653&secret=a13062a2b3c39fa0ec4fd8fb5af63c60