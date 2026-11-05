# SendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | Message text, UTF-8 or UTF-16 string with emoji | 
**QuotedMsgId** | Pointer to **string** | Message ID to quote/reply to (Cloud API wamid) | [optional] 
**ChatId** | Pointer to **string** | Recipient chat ID: phone@c.us, group@g.us, or BSUID@lid | [optional] 
**Phone** | Pointer to **int32** | Phone number starting with country code (alternative to chatId) | [optional] 
**Category** | Pointer to **string** | Optional Direct Send category. Direct Send is in beta: Meta has not enabled it on every business portfolio. &#x60;utility&#x60; or &#x60;authentication&#x60; enables it. &#x60;authentication&#x60; is part of that beta and is not recommended. Omit or &#x60;service&#x60; keeps a normal service message (and the marketing auto-template outside the 24-hour window). &#x60;marketing&#x60; is rejected.  | [optional] 
**DirectSend** | Pointer to [**SendMessageRequestDirectSend**](SendMessageRequestDirectSend.md) |  | [optional] 

## Methods

### NewSendMessageRequest

`func NewSendMessageRequest(body string, ) *SendMessageRequest`

NewSendMessageRequest instantiates a new SendMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestWithDefaults

`func NewSendMessageRequestWithDefaults() *SendMessageRequest`

NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *SendMessageRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendMessageRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendMessageRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetQuotedMsgId

`func (o *SendMessageRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendMessageRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendMessageRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendMessageRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

### GetChatId

`func (o *SendMessageRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendMessageRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendMessageRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendMessageRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetPhone

`func (o *SendMessageRequest) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendMessageRequest) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendMessageRequest) SetPhone(v int32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendMessageRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCategory

`func (o *SendMessageRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SendMessageRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SendMessageRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SendMessageRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDirectSend

`func (o *SendMessageRequest) GetDirectSend() SendMessageRequestDirectSend`

GetDirectSend returns the DirectSend field if non-nil, zero value otherwise.

### GetDirectSendOk

`func (o *SendMessageRequest) GetDirectSendOk() (*SendMessageRequestDirectSend, bool)`

GetDirectSendOk returns a tuple with the DirectSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSend

`func (o *SendMessageRequest) SetDirectSend(v SendMessageRequestDirectSend)`

SetDirectSend sets DirectSend field to given value.

### HasDirectSend

`func (o *SendMessageRequest) HasDirectSend() bool`

HasDirectSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


