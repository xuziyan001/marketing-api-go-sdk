/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 竞争表现维度的曝光评估结果
type CompititionDiagnosisConclusionSpec struct {
	CompititionDiagnosisScore         int64                            `json:"compitition_diagnosis_score,omitempty"`
	CompititionDiagnosisResult        string                           `json:"compitition_diagnosis_result,omitempty"`
	BidAmountDiagnosisScore           int64                            `json:"bid_amount_diagnosis_score,omitempty"`
	BidAmountDiagnosisConclusion      string                           `json:"bid_amount_diagnosis_conclusion,omitempty"`
	ClickDiagnosisScore               int64                            `json:"click_diagnosis_score,omitempty"`
	ClickDiagnosisConclusion          string                           `json:"click_diagnosis_conclusion,omitempty"`
	ConversionDiagnosisScore          int64                            `json:"conversion_diagnosis_score,omitempty"`
	ConversionDiagnosisConclusion     string                           `json:"conversion_diagnosis_conclusion,omitempty"`
	UserAcceptanceDiagnosisScore      int64                            `json:"user_acceptance_diagnosis_score,omitempty"`
	UserAcceptanceDiagnosisConclusion string                           `json:"user_acceptance_diagnosis_conclusion,omitempty"`
	CompititionAnalysisDetailSpecList *[]CompititionAnalysisDetailSpec `json:"compitition_analysis_detail_spec_list,omitempty"`
}
