# Go API client for swagger

k8sdispatcher API generated from k8sdispatcher.yang

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*K8sdispatcherApi* | [**CreateK8sdispatcherByID**](docs/K8sdispatcherApi.md#createk8sdispatcherbyid) | **Post** /k8sdispatcher/{name}/ | Create k8sdispatcher by ID
*K8sdispatcherApi* | [**CreateK8sdispatcherNodeportRuleByID**](docs/K8sdispatcherApi.md#createk8sdispatchernodeportrulebyid) | **Post** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Create nodeport-rule by ID
*K8sdispatcherApi* | [**CreateK8sdispatcherNodeportRuleListByID**](docs/K8sdispatcherApi.md#createk8sdispatchernodeportrulelistbyid) | **Post** /k8sdispatcher/{name}/nodeport-rule/ | Create nodeport-rule by ID
*K8sdispatcherApi* | [**CreateK8sdispatcherPortsByID**](docs/K8sdispatcherApi.md#createk8sdispatcherportsbyid) | **Post** /k8sdispatcher/{name}/ports/{ports_name}/ | Create ports by ID
*K8sdispatcherApi* | [**CreateK8sdispatcherPortsListByID**](docs/K8sdispatcherApi.md#createk8sdispatcherportslistbyid) | **Post** /k8sdispatcher/{name}/ports/ | Create ports by ID
*K8sdispatcherApi* | [**DeleteK8sdispatcherByID**](docs/K8sdispatcherApi.md#deletek8sdispatcherbyid) | **Delete** /k8sdispatcher/{name}/ | Delete k8sdispatcher by ID
*K8sdispatcherApi* | [**DeleteK8sdispatcherNodeportRuleByID**](docs/K8sdispatcherApi.md#deletek8sdispatchernodeportrulebyid) | **Delete** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Delete nodeport-rule by ID
*K8sdispatcherApi* | [**DeleteK8sdispatcherNodeportRuleListByID**](docs/K8sdispatcherApi.md#deletek8sdispatchernodeportrulelistbyid) | **Delete** /k8sdispatcher/{name}/nodeport-rule/ | Delete nodeport-rule by ID
*K8sdispatcherApi* | [**DeleteK8sdispatcherPortsByID**](docs/K8sdispatcherApi.md#deletek8sdispatcherportsbyid) | **Delete** /k8sdispatcher/{name}/ports/{ports_name}/ | Delete ports by ID
*K8sdispatcherApi* | [**DeleteK8sdispatcherPortsListByID**](docs/K8sdispatcherApi.md#deletek8sdispatcherportslistbyid) | **Delete** /k8sdispatcher/{name}/ports/ | Delete ports by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherByID**](docs/K8sdispatcherApi.md#readk8sdispatcherbyid) | **Get** /k8sdispatcher/{name}/ | Read k8sdispatcher by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherInternalSrcIpByID**](docs/K8sdispatcherApi.md#readk8sdispatcherinternalsrcipbyid) | **Get** /k8sdispatcher/{name}/internal-src-ip/ | Read internal-src-ip by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherListByID**](docs/K8sdispatcherApi.md#readk8sdispatcherlistbyid) | **Get** /k8sdispatcher/ | Read k8sdispatcher by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherLoglevelByID**](docs/K8sdispatcherApi.md#readk8sdispatcherloglevelbyid) | **Get** /k8sdispatcher/{name}/loglevel/ | Read loglevel by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherNodeportRangeByID**](docs/K8sdispatcherApi.md#readk8sdispatchernodeportrangebyid) | **Get** /k8sdispatcher/{name}/nodeport-range/ | Read nodeport-range by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherNodeportRuleByID**](docs/K8sdispatcherApi.md#readk8sdispatchernodeportrulebyid) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Read nodeport-rule by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherNodeportRuleExternalTrafficPolicyByID**](docs/K8sdispatcherApi.md#readk8sdispatchernodeportruleexternaltrafficpolicybyid) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/external-traffic-policy/ | Read external-traffic-policy by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherNodeportRuleListByID**](docs/K8sdispatcherApi.md#readk8sdispatchernodeportrulelistbyid) | **Get** /k8sdispatcher/{name}/nodeport-rule/ | Read nodeport-rule by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherNodeportRuleRuleNameByID**](docs/K8sdispatcherApi.md#readk8sdispatchernodeportrulerulenamebyid) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/rule-name/ | Read rule-name by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherPortsByID**](docs/K8sdispatcherApi.md#readk8sdispatcherportsbyid) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/ | Read ports by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherPortsIpByID**](docs/K8sdispatcherApi.md#readk8sdispatcherportsipbyid) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/ip/ | Read ip by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherPortsListByID**](docs/K8sdispatcherApi.md#readk8sdispatcherportslistbyid) | **Get** /k8sdispatcher/{name}/ports/ | Read ports by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherPortsPeerByID**](docs/K8sdispatcherApi.md#readk8sdispatcherportspeerbyid) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/peer/ | Read peer by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherPortsStatusByID**](docs/K8sdispatcherApi.md#readk8sdispatcherportsstatusbyid) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/status/ | Read status by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherPortsTypeByID**](docs/K8sdispatcherApi.md#readk8sdispatcherportstypebyid) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/type/ | Read type by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherPortsUuidByID**](docs/K8sdispatcherApi.md#readk8sdispatcherportsuuidbyid) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/uuid/ | Read uuid by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherServiceNameByID**](docs/K8sdispatcherApi.md#readk8sdispatcherservicenamebyid) | **Get** /k8sdispatcher/{name}/service-name/ | Read service-name by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherSessionRuleByID**](docs/K8sdispatcherApi.md#readk8sdispatchersessionrulebyid) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/ | Read session-rule by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherSessionRuleListByID**](docs/K8sdispatcherApi.md#readk8sdispatchersessionrulelistbyid) | **Get** /k8sdispatcher/{name}/session-rule/ | Read session-rule by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherSessionRuleNewIpByID**](docs/K8sdispatcherApi.md#readk8sdispatchersessionrulenewipbyid) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/new-ip/ | Read new-ip by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherSessionRuleNewPortByID**](docs/K8sdispatcherApi.md#readk8sdispatchersessionrulenewportbyid) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/new-port/ | Read new-port by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherSessionRuleOperationByID**](docs/K8sdispatcherApi.md#readk8sdispatchersessionruleoperationbyid) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/operation/ | Read operation by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherSessionRuleOriginatingRuleByID**](docs/K8sdispatcherApi.md#readk8sdispatchersessionruleoriginatingrulebyid) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/originating-rule/ | Read originating-rule by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherShadowByID**](docs/K8sdispatcherApi.md#readk8sdispatchershadowbyid) | **Get** /k8sdispatcher/{name}/shadow/ | Read shadow by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherSpanByID**](docs/K8sdispatcherApi.md#readk8sdispatcherspanbyid) | **Get** /k8sdispatcher/{name}/span/ | Read span by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherTypeByID**](docs/K8sdispatcherApi.md#readk8sdispatchertypebyid) | **Get** /k8sdispatcher/{name}/type/ | Read type by ID
*K8sdispatcherApi* | [**ReadK8sdispatcherUuidByID**](docs/K8sdispatcherApi.md#readk8sdispatcheruuidbyid) | **Get** /k8sdispatcher/{name}/uuid/ | Read uuid by ID
*K8sdispatcherApi* | [**ReplaceK8sdispatcherByID**](docs/K8sdispatcherApi.md#replacek8sdispatcherbyid) | **Put** /k8sdispatcher/{name}/ | Replace k8sdispatcher by ID
*K8sdispatcherApi* | [**ReplaceK8sdispatcherNodeportRuleByID**](docs/K8sdispatcherApi.md#replacek8sdispatchernodeportrulebyid) | **Put** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Replace nodeport-rule by ID
*K8sdispatcherApi* | [**ReplaceK8sdispatcherNodeportRuleListByID**](docs/K8sdispatcherApi.md#replacek8sdispatchernodeportrulelistbyid) | **Put** /k8sdispatcher/{name}/nodeport-rule/ | Replace nodeport-rule by ID
*K8sdispatcherApi* | [**ReplaceK8sdispatcherPortsByID**](docs/K8sdispatcherApi.md#replacek8sdispatcherportsbyid) | **Put** /k8sdispatcher/{name}/ports/{ports_name}/ | Replace ports by ID
*K8sdispatcherApi* | [**ReplaceK8sdispatcherPortsListByID**](docs/K8sdispatcherApi.md#replacek8sdispatcherportslistbyid) | **Put** /k8sdispatcher/{name}/ports/ | Replace ports by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherByID**](docs/K8sdispatcherApi.md#updatek8sdispatcherbyid) | **Patch** /k8sdispatcher/{name}/ | Update k8sdispatcher by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherListByID**](docs/K8sdispatcherApi.md#updatek8sdispatcherlistbyid) | **Patch** /k8sdispatcher/ | Update k8sdispatcher by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherLoglevelByID**](docs/K8sdispatcherApi.md#updatek8sdispatcherloglevelbyid) | **Patch** /k8sdispatcher/{name}/loglevel/ | Update loglevel by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherNodeportRangeByID**](docs/K8sdispatcherApi.md#updatek8sdispatchernodeportrangebyid) | **Patch** /k8sdispatcher/{name}/nodeport-range/ | Update nodeport-range by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherNodeportRuleByID**](docs/K8sdispatcherApi.md#updatek8sdispatchernodeportrulebyid) | **Patch** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Update nodeport-rule by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherNodeportRuleExternalTrafficPolicyByID**](docs/K8sdispatcherApi.md#updatek8sdispatchernodeportruleexternaltrafficpolicybyid) | **Patch** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/external-traffic-policy/ | Update external-traffic-policy by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherNodeportRuleListByID**](docs/K8sdispatcherApi.md#updatek8sdispatchernodeportrulelistbyid) | **Patch** /k8sdispatcher/{name}/nodeport-rule/ | Update nodeport-rule by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherPortsByID**](docs/K8sdispatcherApi.md#updatek8sdispatcherportsbyid) | **Patch** /k8sdispatcher/{name}/ports/{ports_name}/ | Update ports by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherPortsListByID**](docs/K8sdispatcherApi.md#updatek8sdispatcherportslistbyid) | **Patch** /k8sdispatcher/{name}/ports/ | Update ports by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherPortsPeerByID**](docs/K8sdispatcherApi.md#updatek8sdispatcherportspeerbyid) | **Patch** /k8sdispatcher/{name}/ports/{ports_name}/peer/ | Update peer by ID
*K8sdispatcherApi* | [**UpdateK8sdispatcherSpanByID**](docs/K8sdispatcherApi.md#updatek8sdispatcherspanbyid) | **Patch** /k8sdispatcher/{name}/span/ | Update span by ID


## Documentation For Models

 - [K8sdispatcher](docs/K8sdispatcher.md)
 - [NodeportRule](docs/NodeportRule.md)
 - [Ports](docs/Ports.md)
 - [SessionRule](docs/SessionRule.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



