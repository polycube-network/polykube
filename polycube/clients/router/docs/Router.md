# Router

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the router service | [optional] [default to null]
**Uuid** | **string** | UUID of the Cube | [optional] [default to null]
**Type_** | **string** | Type of the Cube (TC, XDP_SKB, XDP_DRV) | [optional] [default to null]
**ServiceName** | **string** |  | [optional] [default to null]
**Loglevel** | **string** | Defines the logging level of a service instance, from none (OFF) to the most verbose (TRACE) | [optional] [default to null]
**Ports** | [**[]Ports**](Ports.md) | Entry of the ports table | [optional] [default to null]
**Shadow** | **bool** | Defines if the service is visible in Linux | [optional] [default to null]
**Span** | **bool** | Defines if all traffic is sent to Linux | [optional] [default to null]
**Route** | [**[]Route**](Route.md) | Entry associated with the routing table | [optional] [default to null]
**ArpTable** | [**[]ArpTable**](ArpTable.md) | Entry associated with the ARP table | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


