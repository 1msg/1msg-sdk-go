# ListSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | Pointer to [**ListSettings200ResponseWebhookUrl**](ListSettings200ResponseWebhookUrl.md) |  | [optional] 
**GuaranteedHooks** | Pointer to **bool** | Retry webhooks until acknowledged | [optional] 
**AckNotificationsOn** | Pointer to **bool** | Include message ACK / status events in webhooks | [optional] 
**RawHooks** | Pointer to **bool** | Forward raw provider webhook payloads | [optional] 
**MmLiteEnabled** | Pointer to **bool** |  | [optional] 
**TemplateAnalyticsEnabled** | Pointer to **bool** |  | [optional] 
**CtaUrlLinkTrackingOptedOut** | Pointer to **bool** |  | [optional] 
**HaveDisabledWebhooks** | Pointer to **bool** | &#x60;true&#x60; when at least one URL in &#x60;webhookUrl&#x60; is sticky auto-disabled. Not \&quot;the channel is dead\&quot;: other configured URLs still receive events.  | [optional] 
**CircuitState** | Pointer to **string** | Compatibility scalar for **one** URL (first sticky-disabled, else first configured). Prefer &#x60;webhookStatuses[].circuitState&#x60;.  | [optional] 
**Reason** | Pointer to **string** | Last qualifying fail reason of the summary URL | [optional] 
**Fails** | Pointer to **int32** | Diagnostic only (summary URL). Qualifying failures in the rolling ~10 minute breaker window — not lifetime. Prefer &#x60;webhookStatuses&#x60;.  | [optional] 
**Trips** | Pointer to **int32** | Diagnostic only (summary URL). Failed recovery probes (≈15 min → 1 h → 6 h). Sticky disable after 3. Not a blind unblock timer.  | [optional] 
**DisabledAt** | Pointer to **int32** | Epoch milliseconds when the summary URL became sticky-disabled | [optional] 
**FirstFailAt** | Pointer to **int32** | Diagnostic only (summary URL). Epoch ms of the first fail in the current ~10 minute window (pairs with &#x60;fails&#x60;).  | [optional] 
**LastStatus** | Pointer to **int32** | Last HTTP status from the summary URL, when known | [optional] 
**NextProbeAt** | Pointer to **int32** | Epoch milliseconds of the next half-open probe for the summary URL, if any | [optional] 
**DisabledWebhooks** | Pointer to **[]map[string]interface{}** | Currently configured URLs that are sticky auto-disabled (subset of &#x60;webhookUrl&#x60;) | [optional] 
**WebhookStatuses** | Pointer to [**[]WebhookUrlStatus**](WebhookUrlStatus.md) | One entry per currently configured URL. Prefer this over top-level &#x60;haveDisabledWebhooks&#x60; / &#x60;circuitState&#x60;. | [optional] 

## Methods

### NewListSettings200Response

`func NewListSettings200Response() *ListSettings200Response`

NewListSettings200Response instantiates a new ListSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSettings200ResponseWithDefaults

`func NewListSettings200ResponseWithDefaults() *ListSettings200Response`

NewListSettings200ResponseWithDefaults instantiates a new ListSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *ListSettings200Response) GetWebhookUrl() ListSettings200ResponseWebhookUrl`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ListSettings200Response) GetWebhookUrlOk() (*ListSettings200ResponseWebhookUrl, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ListSettings200Response) SetWebhookUrl(v ListSettings200ResponseWebhookUrl)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ListSettings200Response) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetGuaranteedHooks

`func (o *ListSettings200Response) GetGuaranteedHooks() bool`

GetGuaranteedHooks returns the GuaranteedHooks field if non-nil, zero value otherwise.

### GetGuaranteedHooksOk

`func (o *ListSettings200Response) GetGuaranteedHooksOk() (*bool, bool)`

GetGuaranteedHooksOk returns a tuple with the GuaranteedHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedHooks

`func (o *ListSettings200Response) SetGuaranteedHooks(v bool)`

SetGuaranteedHooks sets GuaranteedHooks field to given value.

### HasGuaranteedHooks

`func (o *ListSettings200Response) HasGuaranteedHooks() bool`

HasGuaranteedHooks returns a boolean if a field has been set.

### GetAckNotificationsOn

`func (o *ListSettings200Response) GetAckNotificationsOn() bool`

GetAckNotificationsOn returns the AckNotificationsOn field if non-nil, zero value otherwise.

### GetAckNotificationsOnOk

`func (o *ListSettings200Response) GetAckNotificationsOnOk() (*bool, bool)`

GetAckNotificationsOnOk returns a tuple with the AckNotificationsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckNotificationsOn

`func (o *ListSettings200Response) SetAckNotificationsOn(v bool)`

SetAckNotificationsOn sets AckNotificationsOn field to given value.

### HasAckNotificationsOn

`func (o *ListSettings200Response) HasAckNotificationsOn() bool`

HasAckNotificationsOn returns a boolean if a field has been set.

### GetRawHooks

`func (o *ListSettings200Response) GetRawHooks() bool`

GetRawHooks returns the RawHooks field if non-nil, zero value otherwise.

### GetRawHooksOk

`func (o *ListSettings200Response) GetRawHooksOk() (*bool, bool)`

GetRawHooksOk returns a tuple with the RawHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawHooks

`func (o *ListSettings200Response) SetRawHooks(v bool)`

SetRawHooks sets RawHooks field to given value.

### HasRawHooks

`func (o *ListSettings200Response) HasRawHooks() bool`

HasRawHooks returns a boolean if a field has been set.

### GetMmLiteEnabled

`func (o *ListSettings200Response) GetMmLiteEnabled() bool`

GetMmLiteEnabled returns the MmLiteEnabled field if non-nil, zero value otherwise.

### GetMmLiteEnabledOk

`func (o *ListSettings200Response) GetMmLiteEnabledOk() (*bool, bool)`

GetMmLiteEnabledOk returns a tuple with the MmLiteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmLiteEnabled

`func (o *ListSettings200Response) SetMmLiteEnabled(v bool)`

SetMmLiteEnabled sets MmLiteEnabled field to given value.

### HasMmLiteEnabled

`func (o *ListSettings200Response) HasMmLiteEnabled() bool`

HasMmLiteEnabled returns a boolean if a field has been set.

### GetTemplateAnalyticsEnabled

`func (o *ListSettings200Response) GetTemplateAnalyticsEnabled() bool`

GetTemplateAnalyticsEnabled returns the TemplateAnalyticsEnabled field if non-nil, zero value otherwise.

### GetTemplateAnalyticsEnabledOk

`func (o *ListSettings200Response) GetTemplateAnalyticsEnabledOk() (*bool, bool)`

GetTemplateAnalyticsEnabledOk returns a tuple with the TemplateAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateAnalyticsEnabled

`func (o *ListSettings200Response) SetTemplateAnalyticsEnabled(v bool)`

SetTemplateAnalyticsEnabled sets TemplateAnalyticsEnabled field to given value.

### HasTemplateAnalyticsEnabled

`func (o *ListSettings200Response) HasTemplateAnalyticsEnabled() bool`

HasTemplateAnalyticsEnabled returns a boolean if a field has been set.

### GetCtaUrlLinkTrackingOptedOut

`func (o *ListSettings200Response) GetCtaUrlLinkTrackingOptedOut() bool`

GetCtaUrlLinkTrackingOptedOut returns the CtaUrlLinkTrackingOptedOut field if non-nil, zero value otherwise.

### GetCtaUrlLinkTrackingOptedOutOk

`func (o *ListSettings200Response) GetCtaUrlLinkTrackingOptedOutOk() (*bool, bool)`

GetCtaUrlLinkTrackingOptedOutOk returns a tuple with the CtaUrlLinkTrackingOptedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaUrlLinkTrackingOptedOut

`func (o *ListSettings200Response) SetCtaUrlLinkTrackingOptedOut(v bool)`

SetCtaUrlLinkTrackingOptedOut sets CtaUrlLinkTrackingOptedOut field to given value.

### HasCtaUrlLinkTrackingOptedOut

`func (o *ListSettings200Response) HasCtaUrlLinkTrackingOptedOut() bool`

HasCtaUrlLinkTrackingOptedOut returns a boolean if a field has been set.

### GetHaveDisabledWebhooks

`func (o *ListSettings200Response) GetHaveDisabledWebhooks() bool`

GetHaveDisabledWebhooks returns the HaveDisabledWebhooks field if non-nil, zero value otherwise.

### GetHaveDisabledWebhooksOk

`func (o *ListSettings200Response) GetHaveDisabledWebhooksOk() (*bool, bool)`

GetHaveDisabledWebhooksOk returns a tuple with the HaveDisabledWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaveDisabledWebhooks

`func (o *ListSettings200Response) SetHaveDisabledWebhooks(v bool)`

SetHaveDisabledWebhooks sets HaveDisabledWebhooks field to given value.

### HasHaveDisabledWebhooks

`func (o *ListSettings200Response) HasHaveDisabledWebhooks() bool`

HasHaveDisabledWebhooks returns a boolean if a field has been set.

### GetCircuitState

`func (o *ListSettings200Response) GetCircuitState() string`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *ListSettings200Response) GetCircuitStateOk() (*string, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *ListSettings200Response) SetCircuitState(v string)`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *ListSettings200Response) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### GetReason

