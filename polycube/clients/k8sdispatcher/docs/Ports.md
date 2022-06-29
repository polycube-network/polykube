# Ports

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Port Name | [optional] [default to null]
**Uuid** | **string** | UUID of the port | [optional] [default to null]
**Status** | **string** | Status of the port (UP or DOWN) | [optional] [default to null]
**Peer** | **string** | Peer name, such as a network interfaces (e.g., &#39;veth0&#39;) or another cube (e.g., &#39;br1:port2&#39;) | [optional] [default to null]
**Tcubes** | **[]string** |  | [optional] [default to null]
**Type_** | **string** | Type of the K8s Dispatcher cube port (e.g. BACKEND or FRONTEND) | [optional] [default to null]
**Ip** | **string** | IP address of the node interface (only for FRONTEND port) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


