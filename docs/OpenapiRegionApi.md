# \OpenapiRegionApi

All URIs are relative to *http://0.0.0.0:7070/openapi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegionsCreate**](OpenapiRegionApi.md#RegionsCreate) | **Post** /regions | 
[**RegionsList**](OpenapiRegionApi.md#RegionsList) | **Get** /regions | 
[**RegionsRead**](OpenapiRegionApi.md#RegionsRead) | **Get** /regions/{region_id} | 



## RegionsCreate

> RegionInfo RegionsCreate(ctx, data)



添加数据中心

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**AddRegionRequest**](AddRegionRequest.md)|  | 

### Return type

[**RegionInfo**](RegionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsList

> []RegionInfoResp RegionsList(ctx, )



获取全部数据中心列表

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RegionInfoResp**](RegionInfoResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsRead

> RegionInfoR RegionsRead(ctx, regionId, optional)



获取指定数据中心数据

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **string**|  | 
 **optional** | ***RegionsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegionsReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId2** | **optional.String**| 数据中心名称、id | 
 **extendInfo** | **optional.String**| 是否需要额外数据 | 

### Return type

[**RegionInfoR**](RegionInfoR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

