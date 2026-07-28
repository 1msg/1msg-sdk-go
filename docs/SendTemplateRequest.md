# SendTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **string** |  | [optional] 
**Language** | Pointer to [**SendTemplateRequestLanguage**](SendTemplateRequestLanguage.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **[]interface{}** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**UseMMlite** | Pointer to **bool** | Force Marketing Messages API (POST marketing_messages). If omitted: auto for MARKETING category when channel has mm_lite_enabled and mm_lite_available service settings.  | [optional] 
**MessageActivitySharing** | Pointer to **bool** | Sets message_activity_sharing on WABA payload. Enables click tracking webhooks. Requires MM Lite path. Ignored on Cloud API fallback.  | [optional] 
**MessageSendTtlSeconds** | Pointer to **int32** | Template message TTL in seconds (message_send_ttl_seconds). MARKETING via MM Lite: 43200–2592000. AUTHENTICATION: 30–900. UTILITY: 30–43200. AUTH/UTILITY also accept -1 (30-day custom TTL).  | [optional] 

## Methods

### NewSendTemplateRequest

`func NewSendTemplateRequest() *SendTemplateRequest`

NewSendTemplateRequest instantiates a new SendTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTemplateRequestWithDefaults

`func NewSendTemplateRequestWithDefaults() *SendTemplateRequest`

NewSendTemplateRequestWithDefaults instantiates a new SendTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *SendTemplateRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SendTemplateRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SendTemplateRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SendTemplateRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetLanguage

`func (o *SendTemplateRequest) GetLanguage() SendTemplateRequestLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SendTemplateRequest) GetLanguageOk() (*SendTemplateRequestLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SendTemplateRequest) SetLanguage(v SendTemplateRequestLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SendTemplateRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetNamespace

`func (o *SendTemplateRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SendTemplateRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SendTemplateRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SendTemplateRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParams

`func (o *SendTemplateRequest) GetParams() []interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SendTemplateRequest) GetParamsOk() (*[]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SendTemplateRequest) SetParams(v []interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *SendTemplateRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetPhone

`func (o *SendTemplateRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendTemplateRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendTemplateRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendTemplateRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUseMMlite

`func (o *SendTemplateRequest) GetUseMMlite() bool`

GetUseMMlite returns the UseMMlite field if non-nil, zero value otherwise.

### GetUseMMliteOk

`func (o *SendTemplateRequest) GetUseMMliteOk() (*bool, bool)`

GetUseMMliteOk returns a tuple with the UseMMlite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMMlite

`func (o *SendTemplateRequest) SetUseMMlite(v bool)`

SetUseMMlite sets UseMMlite field to given value.

### HasUseMMlite

`func (o *SendTemplateRequest) HasUseMMlite() bool`

HasUseMMlite returns a boolean if a field has been set.

### GetMessageActivitySharing

`func (o *SendTemplateRequest) GetMessageActivitySharing() bool`

GetMessageActivitySharing returns the MessageActivitySharing field if non-nil, zero value otherwise.

### GetMessageActivitySharingOk

`func (o *SendTemplateRequest) GetMessageActivitySharingOk() (*bool, bool)`

GetMessageActivitySharingOk returns a tuple with the MessageActivitySharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageActivitySharing

`func (o *SendTemplateRequest) SetMessageActivitySharing(v bool)`

SetMessageActivitySharing sets MessageActivitySharing field to given value.

### HasMessageActivitySharing

`func (o *SendTemplateRequest) HasMessageActivitySharing() bool`

HasMessageActivitySharing returns a boolean if a field has been set.

### GetMessageSendTtlSeconds

`func (o *SendTemplateRequest) GetMessageSendTtlSeconds() int32`

GetMessageSendTtlSeconds returns the MessageSendTtlSeconds field if non-nil, zero value otherwise.

### GetMessageSendTtlSecondsOk

`func (o *SendTemplateRequest) GetMessageSendTtlSecondsOk() (*int32, bool)`

GetMessageSendTtlSecondsOk returns a tuple with the MessageSendTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSendTtlSeconds

`func (o *SendTemplateRequest) SetMessageSendTtlSeconds(v int32)`

SetMessageSendTtlSeconds sets MessageSendTtlSeconds field to given value.

### HasMessageSendTtlSeconds

`func (o *SendTemplateRequest) HasMessageSendTtlSeconds() bool`

HasMessageSendTtlSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


