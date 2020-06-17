# \OpenapiAppsApi

All URIs are relative to *http://0.0.0.0:7070/openapi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsRegionsAppsCloseCreate**](OpenapiAppsApi.md#TeamsRegionsAppsCloseCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/close | 
[**TeamsRegionsAppsCopyCreate**](OpenapiAppsApi.md#TeamsRegionsAppsCopyCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
[**TeamsRegionsAppsCopyList**](OpenapiAppsApi.md#TeamsRegionsAppsCopyList) | **Get** /teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
[**TeamsRegionsAppsCreate**](OpenapiAppsApi.md#TeamsRegionsAppsCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps | 
[**TeamsRegionsAppsDelete**](OpenapiAppsApi.md#TeamsRegionsAppsDelete) | **Delete** /teams/{team_id}/regions/{region_name}/apps/{app_id} | 
[**TeamsRegionsAppsHttpdomainsCreate**](OpenapiAppsApi.md#TeamsRegionsAppsHttpdomainsCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
[**TeamsRegionsAppsInstallCreate**](OpenapiAppsApi.md#TeamsRegionsAppsInstallCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/{app_id}/install | 
[**TeamsRegionsAppsList**](OpenapiAppsApi.md#TeamsRegionsAppsList) | **Get** /teams/{team_id}/regions/{region_name}/apps | 
[**TeamsRegionsAppsOperationsCreate**](OpenapiAppsApi.md#TeamsRegionsAppsOperationsCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/{app_id}/operations | 
[**TeamsRegionsAppsRead**](OpenapiAppsApi.md#TeamsRegionsAppsRead) | **Get** /teams/{team_id}/regions/{region_name}/apps/{app_id} | 
[**TeamsRegionsAppsServicesDelete**](OpenapiAppsApi.md#TeamsRegionsAppsServicesDelete) | **Delete** /teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id} | 
[**TeamsRegionsAppsServicesEventsList**](OpenapiAppsApi.md#TeamsRegionsAppsServicesEventsList) | **Get** /teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/events | 
[**TeamsRegionsAppsServicesList**](OpenapiAppsApi.md#TeamsRegionsAppsServicesList) | **Get** /teams/{team_id}/regions/{region_name}/apps/{app_id}/services | 
[**TeamsRegionsAppsServicesRead**](OpenapiAppsApi.md#TeamsRegionsAppsServicesRead) | **Get** /teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id} | 
[**TeamsRegionsAppsServicesTelescopicHorizontalCreate**](OpenapiAppsApi.md#TeamsRegionsAppsServicesTelescopicHorizontalCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/telescopic/horizontal | 
[**TeamsRegionsAppsServicesTelescopicVerticalCreate**](OpenapiAppsApi.md#TeamsRegionsAppsServicesTelescopicVerticalCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/telescopic/vertical | 
[**TeamsRegionsAppsUpgradeCreate**](OpenapiAppsApi.md#TeamsRegionsAppsUpgradeCreate) | **Post** /teams/{team_id}/regions/{region_name}/apps/{app_id}/upgrade | 
[**TeamsRegionsAppsUpgradeList**](OpenapiAppsApi.md#TeamsRegionsAppsUpgradeList) | **Get** /teams/{team_id}/regions/{region_name}/apps/{app_id}/upgrade | 
[**TeamsRegionsResourceList**](OpenapiAppsApi.md#TeamsRegionsResourceList) | **Get** /teams/{team_id}/regions/{region_name}/resource | 



## TeamsRegionsAppsCloseCreate

> TeamAppsCloseSerializers TeamsRegionsAppsCloseCreate(ctx, regionName, teamId, data)



批量关闭应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**data** | [**TeamAppsCloseSerializers**](TeamAppsCloseSerializers.md)|  | 

### Return type

[**TeamAppsCloseSerializers**](TeamAppsCloseSerializers.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsCopyCreate

> AppCopyCRes TeamsRegionsAppsCopyCreate(ctx, appId, regionName, teamId, data)



复制应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**data** | [**AppCopyC**](AppCopyC.md)|  | 

### Return type

[**AppCopyCRes**](AppCopyCRes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsCopyList

> []AppCopyL TeamsRegionsAppsCopyList(ctx, appId, regionName, teamId)



获取需要复制的应用组件信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**regionName** | **string**|  | 
**teamId** | **string**|  | 

### Return type

[**[]AppCopyL**](AppCopyL.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsCreate

> AppBaseInfo TeamsRegionsAppsCreate(ctx, regionName, teamId, data)



创建应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**data** | [**AppPostInfo**](AppPostInfo.md)|  | 

### Return type

[**AppBaseInfo**](AppBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsDelete

> TeamsRegionsAppsDelete(ctx, regionName, teamId, appId, optional)



删除应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
 **optional** | ***TeamsRegionsAppsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsRegionsAppsDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **optional.Int32**| 强制删除 | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsHttpdomainsCreate

> HttpGatewayRule TeamsRegionsAppsHttpdomainsCreate(ctx, appId, regionName, teamId, data)



创建HTTP网关策略

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**regionName** | **string**|  | 
**teamId** | **string**|  | 
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


## TeamsRegionsAppsInstallCreate

> MarketInstall TeamsRegionsAppsInstallCreate(ctx, regionName, teamId, appId, data)



安装云市应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
**data** | [**Install**](Install.md)|  | 

### Return type

[**MarketInstall**](MarketInstall.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsList

> []AppBaseInfo TeamsRegionsAppsList(ctx, regionName, teamId, optional)



团队应用列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
 **optional** | ***TeamsRegionsAppsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsRegionsAppsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **optional.String**| 搜索查询应用名称，团队名称 | 

### Return type

[**[]AppBaseInfo**](AppBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsOperationsCreate

> Success TeamsRegionsAppsOperationsCreate(ctx, regionName, teamId, appId, data)



操作应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
**data** | [**ServiceGroupOperations**](ServiceGroupOperations.md)|  | 

### Return type

[**Success**](Success.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsRead

> AppInfo TeamsRegionsAppsRead(ctx, regionName, teamId, appId)



应用详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 

### Return type

[**AppInfo**](AppInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsServicesDelete

> TeamsRegionsAppsServicesDelete(ctx, regionName, serviceId, teamId, appId, optional)



删除组件

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**serviceId** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
 **optional** | ***TeamsRegionsAppsServicesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsRegionsAppsServicesDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **force** | **optional.Int32**| 强制删除 | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsServicesEventsList

> ListServiceEventsResponse TeamsRegionsAppsServicesEventsList(ctx, regionName, serviceId, teamId, appId, optional)



查询组件事件信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**serviceId** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
 **optional** | ***TeamsRegionsAppsServicesEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsRegionsAppsServicesEventsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **page** | **optional.Int32**| 页码 | 
 **pageSize** | **optional.Int32**| 每页数量 | 

### Return type

[**ListServiceEventsResponse**](ListServiceEventsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsServicesList

> []ServiceBaseInfo TeamsRegionsAppsServicesList(ctx, regionName, teamId, appId)



查询应用下组件列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 

### Return type

[**[]ServiceBaseInfo**](ServiceBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsServicesRead

> ServiceBaseInfo TeamsRegionsAppsServicesRead(ctx, regionName, serviceId, teamId, appId)



查询组件信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**serviceId** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 

### Return type

[**ServiceBaseInfo**](ServiceBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsServicesTelescopicHorizontalCreate

> AppServiceTelescopicHorizontal TeamsRegionsAppsServicesTelescopicHorizontalCreate(ctx, regionName, serviceId, teamId, appId, data)



组件水平伸缩

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**serviceId** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
**data** | [**AppServiceTelescopicHorizontal**](AppServiceTelescopicHorizontal.md)|  | 

### Return type

[**AppServiceTelescopicHorizontal**](AppServiceTelescopicHorizontal.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsServicesTelescopicVerticalCreate

> AppServiceTelescopicVertical TeamsRegionsAppsServicesTelescopicVerticalCreate(ctx, regionName, serviceId, teamId, appId, data)



组件垂直伸缩

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**serviceId** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
**data** | [**AppServiceTelescopicVertical**](AppServiceTelescopicVertical.md)|  | 

### Return type

[**AppServiceTelescopicVertical**](AppServiceTelescopicVertical.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsUpgradeCreate

> ListUpgrade TeamsRegionsAppsUpgradeCreate(ctx, regionName, teamId, appId, data)



升级应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 
**data** | [**Upgrade**](Upgrade.md)|  | 

### Return type

[**ListUpgrade**](ListUpgrade.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsAppsUpgradeList

> ListUpgrade TeamsRegionsAppsUpgradeList(ctx, regionName, teamId, appId)



升级应用

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 
**appId** | **int32**| 应用组id | 

### Return type

[**ListUpgrade**](ListUpgrade.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRegionsResourceList

> TeamAppsResource TeamsRegionsResourceList(ctx, regionName, teamId)



获取团队资源统计

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionName** | **string**|  | 
**teamId** | **string**|  | 

### Return type

[**TeamAppsResource**](TeamAppsResource.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

