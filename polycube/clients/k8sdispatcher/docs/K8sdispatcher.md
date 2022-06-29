# K8sdispatcher

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the k8sdispatcher service | [optional] [default to null]
**Uuid** | **string** | UUID of the Cube | [optional] [default to null]
**Type_** | **string** | Type of the Cube (TC, XDP_SKB, XDP_DRV) | [optional] [default to null]
**ServiceName** | **string** |  | [optional] [default to null]
**Loglevel** | **string** | Logging level of a cube, from none (OFF) to the most verbose (TRACE) | [optional] [default to null]
**Ports** | [**[]Ports**](Ports.md) | Entry of the ports table | [optional] [default to null]
**Shadow** | **bool** | Defines if the service is visible in Linux | [optional] [default to null]
**Span** | **bool** | Defines if all traffic is sent to Linux | [optional] [default to null]
**InternalSrcIp** | **string** | Internal source IP address used for natting incoming packets directed to Kubernetes Services with a CLUSTER external traffic policy | [optional] [default to null]
**NodeportRange** | **string** | Port range used for NodePort Services | [optional] [default to null]
**SessionRule** | [**[]SessionRule**](SessionRule.md) | Session entry related to a specific traffic direction | [optional] [default to null]
**NodeportRule** | [**[]NodeportRule**](NodeportRule.md) | NodePort rule associated with a Kubernetes NodePort Service | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


