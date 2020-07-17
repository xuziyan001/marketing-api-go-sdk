/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type OptimizationGoalPermissionsGetResponseData struct {
	OptimizationGoalPermissionList             []string                                       `json:"optimization_goal_permission_list,omitempty"`
	DeepBehaviorOptimizationGoalPermissionList []DeepBehaviorOptimizationGoalPermissionStruct `json:"deep_behavior_optimization_goal_permission_list,omitempty"`
	DeepWorthOptimizationGoalPermissionList    []DeepWorthOptimizationGoalPermissionStruct    `json:"deep_worth_optimization_goal_permission_list,omitempty"`
}
