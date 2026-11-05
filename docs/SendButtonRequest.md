# SendButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Recipient phone with country code. Use phone or chatId. | [optional] 
**ChatId** | Pointer to **string** | Recipient chat ID (&#x60;phone@c.us&#x60;, group &#x60;@g.us&#x60;, or BSUID &#x60;@lid&#x60;). Use phone or chatId. | [optional] 
**QuotedMsgId** | Pointer to **string** | Optional Cloud API wamid to quote | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **map[string]interface{}** | Optional interactive header (text/image/video/document) | [optional] 
**Buttons** | Pointer to **[]map[string]interface{}** | Legacy alias for sections (&#x60;[{id,title}]&#x60;) | [optional] 
**Sections** | Pointer to [**[]SendButtonRequestSectionsInner**](SendButtonRequestSectionsInner.md) |  | [optional] 

## Methods

### NewSendButtonRequest

`func NewSendButtonRequest() *SendButtonRequest`

NewSendButtonRequest instantiates a new SendButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendButtonRequestWithDefaults

`func NewSendButtonRequestWithDefaults() *SendButtonRequest`

NewSendButtonRequestWithDefaults instantiates a new SendButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendButtonRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendButtonRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendButtonRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendButtonRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendButtonRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendButtonRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendButtonRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendButtonRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendButtonRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendButtonRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendButtonRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendButtonRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

### GetBody

`func (o *SendButtonRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendButtonRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendButtonRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendButtonRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFooter

`func (o *SendButtonRequest) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *SendButtonRequest) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *SendButtonRequest) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *SendButtonRequest) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeader

`func (o *SendButtonRequest) GetHeader() map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SendButtonRequest) GetHeaderOk() (*map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SendButtonRequest) SetHeader(v map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SendButtonRequest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetButtons

`func (o *SendButtonRequest) GetButtons() []map[string]interface{}`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *SendButtonRequest) GetButtonsOk() (*[]map[string]interface{}, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *SendButtonRequest) SetButtons(v []map[string]interface{})`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *SendButtonRequest) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetSections

`func (o *SendButtonRequest) GetSections() []SendButtonRequestSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SendButtonRequest) GetSectionsOk() (*[]SendButtonRequestSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SendButtonRequest) SetSections(v []SendButtonRequestSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *SendButtonRequest) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


