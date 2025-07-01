package query
package pgext

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type condition struct {
	Field string
	Op    string
	Val   interface{}
}

type QueryBuilder[T any] struct {
	table      string
	db         *sqlx.DB
	conditions []condition
	orderBy    string
	limit      int
	offset     int
}

// NewQueryBuilder ...
func NewQueryBuilder[T any](db *sqlx.DB, table string) *QueryBuilder[T] {
	return &QueryBuilder[T]{
		db:    db,
		table: table,
	}
}

// Where adds a condition
func (qb *QueryBuilder[T]) Where(field, op string, val interface{}) *QueryBuilder[T] {
	qb.conditions = append(qb.conditions, condition{field, op, val})
	return qb
}

// Limit sets a limit
func (qb *QueryBuilder[T]) Limit(n int) *QueryBuilder[T] {
	qb.limit = n
	return qb
}

// Offset sets an offset
func (qb *QueryBuilder[T]) Offset(n int) *QueryBuilder[T] {
	qb.offset = n
	return qb
}

// OrderBy sets order
func (qb *QueryBuilder[T]) OrderBy(clause string) *QueryBuilder[T] {
	qb.orderBy = clause
	return qb
}

// Exec executes the query and returns []T
func (qb *QueryBuilder[T]) Exec(ctx context.Context) ([]T, error) {
	var args []interface{}
	var whereClauses []string

	for i, cond := range qb.conditions {
		whereClauses = append(whereClauses, fmt.Sprintf("%s %s $%d", cond.Field, cond.Op, i+1))
		args = append(args, cond.Val)
	}

	query := fmt.Sprintf("SELECT * FROM %s", qb.table)
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	if qb.orderBy != "" {
		query += " ORDER BY " + qb.orderBy
	}
	if qb.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", qb.limit)
	}
	if qb.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", qb.offset)
	}

	var results []T
	err := qb.db.SelectContext(ctx, &results, query, args...)
	return results, err
}
