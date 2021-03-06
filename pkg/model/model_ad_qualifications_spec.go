/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告资质信息，当且仅当qualification_type=AD_QUALIFICATION时可填且必填
type AdQualificationsSpec struct {
	QualificationCode string    `json:"qualification_code,omitempty"`
	ImageIdList       *[]string `json:"image_id_list,omitempty"`
}
