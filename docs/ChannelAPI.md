# \ChannelAPI

All URIs are relative to *https://api.1msg.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSettings**](ChannelAPI.md#CreateSettings) | **Post** /settings | Update channel settings
[**GetConversationalAutomation**](ChannelAPI.md#GetConversationalAutomation) | **Get** /conversationalAutomation | Get conversational automation settings
[**GetStatus**](ChannelAPI.md#GetStatus) | **Get** /status | Get channel status
[**ListSettings**](ChannelAPI.md#ListSettings) | **Get** /settings | Get channel settings
[**SetConversationalAutomation**](ChannelAPI.md#SetConversationalAutomation) | **Post** /conversationalAutomation | Set conversational automation settings



## CreateSettings

> map[string]interface{} CreateSettings(ctx).Token(token).CreateSettingsRequest(createSettingsRequest).Execute()

Update channel settings



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
	createSettingsRequest := *openapiclient.NewCreateSettingsRequest() // CreateSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.CreateSettings(context.Background()).Token(token).CreateSettingsRequest(createSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.CreateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.CreateSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **createSettingsRequest** | [**CreateSettingsRequest**](CreateSettingsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationalAutomation

> ConversationalAutomation GetConversationalAutomation(ctx).Token(token).Execute()

Get conversational automation settings



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
	resp, r, err := apiClient.ChannelAPI.GetConversationalAutomation(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetConversationalAutomation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationalAutomation`: ConversationalAutomation
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetConversationalAutomation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationalAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**ConversationalAutomation**](ConversationalAutomation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> GetStatus200Response GetStatus(ctx).Token(token).Execute()

Get channel status



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
	resp, r, err := apiClient.ChannelAPI.GetStatus(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: GetStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**GetStatus200Response**](GetStatus200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSettings

> ListSettings200Response ListSettings(ctx).Token(token).Execute()

Get channel settings



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
	resp, r, err := apiClient.ChannelAPI.ListSettings(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ListSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSettings`: ListSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ListSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**ListSettings200Response**](ListSettings200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConversationalAutomation

> map[string]interface{} SetConversationalAutomation(ctx).Token(token).ConversationalAutomation(conversationalAutomation).Execute()

Set conversational automation settings



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
	conversationalAutomation := *openapiclient.NewConversationalAutomation() // ConversationalAutomation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.SetConversationalAutomation(context.Background()).Token(token).ConversationalAutomation(conversationalAutomation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.SetConversationalAutomation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConversationalAutomation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.SetConversationalAutomation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConversationalAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **conversationalAutomation** | [**ConversationalAutomation**](ConversationalAutomation.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

