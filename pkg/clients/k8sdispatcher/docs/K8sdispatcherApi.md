# \K8sdispatcherApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateK8sdispatcherByID**](K8sdispatcherApi.md#CreateK8sdispatcherByID) | **Post** /k8sdispatcher/{name}/ | Create k8sdispatcher by ID
[**CreateK8sdispatcherNattingRuleByID**](K8sdispatcherApi.md#CreateK8sdispatcherNattingRuleByID) | **Post** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/ | Create natting-rule by ID
[**CreateK8sdispatcherNattingRuleListByID**](K8sdispatcherApi.md#CreateK8sdispatcherNattingRuleListByID) | **Post** /k8sdispatcher/{name}/natting-rule/ | Create natting-rule by ID
[**CreateK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#CreateK8sdispatcherNodeportRuleByID) | **Post** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Create nodeport-rule by ID
[**CreateK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#CreateK8sdispatcherNodeportRuleListByID) | **Post** /k8sdispatcher/{name}/nodeport-rule/ | Create nodeport-rule by ID
[**CreateK8sdispatcherPortsByID**](K8sdispatcherApi.md#CreateK8sdispatcherPortsByID) | **Post** /k8sdispatcher/{name}/ports/{ports_name}/ | Create ports by ID
[**CreateK8sdispatcherPortsListByID**](K8sdispatcherApi.md#CreateK8sdispatcherPortsListByID) | **Post** /k8sdispatcher/{name}/ports/ | Create ports by ID
[**DeleteK8sdispatcherByID**](K8sdispatcherApi.md#DeleteK8sdispatcherByID) | **Delete** /k8sdispatcher/{name}/ | Delete k8sdispatcher by ID
[**DeleteK8sdispatcherNattingRuleByID**](K8sdispatcherApi.md#DeleteK8sdispatcherNattingRuleByID) | **Delete** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/ | Delete natting-rule by ID
[**DeleteK8sdispatcherNattingRuleListByID**](K8sdispatcherApi.md#DeleteK8sdispatcherNattingRuleListByID) | **Delete** /k8sdispatcher/{name}/natting-rule/ | Delete natting-rule by ID
[**DeleteK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#DeleteK8sdispatcherNodeportRuleByID) | **Delete** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Delete nodeport-rule by ID
[**DeleteK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#DeleteK8sdispatcherNodeportRuleListByID) | **Delete** /k8sdispatcher/{name}/nodeport-rule/ | Delete nodeport-rule by ID
[**DeleteK8sdispatcherPortsByID**](K8sdispatcherApi.md#DeleteK8sdispatcherPortsByID) | **Delete** /k8sdispatcher/{name}/ports/{ports_name}/ | Delete ports by ID
[**DeleteK8sdispatcherPortsListByID**](K8sdispatcherApi.md#DeleteK8sdispatcherPortsListByID) | **Delete** /k8sdispatcher/{name}/ports/ | Delete ports by ID
[**ReadK8sdispatcherByID**](K8sdispatcherApi.md#ReadK8sdispatcherByID) | **Get** /k8sdispatcher/{name}/ | Read k8sdispatcher by ID
[**ReadK8sdispatcherClientSubnetByID**](K8sdispatcherApi.md#ReadK8sdispatcherClientSubnetByID) | **Get** /k8sdispatcher/{name}/client-subnet/ | Read client-subnet by ID
[**ReadK8sdispatcherClusterIpSubnetByID**](K8sdispatcherApi.md#ReadK8sdispatcherClusterIpSubnetByID) | **Get** /k8sdispatcher/{name}/cluster-ip-subnet/ | Read cluster-ip-subnet by ID
[**ReadK8sdispatcherInternalSrcIpByID**](K8sdispatcherApi.md#ReadK8sdispatcherInternalSrcIpByID) | **Get** /k8sdispatcher/{name}/internal-src-ip/ | Read internal-src-ip by ID
[**ReadK8sdispatcherListByID**](K8sdispatcherApi.md#ReadK8sdispatcherListByID) | **Get** /k8sdispatcher/ | Read k8sdispatcher by ID
[**ReadK8sdispatcherLoglevelByID**](K8sdispatcherApi.md#ReadK8sdispatcherLoglevelByID) | **Get** /k8sdispatcher/{name}/loglevel/ | Read loglevel by ID
[**ReadK8sdispatcherNattingRuleByID**](K8sdispatcherApi.md#ReadK8sdispatcherNattingRuleByID) | **Get** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/ | Read natting-rule by ID
[**ReadK8sdispatcherNattingRuleExternalIpByID**](K8sdispatcherApi.md#ReadK8sdispatcherNattingRuleExternalIpByID) | **Get** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/external-ip/ | Read external-ip by ID
[**ReadK8sdispatcherNattingRuleExternalPortByID**](K8sdispatcherApi.md#ReadK8sdispatcherNattingRuleExternalPortByID) | **Get** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/external-port/ | Read external-port by ID
[**ReadK8sdispatcherNattingRuleListByID**](K8sdispatcherApi.md#ReadK8sdispatcherNattingRuleListByID) | **Get** /k8sdispatcher/{name}/natting-rule/ | Read natting-rule by ID
[**ReadK8sdispatcherNodeportRangeByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRangeByID) | **Get** /k8sdispatcher/{name}/nodeport-range/ | Read nodeport-range by ID
[**ReadK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Read nodeport-rule by ID
[**ReadK8sdispatcherNodeportRuleInternalSrcByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleInternalSrcByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/internal-src/ | Read internal-src by ID
[**ReadK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleListByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/ | Read nodeport-rule by ID
[**ReadK8sdispatcherNodeportRuleServiceTypeByID**](K8sdispatcherApi.md#ReadK8sdispatcherNodeportRuleServiceTypeByID) | **Get** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/service-type/ | Read service-type by ID
[**ReadK8sdispatcherPortsByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/ | Read ports by ID
[**ReadK8sdispatcherPortsListByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsListByID) | **Get** /k8sdispatcher/{name}/ports/ | Read ports by ID
[**ReadK8sdispatcherPortsPeerByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsPeerByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/peer/ | Read peer by ID
[**ReadK8sdispatcherPortsStatusByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsStatusByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/status/ | Read status by ID
[**ReadK8sdispatcherPortsTypeByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsTypeByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/type/ | Read type by ID
[**ReadK8sdispatcherPortsUuidByID**](K8sdispatcherApi.md#ReadK8sdispatcherPortsUuidByID) | **Get** /k8sdispatcher/{name}/ports/{ports_name}/uuid/ | Read uuid by ID
[**ReadK8sdispatcherServiceNameByID**](K8sdispatcherApi.md#ReadK8sdispatcherServiceNameByID) | **Get** /k8sdispatcher/{name}/service-name/ | Read service-name by ID
[**ReadK8sdispatcherShadowByID**](K8sdispatcherApi.md#ReadK8sdispatcherShadowByID) | **Get** /k8sdispatcher/{name}/shadow/ | Read shadow by ID
[**ReadK8sdispatcherSpanByID**](K8sdispatcherApi.md#ReadK8sdispatcherSpanByID) | **Get** /k8sdispatcher/{name}/span/ | Read span by ID
[**ReadK8sdispatcherTypeByID**](K8sdispatcherApi.md#ReadK8sdispatcherTypeByID) | **Get** /k8sdispatcher/{name}/type/ | Read type by ID
[**ReadK8sdispatcherUuidByID**](K8sdispatcherApi.md#ReadK8sdispatcherUuidByID) | **Get** /k8sdispatcher/{name}/uuid/ | Read uuid by ID
[**ReplaceK8sdispatcherByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherByID) | **Put** /k8sdispatcher/{name}/ | Replace k8sdispatcher by ID
[**ReplaceK8sdispatcherNattingRuleByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherNattingRuleByID) | **Put** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/ | Replace natting-rule by ID
[**ReplaceK8sdispatcherNattingRuleListByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherNattingRuleListByID) | **Put** /k8sdispatcher/{name}/natting-rule/ | Replace natting-rule by ID
[**ReplaceK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherNodeportRuleByID) | **Put** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Replace nodeport-rule by ID
[**ReplaceK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherNodeportRuleListByID) | **Put** /k8sdispatcher/{name}/nodeport-rule/ | Replace nodeport-rule by ID
[**ReplaceK8sdispatcherPortsByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherPortsByID) | **Put** /k8sdispatcher/{name}/ports/{ports_name}/ | Replace ports by ID
[**ReplaceK8sdispatcherPortsListByID**](K8sdispatcherApi.md#ReplaceK8sdispatcherPortsListByID) | **Put** /k8sdispatcher/{name}/ports/ | Replace ports by ID
[**UpdateK8sdispatcherByID**](K8sdispatcherApi.md#UpdateK8sdispatcherByID) | **Patch** /k8sdispatcher/{name}/ | Update k8sdispatcher by ID
[**UpdateK8sdispatcherClientSubnetByID**](K8sdispatcherApi.md#UpdateK8sdispatcherClientSubnetByID) | **Patch** /k8sdispatcher/{name}/client-subnet/ | Update client-subnet by ID
[**UpdateK8sdispatcherClusterIpSubnetByID**](K8sdispatcherApi.md#UpdateK8sdispatcherClusterIpSubnetByID) | **Patch** /k8sdispatcher/{name}/cluster-ip-subnet/ | Update cluster-ip-subnet by ID
[**UpdateK8sdispatcherInternalSrcIpByID**](K8sdispatcherApi.md#UpdateK8sdispatcherInternalSrcIpByID) | **Patch** /k8sdispatcher/{name}/internal-src-ip/ | Update internal-src-ip by ID
[**UpdateK8sdispatcherListByID**](K8sdispatcherApi.md#UpdateK8sdispatcherListByID) | **Patch** /k8sdispatcher/ | Update k8sdispatcher by ID
[**UpdateK8sdispatcherLoglevelByID**](K8sdispatcherApi.md#UpdateK8sdispatcherLoglevelByID) | **Patch** /k8sdispatcher/{name}/loglevel/ | Update loglevel by ID
[**UpdateK8sdispatcherNattingRuleByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNattingRuleByID) | **Patch** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/ | Update natting-rule by ID
[**UpdateK8sdispatcherNattingRuleExternalIpByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNattingRuleExternalIpByID) | **Patch** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/external-ip/ | Update external-ip by ID
[**UpdateK8sdispatcherNattingRuleExternalPortByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNattingRuleExternalPortByID) | **Patch** /k8sdispatcher/{name}/natting-rule/{internal-src}/{internal-dst}/{internal-sport}/{internal-dport}/{proto}/external-port/ | Update external-port by ID
[**UpdateK8sdispatcherNattingRuleListByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNattingRuleListByID) | **Patch** /k8sdispatcher/{name}/natting-rule/ | Update natting-rule by ID
[**UpdateK8sdispatcherNodeportRangeByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRangeByID) | **Patch** /k8sdispatcher/{name}/nodeport-range/ | Update nodeport-range by ID
[**UpdateK8sdispatcherNodeportRuleByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRuleByID) | **Patch** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/ | Update nodeport-rule by ID
[**UpdateK8sdispatcherNodeportRuleInternalSrcByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRuleInternalSrcByID) | **Patch** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/internal-src/ | Update internal-src by ID
[**UpdateK8sdispatcherNodeportRuleListByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRuleListByID) | **Patch** /k8sdispatcher/{name}/nodeport-rule/ | Update nodeport-rule by ID
[**UpdateK8sdispatcherNodeportRuleServiceTypeByID**](K8sdispatcherApi.md#UpdateK8sdispatcherNodeportRuleServiceTypeByID) | **Patch** /k8sdispatcher/{name}/nodeport-rule/{nodeport-port}/{proto}/service-type/ | Update service-type by ID
[**UpdateK8sdispatcherPortsByID**](K8sdispatcherApi.md#UpdateK8sdispatcherPortsByID) | **Patch** /k8sdispatcher/{name}/ports/{ports_name}/ | Update ports by ID
[**UpdateK8sdispatcherPortsListByID**](K8sdispatcherApi.md#UpdateK8sdispatcherPortsListByID) | **Patch** /k8sdispatcher/{name}/ports/ | Update ports by ID
[**UpdateK8sdispatcherPortsPeerByID**](K8sdispatcherApi.md#UpdateK8sdispatcherPortsPeerByID) | **Patch** /k8sdispatcher/{name}/ports/{ports_name}/peer/ | Update peer by ID
[**UpdateK8sdispatcherPortsTypeByID**](K8sdispatcherApi.md#UpdateK8sdispatcherPortsTypeByID) | **Patch** /k8sdispatcher/{name}/ports/{ports_name}/type/ | Update type by ID
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

# **CreateK8sdispatcherNattingRuleByID**
> CreateK8sdispatcherNattingRuleByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto, nattingRule)
Create natting-rule by ID

Create operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 
  **nattingRule** | [**NattingRule**](NattingRule.md)| natting-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateK8sdispatcherNattingRuleListByID**
> CreateK8sdispatcherNattingRuleListByID(ctx, name, nattingRule)
Create natting-rule by ID

Create operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nattingRule** | [**[]NattingRule**](NattingRule.md)| natting-rulebody object | 

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

# **DeleteK8sdispatcherNattingRuleByID**
> DeleteK8sdispatcherNattingRuleByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto)
Delete natting-rule by ID

Delete operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteK8sdispatcherNattingRuleListByID**
> DeleteK8sdispatcherNattingRuleListByID(ctx, name)
Delete natting-rule by ID

Delete operation of resource: natting-rule

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

# **ReadK8sdispatcherClientSubnetByID**
> string ReadK8sdispatcherClientSubnetByID(ctx, name)
Read client-subnet by ID

Read operation of resource: client-subnet

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

# **ReadK8sdispatcherClusterIpSubnetByID**
> string ReadK8sdispatcherClusterIpSubnetByID(ctx, name)
Read cluster-ip-subnet by ID

Read operation of resource: cluster-ip-subnet

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

# **ReadK8sdispatcherNattingRuleByID**
> NattingRule ReadK8sdispatcherNattingRuleByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto)
Read natting-rule by ID

Read operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 

### Return type

[**NattingRule**](NattingRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherNattingRuleExternalIpByID**
> string ReadK8sdispatcherNattingRuleExternalIpByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto)
Read external-ip by ID

Read operation of resource: external-ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherNattingRuleExternalPortByID**
> int32 ReadK8sdispatcherNattingRuleExternalPortByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto)
Read external-port by ID

Read operation of resource: external-port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadK8sdispatcherNattingRuleListByID**
> []NattingRule ReadK8sdispatcherNattingRuleListByID(ctx, name)
Read natting-rule by ID

Read operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]NattingRule**](NattingRule.md)

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

# **ReadK8sdispatcherNodeportRuleInternalSrcByID**
> string ReadK8sdispatcherNodeportRuleInternalSrcByID(ctx, name, nodeportPort, proto)
Read internal-src by ID

Read operation of resource: internal-src

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

# **ReadK8sdispatcherNodeportRuleServiceTypeByID**
> string ReadK8sdispatcherNodeportRuleServiceTypeByID(ctx, name, nodeportPort, proto)
Read service-type by ID

Read operation of resource: service-type

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

# **ReplaceK8sdispatcherNattingRuleByID**
> ReplaceK8sdispatcherNattingRuleByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto, nattingRule)
Replace natting-rule by ID

Replace operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 
  **nattingRule** | [**NattingRule**](NattingRule.md)| natting-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceK8sdispatcherNattingRuleListByID**
> ReplaceK8sdispatcherNattingRuleListByID(ctx, name, nattingRule)
Replace natting-rule by ID

Replace operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nattingRule** | [**[]NattingRule**](NattingRule.md)| natting-rulebody object | 

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

# **UpdateK8sdispatcherClientSubnetByID**
> UpdateK8sdispatcherClientSubnetByID(ctx, name, clientSubnet)
Update client-subnet by ID

Update operation of resource: client-subnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **clientSubnet** | **string**| Range of IPs of pods in this node | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherClusterIpSubnetByID**
> UpdateK8sdispatcherClusterIpSubnetByID(ctx, name, clusterIpSubnet)
Update cluster-ip-subnet by ID

Update operation of resource: cluster-ip-subnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **clusterIpSubnet** | **string**| Range of VIPs where clusterIP services are exposed | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherInternalSrcIpByID**
> UpdateK8sdispatcherInternalSrcIpByID(ctx, name, internalSrcIp)
Update internal-src-ip by ID

Update operation of resource: internal-src-ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrcIp** | **string**| Internal src ip used for services with externaltrafficpolicy&#x3D;cluster | 

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
  **loglevel** | **string**| Defines the logging level of a service instance, from none (OFF) to the most verbose (TRACE) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNattingRuleByID**
> UpdateK8sdispatcherNattingRuleByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto, nattingRule)
Update natting-rule by ID

Update operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 
  **nattingRule** | [**NattingRule**](NattingRule.md)| natting-rulebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNattingRuleExternalIpByID**
> UpdateK8sdispatcherNattingRuleExternalIpByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto, externalIp)
Update external-ip by ID

Update operation of resource: external-ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 
  **externalIp** | **string**| Translated IP address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNattingRuleExternalPortByID**
> UpdateK8sdispatcherNattingRuleExternalPortByID(ctx, name, internalSrc, internalDst, internalSport, internalDport, proto, externalPort)
Update external-port by ID

Update operation of resource: external-port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **internalSrc** | **string**| ID of internal-src | 
  **internalDst** | **string**| ID of internal-dst | 
  **internalSport** | **int32**| ID of internal-sport | 
  **internalDport** | **int32**| ID of internal-dport | 
  **proto** | **string**| ID of proto | 
  **externalPort** | **int32**| Translated L4 port number | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateK8sdispatcherNattingRuleListByID**
> UpdateK8sdispatcherNattingRuleListByID(ctx, name, nattingRule)
Update natting-rule by ID

Update operation of resource: natting-rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nattingRule** | [**[]NattingRule**](NattingRule.md)| natting-rulebody object | 

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
  **nodeportRange** | **string**| Port range used for NodePort services | 

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

# **UpdateK8sdispatcherNodeportRuleInternalSrcByID**
> UpdateK8sdispatcherNodeportRuleInternalSrcByID(ctx, name, nodeportPort, proto, internalSrc)
Update internal-src by ID

Update operation of resource: internal-src

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 
  **internalSrc** | **string**| Source IP address | 

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

# **UpdateK8sdispatcherNodeportRuleServiceTypeByID**
> UpdateK8sdispatcherNodeportRuleServiceTypeByID(ctx, name, nodeportPort, proto, serviceType)
Update service-type by ID

Update operation of resource: service-type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **nodeportPort** | **int32**| ID of nodeport-port | 
  **proto** | **string**| ID of proto | 
  **serviceType** | **string**| Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoint | 

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

# **UpdateK8sdispatcherPortsTypeByID**
> UpdateK8sdispatcherPortsTypeByID(ctx, name, portsName, type_)
Update type by ID

Update operation of resource: type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **type_** | **string**| Type of the k8sdispatcher port (e.g. BACKEND or FRONTEND) | 

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