`func (o *ListSettings200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListSettings200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListSettings200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ListSettings200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetFails

`func (o *ListSettings200Response) GetFails() int32`

GetFails returns the Fails field if non-nil, zero value otherwise.

### GetFailsOk

`func (o *ListSettings200Response) GetFailsOk() (*int32, bool)`

GetFailsOk returns a tuple with the Fails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFails

`func (o *ListSettings200Response) SetFails(v int32)`

SetFails sets Fails field to given value.

### HasFails

`func (o *ListSettings200Response) HasFails() bool`

HasFails returns a boolean if a field has been set.

### GetTrips

`func (o *ListSettings200Response) GetTrips() int32`

GetTrips returns the Trips field if non-nil, zero value otherwise.

### GetTripsOk

`func (o *ListSettings200Response) GetTripsOk() (*int32, bool)`

GetTripsOk returns a tuple with the Trips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrips

`func (o *ListSettings200Response) SetTrips(v int32)`

SetTrips sets Trips field to given value.

### HasTrips

`func (o *ListSettings200Response) HasTrips() bool`

HasTrips returns a boolean if a field has been set.

### GetDisabledAt

`func (o *ListSettings200Response) GetDisabledAt() int32`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *ListSettings200Response) GetDisabledAtOk() (*int32, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *ListSettings200Response) SetDisabledAt(v int32)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *ListSettings200Response) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetFirstFailAt

`func (o *ListSettings200Response) GetFirstFailAt() int32`

GetFirstFailAt returns the FirstFailAt field if non-nil, zero value otherwise.

### GetFirstFailAtOk

`func (o *ListSettings200Response) GetFirstFailAtOk() (*int32, bool)`

GetFirstFailAtOk returns a tuple with the FirstFailAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFailAt

`func (o *ListSettings200Response) SetFirstFailAt(v int32)`

SetFirstFailAt sets FirstFailAt field to given value.

### HasFirstFailAt

`func (o *ListSettings200Response) HasFirstFailAt() bool`

HasFirstFailAt returns a boolean if a field has been set.

### GetLastStatus

`func (o *ListSettings200Response) GetLastStatus() int32`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ListSettings200Response) GetLastStatusOk() (*int32, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ListSettings200Response) SetLastStatus(v int32)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ListSettings200Response) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetNextProbeAt

`func (o *ListSettings200Response) GetNextProbeAt() int32`

GetNextProbeAt returns the NextProbeAt field if non-nil, zero value otherwise.

### GetNextProbeAtOk

`func (o *ListSettings200Response) GetNextProbeAtOk() (*int32, bool)`

GetNextProbeAtOk returns a tuple with the NextProbeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextProbeAt

`func (o *ListSettings200Response) SetNextProbeAt(v int32)`

SetNextProbeAt sets NextProbeAt field to given value.

### HasNextProbeAt

`func (o *ListSettings200Response) HasNextProbeAt() bool`

HasNextProbeAt returns a boolean if a field has been set.

### GetDisabledWebhooks

`func (o *ListSettings200Response) GetDisabledWebhooks() []map[string]interface{}`

GetDisabledWebhooks returns the DisabledWebhooks field if non-nil, zero value otherwise.

### GetDisabledWebhooksOk

`func (o *ListSettings200Response) GetDisabledWebhooksOk() (*[]map[string]interface{}, bool)`

GetDisabledWebhooksOk returns a tuple with the DisabledWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledWebhooks

`func (o *ListSettings200Response) SetDisabledWebhooks(v []map[string]interface{})`

SetDisabledWebhooks sets DisabledWebhooks field to given value.

### HasDisabledWebhooks

`func (o *ListSettings200Response) HasDisabledWebhooks() bool`

HasDisabledWebhooks returns a boolean if a field has been set.

### GetWebhookStatuses

`func (o *ListSettings200Response) GetWebhookStatuses() []WebhookUrlStatus`

GetWebhookStatuses returns the WebhookStatuses field if non-nil, zero value otherwise.

### GetWebhookStatusesOk

`func (o *ListSettings200Response) GetWebhookStatusesOk() (*[]WebhookUrlStatus, bool)`

GetWebhookStatusesOk returns a tuple with the WebhookStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookStatuses

`func (o *ListSettings200Response) SetWebhookStatuses(v []WebhookUrlStatus)`

SetWebhookStatuses sets WebhookStatuses field to given value.

### HasWebhookStatuses

`func (o *ListSettings200Response) HasWebhookStatuses() bool`

HasWebhookStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


