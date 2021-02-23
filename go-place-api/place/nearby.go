package place

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"googlemaps.github.io/maps"
)

// SearchNearby call geometry and nearbyResult
func SearchNearby(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
	}

	p := r.FormValue("name")
	if p == "" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}
	//fmt.Fprintln(w, p)
	fmt.Println(p)

	// search place to find geometry location
	latlng := geometry(w, r, p)
	//fmt.Fprintln(w, latlng)
	fmt.Println(latlng)

	// take geometry location to search nearby its location
	result := nearbyResult(w, r, latlng)

	w.Header().Set("Conten-type", "application/json")
	//w.Write(result)
	fmt.Fprintf(w, "%s", result)
}

func nearbyResult(_ http.ResponseWriter, _ *http.Request, latlng maps.LatLng) []byte {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	check(err)

	nsr := &maps.NearbySearchRequest{
		Radius: uint(radius),
	}
	nsr.Location = &latlng
	parsePlaceType(placeType, nsr)

	resp, err := c.NearbySearch(context.Background(), nsr)
	check(err)

	j, err := json.Marshal(resp)
	check(err)
	return j
}

func parsePlaceType(placeType string, nsr *maps.NearbySearchRequest) {
	if placeType != "" {
		t, err := maps.ParsePlaceType(placeType)
		check(err)
		nsr.Type = t
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
