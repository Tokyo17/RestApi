// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	products "final/business/products"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FilterByType provides a mock function with given fields: ctx, typeid
func (_m *Repository) FilterByType(ctx context.Context, typeid int) ([]products.ProductDomain, error) {
	ret := _m.Called(ctx, typeid)

	var r0 []products.ProductDomain
	if rf, ok := ret.Get(0).(func(context.Context, int) []products.ProductDomain); ok {
		r0 = rf(ctx, typeid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.ProductDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, typeid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx
func (_m *Repository) Get(ctx context.Context) ([]products.ProductDomain, error) {
	ret := _m.Called(ctx)

	var r0 []products.ProductDomain
	if rf, ok := ret.Get(0).(func(context.Context) []products.ProductDomain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.ProductDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: ctx, domain, id
func (_m *Repository) UpdateProduct(ctx context.Context, domain products.ProductDomain, id int) (products.ProductDomain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 products.ProductDomain
	if rf, ok := ret.Get(0).(func(context.Context, products.ProductDomain, int) products.ProductDomain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(products.ProductDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, products.ProductDomain, int) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadProduct provides a mock function with given fields: ctx, productdomain
func (_m *Repository) UploadProduct(ctx context.Context, productdomain products.ProductDomain) (products.ProductDomain, error) {
	ret := _m.Called(ctx, productdomain)

	var r0 products.ProductDomain
	if rf, ok := ret.Get(0).(func(context.Context, products.ProductDomain) products.ProductDomain); ok {
		r0 = rf(ctx, productdomain)
	} else {
		r0 = ret.Get(0).(products.ProductDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, products.ProductDomain) error); ok {
		r1 = rf(ctx, productdomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadType provides a mock function with given fields: ctx, domain
func (_m *Repository) UploadType(ctx context.Context, domain products.Product_typeDomain) (products.Product_typeDomain, error) {
	ret := _m.Called(ctx, domain)

	var r0 products.Product_typeDomain
	if rf, ok := ret.Get(0).(func(context.Context, products.Product_typeDomain) products.Product_typeDomain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(products.Product_typeDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, products.Product_typeDomain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
