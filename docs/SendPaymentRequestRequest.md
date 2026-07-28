# SendPaymentRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **int32** |  | 
**Region** | **string** |  | 
**Body** | Pointer to **string** |  | [optional] 
**Interactive** | Pointer to **map[string]interface{}** |  | [optional] 
**Action** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSendPaymentRequestRequest

`func NewSendPaymentRequestRequest(phone int32, region string, ) *SendPaymentRequestRequest`

NewSendPaymentRequestRequest instantiates a new SendPaymentRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPaymentRequestRequestWithDefaults

`func NewSendPaymentRequestRequestWithDefaults() *SendPaymentRequestRequest`

NewSendPaymentRequestRequestWithDefaults instantiates a new SendPaymentRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendPaymentRequestRequest) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendPaymentRequestRequest) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendPaymentRequestRequest) SetPhone(v int32)`

SetPhone sets Phone field to given value.


### GetRegion

`func (o *SendPaymentRequestRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SendPaymentRequestRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SendPaymentRequestRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetBody

`func (o *SendPaymentRequestRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendPaymentRequestRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendPaymentRequestRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendPaymentRequestRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetInteractive

`func (o *SendPaymentRequestRequest) GetInteractive() map[string]interface{}`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *SendPaymentRequestRequest) GetInteractiveOk() (*map[string]interface{}, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *SendPaymentRequestRequest) SetInteractive(v map[string]interface{})`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *SendPaymentRequestRequest) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetAction

`func (o *SendPaymentRequestRequest) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SendPaymentRequestRequest) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SendPaymentRequestRequest) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *SendPaymentRequestRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


