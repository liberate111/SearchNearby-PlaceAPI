package place

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type result struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// FindValue show result XYZ
func FindValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
	}
	var re result
	re.X, re.Y, re.Z = findXYZ("X, 5, 9, 15, 23, Y, Z", 7)
	j, err := json.Marshal(re)
	check(err)
	w.Header().Set("Conten-type", "application/json")
	fmt.Fprintf(w, "%s", j)
}

func generate(n int) []int {
	result := []int{}
	if n > 0 {
		// T(n) = 3 + n(n-1)
		for i := 1; i <= n; i++ {
			t := 3 + i*(i-1)
			result = append(result, t)
		}
		return result
	}
	return result
}

func findXYZ(s string, n int) (int, int, int) {
	var x, y, z int
	series := generate(n)
	sl := strings.Split(s, ",")

	for i, v := range sl {
		v = strings.TrimSpace(v)
		if v == "X" {
			x = series[i]
		}
		if v == "Y" {
			y = series[i]
		}
		if v == "Z" {
			z = series[i]
		}
	}
	return x, y, z
}
