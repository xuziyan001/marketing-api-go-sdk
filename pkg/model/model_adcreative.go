/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告创意
type Adcreative struct {
	AdcreativeId           int64                          `json:"adcreative_id,omitempty"`
	AdcreativeName         string                         `json:"adcreative_name,omitempty"`
	CampaignId             int64                          `json:"campaign_id,omitempty"`
	PageType               PageTypeRead                   `json:"page_type,omitempty"`
	PageSpec               PageSpec                       `json:"page_spec,omitempty"`
	SiteSet                []string                       `json:"site_set,omitempty"`
	PromotedObjectType     PromotedObjectType             `json:"promoted_object_type,omitempty"`
	PromotedObjectId       string                         `json:"promoted_object_id,omitempty"`
	CreatedTime            int64                          `json:"created_time,omitempty"`
	LastModifiedTime       int64                          `json:"last_modified_time,omitempty"`
	PlayablePageMaterialId string                         `json:"playable_page_material_id,omitempty"`
	VideoEndPage           VideoEndPageSpec               `json:"video_end_page,omitempty"`
	AdcreativeTemplateId   int64                          `json:"adcreative_template_id,omitempty"`
	AdcreativeElements     AdcreativeCreativeElementsRead `json:"adcreative_elements,omitempty"`
}
