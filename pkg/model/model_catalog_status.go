/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// CatalogStatus : 商品目录状态，详见 <a href='catalog_status' target='_blank'>[商品目录状态]</a>
type CatalogStatus string

// List of CatalogStatus
const (
	CatalogStatus_PREPARE CatalogStatus = "PRODUCT_CATALOG_STATUS_PREPARE"
	CatalogStatus_NORMAL  CatalogStatus = "PRODUCT_CATALOG_STATUS_NORMAL"
	CatalogStatus_SUSPEND CatalogStatus = "PRODUCT_CATALOG_STATUS_SUSPEND"
	CatalogStatus_DELETED CatalogStatus = "PRODUCT_CATALOG_STATUS_DELETED"
)
