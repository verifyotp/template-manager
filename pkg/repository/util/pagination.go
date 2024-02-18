package util

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type paginatorJoin struct {
	query string
	args  []any
}

type PaginatorT[T any] struct {
	Error     error
	db        *gorm.DB
	model     any
	tableName string // Table name to use for the query

	isRaw        bool // is raw query
	disableCount bool // disable count query

	query     any   // query
	queryArgs []any //args

	joins     []paginatorJoin
	selection []string
	group     []string
}

// Default database field for sorting and pagination
const DefaultDatabaseSortField = "created_at"

// Default database sort order
const DefaultDatabaseSortOrder = "desc"

// Default database page size. Default to 10
const DefaultDatabasePageSize = 10

// let T be an array of any type
type PaginationT[T any] struct {
	Page        int   `json:"page"`         // current page
	PageSize    int   `json:"page_size"`    // number of records to return per page
	PageCount   int   `json:"page_count"`   // number of pages available in the database
	RecordCount int64 `json:"record_count"` // number of records available in the database
	Data        T     `json:"data"`         // data to return
}

// Create a new Paginator instance
func NewPaginatorT[T any](db *gorm.DB) *PaginatorT[T] {
	var model T
	return &PaginatorT[T]{db: db, Error: nil, query: nil, queryArgs: nil, model: &model}
}

type orderSetting struct {
	sortField string
	sortOrder string
}

type findSetting struct {
	Order   []orderSetting
	Preload []string
}

type FindOption func(*findSetting)

/*
For multiple find queries, we can use the WithSort function to specify the sort order for each column.

For example, if we want to sort the result set by column1 in ascending order, and then sort by column2 in descending order, we can do the following:

paginator.Find(&pageDetails, 1, 10, WithSortAsc("column1"), WithSortDesc("column2"))

this will generate the following SQL query:

```

SELECT * FROM table
ORDER BY column1 ASC, column2 DESC NULLS LAST
LIMIT 10 OFFSET 0;

```

Here's how it works:

1. The query first sorts the result set by column1 in ascending order (ASC).

2. For rows where column1 values are the same, it will then sort those rows by column2 in descending order (DESC).

note: column2 DESC will only affect the order of rows with the same values in column1
*/
func WithSort(sortField string, sortOrder string) FindOption {
	return func(s *findSetting) {
		s.Order = append(s.Order, orderSetting{sortField: sortField, sortOrder: sortOrder})
	}
}

/*
For multiple find queries, we can use the WithSort function to specify the sort order for each column.

For example, if we want to sort the result set by column1 in ascending order, and then sort by column2 in descending order, we can do the following:

paginator.Find(&pageDetails, 1, 10, WithSortAsc("column1"), WithSortDesc("column2"))

this will generate the following SQL query:

```

SELECT * FROM table
ORDER BY column1 ASC, column2 DESC NULLS LAST
LIMIT 10 OFFSET 0;

```

Here's how it works:

1. The query first sorts the result set by column1 in ascending order (ASC).

2. For rows where column1 values are the same, it will then sort those rows by column2 in descending order (DESC).

note: column2 DESC will only affect the order of rows with the same values in column1
*/
func WithSortAsc(sortField string) FindOption {
	return WithSort(sortField, "asc")
}

/*
For multiple find queries, we can use the WithSort function to specify the sort order for each column.

For example, if we want to sort the result set by column1 in ascending order, and then sort by column2 in descending order, we can do the following:

paginator.Find(&pageDetails, 1, 10, WithSortAsc("column1"), WithSortDesc("column2"))

this will generate the following SQL query:

```

SELECT * FROM table
ORDER BY column1 ASC, column2 DESC NULLS LAST
LIMIT 10 OFFSET 0;

```

Here's how it works:

1. The query first sorts the result set by column1 in ascending order (ASC).

2. For rows where column1 values are the same, it will then sort those rows by column2 in descending order (DESC).
*/
func WithSortDesc(sortField string) FindOption {
	return WithSort(sortField, "desc")
}

func WithPreload(preloads ...string) FindOption {
	return func(s *findSetting) {
		s.Preload = append(s.Preload, preloads...)
	}
}

