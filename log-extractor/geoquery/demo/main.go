package main

import (
	"fmt"

	"taas.com/galaxy/geoquery"
)

func main() {
	err := geoquery.Init("./ipip.ipdb")
	if err != nil {
		fmt.Println("geoquery err:", err)
		return
	}
	// 北京
	fmt.Println(geoquery.GetCCAndPC("39.96.21.158"))
	// 浙江
	fmt.Println(geoquery.GetCCAndPC("47.108.86.98"))
	// 内蒙古
	fmt.Println(geoquery.GetCCAndPC("222.74.126.230"))
	// 法国
	fmt.Println(geoquery.GetCCAndPC("193.105.43.227"))
	// 香港
	fmt.Println(geoquery.GetCCAndPC("47.52.153.193"))
	// 德国
	fmt.Println(geoquery.GetCCAndPC("47.254.171.174"))
	// 韩国
	fmt.Println(geoquery.GetCCAndPC("117.52.35.45"))
	// 台湾
	fmt.Println(geoquery.GetCCAndPC("210.59.131.43"))
	// 吃鸡
	fmt.Println(geoquery.GetCCAndPC("153.254.86.179"))
}
