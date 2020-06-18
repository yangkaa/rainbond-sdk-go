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
// UserInfo struct for UserInfo
type UserInfo struct {
	UserId int32 `json:"user_id"`
	// 邮件地址
	Email *string `json:"email,omitempty"`
	// 用户昵称
	NickName string `json:"nick_name,omitempty"`
	// 手机号码
	Phone *string `json:"phone,omitempty"`
	// 激活状态
	IsActive bool `json:"is_active,omitempty"`
	// 用户来源
	Origion *string `json:"origion,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 注册ip
	ClientIp *string `json:"client_ip,omitempty"`
	// enterprise_id
	EnterpriseId string `json:"enterprise_id,omitempty"`
}
