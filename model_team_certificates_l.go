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
// TeamCertificatesL struct for TeamCertificatesL
type TeamCertificatesL struct {
	List []CertificatesR `json:"list"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Total int32 `json:"total"`
}
