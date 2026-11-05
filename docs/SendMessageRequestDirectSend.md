# SendMessageRequestDirectSend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceFirst** | Pointer to **bool** | When true, send a service message first. Success stops there. A closed-window error (131047) retries as Direct Send. Any other error is returned and Direct Send is not attempted. Omit or false to send Direct Send immediately, even inside an open window.  | [optional] 
**TemplateName** | Pointer to **string** | Utility only. Meta creates or reuses a Direct Send template with this exact name (&#x60;^[a-z0-9_]+$&#x60;, max 512). Do not reuse a template created by hand.  | [optional] 

## Methods

### NewSendMessageRequestDirectSend

`func NewSendMessageRequestDirectSend() *SendMessageRequestDirectSend`

NewSendMessageRequestDirectSend instantiates a new SendMessageRequestDirectSend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestDirectSendWithDefaults

`func NewSendMessageRequestDirectSendWithDefaults() *SendMessageRequestDirectSend`

NewSendMessageRequestDirectSendWithDefaults instantiates a new SendMessageRequestDirectSend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceFirst

`func (o *SendMessageRequestDirectSend) GetServiceFirst() bool`

GetServiceFirst returns the ServiceFirst field if non-nil, zero value otherwise.

### GetServiceFirstOk

`func (o *SendMessageRequestDirectSend) GetServiceFirstOk() (*bool, bool)`

GetServiceFirstOk returns a tuple with the ServiceFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFirst

`func (o *SendMessageRequestDirectSend) SetServiceFirst(v bool)`

SetServiceFirst sets ServiceFirst field to given value.

### HasServiceFirst

`func (o *SendMessageRequestDirectSend) HasServiceFirst() bool`

HasServiceFirst returns a boolean if a field has been set.

### GetTemplateName

`func (o *SendMessageRequestDirectSend) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *SendMessageRequestDirectSend) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *SendMessageRequestDirectSend) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *SendMessageRequestDirectSend) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


