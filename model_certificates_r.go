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
import (
	"time"
)
// CertificatesR struct for CertificatesR
type CertificatesR struct {
	// 是否过期
	HasExpired bool `json:"has_expired"`
	// 域名列表
	IssuedTo []string `json:"issued_to"`
	// 证书名称
	Alias string `json:"alias"`
	// 证书类型
	CertificateType string `json:"certificate_type"`
	// 过期时间
	EndData time.Time `json:"end_data"`
	// id
	Id int32 `json:"id"`
	// 证书来源
	IssuedBy string `json:"issued_by"`
}
