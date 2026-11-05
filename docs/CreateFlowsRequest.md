# CreateFlowsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Flow name | 
**Categories** | **[]string** | One or more Flow categories | 
**EndpointUri** | Pointer to **string** | HTTPS endpoint for the Flow data channel | [optional] 
**Publish** | Pointer to **bool** | Publish immediately after create | [optional] 
**FlowJson** | Pointer to **map[string]interface{}** | Optional Flow JSON document | [optional] 

## Methods

### NewCreateFlowsRequest

`func NewCreateFlowsRequest(name string, categories []string, ) *CreateFlowsRequest`

NewCreateFlowsRequest instantiates a new CreateFlowsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlowsRequestWithDefaults

`func NewCreateFlowsRequestWithDefaults() *CreateFlowsRequest`

NewCreateFlowsRequestWithDefaults instantiates a new CreateFlowsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFlowsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFlowsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFlowsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCategories

`func (o *CreateFlowsRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CreateFlowsRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CreateFlowsRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetEndpointUri

`func (o *CreateFlowsRequest) GetEndpointUri() string`

GetEndpointUri returns the EndpointUri field if non-nil, zero value otherwise.

### GetEndpointUriOk

`func (o *CreateFlowsRequest) GetEndpointUriOk() (*string, bool)`

GetEndpointUriOk returns a tuple with the EndpointUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUri

`func (o *CreateFlowsRequest) SetEndpointUri(v string)`

SetEndpointUri sets EndpointUri field to given value.

### HasEndpointUri

`func (o *CreateFlowsRequest) HasEndpointUri() bool`

HasEndpointUri returns a boolean if a field has been set.

### GetPublish

`func (o *CreateFlowsRequest) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *CreateFlowsRequest) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *CreateFlowsRequest) SetPublish(v bool)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *CreateFlowsRequest) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetFlowJson

`func (o *CreateFlowsRequest) GetFlowJson() map[string]interface{}`

GetFlowJson returns the FlowJson field if non-nil, zero value otherwise.

### GetFlowJsonOk

`func (o *CreateFlowsRequest) GetFlowJsonOk() (*map[string]interface{}, bool)`

GetFlowJsonOk returns a tuple with the FlowJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowJson

`func (o *CreateFlowsRequest) SetFlowJson(v map[string]interface{})`

SetFlowJson sets FlowJson field to given value.

### HasFlowJson

`func (o *CreateFlowsRequest) HasFlowJson() bool`

HasFlowJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


