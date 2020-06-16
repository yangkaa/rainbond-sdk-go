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
// TeamAppsResource struct for TeamAppsResource
type TeamAppsResource struct {
	// cpu总额
	TotalCpu int32 `json:"total_cpu,omitempty"`
	// 内存总额
	TotalMemory int32 `json:"total_memory,omitempty"`
	// 占用cpu
	UsedCpu int32 `json:"used_cpu,omitempty"`
	// 占用内存
	UsedMemory int32 `json:"used_memory,omitempty"`
	// 占用cpu百分比
	UsedCpuPercentage float32 `json:"used_cpu_percentage,omitempty"`
	// 占用内存百分比
	UsedMemoryPercentage float32 `json:"used_memory_percentage,omitempty"`
	// 团队ID
	TeamId string `json:"team_id"`
	// 团队名称
	TeamName string `json:"team_name"`
	// 团队昵称
	TeamAlias string `json:"team_alias"`
}
