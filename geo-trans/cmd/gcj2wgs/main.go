package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	X_PI   = math.Pi * 3000.0 / 180.0
	OFFSET = 0.00669342162296594323
	AXIS   = 6378245.0
)

func main() {

	b, err := os.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(b))
	n := gjson.GetBytes(b, "#").Int()
	// var bCopy = b
	fmt.Println("1:" + string(b))
	for i := 0; i < int(n); i++ {
		v := gjson.GetBytes(b, strconv.Itoa(i)+".lat_and_long").String()
		// fmt.Println(v)
		var lonStr, latStr string
		loc := strings.Split(v, ",")
		lonStr, latStr = loc[0], loc[1]
		var lon, lat float64
		lon, err = strconv.ParseFloat(lonStr, 64)
		if err != nil {
			panic(err)
		}
		lat, err = strconv.ParseFloat(latStr, 64)
		if err != nil {
			panic(err)
		}
		transLon, transLat := GCJ02toWGS84(lon, lat)
		vTrans := fmt.Sprintf("%f,%f", transLon, transLat)
		fmt.Printf("%s --> %s \n", v, vTrans)
		b, err = sjson.SetBytes(b, strconv.Itoa(i)+".lat_and_long", vTrans)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("2:" + string(b))
	// fList := gjson.Get(b)
	// for _, v := range fList {
	// 	fmt.Println(v)
	// }

	err = os.WriteFile("./out.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

// GCJ02toWGS84 火星坐标系->WGS84坐标系
func GCJ02toWGS84(lon, lat float64) (float64, float64) {
	if isOutOFChina(lon, lat) {
		return lon, lat
	}

	mgLon, mgLat := delta(lon, lat)

	return lon*2 - mgLon, lat*2 - mgLat
}

func isOutOFChina(lon, lat float64) bool {
	return !(lon > 72.004 && lon < 135.05 && lat > 3.86 && lat < 53.55)
}

func delta(lon, lat float64) (float64, float64) {
	dlat, dlon := transform(lon-105.0, lat-35.0)
	radlat := lat / 180.0 * math.Pi
	magic := math.Sin(radlat)
	magic = 1 - OFFSET*magic*magic
	sqrtmagic := math.Sqrt(magic)

	dlat = (dlat * 180.0) / ((AXIS * (1 - OFFSET)) / (magic * sqrtmagic) * math.Pi)
	dlon = (dlon * 180.0) / (AXIS / sqrtmagic * math.Cos(radlat) * math.Pi)

	mgLat := lat + dlat
	mgLon := lon + dlon

	return mgLon, mgLat
}

func transform(lon, lat float64) (x, y float64) {
	var lonlat = lon * lat
	var absX = math.Sqrt(math.Abs(lon))
	var lonPi, latPi = lon * math.Pi, lat * math.Pi
	var d = 20.0*math.Sin(6.0*lonPi) + 20.0*math.Sin(2.0*lonPi)
	x, y = d, d
	x += 20.0*math.Sin(latPi) + 40.0*math.Sin(latPi/3.0)
	y += 20.0*math.Sin(lonPi) + 40.0*math.Sin(lonPi/3.0)
	x += 160.0*math.Sin(latPi/12.0) + 320*math.Sin(latPi/30.0)
	y += 150.0*math.Sin(lonPi/12.0) + 300.0*math.Sin(lonPi/30.0)
	x *= 2.0 / 3.0
	y *= 2.0 / 3.0
	x += -100.0 + 2.0*lon + 3.0*lat + 0.2*lat*lat + 0.1*lonlat + 0.2*absX
	y += 300.0 + lon + 2.0*lat + 0.1*lon*lon + 0.1*lonlat + 0.1*absX
	return
}
