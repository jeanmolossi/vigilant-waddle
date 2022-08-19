package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jeanmolossi/vigilant-waddle/src/infra/http"
)

func Example() {
	m := http.NewMeta(
		http.WithBaseURL("http://example.com/api/resource"),
		http.WithPage(2),
		http.WithPageSize(10),
		http.WithItems(make([]interface{}, 10)),
	)

	fmt.Println(toJson(m))
	// Output:
	// {
	// 	"page": 2,
	// 	"items_per_page": 10,
	// 	"next_page": "http://example.com/api/resource?page=3&items_per_page=10",
	// 	"prev_page": "http://example.com/api/resource?page=3&items_per_page=10",
	// 	"total_count": 10
	// }
}

// toJson is a helper function to convert a struct to json
//
// It will return the struct converted as json string
func toJson(m *http.HttpMeta) string {
	b := new(bytes.Buffer)
	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	encoder.Encode(m)

	return b.String()
}
