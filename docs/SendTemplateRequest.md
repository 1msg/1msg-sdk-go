# SendTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **string** | Approved template name (template &#x60;name&#x60;, not the old HSM &#x60;element_name&#x60;) | [optional] 
**Language** | Pointer to [**SendTemplateRequestLanguage**](SendTemplateRequestLanguage.md) |  | [optional] 
**Namespace** | Pointer to **string** | Template namespace from &#x60;list_templates&#x60; | [optional] 
**Params** | Pointer to **[]map[string]interface{}** | Template components in WhatsApp Cloud API format. Each item is an object with &#x60;type&#x60; (&#x60;body&#x60;, &#x60;header&#x60;, &#x60;button&#x60;) and &#x60;parameters&#x60; (objects with &#x60;type&#x60; + &#x60;text&#x60; / media / action). Never a list of strings. Never HSM &#x60;localizable_params&#x60;.  | [optional] 
**Phone** | Pointer to **string** | Recipient phone with country code (use phone or chatId) | [optional] 
**ChatId** | Pointer to **string** | Recipient chat ID (&#x60;phone@c.us&#x60;, group &#x60;@g.us&#x60;, or BSUID &#x60;@lid&#x60;). Use phone or chatId. | [optional] 
**QuotedMsgId** | Pointer to **string** | Optional Cloud API wamid to quote | [optional] 
**UseMMlite** | Pointer to **bool** | Force Marketing Messages API (POST marketing_messages). If omitted, and the channel has mm_lite_enabled and mm_lite_available: MARKETING templates and templates with an unknown category use marketing_messages. An unknown category falls back to messages (mmLiteFallback) when Meta rejects it. Known AUTHENTICATION and UTILITY templates stay on messages. useMMlite:false forces messages.  | [optional] 
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

`func (o *SendTemplateRequest) GetParams() []map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SendTemplateRequest) GetParamsOk() (*[]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SendTemplateRequest) SetParams(v []map[string]interface{})`

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

### GetChatId

`func (o *SendTemplateRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendTemplateRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendTemplateRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendTemplateRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendTemplateRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendTemplateRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendTemplateRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendTemplateRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

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


