/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ProductSaleStatus : 商品在线状态
type ProductSaleStatus string

// List of ProductSaleStatus
const (
	ProductSaleStatus_UNKNOWN ProductSaleStatus = "PRODUCT_SALE_STATUS_UNKNOWN"
	ProductSaleStatus_ONLINE  ProductSaleStatus = "PRODUCT_SALE_STATUS_ONLINE"
	ProductSaleStatus_OFFLINE ProductSaleStatus = "PRODUCT_SALE_STATUS_OFFLINE"
)
