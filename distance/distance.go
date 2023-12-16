package distance

import (
	"math"
	"fmt"
)

type GeoLocation struct {
	Latitude, Longitude float64
}

func GeoDistance(point1, point2 string) float64 {

	var err error

	// convert postcode into long/lat
	geopoint1, err := getLongLatCoordinates(point1)

	if err != nil {
		fmt.Printf ("Problem reading geolocation of '%s'.\n", point1)
		return -1.0		
	}
	
	geopoint2, err := getLongLatCoordinates(point2)

	if err != nil {
		fmt.Printf ("Problem reading geolocation of '%s'.\n", point2)
		return -1.0
	}

	// all is well
	theta := geopoint1.Longitude - geopoint2.Longitude
	
	// convert from degrees into radians
	thetaInRadians := math.Pi * theta / 180.0
	latitude1InRadians := math.Pi * geopoint1.Latitude / 180.0
	latitude2InRadians := math.Pi * geopoint2.Latitude / 180.0
	
	// great circles calculation
	distance := math.Sin(latitude1InRadians) * math.Sin(latitude2InRadians) + math.Cos(latitude1InRadians) * math.Cos(latitude2InRadians) * math.Cos(thetaInRadians);
	if distance > 1 {
		distance = 1
	}
	
	distance = math.Acos(distance)
	distance = (180.0 * distance)/ math.Pi
	distance = (60.0 * distance) * 1.1515
	
	return distance
}

