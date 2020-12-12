# HttpGatewayRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **int32** |  | [optional] [readonly] 
**AutoSsl** | **bool** | 是否自动匹配证书，升级为https，如果开启，由外部服务完成升级 | [optional] 
**AutoSslConfig** | Pointer to **string** | 自动分发证书配置 | [optional] 
**CertificateId** | **int32** | 证书ID | [optional] 
**ContainerPort** | **int32** | 容器端口 | [optional] 
**DomainCookie** | **string** | 域名cookie | [optional] 
**DomainHeander** | **string** | 域名heander | [optional] 
**DomainName** | **string** | 域名 | 
**DomainPath** | **string** | 域名path | [optional] 
**DomainType** | **string** | 组件域名类型 | [optional] 
**HttpRuleId** | **string** | http_rule_id | 
**IsOuterService** | **bool** | 是否已开启对外端口 | [optional] 
**IsSenior** | **bool** | 是否有高级路由 | [optional] 
**Protocol** | **string** | 域名类型 http https httptphttps httpandhttps | [optional] 
**RegionId** | **string** | region id | 
**RuleExtensions** | **[]string** |  | [optional] [readonly] 
**ServiceAlias** | **string** | 组件别名 | [optional] 
**ServiceId** | **string** | 组件id | 
**ServiceName** | **string** | 组件名 | 
**TenantId** | **string** | 租户id | 
**TheWeight** | **int32** | 权重 | [optional] 
**Type** | **int32** | 类型（默认：0， 自定义：1） | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


