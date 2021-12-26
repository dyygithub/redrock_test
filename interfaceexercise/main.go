package main

import (
	"fmt"
	"math/rand"
	"sort"
	_ "sort"
)
//声明一个结构体
type hero struct {
	name string
	age int
}
//声明一个hero结构体的切片类型
type heroslice []hero
//实现Interface接口
func(hs heroslice)Len()int {
	return len(hs)
}
//Less方法就是决定你使用什么标准进行排序
//1.按hero年龄从小到大排序
func(hs heroslice)Less(i,j int)bool{
	return hs[i].age<hs[j].age
}
func(hs heroslice)Swap(i int,j int){
	temp:=hs[i]
	hs[i]=hs[j]
	hs[j]=temp
}

func main(){

	//先定义一个数组/切片
	var intslice=[]int{0,-1,10,7,90}
	//要求对intslice进行排序
	//1.冒泡排序。。
	//2.也可以使用系统提供的方法
	sort.Ints(intslice)
	fmt.Println(intslice)
	//对一个结构体进行排序
	//1.冒泡排序
	//2.使用系统方法
	//测试看看是否可以对结构体切片进行排序
	var heroes heroslice
	for i:=0;i<10;i++{
		hero:=hero{
			name:fmt.Sprintf("英雄~%d",rand.Intn(100)),
			age:rand.Intn(100),
		}
		//将hero append到heroes切片
		heroes=append(heroes,hero)
	}
	//看看排序前的排序
	for _,v:=range heroes{
		fmt.Println(v)
	}
	sort.Sort(heroes)
	fmt.Println("----------排序后---------")
	//看看排序后的顺序
	for _,v:=range heroes{
		fmt.Println(v)
	}
}
