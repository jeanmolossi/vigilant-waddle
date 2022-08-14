package paginator

import "fmt"

func WithBaseURL(format string, args ...interface{}) Option {
	return func(p *paginator) error {
		if format != "" {
			p.baseURL = fmt.Sprintf(format, args...)
		}

		return nil
	}
}

func WithPage(page uint16) Option {
	return func(p *paginator) error {
		if page == 0 {
			page = 1
		}

		if page >= 65535 {
			page = 65535
		}

		p.page = page
		return nil
	}
}

func WithItemsPerPage(itemsPerPage int) Option {
	return func(p *paginator) error {
		if itemsPerPage == 0 {
			itemsPerPage = 10
		}

		if itemsPerPage > 1000 {
			itemsPerPage = 1000
		}

		p.itemsPerPage = itemsPerPage
		return nil
	}
}
