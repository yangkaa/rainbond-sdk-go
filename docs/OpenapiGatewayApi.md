# \OpenapiGatewayApi

All URIs are relative to *http://127.0.0.1:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HttpdomainsList**](OpenapiGatewayApi.md#HttpdomainsList) | **Get** /openapi/v1/httpdomains | 
[**TeamsRegionsAppsDomainsDelete**](OpenapiGatewayApi.md#TeamsRegionsAppsDomainsDelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/domains/{rule_id} | 
[**TeamsRegionsAppsDomainsList**](OpenapiGatewayApi.md#TeamsRegionsAppsDomainsList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/domains | 
[**TeamsRegionsAppsDomainsUpdate**](OpenapiGatewayApi.md#TeamsRegionsAppsDomainsUpdate) | **Put** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/domains/{rule_id} | 
[**TeamsRegionsAppsHttpdomainsCreate**](OpenapiGatewayApi.md#TeamsRegionsAppsHttpdomainsCreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
[**TeamsRegionsAppsHttpdomainsDelete**](OpenapiGatewayApi.md#TeamsRegionsAppsHttpdomainsDelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
[**TeamsRegionsAppsHttpdomainsList**](OpenapiGatewayApi.md#TeamsRegionsAppsHttpdomainsList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
[**TeamsRegionsAppsHttpdomainsRead**](OpenapiGatewayApi.md#TeamsRegionsAppsHttpdomainsRead) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
[**TeamsRegionsAppsHttpdomainsUpdate**](OpenapiGatewayApi.md#TeamsRegionsAppsHttpdomainsUpdate) | **Put** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 



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
**appId** | **int32**| 应用ID | 

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
**appId** | **int32**| 应用ID | 
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


## TeamsRegionsAppsHttpdomainsCreate

> HttpGatewayRule TeamsRegionsAppsHttpdomainsCreate(ctx, regionName, teamId, appId, data)



创建HTTP网关策略

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用ID | 
**data** | [**PostHttpGatewayRule**](PostHttpGatewayRule.md)|  | 

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


## TeamsRegionsAppsHttpdomainsDelete

> HttpGatewayRule TeamsRegionsAppsHttpdomainsDelete(ctx, appId, regionName, ruleId, teamId)



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


## TeamsRegionsAppsHttpdomainsList

> []HttpGatewayRule TeamsRegionsAppsHttpdomainsList(ctx, regionName, teamId, appId)



获取应用http访问策略列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用ID | 

### Return type

[**[]HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsHttpdomainsRead

> HttpGatewayRule TeamsRegionsAppsHttpdomainsRead(ctx, regionName, teamId, appId, ruleId)



获取应用http访问策略详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用ID | 
**ruleId** | **string**| 网关策略id | 

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


## TeamsRegionsAppsHttpdomainsUpdate

> HttpGatewayRule TeamsRegionsAppsHttpdomainsUpdate(ctx, regionName, teamId, appId, ruleId, data)



更新HTTP访问策略

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用ID | 
**ruleId** | **string**| 网关策略id | 
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

