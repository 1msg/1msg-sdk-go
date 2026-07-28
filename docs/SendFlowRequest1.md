# SendFlowRequest1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** |  | 
**Header** | Pointer to [**SendFlowRequest1Header**](SendFlowRequest1Header.md) |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**FlowId** | **string** |  | 
**FlowToken** | **string** |  | 
**FlowCta** | **string** |  | 
**FlowAction** | Pointer to **string** |  | [optional] 
**FlowActionPayload** | Pointer to **map[string]interface{}** |  | [optional] 
**FlowMessageVersion** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**FlowActionData** | Pointer to **map[string]interface{}** |  | [optional] 
**FlowActionScreen** | Pointer to **string** |  | [optional] 
**QuotedMsgId** | Pointer to **string** |  | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **int32** |  | [optional] 

## Methods

### NewSendFlowRequest1

`func NewSendFlowRequest1(body string, flowId string, flowToken string, flowCta string, ) *SendFlowRequest1`

NewSendFlowRequest1 instantiates a new SendFlowRequest1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFlowRequest1WithDefaults

`func NewSendFlowRequest1WithDefaults() *SendFlowRequest1`

NewSendFlowRequest1WithDefaults instantiates a new SendFlowRequest1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *SendFlowRequest1) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendFlowRequest1) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendFlowRequest1) SetBody(v string)`

SetBody sets Body field to given value.


### GetHeader

`func (o *SendFlowRequest1) GetHeader() SendFlowRequest1Header`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SendFlowRequest1) GetHeaderOk() (*SendFlowRequest1Header, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SendFlowRequest1) SetHeader(v SendFlowRequest1Header)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SendFlowRequest1) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetFooter

`func (o *SendFlowRequest1) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *SendFlowRequest1) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *SendFlowRequest1) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *SendFlowRequest1) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetFlowId

`func (o *SendFlowRequest1) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *SendFlowRequest1) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *SendFlowRequest1) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowToken

`func (o *SendFlowRequest1) GetFlowToken() string`

GetFlowToken returns the FlowToken field if non-nil, zero value otherwise.

### GetFlowTokenOk

`func (o *SendFlowRequest1) GetFlowTokenOk() (*string, bool)`

GetFlowTokenOk returns a tuple with the FlowToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowToken

`func (o *SendFlowRequest1) SetFlowToken(v string)`

SetFlowToken sets FlowToken field to given value.


### GetFlowCta

`func (o *SendFlowRequest1) GetFlowCta() string`

GetFlowCta returns the FlowCta field if non-nil, zero value otherwise.

### GetFlowCtaOk

`func (o *SendFlowRequest1) GetFlowCtaOk() (*string, bool)`

GetFlowCtaOk returns a tuple with the FlowCta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCta

`func (o *SendFlowRequest1) SetFlowCta(v string)`

SetFlowCta sets FlowCta field to given value.


### GetFlowAction

`func (o *SendFlowRequest1) GetFlowAction() string`

GetFlowAction returns the FlowAction field if non-nil, zero value otherwise.

### GetFlowActionOk

`func (o *SendFlowRequest1) GetFlowActionOk() (*string, bool)`

GetFlowActionOk returns a tuple with the FlowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAction

`func (o *SendFlowRequest1) SetFlowAction(v string)`

SetFlowAction sets FlowAction field to given value.

### HasFlowAction

`func (o *SendFlowRequest1) HasFlowAction() bool`

HasFlowAction returns a boolean if a field has been set.

### GetFlowActionPayload

`func (o *SendFlowRequest1) GetFlowActionPayload() map[string]interface{}`

GetFlowActionPayload returns the FlowActionPayload field if non-nil, zero value otherwise.

### GetFlowActionPayloadOk

`func (o *SendFlowRequest1) GetFlowActionPayloadOk() (*map[string]interface{}, bool)`

GetFlowActionPayloadOk returns a tuple with the FlowActionPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowActionPayload

`func (o *SendFlowRequest1) SetFlowActionPayload(v map[string]interface{})`

SetFlowActionPayload sets FlowActionPayload field to given value.

### HasFlowActionPayload

`func (o *SendFlowRequest1) HasFlowActionPayload() bool`

HasFlowActionPayload returns a boolean if a field has been set.

### GetFlowMessageVersion

`func (o *SendFlowRequest1) GetFlowMessageVersion() string`

GetFlowMessageVersion returns the FlowMessageVersion field if non-nil, zero value otherwise.

### GetFlowMessageVersionOk

`func (o *SendFlowRequest1) GetFlowMessageVersionOk() (*string, bool)`

GetFlowMessageVersionOk returns a tuple with the FlowMessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowMessageVersion

`func (o *SendFlowRequest1) SetFlowMessageVersion(v string)`

SetFlowMessageVersion sets FlowMessageVersion field to given value.

### HasFlowMessageVersion

`func (o *SendFlowRequest1) HasFlowMessageVersion() bool`

HasFlowMessageVersion returns a boolean if a field has been set.

### GetMode

`func (o *SendFlowRequest1) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SendFlowRequest1) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SendFlowRequest1) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SendFlowRequest1) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetFlowActionData

`func (o *SendFlowRequest1) GetFlowActionData() map[string]interface{}`

GetFlowActionData returns the FlowActionData field if non-nil, zero value otherwise.

### GetFlowActionDataOk

`func (o *SendFlowRequest1) GetFlowActionDataOk() (*map[string]interface{}, bool)`

GetFlowActionDataOk returns a tuple with the FlowActionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowActionData

`func (o *SendFlowRequest1) SetFlowActionData(v map[string]interface{})`

SetFlowActionData sets FlowActionData field to given value.

### HasFlowActionData

`func (o *SendFlowRequest1) HasFlowActionData() bool`

HasFlowActionData returns a boolean if a field has been set.

### GetFlowActionScreen

`func (o *SendFlowRequest1) GetFlowActionScreen() string`

GetFlowActionScreen returns the FlowActionScreen field if non-nil, zero value otherwise.

### GetFlowActionScreenOk

`func (o *SendFlowRequest1) GetFlowActionScreenOk() (*string, bool)`

GetFlowActionScreenOk returns a tuple with the FlowActionScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowActionScreen

`func (o *SendFlowRequest1) SetFlowActionScreen(v string)`

SetFlowActionScreen sets FlowActionScreen field to given value.

### HasFlowActionScreen

`func (o *SendFlowRequest1) HasFlowActionScreen() bool`

HasFlowActionScreen returns a boolean if a field has been set.

### GetQuotedMsgId

`func (o *SendFlowRequest1) GetQuotedMsgId() string`

GetQuotedMsgId returns the QuotedMsgId field if non-nil, zero value otherwise.

### GetQuotedMsgIdOk

`func (o *SendFlowRequest1) GetQuotedMsgIdOk() (*string, bool)`

GetQuotedMsgIdOk returns a tuple with the QuotedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedMsgId

`func (o *SendFlowRequest1) SetQuotedMsgId(v string)`

SetQuotedMsgId sets QuotedMsgId field to given value.

### HasQuotedMsgId

`func (o *SendFlowRequest1) HasQuotedMsgId() bool`

HasQuotedMsgId returns a boolean if a field has been set.

### GetChatId

`func (o *SendFlowRequest1) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendFlowRequest1) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendFlowRequest1) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *SendFlowRequest1) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetPhone

`func (o *SendFlowRequest1) GetPhone() int32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendFlowRequest1) GetPhoneOk() (*int32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendFlowRequest1) SetPhone(v int32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendFlowRequest1) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


