package distance

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// mockedup from ONS API responses
type GeoBodyData struct {
	Postcode 					string		`json:"postcode"`
	Status						string		`json:"status"`
	Usertype					string		`json:"usertype"`
	Easting						int			`json:"easting"`
	Northing					int			`json:"northing"`
	PositionalQualityIndicator 	int 		`json:"positional_quality_indicator"`
	Country						string		`json:"country"`
	Latitude					string		`json:"latitude"`
	Longitude					string		`json:"longitude"`
	PostcodeNoSpace				string		`json:"postcode_no_space"`
	PostcodeFixedWidthSeven 	string		`json:"postcode_fixed_width_seven"`
	PostcodeFixedWidthEight		string		`json:"postcode_fixed_width_eight"`
	PostcodeArea				string		`json:"postcode_area"`
	PostcodeDistrict			string		`json:"postcode_district"`
	PostcodeSector				string		`json:"postcode_sector"`
	Outcode						string		`json:"outcode"`
	Incode						string		`json:"incode"`
}

type GeoBody struct {
	Status 	   string		`json:"status"`
	MatchType string 		`json:"match_type"`
	Input      string		`json:"input"`
	Data	   GeoBodyData 	`json:"data"`
	Copyright  []string		`json:"copyright"`
}

// basic http post request
func getLongLatCoordinates(point string) (GeoLocation, error) {

	var postcode string

	// simple data cleaning to remove spaces
	for i:= 0; i < len(point); i++ {
		if point[i] != ' ' {
			postcode += fmt.Sprintf("%c",point[i])
		}	
	}

	httpUrl := "http://api.getthedata.com/postcode/"+postcode

	response, err := http.Get(httpUrl)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		
		return GeoLocation{}, err
	}
	
	var geoBody GeoBody
    err = json.Unmarshal([]byte(body), &geoBody)
	
	if err != nil {
		return GeoLocation{}, err
	}
		
	// just pull out the fields we want to use
	var geoLocation GeoLocation
	geoLocation.Longitude,_ = strconv.ParseFloat(geoBody.Data.Longitude,64) // string to 64 bit float
	geoLocation.Latitude, _ = strconv.ParseFloat(geoBody.Data.Latitude,64)
 
	return geoLocation, nil
}