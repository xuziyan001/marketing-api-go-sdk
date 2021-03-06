/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// PackageStatus : 渠道包在系统中的状态
type PackageStatus string

// List of PackageStatus
const (
	PackageStatus_NORMAL         PackageStatus = "PACKAGE_STATUS_NORMAL"
	PackageStatus_PENDING        PackageStatus = "PACKAGE_STATUS_PENDING"
	PackageStatus_DENIED         PackageStatus = "PACKAGE_STATUS_DENIED"
	PackageStatus_UPDATE_DENIED  PackageStatus = "PACKAGE_STATUS_UPDATE_DENIED"
	PackageStatus_OFFLINE        PackageStatus = "PACKAGE_STATUS_OFFLINE"
	PackageStatus_DISABLE        PackageStatus = "PACKAGE_STATUS_DISABLE"
	PackageStatus_REVIEW_PENDING PackageStatus = "PACKAGE_STATUS_REVIEW_PENDING"
)
