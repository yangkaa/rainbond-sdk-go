# \OpenapiTeamApi

All URIs are relative to *http://127.0.0.1:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsCertificatesCreate**](OpenapiTeamApi.md#TeamsCertificatesCreate) | **Post** /openapi/v1/teams/{team_id}/certificates | 
[**TeamsCertificatesDelete**](OpenapiTeamApi.md#TeamsCertificatesDelete) | **Delete** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**TeamsCertificatesList**](OpenapiTeamApi.md#TeamsCertificatesList) | **Get** /openapi/v1/teams/{team_id}/certificates | 
[**TeamsCertificatesRead**](OpenapiTeamApi.md#TeamsCertificatesRead) | **Get** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**TeamsCertificatesUpdate**](OpenapiTeamApi.md#TeamsCertificatesUpdate) | **Put** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**TeamsCreate**](OpenapiTeamApi.md#TeamsCreate) | **Post** /openapi/v1/teams | 
[**TeamsDelete**](OpenapiTeamApi.md#TeamsDelete) | **Delete** /openapi/v1/teams/{team_id} | 
[**TeamsList**](OpenapiTeamApi.md#TeamsList) | **Get** /openapi/v1/teams | 
[**TeamsRead**](OpenapiTeamApi.md#TeamsRead) | **Get** /openapi/v1/teams/{team_id} | 
[**TeamsRegionsResourceList**](OpenapiTeamApi.md#TeamsRegionsResourceList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/resource | 
[**TeamsResourceCreate**](OpenapiTeamApi.md#TeamsResourceCreate) | **Post** /openapi/v1/teams/resource | 
[**TeamsUpdate**](OpenapiTeamApi.md#TeamsUpdate) | **Put** /openapi/v1/teams/{team_id} | 



## TeamsCertificatesCreate

> TeamCertificatesR TeamsCertificatesCreate(ctx, teamId, data)



添加证书

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
**data** | [**TeamCertificatesC**](TeamCertificatesC.md)|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCertificatesDelete

> TeamsCertificatesDelete(ctx, certificateId, teamId)



删除证书

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string**|  | 
**teamId** | **string**|  | 

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


## TeamsCertificatesList

> TeamCertificatesL TeamsCertificatesList(ctx, teamId, optional)



获取团队下证书列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
 **optional** | ***TeamsCertificatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsCertificatesListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| 页码 | 
 **pageSize** | **optional.Float32**| 每页数量 | 

### Return type

[**TeamCertificatesL**](TeamCertificatesL.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCertificatesRead

> TeamCertificatesR TeamsCertificatesRead(ctx, certificateId, teamId)



获取团队下证书列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string**|  | 
**teamId** | **string**|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCertificatesUpdate

> TeamCertificatesR TeamsCertificatesUpdate(ctx, certificateId, teamId, data)



更新证书

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string**|  | 
**teamId** | **string**|  | 
**data** | [**TeamCertificatesC**](TeamCertificatesC.md)|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreate

> TeamBaseInfo TeamsCreate(ctx, data)



add team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**CreateTeamReq**](CreateTeamReq.md)|  | 

### Return type

[**TeamBaseInfo**](TeamBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDelete

> TeamsDelete(ctx, teamId, optional)



删除团队

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
 **optional** | ***TeamsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.String**| 团队名称搜索 | 

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


## TeamsList

> ListTeamResp TeamsList(ctx, optional)



获取用户所在团队列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| 团队名称搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListTeamResp**](ListTeamResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsRead

> TeamInfo TeamsRead(ctx, teamId)



获取团队

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 

### Return type

[**TeamInfo**](TeamInfo.md)

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


## TeamsResourceCreate

> []TeamAppsResource TeamsResourceCreate(ctx, data)



批量获取团队资源统计列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**[]TenantRegionList**](TenantRegionList.md)|  | 

### Return type

[**[]TeamAppsResource**](TeamAppsResource.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdate

> UpdateTeamInfoReq TeamsUpdate(ctx, teamId, data)



更新团队信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
**data** | [**UpdateTeamInfoReq**](UpdateTeamInfoReq.md)|  | 

### Return type

[**UpdateTeamInfoReq**](UpdateTeamInfoReq.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

