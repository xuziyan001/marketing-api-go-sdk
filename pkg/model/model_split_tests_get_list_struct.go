/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type SplitTestsGetListStruct struct {
	AccountId                int64             `json:"account_id,omitempty"`
	SplitTestId              int64             `json:"split_test_id,omitempty"`
	SplitTestStatus          SplitTestStatus   `json:"split_test_status,omitempty"`
	SplitTestName            string            `json:"split_test_name,omitempty"`
	BeginTime                int64             `json:"begin_time,omitempty"`
	EndTime                  int64             `json:"end_time,omitempty"`
	SmartExpandEnabled       bool              `json:"smart_expand_enabled,omitempty"`
	AdgroupIdList            []int64           `json:"adgroup_id_list,omitempty"`
	RecommendedRating        RecommendedRating `json:"recommended_rating,omitempty"`
	RecommendedAdgroupIdList []int64           `json:"recommended_adgroup_id_list,omitempty"`
}
