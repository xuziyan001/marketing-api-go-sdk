/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// AccountTypeMap : 资金账户类型
type AccountTypeMap string

// List of AccountTypeMap
const (
	AccountTypeMap_CASH                    AccountTypeMap = "FUND_TYPE_CASH"
	AccountTypeMap_GIFT                    AccountTypeMap = "FUND_TYPE_GIFT"
	AccountTypeMap_SHARED                  AccountTypeMap = "FUND_TYPE_SHARED"
	AccountTypeMap_BANK                    AccountTypeMap = "FUND_TYPE_BANK"
	AccountTypeMap_UNION_GIFT              AccountTypeMap = "FUND_TYPE_UNION_GIFT"
	AccountTypeMap_MP_CASH                 AccountTypeMap = "FUND_TYPE_MP_CASH"
	AccountTypeMap_MP_BANK                 AccountTypeMap = "FUND_TYPE_MP_BANK"
	AccountTypeMap_MP_GIFT                 AccountTypeMap = "FUND_TYPE_MP_GIFT"
	AccountTypeMap_CREDIT_ROLL             AccountTypeMap = "FUND_TYPE_CREDIT_ROLL"
	AccountTypeMap_CREDIT_TEMPORARY        AccountTypeMap = "FUND_TYPE_CREDIT_TEMPORARY"
	AccountTypeMap_CONTRACT_GIFT_VIRTUAL   AccountTypeMap = "FUND_TYPE_CONTRACT_GIFT_VIRTUAL"
	AccountTypeMap_CONTRACT_ASSIGN_VIRTUAL AccountTypeMap = "FUND_TYPE_CONTRACT_ASSIGN_VIRTUAL"
	AccountTypeMap_COMPENSATE_VIRTUAL      AccountTypeMap = "FUND_TYPE_COMPENSATE_VIRTUAL"
	AccountTypeMap_INTERNAL_QUOTA          AccountTypeMap = "FUND_TYPE_INTERNAL_QUOTA"
	AccountTypeMap_TEST_VIRTUAL            AccountTypeMap = "FUND_TYPE_TEST_VIRTUAL"
	AccountTypeMap_UNSUPPORTED             AccountTypeMap = "FUND_TYPE_UNSUPPORTED"
)
