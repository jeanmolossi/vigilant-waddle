package http

// WithBaseURL will define the base url for the meta data
func WithBaseURL(baseURL string) MetaOption {
	return func(m *HttpMeta) {
		m.baseURL = baseURL
	}
}

// WithPage will define the page for the meta data
func WithPage(page uint16) MetaOption {
	return func(m *HttpMeta) {
		if page == 0 {
			page = 1
		}

		m.Page = page
	}
}

// WithPageSize will define the page size for the meta data
// Also known as items per page
func WithPageSize(pageSize int) MetaOption {
	return func(m *HttpMeta) {
		if pageSize == 0 {
			pageSize = 10
		}

		m.PageSize = pageSize
	}
}

// WithItems will define the total items for the meta data
func WithItems(items []interface{}) MetaOption {
	return func(m *HttpMeta) {
		m.TotalCount = len(items)
	}
}
