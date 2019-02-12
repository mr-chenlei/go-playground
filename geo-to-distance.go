package main

import (
	"fmt"
	"math"
)

const EARTH_RADIUS = 6378.137

type GeoData struct {
	longitude float64
	latitude  float64
}

func rad(d float64) float64 {
	return d * math.Pi / 180.0
}

func GetDistance(geoA, geoB *GeoData) float64 {
	radLat1 := rad(geoA.latitude)
	radLat2 := rad(geoB.latitude)
	a := radLat1 - radLat2
	b := rad(geoA.longitude) - rad(geoB.longitude)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+
		math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2)))
	s = s * EARTH_RADIUS
	s = math.Round(s*10000) / 10000
	return s
}

func main() {
	geoA := &GeoData{-17.561556, 14.869265}
	geoB := &GeoData{121.644135, 25.229726}
	distance := GetDistance(geoA, geoB)
	fmt.Println(distance)
}
