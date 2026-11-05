# CreateSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | Pointer to [**CreateSettingsRequestWebhookUrl**](CreateSettingsRequestWebhookUrl.md) |  | [optional] 
**GuaranteedHooks** | Pointer to **bool** |  | [optional] 
**AckNotificationsOn** | Pointer to **bool** |  | [optional] 
**RawHooks** | Pointer to **bool** |  | [optional] 
**TemplateAnalyticsEnabled** | Pointer to **bool** |  | [optional] 
**CtaUrlLinkTrackingOptedOut** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateSettingsRequest

`func NewCreateSettingsRequest() *CreateSettingsRequest`

NewCreateSettingsRequest instantiates a new CreateSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSettingsRequestWithDefaults

`func NewCreateSettingsRequestWithDefaults() *CreateSettingsRequest`

NewCreateSettingsRequestWithDefaults instantiates a new CreateSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *CreateSettingsRequest) GetWebhookUrl() CreateSettingsRequestWebhookUrl`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateSettingsRequest) GetWebhookUrlOk() (*CreateSettingsRequestWebhookUrl, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateSettingsRequest) SetWebhookUrl(v CreateSettingsRequestWebhookUrl)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateSettingsRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetGuaranteedHooks

`func (o *CreateSettingsRequest) GetGuaranteedHooks() bool`

GetGuaranteedHooks returns the GuaranteedHooks field if non-nil, zero value otherwise.

### GetGuaranteedHooksOk

`func (o *CreateSettingsRequest) GetGuaranteedHooksOk() (*bool, bool)`

GetGuaranteedHooksOk returns a tuple with the GuaranteedHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedHooks

`func (o *CreateSettingsRequest) SetGuaranteedHooks(v bool)`

SetGuaranteedHooks sets GuaranteedHooks field to given value.

### HasGuaranteedHooks

`func (o *CreateSettingsRequest) HasGuaranteedHooks() bool`

HasGuaranteedHooks returns a boolean if a field has been set.

### GetAckNotificationsOn

`func (o *CreateSettingsRequest) GetAckNotificationsOn() bool`

GetAckNotificationsOn returns the AckNotificationsOn field if non-nil, zero value otherwise.

### GetAckNotificationsOnOk

`func (o *CreateSettingsRequest) GetAckNotificationsOnOk() (*bool, bool)`

GetAckNotificationsOnOk returns a tuple with the AckNotificationsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckNotificationsOn

`func (o *CreateSettingsRequest) SetAckNotificationsOn(v bool)`

SetAckNotificationsOn sets AckNotificationsOn field to given value.

### HasAckNotificationsOn

`func (o *CreateSettingsRequest) HasAckNotificationsOn() bool`

HasAckNotificationsOn returns a boolean if a field has been set.

### GetRawHooks

`func (o *CreateSettingsRequest) GetRawHooks() bool`

GetRawHooks returns the RawHooks field if non-nil, zero value otherwise.

### GetRawHooksOk

`func (o *CreateSettingsRequest) GetRawHooksOk() (*bool, bool)`

GetRawHooksOk returns a tuple with the RawHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawHooks

`func (o *CreateSettingsRequest) SetRawHooks(v bool)`

SetRawHooks sets RawHooks field to given value.

### HasRawHooks

`func (o *CreateSettingsRequest) HasRawHooks() bool`

HasRawHooks returns a boolean if a field has been set.

### GetTemplateAnalyticsEnabled

`func (o *CreateSettingsRequest) GetTemplateAnalyticsEnabled() bool`

GetTemplateAnalyticsEnabled returns the TemplateAnalyticsEnabled field if non-nil, zero value otherwise.

### GetTemplateAnalyticsEnabledOk

`func (o *CreateSettingsRequest) GetTemplateAnalyticsEnabledOk() (*bool, bool)`

GetTemplateAnalyticsEnabledOk returns a tuple with the TemplateAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateAnalyticsEnabled

`func (o *CreateSettingsRequest) SetTemplateAnalyticsEnabled(v bool)`

SetTemplateAnalyticsEnabled sets TemplateAnalyticsEnabled field to given value.

### HasTemplateAnalyticsEnabled

`func (o *CreateSettingsRequest) HasTemplateAnalyticsEnabled() bool`

HasTemplateAnalyticsEnabled returns a boolean if a field has been set.

### GetCtaUrlLinkTrackingOptedOut

`func (o *CreateSettingsRequest) GetCtaUrlLinkTrackingOptedOut() bool`

GetCtaUrlLinkTrackingOptedOut returns the CtaUrlLinkTrackingOptedOut field if non-nil, zero value otherwise.

### GetCtaUrlLinkTrackingOptedOutOk

`func (o *CreateSettingsRequest) GetCtaUrlLinkTrackingOptedOutOk() (*bool, bool)`

GetCtaUrlLinkTrackingOptedOutOk returns a tuple with the CtaUrlLinkTrackingOptedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaUrlLinkTrackingOptedOut

`func (o *CreateSettingsRequest) SetCtaUrlLinkTrackingOptedOut(v bool)`

SetCtaUrlLinkTrackingOptedOut sets CtaUrlLinkTrackingOptedOut field to given value.

### HasCtaUrlLinkTrackingOptedOut

`func (o *CreateSettingsRequest) HasCtaUrlLinkTrackingOptedOut() bool`

HasCtaUrlLinkTrackingOptedOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


