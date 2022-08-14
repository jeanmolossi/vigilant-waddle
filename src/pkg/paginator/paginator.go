// Package paginator is a helper package for paginating results.
//
// It's made for easy paginate with https://gorm.io
package paginator

import "gorm.io/gorm"

// Option is a function that can be passed to NewPaginator to configure the paginator.
type Option func(p *paginator) error

// Paginator is the interface that wraps the basic Paginate methods.
type Paginator interface {
	// Page returns the current page.
	Page() uint16
	// ShouldPaginate returns true if pagination is needed.
	ShouldPaginate() bool
	// GetItemsPerPage returns the number of items per page.
	GetItemsPerPage() int
	// Paginate is a gorm pagination scope implementation.
	// It returns a gorm.DB with pagination applied.
	//
	// Example:
	//
	//      pagination := paginator.NewPaginator(
	//          paginator.WithPage(2),
	//      )
	//
	//      if pagination.ShouldPaginate() {
	//          gormDB = gormDB.Scopes(pagination.Paginate)
	//      }
	Paginate(db *gorm.DB) *gorm.DB
	// getOffset returns the offset of the current page.
	getOffset() uint16
}
