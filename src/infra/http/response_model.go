// Package http has the meta data for the http response.
//
// What is the meta data?
//
// Meta data is used to define the pagination for the response
// It contains the meta data format and the builder to create the meta data
package http

import (
	"fmt"
	"os"
)

// HttpMeta is the meta data for http response
type HttpMeta struct {
	Page       uint16 `json:"page" example:"2"`
	PageSize   int    `json:"items_per_page" example:"10"`
	NextPage   string `json:"next_page,omitempty" example:"http://example.com/api/resource?page=3&items_per_page=10"`
	PrevPage   string `json:"prev_page,omitempty" example:"http://example.com/api/resource?page=1&items_per_page=10"`
	TotalCount int    `json:"total_count" example:"10"`
	// baseURL is a private prop here but it doesn't have any effect on the meta data
	// because it is not used in the meta data
	// It is used to mount pagination links
	baseURL string `json:"-"`
}

// MetaOption is a function that will be called on the meta data
//
// Used to create options for the meta data like
// 	WithBaseURL, WithPage, WithPageSize, WithItems
//
// Example:
//
// 		// WithBaseURL will define the base url for the meta data
//		func WithBaseURL(baseURL string) MetaOptions {
//			return func(m *HttpMeta) {
//				m.baseURL = baseURL
//			}
//		}
type MetaOption func(*HttpMeta)

// NewMeta creates a new meta data
//
// Look at the example for more details
func NewMeta(opts ...MetaOption) *HttpMeta {
	m := &HttpMeta{
		Page:       1,
		PageSize:   10,
		NextPage:   "",
		PrevPage:   "",
		TotalCount: 0,
		baseURL:    "",
	}

	for _, opt := range opts {
		opt(m)
	}

	if m.baseURL == "" {
		m.baseURL = os.Getenv("API_HOST")
	}

	m.mountPrevPage()
	m.mountNextPage(m.TotalCount)

	return m
}

func (m *HttpMeta) mountPrevPage() {
	if m.Page > 1 {
		m.PrevPage = baseLink(m.baseURL, m.Page+1, m.PageSize)
	}
}

func (m *HttpMeta) mountNextPage(totalCount int) {
	if totalCount >= m.PageSize {
		m.NextPage = baseLink(m.baseURL, m.Page+1, m.PageSize)
	}
}

func baseLink(baseURL string, page uint16, pageSize int) string {
	return fmt.Sprintf("%s?page=%d&items_per_page=%d", baseURL, page, pageSize)
}
