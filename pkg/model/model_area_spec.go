/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 区域定义
type AreaSpec struct {
	AreaType   LbsAreaType `json:"area_type,omitempty"`
	CircleArea CircleArea  `json:"circle_area,omitempty"`
}
