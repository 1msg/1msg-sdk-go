# \CallingAPI

All URIs are relative to *https://api.1msg.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCallingSettings**](CallingAPI.md#GetCallingSettings) | **Get** /callingSettings | Get calling settings
[**InitiateCall**](CallingAPI.md#InitiateCall) | **Post** /initiateCall | Call action (connect / pre_accept / accept / reject / terminate)
[**UpdateCallingSettings**](CallingAPI.md#UpdateCallingSettings) | **Post** /callingSettings | Update calling settings



## GetCallingSettings

> CallingSettings GetCallingSettings(ctx).Token(token).Execute()

Get calling settings



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
	resp, r, err := apiClient.CallingAPI.GetCallingSettings(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallingAPI.GetCallingSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallingSettings`: CallingSettings
	fmt.Fprintf(os.Stdout, "Response from `CallingAPI.GetCallingSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCallingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**CallingSettings**](CallingSettings.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateCall

> InitiateCallResponse InitiateCall(ctx).Token(token).InitiateCallRequest(initiateCallRequest).Execute()

Call action (connect / pre_accept / accept / reject / terminate)



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
	initiateCallRequest := *openapiclient.NewInitiateCallRequest("whatsapp", "connect") // InitiateCallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallingAPI.InitiateCall(context.Background()).Token(token).InitiateCallRequest(initiateCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallingAPI.InitiateCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateCall`: InitiateCallResponse
	fmt.Fprintf(os.Stdout, "Response from `CallingAPI.InitiateCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **initiateCallRequest** | [**InitiateCallRequest**](InitiateCallRequest.md) |  | 

### Return type

[**InitiateCallResponse**](InitiateCallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallingSettings

> UpdateCallingSettings200Response UpdateCallingSettings(ctx).Token(token).CallingSettings(callingSettings).Execute()

Update calling settings



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
	callingSettings := *openapiclient.NewCallingSettings() // CallingSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallingAPI.UpdateCallingSettings(context.Background()).Token(token).CallingSettings(callingSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallingAPI.UpdateCallingSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCallingSettings`: UpdateCallingSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `CallingAPI.UpdateCallingSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **callingSettings** | [**CallingSettings**](CallingSettings.md) |  | 

### Return type

[**UpdateCallingSettings200Response**](UpdateCallingSettings200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

