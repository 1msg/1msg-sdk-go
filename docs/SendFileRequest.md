# SendFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | File source (required unless mediaId is set): - HTTP/HTTPS URL - Base64 data URI  | [optional] 
**Filename** | Pointer to **string** | File name with extension (required with body, not with mediaId) | [optional] 
**MediaId** | Pointer to **string** | WABA media id from uploadMedia (alternative to body) | [optional] 
**MediaType** | Pointer to **string** | Required when using mediaId | [optional] 
**Voice** | Pointer to **bool** | Native WhatsApp voice note (audio/ogg only) | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **int32** |  | [optional] 

## Methods

### NewSendFileRequest

`func NewSendFileRequest() *SendFileRequest`

NewSendFileRequest instantiates a new SendFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFileRequestWithDefaults

`func NewSendFileRequestWithDefaults() *SendFileRequest`

NewSendFileRequestWithDefaults instantiates a new SendFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *SendFileRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendFileRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendFileRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendFileRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFilename

`func (o *SendFileRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SendFileRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SendFileRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SendFileRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetMediaId

`func (o *SendFileRequest) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *SendFileRequest) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *SendFileRequest) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *SendFileRequest) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetMediaType

`func (o *SendFileRequest) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *SendFileRequest) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *SendFileRequest) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *SendFileRequest) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetVoice

`func (o *SendFileRequest) GetVoice() bool`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *SendFileRequest) GetVoiceOk() (*bool, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *SendFileRequest) SetVoice(v bool)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *SendFileRequest) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetCaption

`func (o *SendFileRequest) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *SendFileRequest) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *SendFileRequest) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *SendFileRequest) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendFileRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendFileRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendFileRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendFileRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

### GetChatId

`func (o *SendFileRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendFileRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendFileRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendFileRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetPhone

`func (o *SendFileRequest) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendFileRequest) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendFileRequest) SetPhone(v int32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendFileRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


