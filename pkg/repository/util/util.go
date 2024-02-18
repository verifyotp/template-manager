package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
)

type Query struct {
	Query string
	Args  []any
}

func (q Query) String() string {
	// replace ? with %v
	query := strings.ReplaceAll(q.Query, "?", "%v")
	// fill the query with args
	return fmt.Sprintf(query, q.Args...)
}

func (q *Query) Or(query Query) Query {
	result := OrQuery(*q, query)
	q.Query = result.Query
	q.Args = result.Args
	return *q
}

func (q *Query) And(query Query) Query {
	result := AndQuery(*q, query)
	q.Query = result.Query
	q.Args = result.Args
	return *q
}

func (q *Query) Eq(field string, value any) Query {
	eq := Eq(field, value)
	result := AndQuery(*q, eq)
	q.Query = result.Query
	q.Args = result.Args
	return *q
}

func (q *Query) Gte(field string, value any) Query {
	eq := BuildQueryWithOperator(field, ">=", value)
	result := AndQuery(*q, eq)
	q.Query = result.Query
	q.Args = result.Args
	return *q
}

func (q *Query) Lte(field string, value any) Query {
	eq := BuildQueryWithOperator(field, "<=", value)
	result := AndQuery(*q, eq)
	q.Query = result.Query
	q.Args = result.Args
	return *q
}

