# SendOrderDetailsRequestOrderItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetailerId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Amount** | Pointer to [**SendOrderDetailsRequestOrderItemsInnerAmount**](SendOrderDetailsRequestOrderItemsInnerAmount.md) |  | [optional] 

## Methods

### NewSendOrderDetailsRequestOrderItemsInner

`func NewSendOrderDetailsRequestOrderItemsInner() *SendOrderDetailsRequestOrderItemsInner`

NewSendOrderDetailsRequestOrderItemsInner instantiates a new SendOrderDetailsRequestOrderItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendOrderDetailsRequestOrderItemsInnerWithDefaults

`func NewSendOrderDetailsRequestOrderItemsInnerWithDefaults() *SendOrderDetailsRequestOrderItemsInner`

NewSendOrderDetailsRequestOrderItemsInnerWithDefaults instantiates a new SendOrderDetailsRequestOrderItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetailerId

`func (o *SendOrderDetailsRequestOrderItemsInner) GetRetailerId() string`

GetRetailerId returns the RetailerId field if non-nil, zero value otherwise.

### GetRetailerIdOk

`func (o *SendOrderDetailsRequestOrderItemsInner) GetRetailerIdOk() (*string, bool)`

GetRetailerIdOk returns a tuple with the RetailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailerId

`func (o *SendOrderDetailsRequestOrderItemsInner) SetRetailerId(v string)`

SetRetailerId sets RetailerId field to given value.

### HasRetailerId

`func (o *SendOrderDetailsRequestOrderItemsInner) HasRetailerId() bool`

HasRetailerId returns a boolean if a field has been set.

### GetName

`func (o *SendOrderDetailsRequestOrderItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendOrderDetailsRequestOrderItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendOrderDetailsRequestOrderItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SendOrderDetailsRequestOrderItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *SendOrderDetailsRequestOrderItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SendOrderDetailsRequestOrderItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SendOrderDetailsRequestOrderItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SendOrderDetailsRequestOrderItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAmount

`func (o *SendOrderDetailsRequestOrderItemsInner) GetAmount() SendOrderDetailsRequestOrderItemsInnerAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SendOrderDetailsRequestOrderItemsInner) GetAmountOk() (*SendOrderDetailsRequestOrderItemsInnerAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SendOrderDetailsRequestOrderItemsInner) SetAmount(v SendOrderDetailsRequestOrderItemsInnerAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SendOrderDetailsRequestOrderItemsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


