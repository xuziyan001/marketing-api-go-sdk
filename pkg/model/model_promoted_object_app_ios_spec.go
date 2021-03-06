/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// IOS应用信息
type PromotedObjectAppIosSpec struct {
	Packname           string    `json:"packname,omitempty"`
	Version            string    `json:"version,omitempty"`
	Icon               string    `json:"icon,omitempty"`
	Icon512            string    `json:"icon_512,omitempty"`
	AverageRating      string    `json:"average_rating,omitempty"`
	PackageSize        string    `json:"package_size,omitempty"`
	Genres             *[]string `json:"genres,omitempty"`
	PackageDownloadUrl string    `json:"package_download_url,omitempty"`
}
