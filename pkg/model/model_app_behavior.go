/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// app行为定向
type AppBehavior struct {
	ObjectType   AppActionObjectType `json:"object_type,omitempty"`
	ObjectIdList []int64             `json:"object_id_list,omitempty"`
	TimeWindow   int64               `json:"time_window,omitempty"`
	ActIdList    []string            `json:"act_id_list,omitempty"`
}
