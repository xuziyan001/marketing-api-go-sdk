/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 曝光量分析
type ExposureCountFunnelStruct struct {
	Value        float64 `json:"value,omitempty"`
	Score        int64   `json:"score,omitempty"`
	ScoreDesc    string  `json:"score_desc,omitempty"`
	RankCategory int64   `json:"rank_category,omitempty"`
	RankOverall  int64   `json:"rank_overall,omitempty"`
}
