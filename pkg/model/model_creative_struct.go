/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意规格结构体
type CreativeStruct struct {
	DeepLinkUrl          string                                `json:"deep_link_url,omitempty"`
	AdcreativeTemplateId int64                                 `json:"adcreative_template_id,omitempty"`
	AdcreativeElements   AdcreativeCreativeElementsWithOptions `json:"adcreative_elements,omitempty"`
}
