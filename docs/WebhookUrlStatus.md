# WebhookUrlStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | This webhook URL (same string as in &#x60;webhookUrl&#x60;) | 
**CircuitState** | **string** | Delivery state for this URL only. &#x60;ok&#x60; delivering; &#x60;degraded&#x60; recent qualifying fails (still delivering, warning email on the 2nd fail); &#x60;open&#x60; skipping most deliveries pending a probe; &#x60;half_open&#x60; probe in flight (resume only on HTTP 2xx); &#x60;disabled&#x60; sticky auto-disabled until self-serve re-enable (&#x60;POST /webhook&#x60; / &#x60;POST /settings&#x60; with this URL) or a successful probe before sticky disable.  | 
**Reason** | Pointer to **string** | Last qualifying fail reason, e.g. &#x60;http_404&#x60; or &#x60;timeout_or_network&#x60; | [optional] 
**Fails** | Pointer to **int32** | Diagnostic only — safe for clients to ignore. Number of qualifying delivery failures in the rolling ~10 minute breaker window (not lifetime, not \&quot;how many times the customer retried\&quot;). Prefer &#x60;circuitState&#x60; / &#x60;disabledWebhooks&#x60;.  | [optional] 
**Trips** | Pointer to **int32** | Diagnostic only — safe for clients to ignore. How many recovery probes already failed for this URL. Probes run on a schedule (≈15 min → 1 h → 6 h); sticky disable after 3 failed probes. The 6 h step is the last probe, not a blind \&quot;unblock in 6 hours\&quot; timer.  | [optional] 
**DisabledAt** | Pointer to **int32** | Epoch milliseconds when this URL became sticky-disabled | [optional] 
**FirstFailAt** | Pointer to **int32** | Diagnostic only. Epoch ms of the first fail in the current ~10 minute window (pairs with &#x60;fails&#x60;; not a customer-facing period).  | [optional] 
**LastStatus** | Pointer to **int32** | Last HTTP status from this URL, when known | [optional] 
**NextProbeAt** | Pointer to **int32** | Epoch milliseconds of the next half-open probe, if any | [optional] 

## Methods

### NewWebhookUrlStatus

`func NewWebhookUrlStatus(url string, circuitState string, ) *WebhookUrlStatus`

NewWebhookUrlStatus instantiates a new WebhookUrlStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUrlStatusWithDefaults

`func NewWebhookUrlStatusWithDefaults() *WebhookUrlStatus`

NewWebhookUrlStatusWithDefaults instantiates a new WebhookUrlStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookUrlStatus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookUrlStatus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookUrlStatus) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCircuitState

`func (o *WebhookUrlStatus) GetCircuitState() string`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *WebhookUrlStatus) GetCircuitStateOk() (*string, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *WebhookUrlStatus) SetCircuitState(v string)`

SetCircuitState sets CircuitState field to given value.


### GetReason

`func (o *WebhookUrlStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WebhookUrlStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WebhookUrlStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WebhookUrlStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetFails

`func (o *WebhookUrlStatus) GetFails() int32`

GetFails returns the Fails field if non-nil, zero value otherwise.

### GetFailsOk

`func (o *WebhookUrlStatus) GetFailsOk() (*int32, bool)`

GetFailsOk returns a tuple with the Fails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFails

`func (o *WebhookUrlStatus) SetFails(v int32)`

SetFails sets Fails field to given value.

### HasFails

`func (o *WebhookUrlStatus) HasFails() bool`

HasFails returns a boolean if a field has been set.

### GetTrips

`func (o *WebhookUrlStatus) GetTrips() int32`

GetTrips returns the Trips field if non-nil, zero value otherwise.

### GetTripsOk

`func (o *WebhookUrlStatus) GetTripsOk() (*int32, bool)`

GetTripsOk returns a tuple with the Trips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrips

`func (o *WebhookUrlStatus) SetTrips(v int32)`

SetTrips sets Trips field to given value.

### HasTrips

`func (o *WebhookUrlStatus) HasTrips() bool`

HasTrips returns a boolean if a field has been set.

### GetDisabledAt

`func (o *WebhookUrlStatus) GetDisabledAt() int32`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *WebhookUrlStatus) GetDisabledAtOk() (*int32, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *WebhookUrlStatus) SetDisabledAt(v int32)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *WebhookUrlStatus) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetFirstFailAt

`func (o *WebhookUrlStatus) GetFirstFailAt() int32`

GetFirstFailAt returns the FirstFailAt field if non-nil, zero value otherwise.

### GetFirstFailAtOk

`func (o *WebhookUrlStatus) GetFirstFailAtOk() (*int32, bool)`

GetFirstFailAtOk returns a tuple with the FirstFailAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFailAt

`func (o *WebhookUrlStatus) SetFirstFailAt(v int32)`

SetFirstFailAt sets FirstFailAt field to given value.

### HasFirstFailAt

`func (o *WebhookUrlStatus) HasFirstFailAt() bool`

HasFirstFailAt returns a boolean if a field has been set.

### GetLastStatus

`func (o *WebhookUrlStatus) GetLastStatus() int32`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *WebhookUrlStatus) GetLastStatusOk() (*int32, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *WebhookUrlStatus) SetLastStatus(v int32)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *WebhookUrlStatus) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetNextProbeAt

`func (o *WebhookUrlStatus) GetNextProbeAt() int32`

GetNextProbeAt returns the NextProbeAt field if non-nil, zero value otherwise.

### GetNextProbeAtOk

`func (o *WebhookUrlStatus) GetNextProbeAtOk() (*int32, bool)`

GetNextProbeAtOk returns a tuple with the NextProbeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextProbeAt

`func (o *WebhookUrlStatus) SetNextProbeAt(v int32)`

SetNextProbeAt sets NextProbeAt field to given value.

### HasNextProbeAt

`func (o *WebhookUrlStatus) HasNextProbeAt() bool`

HasNextProbeAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


