# CreateThirdComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServiceKey** | **string** | API 授权Key, 类型为api时有效 | 
**BuildUpgrade** | **bool** | 组件构建后是否升级 | [optional] 
**Category** | **string** | 组件分类：application,cache,store | 
**CheckEventId** | Pointer to **string** | 组件检测事件ID | [optional] 
**CheckUuid** | Pointer to **string** | 组件检测ID | [optional] 
**Cmd** | Pointer to **string** | 启动参数 | [optional] 
**CodeFrom** | Pointer to **string** | 代码来源:gitlab,github | [optional] 
**CodeVersion** | Pointer to **string** | 代码版本 | [optional] 
**CreateStatus** | Pointer to **string** | 组件创建状态 creating|complete | [optional] 
**Creater** | **int32** | 组件创建者 | [optional] 
**Desc** | Pointer to **string** | 描述 | [optional] 
**DockerCmd** | Pointer to **string** | 镜像创建命令 | [optional] 
**ExtendMethod** | **string** | 组件部署类型,stateless or state | [optional] 
**GitProjectId** | **int32** | gitlab 中项目id | [optional] 
**GitUrl** | Pointer to **string** | code代码仓库 | [optional] 
**Image** | **string** | 镜像 | 
**IsService** | **bool** | 是否inner组件 | [optional] 
**IsUpgrate** | **bool** | 是否可以更新 | [optional] 
**Language** | Pointer to **string** | 代码语言 | [optional] 
**MinCpu** | **int32** | cpu个数 | [optional] 
**MinMemory** | **int32** | 内存大小单位（M） | [optional] 
**MinNode** | **int32** | 启动个数 | [optional] 
**OauthServiceId** | Pointer to **int32** | 拉取源码所用的OAuth服务id | [optional] 
**OpenWebhooks** | **bool** | 是否开启自动触发部署功能（兼容老版本组件） | [optional] 
**ServerType** | **string** | 源码仓库类型 | [optional] 
**ServiceAlias** | **string** | 组件别名 | 
**ServiceCname** | **string** | 组件名 | [optional] 
**ServiceId** | **string** | 组件id | 
**ServiceKey** | **string** | 组件key | 
**ServiceOrigin** | **string** | 组件创建类型cloud云市组件,assistant云帮组件 | [optional] 
**ServiceRegion** | **string** | 组件所属区 | 
**ServiceSource** | Pointer to **string** | 组件来源(source_code, market, docker_run, docker_compose) | [optional] 
**ServiceType** | Pointer to **string** | 组件类型:web,mysql,redis,mongodb,phpadmin | [optional] 
**Status** | **string** | 组件状态 | [optional] [default to ]
**TenantId** | **string** | 租户id | 
**TenantServiceGroupId** | **int32** | 组件归属的组件组id | [optional] 
**TotalMemory** | **int32** | 内存使用M | [optional] 
**UpdateVersion** | **int32** | 内部发布次数 | [optional] 
**Url** | **string** | API地址, 类型为api时有效 | 
**Version** | **string** | 版本 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


