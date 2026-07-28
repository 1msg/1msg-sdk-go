# DeleteFlowsFlowId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **interface{}** | Error details returned by the provider. May be a string or a JSON object. | [optional] 

## Methods

### NewDeleteFlowsFlowId200Response

`func NewDeleteFlowsFlowId200Response() *DeleteFlowsFlowId200Response`

NewDeleteFlowsFlowId200Response instantiates a new DeleteFlowsFlowId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFlowsFlowId200ResponseWithDefaults

`func NewDeleteFlowsFlowId200ResponseWithDefaults() *DeleteFlowsFlowId200Response`

NewDeleteFlowsFlowId200ResponseWithDefaults instantiates a new DeleteFlowsFlowId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeleteFlowsFlowId200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeleteFlowsFlowId200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeleteFlowsFlowId200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DeleteFlowsFlowId200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *DeleteFlowsFlowId200Response) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteFlowsFlowId200Response) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteFlowsFlowId200Response) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *DeleteFlowsFlowId200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DeleteFlowsFlowId200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DeleteFlowsFlowId200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


