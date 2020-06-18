# EnterpriseHttpGatewayRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **int32** |  | [optional] [readonly] 
**RegionName** | **string** | 所属集群ID | 
**TeamName** | **string** | 所属团队唯一名称 | 
**AppId** | **int32** | 所属应用ID | 
**AutoSslConfig** | **string** | 自动签发方式 | 
**HttpRuleId** | **string** | http_rule_id | 
**RegionId** | **string** | region id | 
**TenantId** | **string** | 租户id | 
**ServiceId** | **string** | 组件id | 
**ServiceName** | **string** | 组件名 | 
**DomainName** | **string** | 域名 | 
**ContainerPort** | **int32** | 容器端口 | [optional] 
**Protocol** | **string** | 域名类型 http https httptphttps httpandhttps | [optional] 
**CertificateId** | **int32** | 证书ID | [optional] 
**DomainType** | **string** | 组件域名类型 | [optional] 
**ServiceAlias** | **string** | 组件别名 | [optional] 
**IsSenior** | **bool** | 是否有高级路由 | [optional] 
**DomainPath** | **string** | 域名path | [optional] 
**DomainCookie** | **string** | 域名cookie | [optional] 
**DomainHeander** | **string** | 域名heander | [optional] 
**Type** | **int32** | 类型（默认：0， 自定义：1） | [optional] 
**TheWeight** | **int32** | 权重 | [optional] 
**RuleExtensions** | **string** | 扩展功能 | [optional] 
**IsOuterService** | **bool** | 是否已开启对外端口 | [optional] 
**AutoSsl** | **bool** | 是否自动匹配证书，升级为https，如果开启，由外部服务完成升级 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


