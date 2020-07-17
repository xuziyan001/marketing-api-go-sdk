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

type AdvertiserGetExample struct {
	TAds              *ads.SDKClient
	AccessToken       string
	AdvertiserGetOpts *api.AdvertiserGetOpts
}

func (e *AdvertiserGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AdvertiserGetOpts = &api.AdvertiserGetOpts{
		AccountId: optional.NewInt64(int64(0)),

		Fields: optional.NewInterface([]string{"account_id", "daily_budget", "system_status", "reject_message", "corporation_name", "corporation_licence", "certification_image_id", "certification_image", "identity_number", "individual_qualification", "corporate_image_name", "corporate_image_logo", "system_industry_id", "customized_industry", "introduction_url", "industry_qualification_image_id_list", "industry_qualification_image", "ad_qualification_image_id_list", "ad_qualification_image", "contact_person", "contact_person_email", "contact_person_telephone", "contact_person_mobile", "wechat_spec", "websites"}),
	}
}

func (e *AdvertiserGetExample) RunExample() (model.AdvertiserGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Advertiser().Get(ctx, e.AdvertiserGetOpts)
}

func main() {
	e := &AdvertiserGetExample{}
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
