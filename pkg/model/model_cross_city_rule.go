/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 跨城市规则
type CrossCityRule struct {
	Route     []string  `json:"route,omitempty"`
	DateRange DateRange `json:"date_range,omitempty"`
	Frequency int64     `json:"frequency,omitempty"`
}
