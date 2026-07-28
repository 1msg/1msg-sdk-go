# SendStickerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **int32** |  | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**MediaId** | Pointer to **string** | WABA media id from uploadMedia | [optional] 
**Link** | Pointer to **string** | Public URL to webp sticker | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 

## Methods

### NewSendStickerRequest

`func NewSendStickerRequest() *SendStickerRequest`

NewSendStickerRequest instantiates a new SendStickerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendStickerRequestWithDefaults

`func NewSendStickerRequestWithDefaults() *SendStickerRequest`

NewSendStickerRequestWithDefaults instantiates a new SendStickerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendStickerRequest) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendStickerRequest) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendStickerRequest) SetPhone(v int32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendStickerRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendStickerRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendStickerRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendStickerRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendStickerRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMediaId

`func (o *SendStickerRequest) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *SendStickerRequest) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *SendStickerRequest) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *SendStickerRequest) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetLink

`func (o *SendStickerRequest) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *SendStickerRequest) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *SendStickerRequest) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *SendStickerRequest) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendStickerRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendStickerRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendStickerRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendStickerRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


