# SendOrderDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **int32** | Recipient phone (India E.164 digits, no +). Use phone or chatId. | [optional] 
**ChatId** | Pointer to **string** | Recipient chatId (e.g. phone@c.us). Use phone or chatId. | [optional] 
**Template** | **string** | Approved Utility template name that includes an ORDER_DETAILS button | 
**Namespace** | **string** | Template namespace from the channel / template list | 
**Language** | [**SendOrderDetailsRequestLanguage**](SendOrderDetailsRequestLanguage.md) |  | 
**Params** | Pointer to **[]map[string]interface{}** | Extra template components (HEADER / BODY / etc.). If an order_details button is missing, the API appends one from order / referenceId / currency / paymentSettings.  | [optional] 
**ReferenceId** | Pointer to **string** | Unique order / payment reference id (maps to reference_id) | [optional] 
**Currency** | Pointer to **string** | Currency code for India payments | [optional] 
**PaymentSettings** | Pointer to **map[string]interface{}** | Optional payment settings (UPI / payment gateway / payment link). Forwarded as payment_settings on the order_details action.  | [optional] 
**Order** | [**SendOrderDetailsRequestOrder**](SendOrderDetailsRequestOrder.md) |  | 

## Methods

### NewSendOrderDetailsRequest

`func NewSendOrderDetailsRequest(template string, namespace string, language SendOrderDetailsRequestLanguage, order SendOrderDetailsRequestOrder, ) *SendOrderDetailsRequest`

NewSendOrderDetailsRequest instantiates a new SendOrderDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendOrderDetailsRequestWithDefaults

`func NewSendOrderDetailsRequestWithDefaults() *SendOrderDetailsRequest`

NewSendOrderDetailsRequestWithDefaults instantiates a new SendOrderDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendOrderDetailsRequest) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendOrderDetailsRequest) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendOrderDetailsRequest) SetPhone(v int32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendOrderDetailsRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendOrderDetailsRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendOrderDetailsRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendOrderDetailsRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendOrderDetailsRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetTemplate

`func (o *SendOrderDetailsRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SendOrderDetailsRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SendOrderDetailsRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetNamespace

`func (o *SendOrderDetailsRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SendOrderDetailsRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SendOrderDetailsRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetLanguage

`func (o *SendOrderDetailsRequest) GetLanguage() SendOrderDetailsRequestLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SendOrderDetailsRequest) GetLanguageOk() (*SendOrderDetailsRequestLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SendOrderDetailsRequest) SetLanguage(v SendOrderDetailsRequestLanguage)`

SetLanguage sets Language field to given value.


### GetParams

`func (o *SendOrderDetailsRequest) GetParams() []map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SendOrderDetailsRequest) GetParamsOk() (*[]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SendOrderDetailsRequest) SetParams(v []map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *SendOrderDetailsRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetReferenceId

`func (o *SendOrderDetailsRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SendOrderDetailsRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SendOrderDetailsRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SendOrderDetailsRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetCurrency

`func (o *SendOrderDetailsRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SendOrderDetailsRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SendOrderDetailsRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SendOrderDetailsRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentSettings

`func (o *SendOrderDetailsRequest) GetPaymentSettings() map[string]interface{}`

GetPaymentSettings returns the PaymentSettings field if non-nil, zero value otherwise.

### GetPaymentSettingsOk

`func (o *SendOrderDetailsRequest) GetPaymentSettingsOk() (*map[string]interface{}, bool)`

GetPaymentSettingsOk returns a tuple with the PaymentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSettings

`func (o *SendOrderDetailsRequest) SetPaymentSettings(v map[string]interface{})`

SetPaymentSettings sets PaymentSettings field to given value.

### HasPaymentSettings

`func (o *SendOrderDetailsRequest) HasPaymentSettings() bool`

HasPaymentSettings returns a boolean if a field has been set.

### GetOrder

`func (o *SendOrderDetailsRequest) GetOrder() SendOrderDetailsRequestOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SendOrderDetailsRequest) GetOrderOk() (*SendOrderDetailsRequestOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SendOrderDetailsRequest) SetOrder(v SendOrderDetailsRequestOrder)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


