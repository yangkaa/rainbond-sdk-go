# \OpenapiTeamRegionApi

All URIs are relative to *http://127.0.0.1:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsRegionsCreate**](OpenapiTeamRegionApi.md#TeamsRegionsCreate) | **Post** /openapi/v1/teams/{team_id}/regions | 
[**TeamsRegionsList**](OpenapiTeamRegionApi.md#TeamsRegionsList) | **Get** /openapi/v1/teams/{team_id}/regions | 



## TeamsRegionsCreate

> TeamBaseInfo TeamsRegionsCreate(ctx, teamId, data)



开通数据中心

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
**data** | [**TeamRegionReq**](TeamRegionReq.md)|  | 

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


## TeamsRegionsList

> ListTeamRegionsResp TeamsRegionsList(ctx, teamId, optional)



获取团队开通的数据中心列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
 **optional** | ***TeamsRegionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsRegionsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| 根据数据中心名称搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListTeamRegionsResp**](ListTeamRegionsResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

