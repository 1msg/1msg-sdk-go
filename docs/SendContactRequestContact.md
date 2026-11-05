# SendContactRequestContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**SendContactRequestContactName**](SendContactRequestContactName.md) |  | 
**Phone** | **string** | Contact phone in E.164 (&#x60;+79181976551&#x60;), not a timestamp | 
**Email** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **map[string]interface{}** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewSendContactRequestContact

`func NewSendContactRequestContact(name SendContactRequestContactName, phone string, ) *SendContactRequestContact`

NewSendContactRequestContact instantiates a new SendContactRequestContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendContactRequestContactWithDefaults

`func NewSendContactRequestContactWithDefaults() *SendContactRequestContact`

NewSendContactRequestContactWithDefaults instantiates a new SendContactRequestContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SendContactRequestContact) GetName() SendContactRequestContactName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendContactRequestContact) GetNameOk() (*SendContactRequestContactName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendContactRequestContact) SetName(v SendContactRequestContactName)`

SetName sets Name field to given value.


### GetPhone

`func (o *SendContactRequestContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendContactRequestContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendContactRequestContact) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *SendContactRequestContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendContactRequestContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendContactRequestContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SendContactRequestContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrg

`func (o *SendContactRequestContact) GetOrg() map[string]interface{}`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SendContactRequestContact) GetOrgOk() (*map[string]interface{}, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SendContactRequestContact) SetOrg(v map[string]interface{})`

SetOrg sets Org field to given value.

### HasOrg

`func (o *SendContactRequestContact) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetUrl

`func (o *SendContactRequestContact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SendContactRequestContact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SendContactRequestContact) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SendContactRequestContact) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


