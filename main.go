package main

import (
	"crawler/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main(){
	resp,err:=http.Get("http://xiaodiaodaya.cn/xh/jt/12.html")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	//改正则表达式作用是匹配所有的汉字
	reg:= regexp.MustCompile("[\\p{Han}]+")
	if reg == nil{
		fmt.Println("reg err")
		return
	}
	result:=reg.FindAllStringSubmatch(string(body),-1)
	var result1 []string
	var result2 string
	//将多个切片转换成字符串并且append到一个切片当中去
	for i:=0;i<len(result);i++{
		result1=append(result1,fmt.Sprintf(strings.Join(result[i],",")))
	}
	//将切片转换为字符串类型
	result2=fmt.Sprintf(strings.Join(result1,"\n"))
	//将文字信息写入文件，该文件根目录在main根目录下
	err= service.WriteResult(result2,"text.dat")
	if err!=nil{
		fmt.Printf("err is %t",err)
	}

}