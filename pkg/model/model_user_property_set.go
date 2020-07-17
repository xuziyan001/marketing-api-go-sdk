/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// UserPropertySet返回结构
type UserPropertySet struct {
	UserPropertySetId int64  `json:"user_property_set_id,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	CreatedTime       string `json:"created_time,omitempty"`
}
