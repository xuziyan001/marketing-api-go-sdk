/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// custom_audience洞察分析结构
type CustomAudienceInsights struct {
	DimensionType DimensionType       `json:"dimension_type,omitempty"`
	MatchRate     float64             `json:"match_rate,omitempty"`
	Distribution  []RangeDistribution `json:"distribution,omitempty"`
}
