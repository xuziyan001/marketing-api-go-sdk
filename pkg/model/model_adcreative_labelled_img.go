/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 标签图片
type AdcreativeLabelledImg struct {
	Image string            `json:"image,omitempty"`
	Label []AdcreativeLabel `json:"label,omitempty"`
}
