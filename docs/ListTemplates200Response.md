# ListTemplates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Templates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListTemplates200Response

`func NewListTemplates200Response() *ListTemplates200Response`

NewListTemplates200Response instantiates a new ListTemplates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTemplates200ResponseWithDefaults

`func NewListTemplates200ResponseWithDefaults() *ListTemplates200Response`

NewListTemplates200ResponseWithDefaults instantiates a new ListTemplates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListTemplates200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListTemplates200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListTemplates200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListTemplates200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTemplates

`func (o *ListTemplates200Response) GetTemplates() []map[string]interface{}`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *ListTemplates200Response) GetTemplatesOk() (*[]map[string]interface{}, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *ListTemplates200Response) SetTemplates(v []map[string]interface{})`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *ListTemplates200Response) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetError

`func (o *ListTemplates200Response) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListTemplates200Response) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListTemplates200Response) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ListTemplates200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


