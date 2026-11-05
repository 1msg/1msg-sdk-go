# CreateGroupsGroupIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | New group subject / name | [optional] 
**Description** | Pointer to **string** | New group description | [optional] 
**ProfilePictureFile** | Pointer to **string** | Image URL or base64 data URI for the group profile picture | [optional] 

## Methods

### NewCreateGroupsGroupIdRequest

`func NewCreateGroupsGroupIdRequest() *CreateGroupsGroupIdRequest`

NewCreateGroupsGroupIdRequest instantiates a new CreateGroupsGroupIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupsGroupIdRequestWithDefaults

`func NewCreateGroupsGroupIdRequestWithDefaults() *CreateGroupsGroupIdRequest`

NewCreateGroupsGroupIdRequestWithDefaults instantiates a new CreateGroupsGroupIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CreateGroupsGroupIdRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateGroupsGroupIdRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateGroupsGroupIdRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateGroupsGroupIdRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDescription

`func (o *CreateGroupsGroupIdRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupsGroupIdRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupsGroupIdRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupsGroupIdRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProfilePictureFile

`func (o *CreateGroupsGroupIdRequest) GetProfilePictureFile() string`

GetProfilePictureFile returns the ProfilePictureFile field if non-nil, zero value otherwise.

### GetProfilePictureFileOk

`func (o *CreateGroupsGroupIdRequest) GetProfilePictureFileOk() (*string, bool)`

GetProfilePictureFileOk returns a tuple with the ProfilePictureFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureFile

`func (o *CreateGroupsGroupIdRequest) SetProfilePictureFile(v string)`

SetProfilePictureFile sets ProfilePictureFile field to given value.

### HasProfilePictureFile

`func (o *CreateGroupsGroupIdRequest) HasProfilePictureFile() bool`

HasProfilePictureFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


