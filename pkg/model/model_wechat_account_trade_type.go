/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// WechatAccountTradeType : 交易类型
type WechatAccountTradeType string

// List of WechatAccountTradeType
const (
	WechatAccountTradeType_AGENCY_CHARGE                 WechatAccountTradeType = "AGENCY_CHARGE"
	WechatAccountTradeType_AGENCY_TRANSFER_TO_ADVERTISER WechatAccountTradeType = "AGENCY_TRANSFER_TO_ADVERTISER"
	WechatAccountTradeType_AGENCY_REFUND_FROM_ADVERTISER WechatAccountTradeType = "AGENCY_REFUND_FROM_ADVERTISER"
	WechatAccountTradeType_AGENCY_CREDIT_REPAY           WechatAccountTradeType = "AGENCY_CREDIT_REPAY"
	WechatAccountTradeType_AGENCY_REFUND                 WechatAccountTradeType = "AGENCY_REFUND"
	WechatAccountTradeType_ADVERTISER_CHARGE             WechatAccountTradeType = "ADVERTISER_CHARGE"
	WechatAccountTradeType_ADVERTISER_TRANSFER           WechatAccountTradeType = "ADVERTISER_TRANSFER"
)
