# SendCtaUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **int32** |  | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**Body** | **string** |  | 
**DisplayText** | **string** | Button label | 
**Url** | **string** |  | 
**Header** | Pointer to **map[string]interface{}** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 

## Methods

### NewSendCtaUrlRequest

`func NewSendCtaUrlRequest(body string, displayText string, url string, ) *SendCtaUrlRequest`

NewSendCtaUrlRequest instantiates a new SendCtaUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendCtaUrlRequestWithDefaults

`func NewSendCtaUrlRequestWithDefaults() *SendCtaUrlRequest`

NewSendCtaUrlRequestWithDefaults instantiates a new SendCtaUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendCtaUrlRequest) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendCtaUrlRequest) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendCtaUrlRequest) SetPhone(v int32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendCtaUrlRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendCtaUrlRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendCtaUrlRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendCtaUrlRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendCtaUrlRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetBody

`func (o *SendCtaUrlRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendCtaUrlRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendCtaUrlRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetDisplayText

`func (o *SendCtaUrlRequest) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *SendCtaUrlRequest) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *SendCtaUrlRequest) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.


### GetUrl

`func (o *SendCtaUrlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SendCtaUrlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SendCtaUrlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeader

`func (o *SendCtaUrlRequest) GetHeader() map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SendCtaUrlRequest) GetHeaderOk() (*map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SendCtaUrlRequest) SetHeader(v map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SendCtaUrlRequest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetFooter

`func (o *SendCtaUrlRequest) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *SendCtaUrlRequest) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *SendCtaUrlRequest) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *SendCtaUrlRequest) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendCtaUrlRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendCtaUrlRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendCtaUrlRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendCtaUrlRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


