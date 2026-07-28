# PatchFlowsFlowIdAssets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ValidationErrors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Error** | Pointer to **interface{}** | Error details returned by the provider. May be a string or a JSON object. | [optional] 

## Methods

### NewPatchFlowsFlowIdAssets200Response

`func NewPatchFlowsFlowIdAssets200Response() *PatchFlowsFlowIdAssets200Response`

NewPatchFlowsFlowIdAssets200Response instantiates a new PatchFlowsFlowIdAssets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFlowsFlowIdAssets200ResponseWithDefaults

`func NewPatchFlowsFlowIdAssets200ResponseWithDefaults() *PatchFlowsFlowIdAssets200Response`

NewPatchFlowsFlowIdAssets200ResponseWithDefaults instantiates a new PatchFlowsFlowIdAssets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PatchFlowsFlowIdAssets200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PatchFlowsFlowIdAssets200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PatchFlowsFlowIdAssets200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PatchFlowsFlowIdAssets200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetValidationErrors

`func (o *PatchFlowsFlowIdAssets200Response) GetValidationErrors() []map[string]interface{}`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PatchFlowsFlowIdAssets200Response) GetValidationErrorsOk() (*[]map[string]interface{}, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PatchFlowsFlowIdAssets200Response) SetValidationErrors(v []map[string]interface{})`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *PatchFlowsFlowIdAssets200Response) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetError

`func (o *PatchFlowsFlowIdAssets200Response) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PatchFlowsFlowIdAssets200Response) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PatchFlowsFlowIdAssets200Response) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *PatchFlowsFlowIdAssets200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *PatchFlowsFlowIdAssets200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *PatchFlowsFlowIdAssets200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


