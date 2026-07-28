# MessageSentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **bool** | Whether the message was successfully queued for sending | [optional] 
**Id** | Pointer to **string** | Message ID | [optional] 
**Message** | Pointer to **string** | Status message | [optional] 
**Description** | Pointer to **string** | Detailed description | [optional] 

## Methods

### NewMessageSentResponse

`func NewMessageSentResponse() *MessageSentResponse`

NewMessageSentResponse instantiates a new MessageSentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSentResponseWithDefaults

`func NewMessageSentResponseWithDefaults() *MessageSentResponse`

NewMessageSentResponseWithDefaults instantiates a new MessageSentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *MessageSentResponse) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MessageSentResponse) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MessageSentResponse) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MessageSentResponse) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetId

`func (o *MessageSentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageSentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageSentResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageSentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *MessageSentResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageSentResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageSentResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MessageSentResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDescription

`func (o *MessageSentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageSentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageSentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageSentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


