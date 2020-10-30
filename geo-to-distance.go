package main

import (
	"fmt"
	"math"
)

// 赤道半径
const EarthRadius = 6378.137

type GeoData struct {
	longitude float64
	latitude  float64
}

func rad(d float64) float64 {
	return d * math.Pi / 180.0
}

// GetDistance 计算两个地理坐标距离（Km）
// 参数：
//      gA: 地理坐标
//      gB: 地理坐标
// 返回值：
//      距离值，单位千米（Km）
func GetDistance(gA, gB *GeoData) float64 {
	rla := rad(gA.latitude)
	rlb := rad(gB.latitude)
	a := rla - rlb
	b := rad(gA.longitude) - rad(gB.longitude)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+
		math.Cos(rla)*math.Cos(rlb)*math.Pow(math.Sin(b/2), 2)))
	s = s * EarthRadius
	s = math.Round(s*10000) / 10000
	return s
}

func main() {
	gA := &GeoData{-17.561556, 14.869265} // Dakar Senegal, West Africa
	gB := &GeoData{121.644135, 25.229726} // 台北
	distance := GetDistance(gA, gB)
	fmt.Println("Distance from Dakar Senegal to 台北：", distance)
}
