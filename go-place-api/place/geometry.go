package place

import (
	"context"
	"net/http"
	"strings"

	"googlemaps.github.io/maps"
)

// geometry call FindPlaceFromText for find its location
func geometry(_ http.ResponseWriter, _ *http.Request, nPlace string) maps.LatLng {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	check(err)

	fp := &maps.FindPlaceFromTextRequest{
		Input:     nPlace,
		InputType: parseInputType(inputType),
	}

	f, err := parseFields("geometry")
	fp.Fields = f

	resp, err := c.FindPlaceFromText(context.Background(), fp)
	check(err)

	return resp.Candidates[0].Geometry.Location
}

func parseInputType(inputType string) maps.FindPlaceFromTextInputType {
	var it maps.FindPlaceFromTextInputType
	switch inputType {
	case "textquery":
		it = maps.FindPlaceFromTextInputTypeTextQuery
	case "phonenumber":
		it = maps.FindPlaceFromTextInputTypePhoneNumber
	default:
		it = maps.FindPlaceFromTextInputTypeTextQuery
	}
	return it
}

func parseFields(fields string) ([]maps.PlaceSearchFieldMask, error) {
	var res []maps.PlaceSearchFieldMask
	for _, s := range strings.Split(fields, ",") {
		f, err := maps.ParsePlaceSearchFieldMask(s)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
	}
	return res, nil
}
