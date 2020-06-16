/*
 * Rainbond Open API
 *
 * Rainbond open api
 *
 * API version: v1
 * Contact: barnett@goodrain.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// OauthServicesResp struct for OauthServicesResp
type OauthServicesResp struct {
	Enable bool `json:"enable,omitempty"`
	AuthUrl string `json:"auth_url,omitempty"`
	Name string `json:"name"`
	IsConsole bool `json:"is_console,omitempty"`
	IsAutoLogin bool `json:"is_auto_login,omitempty"`
	ServiceId *int32 `json:"service_id,omitempty"`
	OauthType string `json:"oauth_type"`
	Eid string `json:"eid"`
	HomeUrl *string `json:"home_url,omitempty"`
	AccessTokenUrl string `json:"access_token_url,omitempty"`
	ApiUrl string `json:"api_url,omitempty"`
	IsDeleted bool `json:"is_deleted,omitempty"`
	IsGit bool `json:"is_git,omitempty"`
}
