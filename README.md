# Go API client for openapi

Rainbond open api

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:8000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OpenapiAppsApi* | [**TeamsRegionsAppsCloseCreate**](docs/OpenapiAppsApi.md#teamsregionsappsclosecreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/close | 
*OpenapiAppsApi* | [**TeamsRegionsAppsCopyCreate**](docs/OpenapiAppsApi.md#teamsregionsappscopycreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
*OpenapiAppsApi* | [**TeamsRegionsAppsCopyList**](docs/OpenapiAppsApi.md#teamsregionsappscopylist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
*OpenapiAppsApi* | [**TeamsRegionsAppsCreate**](docs/OpenapiAppsApi.md#teamsregionsappscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps | 
*OpenapiAppsApi* | [**TeamsRegionsAppsDelete**](docs/OpenapiAppsApi.md#teamsregionsappsdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id} | 
*OpenapiAppsApi* | [**TeamsRegionsAppsDomainsCreate**](docs/OpenapiAppsApi.md#teamsregionsappsdomainscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/domains | 
*OpenapiAppsApi* | [**TeamsRegionsAppsInstallCreate**](docs/OpenapiAppsApi.md#teamsregionsappsinstallcreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/install | 
*OpenapiAppsApi* | [**TeamsRegionsAppsList**](docs/OpenapiAppsApi.md#teamsregionsappslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps | 
*OpenapiAppsApi* | [**TeamsRegionsAppsMonitorQueryList**](docs/OpenapiAppsApi.md#teamsregionsappsmonitorquerylist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/monitor/query | 
*OpenapiAppsApi* | [**TeamsRegionsAppsMonitorQueryRangeList**](docs/OpenapiAppsApi.md#teamsregionsappsmonitorqueryrangelist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/monitor/query_range | 
*OpenapiAppsApi* | [**TeamsRegionsAppsOperationsCreate**](docs/OpenapiAppsApi.md#teamsregionsappsoperationscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/operations | 
*OpenapiAppsApi* | [**TeamsRegionsAppsRead**](docs/OpenapiAppsApi.md#teamsregionsappsread) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id} | 
*OpenapiAppsApi* | [**TeamsRegionsAppsServicesDelete**](docs/OpenapiAppsApi.md#teamsregionsappsservicesdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id} | 
*OpenapiAppsApi* | [**TeamsRegionsAppsServicesEnvsUpdate**](docs/OpenapiAppsApi.md#teamsregionsappsservicesenvsupdate) | **Put** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/envs | 
*OpenapiAppsApi* | [**TeamsRegionsAppsServicesEventsList**](docs/OpenapiAppsApi.md#teamsregionsappsserviceseventslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/events | 
*OpenapiAppsApi* | [**TeamsRegionsAppsServicesList**](docs/OpenapiAppsApi.md#teamsregionsappsserviceslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services | 
*OpenapiAppsApi* | [**TeamsRegionsAppsServicesRead**](docs/OpenapiAppsApi.md#teamsregionsappsservicesread) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id} | 
*OpenapiAppsApi* | [**TeamsRegionsAppsServicesTelescopicHorizontalCreate**](docs/OpenapiAppsApi.md#teamsregionsappsservicestelescopichorizontalcreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/telescopic/horizontal | 
*OpenapiAppsApi* | [**TeamsRegionsAppsServicesTelescopicVerticalCreate**](docs/OpenapiAppsApi.md#teamsregionsappsservicestelescopicverticalcreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/telescopic/vertical | 
*OpenapiAppsApi* | [**TeamsRegionsAppsThirdComponentsCreate**](docs/OpenapiAppsApi.md#teamsregionsappsthirdcomponentscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/third-components | 
*OpenapiAppsApi* | [**TeamsRegionsAppsUpgradeCreate**](docs/OpenapiAppsApi.md#teamsregionsappsupgradecreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/upgrade | 
*OpenapiAppsApi* | [**TeamsRegionsAppsUpgradeList**](docs/OpenapiAppsApi.md#teamsregionsappsupgradelist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/upgrade | 
*OpenapiEntrepriseApi* | [**ConfigsList**](docs/OpenapiEntrepriseApi.md#configslist) | **Get** /openapi/v1/configs | 
*OpenapiGatewayApi* | [**HttpdomainsList**](docs/OpenapiGatewayApi.md#httpdomainslist) | **Get** /openapi/v1/httpdomains | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsDomainsDelete**](docs/OpenapiGatewayApi.md#teamsregionsappsdomainsdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/domains/{rule_id} | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsDomainsList**](docs/OpenapiGatewayApi.md#teamsregionsappsdomainslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/domains | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsDomainsUpdate**](docs/OpenapiGatewayApi.md#teamsregionsappsdomainsupdate) | **Put** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/domains/{rule_id} | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsHttpdomainsCreate**](docs/OpenapiGatewayApi.md#teamsregionsappshttpdomainscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsHttpdomainsDelete**](docs/OpenapiGatewayApi.md#teamsregionsappshttpdomainsdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsHttpdomainsList**](docs/OpenapiGatewayApi.md#teamsregionsappshttpdomainslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsHttpdomainsRead**](docs/OpenapiGatewayApi.md#teamsregionsappshttpdomainsread) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
*OpenapiGatewayApi* | [**TeamsRegionsAppsHttpdomainsUpdate**](docs/OpenapiGatewayApi.md#teamsregionsappshttpdomainsupdate) | **Put** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
*OpenapiRegionApi* | [**RegionsCreate**](docs/OpenapiRegionApi.md#regionscreate) | **Post** /openapi/v1/regions | 
*OpenapiRegionApi* | [**RegionsList**](docs/OpenapiRegionApi.md#regionslist) | **Get** /openapi/v1/regions | 
*OpenapiRegionApi* | [**RegionsRead**](docs/OpenapiRegionApi.md#regionsread) | **Get** /openapi/v1/regions/{region_id} | 
*OpenapiTeamApi* | [**TeamsCertificatesCreate**](docs/OpenapiTeamApi.md#teamscertificatescreate) | **Post** /openapi/v1/teams/{team_id}/certificates | 
*OpenapiTeamApi* | [**TeamsCertificatesDelete**](docs/OpenapiTeamApi.md#teamscertificatesdelete) | **Delete** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
*OpenapiTeamApi* | [**TeamsCertificatesList**](docs/OpenapiTeamApi.md#teamscertificateslist) | **Get** /openapi/v1/teams/{team_id}/certificates | 
*OpenapiTeamApi* | [**TeamsCertificatesRead**](docs/OpenapiTeamApi.md#teamscertificatesread) | **Get** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
*OpenapiTeamApi* | [**TeamsCertificatesUpdate**](docs/OpenapiTeamApi.md#teamscertificatesupdate) | **Put** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
*OpenapiTeamApi* | [**TeamsCreate**](docs/OpenapiTeamApi.md#teamscreate) | **Post** /openapi/v1/teams | 
*OpenapiTeamApi* | [**TeamsDelete**](docs/OpenapiTeamApi.md#teamsdelete) | **Delete** /openapi/v1/teams/{team_id} | 
*OpenapiTeamApi* | [**TeamsList**](docs/OpenapiTeamApi.md#teamslist) | **Get** /openapi/v1/teams | 
*OpenapiTeamApi* | [**TeamsRead**](docs/OpenapiTeamApi.md#teamsread) | **Get** /openapi/v1/teams/{team_id} | 
*OpenapiTeamApi* | [**TeamsRegionsOverviewList**](docs/OpenapiTeamApi.md#teamsregionsoverviewlist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/overview | 
*OpenapiTeamApi* | [**TeamsRegionsResourceList**](docs/OpenapiTeamApi.md#teamsregionsresourcelist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/resource | 
*OpenapiTeamApi* | [**TeamsResourceCreate**](docs/OpenapiTeamApi.md#teamsresourcecreate) | **Post** /openapi/v1/teams/resource | 
*OpenapiTeamApi* | [**TeamsUpdate**](docs/OpenapiTeamApi.md#teamsupdate) | **Put** /openapi/v1/teams/{team_id} | 
*OpenapiTeamRegionApi* | [**TeamsRegionsCreate**](docs/OpenapiTeamRegionApi.md#teamsregionscreate) | **Post** /openapi/v1/teams/{team_id}/regions | 
*OpenapiTeamRegionApi* | [**TeamsRegionsList**](docs/OpenapiTeamRegionApi.md#teamsregionslist) | **Get** /openapi/v1/teams/{team_id}/regions | 
*OpenapiUserApi* | [**AdministratorsCreate**](docs/OpenapiUserApi.md#administratorscreate) | **Post** /openapi/v1/administrators | 
*OpenapiUserApi* | [**AdministratorsDelete**](docs/OpenapiUserApi.md#administratorsdelete) | **Delete** /openapi/v1/administrators/{user_id} | 
*OpenapiUserApi* | [**AdministratorsList**](docs/OpenapiUserApi.md#administratorslist) | **Get** /openapi/v1/administrators | 
*OpenapiUserApi* | [**ChangepwdUpdate**](docs/OpenapiUserApi.md#changepwdupdate) | **Put** /openapi/v1/changepwd | 
*OpenapiUserApi* | [**UsersChangepwdUpdate**](docs/OpenapiUserApi.md#userschangepwdupdate) | **Put** /openapi/v1/users/{user_id}/changepwd | 
*OpenapiUserApi* | [**UsersCreate**](docs/OpenapiUserApi.md#userscreate) | **Post** /openapi/v1/users | 
*OpenapiUserApi* | [**UsersDelete**](docs/OpenapiUserApi.md#usersdelete) | **Delete** /openapi/v1/users/{user_id} | 
*OpenapiUserApi* | [**UsersList**](docs/OpenapiUserApi.md#userslist) | **Get** /openapi/v1/users | 
*OpenapiUserApi* | [**UsersRead**](docs/OpenapiUserApi.md#usersread) | **Get** /openapi/v1/users/{user_id} | 
*OpenapiUserApi* | [**UsersUpdate**](docs/OpenapiUserApi.md#usersupdate) | **Put** /openapi/v1/users/{user_id} | 


## Documentation For Models

 - [AddRegionRequest](docs/AddRegionRequest.md)
 - [AppBaseInfo](docs/AppBaseInfo.md)
 - [AppCopyC](docs/AppCopyC.md)
 - [AppCopyCRes](docs/AppCopyCRes.md)
 - [AppCopyL](docs/AppCopyL.md)
 - [AppCopyModify](docs/AppCopyModify.md)
 - [AppInfo](docs/AppInfo.md)
 - [AppPostInfo](docs/AppPostInfo.md)
 - [AppServiceEvents](docs/AppServiceEvents.md)
 - [AppServiceTelescopicHorizontal](docs/AppServiceTelescopicHorizontal.md)
 - [AppServiceTelescopicVertical](docs/AppServiceTelescopicVertical.md)
 - [AppStoreImageHubBaseResp](docs/AppStoreImageHubBaseResp.md)
 - [AppStoreImageHubResp](docs/AppStoreImageHubResp.md)
 - [AutoSsl](docs/AutoSsl.md)
 - [CertificatesR](docs/CertificatesR.md)
 - [ChangePassWd](docs/ChangePassWd.md)
 - [ChangePassWdUser](docs/ChangePassWdUser.md)
 - [CloudMarketBaseResp](docs/CloudMarketBaseResp.md)
 - [ComponentEnvsBaseSerializers](docs/ComponentEnvsBaseSerializers.md)
 - [ComponentEnvsSerializers](docs/ComponentEnvsSerializers.md)
 - [ComponentMonitorBaseSerializers](docs/ComponentMonitorBaseSerializers.md)
 - [ComponentMonitorItemsSerializers](docs/ComponentMonitorItemsSerializers.md)
 - [ComponentMonitorSerializers](docs/ComponentMonitorSerializers.md)
 - [CreateAdminUserReq](docs/CreateAdminUserReq.md)
 - [CreateTeamReq](docs/CreateTeamReq.md)
 - [CreateThirdComponent](docs/CreateThirdComponent.md)
 - [CreateThirdComponentResponse](docs/CreateThirdComponentResponse.md)
 - [CreateUser](docs/CreateUser.md)
 - [CustomJwt](docs/CustomJwt.md)
 - [EnterpriseConfigSeralizer](docs/EnterpriseConfigSeralizer.md)
 - [EnterpriseHttpGatewayRule](docs/EnterpriseHttpGatewayRule.md)
 - [ExportAppBaseResp](docs/ExportAppBaseResp.md)
 - [Fail](docs/Fail.md)
 - [GatewayRule](docs/GatewayRule.md)
 - [HttpConfiguration](docs/HttpConfiguration.md)
 - [HttpGatewayRule](docs/HttpGatewayRule.md)
 - [HttpHeader](docs/HttpHeader.md)
 - [Install](docs/Install.md)
 - [ListServiceEventsResponse](docs/ListServiceEventsResponse.md)
 - [ListTeamRegionsResp](docs/ListTeamRegionsResp.md)
 - [ListTeamResp](docs/ListTeamResp.md)
 - [ListUpgrade](docs/ListUpgrade.md)
 - [ListUsersRespView](docs/ListUsersRespView.md)
 - [MarketInstall](docs/MarketInstall.md)
 - [MonitorDataSerializers](docs/MonitorDataSerializers.md)
 - [NewBieGuideBaseResp](docs/NewBieGuideBaseResp.md)
 - [OauthServicesBaseResp](docs/OauthServicesBaseResp.md)
 - [OauthServicesResp](docs/OauthServicesResp.md)
 - [ObjectStorageBaseResp](docs/ObjectStorageBaseResp.md)
 - [ObjectStorageResp](docs/ObjectStorageResp.md)
 - [PostGatewayRule](docs/PostGatewayRule.md)
 - [PostHttpGatewayRule](docs/PostHttpGatewayRule.md)
 - [PostTcpGatewayRule](docs/PostTcpGatewayRule.md)
 - [PostTcpGatewayRuleExtensions](docs/PostTcpGatewayRuleExtensions.md)
 - [RegionInfo](docs/RegionInfo.md)
 - [RegionInfoR](docs/RegionInfoR.md)
 - [RegionInfoResp](docs/RegionInfoResp.md)
 - [RoleInfo](docs/RoleInfo.md)
 - [ServiceBaseInfo](docs/ServiceBaseInfo.md)
 - [ServiceGroupOperations](docs/ServiceGroupOperations.md)
 - [Success](docs/Success.md)
 - [TcpGatewayRule](docs/TcpGatewayRule.md)
 - [TeamAppsCloseSerializers](docs/TeamAppsCloseSerializers.md)
 - [TeamAppsResource](docs/TeamAppsResource.md)
 - [TeamBaseInfo](docs/TeamBaseInfo.md)
 - [TeamCertificatesC](docs/TeamCertificatesC.md)
 - [TeamCertificatesL](docs/TeamCertificatesL.md)
 - [TeamCertificatesR](docs/TeamCertificatesR.md)
 - [TeamInfo](docs/TeamInfo.md)
 - [TeamOverview](docs/TeamOverview.md)
 - [TeamRegionReq](docs/TeamRegionReq.md)
 - [TeamRegionsResp](docs/TeamRegionsResp.md)
 - [TenantRegionList](docs/TenantRegionList.md)
 - [UpdatePostHttpGatewayRule](docs/UpdatePostHttpGatewayRule.md)
 - [UpdateTeamInfoReq](docs/UpdateTeamInfoReq.md)
 - [UpdateUser](docs/UpdateUser.md)
 - [Upgrade](docs/Upgrade.md)
 - [UpgradeBase](docs/UpgradeBase.md)
 - [UserInfo](docs/UserInfo.md)


## Documentation For Authorization



## Bearer

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```



## Author

barnett@goodrain.com

