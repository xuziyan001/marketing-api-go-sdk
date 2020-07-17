/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ClaimType : 归因方式，SDK上报方式时归因方式只能为CLAIM_TYPE_ACTIVATION
type ClaimType string

// List of ClaimType
const (
	ClaimType_ACTIVATION ClaimType = "CLAIM_TYPE_ACTIVATION"
	ClaimType_CLICK      ClaimType = "CLAIM_TYPE_CLICK"
)
