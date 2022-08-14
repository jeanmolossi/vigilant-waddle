package paginator

import "gorm.io/gorm"

type paginator struct {
	baseURL      string
	page         uint16
	itemsPerPage int
}

// NewPaginator creates a new paginator.
// It accepts a variadic list of options to configure the paginator.
func NewPaginator(opts ...Option) Paginator {
	p := &paginator{}
	p.buildOptions(opts...)
	return p
}

// Page returns the current page.
func (p *paginator) Page() uint16 {
	return p.page
}

// ShouldPaginate returns true if pagination is needed.
func (p *paginator) ShouldPaginate() bool {
	return p.page > 1 || p.itemsPerPage != 10
}

// GetItemsPerPage returns the number of items per page.
func (p *paginator) GetItemsPerPage() int {
	return p.itemsPerPage
}

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
func (p *paginator) Paginate(db *gorm.DB) *gorm.DB {
	offset := int(p.getOffset())
	return db.Offset(offset).Limit(p.GetItemsPerPage())
}

// getOffset returns the offset of the current page.
func (p *paginator) getOffset() uint16 {
	return uint16(p.itemsPerPage) * (p.page - 1)
}

// buildOptions applies the given options to the paginator.
func (p *paginator) buildOptions(opts ...Option) error {
	if len(opts) == 0 {
		panic("no options provided")
	}

	for _, opt := range opts {
		if err := opt(p); err != nil {
			return err
		}
	}

	return nil
}
