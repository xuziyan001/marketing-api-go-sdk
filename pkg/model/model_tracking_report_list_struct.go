/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 点击追踪报表结构体
type TrackingReportListStruct struct {
	Date            string     `json:"date,omitempty"`
	Hour            int64      `json:"hour,omitempty"`
	RawClicks       int64      `json:"raw_clicks,omitempty"`
	Var1minClicks   int64      `json:"1min_clicks,omitempty"`
	Var2minClicks   int64      `json:"2min_clicks,omitempty"`
	Var15minClicks  int64      `json:"15min_clicks,omitempty"`
	RequestFail     int64      `json:"request_fail,omitempty"`
	ResponseFail    int64      `json:"response_fail,omitempty"`
	HttpStatusError int64      `json:"http_status_error,omitempty"`
	OtherError      int64      `json:"other_error,omitempty"`
	SourceType      SourceType `json:"source_type,omitempty"`
}
