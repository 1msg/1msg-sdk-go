# SendProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Recipient phone with country code. Use phone or chatId. | [optional] 
**ChatId** | Pointer to **string** | Recipient chat ID (&#x60;phone@c.us&#x60;, group &#x60;@g.us&#x60;, or BSUID &#x60;@lid&#x60;). Use phone or chatId. | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** | Optional message text shown with the product | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** | Text header (product list only) | [optional] 
**Action** | **map[string]interface{}** | Catalog action (&#x60;catalog_id&#x60; + &#x60;product_retailer_id&#x60;, or &#x60;sections&#x60; for a list) | 

## Methods

### NewSendProductRequest

`func NewSendProductRequest(action map[string]interface{}, ) *SendProductRequest`

NewSendProductRequest instantiates a new SendProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendProductRequestWithDefaults

`func NewSendProductRequestWithDefaults() *SendProductRequest`

NewSendProductRequestWithDefaults instantiates a new SendProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendProductRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendProductRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendProductRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendProductRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendProductRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendProductRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendProductRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendProductRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendProductRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendProductRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendProductRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendProductRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

### GetBody

`func (o *SendProductRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendProductRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendProductRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendProductRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFooter

`func (o *SendProductRequest) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *SendProductRequest) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *SendProductRequest) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *SendProductRequest) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeader

`func (o *SendProductRequest) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SendProductRequest) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SendProductRequest) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SendProductRequest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetAction

`func (o *SendProductRequest) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SendProductRequest) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SendProductRequest) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


