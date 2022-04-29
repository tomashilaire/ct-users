// Package classification Entity API
//
// Documentation for Entity API
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
package entityhttphdl

import "root/internal/core/domain"

// swagger:parameters findEntity deleteEntity updateEntity
type IdParameter struct {
	// in: path
	// required: true
	Id string `bson:"_id" json:"id"`
}

// swagger:response entityResponse
type EntityResponse struct {
	// in: body
	Body domain.Entity
}

// swagger:response entitiesResponse
type EntitiesResponse struct {
	// in: body
	Body []domain.Entity
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

// swagger:parameters createEntity
type BodyCreate struct {
	// in: body
	// required: true
	Body BodyWrapper
}

// swagger:parameters updateEntity
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
