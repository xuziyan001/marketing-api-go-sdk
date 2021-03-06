/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 诊断详细信息
type ResponseDetailStruct struct {
	Funnel           *DetailFunnelStruct         `json:"funnel,omitempty"`
	TargetDetail     *DetailTargetDetailStruct   `json:"target_detail,omitempty"`
	CompeteDetail    *DetailCompeteDetailStruct  `json:"compete_detail,omitempty"`
	ExposureDetail   *ExposureDetailStruct       `json:"exposure_detail,omitempty"`
	ClickDetail      *ClickDetailStruct          `json:"click_detail,omitempty"`
	ConversionDetail *ConversionDetailStruct     `json:"conversion_detail,omitempty"`
	CpaDetail        *DetailCpaDetailStruct      `json:"cpa_detail,omitempty"`
	CreativeDetail   *DetailCreativeDetailStruct `json:"creative_detail,omitempty"`
	Optimization     *OptimizeContentMainStruct  `json:"optimization,omitempty"`
}
