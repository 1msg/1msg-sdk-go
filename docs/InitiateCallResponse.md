# InitiateCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingProduct** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Calls** | Pointer to [**[]InitiateCallResponseCallsInner**](InitiateCallResponseCallsInner.md) | Present on outbound &#x60;connect&#x60; — contains the new call id | [optional] 
**Result** | Pointer to **string** | Legacy success marker when upstream returns an empty body | [optional] 
**Response** | Pointer to [**InitiateCallResponseResponse**](InitiateCallResponseResponse.md) |  | [optional] 

## Methods

### NewInitiateCallResponse

`func NewInitiateCallResponse() *InitiateCallResponse`

NewInitiateCallResponse instantiates a new InitiateCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateCallResponseWithDefaults

`func NewInitiateCallResponseWithDefaults() *InitiateCallResponse`

NewInitiateCallResponseWithDefaults instantiates a new InitiateCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingProduct

`func (o *InitiateCallResponse) GetMessagingProduct() string`

GetMessagingProduct returns the MessagingProduct field if non-nil, zero value otherwise.

### GetMessagingProductOk

`func (o *InitiateCallResponse) GetMessagingProductOk() (*string, bool)`

GetMessagingProductOk returns a tuple with the MessagingProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProduct

`func (o *InitiateCallResponse) SetMessagingProduct(v string)`

SetMessagingProduct sets MessagingProduct field to given value.

### HasMessagingProduct

`func (o *InitiateCallResponse) HasMessagingProduct() bool`

HasMessagingProduct returns a boolean if a field has been set.

### GetSuccess

`func (o *InitiateCallResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InitiateCallResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InitiateCallResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *InitiateCallResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCalls

`func (o *InitiateCallResponse) GetCalls() []InitiateCallResponseCallsInner`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *InitiateCallResponse) GetCallsOk() (*[]InitiateCallResponseCallsInner, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *InitiateCallResponse) SetCalls(v []InitiateCallResponseCallsInner)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *InitiateCallResponse) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetResult

`func (o *InitiateCallResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InitiateCallResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InitiateCallResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *InitiateCallResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResponse

`func (o *InitiateCallResponse) GetResponse() InitiateCallResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *InitiateCallResponse) GetResponseOk() (*InitiateCallResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *InitiateCallResponse) SetResponse(v InitiateCallResponseResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *InitiateCallResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


