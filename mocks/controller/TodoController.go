// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// TodoController is an autogenerated mock type for the TodoController type
type TodoController struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *TodoController) Create(c *gin.Context) {
	_m.Called(c)
}

// Delete provides a mock function with given fields: c
func (_m *TodoController) Delete(c *gin.Context) {
	_m.Called(c)
}

// FindAll provides a mock function with given fields: c
func (_m *TodoController) FindAll(c *gin.Context) {
	_m.Called(c)
}

// FindOne provides a mock function with given fields: c
func (_m *TodoController) FindOne(c *gin.Context) {
	_m.Called(c)
}

// Update provides a mock function with given fields: c
func (_m *TodoController) Update(c *gin.Context) {
	_m.Called(c)
}