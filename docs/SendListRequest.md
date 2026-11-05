# SendListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Recipient phone with country code. Use phone or chatId. | [optional] 
**ChatId** | Pointer to **string** | Recipient chat ID (&#x60;phone@c.us&#x60;, group &#x60;@g.us&#x60;, or BSUID &#x60;@lid&#x60;). Use phone or chatId. | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**ButtonText** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Sections** | Pointer to [**[]SendListRequestSectionsInner**](SendListRequestSectionsInner.md) |  | [optional] 

## Methods

### NewSendListRequest

`func NewSendListRequest() *SendListRequest`

NewSendListRequest instantiates a new SendListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendListRequestWithDefaults

`func NewSendListRequestWithDefaults() *SendListRequest`

NewSendListRequestWithDefaults instantiates a new SendListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendListRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendListRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendListRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendListRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendListRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendListRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendListRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendListRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendListRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendListRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendListRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendListRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

### GetBody

`func (o *SendListRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendListRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendListRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendListRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetButtonText

`func (o *SendListRequest) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *SendListRequest) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *SendListRequest) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *SendListRequest) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### GetAction

`func (o *SendListRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SendListRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SendListRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SendListRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSections

`func (o *SendListRequest) GetSections() []SendListRequestSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SendListRequest) GetSectionsOk() (*[]SendListRequestSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SendListRequest) SetSections(v []SendListRequestSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *SendListRequest) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


