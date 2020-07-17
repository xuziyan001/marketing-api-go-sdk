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

type AndroidChannelPackagesGetExample struct {
	TAds                          *ads.SDKClient
	AccessToken                   string
	AccountId                     int64
	MyappAuthKey                  string
	AndroidAppId                  int64
	AndroidChannelPackagesGetOpts *api.AndroidChannelPackagesGetOpts
}

func (e *AndroidChannelPackagesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.MyappAuthKey = "YOUR MYAPP AUTH KEY"
	e.AndroidAppId = int64(0)
	e.AndroidChannelPackagesGetOpts = &api.AndroidChannelPackagesGetOpts{

		Fields: optional.NewInterface([]string{"android_app_id", "package_name", "channel_package_id", "version_code", "version_name", "created_time", "last_modified_time", "system_status", "audit_status"}),
	}
}

func (e *AndroidChannelPackagesGetExample) RunExample() (model.AndroidChannelPackagesGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AndroidChannelPackages().Get(ctx, e.AccountId, e.MyappAuthKey, e.AndroidAppId, e.AndroidChannelPackagesGetOpts)
}

func main() {
	e := &AndroidChannelPackagesGetExample{}
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
