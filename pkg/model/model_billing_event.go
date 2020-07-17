/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// BillingEvent : 计费类型
type BillingEvent string

// List of BillingEvent
const (
	BillingEvent_NONE         BillingEvent = "BILLINGEVENT_NONE"
	BillingEvent_CLICK        BillingEvent = "BILLINGEVENT_CLICK"
	BillingEvent_APP_DOWNLOAD BillingEvent = "BILLINGEVENT_APP_DOWNLOAD"
	BillingEvent_IMPRESSION   BillingEvent = "BILLINGEVENT_IMPRESSION"
	BillingEvent_APP_INSTALL  BillingEvent = "BILLINGEVENT_APP_INSTALL"
	BillingEvent_DURATION     BillingEvent = "BILLINGEVENT_DURATION"
	BillingEvent_DAY          BillingEvent = "BILLINGEVENT_DAY"
)
