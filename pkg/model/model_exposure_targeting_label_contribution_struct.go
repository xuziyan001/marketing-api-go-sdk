/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 定向标签（曝光）贡献度
type ExposureTargetingLabelContributionStruct struct {
	Locard   []LabelContributionListItem `json:"locard,omitempty"`
	Audience []LabelContributionListItem `json:"audience,omitempty"`
}
