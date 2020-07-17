/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProductItemsUpdateRequest struct {
	AccountId         int64             `json:"account_id,omitempty"`
	CatalogId         int64             `json:"catalog_id,omitempty"`
	ProductOuterId    string            `json:"product_outer_id,omitempty"`
	Price             float64           `json:"price,omitempty"`
	SalePrice         float64           `json:"sale_price,omitempty"`
	PricePc           float64           `json:"price_pc,omitempty"`
	PriceMobile       float64           `json:"price_mobile,omitempty"`
	PriceApp          float64           `json:"price_app,omitempty"`
	StockVolume       int64             `json:"stock_volume,omitempty"`
	Discount          float64           `json:"discount,omitempty"`
	ExpirationTime    int64             `json:"expiration_time,omitempty"`
	ProductSaleStatus ProductSaleStatus `json:"product_sale_status,omitempty"`
	ProductVisibility ProductVisibility `json:"product_visibility,omitempty"`
}
