# GetWhatsappBusinessEncryption200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessPublicKey** | Pointer to **string** | Stored 2048-bit RSA public key (PEM) | [optional] 
**BusinessPublicKeySignatureStatus** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetWhatsappBusinessEncryption200Response

`func NewGetWhatsappBusinessEncryption200Response() *GetWhatsappBusinessEncryption200Response`

NewGetWhatsappBusinessEncryption200Response instantiates a new GetWhatsappBusinessEncryption200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsappBusinessEncryption200ResponseWithDefaults

`func NewGetWhatsappBusinessEncryption200ResponseWithDefaults() *GetWhatsappBusinessEncryption200Response`

NewGetWhatsappBusinessEncryption200ResponseWithDefaults instantiates a new GetWhatsappBusinessEncryption200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessPublicKey

`func (o *GetWhatsappBusinessEncryption200Response) GetBusinessPublicKey() string`

GetBusinessPublicKey returns the BusinessPublicKey field if non-nil, zero value otherwise.

### GetBusinessPublicKeyOk

`func (o *GetWhatsappBusinessEncryption200Response) GetBusinessPublicKeyOk() (*string, bool)`

GetBusinessPublicKeyOk returns a tuple with the BusinessPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPublicKey

`func (o *GetWhatsappBusinessEncryption200Response) SetBusinessPublicKey(v string)`

SetBusinessPublicKey sets BusinessPublicKey field to given value.

### HasBusinessPublicKey

`func (o *GetWhatsappBusinessEncryption200Response) HasBusinessPublicKey() bool`

HasBusinessPublicKey returns a boolean if a field has been set.

### GetBusinessPublicKeySignatureStatus

`func (o *GetWhatsappBusinessEncryption200Response) GetBusinessPublicKeySignatureStatus() string`

GetBusinessPublicKeySignatureStatus returns the BusinessPublicKeySignatureStatus field if non-nil, zero value otherwise.

### GetBusinessPublicKeySignatureStatusOk

`func (o *GetWhatsappBusinessEncryption200Response) GetBusinessPublicKeySignatureStatusOk() (*string, bool)`

GetBusinessPublicKeySignatureStatusOk returns a tuple with the BusinessPublicKeySignatureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPublicKeySignatureStatus

`func (o *GetWhatsappBusinessEncryption200Response) SetBusinessPublicKeySignatureStatus(v string)`

SetBusinessPublicKeySignatureStatus sets BusinessPublicKeySignatureStatus field to given value.

### HasBusinessPublicKeySignatureStatus

`func (o *GetWhatsappBusinessEncryption200Response) HasBusinessPublicKeySignatureStatus() bool`

HasBusinessPublicKeySignatureStatus returns a boolean if a field has been set.

### GetData

`func (o *GetWhatsappBusinessEncryption200Response) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetWhatsappBusinessEncryption200Response) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetWhatsappBusinessEncryption200Response) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GetWhatsappBusinessEncryption200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