// BuildQuery is a helper function to build query string
// for gorm where clause
// example:
//
//	BuildQuery("id", 1) => "id = 1"
//	BuildQuery("name", "John") => {"name = ?", "John" }
//	BuildQuery("is_active", true) => {"is_active = ?", true }
//	BuildQuery("is_active", false) => {"is_active = ?", false }
//
// BuildQuery("id", nil) => { "id IS NULL", []any{} }
//
//	// for more complex query
//	BuildQuery("id", []int{1,2,3}) => { "id IN (?)", [1,2,3] }
//	BuildQuery("id", []string{"1","2","3"}) => { "id IN (?)", ['1','2','3'] }
//
// you can use this with FindByFieldName function
// example:
//
//		FindByFieldName(ctx, BuildQuery("id", 1))
//		FindByFieldName(ctx, BuildQuery("id", []int{1,2,3}))
//
//
//	 You can combine multiple query
//	 example:
//
//		AndQuery(BuildQuery("id", 1), BuildQuery("name", "John")) => { "id = ? AND name = ?", [1, "John"] }
//
//		OrQuery(BuildQuery("id", 1), BuildQuery("name", "John")) => { "id = ? OR name = ?", [1, "John"] }
func BuildQuery(field string, value any) Query {

	if field == "" {
		return Query{}
	}

	switch v := value.(type) {
	case string:
		return Query{
			Query: fmt.Sprintf("%s = ?", field),
			Args:  []any{v},
		}
	case *string:
		return Query{
			Query: fmt.Sprintf("%s = ?", field),
			Args:  []any{v},
		}
	case int, int32, int64, float32, float64, uint, uint64, bool:
		return Query{
			Query: fmt.Sprintf("%s = ?", field),
			Args:  []any{v},
		}
	case time.Time:
		return Query{
			Query: fmt.Sprintf("%s = ?", field),
			Args:  []any{v},
		}
	case pq.StringArray:
		var vars []any
		for i := range v {
			vars = append(vars, v[i])
		}
		return Query{
			Query: fmt.Sprintf("%s IN (?)", field),
			Args:  []any{vars},
		}
	case []string:
		var vars []any
		for i := range v {
			vars = append(vars, v[i])
		}
		return Query{
			Query: fmt.Sprintf("%s IN (?)", field),
			Args:  []any{vars},
		}
	case []int, []int32, []int64, []float32, []float64, []uint, []uint64, []bool:
		var values []string
		// convert all values to string
		value := reflect.ValueOf(v)
		for i := 0; i < value.Len(); i++ {
			values = append(values, fmt.Sprintf("%v", value.Index(i)))
		}

		var vars []any
		for i := range values {
			vars = append(vars, values[i])
		}
		return Query{
			Query: fmt.Sprintf("%s IN (?)", field),
			Args:  []any{vars},
		}

	case []interface{}:
		return Query{
			Query: fmt.Sprintf("%s IN (?)", field),
			Args:  v,
		}
	case nil:
		return Query{
			Query: fmt.Sprintf("%s IS NULL", field),
			Args:  []any{},
		}
	case decimal.Decimal:
		return Query{
			Query: fmt.Sprintf("%s = ?", field),
			Args:  []any{v},
		}
	case *decimal.Decimal:
		return Query{
			Query: fmt.Sprintf("%s = ?", field),
			Args:  []any{v},
		}
	default:
		{
			reflectValue := reflect.ValueOf(value)
			switch reflectValue.Kind() {
			case reflect.String:
				return Query{
					Query: fmt.Sprintf("%s = ?", field),
					Args:  []any{reflectValue.String()},
				}
			case reflect.Array, reflect.Map, reflect.Slice:
			case reflect.Bool:
				return Query{
					Query: fmt.Sprintf("%s = ?", field),
					Args:  []any{reflectValue.Bool()},
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return Query{
					Query: fmt.Sprintf("%s = ?", field),
					Args:  []any{reflectValue.Int()},
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return Query{
					Query: fmt.Sprintf("%s = ?", field),
					Args:  []any{reflectValue.Uint()},
				}
			case reflect.Float32, reflect.Float64:
				return Query{
					Query: fmt.Sprintf("%s = ?", field),
					Args:  []any{reflectValue.Float()},
				}
			case reflect.Interface, reflect.Ptr:
				if reflectValue.IsNil() {
					return Query{
						Query: fmt.Sprintf("%s IS NULL", field),
						Args:  []any{},
					}
				}
				return BuildQuery(field, reflectValue.Elem().Interface())
			case reflect.Struct:
				v, ok := value.(time.Time)
				if ok {
					return Query{
						Query: fmt.Sprintf("%s = ?", field),
						Args:  []any{v},
					}
				}
			}
		}

		var vars = []any{v}
		return Query{
			Query: fmt.Sprintf("%s IN (?)", field),
			Args:  vars,
		}
	}
}

// Eq is alias for BuildQuery
var Eq = BuildQuery

// EqArrayField is a helper function to build query string
//
// for array fields on gorm where clause
//
// example:
//
// EqArrayField("id", 1) => "id = ANY(1)"
// EqArrayField("name", "John") => {"name = ANY(?)", "John" }
func EqArrayField(field string, value any) Query {
	return Query{
		Query: fmt.Sprintf("? = ANY(%s)", field),
		Args:  []any{value},
	}
}

func Between(field string, from, to time.Time) Query {
	return Query{
		Query: fmt.Sprintf("%s BETWEEN ? AND ?", field),
		Args:  []any{from, to},
	}
}

// BuildQueryWithOperator is a helper function to build query string
// for gorm where clause
// example:
//
//		BuildQueryWithOperator("id", "=", 1) => { "id = ?", 1 }
//		BuildQueryWithOperator("name", "=", "John") => { "name = ?", "John" }
//		BuildQueryWithOperator("is_active", "=", true) => { "is_active = ?", true }
//		BuildQueryWithOperator("is_active", "=", false) => { "is_active = ?", false }
//	 BuildQueryWithOperator("id", "IN", []int{1,2,3}) => { "id IN (?)", [1,2,3]  }
//	 BuildQueryWithOperator("id", "IN", []string{"1","2","3"}) => { "id IN (?)", ['1','2','3'] }
//	 BuildQueryWithOperator("id", "IN", []interface{}{1,2,3}) => { "id IN (?)", [1,2,3] }
//	 BuildQueryWithOperator("id", "=>", 1) => { "id => ?", 1 }
//	 BuildQueryWithOperator("id", "<=", 1) => { "id <= ?", 1 }
//	 BuildQueryWithOperator("id", "!=", 1) => { "id != ?", 1 }
//	 BuildQueryWithOperator("id", "LIKE", "John") => { "id LIKE ?", "John" }
func BuildQueryWithOperator(field string, operator string, value any) Query {

	switch v := value.(type) {
	case nil:
		return Query{
			Query: fmt.Sprintf("%s %s NULL", field, operator),
			Args:  []any{},
		}
	case string:
		return Query{
			Query: fmt.Sprintf("%s %s ?", field, operator),
			Args:  []any{v},
		}
	case int, int32, int64, float32, float64, uint, uint64, bool:
		return Query{
			Query: fmt.Sprintf("%s %s ?", field, operator),
			Args:  []any{v},
		}
	case time.Time:
		return Query{
			Query: fmt.Sprintf("%s %s ?", field, operator),
			Args:  []any{v},
		}
	case []string, pq.StringArray:
		var values []string

		// convert all values to string
		value := reflect.ValueOf(v)
		for i := 0; i < value.Len(); i++ {
			values = append(values, fmt.Sprintf("%v", value.Index(i)))
		}

		var vars []any
		for i := range values {
			vars = append(vars, values[i])
		}
		return Query{
			Query: fmt.Sprintf("%s %s (?)", field, operator),
			Args:  vars,
		}

	case []int, []int32, []int64, []float32, []float64, []uint, []uint64, []bool:

		var values []string

		// convert all values to string
		value := reflect.ValueOf(v)

		for i := 0; i < value.Len(); i++ {
			values = append(values, fmt.Sprintf("%v", value.Index(i)))
		}

		var vars []any
		for i := range values {
			vars = append(vars, values[i])
		}
		return Query{
			Query: fmt.Sprintf("%s %s (?)", field, operator),
			Args:  vars,
		}

	case []interface{}:
		return Query{
			Query: fmt.Sprintf("%s %s (?)", field, operator),
			Args:  v,
		}
	default:
		return Query{
			Query: fmt.Sprintf("%s %s ?", field, operator),
			Args:  []any{v},
		}
	}

}

var And = AndQuery

// Combine multiple query string with AND operator
// example:
//
//	AndQuery(BuildQuery("id", 1), BuildQuery("name", "John")) => { "id = ? AND name = ?", [1, "John"] }
func AndQuery(queries ...Query) Query {
	if len(queries) == 1 {
		return queries[0]
	}
	var prunned []string
	var prunnedVars []any
	for index := range queries {
		// remove empty strings
		if queries[index].Query != "" {
			prunned = append(prunned, queries[index].Query)
			prunnedVars = append(prunnedVars, queries[index].Args...)
		}
	}
	if len(prunned) == 1 {
		return Query{
			Query: prunned[0],
			Args:  prunnedVars,
		}
	}
	newSQL := strings.Join(prunned, " AND ")
	newSQL = " ( " + newSQL + " ) "

	return Query{
		Query: newSQL,
		Args:  prunnedVars,
	}
}

var Or = OrQuery

// Combine multiple query string with OR operator
// example:
//
//	OrQuery(BuildQuery("id", 1), BuildQuery("name", "John")) => { "id = ? AND name = ?", [1, "John"] }
func OrQuery(queries ...Query) Query {
	if len(queries) == 1 {
		return queries[0]
	}
	var prunned []string
	var prunnedVars []any
	for index := range queries {
		// remove empty strings
		if queries[index].Query != "" {
			prunned = append(prunned, queries[index].Query)
			prunnedVars = append(prunnedVars, queries[index].Args...)
		}
	}
	if len(prunned) == 1 {
		return Query{
			Query: prunned[0],
			Args:  prunnedVars,
		}
	}
	newSQL := strings.Join(prunned, " OR ")
	newSQL = " ( " + newSQL + " ) "

	return Query{
		Query: newSQL,
		Args:  prunnedVars,
	}
}
