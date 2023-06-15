package database

import "gorm.io/gorm"

type OrderBy string

const (
	OrderByAsc  OrderBy = "asc"
	OrderByDesc OrderBy = "desc"
)

type Where struct {
	column string
	op     string
	value  interface{}
}

type Filter struct {
	where   []Where
	offset  int
	limit   int
	orderBy map[string]OrderBy
}

func NewFilter() *Filter {
	return &Filter{
		where:   make([]Where, 0),
		orderBy: make(map[string]OrderBy),
	}
}

func (f *Filter) AddWhere(column, op string, value interface{}) *Filter {
	f.where = append(f.where, Where{column: column, op: op, value: value})
	return f
}

func (f *Filter) AddOrderBy(column string, orderBy OrderBy) *Filter {
	f.orderBy[column] = orderBy

	return f
}

func (f *Filter) SetOffset(offset int) *Filter {
	f.offset = offset
	return f
}

func (f *Filter) SetLimit(limit int) *Filter {
	f.limit = limit
	return f
}

func (f *Filter) Build(db *gorm.DB) *gorm.DB {

	if len(f.where) > 0 {
		for _, v := range f.where {
			db = db.Where(v.column+" "+v.op+" ?", v.value)
		}
	}

	if f.limit > 0 && f.offset > 0 {
		db = db.Limit(f.limit).Offset((f.offset - 1) * f.limit)
	} else if f.limit > 0 {
		db = db.Limit(f.limit)
	} else if f.offset > 0 {
		db = db.Offset((f.offset - 1) * f.limit)
	}

	if len(f.orderBy) > 0 {
		for k, v := range f.orderBy {
			db = db.Order(k + " " + string(v))
		}
	} else {
		db = db.Order("id desc")
	}

	return db
}

func (f *Filter) Count(db *gorm.DB) int64 {
	var count int64
	f.
		SetOffset(0).
		SetLimit(0).
		Build(db).
		Count(&count)
	return count
}
