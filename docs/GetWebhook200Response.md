# GetWebhook200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | Pointer to [**NullableGetWebhook200ResponseWebhookUrl**](GetWebhook200ResponseWebhookUrl.md) |  | [optional] 
**HaveDisabledWebhooks** | Pointer to **bool** | &#x60;true&#x60; when at least one URL in &#x60;webhookUrl&#x60; is sticky auto-disabled. Not \&quot;the channel is dead\&quot;: other configured URLs still receive events. Send/receive are unaffected.  | [optional] 
**CircuitState** | Pointer to **string** | Compatibility scalar for **one** URL (first sticky-disabled, else first configured). Not the channel as a whole. Prefer &#x60;webhookStatuses[].circuitState&#x60;.  | [optional] 
**Reason** | Pointer to **string** | Last qualifying fail reason of the summary URL (same URL as &#x60;circuitState&#x60;) | [optional] 
**Fails** | Pointer to **int32** | Diagnostic only (summary URL). Qualifying failures in the rolling ~10 minute breaker window — not lifetime and not customer retries. Prefer &#x60;webhookStatuses&#x60; / &#x60;disabledWebhooks&#x60;.  | [optional] 
**Trips** | Pointer to **int32** | Diagnostic only (summary URL). Failed recovery probes (≈15 min → 1 h → 6 h). Sticky disable after 3. Not a blind unblock timer.  | [optional] 
**DisabledAt** | Pointer to **int32** | Epoch milliseconds when the summary URL became sticky-disabled | [optional] 
**FirstFailAt** | Pointer to **int32** | Diagnostic only (summary URL). Epoch ms of the first fail in the current ~10 minute window (pairs with &#x60;fails&#x60;).  | [optional] 
**LastStatus** | Pointer to **int32** | Last HTTP status from the summary URL, when known | [optional] 
**NextProbeAt** | Pointer to **int32** | Epoch milliseconds of the next half-open probe for the summary URL, if any | [optional] 
**DisabledWebhooks** | Pointer to **[]map[string]interface{}** | Currently configured URLs that are sticky auto-disabled (subset of &#x60;webhookUrl&#x60;). Does not include leftover URLs that are no longer configured.  | [optional] 
**WebhookStatuses** | Pointer to [**[]WebhookUrlStatus**](WebhookUrlStatus.md) | One entry per currently configured URL. Prefer this over top-level &#x60;haveDisabledWebhooks&#x60; / &#x60;circuitState&#x60; when more than one webhook is set.  | [optional] 

## Methods

### NewGetWebhook200Response

`func NewGetWebhook200Response() *GetWebhook200Response`

NewGetWebhook200Response instantiates a new GetWebhook200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhook200ResponseWithDefaults

`func NewGetWebhook200ResponseWithDefaults() *GetWebhook200Response`

NewGetWebhook200ResponseWithDefaults instantiates a new GetWebhook200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *GetWebhook200Response) GetWebhookUrl() GetWebhook200ResponseWebhookUrl`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *GetWebhook200Response) GetWebhookUrlOk() (*GetWebhook200ResponseWebhookUrl, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *GetWebhook200Response) SetWebhookUrl(v GetWebhook200ResponseWebhookUrl)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *GetWebhook200Response) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *GetWebhook200Response) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *GetWebhook200Response) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetHaveDisabledWebhooks

`func (o *GetWebhook200Response) GetHaveDisabledWebhooks() bool`

GetHaveDisabledWebhooks returns the HaveDisabledWebhooks field if non-nil, zero value otherwise.

### GetHaveDisabledWebhooksOk

`func (o *GetWebhook200Response) GetHaveDisabledWebhooksOk() (*bool, bool)`

GetHaveDisabledWebhooksOk returns a tuple with the HaveDisabledWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaveDisabledWebhooks

`func (o *GetWebhook200Response) SetHaveDisabledWebhooks(v bool)`

SetHaveDisabledWebhooks sets HaveDisabledWebhooks field to given value.

### HasHaveDisabledWebhooks

`func (o *GetWebhook200Response) HasHaveDisabledWebhooks() bool`

HasHaveDisabledWebhooks returns a boolean if a field has been set.

### GetCircuitState

`func (o *GetWebhook200Response) GetCircuitState() string`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *GetWebhook200Response) GetCircuitStateOk() (*string, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *GetWebhook200Response) SetCircuitState(v string)`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *GetWebhook200Response) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### GetReason

`func (o *GetWebhook200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetWebhook200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetWebhook200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetWebhook200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetFails

`func (o *GetWebhook200Response) GetFails() int32`

GetFails returns the Fails field if non-nil, zero value otherwise.

### GetFailsOk

`func (o *GetWebhook200Response) GetFailsOk() (*int32, bool)`

GetFailsOk returns a tuple with the Fails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFails

`func (o *GetWebhook200Response) SetFails(v int32)`

SetFails sets Fails field to given value.

### HasFails

`func (o *GetWebhook200Response) HasFails() bool`

HasFails returns a boolean if a field has been set.

### GetTrips

`func (o *GetWebhook200Response) GetTrips() int32`

GetTrips returns the Trips field if non-nil, zero value otherwise.

### GetTripsOk

`func (o *GetWebhook200Response) GetTripsOk() (*int32, bool)`

GetTripsOk returns a tuple with the Trips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrips

`func (o *GetWebhook200Response) SetTrips(v int32)`

SetTrips sets Trips field to given value.

### HasTrips

`func (o *GetWebhook200Response) HasTrips() bool`

HasTrips returns a boolean if a field has been set.

### GetDisabledAt

`func (o *GetWebhook200Response) GetDisabledAt() int32`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GetWebhook200Response) GetDisabledAtOk() (*int32, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GetWebhook200Response) SetDisabledAt(v int32)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GetWebhook200Response) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetFirstFailAt

`func (o *GetWebhook200Response) GetFirstFailAt() int32`

GetFirstFailAt returns the FirstFailAt field if non-nil, zero value otherwise.

### GetFirstFailAtOk

`func (o *GetWebhook200Response) GetFirstFailAtOk() (*int32, bool)`

GetFirstFailAtOk returns a tuple with the FirstFailAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFailAt

`func (o *GetWebhook200Response) SetFirstFailAt(v int32)`

SetFirstFailAt sets FirstFailAt field to given value.

### HasFirstFailAt

`func (o *GetWebhook200Response) HasFirstFailAt() bool`

HasFirstFailAt returns a boolean if a field has been set.

### GetLastStatus

`func (o *GetWebhook200Response) GetLastStatus() int32`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *GetWebhook200Response) GetLastStatusOk() (*int32, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *GetWebhook200Response) SetLastStatus(v int32)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *GetWebhook200Response) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetNextProbeAt

`func (o *GetWebhook200Response) GetNextProbeAt() int32`

GetNextProbeAt returns the NextProbeAt field if non-nil, zero value otherwise.

### GetNextProbeAtOk

`func (o *GetWebhook200Response) GetNextProbeAtOk() (*int32, bool)`

GetNextProbeAtOk returns a tuple with the NextProbeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextProbeAt

`func (o *GetWebhook200Response) SetNextProbeAt(v int32)`

SetNextProbeAt sets NextProbeAt field to given value.

### HasNextProbeAt

`func (o *GetWebhook200Response) HasNextProbeAt() bool`

HasNextProbeAt returns a boolean if a field has been set.

### GetDisabledWebhooks

`func (o *GetWebhook200Response) GetDisabledWebhooks() []map[string]interface{}`

GetDisabledWebhooks returns the DisabledWebhooks field if non-nil, zero value otherwise.

### GetDisabledWebhooksOk

`func (o *GetWebhook200Response) GetDisabledWebhooksOk() (*[]map[string]interface{}, bool)`

GetDisabledWebhooksOk returns a tuple with the DisabledWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledWebhooks

`func (o *GetWebhook200Response) SetDisabledWebhooks(v []map[string]interface{})`

SetDisabledWebhooks sets DisabledWebhooks field to given value.

### HasDisabledWebhooks

`func (o *GetWebhook200Response) HasDisabledWebhooks() bool`

HasDisabledWebhooks returns a boolean if a field has been set.

### GetWebhookStatuses

`func (o *GetWebhook200Response) GetWebhookStatuses() []WebhookUrlStatus`

GetWebhookStatuses returns the WebhookStatuses field if non-nil, zero value otherwise.

### GetWebhookStatusesOk

`func (o *GetWebhook200Response) GetWebhookStatusesOk() (*[]WebhookUrlStatus, bool)`

GetWebhookStatusesOk returns a tuple with the WebhookStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookStatuses

`func (o *GetWebhook200Response) SetWebhookStatuses(v []WebhookUrlStatus)`

SetWebhookStatuses sets WebhookStatuses field to given value.

### HasWebhookStatuses

`func (o *GetWebhook200Response) HasWebhookStatuses() bool`

HasWebhookStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


