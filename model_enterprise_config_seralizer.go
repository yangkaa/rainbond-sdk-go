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
// EnterpriseConfigSeralizer struct for EnterpriseConfigSeralizer
type EnterpriseConfigSeralizer struct {
	ExportApp ExportAppBaseResp `json:"export_app,omitempty"`
	AutoSsl AutoSsl `json:"auto_ssl,omitempty"`
	OauthServices OauthServicesBaseResp `json:"oauth_services,omitempty"`
	CloudMarket CloudMarketBaseResp `json:"cloud_market,omitempty"`
	ObjectStorage ObjectStorageBaseResp `json:"object_storage,omitempty"`
	AppstoreImageHub AppStoreImageHubBaseResp `json:"appstore_image_hub,omitempty"`
	NewbieGuide NewBieGuideBaseResp `json:"newbie_guide,omitempty"`
}
