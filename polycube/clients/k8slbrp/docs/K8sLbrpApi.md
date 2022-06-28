# \K8sLbrpApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request                                                                 | Description
------------- |------------------------------------------------------------------------------| -------------
[**CreateK8sLbrpByID**](K8sK8sLbrpApi.md#CreateK8sLbrpByID) | **Post** /k8slbrp/{name}/                                                    | Create k8slbrp by ID
[**CreateK8sLbrpPortsByID**](K8sK8sLbrpApi.md#CreateK8sLbrpPortsByID) | **Post** /k8slbrp/{name}/ports/{ports_name}/                                 | Create ports by ID
[**CreateK8sLbrpPortsListByID**](K8sK8sLbrpApi.md#CreateK8sLbrpPortsListByID) | **Post** /k8slbrp/{name}/ports/                                              | Create ports by ID
[**CreateK8sLbrpServiceBackendByID**](K8sK8sLbrpApi.md#CreateK8sLbrpServiceBackendByID) | **Post** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/         | Create backend by ID
[**CreateK8sLbrpServiceBackendListByID**](K8sK8sLbrpApi.md#CreateK8sLbrpServiceBackendListByID) | **Post** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/              | Create backend by ID
[**CreateK8sLbrpServiceByID**](K8sK8sLbrpApi.md#CreateK8sLbrpServiceByID) | **Post** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/                      | Create service by ID
[**CreateK8sLbrpServiceListByID**](K8sK8sLbrpApi.md#CreateK8sLbrpServiceListByID) | **Post** /k8slbrp/{name}/service/                                            | Create service by ID
[**CreateK8sLbrpSrcIpRewriteByID**](K8sK8sLbrpApi.md#CreateK8sLbrpSrcIpRewriteByID) | **Post** /k8slbrp/{name}/src-ip-rewrite/                                     | Create src-ip-rewrite by ID
[**DeleteK8sLbrpByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpByID) | **Delete** /k8slbrp/{name}/                                                  | Delete k8slbrp by ID
[**DeleteK8sLbrpPortsByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpPortsByID) | **Delete** /k8slbrp/{name}/ports/{ports_name}/                               | Delete ports by ID
[**DeleteK8sLbrpPortsListByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpPortsListByID) | **Delete** /k8slbrp/{name}/ports/                                            | Delete ports by ID
[**DeleteK8sLbrpServiceBackendByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpServiceBackendByID) | **Delete** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/       | Delete backend by ID
[**DeleteK8sLbrpServiceBackendListByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpServiceBackendListByID) | **Delete** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/            | Delete backend by ID
[**DeleteK8sLbrpServiceByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpServiceByID) | **Delete** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/                    | Delete service by ID
[**DeleteK8sLbrpServiceListByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpServiceListByID) | **Delete** /k8slbrp/{name}/service/                                          | Delete service by ID
[**DeleteK8sLbrpSrcIpRewriteByID**](K8sK8sLbrpApi.md#DeleteK8sLbrpSrcIpRewriteByID) | **Delete** /k8slbrp/{name}/src-ip-rewrite/                                   | Delete src-ip-rewrite by ID
[**ReadK8sLbrpByID**](K8sK8sLbrpApi.md#ReadK8sLbrpByID) | **Get** /k8slbrp/{name}/                                                     | Read k8slbrp by ID
[**ReadK8sLbrpListByID**](K8sK8sLbrpApi.md#ReadK8sLbrpListByID) | **Get** /k8slbrp/                                                            | Read k8slbrp by ID
[**ReadK8sLbrpLoglevelByID**](K8sK8sLbrpApi.md#ReadK8sLbrpLoglevelByID) | **Get** /k8slbrp/{name}/loglevel/                                            | Read loglevel by ID
[**ReadK8sLbrpPortsByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortsByID) | **Get** /k8slbrp/{name}/ports/{ports_name}/                                  | Read ports by ID
[**ReadK8sLbrpPortsListByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortsListByID) | **Get** /k8slbrp/{name}/ports/                                               | Read ports by ID
[**ReadK8sLbrpPortsPeerByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortsPeerByID) | **Get** /k8slbrp/{name}/ports/{ports_name}/peer/                             | Read peer by ID
[**ReadK8sLbrpPortsStatusByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortsStatusByID) | **Get** /k8slbrp/{name}/ports/{ports_name}/status/                           | Read status by ID
[**ReadK8sLbrpPortsTypeByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortsTypeByID) | **Get** /k8slbrp/{name}/ports/{ports_name}/type/                             | Read type by ID
[**ReadK8sLbrpPortsIpByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortsIpByID) | **Get** /k8slbrp/{name}/ports/{ports_name}/ip/                               | Read ip by ID
[**ReadK8sLbrpPortsUuidByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortsUuidByID) | **Get** /k8slbrp/{name}/ports/{ports_name}/uuid/                             | Read uuid by ID
[**ReadK8sLbrpServiceBackendByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceBackendByID) | **Get** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/          | Read backend by ID
[**ReadK8sLbrpServiceBackendListByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceBackendListByID) | **Get** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/               | Read backend by ID
[**ReadK8sLbrpServiceBackendNameByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceBackendNameByID) | **Get** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/name/     | Read name by ID
[**ReadK8sLbrpServiceBackendPortByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceBackendPortByID) | **Get** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/port/     | Read port by ID
[**ReadK8sLbrpServiceBackendWeightByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceBackendWeightByID) | **Get** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/weight/   | Read weight by ID
[**ReadK8sLbrpServiceByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceByID) | **Get** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/                       | Read service by ID
[**ReadK8sLbrpServiceListByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceListByID) | **Get** /k8slbrp/{name}/service/                                             | Read service by ID
[**ReadK8sLbrpServiceNameByID**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceNameByID) | **Get** /k8slbrp/{name}/service-name/                                        | Read service-name by ID
[**ReadK8sLbrpServiceNameByID_0**](K8sK8sLbrpApi.md#ReadK8sLbrpServiceNameByID_0) | **Get** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/name/                  | Read name by ID
[**ReadK8sLbrpShadowByID**](K8sK8sLbrpApi.md#ReadK8sLbrpShadowByID) | **Get** /k8slbrp/{name}/shadow/                                              | Read shadow by ID
[**ReadK8sLbrpSpanByID**](K8sK8sLbrpApi.md#ReadK8sLbrpSpanByID) | **Get** /k8slbrp/{name}/span/                                                | Read span by ID
[**ReadK8sLbrpSrcIpRewriteByID**](K8sK8sLbrpApi.md#ReadK8sLbrpSrcIpRewriteByID) | **Get** /k8slbrp/{name}/src-ip-rewrite/                                      | Read src-ip-rewrite by ID
[**ReadK8sLbrpSrcIpRewriteIpRangeByID**](K8sK8sLbrpApi.md#ReadK8sLbrpSrcIpRewriteIpRangeByID) | **Get** /k8slbrp/{name}/src-ip-rewrite/ip-range/                             | Read ip-range by ID
[**ReadK8sLbrpSrcIpRewriteNewIpRangeByID**](K8sK8sLbrpApi.md#ReadK8sLbrpSrcIpRewriteNewIpRangeByID) | **Get** /k8slbrp/{name}/src-ip-rewrite/new_ip_range/                         | Read new_ip_range by ID
[**ReadK8sLbrpTypeByID**](K8sK8sLbrpApi.md#ReadK8sLbrpTypeByID) | **Get** /k8slbrp/{name}/type/                                                | Read type by ID
[**ReadK8sLbrpUuidByID**](K8sK8sLbrpApi.md#ReadK8sLbrpUuidByID) | **Get** /k8slbrp/{name}/uuid/                                                | Read uuid by ID
[**ReadK8sLbrpPortModeByID**](K8sK8sLbrpApi.md#ReadK8sLbrpPortModeByID) | **Get** /k8slbrp/{name}/port_mode/                                           | Read port mode by ID
[**ReplaceK8sLbrpByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpByID) | **Put** /k8slbrp/{name}/                                                     | Replace k8slbrp by ID
[**ReplaceK8sLbrpPortsByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpPortsByID) | **Put** /k8slbrp/{name}/ports/{ports_name}/                                  | Replace ports by ID
[**ReplaceK8sLbrpPortsListByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpPortsListByID) | **Put** /k8slbrp/{name}/ports/                                               | Replace ports by ID
[**ReplaceK8sLbrpServiceBackendByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpServiceBackendByID) | **Put** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/          | Replace backend by ID
[**ReplaceK8sLbrpServiceBackendListByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpServiceBackendListByID) | **Put** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/               | Replace backend by ID
[**ReplaceK8sLbrpServiceByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpServiceByID) | **Put** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/                       | Replace service by ID
[**ReplaceK8sLbrpServiceListByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpServiceListByID) | **Put** /k8slbrp/{name}/service/                                             | Replace service by ID
[**ReplaceK8sLbrpSrcIpRewriteByID**](K8sK8sLbrpApi.md#ReplaceK8sLbrpSrcIpRewriteByID) | **Put** /k8slbrp/{name}/src-ip-rewrite/                                      | Replace src-ip-rewrite by ID
[**UpdateK8sLbrpByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpByID) | **Patch** /k8slbrp/{name}/                                                   | Update k8slbrp by ID
[**UpdateK8sLbrpListByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpListByID) | **Patch** /k8slbrp/                                                          | Update k8slbrp by ID
[**UpdateK8sLbrpLoglevelByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpLoglevelByID) | **Patch** /k8slbrp/{name}/loglevel/                                          | Update loglevel by ID
[**UpdateK8sLbrpPortsByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpPortsByID) | **Patch** /k8slbrp/{name}/ports/{ports_name}/                                | Update ports by ID
[**UpdateK8sLbrpPortsListByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpPortsListByID) | **Patch** /k8slbrp/{name}/ports/                                             | Update ports by ID
[**UpdateK8sLbrpPortsPeerByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpPortsPeerByID) | **Patch** /k8slbrp/{name}/ports/{ports_name}/peer/                           | Update peer by ID
[**UpdateK8sLbrpPortsTypeByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpPortsTypeByID) | **Patch** /k8slbrp/{name}/ports/{ports_name}/type/                           | Update type by ID
[**UpdateK8sLbrpServiceBackendByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceBackendByID) | **Patch** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/        | Update backend by ID
[**UpdateK8sLbrpServiceBackendListByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceBackendListByID) | **Patch** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/             | Update backend by ID
[**UpdateK8sLbrpServiceBackendNameByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceBackendNameByID) | **Patch** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/name/   | Update name by ID
[**UpdateK8sLbrpServiceBackendPortByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceBackendPortByID) | **Patch** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/port/   | Update port by ID
[**UpdateK8sLbrpServiceBackendWeightByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceBackendWeightByID) | **Patch** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/weight/ | Update weight by ID
[**UpdateK8sLbrpServiceByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceByID) | **Patch** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/                     | Update service by ID
[**UpdateK8sLbrpServiceListByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceListByID) | **Patch** /k8slbrp/{name}/service/                                           | Update service by ID
[**UpdateK8sLbrpServiceNameByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpServiceNameByID) | **Patch** /k8slbrp/{name}/service/{vip}/{vport}/{proto}/name/                | Update name by ID
[**UpdateK8sLbrpSpanByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpSpanByID) | **Patch** /k8slbrp/{name}/span/                                              | Update span by ID
[**UpdateK8sLbrpSrcIpRewriteByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpSrcIpRewriteByID) | **Patch** /k8slbrp/{name}/src-ip-rewrite/                                    | Update src-ip-rewrite by ID
[**UpdateK8sLbrpSrcIpRewriteIpRangeByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpSrcIpRewriteIpRangeByID) | **Patch** /k8slbrp/{name}/src-ip-rewrite/ip-range/                           | Update ip-range by ID
[**UpdateK8sLbrpSrcIpRewriteNewIpRangeByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpSrcIpRewriteNewIpRangeByID) | **Patch** /k8slbrp/{name}/src-ip-rewrite/new_ip_range/                       | Update new_ip_range by ID
[**UpdateK8sLbrpPortModeByID**](K8sK8sLbrpApi.md#UpdateK8sLbrpPortModeByID) | **Patch** /k8slbrp/{name}/port_mode/                                         | Update port mode by ID

# **CreateK8sLbrpByID**
> CreateK8sLbrpByID(ctx, name, k8slbrp)
Create k8slbrp by ID

Create operation of resource: k8slbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **k8slbrp** | [**K8sLbrp**](K8sK8sLbrp.md)| k8slbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sLbrpPortsByID**
> CreateK8sLbrpPortsByID(ctx, name, portsName, ports)
Create ports by ID

Create operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ports** | [**Ports**](Ports.md)| portsbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sLbrpPortsListByID**
> CreateK8sLbrpPortsListByID(ctx, name, ports)
Create ports by ID

Create operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **ports** | [**[]Ports**](Ports.md)| portsbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sLbrpServiceBackendByID**
> CreateK8sLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip, backend)
Create backend by ID

Create operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 
  **backend** | [**ServiceBackend**](ServiceBackend.md)| backendbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sLbrpServiceBackendListByID**
> CreateK8sLbrpServiceBackendListByID(ctx, name, vip, vport, proto, backend)
Create backend by ID

Create operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **backend** | [**[]ServiceBackend**](ServiceBackend.md)| backendbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sLbrpServiceByID**
> CreateK8sLbrpServiceByID(ctx, name, vip, vport, proto, service)
Create service by ID

Create operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **service** | [**Service**](Service.md)| servicebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sLbrpServiceListByID**
> CreateK8sLbrpServiceListByID(ctx, name, service)
Create service by ID

Create operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **service** | [**[]Service**](Service.md)| servicebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sLbrpSrcIpRewriteByID**
> CreateK8sLbrpSrcIpRewriteByID(ctx, name, srcIpRewrite)
Create src-ip-rewrite by ID

Create operation of resource: src-ip-rewrite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **srcIpRewrite** | [**SrcIpRewrite**](SrcIpRewrite.md)| src-ip-rewritebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpByID**
> DeleteK8sLbrpByID(ctx, name)
Delete k8slbrp by ID

Delete operation of resource: k8slbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpPortsByID**
> DeleteK8sLbrpPortsByID(ctx, name, portsName)
Delete ports by ID

Delete operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpPortsListByID**
> DeleteK8sLbrpPortsListByID(ctx, name)
Delete ports by ID

Delete operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpServiceBackendByID**
> DeleteK8sLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip)
Delete backend by ID

Delete operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpServiceBackendListByID**
> DeleteK8sLbrpServiceBackendListByID(ctx, name, vip, vport, proto)
Delete backend by ID

Delete operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpServiceByID**
> DeleteK8sLbrpServiceByID(ctx, name, vip, vport, proto)
Delete service by ID

Delete operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpServiceListByID**
> DeleteK8sLbrpServiceListByID(ctx, name)
Delete service by ID

Delete operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sLbrpSrcIpRewriteByID**
> DeleteK8sLbrpSrcIpRewriteByID(ctx, name)
Delete src-ip-rewrite by ID

Delete operation of resource: src-ip-rewrite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpByID**
> K8sLbrp ReadK8sLbrpByID(ctx, name)
Read k8slbrp by ID

Read operation of resource: k8slbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**K8sLbrp**](K8sK8sLbrp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpListByID**
> []K8sLbrp ReadK8sLbrpListByID(ctx, )
Read k8slbrp by ID

Read operation of resource: k8slbrp

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]K8sLbrp**](K8sK8sLbrp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpLoglevelByID**
> string ReadK8sLbrpLoglevelByID(ctx, name)
Read loglevel by ID

Read operation of resource: loglevel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortsByID**
> Ports ReadK8sLbrpPortsByID(ctx, name, portsName)
Read ports by ID

Read operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 

### Return type

[**Ports**](Ports.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortsListByID**
> []Ports ReadK8sLbrpPortsListByID(ctx, name)
Read ports by ID

Read operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]Ports**](Ports.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortsPeerByID**
> string ReadK8sLbrpPortsPeerByID(ctx, name, portsName)
Read peer by ID

Read operation of resource: peer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortsStatusByID**
> string ReadK8sLbrpPortsStatusByID(ctx, name, portsName)
Read status by ID

Read operation of resource: status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortsTypeByID**
> string ReadK8sLbrpPortsTypeByID(ctx, name, portsName)
Read type by ID

Read operation of resource: type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortsIpByID**
> string ReadK8sLbrpPortsIpByID(ctx, name, portsName)
Read ip by ID

Read operation of resource: ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| ID of name |
**portsName** | **string**| ID of ports_name |

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortsUuidByID**
> string ReadK8sLbrpPortsUuidByID(ctx, name, portsName)
Read uuid by ID

Read operation of resource: uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceBackendByID**
> ServiceBackend ReadK8sLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip)
Read backend by ID

Read operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 

### Return type

[**ServiceBackend**](ServiceBackend.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceBackendListByID**
> []ServiceBackend ReadK8sLbrpServiceBackendListByID(ctx, name, vip, vport, proto)
Read backend by ID

Read operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 

### Return type

[**[]ServiceBackend**](ServiceBackend.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceBackendNameByID**
> string ReadK8sLbrpServiceBackendNameByID(ctx, name, vip, vport, proto, ip)
Read name by ID

Read operation of resource: name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceBackendPortByID**
> int32 ReadK8sLbrpServiceBackendPortByID(ctx, name, vip, vport, proto, ip)
Read port by ID

Read operation of resource: port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceBackendWeightByID**
> int32 ReadK8sLbrpServiceBackendWeightByID(ctx, name, vip, vport, proto, ip)
Read weight by ID

Read operation of resource: weight

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceByID**
> Service ReadK8sLbrpServiceByID(ctx, name, vip, vport, proto)
Read service by ID

Read operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceListByID**
> []Service ReadK8sLbrpServiceListByID(ctx, name)
Read service by ID

Read operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceNameByID**
> string ReadK8sLbrpServiceNameByID(ctx, name)
Read service-name by ID

Read operation of resource: service-name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpServiceNameByID_0**
> string ReadK8sLbrpServiceNameByID_0(ctx, name, vip, vport, proto)
Read name by ID

Read operation of resource: name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpShadowByID**
> bool ReadK8sLbrpShadowByID(ctx, name)
Read shadow by ID

Read operation of resource: shadow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpSpanByID**
> bool ReadK8sLbrpSpanByID(ctx, name)
Read span by ID

Read operation of resource: span

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpSrcIpRewriteByID**
> SrcIpRewrite ReadK8sLbrpSrcIpRewriteByID(ctx, name)
Read src-ip-rewrite by ID

Read operation of resource: src-ip-rewrite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**SrcIpRewrite**](SrcIpRewrite.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpSrcIpRewriteIpRangeByID**
> string ReadK8sLbrpSrcIpRewriteIpRangeByID(ctx, name)
Read ip-range by ID

Read operation of resource: ip-range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpSrcIpRewriteNewIpRangeByID**
> string ReadK8sLbrpSrcIpRewriteNewIpRangeByID(ctx, name)
Read new_ip_range by ID

Read operation of resource: new_ip_range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpTypeByID**
> string ReadK8sLbrpTypeByID(ctx, name)
Read type by ID

Read operation of resource: type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpUuidByID**
> string ReadK8sLbrpUuidByID(ctx, name)
Read uuid by ID

Read operation of resource: uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sLbrpPortModeByID**
> string ReadK8sLbrpPortModeByID(ctx, name)
Read port mode by ID

Read operation of resource: port_mode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| ID of name |

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpByID**
> ReplaceK8sLbrpByID(ctx, name, k8slbrp)
Replace k8slbrp by ID

Replace operation of resource: k8slbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **k8slbrp** | [**K8sLbrp**](K8sK8sLbrp.md)| k8slbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpPortsByID**
> ReplaceK8sLbrpPortsByID(ctx, name, portsName, ports)
Replace ports by ID

Replace operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ports** | [**Ports**](Ports.md)| portsbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpPortsListByID**
> ReplaceK8sLbrpPortsListByID(ctx, name, ports)
Replace ports by ID

Replace operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **ports** | [**[]Ports**](Ports.md)| portsbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpServiceBackendByID**
> ReplaceK8sLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip, backend)
Replace backend by ID

Replace operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 
  **backend** | [**ServiceBackend**](ServiceBackend.md)| backendbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpServiceBackendListByID**
> ReplaceK8sLbrpServiceBackendListByID(ctx, name, vip, vport, proto, backend)
Replace backend by ID

Replace operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **backend** | [**[]ServiceBackend**](ServiceBackend.md)| backendbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpServiceByID**
> ReplaceK8sLbrpServiceByID(ctx, name, vip, vport, proto, service)
Replace service by ID

Replace operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **service** | [**Service**](Service.md)| servicebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpServiceListByID**
> ReplaceK8sLbrpServiceListByID(ctx, name, service)
Replace service by ID

Replace operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **service** | [**[]Service**](Service.md)| servicebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sLbrpSrcIpRewriteByID**
> ReplaceK8sLbrpSrcIpRewriteByID(ctx, name, srcIpRewrite)
Replace src-ip-rewrite by ID

Replace operation of resource: src-ip-rewrite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **srcIpRewrite** | [**SrcIpRewrite**](SrcIpRewrite.md)| src-ip-rewritebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpByID**
> UpdateK8sLbrpByID(ctx, name, k8slbrp)
Update k8slbrp by ID

Update operation of resource: k8slbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **k8slbrp** | [**K8sLbrp**](K8sK8sLbrp.md)| k8slbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpListByID**
> UpdateK8sLbrpListByID(ctx, k8slbrp)
Update k8slbrp by ID

Update operation of resource: k8slbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8slbrp** | [**[]K8sLbrp**](K8sK8sLbrp.md)| k8slbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpLoglevelByID**
> UpdateK8sLbrpLoglevelByID(ctx, name, loglevel)
Update loglevel by ID

Update operation of resource: loglevel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **loglevel** | **string**| Defines the logging level of a service instance, from none (OFF) to the most verbose (TRACE) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpPortsByID**
> UpdateK8sLbrpPortsByID(ctx, name, portsName, ports)
Update ports by ID

Update operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ports** | [**Ports**](Ports.md)| portsbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpPortsListByID**
> UpdateK8sLbrpPortsListByID(ctx, name, ports)
Update ports by ID

Update operation of resource: ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **ports** | [**[]Ports**](Ports.md)| portsbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpPortsPeerByID**
> UpdateK8sLbrpPortsPeerByID(ctx, name, portsName, peer)
Update peer by ID

Update operation of resource: peer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **peer** | **string**| Peer name, such as a network interfaces (e.g., &#39;veth0&#39;) or another cube (e.g., &#39;br1:port2&#39;) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpPortsTypeByID**
> UpdateK8sLbrpPortsTypeByID(ctx, name, portsName, type_)
Update type by ID

Update operation of resource: type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **type_** | **string**| Type of the LB port (e.g. FRONTEND or BACKEND) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpPortsIpByID**
> UpdateK8sLbrpPortsIpByID(ctx, name, portsName, ip_)
Update ip by ID

Update operation of resource: ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| ID of name |
**portsName** | **string**| ID of ports_name |
**ip_** | **string**| IP address of the client interface (only for FRONTEND port) |

### Return type

(empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# **UpdateK8sLbrpServiceBackendByID**
> UpdateK8sLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip, backend)
Update backend by ID

Update operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 
  **backend** | [**ServiceBackend**](ServiceBackend.md)| backendbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpServiceBackendListByID**
> UpdateK8sLbrpServiceBackendListByID(ctx, name, vip, vport, proto, backend)
Update backend by ID

Update operation of resource: backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **backend** | [**[]ServiceBackend**](ServiceBackend.md)| backendbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpServiceBackendNameByID**
> UpdateK8sLbrpServiceBackendNameByID(ctx, name, vip, vport, proto, ip, name2)
Update name by ID

Update operation of resource: name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 
  **name2** | **string**| name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpServiceBackendPortByID**
> UpdateK8sLbrpServiceBackendPortByID(ctx, name, vip, vport, proto, ip, port)
Update port by ID

Update operation of resource: port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 
  **port** | **int32**| Port where the server listen to (this value is ignored in case of ICMP) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpServiceBackendWeightByID**
> UpdateK8sLbrpServiceBackendWeightByID(ctx, name, vip, vport, proto, ip, weight)
Update weight by ID

Update operation of resource: weight

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **ip** | **string**| ID of ip | 
  **weight** | **int32**| Weight of the backend in the pool | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpServiceByID**
> UpdateK8sLbrpServiceByID(ctx, name, vip, vport, proto, service)
Update service by ID

Update operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **service** | [**Service**](Service.md)| servicebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpServiceListByID**
> UpdateK8sLbrpServiceListByID(ctx, name, service)
Update service by ID

Update operation of resource: service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **service** | [**[]Service**](Service.md)| servicebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpServiceNameByID**
> UpdateK8sLbrpServiceNameByID(ctx, name, vip, vport, proto, name2)
Update name by ID

Update operation of resource: name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **vip** | **string**| ID of vip | 
  **vport** | **int32**| ID of vport | 
  **proto** | **string**| ID of proto | 
  **name2** | **string**| Service name related to the backend server of the pool is connected to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpSpanByID**
> UpdateK8sLbrpSpanByID(ctx, name, span)
Update span by ID

Update operation of resource: span

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **span** | **bool**| Defines if all traffic is sent to Linux | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpSrcIpRewriteByID**
> UpdateK8sLbrpSrcIpRewriteByID(ctx, name, srcIpRewrite)
Update src-ip-rewrite by ID

Update operation of resource: src-ip-rewrite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **srcIpRewrite** | [**SrcIpRewrite**](SrcIpRewrite.md)| src-ip-rewritebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpSrcIpRewriteIpRangeByID**
> UpdateK8sLbrpSrcIpRewriteIpRangeByID(ctx, name, ipRange)
Update ip-range by ID

Update operation of resource: ip-range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **ipRange** | **string**| Range of IP addresses of the clients that must be replaced | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpSrcIpRewriteNewIpRangeByID**
> UpdateK8sLbrpSrcIpRewriteNewIpRangeByID(ctx, name, newIpRange)
Update new_ip_range by ID

Update operation of resource: new_ip_range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **newIpRange** | **string**| Range of IP addresses of the that must be used to replace client addresses | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sLbrpPortModeByID**
> UpdateK8sLbrpPortModeByID(ctx, name, portMode)
Update port mode by ID

Update operation of resource: port_mode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| ID of name |
**portMode** | **string**| K8s lbrp mode of operation. 'MULTI' allows to manage multiple FRONTEND port. 'SINGLE' is optimized for working with a single FRONTEND port |

### Return type

(empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

