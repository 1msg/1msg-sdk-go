# ProfileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** | Profile&#39;s About section | [optional] 
**Address** | Pointer to **string** | Address of the business | [optional] 
**Description** | Pointer to **string** | Description of the business | [optional] 
**Email** | Pointer to **string** | Business email | [optional] 
**Phone** | Pointer to **string** | Linked phone number | [optional] 
**Vertical** | Pointer to **string** | Industry of the business | [optional] 
**Photo** | Pointer to **string** | Profile photo URL | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProfileInfo

`func NewProfileInfo() *ProfileInfo`

NewProfileInfo instantiates a new ProfileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileInfoWithDefaults

`func NewProfileInfoWithDefaults() *ProfileInfo`

NewProfileInfoWithDefaults instantiates a new ProfileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *ProfileInfo) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ProfileInfo) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ProfileInfo) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ProfileInfo) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetAddress

`func (o *ProfileInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ProfileInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ProfileInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ProfileInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *ProfileInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProfileInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProfileInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProfileInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *ProfileInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProfileInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProfileInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProfileInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *ProfileInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ProfileInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ProfileInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ProfileInfo) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetVertical

`func (o *ProfileInfo) GetVertical() string`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *ProfileInfo) GetVerticalOk() (*string, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *ProfileInfo) SetVertical(v string)`

SetVertical sets Vertical field to given value.

### HasVertical

`func (o *ProfileInfo) HasVertical() bool`

HasVertical returns a boolean if a field has been set.

### GetPhoto

`func (o *ProfileInfo) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ProfileInfo) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ProfileInfo) SetPhoto(v string)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ProfileInfo) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetWebsites

`func (o *ProfileInfo) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *ProfileInfo) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *ProfileInfo) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *ProfileInfo) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


