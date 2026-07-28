# ConversationalAutomation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableWelcomeMessage** | Pointer to **bool** | Enable the automatic welcome message for new conversations | [optional] 
**Prompts** | Pointer to **[]string** | Ice-breaker prompt strings (max 4, each ≤ 80 characters) | [optional] 
**Commands** | Pointer to [**[]ConversationalAutomationCommandsInner**](ConversationalAutomationCommandsInner.md) | Slash commands shown in WhatsApp | [optional] 

## Methods

### NewConversationalAutomation

`func NewConversationalAutomation() *ConversationalAutomation`

NewConversationalAutomation instantiates a new ConversationalAutomation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationalAutomationWithDefaults

`func NewConversationalAutomationWithDefaults() *ConversationalAutomation`

NewConversationalAutomationWithDefaults instantiates a new ConversationalAutomation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableWelcomeMessage

`func (o *ConversationalAutomation) GetEnableWelcomeMessage() bool`

GetEnableWelcomeMessage returns the EnableWelcomeMessage field if non-nil, zero value otherwise.

### GetEnableWelcomeMessageOk

`func (o *ConversationalAutomation) GetEnableWelcomeMessageOk() (*bool, bool)`

GetEnableWelcomeMessageOk returns a tuple with the EnableWelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWelcomeMessage

`func (o *ConversationalAutomation) SetEnableWelcomeMessage(v bool)`

SetEnableWelcomeMessage sets EnableWelcomeMessage field to given value.

### HasEnableWelcomeMessage

`func (o *ConversationalAutomation) HasEnableWelcomeMessage() bool`

HasEnableWelcomeMessage returns a boolean if a field has been set.

### GetPrompts

`func (o *ConversationalAutomation) GetPrompts() []string`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *ConversationalAutomation) GetPromptsOk() (*[]string, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *ConversationalAutomation) SetPrompts(v []string)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *ConversationalAutomation) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### GetCommands

`func (o *ConversationalAutomation) GetCommands() []ConversationalAutomationCommandsInner`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ConversationalAutomation) GetCommandsOk() (*[]ConversationalAutomationCommandsInner, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ConversationalAutomation) SetCommands(v []ConversationalAutomationCommandsInner)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *ConversationalAutomation) HasCommands() bool`

HasCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


