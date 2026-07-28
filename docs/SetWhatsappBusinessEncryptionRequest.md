# SetWhatsappBusinessEncryptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessPublicKey** | **string** | 2048-bit RSA public key in PEM format | 

## Methods

### NewSetWhatsappBusinessEncryptionRequest

`func NewSetWhatsappBusinessEncryptionRequest(businessPublicKey string, ) *SetWhatsappBusinessEncryptionRequest`

NewSetWhatsappBusinessEncryptionRequest instantiates a new SetWhatsappBusinessEncryptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWhatsappBusinessEncryptionRequestWithDefaults

`func NewSetWhatsappBusinessEncryptionRequestWithDefaults() *SetWhatsappBusinessEncryptionRequest`

NewSetWhatsappBusinessEncryptionRequestWithDefaults instantiates a new SetWhatsappBusinessEncryptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessPublicKey

`func (o *SetWhatsappBusinessEncryptionRequest) GetBusinessPublicKey() string`

GetBusinessPublicKey returns the BusinessPublicKey field if non-nil, zero value otherwise.

### GetBusinessPublicKeyOk

`func (o *SetWhatsappBusinessEncryptionRequest) GetBusinessPublicKeyOk() (*string, bool)`

GetBusinessPublicKeyOk returns a tuple with the BusinessPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPublicKey

`func (o *SetWhatsappBusinessEncryptionRequest) SetBusinessPublicKey(v string)`

SetBusinessPublicKey sets BusinessPublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


