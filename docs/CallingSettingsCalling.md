# CallingSettingsCalling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Enable or disable Calling API on this number | [optional] 
**CallIconVisibility** | Pointer to **string** | Controls call icon visibility in the WhatsApp client | [optional] 
**CallbackPermissionStatus** | Pointer to **string** | When ENABLED, a user who calls your business automatically grants call permission for business-initiated callbacks (subject to Meta rules).  | [optional] 
**SrtpKeyExchangeProtocol** | Pointer to **string** | SRTP key exchange. DTLS is default/recommended. SDES is only valid when SIP signaling is enabled.  | [optional] 
**CallHours** | Pointer to [**CallingSettingsCallingCallHours**](CallingSettingsCallingCallHours.md) |  | [optional] 
**Sip** | Pointer to [**CallingSettingsCallingSip**](CallingSettingsCallingSip.md) |  | [optional] 
**Video** | Pointer to [**CallingSettingsCallingVideo**](CallingSettingsCallingVideo.md) |  | [optional] 
**Audio** | Pointer to [**CallingSettingsCallingAudio**](CallingSettingsCallingAudio.md) |  | [optional] 
**Restrictions** | Pointer to **map[string]interface{}** | Calling restrictions (response-only) | [optional] 
**IpAddresses** | Pointer to **map[string]interface{}** | Meta media IP ranges (response-only) | [optional] 
**CallIcons** | Pointer to **map[string]interface{}** | Call icon country restrictions (response-only) | [optional] 

## Methods

### NewCallingSettingsCalling

`func NewCallingSettingsCalling() *CallingSettingsCalling`

NewCallingSettingsCalling instantiates a new CallingSettingsCalling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallingSettingsCallingWithDefaults

`func NewCallingSettingsCallingWithDefaults() *CallingSettingsCalling`

NewCallingSettingsCallingWithDefaults instantiates a new CallingSettingsCalling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CallingSettingsCalling) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallingSettingsCalling) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallingSettingsCalling) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallingSettingsCalling) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCallIconVisibility

`func (o *CallingSettingsCalling) GetCallIconVisibility() string`

GetCallIconVisibility returns the CallIconVisibility field if non-nil, zero value otherwise.

### GetCallIconVisibilityOk

`func (o *CallingSettingsCalling) GetCallIconVisibilityOk() (*string, bool)`

GetCallIconVisibilityOk returns a tuple with the CallIconVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIconVisibility

`func (o *CallingSettingsCalling) SetCallIconVisibility(v string)`

SetCallIconVisibility sets CallIconVisibility field to given value.

### HasCallIconVisibility

`func (o *CallingSettingsCalling) HasCallIconVisibility() bool`

HasCallIconVisibility returns a boolean if a field has been set.

### GetCallbackPermissionStatus

`func (o *CallingSettingsCalling) GetCallbackPermissionStatus() string`

GetCallbackPermissionStatus returns the CallbackPermissionStatus field if non-nil, zero value otherwise.

### GetCallbackPermissionStatusOk

`func (o *CallingSettingsCalling) GetCallbackPermissionStatusOk() (*string, bool)`

GetCallbackPermissionStatusOk returns a tuple with the CallbackPermissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackPermissionStatus

`func (o *CallingSettingsCalling) SetCallbackPermissionStatus(v string)`

SetCallbackPermissionStatus sets CallbackPermissionStatus field to given value.

### HasCallbackPermissionStatus

`func (o *CallingSettingsCalling) HasCallbackPermissionStatus() bool`

HasCallbackPermissionStatus returns a boolean if a field has been set.

### GetSrtpKeyExchangeProtocol

`func (o *CallingSettingsCalling) GetSrtpKeyExchangeProtocol() string`

GetSrtpKeyExchangeProtocol returns the SrtpKeyExchangeProtocol field if non-nil, zero value otherwise.

### GetSrtpKeyExchangeProtocolOk

`func (o *CallingSettingsCalling) GetSrtpKeyExchangeProtocolOk() (*string, bool)`

GetSrtpKeyExchangeProtocolOk returns a tuple with the SrtpKeyExchangeProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrtpKeyExchangeProtocol

`func (o *CallingSettingsCalling) SetSrtpKeyExchangeProtocol(v string)`

SetSrtpKeyExchangeProtocol sets SrtpKeyExchangeProtocol field to given value.

### HasSrtpKeyExchangeProtocol

`func (o *CallingSettingsCalling) HasSrtpKeyExchangeProtocol() bool`

HasSrtpKeyExchangeProtocol returns a boolean if a field has been set.

### GetCallHours

`func (o *CallingSettingsCalling) GetCallHours() CallingSettingsCallingCallHours`

GetCallHours returns the CallHours field if non-nil, zero value otherwise.

### GetCallHoursOk

`func (o *CallingSettingsCalling) GetCallHoursOk() (*CallingSettingsCallingCallHours, bool)`

GetCallHoursOk returns a tuple with the CallHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallHours

`func (o *CallingSettingsCalling) SetCallHours(v CallingSettingsCallingCallHours)`

SetCallHours sets CallHours field to given value.

### HasCallHours

`func (o *CallingSettingsCalling) HasCallHours() bool`

HasCallHours returns a boolean if a field has been set.

### GetSip

`func (o *CallingSettingsCalling) GetSip() CallingSettingsCallingSip`

GetSip returns the Sip field if non-nil, zero value otherwise.

### GetSipOk

`func (o *CallingSettingsCalling) GetSipOk() (*CallingSettingsCallingSip, bool)`

GetSipOk returns a tuple with the Sip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSip

`func (o *CallingSettingsCalling) SetSip(v CallingSettingsCallingSip)`

SetSip sets Sip field to given value.

### HasSip

`func (o *CallingSettingsCalling) HasSip() bool`

HasSip returns a boolean if a field has been set.

### GetVideo

`func (o *CallingSettingsCalling) GetVideo() CallingSettingsCallingVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CallingSettingsCalling) GetVideoOk() (*CallingSettingsCallingVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CallingSettingsCalling) SetVideo(v CallingSettingsCallingVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CallingSettingsCalling) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetAudio

`func (o *CallingSettingsCalling) GetAudio() CallingSettingsCallingAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *CallingSettingsCalling) GetAudioOk() (*CallingSettingsCallingAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *CallingSettingsCalling) SetAudio(v CallingSettingsCallingAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *CallingSettingsCalling) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetRestrictions

`func (o *CallingSettingsCalling) GetRestrictions() map[string]interface{}`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *CallingSettingsCalling) GetRestrictionsOk() (*map[string]interface{}, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *CallingSettingsCalling) SetRestrictions(v map[string]interface{})`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *CallingSettingsCalling) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetIpAddresses

`func (o *CallingSettingsCalling) GetIpAddresses() map[string]interface{}`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CallingSettingsCalling) GetIpAddressesOk() (*map[string]interface{}, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CallingSettingsCalling) SetIpAddresses(v map[string]interface{})`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CallingSettingsCalling) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetCallIcons

`func (o *CallingSettingsCalling) GetCallIcons() map[string]interface{}`

GetCallIcons returns the CallIcons field if non-nil, zero value otherwise.

### GetCallIconsOk

`func (o *CallingSettingsCalling) GetCallIconsOk() (*map[string]interface{}, bool)`

GetCallIconsOk returns a tuple with the CallIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIcons

`func (o *CallingSettingsCalling) SetCallIcons(v map[string]interface{})`

SetCallIcons sets CallIcons field to given value.

### HasCallIcons

`func (o *CallingSettingsCalling) HasCallIcons() bool`

HasCallIcons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


