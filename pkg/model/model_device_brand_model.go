/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 设备品牌型号定向
type DeviceBrandModel struct {
	IncludedList []int64 `json:"included_list,omitempty"`
	ExcludedList []int64 `json:"excluded_list,omitempty"`
}
