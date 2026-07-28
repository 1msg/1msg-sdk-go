# CreateReadMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** |  | [optional] 
**MsgId** | Pointer to **string** |  | [optional] 
**TypingIndicator** | Pointer to **bool** | Show WhatsApp typing indicator (max 25s or until reply) | [optional] [default to false]

## Methods

### NewCreateReadMessageRequest

`func NewCreateReadMessageRequest() *CreateReadMessageRequest`

NewCreateReadMessageRequest instantiates a new CreateReadMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReadMessageRequestWithDefaults

`func NewCreateReadMessageRequestWithDefaults() *CreateReadMessageRequest`

NewCreateReadMessageRequestWithDefaults instantiates a new CreateReadMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *CreateReadMessageRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *CreateReadMessageRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *CreateReadMessageRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *CreateReadMessageRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMsgId

`func (o *CreateReadMessageRequest) GetMsgId() string`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *CreateReadMessageRequest) GetMsgIdOk() (*string, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *CreateReadMessageRequest) SetMsgId(v string)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *CreateReadMessageRequest) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetTypingIndicator

`func (o *CreateReadMessageRequest) GetTypingIndicator() bool`

GetTypingIndicator returns the TypingIndicator field if non-nil, zero value otherwise.

### GetTypingIndicatorOk

`func (o *CreateReadMessageRequest) GetTypingIndicatorOk() (*bool, bool)`

GetTypingIndicatorOk returns a tuple with the TypingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypingIndicator

`func (o *CreateReadMessageRequest) SetTypingIndicator(v bool)`

SetTypingIndicator sets TypingIndicator field to given value.

### HasTypingIndicator

`func (o *CreateReadMessageRequest) HasTypingIndicator() bool`

HasTypingIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


