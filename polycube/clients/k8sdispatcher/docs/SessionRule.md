# SessionRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | Session entry direction (e.g. INGRESS or EGRESS) | [optional] [default to null]
**SrcIp** | **string** | Session entry source IP address | [optional] [default to null]
**DstIp** | **string** | Session entry destination IP address | [optional] [default to null]
**SrcPort** | **int32** | Session entry source L4 port number | [optional] [default to null]
**DstPort** | **int32** | Session entry destination L4 port number | [optional] [default to null]
**Proto** | **string** | Session entry L4 protocol | [optional] [default to null]
**NewIp** | **string** | Translated IP address | [optional] [default to null]
**NewPort** | **int32** | Translated L4 port number | [optional] [default to null]
**Operation** | **string** | Operation applied on the original packet | [optional] [default to null]
**OriginatingRule** | **string** | Rule originating the session entry | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


