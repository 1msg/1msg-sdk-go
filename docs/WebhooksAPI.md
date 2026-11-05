# \WebhooksAPI

All URIs are relative to *https://api.1msg.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhook**](WebhooksAPI.md#GetWebhook) | **Get** /webhook | Get webhook URL
[**SetWebhook**](WebhooksAPI.md#SetWebhook) | **Post** /webhook | Set webhook URL



## GetWebhook

> GetWebhook200Response GetWebhook(ctx).Token(token).Execute()

Get webhook URL



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
	resp, r, err := apiClient.WebhooksAPI.GetWebhook(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: GetWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**GetWebhook200Response**](GetWebhook200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWebhook

> SetWebhook200Response SetWebhook(ctx).Token(token).SetWebhookRequest(setWebhookRequest).Execute()

Set webhook URL



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
	setWebhookRequest := *openapiclient.NewSetWebhookRequest(openapiclient.setWebhook_request_webhookUrl{ArrayOfString: new([]string)}) // SetWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.SetWebhook(context.Background()).Token(token).SetWebhookRequest(setWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.SetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWebhook`: SetWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.SetWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **setWebhookRequest** | [**SetWebhookRequest**](SetWebhookRequest.md) |  | 

### Return type

[**SetWebhook200Response**](SetWebhook200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

