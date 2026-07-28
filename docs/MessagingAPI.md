# \MessagingAPI

All URIs are relative to *https://api.1msg.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReadMessage**](MessagingAPI.md#CreateReadMessage) | **Post** /readMessage | Mark message as read
[**CreateUploadMedia**](MessagingAPI.md#CreateUploadMedia) | **Post** /uploadMedia | Upload media
[**DeleteMedia**](MessagingAPI.md#DeleteMedia) | **Delete** /media/{mediaId} | Delete media from WABA storage
[**DeleteMediaLegacy**](MessagingAPI.md#DeleteMediaLegacy) | **Post** /deleteMedia | Delete media from WABA storage (deprecated alias)
[**GetMmLiteStatus**](MessagingAPI.md#GetMmLiteStatus) | **Get** /mmLiteStatus | Get MM Lite availability and status
[**ListMessages**](MessagingAPI.md#ListMessages) | **Get** /messages | Get messages list
[**RetrieveMedia**](MessagingAPI.md#RetrieveMedia) | **Get** /retrieveMedia | Retrieve uploaded media metadata
[**SendAddressMessage**](MessagingAPI.md#SendAddressMessage) | **Post** /sendAddressMessage | Send address request message
[**SendButton**](MessagingAPI.md#SendButton) | **Post** /sendButton | Send Reply Buttons Message
[**SendCarousel**](MessagingAPI.md#SendCarousel) | **Post** /sendCarousel | Send Carousel
[**SendContact**](MessagingAPI.md#SendContact) | **Post** /sendContact | Send a Contact
[**SendCtaUrl**](MessagingAPI.md#SendCtaUrl) | **Post** /sendCtaUrl | Send CTA URL interactive message
[**SendFile**](MessagingAPI.md#SendFile) | **Post** /sendFile | Send a File
[**SendFlow**](MessagingAPI.md#SendFlow) | **Post** /sendFlow | Send WhatsApp Flow Message
[**SendList**](MessagingAPI.md#SendList) | **Post** /sendList | Send List Message
[**SendLocation**](MessagingAPI.md#SendLocation) | **Post** /sendLocation | Send a Location
[**SendLocationRequest**](MessagingAPI.md#SendLocationRequest) | **Post** /sendLocationRequest | Send Location Request Message
[**SendMessage**](MessagingAPI.md#SendMessage) | **Post** /sendMessage | Send a Message
[**SendOrderDetails**](MessagingAPI.md#SendOrderDetails) | **Post** /sendOrderDetails | Send order details (India payments template)
[**SendPaymentRequest**](MessagingAPI.md#SendPaymentRequest) | **Post** /sendPaymentRequest | Send payment request (regional)
[**SendProduct**](MessagingAPI.md#SendProduct) | **Post** /sendProduct | Send a Product
[**SendReaction**](MessagingAPI.md#SendReaction) | **Post** /sendReaction | Send Reaction
[**SendSticker**](MessagingAPI.md#SendSticker) | **Post** /sendSticker | Send sticker message



## CreateReadMessage

> CreateReadMessage200Response CreateReadMessage(ctx).Token(token).MessageId(messageId).MsgId(msgId).TypingIndicator(typingIndicator).Execute()

Mark message as read

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	messageId := "messageId_example" // string |  (optional)
	msgId := "msgId_example" // string |  (optional)
	typingIndicator := true // bool | Show WhatsApp typing indicator (max 25s or until reply) (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.CreateReadMessage(context.Background()).Token(token).MessageId(messageId).MsgId(msgId).TypingIndicator(typingIndicator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.CreateReadMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReadMessage`: CreateReadMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.CreateReadMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReadMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **messageId** | **string** |  | 
 **msgId** | **string** |  | 
 **typingIndicator** | **bool** | Show WhatsApp typing indicator (max 25s or until reply) | [default to false]

### Return type

[**CreateReadMessage200Response**](CreateReadMessage200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUploadMedia

> CreateUploadMedia(ctx).Token(token).RequestBody(requestBody).Execute()

Upload media

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MessagingAPI.CreateUploadMedia(context.Background()).Token(token).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.CreateUploadMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **requestBody** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMedia

> SuccessResponse DeleteMedia(ctx, mediaId).Token(token).Execute()

Delete media from WABA storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	mediaId := "mediaId_example" // string | Numeric WABA media id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.DeleteMedia(context.Background(), mediaId).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.DeleteMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMedia`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.DeleteMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Numeric WABA media id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMediaLegacy

> SuccessResponse DeleteMediaLegacy(ctx).Token(token).DeleteMediaLegacyRequest(deleteMediaLegacyRequest).Execute()

Delete media from WABA storage (deprecated alias)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	deleteMediaLegacyRequest := *openapiclient.NewDeleteMediaLegacyRequest("MediaId_example") // DeleteMediaLegacyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.DeleteMediaLegacy(context.Background()).Token(token).DeleteMediaLegacyRequest(deleteMediaLegacyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.DeleteMediaLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMediaLegacy`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.DeleteMediaLegacy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **deleteMediaLegacyRequest** | [**DeleteMediaLegacyRequest**](DeleteMediaLegacyRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMmLiteStatus

> GetMmLiteStatus200Response GetMmLiteStatus(ctx).Token(token).Execute()

Get MM Lite availability and status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.GetMmLiteStatus(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.GetMmLiteStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMmLiteStatus`: GetMmLiteStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.GetMmLiteStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMmLiteStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**GetMmLiteStatus200Response**](GetMmLiteStatus200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessages

> ListMessages200Response ListMessages(ctx).Token(token).Execute()

Get messages list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.ListMessages(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.ListMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessages`: ListMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.ListMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**ListMessages200Response**](ListMessages200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMedia

> RetrieveMedia200Response RetrieveMedia(ctx).Token(token).MediaId(mediaId).Execute()

Retrieve uploaded media metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	mediaId := "mediaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.RetrieveMedia(context.Background()).Token(token).MediaId(mediaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.RetrieveMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveMedia`: RetrieveMedia200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.RetrieveMedia`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **mediaId** | **string** |  | 

### Return type

[**RetrieveMedia200Response**](RetrieveMedia200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendAddressMessage

> MessageSentResponse SendAddressMessage(ctx).Token(token).SendAddressMessageRequest(sendAddressMessageRequest).Execute()

Send address request message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendAddressMessageRequest := *openapiclient.NewSendAddressMessageRequest("Body_example") // SendAddressMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendAddressMessage(context.Background()).Token(token).SendAddressMessageRequest(sendAddressMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendAddressMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendAddressMessage`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendAddressMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendAddressMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendAddressMessageRequest** | [**SendAddressMessageRequest**](SendAddressMessageRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendButton

> MessageSentResponse SendButton(ctx).Token(token).SendButtonRequest(sendButtonRequest).Execute()

Send Reply Buttons Message

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendButtonRequest := *openapiclient.NewSendButtonRequest() // SendButtonRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendButton(context.Background()).Token(token).SendButtonRequest(sendButtonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendButton``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendButton`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendButton`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendButtonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendButtonRequest** | [**SendButtonRequest**](SendButtonRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCarousel

> MessageSentResponse SendCarousel(ctx).Token(token).Params(params).Body(body).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()

Send Carousel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	params := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Required. Template-like structure (same as sendTemplate params). Must include a CAROUSEL component and its cards.  Structure: - params[] item with type CAROUSEL and cards (2..10) - cards[] item:   - components[] items:     - header: parameters[0] with type image/video/product       - image: {\\\"image\\\": {\\\"link\\\": \\\"https://...\\\"}}       - video: {\\\"video\\\": {\\\"link\\\": \\\"https://...\\\"}}       - product: {\\\"product\\\": {\\\"catalog_id\\\": \\\"...\\\", \\\"product_retailer_id\\\": \\\"...\\\"}}     - button: sub_type url or quick_reply       - url: parameters {\\\"text\\\": \\\"Button\\\", \\\"url\\\": \\\"https://...\\\"}       - quick_reply: parameters[] of {\\\"id\\\": \\\"...\\\", \\\"text\\\": \\\"...\\\"} 
	body := "body_example" // string | Text shown above the carousel. Optional. If omitted and params include a body component, the body will be taken from params. (optional)
	quotedMsgId := "quotedMsgId_example" // string | Quoted message ID (Cloud API) (optional)
	chatId := "chatId_example" // string | Required if phone is not set. Chat ID from the message list. Examples: 12020721369@c.us or 120363046942338209@g.us(group) (optional)
	phone := int32(56) // int32 | Required if chatId is not set. A phone number starting with the country code. USA example: 12020721369. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendCarousel(context.Background()).Token(token).Params(params).Body(body).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendCarousel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendCarousel`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendCarousel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendCarouselRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **params** | **[]map[string]interface{}** | Required. Template-like structure (same as sendTemplate params). Must include a CAROUSEL component and its cards.  Structure: - params[] item with type CAROUSEL and cards (2..10) - cards[] item:   - components[] items:     - header: parameters[0] with type image/video/product       - image: {\\\&quot;image\\\&quot;: {\\\&quot;link\\\&quot;: \\\&quot;https://...\\\&quot;}}       - video: {\\\&quot;video\\\&quot;: {\\\&quot;link\\\&quot;: \\\&quot;https://...\\\&quot;}}       - product: {\\\&quot;product\\\&quot;: {\\\&quot;catalog_id\\\&quot;: \\\&quot;...\\\&quot;, \\\&quot;product_retailer_id\\\&quot;: \\\&quot;...\\\&quot;}}     - button: sub_type url or quick_reply       - url: parameters {\\\&quot;text\\\&quot;: \\\&quot;Button\\\&quot;, \\\&quot;url\\\&quot;: \\\&quot;https://...\\\&quot;}       - quick_reply: parameters[] of {\\\&quot;id\\\&quot;: \\\&quot;...\\\&quot;, \\\&quot;text\\\&quot;: \\\&quot;...\\\&quot;}  | 
 **body** | **string** | Text shown above the carousel. Optional. If omitted and params include a body component, the body will be taken from params. | 
 **quotedMsgId** | **string** | Quoted message ID (Cloud API) | 
 **chatId** | **string** | Required if phone is not set. Chat ID from the message list. Examples: 12020721369@c.us or 120363046942338209@g.us(group) | 
 **phone** | **int32** | Required if chatId is not set. A phone number starting with the country code. USA example: 12020721369. | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendContact

> MessageSentResponse SendContact(ctx).Token(token).SendContactRequest(sendContactRequest).Execute()

Send a Contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendContactRequest := *openapiclient.NewSendContactRequest() // SendContactRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendContact(context.Background()).Token(token).SendContactRequest(sendContactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendContact`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendContactRequest** | [**SendContactRequest**](SendContactRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCtaUrl

> MessageSentResponse SendCtaUrl(ctx).Token(token).SendCtaUrlRequest(sendCtaUrlRequest).Execute()

Send CTA URL interactive message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendCtaUrlRequest := *openapiclient.NewSendCtaUrlRequest("Body_example", "DisplayText_example", "Url_example") // SendCtaUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendCtaUrl(context.Background()).Token(token).SendCtaUrlRequest(sendCtaUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendCtaUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendCtaUrl`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendCtaUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendCtaUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendCtaUrlRequest** | [**SendCtaUrlRequest**](SendCtaUrlRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendFile

> MessageSentResponse SendFile(ctx).Token(token).Body(body).Filename(filename).MediaId(mediaId).MediaType(mediaType).Voice(voice).Caption(caption).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()

Send a File



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	body := "body_example" // string | File source (required unless mediaId is set): - HTTP/HTTPS URL - Base64 data URI  (optional)
	filename := "filename_example" // string | File name with extension (required with body, not with mediaId) (optional)
	mediaId := "mediaId_example" // string | WABA media id from uploadMedia (alternative to body) (optional)
	mediaType := "mediaType_example" // string | Required when using mediaId (optional)
	voice := true // bool | Native WhatsApp voice note (audio/ogg only) (optional)
	caption := "caption_example" // string |  (optional)
	quotedMsgId := "quotedMsgId_example" // string |  (optional)
	chatId := "chatId_example" // string |  (optional)
	phone := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendFile(context.Background()).Token(token).Body(body).Filename(filename).MediaId(mediaId).MediaType(mediaType).Voice(voice).Caption(caption).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFile`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **body** | **string** | File source (required unless mediaId is set): - HTTP/HTTPS URL - Base64 data URI  | 
 **filename** | **string** | File name with extension (required with body, not with mediaId) | 
 **mediaId** | **string** | WABA media id from uploadMedia (alternative to body) | 
 **mediaType** | **string** | Required when using mediaId | 
 **voice** | **bool** | Native WhatsApp voice note (audio/ogg only) | 
 **caption** | **string** |  | 
 **quotedMsgId** | **string** |  | 
 **chatId** | **string** |  | 
 **phone** | **int32** |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendFlow

> MessageSentResponse SendFlow(ctx).Token(token).Body(body).FlowId(flowId).FlowToken(flowToken).FlowCta(flowCta).Header(header).Footer(footer).FlowAction(flowAction).FlowActionPayload(flowActionPayload).FlowMessageVersion(flowMessageVersion).Mode(mode).FlowActionData(flowActionData).FlowActionScreen(flowActionScreen).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()

Send WhatsApp Flow Message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	body := "body_example" // string | Flow message body text
	flowId := "flowId_example" // string | Published Flow ID
	flowToken := "flowToken_example" // string | Flow token generated by the business
	flowCta := "flowCta_example" // string | CTA button text
	header := openapiclient.sendFlow_request_header{MapmapOfStringAny: new(map[string]interface{})} // SendFlowRequestHeader |  (optional)
	footer := "footer_example" // string | Footer text (optional)
	flowAction := "flowAction_example" // string | Flow action type (optional)
	flowActionPayload := map[string]interface{}{ ... } // map[string]interface{} | Required for flowAction=navigate (screen is required). Ignored for data_exchange. If data is provided, it must be a non-empty object. (optional)
	flowMessageVersion := "flowMessageVersion_example" // string | Flow message version (default \\\"3\\\") (optional)
	mode := "mode_example" // string | Flow mode (draft or published). If omitted, provider default applies (optional)
	flowActionData := map[string]interface{}{ ... } // map[string]interface{} | Shortcut for flowActionPayload.data (optional) (optional)
	flowActionScreen := "flowActionScreen_example" // string | Shortcut for flowActionPayload.screen (optional) (optional)
	quotedMsgId := "quotedMsgId_example" // string | Quoted message ID (Cloud API) (optional)
	chatId := "chatId_example" // string | Required if phone is not set. Chat ID from the message list. Examples: 12020721369@c.us or 120363046942338209@g.us(group) (optional)
	phone := int32(56) // int32 | Required if chatId is not set. A phone number starting with the country code. USA example: 12020721369. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendFlow(context.Background()).Token(token).Body(body).FlowId(flowId).FlowToken(flowToken).FlowCta(flowCta).Header(header).Footer(footer).FlowAction(flowAction).FlowActionPayload(flowActionPayload).FlowMessageVersion(flowMessageVersion).Mode(mode).FlowActionData(flowActionData).FlowActionScreen(flowActionScreen).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFlow`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **body** | **string** | Flow message body text | 
 **flowId** | **string** | Published Flow ID | 
 **flowToken** | **string** | Flow token generated by the business | 
 **flowCta** | **string** | CTA button text | 
 **header** | [**SendFlowRequestHeader**](SendFlowRequestHeader.md) |  | 
 **footer** | **string** | Footer text | 
 **flowAction** | **string** | Flow action type | 
 **flowActionPayload** | [**map[string]interface{}**](map[string]interface{}.md) | Required for flowAction&#x3D;navigate (screen is required). Ignored for data_exchange. If data is provided, it must be a non-empty object. | 
 **flowMessageVersion** | **string** | Flow message version (default \\\&quot;3\\\&quot;) | 
 **mode** | **string** | Flow mode (draft or published). If omitted, provider default applies | 
 **flowActionData** | [**map[string]interface{}**](map[string]interface{}.md) | Shortcut for flowActionPayload.data (optional) | 
 **flowActionScreen** | **string** | Shortcut for flowActionPayload.screen (optional) | 
 **quotedMsgId** | **string** | Quoted message ID (Cloud API) | 
 **chatId** | **string** | Required if phone is not set. Chat ID from the message list. Examples: 12020721369@c.us or 120363046942338209@g.us(group) | 
 **phone** | **int32** | Required if chatId is not set. A phone number starting with the country code. USA example: 12020721369. | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendList

> MessageSentResponse SendList(ctx).Token(token).SendListRequest(sendListRequest).Execute()

Send List Message

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendListRequest := *openapiclient.NewSendListRequest() // SendListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendList(context.Background()).Token(token).SendListRequest(sendListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendList`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendListRequest** | [**SendListRequest**](SendListRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLocation

> MessageSentResponse SendLocation(ctx).Token(token).Lat(lat).Lng(lng).Address(address).Name(name).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()

Send a Location



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	lat := "lat_example" // string | Latitude of the location. Example: 45.018337
	lng := "lng_example" // string | Longitude of the location. Example: -73.968285
	address := "address_example" // string | Address of the location. Only displayed if name is present. Example: 9766 Valley View St., New York, NY 10024 (optional)
	name := "name_example" // string | Name of the location. Example: Facebook HQ (optional)
	quotedMsgId := "quotedMsgId_example" // string | Quoted message ID (Cloud API) (optional)
	chatId := "chatId_example" // string | Required if phone is not set. Chat ID from the message list. Examples: 12020721369@c.us or 120363046942338209@g.us(group) (optional)
	phone := int32(56) // int32 | Required if chatId is not set. A phone number starting with the country code. USA example: 12020721369. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendLocation(context.Background()).Token(token).Lat(lat).Lng(lng).Address(address).Name(name).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendLocation`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **lat** | **string** | Latitude of the location. Example: 45.018337 | 
 **lng** | **string** | Longitude of the location. Example: -73.968285 | 
 **address** | **string** | Address of the location. Only displayed if name is present. Example: 9766 Valley View St., New York, NY 10024 | 
 **name** | **string** | Name of the location. Example: Facebook HQ | 
 **quotedMsgId** | **string** | Quoted message ID (Cloud API) | 
 **chatId** | **string** | Required if phone is not set. Chat ID from the message list. Examples: 12020721369@c.us or 120363046942338209@g.us(group) | 
 **phone** | **int32** | Required if chatId is not set. A phone number starting with the country code. USA example: 12020721369. | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLocationRequest

> MessageSentResponse SendLocationRequest(ctx).Token(token).SendLocationRequestRequest(sendLocationRequestRequest).Execute()

Send Location Request Message

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendLocationRequestRequest := *openapiclient.NewSendLocationRequestRequest() // SendLocationRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendLocationRequest(context.Background()).Token(token).SendLocationRequestRequest(sendLocationRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendLocationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendLocationRequest`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendLocationRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendLocationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendLocationRequestRequest** | [**SendLocationRequestRequest**](SendLocationRequestRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> MessageSentResponse SendMessage(ctx).Token(token).Body(body).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()

Send a Message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	body := "body_example" // string | Message text, UTF-8 or UTF-16 string with emoji
	quotedMsgId := "quotedMsgId_example" // string | Message ID to quote/reply to (Cloud API wamid) (optional)
	chatId := "chatId_example" // string | Chat ID in format: phone@c.us (individual) or phone@g.us (group) (optional)
	phone := int32(56) // int32 | Phone number starting with country code (alternative to chatId) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendMessage(context.Background()).Token(token).Body(body).QuotedMsgId(quotedMsgId).ChatId(chatId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMessage`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **body** | **string** | Message text, UTF-8 or UTF-16 string with emoji | 
 **quotedMsgId** | **string** | Message ID to quote/reply to (Cloud API wamid) | 
 **chatId** | **string** | Chat ID in format: phone@c.us (individual) or phone@g.us (group) | 
 **phone** | **int32** | Phone number starting with country code (alternative to chatId) | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendOrderDetails

> MessageSentResponse SendOrderDetails(ctx).Token(token).SendOrderDetailsRequest(sendOrderDetailsRequest).Execute()

Send order details (India payments template)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendOrderDetailsRequest := *openapiclient.NewSendOrderDetailsRequest("order_details_utility", "your_namespace_uuid", *openapiclient.NewSendOrderDetailsRequestLanguage("en"), *openapiclient.NewSendOrderDetailsRequestOrder([]openapiclient.SendOrderDetailsRequestOrderItemsInner{*openapiclient.NewSendOrderDetailsRequestOrderItemsInner()})) // SendOrderDetailsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendOrderDetails(context.Background()).Token(token).SendOrderDetailsRequest(sendOrderDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendOrderDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendOrderDetails`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendOrderDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendOrderDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendOrderDetailsRequest** | [**SendOrderDetailsRequest**](SendOrderDetailsRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPaymentRequest

> MessageSentResponse SendPaymentRequest(ctx).Token(token).SendPaymentRequestRequest(sendPaymentRequestRequest).Execute()

Send payment request (regional)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendPaymentRequestRequest := *openapiclient.NewSendPaymentRequestRequest(int32(123), "Region_example") // SendPaymentRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendPaymentRequest(context.Background()).Token(token).SendPaymentRequestRequest(sendPaymentRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendPaymentRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPaymentRequest`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendPaymentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendPaymentRequestRequest** | [**SendPaymentRequestRequest**](SendPaymentRequestRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendProduct

> MessageSentResponse SendProduct(ctx).Token(token).RequestBody(requestBody).Execute()

Send a Product

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendProduct(context.Background()).Token(token).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendProduct`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendReaction

> MessageSentResponse SendReaction(ctx).Token(token).SendReactionRequest(sendReactionRequest).Execute()

Send Reaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendReactionRequest := *openapiclient.NewSendReactionRequest() // SendReactionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendReaction(context.Background()).Token(token).SendReactionRequest(sendReactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendReaction`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendReaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendReactionRequest** | [**SendReactionRequest**](SendReactionRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSticker

> MessageSentResponse SendSticker(ctx).Token(token).SendStickerRequest(sendStickerRequest).Execute()

Send sticker message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	token := "token_example" // string | JWT token or API key for authorization
	sendStickerRequest := *openapiclient.NewSendStickerRequest() // SendStickerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.SendSticker(context.Background()).Token(token).SendStickerRequest(sendStickerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.SendSticker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSticker`: MessageSentResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.SendSticker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendStickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **sendStickerRequest** | [**SendStickerRequest**](SendStickerRequest.md) |  | 

### Return type

[**MessageSentResponse**](MessageSentResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

