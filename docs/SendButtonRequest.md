# SendButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Sections** | Pointer to [**[]SendButtonRequestSectionsInner**](SendButtonRequestSectionsInner.md) |  | [optional] 

## Methods

### NewSendButtonRequest

`func NewSendButtonRequest() *SendButtonRequest`

NewSendButtonRequest instantiates a new SendButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendButtonRequestWithDefaults

`func NewSendButtonRequestWithDefaults() *SendButtonRequest`

NewSendButtonRequestWithDefaults instantiates a new SendButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendButtonRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendButtonRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendButtonRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendButtonRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetBody

`func (o *SendButtonRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendButtonRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendButtonRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendButtonRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFooter

`func (o *SendButtonRequest) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *SendButtonRequest) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *SendButtonRequest) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *SendButtonRequest) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetSections

`func (o *SendButtonRequest) GetSections() []SendButtonRequestSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SendButtonRequest) GetSectionsOk() (*[]SendButtonRequestSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SendButtonRequest) SetSections(v []SendButtonRequestSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *SendButtonRequest) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


