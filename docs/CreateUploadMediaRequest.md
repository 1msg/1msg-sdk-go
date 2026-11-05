# CreateUploadMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | File HTTP(S) URL or base64 data URI (&#x60;data:&lt;mime&gt;;base64,...&#x60;) | [optional] 
**Url** | Pointer to **string** | Alias for &#x60;body&#x60; — public HTTP(S) URL of the file to upload | [optional] 

## Methods

### NewCreateUploadMediaRequest

`func NewCreateUploadMediaRequest() *CreateUploadMediaRequest`

NewCreateUploadMediaRequest instantiates a new CreateUploadMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUploadMediaRequestWithDefaults

`func NewCreateUploadMediaRequestWithDefaults() *CreateUploadMediaRequest`

NewCreateUploadMediaRequestWithDefaults instantiates a new CreateUploadMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CreateUploadMediaRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateUploadMediaRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateUploadMediaRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateUploadMediaRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetUrl

`func (o *CreateUploadMediaRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateUploadMediaRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateUploadMediaRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateUploadMediaRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


