# ListFlows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Paging** | Pointer to **map[string]interface{}** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **interface{}** | Error details returned by the provider. May be a string or a JSON object. | [optional] 

## Methods

### NewListFlows200Response

`func NewListFlows200Response() *ListFlows200Response`

NewListFlows200Response instantiates a new ListFlows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlows200ResponseWithDefaults

`func NewListFlows200ResponseWithDefaults() *ListFlows200Response`

NewListFlows200ResponseWithDefaults instantiates a new ListFlows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListFlows200Response) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListFlows200Response) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListFlows200Response) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *ListFlows200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPaging

`func (o *ListFlows200Response) GetPaging() map[string]interface{}`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ListFlows200Response) GetPagingOk() (*map[string]interface{}, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ListFlows200Response) SetPaging(v map[string]interface{})`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ListFlows200Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetCount

`func (o *ListFlows200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListFlows200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListFlows200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListFlows200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *ListFlows200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListFlows200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListFlows200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListFlows200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetError

`func (o *ListFlows200Response) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListFlows200Response) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListFlows200Response) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ListFlows200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ListFlows200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListFlows200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


