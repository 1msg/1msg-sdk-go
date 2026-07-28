# UpdateMeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** | Profile About section | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Business description (max 256 characters) | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Vertical** | Pointer to **string** | Industry of the business | [optional] 
**Photo** | Pointer to **string** | HTTPS URL or data:image/...;base64,... payload | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateMeRequest

`func NewUpdateMeRequest() *UpdateMeRequest`

NewUpdateMeRequest instantiates a new UpdateMeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMeRequestWithDefaults

`func NewUpdateMeRequestWithDefaults() *UpdateMeRequest`

NewUpdateMeRequestWithDefaults instantiates a new UpdateMeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *UpdateMeRequest) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *UpdateMeRequest) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *UpdateMeRequest) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *UpdateMeRequest) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateMeRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateMeRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateMeRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateMeRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateMeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateMeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateMeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateMeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateMeRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetVertical

`func (o *UpdateMeRequest) GetVertical() string`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *UpdateMeRequest) GetVerticalOk() (*string, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *UpdateMeRequest) SetVertical(v string)`

SetVertical sets Vertical field to given value.

### HasVertical

`func (o *UpdateMeRequest) HasVertical() bool`

HasVertical returns a boolean if a field has been set.

### GetPhoto

`func (o *UpdateMeRequest) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *UpdateMeRequest) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *UpdateMeRequest) SetPhoto(v string)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *UpdateMeRequest) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetWebsites

`func (o *UpdateMeRequest) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *UpdateMeRequest) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *UpdateMeRequest) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *UpdateMeRequest) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


