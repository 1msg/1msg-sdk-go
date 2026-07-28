# SendReactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 

## Methods

### NewSendReactionRequest

`func NewSendReactionRequest() *SendReactionRequest`

NewSendReactionRequest instantiates a new SendReactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendReactionRequestWithDefaults

`func NewSendReactionRequestWithDefaults() *SendReactionRequest`

NewSendReactionRequestWithDefaults instantiates a new SendReactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendReactionRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendReactionRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendReactionRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendReactionRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetBody

`func (o *SendReactionRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendReactionRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendReactionRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendReactionRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendReactionRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendReactionRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendReactionRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendReactionRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


