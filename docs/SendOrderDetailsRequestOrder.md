# SendOrderDetailsRequestOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Order status | [optional] 
**Items** | [**[]SendOrderDetailsRequestOrderItemsInner**](SendOrderDetailsRequestOrderItemsInner.md) | Line items | 
**Subtotal** | Pointer to [**SendOrderDetailsRequestOrderItemsInnerAmount**](SendOrderDetailsRequestOrderItemsInnerAmount.md) |  | [optional] 
**Tax** | Pointer to **map[string]interface{}** |  | [optional] 
**Shipping** | Pointer to **map[string]interface{}** |  | [optional] 
**Discount** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSendOrderDetailsRequestOrder

`func NewSendOrderDetailsRequestOrder(items []SendOrderDetailsRequestOrderItemsInner, ) *SendOrderDetailsRequestOrder`

NewSendOrderDetailsRequestOrder instantiates a new SendOrderDetailsRequestOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendOrderDetailsRequestOrderWithDefaults

`func NewSendOrderDetailsRequestOrderWithDefaults() *SendOrderDetailsRequestOrder`

NewSendOrderDetailsRequestOrderWithDefaults instantiates a new SendOrderDetailsRequestOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SendOrderDetailsRequestOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendOrderDetailsRequestOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendOrderDetailsRequestOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SendOrderDetailsRequestOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetItems

`func (o *SendOrderDetailsRequestOrder) GetItems() []SendOrderDetailsRequestOrderItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SendOrderDetailsRequestOrder) GetItemsOk() (*[]SendOrderDetailsRequestOrderItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SendOrderDetailsRequestOrder) SetItems(v []SendOrderDetailsRequestOrderItemsInner)`

SetItems sets Items field to given value.


### GetSubtotal

`func (o *SendOrderDetailsRequestOrder) GetSubtotal() SendOrderDetailsRequestOrderItemsInnerAmount`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *SendOrderDetailsRequestOrder) GetSubtotalOk() (*SendOrderDetailsRequestOrderItemsInnerAmount, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *SendOrderDetailsRequestOrder) SetSubtotal(v SendOrderDetailsRequestOrderItemsInnerAmount)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *SendOrderDetailsRequestOrder) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTax

`func (o *SendOrderDetailsRequestOrder) GetTax() map[string]interface{}`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *SendOrderDetailsRequestOrder) GetTaxOk() (*map[string]interface{}, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *SendOrderDetailsRequestOrder) SetTax(v map[string]interface{})`

SetTax sets Tax field to given value.

### HasTax

`func (o *SendOrderDetailsRequestOrder) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetShipping

`func (o *SendOrderDetailsRequestOrder) GetShipping() map[string]interface{}`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *SendOrderDetailsRequestOrder) GetShippingOk() (*map[string]interface{}, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *SendOrderDetailsRequestOrder) SetShipping(v map[string]interface{})`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *SendOrderDetailsRequestOrder) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetDiscount

`func (o *SendOrderDetailsRequestOrder) GetDiscount() map[string]interface{}`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SendOrderDetailsRequestOrder) GetDiscountOk() (*map[string]interface{}, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SendOrderDetailsRequestOrder) SetDiscount(v map[string]interface{})`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SendOrderDetailsRequestOrder) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


