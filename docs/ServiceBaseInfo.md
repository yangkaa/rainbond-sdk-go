# ServiceBaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | 组件状态 | [optional] [default to ]
**ServiceId** | **string** | 组件id | 
**TenantId** | **string** | 租户id | 
**ServiceKey** | **string** | 组件key | 
**ServiceAlias** | **string** | 组件别名 | 
**ServiceCname** | **string** | 组件名 | [optional] 
**ServiceRegion** | **string** | 组件所属区 | 
**Desc** | Pointer to **string** | 描述 | [optional] 
**Category** | **string** | 组件分类：application,cache,store | 
**Version** | **string** | 版本 | 
**UpdateVersion** | **int32** | 内部发布次数 | [optional] 
**Image** | **string** | 镜像 | 
**Cmd** | Pointer to **string** | 启动参数 | [optional] 
**ExtendMethod** | **string** | 组件部署类型,stateless or state | [optional] 
**MinNode** | **int32** | 启动个数 | [optional] 
**MinCpu** | **int32** | cpu个数 | [optional] 
**MinMemory** | **int32** | 内存大小单位（M） | [optional] 
**CodeFrom** | Pointer to **string** | 代码来源:gitlab,github | [optional] 
**GitUrl** | Pointer to **string** | code代码仓库 | [optional] 
**GitProjectId** | **int32** | gitlab 中项目id | [optional] 
**CodeVersion** | Pointer to **string** | 代码版本 | [optional] 
**ServiceType** | Pointer to **string** | 组件类型:web,mysql,redis,mongodb,phpadmin | [optional] 
**Creater** | **int32** | 组件创建者 | [optional] 
**Language** | Pointer to **string** | 代码语言 | [optional] 
**TotalMemory** | **int32** | 内存使用M | [optional] 
**IsService** | **bool** | 是否inner组件 | [optional] 
**ServiceOrigin** | **string** | 组件创建类型cloud云市组件,assistant云帮组件 | [optional] 
**TenantServiceGroupId** | **int32** | 组件归属的组件组id | [optional] 
**OpenWebhooks** | **bool** | 是否开启自动触发部署功能（兼容老版本组件） | [optional] 
**ServiceSource** | Pointer to **string** | 组件来源(source_code, market, docker_run, docker_compose) | [optional] 
**CreateStatus** | Pointer to **string** | 组件创建状态 creating|complete | [optional] 
**CheckUuid** | Pointer to **string** | 组件检测ID | [optional] 
**CheckEventId** | Pointer to **string** | 组件检测事件ID | [optional] 
**DockerCmd** | Pointer to **string** | 镜像创建命令 | [optional] 
**ServerType** | **string** | 源码仓库类型 | [optional] 
**IsUpgrate** | **bool** | 是否可以更新 | [optional] 
**BuildUpgrade** | **bool** | 组件构建后是否升级 | [optional] 
**OauthServiceId** | Pointer to **int32** | 拉取源码所用的OAuth服务id | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


