# Lbrp

## Properties
Name | Type | Description                                                                                                                                     | Notes
------------ | ------------- |-------------------------------------------------------------------------------------------------------------------------------------------------| -------------
**Name** | **string** | Name of the lbrp service                                                                                                                     | [optional] [default to null]
**Uuid** | **string** | UUID of the Cube                                                                                                                                | [optional] [default to null]
**Type_** | **string** | Type of the Cube (TC, XDP_SKB, XDP_DRV)                                                                                                         | [optional] [default to null]
**ServiceName** | **string** |                                                                                                                                                 | [optional] [default to null]
**Loglevel** | **string** | Defines the logging level of a service instance, from none (OFF) to the most verbose (TRACE)                                                    | [optional] [default to null]
**Ports** | [**[]Ports**](Ports.md) | Entry of the ports table                                                                                                                        | [optional] [default to null]
**Shadow** | **bool** | Defines if the service is visible in Linux                                                                                                      | [optional] [default to null]
**Span** | **bool** | Defines if all traffic is sent to Linux                                                                                                         | [optional] [default to null]
**SrcIpRewrite** | [***SrcIpRewrite**](SrcIpRewrite.md) | If configured, when a client request arrives to the LB, the source IP addrress is replaced with another IP address from the &#39;new&#39; range | [optional] [default to null]
**Service** | [**[]Service**](Service.md) | Services (i.e., virtual ip:protocol:port) exported to the client                                                                                | [optional] [default to null]
**PortMode_** | **string** | K8s lbrp mode of operation. 'MULTI' allows to manage multiple FRONTEND port. 'SINGLE' is optimized for working with a single FRONTEND port                                                    | [optional] [default to MULTI]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


