// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	books "github.com/rafaelhl/library-gorm-lib/books"

	mock "github.com/stretchr/testify/mock"
)

// BookInserter is an autogenerated mock type for the BookInserter type
type BookInserter struct {
	mock.Mock
}

// InsertBook provides a mock function with given fields: ctx, book
func (_m *BookInserter) InsertBook(ctx context.Context, book books.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, books.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}