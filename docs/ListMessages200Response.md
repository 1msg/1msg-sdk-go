# ListMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notice** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListMessages200Response

`func NewListMessages200Response() *ListMessages200Response`

NewListMessages200Response instantiates a new ListMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessages200ResponseWithDefaults

`func NewListMessages200ResponseWithDefaults() *ListMessages200Response`

NewListMessages200ResponseWithDefaults instantiates a new ListMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotice

`func (o *ListMessages200Response) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *ListMessages200Response) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *ListMessages200Response) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *ListMessages200Response) HasNotice() bool`

HasNotice returns a boolean if a field has been set.

### GetMessages

`func (o *ListMessages200Response) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListMessages200Response) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListMessages200Response) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ListMessages200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


