# CallingSettingsCallingSipServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RequestUriUserParams** | Pointer to **map[string]string** |  | [optional] 
**SipUserPassword** | Pointer to **string** | Present on GET when SIP credentials are returned by Meta | [optional] 
**Password** | Pointer to **string** | Alternate password field name from some Meta responses | [optional] 
**AppId** | Pointer to **string** |  | [optional] 

## Methods

### NewCallingSettingsCallingSipServersInner

`func NewCallingSettingsCallingSipServersInner() *CallingSettingsCallingSipServersInner`

NewCallingSettingsCallingSipServersInner instantiates a new CallingSettingsCallingSipServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallingSettingsCallingSipServersInnerWithDefaults

`func NewCallingSettingsCallingSipServersInnerWithDefaults() *CallingSettingsCallingSipServersInner`

NewCallingSettingsCallingSipServersInnerWithDefaults instantiates a new CallingSettingsCallingSipServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *CallingSettingsCallingSipServersInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CallingSettingsCallingSipServersInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CallingSettingsCallingSipServersInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CallingSettingsCallingSipServersInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *CallingSettingsCallingSipServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CallingSettingsCallingSipServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CallingSettingsCallingSipServersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CallingSettingsCallingSipServersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRequestUriUserParams

`func (o *CallingSettingsCallingSipServersInner) GetRequestUriUserParams() map[string]string`

GetRequestUriUserParams returns the RequestUriUserParams field if non-nil, zero value otherwise.

### GetRequestUriUserParamsOk

`func (o *CallingSettingsCallingSipServersInner) GetRequestUriUserParamsOk() (*map[string]string, bool)`

GetRequestUriUserParamsOk returns a tuple with the RequestUriUserParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUriUserParams

`func (o *CallingSettingsCallingSipServersInner) SetRequestUriUserParams(v map[string]string)`

SetRequestUriUserParams sets RequestUriUserParams field to given value.

### HasRequestUriUserParams

`func (o *CallingSettingsCallingSipServersInner) HasRequestUriUserParams() bool`

HasRequestUriUserParams returns a boolean if a field has been set.

### GetSipUserPassword

`func (o *CallingSettingsCallingSipServersInner) GetSipUserPassword() string`

GetSipUserPassword returns the SipUserPassword field if non-nil, zero value otherwise.

### GetSipUserPasswordOk

`func (o *CallingSettingsCallingSipServersInner) GetSipUserPasswordOk() (*string, bool)`

GetSipUserPasswordOk returns a tuple with the SipUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUserPassword

`func (o *CallingSettingsCallingSipServersInner) SetSipUserPassword(v string)`

SetSipUserPassword sets SipUserPassword field to given value.

### HasSipUserPassword

`func (o *CallingSettingsCallingSipServersInner) HasSipUserPassword() bool`

HasSipUserPassword returns a boolean if a field has been set.

### GetPassword

`func (o *CallingSettingsCallingSipServersInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CallingSettingsCallingSipServersInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CallingSettingsCallingSipServersInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CallingSettingsCallingSipServersInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAppId

`func (o *CallingSettingsCallingSipServersInner) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CallingSettingsCallingSipServersInner) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CallingSettingsCallingSipServersInner) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CallingSettingsCallingSipServersInner) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


