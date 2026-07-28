# SendAddressMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **int32** | Recipient phone (E.164 digits, no +). Must match country. | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**Body** | **string** | Body text shown with the address request | 
**Country** | Pointer to **string** | Address form country. Defaults to IN if omitted. | [optional] 
**Values** | Pointer to **map[string]interface{}** | Optional prefilled address fields | [optional] 
**SavedAddresses** | Pointer to **[]map[string]interface{}** | Optional previously saved addresses for the user | [optional] 
**ValidationErrors** | Pointer to **map[string]interface{}** | Optional field validation errors when re-prompting | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 

## Methods

### NewSendAddressMessageRequest

`func NewSendAddressMessageRequest(body string, ) *SendAddressMessageRequest`

NewSendAddressMessageRequest instantiates a new SendAddressMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendAddressMessageRequestWithDefaults

`func NewSendAddressMessageRequestWithDefaults() *SendAddressMessageRequest`

NewSendAddressMessageRequestWithDefaults instantiates a new SendAddressMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendAddressMessageRequest) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendAddressMessageRequest) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendAddressMessageRequest) SetPhone(v int32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendAddressMessageRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetChatId

`func (o *SendAddressMessageRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendAddressMessageRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendAddressMessageRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendAddressMessageRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetBody

`func (o *SendAddressMessageRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendAddressMessageRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendAddressMessageRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetCountry

`func (o *SendAddressMessageRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SendAddressMessageRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SendAddressMessageRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SendAddressMessageRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetValues

`func (o *SendAddressMessageRequest) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SendAddressMessageRequest) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SendAddressMessageRequest) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *SendAddressMessageRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetSavedAddresses

`func (o *SendAddressMessageRequest) GetSavedAddresses() []map[string]interface{}`

GetSavedAddresses returns the SavedAddresses field if non-nil, zero value otherwise.

### GetSavedAddressesOk

`func (o *SendAddressMessageRequest) GetSavedAddressesOk() (*[]map[string]interface{}, bool)`

GetSavedAddressesOk returns a tuple with the SavedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedAddresses

`func (o *SendAddressMessageRequest) SetSavedAddresses(v []map[string]interface{})`

SetSavedAddresses sets SavedAddresses field to given value.

### HasSavedAddresses

`func (o *SendAddressMessageRequest) HasSavedAddresses() bool`

HasSavedAddresses returns a boolean if a field has been set.

### GetValidationErrors

`func (o *SendAddressMessageRequest) GetValidationErrors() map[string]interface{}`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *SendAddressMessageRequest) GetValidationErrorsOk() (*map[string]interface{}, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *SendAddressMessageRequest) SetValidationErrors(v map[string]interface{})`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *SendAddressMessageRequest) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendAddressMessageRequest) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendAddressMessageRequest) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendAddressMessageRequest) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendAddressMessageRequest) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


