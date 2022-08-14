package paginator_test

import (
	"fmt"

	"github.com/jeanmolossi/vigilant-waddle/src/pkg/paginator"
)

func Example() {
	p := paginator.New(
		paginator.WithBaseURL("https://example.com/article/%s", "article-slug"),
		paginator.WithPage(2),
		paginator.WithItemsPerPage(15),
	)

	fmt.Println(p.Page(), p.ShouldPaginate(), p.GetItemsPerPage())

	// Output:
	// 2 true 15
}
