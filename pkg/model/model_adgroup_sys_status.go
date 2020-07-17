/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// AdgroupSysStatus : 广告组在系统中的状态
type AdgroupSysStatus string

// List of AdgroupSysStatus
const (
	AdgroupSysStatus_NORMAL            AdgroupSysStatus = "AD_GROUP_STATUS_NORMAL"
	AdgroupSysStatus_PENDING           AdgroupSysStatus = "AD_GROUP_STATUS_PENDING"
	AdgroupSysStatus_DENIED            AdgroupSysStatus = "AD_GROUP_STATUS_DENIED"
	AdgroupSysStatus_FROZEN            AdgroupSysStatus = "AD_GROUP_STATUS_FROZEN"
	AdgroupSysStatus_PARTIALLY_PENDING AdgroupSysStatus = "AD_GROUP_STATUS_PARTIALLY_PENDING"
	AdgroupSysStatus_PARTIALLY_NORMAL  AdgroupSysStatus = "AD_GROUP_STATUS_PARTIALLY_NORMAL"
	AdgroupSysStatus_PREPARE           AdgroupSysStatus = "AD_GROUP_STATUS_PREPARE"
	AdgroupSysStatus_DELETED           AdgroupSysStatus = "AD_GROUP_STATUS_DELETED"
	AdgroupSysStatus_INVALID           AdgroupSysStatus = "AD_GROUP_STATUS_INVALID"
)
