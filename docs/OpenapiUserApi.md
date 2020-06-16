# \OpenapiUserApi

All URIs are relative to *http://0.0.0.0:7070/openapi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdministratorsCreate**](OpenapiUserApi.md#AdministratorsCreate) | **Post** /administrators | 
[**AdministratorsDelete**](OpenapiUserApi.md#AdministratorsDelete) | **Delete** /administrators/{user_id} | 
[**AdministratorsList**](OpenapiUserApi.md#AdministratorsList) | **Get** /administrators | 
[**ChangepwdUpdate**](OpenapiUserApi.md#ChangepwdUpdate) | **Put** /changepwd | 
[**UsersChangepwdUpdate**](OpenapiUserApi.md#UsersChangepwdUpdate) | **Put** /users/{user_id}/changepwd | 
[**UsersCreate**](OpenapiUserApi.md#UsersCreate) | **Post** /users | 
[**UsersDelete**](OpenapiUserApi.md#UsersDelete) | **Delete** /users/{user_id} | 
[**UsersList**](OpenapiUserApi.md#UsersList) | **Get** /users | 
[**UsersRead**](OpenapiUserApi.md#UsersRead) | **Get** /users/{user_id} | 
[**UsersUpdate**](OpenapiUserApi.md#UsersUpdate) | **Put** /users/{user_id} | 



## AdministratorsCreate

> CreateAdminUserReq AdministratorsCreate(ctx, data)



添加企业用户

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**CreateAdminUserReq**](CreateAdminUserReq.md)|  | 

### Return type

[**CreateAdminUserReq**](CreateAdminUserReq.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdministratorsDelete

> AdministratorsDelete(ctx, userId)



删除企业管理员

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdministratorsList

> ListUsersRespView AdministratorsList(ctx, optional)



获取企业管理员列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdministratorsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AdministratorsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **optional.String**| 企业ID | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListUsersRespView**](ListUsersRespView.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangepwdUpdate

> ChangePassWd ChangepwdUpdate(ctx, data)



修改自己账号密码

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**ChangePassWd**](ChangePassWd.md)|  | 

### Return type

[**ChangePassWd**](ChangePassWd.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersChangepwdUpdate

> ChangePassWdUser UsersChangepwdUpdate(ctx, userId, data)



修改用户密码

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 
**data** | [**ChangePassWdUser**](ChangePassWdUser.md)|  | 

### Return type

[**ChangePassWdUser**](ChangePassWdUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCreate

> CreateUser UsersCreate(ctx, data)



添加普通用户

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | [**CreateUser**](CreateUser.md)|  | 

### Return type

[**CreateUser**](CreateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDelete

> UsersDelete(ctx, userId)



删除用户

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

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


## UsersList

> ListUsersRespView UsersList(ctx, optional)



获取用户列表

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| 用户名、邮箱、手机号搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListUsersRespView**](ListUsersRespView.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRead

> UserInfo UsersRead(ctx, userId)



根据用户ID获取用户信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdate

> UpdateUser UsersUpdate(ctx, userId, data)



更新用户信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 
**data** | [**UpdateUser**](UpdateUser.md)|  | 

### Return type

[**UpdateUser**](UpdateUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

