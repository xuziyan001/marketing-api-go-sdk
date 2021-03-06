/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 拉取创意结构
type DpPageInfo struct {
	Page      int64 `json:"page,omitempty"`
	PageSize  int64 `json:"page_size,omitempty"`
	TotalPage int64 `json:"total_page,omitempty"`
	TotalNum  int64 `json:"total_num,omitempty"`
}
