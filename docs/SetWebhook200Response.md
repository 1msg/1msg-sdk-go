# SetWebhook200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **bool** | True when the stored webhook URL(s) match the request | [optional] 
**WebhookUrl** | Pointer to [**SetWebhook200ResponseWebhookUrl**](SetWebhook200ResponseWebhookUrl.md) |  | [optional] 

## Methods

### NewSetWebhook200Response

`func NewSetWebhook200Response() *SetWebhook200Response`

NewSetWebhook200Response instantiates a new SetWebhook200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWebhook200ResponseWithDefaults

`func NewSetWebhook200ResponseWithDefaults() *SetWebhook200Response`

NewSetWebhook200ResponseWithDefaults instantiates a new SetWebhook200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *SetWebhook200Response) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *SetWebhook200Response) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *SetWebhook200Response) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *SetWebhook200Response) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *SetWebhook200Response) GetWebhookUrl() SetWebhook200ResponseWebhookUrl`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *SetWebhook200Response) GetWebhookUrlOk() (*SetWebhook200ResponseWebhookUrl, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *SetWebhook200Response) SetWebhookUrl(v SetWebhook200ResponseWebhookUrl)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *SetWebhook200Response) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


