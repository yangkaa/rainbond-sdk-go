# \OpenapiGatewayApi

All URIs are relative to *http://0.0.0.0:7070/openapi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HttpdomainsList**](OpenapiGatewayApi.md#HttpdomainsList) | **Get** /httpdomains | 
[**TeamsRegionsAppsDomainsDelete**](OpenapiGatewayApi.md#TeamsRegionsAppsDomainsDelete) | **Delete** /teams/{team_id}/regions/{region_name}/apps/{app_id}/domains/{rule_id} | 
[**TeamsRegionsAppsDomainsList**](OpenapiGatewayApi.md#TeamsRegionsAppsDomainsList) | **Get** /teams/{team_id}/regions/{region_name}/apps/{app_id}/domains | 
[**TeamsRegionsAppsDomainsUpdate**](OpenapiGatewayApi.md#TeamsRegionsAppsDomainsUpdate) | **Put** /teams/{team_id}/regions/{region_name}/apps/{app_id}/domains/{rule_id} | 



## HttpdomainsList

> []EnterpriseHttpGatewayRule HttpdomainsList(ctx, optional)



获取企业应用http访问策略列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HttpdomainsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HttpdomainsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoSsl** | **optional.Bool**| 查询条件，是否为需要自动匹配证书的策略 | 

### Return type

[**[]EnterpriseHttpGatewayRule**](EnterpriseHTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsDomainsDelete

> HttpGatewayRule TeamsRegionsAppsDomainsDelete(ctx, appId, regionName, ruleId, teamId)



删除HTTP访问策略

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**regionName** | **string**|  | 
**ruleId** | **string**|  | 
**teamId** | **string**|  | 

### Return type

[**HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsDomainsList

> GatewayRule TeamsRegionsAppsDomainsList(ctx, regionName, teamId, appId)



获取应用访问策略列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 

### Return type

[**GatewayRule**](GatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsDomainsUpdate

> HttpGatewayRule TeamsRegionsAppsDomainsUpdate(ctx, regionName, ruleId, teamId, appId, data)



更新访问策略

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**ruleId** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
**data** | [**UpdatePostHttpGatewayRule**](UpdatePostHttpGatewayRule.md)|  | 

### Return type

[**HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

