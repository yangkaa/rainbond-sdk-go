# RegionInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | region id | [optional] 
**RegionName** | **string** | 集群名称 | 
**RegionAlias** | **string** | 集群别名 | 
**Url** | **string** | 集群API url | 
**Wsurl** | **string** | 集群Websocket url | [optional] 
**Httpdomain** | **string** | 集群http应用访问根域名 | [optional] 
**Tcpdomain** | **string** | 集群tcp应用访问根域名 | [optional] 
**Status** | **string** | 集群状态 0：编辑中 1:启用 2：停用 3:维护中 | 
**Desc** | **string** | 集群描述 | 
**Scope** | **string** | 数据中心范围 private|public | [optional] [default to private]
**SslCaCert** | Pointer to **string** | api ca file | [optional] 
**CertFile** | Pointer to **string** | api cert file | [optional] 
**KeyFile** | Pointer to **string** | api cert key file | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


