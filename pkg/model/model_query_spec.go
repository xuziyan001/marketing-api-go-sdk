/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 查询条件
type QuerySpec struct {
	WechatLinkAd WechatLinkAd `json:"wechat_link_ad,omitempty"`
	WechatOcpa   WechatOcpa   `json:"wechat_ocpa,omitempty"`
}
