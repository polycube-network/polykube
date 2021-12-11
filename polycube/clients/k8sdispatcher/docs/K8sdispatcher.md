# K8sdispatcher

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the k8sdispatcher service | [optional] [default to null]
**Uuid** | **string** | UUID of the Cube | [optional] [default to null]
**Type_** | **string** | Type of the Cube (TC, XDP_SKB, XDP_DRV) | [optional] [default to null]
**ServiceName** | **string** |  | [optional] [default to null]
**Loglevel** | **string** | Defines the logging level of a service instance, from none (OFF) to the most verbose (TRACE) | [optional] [default to null]
**Ports** | [**[]Ports**](Ports.md) | Entry of the ports table | [optional] [default to null]
**Shadow** | **bool** | Defines if the service is visible in Linux | [optional] [default to null]
**Span** | **bool** | Defines if all traffic is sent to Linux | [optional] [default to null]
**ClusterIpSubnet** | **string** | Range of VIPs where clusterIP services are exposed | [optional] [default to null]
**ClientSubnet** | **string** | Range of IPs of pods in this node | [optional] [default to null]
**InternalSrcIp** | **string** | Internal src ip used for services with externaltrafficpolicy&#x3D;cluster | [optional] [default to null]
**NattingRule** | [**[]NattingRule**](NattingRule.md) |  | [optional] [default to null]
**NodeportRule** | [**[]NodeportRule**](NodeportRule.md) |  | [optional] [default to null]
**NodeportRange** | **string** | Port range used for NodePort services | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


