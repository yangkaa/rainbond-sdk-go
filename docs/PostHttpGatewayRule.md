# PostHttpGatewayRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSsl** | **bool** | 是否自动匹配证书，升级为https，如果开启，由外部服务完成升级 | [optional] [default to false]
**AutoSslConfig** | **string** | 自动分发证书配置 | [optional] 
**CertificateId** | **int32** | 证书id | [optional] 
**ContainerPort** | **int32** | 绑定端口 | 
**DomainCookie** | **string** | 域名cookie | [optional] 
**DomainHeader** | **string** | 域名header | [optional] 
**DomainName** | **string** | 域名 | 
**DomainPath** | **string** | 域名路径 | [optional] [default to /]
**RuleExtensions** | **[]string** | 规则扩展 | [optional] 
**ServiceId** | **string** | 应用组件id | 
**TheWeight** | **int32** |  | [optional] 
**WhetherOpen** | **bool** | 是否开放 | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


