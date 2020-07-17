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

type CapabilitiesGetExample struct {
	TAds                *ads.SDKClient
	AccessToken         string
	AccountId           int64
	Capability          string
	CapabilitiesGetOpts *api.CapabilitiesGetOpts
}

func (e *CapabilitiesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.Capability = "CAPABILITY_WECHAT_ECOMMERCE_PRODUCT"
	e.CapabilitiesGetOpts = &api.CapabilitiesGetOpts{

		Fields: optional.NewInterface([]string{"wechat_ecommerce_product_spec", "wechat_link_ad_spec", "wechat_ocpa_spec"}),
	}
}

func (e *CapabilitiesGetExample) RunExample() (model.CapabilitiesGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Capabilities().Get(ctx, e.AccountId, e.Capability, e.CapabilitiesGetOpts)
}

func main() {
	e := &CapabilitiesGetExample{}
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
