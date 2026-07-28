# SendContactRequestContactsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**SendContactRequestContactsInnerName**](SendContactRequestContactsInnerName.md) |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to **[]interface{}** |  | [optional] 
**Emails** | Pointer to **[]interface{}** |  | [optional] 
**Org** | Pointer to [**SendContactRequestContactsInnerOrg**](SendContactRequestContactsInnerOrg.md) |  | [optional] 
**Phones** | Pointer to [**[]SendContactRequestContactsInnerPhonesInner**](SendContactRequestContactsInnerPhonesInner.md) |  | [optional] 
**Urls** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewSendContactRequestContactsInner

`func NewSendContactRequestContactsInner() *SendContactRequestContactsInner`

NewSendContactRequestContactsInner instantiates a new SendContactRequestContactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendContactRequestContactsInnerWithDefaults

`func NewSendContactRequestContactsInnerWithDefaults() *SendContactRequestContactsInner`

NewSendContactRequestContactsInnerWithDefaults instantiates a new SendContactRequestContactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SendContactRequestContactsInner) GetName() SendContactRequestContactsInnerName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendContactRequestContactsInner) GetNameOk() (*SendContactRequestContactsInnerName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendContactRequestContactsInner) SetName(v SendContactRequestContactsInnerName)`

SetName sets Name field to given value.

### HasName

`func (o *SendContactRequestContactsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBirthday

`func (o *SendContactRequestContactsInner) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *SendContactRequestContactsInner) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *SendContactRequestContactsInner) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *SendContactRequestContactsInner) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetAddresses

`func (o *SendContactRequestContactsInner) GetAddresses() []interface{}`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *SendContactRequestContactsInner) GetAddressesOk() (*[]interface{}, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *SendContactRequestContactsInner) SetAddresses(v []interface{})`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *SendContactRequestContactsInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetEmails

`func (o *SendContactRequestContactsInner) GetEmails() []interface{}`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SendContactRequestContactsInner) GetEmailsOk() (*[]interface{}, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SendContactRequestContactsInner) SetEmails(v []interface{})`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *SendContactRequestContactsInner) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetOrg

`func (o *SendContactRequestContactsInner) GetOrg() SendContactRequestContactsInnerOrg`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SendContactRequestContactsInner) GetOrgOk() (*SendContactRequestContactsInnerOrg, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SendContactRequestContactsInner) SetOrg(v SendContactRequestContactsInnerOrg)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *SendContactRequestContactsInner) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhones

`func (o *SendContactRequestContactsInner) GetPhones() []SendContactRequestContactsInnerPhonesInner`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *SendContactRequestContactsInner) GetPhonesOk() (*[]SendContactRequestContactsInnerPhonesInner, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *SendContactRequestContactsInner) SetPhones(v []SendContactRequestContactsInnerPhonesInner)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *SendContactRequestContactsInner) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetUrls

`func (o *SendContactRequestContactsInner) GetUrls() []interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SendContactRequestContactsInner) GetUrlsOk() (*[]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SendContactRequestContactsInner) SetUrls(v []interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *SendContactRequestContactsInner) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