// Find database record and paginate based on the provided  sortOptions parameters
// If optional sortOptions are not provided, the system uses the default sortOptions
//
//	DefaultDatabaseSortField and  DefaultDatabaseSortOrder for field and sorting order
//
// db -  DB instance to
// page - the size of record to return
// pageSize - the size of record page to return
// result - the database find result
// sortOptions - the pagination options
//
//	sortOptions[0]- the sortOrder . Could be `asc` or `desc`
//	sortOptions[1]- the filed to user for sorting. Defaults to `date_created`
func (p *PaginatorT[T]) Find(pageDetails *PaginationT[[]T], page int, pageSize int, opts ...FindOption) *PaginatorT[T] {

	var err error
	var pageCount int
	var recordCount int64
	fetchedData := make([]T, 0, pageSize)
	var setting = new(findSetting)
	for _, opt := range opts {
		opt(setting)
	}

	if len(setting.Order) < 1 {
		setting.Order = []orderSetting{{sortField: DefaultDatabaseSortField, sortOrder: DefaultDatabaseSortOrder}}
	}
	model := p.db.Model(&p.model)
	if p.tableName != "" {
		model = model.Table(p.tableName)
	}

	if len(p.selection) > 0 {
		model = model.Select(p.selection)
	}
	// Add the custom query provided via the package Where method. This can be use to further
	// fine-tune selection without using database filter implementation.
	if p.query != nil {
		if p.isRaw {
			model = model.Raw(p.query.(string), p.queryArgs...)
		} else if p.queryArgs != nil {
			model = model.Where(p.query, p.queryArgs...)
		} else {
			model = model.Where(p.query)
		}
	}

	for _, join := range p.joins {
		model = model.Joins(join.query, join.args...)
	}

	for _, preload := range setting.Preload {
		model = model.Preload(preload)
	}

	for _, group := range p.group {
		model = model.Group(group)
	}

	// count the total number of pages record could fetch based on the total selection possible
	if !p.disableCount { // if
		model = model.Count(&recordCount)
		pageCount = CalculatePageCount(recordCount, pageSize)
	}
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		model = model.Offset(offset).Limit(pageSize)
	}
	for _, sort := range setting.Order {
		model = model.Order(fmt.Sprintf("%s %s NULLS LAST", sort.sortField, sort.sortOrder))
	}

	err = model.Find(&fetchedData).Error
	if err != nil {
		p.Error = err
	} else {
		p.clearError()
	}
	// Assign additional page information to pageDetails
	pageDetails.Page = page
	pageDetails.PageSize = pageSize
	pageDetails.PageCount = pageCount
	pageDetails.PageSize = pageSize
	pageDetails.RecordCount = recordCount
	pageDetails.Data = fetchedData

	p.clearQuery()
	return p
}

func CalculatePageCount(recordCount int64, pageSize int) int {
	var pageCount int
	if recordCount > 0 {
		if recordCount <= int64(pageSize) {
			pageCount = 1
		} else {
			if pageSize > 0 {
				pages := float64(recordCount) / float64(pageSize)
				pageCount = int(math.Ceil(pages))
			} else {
				// if not page sixing provided, we will assume pages size is equivalent to
				// record size since we cannot reasonably know the page size. This will happen
				// when we are not performing pagination but called the paginator find method
				pageCount = int(recordCount)
			}
		}
	}
	return pageCount
}

func (p PaginatorT[T]) Copy() *PaginatorT[T] {
	return &p
}

func (p *PaginatorT[T]) Count(pageSize int) (pageCount int, recordCount int64, err error) {

	model := p.db.Model(&p.model)
	if p.tableName != "" {
		model = model.Table(p.tableName)
	}

	if len(p.selection) > 0 {
		model = model.Select(p.selection)
	}

	// Add the custom query provided via the package Where method. This can be use to further
	// fine-tune selection without using database filter implementation.
	if p.query != nil {
		if p.isRaw {
			model = model.Raw(p.query.(string), p.queryArgs...)
		} else {
			model = model.Where(p.query, p.queryArgs...)
		}
	}

	for _, join := range p.joins {
		model = model.Joins(join.query, join.args...)
	}

	for _, group := range p.group {
		model = model.Group(group)
	}

	// count the total number of pages record could fetch based on the total selection possible
	err = model.Count(&recordCount).Error
	pageCount = CalculatePageCount(recordCount, pageSize)
	return pageCount, recordCount, err
}

func (p *PaginatorT[T]) DisableCount() *PaginatorT[T] {
	p.disableCount = true
	return p
}

func (p *PaginatorT[T]) Select(selection ...string) *PaginatorT[T] {
	p.selection = append(p.selection, selection...)
	return p
}

func (p *PaginatorT[T]) Group(group ...string) *PaginatorT[T] {
	p.group = append(p.group, group...)
	return p
}

func (p *PaginatorT[T]) Joins(query string, args ...any) *PaginatorT[T] {
	p.joins = append(p.joins, paginatorJoin{query: query, args: args})
	return p
}

func (p *PaginatorT[T]) Model(modelPtr any) *PaginatorT[T] {
	p.model = modelPtr
	return p
}

func (p *PaginatorT[T]) Raw(query string, args ...any) *PaginatorT[T] {
	p.isRaw = true
	p.query = query
	if len(args) > 0 {
		p.queryArgs = args
	} else {
		p.clearQuery()
	}
	return p
}

func (p *PaginatorT[T]) Where(query any, args ...any) *PaginatorT[T] {
	p.query = query
	if len(args) > 0 {
		p.queryArgs = args
	} else {
		p.clearQuery()
	}
	return p
}

func (p *PaginatorT[T]) WhereNoArgs(query any) *PaginatorT[T] {
	p.query = query
	return p
}

func (p *PaginatorT[T]) SetTableName(tableName string) *PaginatorT[T] {
	p.tableName = tableName
	return p
}

// clear query after each call
func (p *PaginatorT[T]) clearQuery() {
	p.query = nil
	p.queryArgs = nil
}

// Clear error value
func (p *PaginatorT[T]) clearError() {
	p.Error = nil
}
