# InitiateCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingProduct** | **string** | Must be &#x60;whatsapp&#x60; | 
**Action** | **string** | Call control action | 
**CallId** | Pointer to **string** | WhatsApp call id from the inbound/outbound calls webhook (&#x60;calls[].id&#x60;). Required for &#x60;pre_accept&#x60;, &#x60;accept&#x60;, &#x60;reject&#x60;, &#x60;terminate&#x60;.  | [optional] 
**To** | Pointer to **string** | Recipient WhatsApp user phone (digits, country code, no +). Required for outbound &#x60;connect&#x60;.  | [optional] 
**BizOpaqueCallbackData** | Pointer to **string** | Optional opaque string echoed on later call webhooks for correlation | [optional] 
**Session** | Pointer to [**InitiateCallRequestSession**](InitiateCallRequestSession.md) |  | [optional] 

## Methods

### NewInitiateCallRequest

`func NewInitiateCallRequest(messagingProduct string, action string, ) *InitiateCallRequest`

NewInitiateCallRequest instantiates a new InitiateCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateCallRequestWithDefaults

`func NewInitiateCallRequestWithDefaults() *InitiateCallRequest`

NewInitiateCallRequestWithDefaults instantiates a new InitiateCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingProduct

`func (o *InitiateCallRequest) GetMessagingProduct() string`

GetMessagingProduct returns the MessagingProduct field if non-nil, zero value otherwise.

### GetMessagingProductOk

`func (o *InitiateCallRequest) GetMessagingProductOk() (*string, bool)`

GetMessagingProductOk returns a tuple with the MessagingProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProduct

`func (o *InitiateCallRequest) SetMessagingProduct(v string)`

SetMessagingProduct sets MessagingProduct field to given value.


### GetAction

`func (o *InitiateCallRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InitiateCallRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InitiateCallRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetCallId

`func (o *InitiateCallRequest) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *InitiateCallRequest) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *InitiateCallRequest) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *InitiateCallRequest) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetTo

`func (o *InitiateCallRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InitiateCallRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InitiateCallRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *InitiateCallRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetBizOpaqueCallbackData

`func (o *InitiateCallRequest) GetBizOpaqueCallbackData() string`

GetBizOpaqueCallbackData returns the BizOpaqueCallbackData field if non-nil, zero value otherwise.

### GetBizOpaqueCallbackDataOk

`func (o *InitiateCallRequest) GetBizOpaqueCallbackDataOk() (*string, bool)`

GetBizOpaqueCallbackDataOk returns a tuple with the BizOpaqueCallbackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizOpaqueCallbackData

`func (o *InitiateCallRequest) SetBizOpaqueCallbackData(v string)`

SetBizOpaqueCallbackData sets BizOpaqueCallbackData field to given value.

### HasBizOpaqueCallbackData

`func (o *InitiateCallRequest) HasBizOpaqueCallbackData() bool`

HasBizOpaqueCallbackData returns a boolean if a field has been set.

### GetSession

`func (o *InitiateCallRequest) GetSession() InitiateCallRequestSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *InitiateCallRequest) GetSessionOk() (*InitiateCallRequestSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *InitiateCallRequest) SetSession(v InitiateCallRequestSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *InitiateCallRequest) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


