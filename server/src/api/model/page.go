package model

type Pageable interface {
	GetOffset() int
	GetPageSize() int
	GetPageNumber() int
}

type pageable struct {
	page int
	size int
}

func (p *pageable) GetOffset() int {
	return int((p.page - 1) * p.size)
}

func (p *pageable) GetPageSize() int {
	return int(p.size)
}

func (p *pageable) GetPageNumber() int {
	return int(p.page)
}

func NewPageable(page int, size int) Pageable {
	if page == 0 {
		page = 1
	}
	return &pageable{
		page: page,
		size: size,
	}
}

type Page[T any] interface {
	GetContents() []T
	GetTotalCount() int
	GetPageNumber() int
	GetPageSize() int
}

type page[T any] struct {
	values     []T
	total      int
	pageNumber int
	size       int
}

func (p *page[T]) GetContents() []T {
	return p.values
}

func (p *page[T]) GetTotalCount() int {
	return int(p.total)
}

func (p *page[T]) GetPageNumber() int {
	return int(p.pageNumber)
}

func (p *page[T]) GetPageSize() int {
	return int(p.size)
}

func NewPage[T any](values []T, total int, pageNumber int, size int) Page[T] {
	return &page[T]{
		values:     values,
		total:      total,
		pageNumber: pageNumber,
		size:       size,
	}
}

type StudentPage = Page[*Student]
