# InitiateCallRequestSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdpType** | **string** | offer for outbound connect; answer for pre_accept/accept | 
**Sdp** | **string** | Full SDP (RFC 8866). Replace placeholders with SDP from your WebRTC stack. Must negotiate ICE, DTLS-SRTP, and OPUS for WhatsApp media.  | 

## Methods

### NewInitiateCallRequestSession

`func NewInitiateCallRequestSession(sdpType string, sdp string, ) *InitiateCallRequestSession`

NewInitiateCallRequestSession instantiates a new InitiateCallRequestSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateCallRequestSessionWithDefaults

`func NewInitiateCallRequestSessionWithDefaults() *InitiateCallRequestSession`

NewInitiateCallRequestSessionWithDefaults instantiates a new InitiateCallRequestSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdpType

`func (o *InitiateCallRequestSession) GetSdpType() string`

GetSdpType returns the SdpType field if non-nil, zero value otherwise.

### GetSdpTypeOk

`func (o *InitiateCallRequestSession) GetSdpTypeOk() (*string, bool)`

GetSdpTypeOk returns a tuple with the SdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdpType

`func (o *InitiateCallRequestSession) SetSdpType(v string)`

SetSdpType sets SdpType field to given value.


### GetSdp

`func (o *InitiateCallRequestSession) GetSdp() string`

GetSdp returns the Sdp field if non-nil, zero value otherwise.

### GetSdpOk

`func (o *InitiateCallRequestSession) GetSdpOk() (*string, bool)`

GetSdpOk returns a tuple with the Sdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdp

`func (o *InitiateCallRequestSession) SetSdp(v string)`

SetSdp sets Sdp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


