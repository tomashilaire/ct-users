// Package classification Test API
//
// Documentation for Test API
//
//	Schemes: http
//	BasePath: /
//  Version: 0.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
package testhttphdl

import "test/internal/core/domain"

// swagger:parameters findTest deleteTest updateTest
type IdParameter struct {
	// in: path
	// required: true
	Id string `bson:"_id" json:"id"`
}

// swagger:response testResponse
type TestResponse struct {
	// in: body
	Body domain.Test
}

// swagger:response testsResponse
type TestsResponse struct {
	// in: body
	Body []domain.Test
}

// swagger:response idResponse
type IdResponse struct {
	// in: body
	Id IdWrapper
}

// swagger:response errorResponse
type GenericError struct {
	// in: body
	Message MessageWrapper
}

// swagger:parameters createTest
type BodyCreate struct {
	// in: body
	// required: true
	Body BodyWrapper
}

// swagger:parameters updateTest
type BodyUpdate struct {
	// in: body
	Body BodyWrapper
}

type IdWrapper struct {
	Id string `bson:"_id" json:"id"`
}

type BodyWrapper struct {
	Name   string `bson:"name" json:"name"`
	Action string `bson:"action" json:"action"`
}

type MessageWrapper struct {
	Message string `json:"message"`
}
