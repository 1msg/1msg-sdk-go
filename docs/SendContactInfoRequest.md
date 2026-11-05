# SendContactInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Recipient phone with country code. Use phone, chatId, or bsuid. | [optional] 
**ChatId** | Pointer to **string** | Recipient chat ID (&#x60;phone@c.us&#x60;, BSUID &#x60;@lid&#x60;, or username &#x60;@username&#x60;). | [optional] 
**Bsuid** | Pointer to **string** | Business-scoped user id, for example &#x60;US.13491208655302741918&#x60;. | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 
**Body** | **string** | Text shown above the share-contact button. WhatsApp sets the button label; do not send button text.  | 

## Methods

### NewSendContactInfoRequest

`func NewSendContactInfoRequest(body string, ) *SendContactInfoRequest`

NewSendContactInfoRequest instantiates a new SendContactInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendContactInfoRequestWithDefaults

`func NewSendContactInfoRequestWithDefaults() *SendContactInfoRequest`

NewSendContactInfoRequestWithDefaults instantiates a new SendContactInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendContactInfoRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendContactInfoRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendContactInfoRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendContactInfoRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendContactInfoRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendContactInfoRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendContactInfoRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendContactInfoRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetBsuid

`func (o *SendContactInfoRequest) GetBsuid() string`

GetBsuid returns the Bsuid field if non-nil, zero value otherwise.

### GetBsuidOk

`func (o *SendContactInfoRequest) GetBsuidOk() (*string, bool)`

GetBsuidOk returns a tuple with the Bsuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsuid

`func (o *SendContactInfoRequest) SetBsuid(v string)`

SetBsuid sets Bsuid field to given value.

### HasBsuid

`func (o *SendContactInfoRequest) HasBsuid() bool`

HasBsuid returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendContactInfoRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendContactInfoRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendContactInfoRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendContactInfoRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

### GetBody

`func (o *SendContactInfoRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendContactInfoRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendContactInfoRequest) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


