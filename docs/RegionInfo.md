# RegionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | region id | 
**RegionName** | **string** | 数据中心名称,不可修改 | 
**RegionAlias** | **string** | 数据中心别名 | 
**Url** | **string** | 数据中心API url | 
**Token** | Pointer to **string** | 数据中心token | [optional] 
**Wsurl** | **string** | 数据中心Websocket url | 
**Httpdomain** | **string** | 数据中心http应用访问根域名 | 
**Tcpdomain** | **string** | 数据中心tcp应用访问根域名 | 
**Scope** | **string** | 数据中心范围 private|public | [optional] 
**SslCaCert** | Pointer to **string** | 数据中心访问ca证书地址 | [optional] 
**CertFile** | Pointer to **string** | 验证文件 | [optional] 
**KeyFile** | Pointer to **string** | 验证的key | [optional] 
**Status** | **string** | 数据中心状态 0：编辑中 1:启用 2：停用 3:维护中 | 
**Desc** | **string** | 数据中心描述 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


