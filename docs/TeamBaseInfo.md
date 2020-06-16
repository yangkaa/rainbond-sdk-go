# TeamBaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **int32** |  | [optional] [readonly] 
**TenantId** | **string** | 租户id | [optional] 
**TenantName** | **string** | 租户名称 | 
**Region** | **string** | 区域中心,弃用 | [optional] 
**IsActive** | **bool** | 激活状态 | [optional] 
**CreateTime** | [**time.Time**](time.Time.md) | 创建时间 | [optional] [readonly] 
**Creater** | **int32** | 租户创建者 | [optional] 
**LimitMemory** | **int32** | 内存大小单位（M） | [optional] 
**UpdateTime** | [**time.Time**](time.Time.md) | 更新时间 | [optional] [readonly] 
**ExpiredTime** | Pointer to [**time.Time**](time.Time.md) | 过期时间 | [optional] 
**TenantAlias** | Pointer to **string** | 团队别名 | [optional] 
**EnterpriseId** | Pointer to **string** | 企业id | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


