# GetStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCloud** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** | typically &#x60;connected&#x60; when the number can send | [optional] 
**AccountStatus** | Pointer to **string** |  | [optional] 
**MmLiteAvailable** | Pointer to **bool** |  | [optional] 
**HealthStatusWarning** | Pointer to **string** | Optional soft-fail warning when Meta Graph health_status could not be loaded. Never includes internal provider names. | [optional] 

## Methods

### NewGetStatus200Response

`func NewGetStatus200Response() *GetStatus200Response`

NewGetStatus200Response instantiates a new GetStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatus200ResponseWithDefaults

`func NewGetStatus200ResponseWithDefaults() *GetStatus200Response`

NewGetStatus200ResponseWithDefaults instantiates a new GetStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCloud

`func (o *GetStatus200Response) GetIsCloud() bool`

GetIsCloud returns the IsCloud field if non-nil, zero value otherwise.

### GetIsCloudOk

`func (o *GetStatus200Response) GetIsCloudOk() (*bool, bool)`

GetIsCloudOk returns a tuple with the IsCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloud

`func (o *GetStatus200Response) SetIsCloud(v bool)`

SetIsCloud sets IsCloud field to given value.

### HasIsCloud

`func (o *GetStatus200Response) HasIsCloud() bool`

HasIsCloud returns a boolean if a field has been set.

### GetStatus

`func (o *GetStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccountStatus

`func (o *GetStatus200Response) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *GetStatus200Response) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *GetStatus200Response) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *GetStatus200Response) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetMmLiteAvailable

`func (o *GetStatus200Response) GetMmLiteAvailable() bool`

GetMmLiteAvailable returns the MmLiteAvailable field if non-nil, zero value otherwise.

### GetMmLiteAvailableOk

`func (o *GetStatus200Response) GetMmLiteAvailableOk() (*bool, bool)`

GetMmLiteAvailableOk returns a tuple with the MmLiteAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmLiteAvailable

`func (o *GetStatus200Response) SetMmLiteAvailable(v bool)`

SetMmLiteAvailable sets MmLiteAvailable field to given value.

### HasMmLiteAvailable

`func (o *GetStatus200Response) HasMmLiteAvailable() bool`

HasMmLiteAvailable returns a boolean if a field has been set.

### GetHealthStatusWarning

`func (o *GetStatus200Response) GetHealthStatusWarning() string`

GetHealthStatusWarning returns the HealthStatusWarning field if non-nil, zero value otherwise.

### GetHealthStatusWarningOk

`func (o *GetStatus200Response) GetHealthStatusWarningOk() (*string, bool)`

GetHealthStatusWarningOk returns a tuple with the HealthStatusWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatusWarning

`func (o *GetStatus200Response) SetHealthStatusWarning(v string)`

SetHealthStatusWarning sets HealthStatusWarning field to given value.

### HasHealthStatusWarning

`func (o *GetStatus200Response) HasHealthStatusWarning() bool`

HasHealthStatusWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


