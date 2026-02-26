package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// R 地球半径，单位米
const R = 6367000

// Distance 经纬度算距离
// lonA, latA分别为A点的纬度和经度
// lonB, latB分别为B点的纬度和经度
// 返回的距离单位为米
func Distance(lngA, latA, lngB, latB float64) float64 {
	c := math.Sin(latA)*math.Sin(latB)*math.Cos(lngA-lngB) + math.Cos(latA)*math.Cos(latB)
	return R * math.Acos(c) * math.Pi / 180
}

func main() {

	b, err := os.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}

	var dataList []struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}

	err = json.Unmarshal(b, &dataList)
	if err != nil {
		panic(err)
	}

	var dis float64
	if len(dataList) > 0 {
		p := struct {
			Lat float64
			Lon float64
		}{
			Lat: dataList[0].Lat,
			Lon: dataList[0].Lon,
		}

		if len(dataList) > 1 {
			for _, v := range dataList[1:] {
				// dis += CalcCoordsDistance(Point{p.Lon, p.Lat}, Point{v.Lon, v.Lat})
				dis += Distance(p.Lon, p.Lat, v.Lon, v.Lat)
				p.Lat = v.Lat
				p.Lon = v.Lon
			}
		}

	}

	fmt.Println("dis", dis)
}

// Point 表示一个点的经纬度坐标 [lng, lat]
type Point [2]float64

// CalcCoordsDistance 计算两个点(经纬度)的距离
// startDot: 开始点的经纬度[lng, lat]
// endDot: 结束点的经纬度[lng, lat]
// 返回: 距离，单位为米
func CalcCoordsDistance(startDot Point, endDot Point) float64 {
	earthRadius := 6378137.0 // 地球半径
	// PI := math.Pi                                                      // 圆周率π
	startRadianLat := getRadian(startDot[1])                           // 纬度 - 开始
	endRadianLat := getRadian(endDot[1])                               // 纬度 - 结束
	latDiffVal := startRadianLat - endRadianLat                        // 维度差值
	lngDiffVal := getRadian(startDot[0]) - getRadian(endDot[0])        // 经度差值
	latDiffSinVal := math.Sin(latDiffVal / 2)                          // 维度差值的正弦值
	lngDiffSinVal := math.Sin(lngDiffVal / 2)                          // 经度差值的正弦值
	latCosProduct := math.Cos(startRadianLat) * math.Cos(endRadianLat) // 维度的余弦值乘积
	powVal := latCosProduct * math.Pow(lngDiffSinVal, 2)
	sqrtVal := math.Pow(latDiffSinVal, 2) + powVal            // 开平方根的值
	result := 2 * math.Asin(math.Sqrt(sqrtVal)) * earthRadius // 结果值

	return math.Round(result*100) / 100 // 保留两位小数
}

// getRadian 将角度转换为弧度
func getRadian(d float64) float64 {
	return (d * math.Pi) / 180.0
}
