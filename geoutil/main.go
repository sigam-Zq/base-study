package main

func main() {
	// Example usage of IsPointInPolygon
	polygon := []Point{{117.141588, 36.661758}, {117.146921, 36.663282}, {117.151975, 36.664917}, {117.153651, 36.662005}, {117.155131, 36.659026}, {117.150384, 36.65701}, {117.145022, 36.655643}, {117.141671, 36.661646}}
	pointT := []Point{{117.145278, 36.657762}, {117.146867, 36.660175}, {117.148556, 36.660935}}
	tolerance := 500.0 // in meters

	for _, v := range pointT {
		wgs84Point := GCJ02ToWGS84(v[0], v[1])
		inPolygon := IsPointInPolygon(polygon, wgs84Point, tolerance)
		println("Is point in polygon:", inPolygon, wgs84Point[0], wgs84Point[1])
	}
}
