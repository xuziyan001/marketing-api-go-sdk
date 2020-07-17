/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type FundStatementsDetailedGetListStruct struct {
	Time           int64     `json:"time,omitempty"`
	ExternalBillNo string    `json:"external_bill_no,omitempty"`
	TradeType      TradeType `json:"trade_type,omitempty"`
	Amount         int64     `json:"amount,omitempty"`
	Description    string    `json:"description,omitempty"`
}
