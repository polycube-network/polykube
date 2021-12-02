# \RouterApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouterArpTableByID**](RouterApi.md#CreateRouterArpTableByID) | **Post** /router/{name}/arp-table/{address}/ | Create arp-table by ID
[**CreateRouterArpTableListByID**](RouterApi.md#CreateRouterArpTableListByID) | **Post** /router/{name}/arp-table/ | Create arp-table by ID
[**CreateRouterByID**](RouterApi.md#CreateRouterByID) | **Post** /router/{name}/ | Create router by ID
[**CreateRouterPortsByID**](RouterApi.md#CreateRouterPortsByID) | **Post** /router/{name}/ports/{ports_name}/ | Create ports by ID
[**CreateRouterPortsListByID**](RouterApi.md#CreateRouterPortsListByID) | **Post** /router/{name}/ports/ | Create ports by ID
[**CreateRouterPortsSecondaryipByID**](RouterApi.md#CreateRouterPortsSecondaryipByID) | **Post** /router/{name}/ports/{ports_name}/secondaryip/{ip}/ | Create secondaryip by ID
[**CreateRouterPortsSecondaryipListByID**](RouterApi.md#CreateRouterPortsSecondaryipListByID) | **Post** /router/{name}/ports/{ports_name}/secondaryip/ | Create secondaryip by ID
[**CreateRouterRouteByID**](RouterApi.md#CreateRouterRouteByID) | **Post** /router/{name}/route/{network}/{nexthop}/ | Create route by ID
[**CreateRouterRouteListByID**](RouterApi.md#CreateRouterRouteListByID) | **Post** /router/{name}/route/ | Create route by ID
[**DeleteRouterArpTableByID**](RouterApi.md#DeleteRouterArpTableByID) | **Delete** /router/{name}/arp-table/{address}/ | Delete arp-table by ID
[**DeleteRouterArpTableListByID**](RouterApi.md#DeleteRouterArpTableListByID) | **Delete** /router/{name}/arp-table/ | Delete arp-table by ID
[**DeleteRouterByID**](RouterApi.md#DeleteRouterByID) | **Delete** /router/{name}/ | Delete router by ID
[**DeleteRouterPortsByID**](RouterApi.md#DeleteRouterPortsByID) | **Delete** /router/{name}/ports/{ports_name}/ | Delete ports by ID
[**DeleteRouterPortsListByID**](RouterApi.md#DeleteRouterPortsListByID) | **Delete** /router/{name}/ports/ | Delete ports by ID
[**DeleteRouterPortsSecondaryipByID**](RouterApi.md#DeleteRouterPortsSecondaryipByID) | **Delete** /router/{name}/ports/{ports_name}/secondaryip/{ip}/ | Delete secondaryip by ID
[**DeleteRouterPortsSecondaryipListByID**](RouterApi.md#DeleteRouterPortsSecondaryipListByID) | **Delete** /router/{name}/ports/{ports_name}/secondaryip/ | Delete secondaryip by ID
[**DeleteRouterRouteByID**](RouterApi.md#DeleteRouterRouteByID) | **Delete** /router/{name}/route/{network}/{nexthop}/ | Delete route by ID
[**DeleteRouterRouteListByID**](RouterApi.md#DeleteRouterRouteListByID) | **Delete** /router/{name}/route/ | Delete route by ID
[**ReadRouterArpTableByID**](RouterApi.md#ReadRouterArpTableByID) | **Get** /router/{name}/arp-table/{address}/ | Read arp-table by ID
[**ReadRouterArpTableInterfaceByID**](RouterApi.md#ReadRouterArpTableInterfaceByID) | **Get** /router/{name}/arp-table/{address}/interface/ | Read interface by ID
[**ReadRouterArpTableListByID**](RouterApi.md#ReadRouterArpTableListByID) | **Get** /router/{name}/arp-table/ | Read arp-table by ID
[**ReadRouterArpTableMacByID**](RouterApi.md#ReadRouterArpTableMacByID) | **Get** /router/{name}/arp-table/{address}/mac/ | Read mac by ID
[**ReadRouterByID**](RouterApi.md#ReadRouterByID) | **Get** /router/{name}/ | Read router by ID
[**ReadRouterListByID**](RouterApi.md#ReadRouterListByID) | **Get** /router/ | Read router by ID
[**ReadRouterLoglevelByID**](RouterApi.md#ReadRouterLoglevelByID) | **Get** /router/{name}/loglevel/ | Read loglevel by ID
[**ReadRouterPortsByID**](RouterApi.md#ReadRouterPortsByID) | **Get** /router/{name}/ports/{ports_name}/ | Read ports by ID
[**ReadRouterPortsIpByID**](RouterApi.md#ReadRouterPortsIpByID) | **Get** /router/{name}/ports/{ports_name}/ip/ | Read ip by ID
[**ReadRouterPortsListByID**](RouterApi.md#ReadRouterPortsListByID) | **Get** /router/{name}/ports/ | Read ports by ID
[**ReadRouterPortsMacByID**](RouterApi.md#ReadRouterPortsMacByID) | **Get** /router/{name}/ports/{ports_name}/mac/ | Read mac by ID
[**ReadRouterPortsPeerByID**](RouterApi.md#ReadRouterPortsPeerByID) | **Get** /router/{name}/ports/{ports_name}/peer/ | Read peer by ID
[**ReadRouterPortsSecondaryipByID**](RouterApi.md#ReadRouterPortsSecondaryipByID) | **Get** /router/{name}/ports/{ports_name}/secondaryip/{ip}/ | Read secondaryip by ID
[**ReadRouterPortsSecondaryipListByID**](RouterApi.md#ReadRouterPortsSecondaryipListByID) | **Get** /router/{name}/ports/{ports_name}/secondaryip/ | Read secondaryip by ID
[**ReadRouterPortsStatusByID**](RouterApi.md#ReadRouterPortsStatusByID) | **Get** /router/{name}/ports/{ports_name}/status/ | Read status by ID
[**ReadRouterPortsUuidByID**](RouterApi.md#ReadRouterPortsUuidByID) | **Get** /router/{name}/ports/{ports_name}/uuid/ | Read uuid by ID
[**ReadRouterRouteByID**](RouterApi.md#ReadRouterRouteByID) | **Get** /router/{name}/route/{network}/{nexthop}/ | Read route by ID
[**ReadRouterRouteInterfaceByID**](RouterApi.md#ReadRouterRouteInterfaceByID) | **Get** /router/{name}/route/{network}/{nexthop}/interface/ | Read interface by ID
[**ReadRouterRouteListByID**](RouterApi.md#ReadRouterRouteListByID) | **Get** /router/{name}/route/ | Read route by ID
[**ReadRouterRoutePathcostByID**](RouterApi.md#ReadRouterRoutePathcostByID) | **Get** /router/{name}/route/{network}/{nexthop}/pathcost/ | Read pathcost by ID
[**ReadRouterServiceNameByID**](RouterApi.md#ReadRouterServiceNameByID) | **Get** /router/{name}/service-name/ | Read service-name by ID
[**ReadRouterShadowByID**](RouterApi.md#ReadRouterShadowByID) | **Get** /router/{name}/shadow/ | Read shadow by ID
[**ReadRouterSpanByID**](RouterApi.md#ReadRouterSpanByID) | **Get** /router/{name}/span/ | Read span by ID
[**ReadRouterTypeByID**](RouterApi.md#ReadRouterTypeByID) | **Get** /router/{name}/type/ | Read type by ID
[**ReadRouterUuidByID**](RouterApi.md#ReadRouterUuidByID) | **Get** /router/{name}/uuid/ | Read uuid by ID
[**ReplaceRouterArpTableByID**](RouterApi.md#ReplaceRouterArpTableByID) | **Put** /router/{name}/arp-table/{address}/ | Replace arp-table by ID
[**ReplaceRouterArpTableListByID**](RouterApi.md#ReplaceRouterArpTableListByID) | **Put** /router/{name}/arp-table/ | Replace arp-table by ID
[**ReplaceRouterByID**](RouterApi.md#ReplaceRouterByID) | **Put** /router/{name}/ | Replace router by ID
[**ReplaceRouterPortsByID**](RouterApi.md#ReplaceRouterPortsByID) | **Put** /router/{name}/ports/{ports_name}/ | Replace ports by ID
[**ReplaceRouterPortsListByID**](RouterApi.md#ReplaceRouterPortsListByID) | **Put** /router/{name}/ports/ | Replace ports by ID
[**ReplaceRouterPortsSecondaryipByID**](RouterApi.md#ReplaceRouterPortsSecondaryipByID) | **Put** /router/{name}/ports/{ports_name}/secondaryip/{ip}/ | Replace secondaryip by ID
[**ReplaceRouterPortsSecondaryipListByID**](RouterApi.md#ReplaceRouterPortsSecondaryipListByID) | **Put** /router/{name}/ports/{ports_name}/secondaryip/ | Replace secondaryip by ID
[**ReplaceRouterRouteByID**](RouterApi.md#ReplaceRouterRouteByID) | **Put** /router/{name}/route/{network}/{nexthop}/ | Replace route by ID
[**ReplaceRouterRouteListByID**](RouterApi.md#ReplaceRouterRouteListByID) | **Put** /router/{name}/route/ | Replace route by ID
[**UpdateRouterArpTableByID**](RouterApi.md#UpdateRouterArpTableByID) | **Patch** /router/{name}/arp-table/{address}/ | Update arp-table by ID
[**UpdateRouterArpTableInterfaceByID**](RouterApi.md#UpdateRouterArpTableInterfaceByID) | **Patch** /router/{name}/arp-table/{address}/interface/ | Update interface by ID
[**UpdateRouterArpTableListByID**](RouterApi.md#UpdateRouterArpTableListByID) | **Patch** /router/{name}/arp-table/ | Update arp-table by ID
[**UpdateRouterArpTableMacByID**](RouterApi.md#UpdateRouterArpTableMacByID) | **Patch** /router/{name}/arp-table/{address}/mac/ | Update mac by ID
[**UpdateRouterByID**](RouterApi.md#UpdateRouterByID) | **Patch** /router/{name}/ | Update router by ID
[**UpdateRouterListByID**](RouterApi.md#UpdateRouterListByID) | **Patch** /router/ | Update router by ID
[**UpdateRouterLoglevelByID**](RouterApi.md#UpdateRouterLoglevelByID) | **Patch** /router/{name}/loglevel/ | Update loglevel by ID
[**UpdateRouterPortsByID**](RouterApi.md#UpdateRouterPortsByID) | **Patch** /router/{name}/ports/{ports_name}/ | Update ports by ID
[**UpdateRouterPortsIpByID**](RouterApi.md#UpdateRouterPortsIpByID) | **Patch** /router/{name}/ports/{ports_name}/ip/ | Update ip by ID
[**UpdateRouterPortsListByID**](RouterApi.md#UpdateRouterPortsListByID) | **Patch** /router/{name}/ports/ | Update ports by ID
[**UpdateRouterPortsMacByID**](RouterApi.md#UpdateRouterPortsMacByID) | **Patch** /router/{name}/ports/{ports_name}/mac/ | Update mac by ID
[**UpdateRouterPortsPeerByID**](RouterApi.md#UpdateRouterPortsPeerByID) | **Patch** /router/{name}/ports/{ports_name}/peer/ | Update peer by ID
[**UpdateRouterPortsSecondaryipByID**](RouterApi.md#UpdateRouterPortsSecondaryipByID) | **Patch** /router/{name}/ports/{ports_name}/secondaryip/{ip}/ | Update secondaryip by ID
[**UpdateRouterPortsSecondaryipListByID**](RouterApi.md#UpdateRouterPortsSecondaryipListByID) | **Patch** /router/{name}/ports/{ports_name}/secondaryip/ | Update secondaryip by ID
[**UpdateRouterRouteByID**](RouterApi.md#UpdateRouterRouteByID) | **Patch** /router/{name}/route/{network}/{nexthop}/ | Update route by ID
[**UpdateRouterRouteListByID**](RouterApi.md#UpdateRouterRouteListByID) | **Patch** /router/{name}/route/ | Update route by ID
[**UpdateRouterRoutePathcostByID**](RouterApi.md#UpdateRouterRoutePathcostByID) | **Patch** /router/{name}/route/{network}/{nexthop}/pathcost/ | Update pathcost by ID
[**UpdateRouterSpanByID**](RouterApi.md#UpdateRouterSpanByID) | **Patch** /router/{name}/span/ | Update span by ID


# **CreateRouterArpTableByID**
> CreateRouterArpTableByID(ctx, name, address, arpTable)
Create arp-table by ID

Create operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **arpTable** | [**ArpTable**](ArpTable.md)| arp-tablebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouterArpTableListByID**
> CreateRouterArpTableListByID(ctx, name, arpTable)
Create arp-table by ID

Create operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **arpTable** | [**[]ArpTable**](ArpTable.md)| arp-tablebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouterByID**
> CreateRouterByID(ctx, name, router)
Create router by ID

Create operation of resource: router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **router** | [**Router**](Router.md)| routerbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouterPortsByID**
> CreateRouterPortsByID(ctx, name, portsName, ports)
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

# **CreateRouterPortsListByID**
> CreateRouterPortsListByID(ctx, name, ports)
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

# **CreateRouterPortsSecondaryipByID**
> CreateRouterPortsSecondaryipByID(ctx, name, portsName, ip, secondaryip)
Create secondaryip by ID

Create operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ip** | **string**| ID of ip | 
  **secondaryip** | [**PortsSecondaryip**](PortsSecondaryip.md)| secondaryipbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouterPortsSecondaryipListByID**
> CreateRouterPortsSecondaryipListByID(ctx, name, portsName, secondaryip)
Create secondaryip by ID

Create operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **secondaryip** | [**[]PortsSecondaryip**](PortsSecondaryip.md)| secondaryipbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouterRouteByID**
> CreateRouterRouteByID(ctx, name, network, nexthop, route)
Create route by ID

Create operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 
  **route** | [**Route**](Route.md)| routebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouterRouteListByID**
> CreateRouterRouteListByID(ctx, name, route)
Create route by ID

Create operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **route** | [**[]Route**](Route.md)| routebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouterArpTableByID**
> DeleteRouterArpTableByID(ctx, name, address)
Delete arp-table by ID

Delete operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouterArpTableListByID**
> DeleteRouterArpTableListByID(ctx, name)
Delete arp-table by ID

Delete operation of resource: arp-table

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

# **DeleteRouterByID**
> DeleteRouterByID(ctx, name)
Delete router by ID

Delete operation of resource: router

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

# **DeleteRouterPortsByID**
> DeleteRouterPortsByID(ctx, name, portsName)
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

# **DeleteRouterPortsListByID**
> DeleteRouterPortsListByID(ctx, name)
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

# **DeleteRouterPortsSecondaryipByID**
> DeleteRouterPortsSecondaryipByID(ctx, name, portsName, ip)
Delete secondaryip by ID

Delete operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ip** | **string**| ID of ip | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouterPortsSecondaryipListByID**
> DeleteRouterPortsSecondaryipListByID(ctx, name, portsName)
Delete secondaryip by ID

Delete operation of resource: secondaryip

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

# **DeleteRouterRouteByID**
> DeleteRouterRouteByID(ctx, name, network, nexthop)
Delete route by ID

Delete operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouterRouteListByID**
> DeleteRouterRouteListByID(ctx, name)
Delete route by ID

Delete operation of resource: route

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

# **ReadRouterArpTableByID**
> ArpTable ReadRouterArpTableByID(ctx, name, address)
Read arp-table by ID

Read operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 

### Return type

[**ArpTable**](ArpTable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterArpTableInterfaceByID**
> string ReadRouterArpTableInterfaceByID(ctx, name, address)
Read interface by ID

Read operation of resource: interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterArpTableListByID**
> []ArpTable ReadRouterArpTableListByID(ctx, name)
Read arp-table by ID

Read operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]ArpTable**](ArpTable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterArpTableMacByID**
> string ReadRouterArpTableMacByID(ctx, name, address)
Read mac by ID

Read operation of resource: mac

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterByID**
> Router ReadRouterByID(ctx, name)
Read router by ID

Read operation of resource: router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**Router**](Router.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterListByID**
> []Router ReadRouterListByID(ctx, )
Read router by ID

Read operation of resource: router

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Router**](Router.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterLoglevelByID**
> string ReadRouterLoglevelByID(ctx, name)
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

# **ReadRouterPortsByID**
> Ports ReadRouterPortsByID(ctx, name, portsName)
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

# **ReadRouterPortsIpByID**
> string ReadRouterPortsIpByID(ctx, name, portsName)
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

# **ReadRouterPortsListByID**
> []Ports ReadRouterPortsListByID(ctx, name)
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

# **ReadRouterPortsMacByID**
> string ReadRouterPortsMacByID(ctx, name, portsName)
Read mac by ID

Read operation of resource: mac

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

# **ReadRouterPortsPeerByID**
> string ReadRouterPortsPeerByID(ctx, name, portsName)
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

# **ReadRouterPortsSecondaryipByID**
> PortsSecondaryip ReadRouterPortsSecondaryipByID(ctx, name, portsName, ip)
Read secondaryip by ID

Read operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ip** | **string**| ID of ip | 

### Return type

[**PortsSecondaryip**](PortsSecondaryip.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterPortsSecondaryipListByID**
> []PortsSecondaryip ReadRouterPortsSecondaryipListByID(ctx, name, portsName)
Read secondaryip by ID

Read operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 

### Return type

[**[]PortsSecondaryip**](PortsSecondaryip.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterPortsStatusByID**
> string ReadRouterPortsStatusByID(ctx, name, portsName)
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

# **ReadRouterPortsUuidByID**
> string ReadRouterPortsUuidByID(ctx, name, portsName)
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

# **ReadRouterRouteByID**
> Route ReadRouterRouteByID(ctx, name, network, nexthop)
Read route by ID

Read operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 

### Return type

[**Route**](Route.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterRouteInterfaceByID**
> string ReadRouterRouteInterfaceByID(ctx, name, network, nexthop)
Read interface by ID

Read operation of resource: interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterRouteListByID**
> []Route ReadRouterRouteListByID(ctx, name)
Read route by ID

Read operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]Route**](Route.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterRoutePathcostByID**
> int32 ReadRouterRoutePathcostByID(ctx, name, network, nexthop)
Read pathcost by ID

Read operation of resource: pathcost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouterServiceNameByID**
> string ReadRouterServiceNameByID(ctx, name)
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

# **ReadRouterShadowByID**
> bool ReadRouterShadowByID(ctx, name)
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

# **ReadRouterSpanByID**
> bool ReadRouterSpanByID(ctx, name)
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

# **ReadRouterTypeByID**
> string ReadRouterTypeByID(ctx, name)
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

# **ReadRouterUuidByID**
> string ReadRouterUuidByID(ctx, name)
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

# **ReplaceRouterArpTableByID**
> ReplaceRouterArpTableByID(ctx, name, address, arpTable)
Replace arp-table by ID

Replace operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **arpTable** | [**ArpTable**](ArpTable.md)| arp-tablebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRouterArpTableListByID**
> ReplaceRouterArpTableListByID(ctx, name, arpTable)
Replace arp-table by ID

Replace operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **arpTable** | [**[]ArpTable**](ArpTable.md)| arp-tablebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRouterByID**
> ReplaceRouterByID(ctx, name, router)
Replace router by ID

Replace operation of resource: router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **router** | [**Router**](Router.md)| routerbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRouterPortsByID**
> ReplaceRouterPortsByID(ctx, name, portsName, ports)
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

# **ReplaceRouterPortsListByID**
> ReplaceRouterPortsListByID(ctx, name, ports)
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

# **ReplaceRouterPortsSecondaryipByID**
> ReplaceRouterPortsSecondaryipByID(ctx, name, portsName, ip, secondaryip)
Replace secondaryip by ID

Replace operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ip** | **string**| ID of ip | 
  **secondaryip** | [**PortsSecondaryip**](PortsSecondaryip.md)| secondaryipbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRouterPortsSecondaryipListByID**
> ReplaceRouterPortsSecondaryipListByID(ctx, name, portsName, secondaryip)
Replace secondaryip by ID

Replace operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **secondaryip** | [**[]PortsSecondaryip**](PortsSecondaryip.md)| secondaryipbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRouterRouteByID**
> ReplaceRouterRouteByID(ctx, name, network, nexthop, route)
Replace route by ID

Replace operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 
  **route** | [**Route**](Route.md)| routebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRouterRouteListByID**
> ReplaceRouterRouteListByID(ctx, name, route)
Replace route by ID

Replace operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **route** | [**[]Route**](Route.md)| routebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterArpTableByID**
> UpdateRouterArpTableByID(ctx, name, address, arpTable)
Update arp-table by ID

Update operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **arpTable** | [**ArpTable**](ArpTable.md)| arp-tablebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterArpTableInterfaceByID**
> UpdateRouterArpTableInterfaceByID(ctx, name, address, interface_)
Update interface by ID

Update operation of resource: interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **interface_** | **string**| Outgoing interface | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterArpTableListByID**
> UpdateRouterArpTableListByID(ctx, name, arpTable)
Update arp-table by ID

Update operation of resource: arp-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **arpTable** | [**[]ArpTable**](ArpTable.md)| arp-tablebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterArpTableMacByID**
> UpdateRouterArpTableMacByID(ctx, name, address, mac)
Update mac by ID

Update operation of resource: mac

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **mac** | **string**| Destination MAC address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterByID**
> UpdateRouterByID(ctx, name, router)
Update router by ID

Update operation of resource: router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **router** | [**Router**](Router.md)| routerbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterListByID**
> UpdateRouterListByID(ctx, router)
Update router by ID

Update operation of resource: router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **router** | [**[]Router**](Router.md)| routerbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterLoglevelByID**
> UpdateRouterLoglevelByID(ctx, name, loglevel)
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

# **UpdateRouterPortsByID**
> UpdateRouterPortsByID(ctx, name, portsName, ports)
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

# **UpdateRouterPortsIpByID**
> UpdateRouterPortsIpByID(ctx, name, portsName, ip)
Update ip by ID

Update operation of resource: ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ip** | **string**| IP address and prefix of the port | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterPortsListByID**
> UpdateRouterPortsListByID(ctx, name, ports)
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

# **UpdateRouterPortsMacByID**
> UpdateRouterPortsMacByID(ctx, name, portsName, mac)
Update mac by ID

Update operation of resource: mac

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **mac** | **string**| MAC address of the port | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterPortsPeerByID**
> UpdateRouterPortsPeerByID(ctx, name, portsName, peer)
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

# **UpdateRouterPortsSecondaryipByID**
> UpdateRouterPortsSecondaryipByID(ctx, name, portsName, ip, secondaryip)
Update secondaryip by ID

Update operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **ip** | **string**| ID of ip | 
  **secondaryip** | [**PortsSecondaryip**](PortsSecondaryip.md)| secondaryipbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterPortsSecondaryipListByID**
> UpdateRouterPortsSecondaryipListByID(ctx, name, portsName, secondaryip)
Update secondaryip by ID

Update operation of resource: secondaryip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **portsName** | **string**| ID of ports_name | 
  **secondaryip** | [**[]PortsSecondaryip**](PortsSecondaryip.md)| secondaryipbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterRouteByID**
> UpdateRouterRouteByID(ctx, name, network, nexthop, route)
Update route by ID

Update operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 
  **route** | [**Route**](Route.md)| routebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterRouteListByID**
> UpdateRouterRouteListByID(ctx, name, route)
Update route by ID

Update operation of resource: route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **route** | [**[]Route**](Route.md)| routebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterRoutePathcostByID**
> UpdateRouterRoutePathcostByID(ctx, name, network, nexthop, pathcost)
Update pathcost by ID

Update operation of resource: pathcost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **network** | **string**| ID of network | 
  **nexthop** | **string**| ID of nexthop | 
  **pathcost** | **int32**| Cost of this route | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouterSpanByID**
> UpdateRouterSpanByID(ctx, name, span)
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

