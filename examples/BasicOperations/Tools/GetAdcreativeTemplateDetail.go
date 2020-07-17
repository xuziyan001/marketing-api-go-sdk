/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type AdcreativeTemplateDetailGetExample struct {
	TAds                            *ads.SDKClient
	AccessToken                     string
	AdcreativeTemplateId            int64
	PromotedObjectType              string
	AdcreativeTemplateDetailGetOpts *api.AdcreativeTemplateDetailGetOpts
}

func (e *AdcreativeTemplateDetailGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AdcreativeTemplateId = int64(968)
	e.PromotedObjectType = "PROMOTED_OBJECT_TYPE_LINK"
	e.AdcreativeTemplateDetailGetOpts = &api.AdcreativeTemplateDetailGetOpts{

		AccountId: optional.NewInt64(int64(0)),

		SiteSet: optional.NewInterface([]string{"SITE_SET_QZONE"}),

		Fields: optional.NewInterface([]string{"adcreative_template_id", "adcreative_template_name", "adcreative_template_description", "adcreative_template_size", "adcreative_template_style", "adcreative_sample_image_list", "ad_attributes", "adcreative_attributes", "adcreative_elements", "support_billing_spec_list", "support_page_type", "unsupport_billing_spec_list", "unsupport_ad_attributes_spec_list", "unsupport_adcreative_attributes_spec_list", "unsupport_siteset_detail_spec", "support_dynamic_ability_spec_list"}),
	}
}

func (e *AdcreativeTemplateDetailGetExample) RunExample() (model.AdcreativeTemplateDetailGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AdcreativeTemplateDetail().Get(ctx, e.AdcreativeTemplateId, e.PromotedObjectType, e.AdcreativeTemplateDetailGetOpts)
}

func main() {
	e := &AdcreativeTemplateDetailGetExample{}
	e.Init()
	response, httpResponse, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Http response:", httpResponse)
}
