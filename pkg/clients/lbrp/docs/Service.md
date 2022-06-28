# Service

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service name related to the backend server of the pool is connected to | [optional] [default to null]
**Vip** | **string** | Virtual IP (vip) of the service where clients connect to | [optional] [default to null]
**Vport** | **int32** | Port of the virtual server where clients connect to (this value is ignored in case of ICMP) | [optional] [default to null]
**Proto** | **string** | Upper-layer protocol associated with a loadbalancing service instance. &#39;ALL&#39; creates an entry for all the supported protocols | [optional] [default to null]
**Backend** | [**[]ServiceBackend**](ServiceBackend.md) | Pool of backend servers that actually serve requests | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


