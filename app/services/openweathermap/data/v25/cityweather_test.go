package cityweather

import (
	"errors"
	"testing"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gf.v1/mock/gophermocks"
	"git.target.com/gophersaurus/gf.v1/mock/mockstar"
)

func Test_Find(t *testing.T) {
	mockstar.T = t

	mockHTTP := gophermocks.NewMockClient()
	gf.HTTP = mockHTTP

	city := "This_is_a_city"
	country := "This_is_the_country"
	path := "http://api.openweathermap.org/data/2.5/weather?q="

	url := path + city + "," + country
	json := &gf.HTTPResponse{Body: []byte("{\"name\":\"result\",\"main\":{\"temp\":78.2,\"temp_max\":84.6,\"temp_min\":42.3},\"sys\":{\"sunrise\":15,\"sunset\":64651764}}")}

	intended := &Result{
		Name: "result",
		Main: Main{
			Temp:    78.2,
			TempMax: 84.6,
			TempMin: 42.3,
		},
		Sys: Sys{
			Sunrise: 15,
			Sunset:  64651764,
		},
	}

	mockHTTP.When("Get", url).Return(json, nil)

	result, err := Find(city, country)

	mockstar.Expect(result).ToEqual(intended)
	mockstar.Expect(err).ToBeNil()
}

func Test_Find_Err(t *testing.T) {
	mockstar.T = t

	mockHTTP := gophermocks.NewMockClient()
	gf.HTTP = mockHTTP

	city := "This_is_a_city"
	country := "This_is_the_country"
	path := "http://api.openweathermap.org/data/2.5/weather?q="

	url := path + city + "," + country

	expectedErr := errors.New("this is an error")

	mockHTTP.When("Get", url).Return(nil, expectedErr)

	_, err := Find(city, country)

	mockstar.Expect(err).ToEqual(expectedErr)
}
