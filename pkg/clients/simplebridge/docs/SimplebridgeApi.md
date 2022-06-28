# \SimplebridgeApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSimplebridgeByID**](SimplebridgeApi.md#CreateSimplebridgeByID) | **Post** /simplebridge/{name}/ | Create simplebridge by ID
[**CreateSimplebridgeFdbByID**](SimplebridgeApi.md#CreateSimplebridgeFdbByID) | **Post** /simplebridge/{name}/fdb/ | Create fdb by ID
[**CreateSimplebridgeFdbEntryByID**](SimplebridgeApi.md#CreateSimplebridgeFdbEntryByID) | **Post** /simplebridge/{name}/fdb/entry/{address}/ | Create entry by ID
[**CreateSimplebridgeFdbEntryListByID**](SimplebridgeApi.md#CreateSimplebridgeFdbEntryListByID) | **Post** /simplebridge/{name}/fdb/entry/ | Create entry by ID
[**CreateSimplebridgeFdbFlushByID**](SimplebridgeApi.md#CreateSimplebridgeFdbFlushByID) | **Post** /simplebridge/{name}/fdb/flush/ | Create flush by ID
[**CreateSimplebridgePortsByID**](SimplebridgeApi.md#CreateSimplebridgePortsByID) | **Post** /simplebridge/{name}/ports/{ports_name}/ | Create ports by ID
[**CreateSimplebridgePortsListByID**](SimplebridgeApi.md#CreateSimplebridgePortsListByID) | **Post** /simplebridge/{name}/ports/ | Create ports by ID
[**DeleteSimplebridgeByID**](SimplebridgeApi.md#DeleteSimplebridgeByID) | **Delete** /simplebridge/{name}/ | Delete simplebridge by ID
[**DeleteSimplebridgeFdbByID**](SimplebridgeApi.md#DeleteSimplebridgeFdbByID) | **Delete** /simplebridge/{name}/fdb/ | Delete fdb by ID
[**DeleteSimplebridgeFdbEntryByID**](SimplebridgeApi.md#DeleteSimplebridgeFdbEntryByID) | **Delete** /simplebridge/{name}/fdb/entry/{address}/ | Delete entry by ID
[**DeleteSimplebridgeFdbEntryListByID**](SimplebridgeApi.md#DeleteSimplebridgeFdbEntryListByID) | **Delete** /simplebridge/{name}/fdb/entry/ | Delete entry by ID
[**DeleteSimplebridgePortsByID**](SimplebridgeApi.md#DeleteSimplebridgePortsByID) | **Delete** /simplebridge/{name}/ports/{ports_name}/ | Delete ports by ID
[**DeleteSimplebridgePortsListByID**](SimplebridgeApi.md#DeleteSimplebridgePortsListByID) | **Delete** /simplebridge/{name}/ports/ | Delete ports by ID
[**ReadSimplebridgeByID**](SimplebridgeApi.md#ReadSimplebridgeByID) | **Get** /simplebridge/{name}/ | Read simplebridge by ID
[**ReadSimplebridgeFdbAgingTimeByID**](SimplebridgeApi.md#ReadSimplebridgeFdbAgingTimeByID) | **Get** /simplebridge/{name}/fdb/aging-time/ | Read aging-time by ID
[**ReadSimplebridgeFdbByID**](SimplebridgeApi.md#ReadSimplebridgeFdbByID) | **Get** /simplebridge/{name}/fdb/ | Read fdb by ID
[**ReadSimplebridgeFdbEntryAgeByID**](SimplebridgeApi.md#ReadSimplebridgeFdbEntryAgeByID) | **Get** /simplebridge/{name}/fdb/entry/{address}/age/ | Read age by ID
[**ReadSimplebridgeFdbEntryByID**](SimplebridgeApi.md#ReadSimplebridgeFdbEntryByID) | **Get** /simplebridge/{name}/fdb/entry/{address}/ | Read entry by ID
[**ReadSimplebridgeFdbEntryListByID**](SimplebridgeApi.md#ReadSimplebridgeFdbEntryListByID) | **Get** /simplebridge/{name}/fdb/entry/ | Read entry by ID
[**ReadSimplebridgeFdbEntryPortByID**](SimplebridgeApi.md#ReadSimplebridgeFdbEntryPortByID) | **Get** /simplebridge/{name}/fdb/entry/{address}/port/ | Read port by ID
[**ReadSimplebridgeListByID**](SimplebridgeApi.md#ReadSimplebridgeListByID) | **Get** /simplebridge/ | Read simplebridge by ID
[**ReadSimplebridgeLoglevelByID**](SimplebridgeApi.md#ReadSimplebridgeLoglevelByID) | **Get** /simplebridge/{name}/loglevel/ | Read loglevel by ID
[**ReadSimplebridgePortsByID**](SimplebridgeApi.md#ReadSimplebridgePortsByID) | **Get** /simplebridge/{name}/ports/{ports_name}/ | Read ports by ID
[**ReadSimplebridgePortsListByID**](SimplebridgeApi.md#ReadSimplebridgePortsListByID) | **Get** /simplebridge/{name}/ports/ | Read ports by ID
[**ReadSimplebridgePortsPeerByID**](SimplebridgeApi.md#ReadSimplebridgePortsPeerByID) | **Get** /simplebridge/{name}/ports/{ports_name}/peer/ | Read peer by ID
[**ReadSimplebridgePortsStatusByID**](SimplebridgeApi.md#ReadSimplebridgePortsStatusByID) | **Get** /simplebridge/{name}/ports/{ports_name}/status/ | Read status by ID
[**ReadSimplebridgePortsUuidByID**](SimplebridgeApi.md#ReadSimplebridgePortsUuidByID) | **Get** /simplebridge/{name}/ports/{ports_name}/uuid/ | Read uuid by ID
[**ReadSimplebridgeServiceNameByID**](SimplebridgeApi.md#ReadSimplebridgeServiceNameByID) | **Get** /simplebridge/{name}/service-name/ | Read service-name by ID
[**ReadSimplebridgeShadowByID**](SimplebridgeApi.md#ReadSimplebridgeShadowByID) | **Get** /simplebridge/{name}/shadow/ | Read shadow by ID
[**ReadSimplebridgeSpanByID**](SimplebridgeApi.md#ReadSimplebridgeSpanByID) | **Get** /simplebridge/{name}/span/ | Read span by ID
[**ReadSimplebridgeTypeByID**](SimplebridgeApi.md#ReadSimplebridgeTypeByID) | **Get** /simplebridge/{name}/type/ | Read type by ID
[**ReadSimplebridgeUuidByID**](SimplebridgeApi.md#ReadSimplebridgeUuidByID) | **Get** /simplebridge/{name}/uuid/ | Read uuid by ID
[**ReplaceSimplebridgeByID**](SimplebridgeApi.md#ReplaceSimplebridgeByID) | **Put** /simplebridge/{name}/ | Replace simplebridge by ID
[**ReplaceSimplebridgeFdbByID**](SimplebridgeApi.md#ReplaceSimplebridgeFdbByID) | **Put** /simplebridge/{name}/fdb/ | Replace fdb by ID
[**ReplaceSimplebridgeFdbEntryByID**](SimplebridgeApi.md#ReplaceSimplebridgeFdbEntryByID) | **Put** /simplebridge/{name}/fdb/entry/{address}/ | Replace entry by ID
[**ReplaceSimplebridgeFdbEntryListByID**](SimplebridgeApi.md#ReplaceSimplebridgeFdbEntryListByID) | **Put** /simplebridge/{name}/fdb/entry/ | Replace entry by ID
[**ReplaceSimplebridgePortsByID**](SimplebridgeApi.md#ReplaceSimplebridgePortsByID) | **Put** /simplebridge/{name}/ports/{ports_name}/ | Replace ports by ID
[**ReplaceSimplebridgePortsListByID**](SimplebridgeApi.md#ReplaceSimplebridgePortsListByID) | **Put** /simplebridge/{name}/ports/ | Replace ports by ID
[**UpdateSimplebridgeByID**](SimplebridgeApi.md#UpdateSimplebridgeByID) | **Patch** /simplebridge/{name}/ | Update simplebridge by ID
[**UpdateSimplebridgeFdbAgingTimeByID**](SimplebridgeApi.md#UpdateSimplebridgeFdbAgingTimeByID) | **Patch** /simplebridge/{name}/fdb/aging-time/ | Update aging-time by ID
[**UpdateSimplebridgeFdbByID**](SimplebridgeApi.md#UpdateSimplebridgeFdbByID) | **Patch** /simplebridge/{name}/fdb/ | Update fdb by ID
[**UpdateSimplebridgeFdbEntryByID**](SimplebridgeApi.md#UpdateSimplebridgeFdbEntryByID) | **Patch** /simplebridge/{name}/fdb/entry/{address}/ | Update entry by ID
[**UpdateSimplebridgeFdbEntryListByID**](SimplebridgeApi.md#UpdateSimplebridgeFdbEntryListByID) | **Patch** /simplebridge/{name}/fdb/entry/ | Update entry by ID
[**UpdateSimplebridgeFdbEntryPortByID**](SimplebridgeApi.md#UpdateSimplebridgeFdbEntryPortByID) | **Patch** /simplebridge/{name}/fdb/entry/{address}/port/ | Update port by ID
[**UpdateSimplebridgeListByID**](SimplebridgeApi.md#UpdateSimplebridgeListByID) | **Patch** /simplebridge/ | Update simplebridge by ID
[**UpdateSimplebridgeLoglevelByID**](SimplebridgeApi.md#UpdateSimplebridgeLoglevelByID) | **Patch** /simplebridge/{name}/loglevel/ | Update loglevel by ID
[**UpdateSimplebridgePortsByID**](SimplebridgeApi.md#UpdateSimplebridgePortsByID) | **Patch** /simplebridge/{name}/ports/{ports_name}/ | Update ports by ID
[**UpdateSimplebridgePortsListByID**](SimplebridgeApi.md#UpdateSimplebridgePortsListByID) | **Patch** /simplebridge/{name}/ports/ | Update ports by ID
[**UpdateSimplebridgePortsPeerByID**](SimplebridgeApi.md#UpdateSimplebridgePortsPeerByID) | **Patch** /simplebridge/{name}/ports/{ports_name}/peer/ | Update peer by ID
[**UpdateSimplebridgeSpanByID**](SimplebridgeApi.md#UpdateSimplebridgeSpanByID) | **Patch** /simplebridge/{name}/span/ | Update span by ID


# **CreateSimplebridgeByID**
> CreateSimplebridgeByID(ctx, name, simplebridge)
Create simplebridge by ID

Create operation of resource: simplebridge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **simplebridge** | [**Simplebridge**](Simplebridge.md)| simplebridgebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSimplebridgeFdbByID**
> CreateSimplebridgeFdbByID(ctx, name, fdb)
Create fdb by ID

Create operation of resource: fdb

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **fdb** | [**Fdb**](Fdb.md)| fdbbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSimplebridgeFdbEntryByID**
> CreateSimplebridgeFdbEntryByID(ctx, name, address, entry)
Create entry by ID

Create operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **entry** | [**FdbEntry**](FdbEntry.md)| entrybody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSimplebridgeFdbEntryListByID**
> CreateSimplebridgeFdbEntryListByID(ctx, name, entry)
Create entry by ID

Create operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **entry** | [**[]FdbEntry**](FdbEntry.md)| entrybody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSimplebridgeFdbFlushByID**
> FdbFlushOutput CreateSimplebridgeFdbFlushByID(ctx, name)
Create flush by ID

Create operation of resource: flush

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**FdbFlushOutput**](FdbFlushOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSimplebridgePortsByID**
> CreateSimplebridgePortsByID(ctx, name, portsName, ports)
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

# **CreateSimplebridgePortsListByID**
> CreateSimplebridgePortsListByID(ctx, name, ports)
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

# **DeleteSimplebridgeByID**
> DeleteSimplebridgeByID(ctx, name)
Delete simplebridge by ID

Delete operation of resource: simplebridge

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

# **DeleteSimplebridgeFdbByID**
> DeleteSimplebridgeFdbByID(ctx, name)
Delete fdb by ID

Delete operation of resource: fdb

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

# **DeleteSimplebridgeFdbEntryByID**
> DeleteSimplebridgeFdbEntryByID(ctx, name, address)
Delete entry by ID

Delete operation of resource: entry

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

# **DeleteSimplebridgeFdbEntryListByID**
> DeleteSimplebridgeFdbEntryListByID(ctx, name)
Delete entry by ID

Delete operation of resource: entry

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

# **DeleteSimplebridgePortsByID**
> DeleteSimplebridgePortsByID(ctx, name, portsName)
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

# **DeleteSimplebridgePortsListByID**
> DeleteSimplebridgePortsListByID(ctx, name)
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

# **ReadSimplebridgeByID**
> Simplebridge ReadSimplebridgeByID(ctx, name)
Read simplebridge by ID

Read operation of resource: simplebridge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**Simplebridge**](Simplebridge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimplebridgeFdbAgingTimeByID**
> int32 ReadSimplebridgeFdbAgingTimeByID(ctx, name)
Read aging-time by ID

Read operation of resource: aging-time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimplebridgeFdbByID**
> Fdb ReadSimplebridgeFdbByID(ctx, name)
Read fdb by ID

Read operation of resource: fdb

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**Fdb**](Fdb.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimplebridgeFdbEntryAgeByID**
> int32 ReadSimplebridgeFdbEntryAgeByID(ctx, name, address)
Read age by ID

Read operation of resource: age

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimplebridgeFdbEntryByID**
> FdbEntry ReadSimplebridgeFdbEntryByID(ctx, name, address)
Read entry by ID

Read operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 

### Return type

[**FdbEntry**](FdbEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimplebridgeFdbEntryListByID**
> []FdbEntry ReadSimplebridgeFdbEntryListByID(ctx, name)
Read entry by ID

Read operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 

### Return type

[**[]FdbEntry**](FdbEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimplebridgeFdbEntryPortByID**
> string ReadSimplebridgeFdbEntryPortByID(ctx, name, address)
Read port by ID

Read operation of resource: port

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

# **ReadSimplebridgeListByID**
> []Simplebridge ReadSimplebridgeListByID(ctx, )
Read simplebridge by ID

Read operation of resource: simplebridge

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Simplebridge**](Simplebridge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimplebridgeLoglevelByID**
> string ReadSimplebridgeLoglevelByID(ctx, name)
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

# **ReadSimplebridgePortsByID**
> Ports ReadSimplebridgePortsByID(ctx, name, portsName)
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

# **ReadSimplebridgePortsListByID**
> []Ports ReadSimplebridgePortsListByID(ctx, name)
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

# **ReadSimplebridgePortsPeerByID**
> string ReadSimplebridgePortsPeerByID(ctx, name, portsName)
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

# **ReadSimplebridgePortsStatusByID**
> string ReadSimplebridgePortsStatusByID(ctx, name, portsName)
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

# **ReadSimplebridgePortsUuidByID**
> string ReadSimplebridgePortsUuidByID(ctx, name, portsName)
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

# **ReadSimplebridgeServiceNameByID**
> string ReadSimplebridgeServiceNameByID(ctx, name)
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

# **ReadSimplebridgeShadowByID**
> bool ReadSimplebridgeShadowByID(ctx, name)
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

# **ReadSimplebridgeSpanByID**
> bool ReadSimplebridgeSpanByID(ctx, name)
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

# **ReadSimplebridgeTypeByID**
> string ReadSimplebridgeTypeByID(ctx, name)
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

# **ReadSimplebridgeUuidByID**
> string ReadSimplebridgeUuidByID(ctx, name)
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

# **ReplaceSimplebridgeByID**
> ReplaceSimplebridgeByID(ctx, name, simplebridge)
Replace simplebridge by ID

Replace operation of resource: simplebridge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **simplebridge** | [**Simplebridge**](Simplebridge.md)| simplebridgebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSimplebridgeFdbByID**
> ReplaceSimplebridgeFdbByID(ctx, name, fdb)
Replace fdb by ID

Replace operation of resource: fdb

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **fdb** | [**Fdb**](Fdb.md)| fdbbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSimplebridgeFdbEntryByID**
> ReplaceSimplebridgeFdbEntryByID(ctx, name, address, entry)
Replace entry by ID

Replace operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **entry** | [**FdbEntry**](FdbEntry.md)| entrybody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSimplebridgeFdbEntryListByID**
> ReplaceSimplebridgeFdbEntryListByID(ctx, name, entry)
Replace entry by ID

Replace operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **entry** | [**[]FdbEntry**](FdbEntry.md)| entrybody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSimplebridgePortsByID**
> ReplaceSimplebridgePortsByID(ctx, name, portsName, ports)
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

# **ReplaceSimplebridgePortsListByID**
> ReplaceSimplebridgePortsListByID(ctx, name, ports)
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

# **UpdateSimplebridgeByID**
> UpdateSimplebridgeByID(ctx, name, simplebridge)
Update simplebridge by ID

Update operation of resource: simplebridge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **simplebridge** | [**Simplebridge**](Simplebridge.md)| simplebridgebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSimplebridgeFdbAgingTimeByID**
> UpdateSimplebridgeFdbAgingTimeByID(ctx, name, agingTime)
Update aging-time by ID

Update operation of resource: aging-time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **agingTime** | **int32**| Aging time of the filtering database (in seconds) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSimplebridgeFdbByID**
> UpdateSimplebridgeFdbByID(ctx, name, fdb)
Update fdb by ID

Update operation of resource: fdb

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **fdb** | [**Fdb**](Fdb.md)| fdbbody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSimplebridgeFdbEntryByID**
> UpdateSimplebridgeFdbEntryByID(ctx, name, address, entry)
Update entry by ID

Update operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **entry** | [**FdbEntry**](FdbEntry.md)| entrybody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSimplebridgeFdbEntryListByID**
> UpdateSimplebridgeFdbEntryListByID(ctx, name, entry)
Update entry by ID

Update operation of resource: entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **entry** | [**[]FdbEntry**](FdbEntry.md)| entrybody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSimplebridgeFdbEntryPortByID**
> UpdateSimplebridgeFdbEntryPortByID(ctx, name, address, port)
Update port by ID

Update operation of resource: port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ID of name | 
  **address** | **string**| ID of address | 
  **port** | **string**| Output port name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSimplebridgeListByID**
> UpdateSimplebridgeListByID(ctx, simplebridge)
Update simplebridge by ID

Update operation of resource: simplebridge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **simplebridge** | [**[]Simplebridge**](Simplebridge.md)| simplebridgebody object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSimplebridgeLoglevelByID**
> UpdateSimplebridgeLoglevelByID(ctx, name, loglevel)
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

# **UpdateSimplebridgePortsByID**
> UpdateSimplebridgePortsByID(ctx, name, portsName, ports)
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

# **UpdateSimplebridgePortsListByID**
> UpdateSimplebridgePortsListByID(ctx, name, ports)
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

# **UpdateSimplebridgePortsPeerByID**
> UpdateSimplebridgePortsPeerByID(ctx, name, portsName, peer)
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

# **UpdateSimplebridgeSpanByID**
> UpdateSimplebridgeSpanByID(ctx, name, span)
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

