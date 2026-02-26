package main

import (
	"math"
)

// Point 表示一个点的经纬度坐标 [lng, lat]
type Point [2]float64

// IsPointInPolygon 判断一个点是否在一个多边形内
// pts: 多边形的顶点数组，每个顶点是一个包含经纬度的数组 [lng, lat]
// point: 要判断的点，包含经纬度的数组 [lng, lat]
// tolerance: 误差范围，单位为米，默认值为 100 米
// 返回: 如果点在多边形内或在误差范围内，返回 true；否则返回 false
func IsPointInPolygon(pts []Point, point Point, tolerance float64) bool {
	// 遍历所有顶点与点之间的距离是否小于等于tolerance
	if tolerance > 0 {
		for i := 0; i < len(pts); i++ {
			if CalcCoordsDistance(pts[i], point) <= tolerance {
				return true
			}
		}
	}

	N := len(pts)
	boundOrVertex := true // 如果点位于多边形的顶点或边上，也算做点在多边形内，直接返回true
	intersectCount := 0   // cross points count of x
	precision := 2e-10    // 浮点类型计算时候与0比较时候的容差
	var p1, p2 Point      // neighbour bound vertices
	p := point
	p1 = pts[0] // left vertex

	for i := 1; i <= N; i++ { // check all rays
		if (p[1] == p1[1]) && (p[0] == p1[0]) {
			return boundOrVertex // p is a vertex
		}
		p2 = pts[i%N] // right vertex

		if p[1] < math.Min(p1[1], p2[1]) || p[1] > math.Max(p1[1], p2[1]) {
			// ray is outside of our interests
			p1 = p2
			continue // next ray left point
		}

		if p[1] > math.Min(p1[1], p2[1]) && p[1] < math.Max(p1[1], p2[1]) {
			// ray is crossing over by the algorithm (common part of)
			if p[0] <= math.Max(p1[0], p2[0]) {
				// x is before of ray
				if p1[1] == p2[1] && p[0] >= math.Min(p1[0], p2[0]) {
					// overlies on a horizontal ray
					return boundOrVertex
				}
				if p1[0] == p2[0] { // ray is vertical
					if p1[0] == p[0] { // overlies on a vertical ray
						return boundOrVertex
					} else { // before ray
						intersectCount++
					}
				} else { // cross point on the left side
					xinters := (p[1]-p1[1])*(p2[0]-p1[0])/(p2[1]-p1[1]) + p1[0] // cross point of lng
					if math.Abs(p[0]-xinters) < precision {                     // overlies on a ray
						return boundOrVertex
					}
					if p[0] < xinters { // before ray
						intersectCount++
					}
				}
			}
		} else {
			// special case when ray is crossing through the vertex
			if p[1] == p2[1] && p[0] <= p2[0] {
				// p crossing over p2
				p3 := pts[(i+1)%N] // next vertex
				if p[1] >= math.Min(p1[1], p3[1]) && p[1] <= math.Max(p1[1], p3[1]) {
					// p[1] lies between p1[1] & p3[1]
					intersectCount++
				} else {
					intersectCount += 2
				}
			}
		}
		p1 = p2 // next ray left point
	}

	if intersectCount%2 == 0 { // 偶数在多边形外
		return false
	} else { // 奇数在多边形内
		return true
	}
}

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

// 坐标系转换相关常量
const (
	PI = 3.1415926535897932384626
	a  = 6378245.0              // 卫星椭球坐标投影到平面地图坐标系的投影因子
	ee = 0.00669342162296594323 // 椭球的偏心率
)

// transformLng 转化经度
func transformLng(lng, lat float64) float64 {
	ret := 300.0 + lng + 2.0*lat + 0.1*lng*lng + 0.1*lng*lat + 0.1*math.Sqrt(math.Abs(lng))
	ret += ((20.0*math.Sin(6.0*lng*PI) + 20.0*math.Sin(2.0*lng*PI)) * 2.0) / 3.0
	ret += ((20.0*math.Sin(lng*PI) + 40.0*math.Sin((lng/3.0)*PI)) * 2.0) / 3.0
	ret += ((150.0*math.Sin((lng/12.0)*PI) + 300.0*math.Sin((lng/30.0)*PI)) * 2.0) / 3.0
	return ret
}

// transformLat 转化纬度
func transformLat(lng, lat float64) float64 {
	ret := -100.0 + 2.0*lng + 3.0*lat + 0.2*lat*lat + 0.1*lng*lat + 0.2*math.Sqrt(math.Abs(lng))
	ret += ((20.0*math.Sin(6.0*lng*PI) + 20.0*math.Sin(2.0*lng*PI)) * 2.0) / 3.0
	ret += ((20.0*math.Sin(lat*PI) + 40.0*math.Sin((lat/3.0)*PI)) * 2.0) / 3.0
	ret += ((160.0*math.Sin((lat/12.0)*PI) + 320*math.Sin((lat*PI)/30.0)) * 2.0) / 3.0
	return ret
}

// OutOfChina 判断是否在国内还是国外
func OutOfChina(lon, lat float64) bool {
	if lon < 72.004 || lon > 137.8347 {
		return true
	}
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	return false
}

// WGS84ToGCJ02 地球坐标系转火星坐标系
func WGS84ToGCJ02(lng, lat float64) Point {
	dlat := transformLat(lng-105.0, lat-35.0)
	dlng := transformLng(lng-105.0, lat-35.0)
	radlat := (lat / 180.0) * PI
	magic := math.Sin(radlat)
	magic = 1 - ee*magic*magic
	sqrtmagic := math.Sqrt(magic)
	dlat = (dlat * 180.0) / (((a * (1 - ee)) / (magic * sqrtmagic)) * PI)
	dlng = (dlng * 180.0) / ((a / sqrtmagic) * math.Cos(radlat) * PI)
	mglat := lat + dlat
	mglng := lng + dlng

	return Point{mglng, mglat}
}

// GCJ02ToWGS84 火星坐标系转地球坐标系
func GCJ02ToWGS84(lng, lat float64) Point {
	originalLngSign := math.Copysign(1, lng)
	originalLatSign := math.Copysign(1, lat)
	lat = math.Abs(lat)
	lng = math.Abs(lng)
	dlat := transformLat(lng-105.0, lat-35.0)
	dlng := transformLng(lng-105.0, lat-35.0)
	radlat := lat / 180.0 * PI
	magic := math.Sin(radlat)
	magic = 1 - ee*magic*magic
	sqrtmagic := math.Sqrt(magic)
	dlat = (dlat * 180.0) / ((a * (1 - ee)) / (magic * sqrtmagic) * PI)
	dlng = (dlng * 180.0) / (a / sqrtmagic * math.Cos(radlat) * PI)
	mglat := lat + dlat
	mglng := lng + dlng
	lngs := lng*2 - mglng
	lats := lat*2 - mglat
	finalLng := originalLngSign * lngs
	finalLat := originalLatSign * lats

	return Point{finalLng, finalLat}
}

// ConvertWGS84ArrayToGCJ02 批量转换WGS84坐标数组到GCJ02坐标系
func ConvertWGS84ArrayToGCJ02(points []Point) []Point {
	result := make([]Point, len(points))
	for i, point := range points {
		result[i] = WGS84ToGCJ02(point[0], point[1])
	}
	return result
}
