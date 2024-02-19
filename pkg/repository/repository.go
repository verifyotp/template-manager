package repository

import (
	"context"
	"strings"

	"template-manager/pkg/repository/util"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository[T any] interface {
	Count(ctx context.Context, query any) (count int64, err error)
	Create(ctx context.Context, t *T) error
	CreateMany(ctx context.Context, data []T, opts ...CreateOpt) error
	Delete(ctx context.Context, t *T) error
	DeleteByFieldName(ctx context.Context, query any) error
	Find(ctx context.Context, conds ...interface{}) ([]T, error)
	FindByFieldName(ctx context.Context, query any) (*T, error)
	FindByFieldNameWithPreload(ctx context.Context, query any, preloads ...string) (*T, error)
	FindMany(ctx context.Context, query any, page int, pageSize int, preloads ...string) ([]T, error)
	FindManyWithOptions(ctx context.Context, query any, opts ...Opt) ([]T, error)
	FindManyWithOrder(ctx context.Context, query any, page int, pageSize int, order string, preloads ...string) ([]T, error)
	FindWithPagination(ctx context.Context, query any, opts ...Opt) (*util.PaginationT[[]T], error)
	Get(ctx context.Context, conds ...interface{}) (*T, error)
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	Update(ctx context.Context, E *T) error
	UpdateMany(ctx context.Context, query any, data any) error
}

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *repository[T] {
	var model T
	err := db.AutoMigrate(model)
	if err != nil {
		return nil
	}
	return &repository[T]{db: db}
}

var _ Repository[any] = (*repository[any])(nil)

func (r *repository[T]) Create(ctx context.Context, t *T) error {
	return r.db.WithContext(ctx).Create(t).Error

}

func (r *repository[T]) Find(ctx context.Context, conds ...interface{}) ([]T, error) {
	var (
		dest []T
	)

	if err := r.db.WithContext(ctx).Find(&dest, conds...).Error; err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *repository[T]) Get(ctx context.Context, conds ...interface{}) (*T, error) {
	var (
		dest T
	)

	if err := r.db.WithContext(ctx).Model(&dest).First(&dest, conds...).Error; err != nil {
		return nil, err
	}

	return &dest, nil
}

func (r *repository[T]) Update(ctx context.Context, E *T) error {
	return r.db.WithContext(ctx).Updates(E).Error
}

func (r *repository[T]) Delete(ctx context.Context, t *T) error {
	if err := r.db.WithContext(ctx).Delete(t).Error; err != nil {
		return err
	}
	return nil
}

type findManyOptions struct {
	Page     int
	PageSize int
	Preloads []string
	OrderBy  string
	Order    string
}

// applyFilter applies the filter to the query
//
//	applyFilter(db).Find(&result)
func (f findManyOptions) applyFilter(tx *gorm.DB) *gorm.DB {
	if f.isPaginationSet() {
		tx = tx.Limit(f.PageSize).Offset((f.Page - 1) * f.PageSize)
	}
	if f.isSortSet() {
		tx = tx.Order(f.OrderBy + " " + f.Order)
	}
	if f.isPreloadSet() {
		tx = tx.Set("gorm:auto_preload", true)
		for _, preload := range f.Preloads {
			tx = tx.Preload(preload)
		}
	}
	return tx
}

func (f findManyOptions) isSortSet() bool {
	return f.OrderBy != "" && f.Order != ""
}

func (f findManyOptions) isPaginationSet() bool {
	return f.Page != 0 && f.PageSize != 0
}

func (f findManyOptions) isPreloadSet() bool {
	return len(f.Preloads) > 0
}

type Opt func(*findManyOptions)

// WithPagination is used to paginate the result
//
//	WithPagination(1, 10)
func WithPagination(page, pageSize int) Opt {
	return func(o *findManyOptions) {
		o.Page = page
		o.PageSize = pageSize
	}
}

// WithPreloads is used to preload the result
//
//	WithPreloads("user", "user.profile")
func WithPreloads(preloads ...string) Opt {
	return func(o *findManyOptions) {
		o.Preloads = preloads
	}
}

// WithOrderBy is used to order the result by a field
//
//	WithOrderBy("created_at", "desc") // order by created_at desc
func WithOrderBy(orderBy string, order string) Opt {
	return func(o *findManyOptions) {
		o.OrderBy = orderBy
		o.Order = strings.ToUpper(order)
	}
}

type createOptions struct {
	clause []clause.Expression
}

func (c createOptions) applyClause(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(c.clause...)
}

type CreateOpt func(*createOptions)

func WithClause(clause ...clause.Expression) CreateOpt {
	return func(o *createOptions) {
		o.clause = clause
	}
}

func (r repository[T]) FindManyWithOptions(ctx context.Context, query any, opts ...Opt) ([]T, error) {
	var result []T
	var Model T

	var opt findManyOptions
	for _, o := range opts { // apply the options
		o(&opt)
	}
	db := r.db.Model(&Model)
	if q, ok := query.(util.Query); ok {
		db = db.Where(q.Query, q.Args...)
	} else {
		db = db.Where(query)
	}
	db = opt.applyFilter(db) // apply the filter
	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindWithPagination is used to find many with pagination
//
//	FindWithPagination(ctx, query,
//				 		WithPagination(1, 10),
//						WithPreloads("user", "user.profile"),
//						WithOrderBy("created_at", "desc"))
func (r repository[T]) FindWithPagination(ctx context.Context, query any, opts ...Opt) (*util.PaginationT[[]T], error) {
	var opt findManyOptions
	for _, o := range opts { // apply the options
		o(&opt)
	}
	var pgnOpts []util.FindOption
	if opt.isSortSet() {
		pgnOpts = append(pgnOpts, util.WithSort(opt.OrderBy, opt.Order))
	}
	if opt.isPreloadSet() {
		pgnOpts = append(pgnOpts, util.WithPreload(opt.Preloads...))
	}

	pgn := util.NewPaginatorT[T](r.db)
	var result util.PaginationT[[]T]
	if q, ok := query.(util.Query); ok {
		pgn = pgn.Where(q.Query, q.Args...)
	} else {
		pgn = pgn.WhereNoArgs(query)
	}

	err := pgn.Find(&result, opt.Page, opt.PageSize, pgnOpts...).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r repository[T]) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r repository[T]) FindMany(ctx context.Context, query any, page, pageSize int, preloads ...string) ([]T, error) {

	var result []T
	db := r.db.Set("gorm:auto_preload", true)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	if q, ok := query.(util.Query); ok {
		db = db.Where(q.Query, q.Args...)
	} else {
		db = db.Where(query)
	}

	err := db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository[T]) FindManyWithOrder(ctx context.Context, query any, page, pageSize int, order string, preloads ...string) ([]T, error) {

	var result []T
	db := r.db.Set("gorm:auto_preload", true)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	if order != "" {
		db = db.Order(order)
	}

	if q, ok := query.(util.Query); ok {
		db = db.Where(q.Query, q.Args...)
	} else {
		db = db.Where(query)
	}

	err := db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository[T]) FindByFieldName(ctx context.Context, query any) (*T, error) {
	var result T

	db := r.db.Set("gorm:auto_preload", true)

	if q, ok := query.(util.Query); ok {
		db = db.Where(q.Query, q.Args...)
	} else {
		db = db.Where(query)
	}

	err := db.First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository[T]) FindByFieldNameWithPreload(ctx context.Context, query any, preloads ...string) (*T, error) {

	var result T
	db := r.db.Set("gorm:auto_preload", true)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	if q, ok := query.(util.Query); ok {
		db = db.Where(q.Query, q.Args...)
	} else {
		db = db.Where(query)
	}

	err := db.First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository[T]) CreateMany(ctx context.Context, data []T, opts ...CreateOpt) error {
	var opt createOptions
	for _, o := range opts { // apply the options
		o(&opt)
	}
	db := r.db
	if opt.clause != nil {
		db = opt.applyClause(db)
	}
	err := db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository[T]) UpdateMany(ctx context.Context, query any, data any) error {
	var a T
	err := r.db.Model(a).Where(query).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository[T]) DeleteByFieldName(ctx context.Context, query any) error {
	var a T
	db := r.db
	if q, ok := query.(util.Query); ok {
		db = db.Where(q.Query, q.Args...)
	} else {
		db = db.Where(query)
	}

	err := db.Delete(&a).Error
	return err
}

func (r *repository[T]) Count(ctx context.Context, query any) (count int64, err error) {
	var a T

	db := r.db.Model(a)
	if q, ok := query.(util.Query); ok {
		db = db.Where(q.Query, q.Args...)
	} else {
		db = db.Where(query)
	}

	return count, db.Count(&count).Error
}
