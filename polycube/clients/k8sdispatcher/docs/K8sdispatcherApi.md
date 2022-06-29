# \K8sdispatcherApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateK8sdispatcherByID**](K8sdispatcherApi.md#CreateK8sdispatcherByID) | **Post** /k8sdispatcher/{name}/ | Create k8sdispatcher by ID
[**CreateK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#CreateK8sdispatcherNodeportRuleByID) | **Post** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Create nodeport-rule by ID
[**CreateK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#CreateK8sdispatcherNodeportRuleListByID) | **Post** /k8sdispatcher/{name}/nodeport-rule/ | Create nodeport-rule by ID
[**CreateK8sdispatcherPortsByID**](K8sdispatcherApi.md#CreateK8sdispatcherPortsByID) | **Post** /k8sdispatcher/{name}/ports/{ports_name}/ | Create ports by ID
[**CreateK8sdispatcherPortsListByID**](K8sdispatcherApi.md#CreateK8sdispatcherPortsListByID) | **Post** /k8sdispatcher/{name}/ports/ | Create ports by ID
[**DeleteK8sdispatcherByID**](K8sdispatcherApi.md#DeleteK8sdispatcherByID) | **Delete** /k8sdispatcher/{name}/ | Delete k8sdispatcher by ID
[**DeleteK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#DeleteK8sdispatcherNodeportRuleByID) | **Delete** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Delete nodeport-rule by ID
[**DeleteK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#DeleteK8sdispatcherNodeportRuleListByID) | **Delete** /k8sdispatcher/{name}/nodeport-rule/ | Delete nodeport-rule by ID
[**DeleteK8sdispatcherPortsByID**](K8sdispatcherApi.md#DeleteK8sdispatcherPortsByID) | **Delete** /k8sdispatcher/{name}/ports/{ports_name}/ | Delete ports by ID
[**DeleteK8sdispatcherPortsListByID**](K8sdispatcherApi.md#DeleteK8sdispatcherPortsListByID) | **Delete** /k8sdispatcher/{name}/ports/ | Delete ports by ID
[**ReadK8sdispatcherByID**](K8sdispatcherApi.md#ReadK8sdispatcherByID) | **Get** /k8sdispatcher/{name}/ | Read k8sdispatcher by ID
[**ReadK8sdispatcherInternalSrcIpByID**](K8sdispatcherApi.md#ReadK8sdispatcherInternalSrcIpByID) | **Get** /k8sdispatcher/{name}/internal-src-ip/ | Read internal-src-ip by ID
[**ReadK8sdispatcherListByID**](K8sdispatcherApi.md#ReadK8sdispatcherListByID) | **Get** /k8sdispatcher/ | Read k8sdispatcher by ID
[**ReadK8sdispatcherLoglevelByID**](K8sdispatcherApi.md#ReadK8sdispatcherLoglevelByID) | **Get** /k8sdispatcher/{name}/loglevel/ | Read loglevel by ID
[**ReadK8sdispatcherNodeportRangeByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRangeByID) | **Get** /k8sdispatcher/{name}/nodeport-range/ | Read nodeport-range by ID
[**ReadK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Read nodeport-rule by ID
[**ReadK8sdispatcherNodeportRuleExternalTrafficPolicyByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleExternalTrafficPolicyByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/external-traffic-policy/ | Read external-traffic-policy by ID
[**ReadK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleListByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/ | Read nodeport-rule by ID
[**ReadK8sdispatcherNodeportRuleRuleNameByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleRuleNameByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/rule-name/ | Read rule-name by ID
[**ReadK8sdispatcherPortsByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/ | Read ports by ID
[**ReadK8sdispatcherPortsIpByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsIpByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/ip/ | Read ip by ID
[**ReadK8sdispatcherPortsListByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsListByID) | **Get** /k8sdispatcher/{name}/ports/ | Read ports by ID
[**ReadK8sdispatcherPortsPeerByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsPeerByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/peer/ | Read peer by ID
[**ReadK8sdispatcherPortsStatusByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsStatusByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/status/ | Read status by ID
[**ReadK8sdispatcherPortsTypeByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsTypeByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/type/ | Read type by ID
[**ReadK8sdispatcherPortsUuidByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsUuidByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/uuid/ | Read uuid by ID
[**ReadK8sdispatcherServiceNameByID**](K8sdispatcherApi.md#ReadK8sdispatcherServiceNameByID) | **Get** /k8sdispatcher/{name}/service-name/ | Read service-name by ID
[**ReadK8sdispatcherSessionRuleByID**](K8sdispatcherApi.md#ReadK8sdispatcherSessionRuleByID) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/ | Read session-rule by ID
[**ReadK8sdispatcherSessionRuleListByID**](K8sdispatcherApi.md#ReadK8sdispatcherSessionRuleListByID) | **Get** /k8sdispatcher/{name}/session-rule/ | Read session-rule by ID
[**ReadK8sdispatcherSessionRuleNewIpByID**](K8sdispatcherApi.md#ReadK8sdispatcherSessionRuleNewIpByID) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/new-ip/ | Read new-ip by ID
[**ReadK8sdispatcherSessionRuleNewPortByID**](K8sdispatcherApi.md#ReadK8sdispatcherSessionRuleNewPortByID) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/new-port/ | Read new-port by ID
[**ReadK8sdispatcherSessionRuleOperationByID**](K8sdispatcherApi.md#ReadK8sdispatcherSessionRuleOperationByID) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/operation/ | Read operation by ID
[**ReadK8sdispatcherSessionRuleOriginatingRuleByID**](K8sdispatcherApi.md#ReadK8sdispatcherSessionRuleOriginatingRuleByID) | **Get** /k8sdispatcher/{name}/session-rule/{direction}/{src-ip}/{dst-ip}/{src-port}/{dst-port}/{proto}/originating-rule/ | Read originating-rule by ID
[**ReadK8sdispatcherShadowByID**](K8sdispatcherApi.md#ReadK8sdispatcherShadowByID) | **Get** /k8sdispatcher/{name}/shadow/ | Read shadow by ID
[**ReadK8sdispatcherSpanByID**](K8sdispatcherApi.md#ReadK8sdispatcherSpanByID) | **Get** /k8sdispatcher/{name}/span/ | Read span by ID
[**ReadK8sdispatcherTypeByID**](K8sdispatcherApi.md#ReadK8sdispatcherTypeByID) | **Get** /k8sdispatcher/{name}/type/ | Read type by ID
[**ReadK8sdispatcherUuidByID**](K8sdispatcherApi.md#ReadK8sdispatcherUuidByID) | **Get** /k8sdispatcher/{name}/uuid/ | Read uuid by ID
[**ReplaceK8sdispatcherByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherByID) | **Put** /k8sdispatcher/{name}/ | Replace k8sdispatcher by ID
[**ReplaceK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherNodeportRuleByID) | **Put** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Replace nodeport-rule by ID
[**ReplaceK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherNodeportRuleListByID) | **Put** /k8sdispatcher/{name}/nodeport-rule/ | Replace nodeport-rule by ID
[**ReplaceK8sdispatcherPortsByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherPortsByID) | **Put** /k8sdispatcher/{name}/ports/{ports_name}/ | Replace ports by ID
[**ReplaceK8sdispatcherPortsListByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherPortsListByID) | **Put** /k8sdispatcher/{name}/ports/ | Replace ports by ID
[**UpdateK8sdispatcherByID**](K8sdispatcherApi.md#UpdateK8sdispatcherByID) | **Patch** /k8sdispatcher/{name}/ | Update k8sdispatcher by ID
[**UpdateK8sdispatcherListByID**](K8sdispatcherApi.md#UpdateK8sdispatcherListByID) | **Patch** /k8sdispatcher/ | Update k8sdispatcher by ID
[**UpdateK8sdispatcherLoglevelByID**](K8sdispatcherApi.md#UpdateK8sdispatcherLoglevelByID) | **Patch** /k8sdispatcher/{name}/loglevel/ | Update loglevel by ID
[**UpdateK8sdispatcherNodeportRangeByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRangeByID) | **Patch** /k8sdispatcher/{name}/nodeport-range/ | Update nodeport-range by ID
[**UpdateK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRuleByID) | **Patch** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Update nodeport-rule by ID
[**UpdateK8sdispatcherNodeportRuleExternalTrafficPolicyByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRuleExternalTrafficPolicyByID) | **Patch** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/external-traffic-policy/ | Update external-traffic-policy by ID
[**UpdateK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRuleListByID) | **Patch** /k8sdispatcher/{name}/nodeport-rule/ | Update nodeport-rule by ID
[**UpdateK8sdispatcherPortsByID**](K8sdispatcherApi.md#UpdateK8sdispatcherPortsByID) | **Patch** /k8sdispatcher/{name}/ports/{ports_name}/ | Update ports by ID
[**UpdateK8sdispatcherPortsListByID**](K8sdispatcherApi.md#UpdateK8sdispatcherPortsListByID) | **Patch** /k8sdispatcher/{name}/ports/ | Update ports by ID
[**UpdateK8sdispatcherPortsPeerByID**](K8sdispatcherApi.md#UpdateK8sdispatcherPortsPeerByID) | **Patch** /k8sdispatcher/{name}/ports/{ports_name}/peer/ | Update peer by ID
[**UpdateK8sdispatcherSpanByID**](K8sdispatcherApi.md#UpdateK8sdispatcherSpanByID) | **Patch** /k8sdispatcher/{name}/span/ | Update span by ID


# **CreateK8sdispatcherByID**
> CreateK8sdispatcherByID(ctx, name, k8sdispatcher)
Create k8sdispatcher by ID

Create operation of resource: k8sdispatcher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **k8sdispatcher** | [**K8sdispatcher**](K8sdispatcher.md)| k8sdispatcherbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sdispatcherNodeportRuleByID**
> CreateK8sdispatcherNodeportRuleByID(ctx, name, nodeportPort, proto, nodeportRule)
Create nodeport-rule by ID

Create operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 
  **nodeportRule** | [**NodeportRule**](NodeportRule.md)| nodeport-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sdispatcherNodeportRuleListByID**
> CreateK8sdispatcherNodeportRuleListByID(ctx, name, nodeportRule)
Create nodeport-rule by ID

Create operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportRule** | [**[]NodeportRule**](NodeportRule.md)| nodeport-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sdispatcherPortsByID**
> CreateK8sdispatcherPortsByID(ctx, name, portsName, ports)
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

# **CreateK8sdispatcherPortsListByID**
> CreateK8sdispatcherPortsListByID(ctx, name, ports)
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

# **DeleteK8sdispatcherByID**
> DeleteK8sdispatcherByID(ctx, name)
Delete k8sdispatcher by ID

Delete operation of resource: k8sdispatcher

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

# **DeleteK8sdispatcherNodeportRuleByID**
> DeleteK8sdispatcherNodeportRuleByID(ctx, name, nodeportPort, proto)
Delete nodeport-rule by ID

Delete operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sdispatcherNodeportRuleListByID**
> DeleteK8sdispatcherNodeportRuleListByID(ctx, name)
Delete nodeport-rule by ID

Delete operation of resource: nodeport-rule

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

# **DeleteK8sdispatcherPortsByID**
> DeleteK8sdispatcherPortsByID(ctx, name, portsName)
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

# **DeleteK8sdispatcherPortsListByID**
> DeleteK8sdispatcherPortsListByID(ctx, name)
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

# **ReadK8sdispatcherByID**
> K8sdispatcher ReadK8sdispatcherByID(ctx, name)
Read k8sdispatcher by ID

Read operation of resource: k8sdispatcher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**K8sdispatcher**](K8sdispatcher.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherInternalSrcIpByID**
> string ReadK8sdispatcherInternalSrcIpByID(ctx, name)
Read internal-src-ip by ID

Read operation of resource: internal-src-ip

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

# **ReadK8sdispatcherListByID**
> []K8sdispatcher ReadK8sdispatcherListByID(ctx, )
Read k8sdispatcher by ID

Read operation of resource: k8sdispatcher

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]K8sdispatcher**](K8sdispatcher.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherLoglevelByID**
> string ReadK8sdispatcherLoglevelByID(ctx, name)
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

# **ReadK8sdispatcherNodeportRangeByID**
> string ReadK8sdispatcherNodeportRangeByID(ctx, name)
Read nodeport-range by ID

Read operation of resource: nodeport-range

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

# **ReadK8sdispatcherNodeportRuleByID**
> NodeportRule ReadK8sdispatcherNodeportRuleByID(ctx, name, nodeportPort, proto)
Read nodeport-rule by ID

Read operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 

### Return type

[**NodeportRule**](NodeportRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherNodeportRuleExternalTrafficPolicyByID**
> string ReadK8sdispatcherNodeportRuleExternalTrafficPolicyByID(ctx, name, nodeportPort, proto)
Read external-traffic-policy by ID

Read operation of resource: external-traffic-policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherNodeportRuleListByID**
> []NodeportRule ReadK8sdispatcherNodeportRuleListByID(ctx, name)
Read nodeport-rule by ID

Read operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]NodeportRule**](NodeportRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherNodeportRuleRuleNameByID**
> string ReadK8sdispatcherNodeportRuleRuleNameByID(ctx, name, nodeportPort, proto)
Read rule-name by ID

Read operation of resource: rule-name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherPortsByID**
> Ports ReadK8sdispatcherPortsByID(ctx, name, portsName)
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

# **ReadK8sdispatcherPortsIpByID**
> string ReadK8sdispatcherPortsIpByID(ctx, name, portsName)
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

# **ReadK8sdispatcherPortsListByID**
> []Ports ReadK8sdispatcherPortsListByID(ctx, name)
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

# **ReadK8sdispatcherPortsPeerByID**
> string ReadK8sdispatcherPortsPeerByID(ctx, name, portsName)
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

# **ReadK8sdispatcherPortsStatusByID**
> string ReadK8sdispatcherPortsStatusByID(ctx, name, portsName)
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

# **ReadK8sdispatcherPortsTypeByID**
> string ReadK8sdispatcherPortsTypeByID(ctx, name, portsName)
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

# **ReadK8sdispatcherPortsUuidByID**
> string ReadK8sdispatcherPortsUuidByID(ctx, name, portsName)
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

# **ReadK8sdispatcherServiceNameByID**
> string ReadK8sdispatcherServiceNameByID(ctx, name)
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

# **ReadK8sdispatcherSessionRuleByID**
> SessionRule ReadK8sdispatcherSessionRuleByID(ctx, name, direction, srcIp, dstIp, srcPort, dstPort, proto)
Read session-rule by ID

Read operation of resource: session-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **direction** | **string**| ID of direction | 
  **srcIp** | **string**| ID of src-ip | 
  **dstIp** | **string**| ID of dst-ip | 
  **srcPort** | **int32**| ID of src-port | 
  **dstPort** | **int32**| ID of dst-port | 
  **proto** | **string**| ID of proto | 

### Return type

[**SessionRule**](SessionRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherSessionRuleListByID**
> []SessionRule ReadK8sdispatcherSessionRuleListByID(ctx, name)
Read session-rule by ID

Read operation of resource: session-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]SessionRule**](SessionRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherSessionRuleNewIpByID**
> string ReadK8sdispatcherSessionRuleNewIpByID(ctx, name, direction, srcIp, dstIp, srcPort, dstPort, proto)
Read new-ip by ID

Read operation of resource: new-ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **direction** | **string**| ID of direction | 
  **srcIp** | **string**| ID of src-ip | 
  **dstIp** | **string**| ID of dst-ip | 
  **srcPort** | **int32**| ID of src-port | 
  **dstPort** | **int32**| ID of dst-port | 
  **proto** | **string**| ID of proto | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherSessionRuleNewPortByID**
> int32 ReadK8sdispatcherSessionRuleNewPortByID(ctx, name, direction, srcIp, dstIp, srcPort, dstPort, proto)
Read new-port by ID

Read operation of resource: new-port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **direction** | **string**| ID of direction | 
  **srcIp** | **string**| ID of src-ip | 
  **dstIp** | **string**| ID of dst-ip | 
  **srcPort** | **int32**| ID of src-port | 
  **dstPort** | **int32**| ID of dst-port | 
  **proto** | **string**| ID of proto | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherSessionRuleOperationByID**
> string ReadK8sdispatcherSessionRuleOperationByID(ctx, name, direction, srcIp, dstIp, srcPort, dstPort, proto)
Read operation by ID

Read operation of resource: operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **direction** | **string**| ID of direction | 
  **srcIp** | **string**| ID of src-ip | 
  **dstIp** | **string**| ID of dst-ip | 
  **srcPort** | **int32**| ID of src-port | 
  **dstPort** | **int32**| ID of dst-port | 
  **proto** | **string**| ID of proto | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherSessionRuleOriginatingRuleByID**
> string ReadK8sdispatcherSessionRuleOriginatingRuleByID(ctx, name, direction, srcIp, dstIp, srcPort, dstPort, proto)
Read originating-rule by ID

Read operation of resource: originating-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **direction** | **string**| ID of direction | 
  **srcIp** | **string**| ID of src-ip | 
  **dstIp** | **string**| ID of dst-ip | 
  **srcPort** | **int32**| ID of src-port | 
  **dstPort** | **int32**| ID of dst-port | 
  **proto** | **string**| ID of proto | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherShadowByID**
> bool ReadK8sdispatcherShadowByID(ctx, name)
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

# **ReadK8sdispatcherSpanByID**
> bool ReadK8sdispatcherSpanByID(ctx, name)
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

# **ReadK8sdispatcherTypeByID**
> string ReadK8sdispatcherTypeByID(ctx, name)
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

# **ReadK8sdispatcherUuidByID**
> string ReadK8sdispatcherUuidByID(ctx, name)
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

# **ReplaceK8sdispatcherByID**
> ReplaceK8sdispatcherByID(ctx, name, k8sdispatcher)
Replace k8sdispatcher by ID

Replace operation of resource: k8sdispatcher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **k8sdispatcher** | [**K8sdispatcher**](K8sdispatcher.md)| k8sdispatcherbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sdispatcherNodeportRuleByID**
> ReplaceK8sdispatcherNodeportRuleByID(ctx, name, nodeportPort, proto, nodeportRule)
Replace nodeport-rule by ID

Replace operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 
  **nodeportRule** | [**NodeportRule**](NodeportRule.md)| nodeport-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sdispatcherNodeportRuleListByID**
> ReplaceK8sdispatcherNodeportRuleListByID(ctx, name, nodeportRule)
Replace nodeport-rule by ID

Replace operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportRule** | [**[]NodeportRule**](NodeportRule.md)| nodeport-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sdispatcherPortsByID**
> ReplaceK8sdispatcherPortsByID(ctx, name, portsName, ports)
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

# **ReplaceK8sdispatcherPortsListByID**
> ReplaceK8sdispatcherPortsListByID(ctx, name, ports)
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

# **UpdateK8sdispatcherByID**
> UpdateK8sdispatcherByID(ctx, name, k8sdispatcher)
Update k8sdispatcher by ID

Update operation of resource: k8sdispatcher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **k8sdispatcher** | [**K8sdispatcher**](K8sdispatcher.md)| k8sdispatcherbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherListByID**
> UpdateK8sdispatcherListByID(ctx, k8sdispatcher)
Update k8sdispatcher by ID

Update operation of resource: k8sdispatcher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k8sdispatcher** | [**[]K8sdispatcher**](K8sdispatcher.md)| k8sdispatcherbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherLoglevelByID**
> UpdateK8sdispatcherLoglevelByID(ctx, name, loglevel)
Update loglevel by ID

Update operation of resource: loglevel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **loglevel** | **string**| Logging level of a cube, from none (OFF) to the most verbose (TRACE) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNodeportRangeByID**
> UpdateK8sdispatcherNodeportRangeByID(ctx, name, nodeportRange)
Update nodeport-range by ID

Update operation of resource: nodeport-range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportRange** | **string**| Port range used for NodePort Services | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNodeportRuleByID**
> UpdateK8sdispatcherNodeportRuleByID(ctx, name, nodeportPort, proto, nodeportRule)
Update nodeport-rule by ID

Update operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 
  **nodeportRule** | [**NodeportRule**](NodeportRule.md)| nodeport-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNodeportRuleExternalTrafficPolicyByID**
> UpdateK8sdispatcherNodeportRuleExternalTrafficPolicyByID(ctx, name, nodeportPort, proto, externalTrafficPolicy)
Update external-traffic-policy by ID

Update operation of resource: external-traffic-policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 
  **externalTrafficPolicy** | **string**| The external traffic policy of the Kubernetes NodePort Service | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNodeportRuleListByID**
> UpdateK8sdispatcherNodeportRuleListByID(ctx, name, nodeportRule)
Update nodeport-rule by ID

Update operation of resource: nodeport-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportRule** | [**[]NodeportRule**](NodeportRule.md)| nodeport-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherPortsByID**
> UpdateK8sdispatcherPortsByID(ctx, name, portsName, ports)
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

# **UpdateK8sdispatcherPortsListByID**
> UpdateK8sdispatcherPortsListByID(ctx, name, ports)
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

# **UpdateK8sdispatcherPortsPeerByID**
> UpdateK8sdispatcherPortsPeerByID(ctx, name, portsName, peer)
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

# **UpdateK8sdispatcherSpanByID**
> UpdateK8sdispatcherSpanByID(ctx, name, span)
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

