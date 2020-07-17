/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PromotedObjectsAddRequest struct {
	PromotedObjectType PromotedObjectType `json:"promoted_object_type,omitempty"`
	PromotedObjectId   string             `json:"promoted_object_id,omitempty"`
	PromotedObjectSpec PromotedObjectSpec `json:"promoted_object_spec,omitempty"`
	AccountId          int64              `json:"account_id,omitempty"`
}
