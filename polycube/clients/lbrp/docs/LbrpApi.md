# \LbrpApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request                                                                 | Description
------------- |------------------------------------------------------------------------------| -------------
[**CreateLbrpByID**](LbrpApi.md#CreateLbrpByID) | **Post** /lbrp/{name}/                                                    | Create lbrp by ID
[**CreateLbrpPortsByID**](LbrpApi.md#CreateLbrpPortsByID) | **Post** /lbrp/{name}/ports/{ports_name}/                                 | Create ports by ID
[**CreateLbrpPortsListByID**](LbrpApi.md#CreateLbrpPortsListByID) | **Post** /lbrp/{name}/ports/                                              | Create ports by ID
[**CreateLbrpServiceBackendByID**](LbrpApi.md#CreateLbrpServiceBackendByID) | **Post** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/         | Create backend by ID
[**CreateLbrpServiceBackendListByID**](LbrpApi.md#CreateLbrpServiceBackendListByID) | **Post** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/              | Create backend by ID
[**CreateLbrpServiceByID**](LbrpApi.md#CreateLbrpServiceByID) | **Post** /lbrp/{name}/service/{vip}/{vport}/{proto}/                      | Create service by ID
[**CreateLbrpServiceListByID**](LbrpApi.md#CreateLbrpServiceListByID) | **Post** /lbrp/{name}/service/                                            | Create service by ID
[**CreateLbrpSrcIpRewriteByID**](LbrpApi.md#CreateLbrpSrcIpRewriteByID) | **Post** /lbrp/{name}/src-ip-rewrite/                                     | Create src-ip-rewrite by ID
[**DeleteLbrpByID**](LbrpApi.md#DeleteLbrpByID) | **Delete** /lbrp/{name}/                                                  | Delete lbrp by ID
[**DeleteLbrpPortsByID**](LbrpApi.md#DeleteLbrpPortsByID) | **Delete** /lbrp/{name}/ports/{ports_name}/                               | Delete ports by ID
[**DeleteLbrpPortsListByID**](LbrpApi.md#DeleteLbrpPortsListByID) | **Delete** /lbrp/{name}/ports/                                            | Delete ports by ID
[**DeleteLbrpServiceBackendByID**](LbrpApi.md#DeleteLbrpServiceBackendByID) | **Delete** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/       | Delete backend by ID
[**DeleteLbrpServiceBackendListByID**](LbrpApi.md#DeleteLbrpServiceBackendListByID) | **Delete** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/            | Delete backend by ID
[**DeleteLbrpServiceByID**](LbrpApi.md#DeleteLbrpServiceByID) | **Delete** /lbrp/{name}/service/{vip}/{vport}/{proto}/                    | Delete service by ID
[**DeleteLbrpServiceListByID**](LbrpApi.md#DeleteLbrpServiceListByID) | **Delete** /lbrp/{name}/service/                                          | Delete service by ID
[**DeleteLbrpSrcIpRewriteByID**](LbrpApi.md#DeleteLbrpSrcIpRewriteByID) | **Delete** /lbrp/{name}/src-ip-rewrite/                                   | Delete src-ip-rewrite by ID
[**ReadLbrpByID**](LbrpApi.md#ReadLbrpByID) | **Get** /lbrp/{name}/                                                     | Read lbrp by ID
[**ReadLbrpListByID**](LbrpApi.md#ReadLbrpListByID) | **Get** /lbrp/                                                            | Read lbrp by ID
[**ReadLbrpLoglevelByID**](LbrpApi.md#ReadLbrpLoglevelByID) | **Get** /lbrp/{name}/loglevel/                                            | Read loglevel by ID
[**ReadLbrpPortsByID**](LbrpApi.md#ReadLbrpPortsByID) | **Get** /lbrp/{name}/ports/{ports_name}/                                  | Read ports by ID
[**ReadLbrpPortsListByID**](LbrpApi.md#ReadLbrpPortsListByID) | **Get** /lbrp/{name}/ports/                                               | Read ports by ID
[**ReadLbrpPortsPeerByID**](LbrpApi.md#ReadLbrpPortsPeerByID) | **Get** /lbrp/{name}/ports/{ports_name}/peer/                             | Read peer by ID
[**ReadLbrpPortsStatusByID**](LbrpApi.md#ReadLbrpPortsStatusByID) | **Get** /lbrp/{name}/ports/{ports_name}/status/                           | Read status by ID
[**ReadLbrpPortsTypeByID**](LbrpApi.md#ReadLbrpPortsTypeByID) | **Get** /lbrp/{name}/ports/{ports_name}/type/                             | Read type by ID
[**ReadLbrpPortsIpByID**](LbrpApi.md#ReadLbrpPortsIpByID) | **Get** /lbrp/{name}/ports/{ports_name}/ip/                               | Read ip by ID
[**ReadLbrpPortsUuidByID**](LbrpApi.md#ReadLbrpPortsUuidByID) | **Get** /lbrp/{name}/ports/{ports_name}/uuid/                             | Read uuid by ID
[**ReadLbrpServiceBackendByID**](LbrpApi.md#ReadLbrpServiceBackendByID) | **Get** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/          | Read backend by ID
[**ReadLbrpServiceBackendListByID**](LbrpApi.md#ReadLbrpServiceBackendListByID) | **Get** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/               | Read backend by ID
[**ReadLbrpServiceBackendNameByID**](LbrpApi.md#ReadLbrpServiceBackendNameByID) | **Get** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/name/     | Read name by ID
[**ReadLbrpServiceBackendPortByID**](LbrpApi.md#ReadLbrpServiceBackendPortByID) | **Get** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/port/     | Read port by ID
[**ReadLbrpServiceBackendWeightByID**](LbrpApi.md#ReadLbrpServiceBackendWeightByID) | **Get** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/weight/   | Read weight by ID
[**ReadLbrpServiceByID**](LbrpApi.md#ReadLbrpServiceByID) | **Get** /lbrp/{name}/service/{vip}/{vport}/{proto}/                       | Read service by ID
[**ReadLbrpServiceListByID**](LbrpApi.md#ReadLbrpServiceListByID) | **Get** /lbrp/{name}/service/                                             | Read service by ID
[**ReadLbrpServiceNameByID**](LbrpApi.md#ReadLbrpServiceNameByID) | **Get** /lbrp/{name}/service-name/                                        | Read service-name by ID
[**ReadLbrpServiceNameByID_0**](LbrpApi.md#ReadLbrpServiceNameByID_0) | **Get** /lbrp/{name}/service/{vip}/{vport}/{proto}/name/                  | Read name by ID
[**ReadLbrpShadowByID**](LbrpApi.md#ReadLbrpShadowByID) | **Get** /lbrp/{name}/shadow/                                              | Read shadow by ID
[**ReadLbrpSpanByID**](LbrpApi.md#ReadLbrpSpanByID) | **Get** /lbrp/{name}/span/                                                | Read span by ID
[**ReadLbrpSrcIpRewriteByID**](LbrpApi.md#ReadLbrpSrcIpRewriteByID) | **Get** /lbrp/{name}/src-ip-rewrite/                                      | Read src-ip-rewrite by ID
[**ReadLbrpSrcIpRewriteIpRangeByID**](LbrpApi.md#ReadLbrpSrcIpRewriteIpRangeByID) | **Get** /lbrp/{name}/src-ip-rewrite/ip-range/                             | Read ip-range by ID
[**ReadLbrpSrcIpRewriteNewIpRangeByID**](LbrpApi.md#ReadLbrpSrcIpRewriteNewIpRangeByID) | **Get** /lbrp/{name}/src-ip-rewrite/new_ip_range/                         | Read new_ip_range by ID
[**ReadLbrpTypeByID**](LbrpApi.md#ReadLbrpTypeByID) | **Get** /lbrp/{name}/type/                                                | Read type by ID
[**ReadLbrpUuidByID**](LbrpApi.md#ReadLbrpUuidByID) | **Get** /lbrp/{name}/uuid/                                                | Read uuid by ID
[**ReadLbrpPortModeByID**](LbrpApi.md#ReadLbrpPortModeByID) | **Get** /lbrp/{name}/port_mode/                                           | Read port mode by ID
[**ReplaceLbrpByID**](LbrpApi.md#ReplaceLbrpByID) | **Put** /lbrp/{name}/                                                     | Replace lbrp by ID
[**ReplaceLbrpPortsByID**](LbrpApi.md#ReplaceLbrpPortsByID) | **Put** /lbrp/{name}/ports/{ports_name}/                                  | Replace ports by ID
[**ReplaceLbrpPortsListByID**](LbrpApi.md#ReplaceLbrpPortsListByID) | **Put** /lbrp/{name}/ports/                                               | Replace ports by ID
[**ReplaceLbrpServiceBackendByID**](LbrpApi.md#ReplaceLbrpServiceBackendByID) | **Put** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/          | Replace backend by ID
[**ReplaceLbrpServiceBackendListByID**](LbrpApi.md#ReplaceLbrpServiceBackendListByID) | **Put** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/               | Replace backend by ID
[**ReplaceLbrpServiceByID**](LbrpApi.md#ReplaceLbrpServiceByID) | **Put** /lbrp/{name}/service/{vip}/{vport}/{proto}/                       | Replace service by ID
[**ReplaceLbrpServiceListByID**](LbrpApi.md#ReplaceLbrpServiceListByID) | **Put** /lbrp/{name}/service/                                             | Replace service by ID
[**ReplaceLbrpSrcIpRewriteByID**](LbrpApi.md#ReplaceLbrpSrcIpRewriteByID) | **Put** /lbrp/{name}/src-ip-rewrite/                                      | Replace src-ip-rewrite by ID
[**UpdateLbrpByID**](LbrpApi.md#UpdateLbrpByID) | **Patch** /lbrp/{name}/                                                   | Update lbrp by ID
[**UpdateLbrpListByID**](LbrpApi.md#UpdateLbrpListByID) | **Patch** /lbrp/                                                          | Update lbrp by ID
[**UpdateLbrpLoglevelByID**](LbrpApi.md#UpdateLbrpLoglevelByID) | **Patch** /lbrp/{name}/loglevel/                                          | Update loglevel by ID
[**UpdateLbrpPortsByID**](LbrpApi.md#UpdateLbrpPortsByID) | **Patch** /lbrp/{name}/ports/{ports_name}/                                | Update ports by ID
[**UpdateLbrpPortsListByID**](LbrpApi.md#UpdateLbrpPortsListByID) | **Patch** /lbrp/{name}/ports/                                             | Update ports by ID
[**UpdateLbrpPortsPeerByID**](LbrpApi.md#UpdateLbrpPortsPeerByID) | **Patch** /lbrp/{name}/ports/{ports_name}/peer/                           | Update peer by ID
[**UpdateLbrpPortsTypeByID**](LbrpApi.md#UpdateLbrpPortsTypeByID) | **Patch** /lbrp/{name}/ports/{ports_name}/type/                           | Update type by ID
[**UpdateLbrpServiceBackendByID**](LbrpApi.md#UpdateLbrpServiceBackendByID) | **Patch** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/        | Update backend by ID
[**UpdateLbrpServiceBackendListByID**](LbrpApi.md#UpdateLbrpServiceBackendListByID) | **Patch** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/             | Update backend by ID
[**UpdateLbrpServiceBackendNameByID**](LbrpApi.md#UpdateLbrpServiceBackendNameByID) | **Patch** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/name/   | Update name by ID
[**UpdateLbrpServiceBackendPortByID**](LbrpApi.md#UpdateLbrpServiceBackendPortByID) | **Patch** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/port/   | Update port by ID
[**UpdateLbrpServiceBackendWeightByID**](LbrpApi.md#UpdateLbrpServiceBackendWeightByID) | **Patch** /lbrp/{name}/service/{vip}/{vport}/{proto}/backend/{ip}/weight/ | Update weight by ID
[**UpdateLbrpServiceByID**](LbrpApi.md#UpdateLbrpServiceByID) | **Patch** /lbrp/{name}/service/{vip}/{vport}/{proto}/                     | Update service by ID
[**UpdateLbrpServiceListByID**](LbrpApi.md#UpdateLbrpServiceListByID) | **Patch** /lbrp/{name}/service/                                           | Update service by ID
[**UpdateLbrpServiceNameByID**](LbrpApi.md#UpdateLbrpServiceNameByID) | **Patch** /lbrp/{name}/service/{vip}/{vport}/{proto}/name/                | Update name by ID
[**UpdateLbrpSpanByID**](LbrpApi.md#UpdateLbrpSpanByID) | **Patch** /lbrp/{name}/span/                                              | Update span by ID
[**UpdateLbrpSrcIpRewriteByID**](LbrpApi.md#UpdateLbrpSrcIpRewriteByID) | **Patch** /lbrp/{name}/src-ip-rewrite/                                    | Update src-ip-rewrite by ID
[**UpdateLbrpSrcIpRewriteIpRangeByID**](LbrpApi.md#UpdateLbrpSrcIpRewriteIpRangeByID) | **Patch** /lbrp/{name}/src-ip-rewrite/ip-range/                           | Update ip-range by ID
[**UpdateLbrpSrcIpRewriteNewIpRangeByID**](LbrpApi.md#UpdateLbrpSrcIpRewriteNewIpRangeByID) | **Patch** /lbrp/{name}/src-ip-rewrite/new_ip_range/                       | Update new_ip_range by ID
[**UpdateLbrpPortModeByID**](LbrpApi.md#UpdateLbrpPortModeByID) | **Patch** /lbrp/{name}/port_mode/                                         | Update port mode by ID

# **CreateLbrpByID**
> CreateLbrpByID(ctx, name, lbrp)
Create lbrp by ID

Create operation of resource: lbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **lbrp** | [**Lbrp**](Lbrp.md)| lbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLbrpPortsByID**
> CreateLbrpPortsByID(ctx, name, portsName, ports)
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

# **CreateLbrpPortsListByID**
> CreateLbrpPortsListByID(ctx, name, ports)
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

# **CreateLbrpServiceBackendByID**
> CreateLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip, backend)
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

# **CreateLbrpServiceBackendListByID**
> CreateLbrpServiceBackendListByID(ctx, name, vip, vport, proto, backend)
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

# **CreateLbrpServiceByID**
> CreateLbrpServiceByID(ctx, name, vip, vport, proto, service)
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

# **CreateLbrpServiceListByID**
> CreateLbrpServiceListByID(ctx, name, service)
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

# **CreateLbrpSrcIpRewriteByID**
> CreateLbrpSrcIpRewriteByID(ctx, name, srcIpRewrite)
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

# **DeleteLbrpByID**
> DeleteLbrpByID(ctx, name)
Delete lbrp by ID

Delete operation of resource: lbrp

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

# **DeleteLbrpPortsByID**
> DeleteLbrpPortsByID(ctx, name, portsName)
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

# **DeleteLbrpPortsListByID**
> DeleteLbrpPortsListByID(ctx, name)
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

# **DeleteLbrpServiceBackendByID**
> DeleteLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip)
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

# **DeleteLbrpServiceBackendListByID**
> DeleteLbrpServiceBackendListByID(ctx, name, vip, vport, proto)
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

# **DeleteLbrpServiceByID**
> DeleteLbrpServiceByID(ctx, name, vip, vport, proto)
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

# **DeleteLbrpServiceListByID**
> DeleteLbrpServiceListByID(ctx, name)
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

# **DeleteLbrpSrcIpRewriteByID**
> DeleteLbrpSrcIpRewriteByID(ctx, name)
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

# **ReadLbrpByID**
> Lbrp ReadLbrpByID(ctx, name)
Read lbrp by ID

Read operation of resource: lbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**Lbrp**](Lbrp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLbrpListByID**
> []Lbrp ReadLbrpListByID(ctx, )
Read lbrp by ID

Read operation of resource: lbrp

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Lbrp**](Lbrp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLbrpLoglevelByID**
> string ReadLbrpLoglevelByID(ctx, name)
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

# **ReadLbrpPortsByID**
> Ports ReadLbrpPortsByID(ctx, name, portsName)
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

# **ReadLbrpPortsListByID**
> []Ports ReadLbrpPortsListByID(ctx, name)
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

# **ReadLbrpPortsPeerByID**
> string ReadLbrpPortsPeerByID(ctx, name, portsName)
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

# **ReadLbrpPortsStatusByID**
> string ReadLbrpPortsStatusByID(ctx, name, portsName)
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

# **ReadLbrpPortsTypeByID**
> string ReadLbrpPortsTypeByID(ctx, name, portsName)
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

# **ReadLbrpPortsIpByID**
> string ReadLbrpPortsIpByID(ctx, name, portsName)
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

# **ReadLbrpPortsUuidByID**
> string ReadLbrpPortsUuidByID(ctx, name, portsName)
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

# **ReadLbrpServiceBackendByID**
> ServiceBackend ReadLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip)
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

# **ReadLbrpServiceBackendListByID**
> []ServiceBackend ReadLbrpServiceBackendListByID(ctx, name, vip, vport, proto)
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

# **ReadLbrpServiceBackendNameByID**
> string ReadLbrpServiceBackendNameByID(ctx, name, vip, vport, proto, ip)
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

# **ReadLbrpServiceBackendPortByID**
> int32 ReadLbrpServiceBackendPortByID(ctx, name, vip, vport, proto, ip)
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

# **ReadLbrpServiceBackendWeightByID**
> int32 ReadLbrpServiceBackendWeightByID(ctx, name, vip, vport, proto, ip)
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

# **ReadLbrpServiceByID**
> Service ReadLbrpServiceByID(ctx, name, vip, vport, proto)
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

# **ReadLbrpServiceListByID**
> []Service ReadLbrpServiceListByID(ctx, name)
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

# **ReadLbrpServiceNameByID**
> string ReadLbrpServiceNameByID(ctx, name)
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

# **ReadLbrpServiceNameByID_0**
> string ReadLbrpServiceNameByID_0(ctx, name, vip, vport, proto)
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

# **ReadLbrpShadowByID**
> bool ReadLbrpShadowByID(ctx, name)
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

# **ReadLbrpSpanByID**
> bool ReadLbrpSpanByID(ctx, name)
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

# **ReadLbrpSrcIpRewriteByID**
> SrcIpRewrite ReadLbrpSrcIpRewriteByID(ctx, name)
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

# **ReadLbrpSrcIpRewriteIpRangeByID**
> string ReadLbrpSrcIpRewriteIpRangeByID(ctx, name)
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

# **ReadLbrpSrcIpRewriteNewIpRangeByID**
> string ReadLbrpSrcIpRewriteNewIpRangeByID(ctx, name)
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

# **ReadLbrpTypeByID**
> string ReadLbrpTypeByID(ctx, name)
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

# **ReadLbrpUuidByID**
> string ReadLbrpUuidByID(ctx, name)
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

# **ReadLbrpPortModeByID**
> string ReadLbrpPortModeByID(ctx, name)
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

# **ReplaceLbrpByID**
> ReplaceLbrpByID(ctx, name, lbrp)
Replace lbrp by ID

Replace operation of resource: lbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **lbrp** | [**Lbrp**](Lbrp.md)| lbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceLbrpPortsByID**
> ReplaceLbrpPortsByID(ctx, name, portsName, ports)
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

# **ReplaceLbrpPortsListByID**
> ReplaceLbrpPortsListByID(ctx, name, ports)
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

# **ReplaceLbrpServiceBackendByID**
> ReplaceLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip, backend)
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

# **ReplaceLbrpServiceBackendListByID**
> ReplaceLbrpServiceBackendListByID(ctx, name, vip, vport, proto, backend)
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

# **ReplaceLbrpServiceByID**
> ReplaceLbrpServiceByID(ctx, name, vip, vport, proto, service)
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

# **ReplaceLbrpServiceListByID**
> ReplaceLbrpServiceListByID(ctx, name, service)
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

# **ReplaceLbrpSrcIpRewriteByID**
> ReplaceLbrpSrcIpRewriteByID(ctx, name, srcIpRewrite)
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

# **UpdateLbrpByID**
> UpdateLbrpByID(ctx, name, lbrp)
Update lbrp by ID

Update operation of resource: lbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **lbrp** | [**Lbrp**](Lbrp.md)| lbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLbrpListByID**
> UpdateLbrpListByID(ctx, lbrp)
Update lbrp by ID

Update operation of resource: lbrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lbrp** | [**[]Lbrp**](Lbrp.md)| lbrpbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLbrpLoglevelByID**
> UpdateLbrpLoglevelByID(ctx, name, loglevel)
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

# **UpdateLbrpPortsByID**
> UpdateLbrpPortsByID(ctx, name, portsName, ports)
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

# **UpdateLbrpPortsListByID**
> UpdateLbrpPortsListByID(ctx, name, ports)
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

# **UpdateLbrpPortsPeerByID**
> UpdateLbrpPortsPeerByID(ctx, name, portsName, peer)
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

# **UpdateLbrpPortsTypeByID**
> UpdateLbrpPortsTypeByID(ctx, name, portsName, type_)
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

# **UpdateLbrpPortsIpByID**
> UpdateLbrpPortsIpByID(ctx, name, portsName, ip_)
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


# **UpdateLbrpServiceBackendByID**
> UpdateLbrpServiceBackendByID(ctx, name, vip, vport, proto, ip, backend)
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

# **UpdateLbrpServiceBackendListByID**
> UpdateLbrpServiceBackendListByID(ctx, name, vip, vport, proto, backend)
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

# **UpdateLbrpServiceBackendNameByID**
> UpdateLbrpServiceBackendNameByID(ctx, name, vip, vport, proto, ip, name2)
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

# **UpdateLbrpServiceBackendPortByID**
> UpdateLbrpServiceBackendPortByID(ctx, name, vip, vport, proto, ip, port)
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

# **UpdateLbrpServiceBackendWeightByID**
> UpdateLbrpServiceBackendWeightByID(ctx, name, vip, vport, proto, ip, weight)
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

# **UpdateLbrpServiceByID**
> UpdateLbrpServiceByID(ctx, name, vip, vport, proto, service)
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

# **UpdateLbrpServiceListByID**
> UpdateLbrpServiceListByID(ctx, name, service)
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

# **UpdateLbrpServiceNameByID**
> UpdateLbrpServiceNameByID(ctx, name, vip, vport, proto, name2)
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

# **UpdateLbrpSpanByID**
> UpdateLbrpSpanByID(ctx, name, span)
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

# **UpdateLbrpSrcIpRewriteByID**
> UpdateLbrpSrcIpRewriteByID(ctx, name, srcIpRewrite)
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

# **UpdateLbrpSrcIpRewriteIpRangeByID**
> UpdateLbrpSrcIpRewriteIpRangeByID(ctx, name, ipRange)
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

# **UpdateLbrpSrcIpRewriteNewIpRangeByID**
> UpdateLbrpSrcIpRewriteNewIpRangeByID(ctx, name, newIpRange)
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

# **UpdateLbrpPortModeByID**
> UpdateLbrpPortModeByID(ctx, name, portMode)
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

