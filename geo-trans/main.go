package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	// Krasovsky 1940
	a  = 6378245.0
	ee = 0.00669342162296594323
)

type location struct {
	lon float64
	lat float64
}

func main() {

	b, err := os.ReadFile("./raw.json")
	if err != nil {
		panic(err)
	}
	jsonStr := string(b)
	fList := gjson.Get(jsonStr, "features")
	res := make([]string, 0)
	fList.ForEach(func(i, f gjson.Result) bool {

		fmt.Println("i", i.Int())
		locList := make([]location, f.Get("geometry.coordinates.#").Int())
		for i, v := range f.Get("geometry.coordinates").Array() {

			lon := v.Get("0").Float()
			lat := v.Get("1").Float()
			locList[i] = location{lon: lon, lat: lat}
			fmt.Printf("lon: %f , lat: %f\n", lon, lat)

		}
		locsRes := Wgs2gcjs(locList)
		target := make([][2]float64, len(locsRes))
		for i, loc := range locsRes {
			target[i] = [2]float64{loc.lon, loc.lat}
		}

		tar, err := sjson.Set(jsonStr, fmt.Sprintf("features.%d.geometry.coordinates", i.Int()), target)
		if err != nil {
			log.Fatalf("failed to set json: %v", err)
		}
		res = append(res, tar)
		fmt.Printf("from %v  \n to %v\n", f.String(), tar)
		return true
	})
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("failed to set json: %v", err)
	}
	if err := os.WriteFile("./out.json", jsonBytes, 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}

}

func Wgs2gcjs(locsReq []location) (locsRes []location) {
	for _, loc := range locsReq {
		gcjLat, gcjLon, err := wgs2gcj(loc.lon, loc.lat)
		if err != nil {
			fmt.Println(err)
			continue
		}
		locsRes = append(locsRes, location{lon: gcjLon, lat: gcjLat})
	}
	return
}

func wgs2gcj(lon, lat float64) (gcjLat, gcjLon float64, err error) {
	if outOfChina(lat, lon) {

		log.Println("=================out of China   LON LAT", lon, lat)
		return 0, 0, errors.New("out of China")
	}
	dLat := transformLat(lat-35.0, lon-105.0)
	dLon := transformLon(lat-35.0, lon-105.0)
	rLat := lat / 180.0 * math.Pi
	magic := math.Sin(rLat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * math.Pi)
	dLon = (dLon * 180.0) / (a / sqrtMagic * math.Cos(rLat) * math.Pi)
	gcjLat = lat + dLat
	gcjLon = lon + dLon
	return
}

func outOfChina(lat, lon float64) bool {
	if lat < 0.8293 || lat > 55.8271 {
		return true
	}
	if lon < 72.004 || lon > 137.8347 {
		return true
	}
	return false
}

func transformLat(lat, lon float64) float64 {
	ret := -100.0 + 2.0*lon + 3.0*lat + 0.2*lat*lat + 0.1*lon*lat + 0.2*math.Sqrt(math.Abs(lon))
	ret += (20.0*math.Sin(6.0*lon*math.Pi) + 20.0*math.Sin(2.0*lon*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lat*math.Pi) + 40.0*math.Sin(lat/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(lat/12.0*math.Pi) + 320*math.Sin(lat*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLon(lat, lon float64) float64 {
	ret := 300.0 + lon + 2.0*lat + 0.1*lon*lon + 0.1*lon*lat + 0.1*math.Sqrt(math.Abs(lon))
	ret += (20.0*math.Sin(6.0*lon*math.Pi) + 20.0*math.Sin(2.0*lon*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lon*math.Pi) + 40.0*math.Sin(lon/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(lon/12.0*math.Pi) + 300.0*math.Sin(lon/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}
