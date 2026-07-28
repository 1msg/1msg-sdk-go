# \UsersAPI

All URIs are relative to *https://api.1msg.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockUser**](UsersAPI.md#BlockUser) | **Post** /blockUser | Block WhatsApp user
[**ListBlockedUsers**](UsersAPI.md#ListBlockedUsers) | **Get** /blockedUsers | List blocked WhatsApp users
[**UnblockUser**](UsersAPI.md#UnblockUser) | **Post** /unblockUser | Unblock WhatsApp user



## BlockUser

> SuccessResponse BlockUser(ctx).Token(token).BlockUserRequest(blockUserRequest).Execute()

Block WhatsApp user

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
	blockUserRequest := *openapiclient.NewBlockUserRequest(int32(123)) // BlockUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.BlockUser(context.Background()).Token(token).BlockUserRequest(blockUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.BlockUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockUser`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.BlockUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **blockUserRequest** | [**BlockUserRequest**](BlockUserRequest.md) |  | 

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


## ListBlockedUsers

> ListBlockedUsers200Response ListBlockedUsers(ctx).Token(token).Execute()

List blocked WhatsApp users



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
	resp, r, err := apiClient.UsersAPI.ListBlockedUsers(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListBlockedUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockedUsers`: ListBlockedUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListBlockedUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockedUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**ListBlockedUsers200Response**](ListBlockedUsers200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnblockUser

> SuccessResponse UnblockUser(ctx).Token(token).BlockUserRequest(blockUserRequest).Execute()

Unblock WhatsApp user

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
	blockUserRequest := *openapiclient.NewBlockUserRequest(int32(123)) // BlockUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UnblockUser(context.Background()).Token(token).BlockUserRequest(blockUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnblockUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnblockUser`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UnblockUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnblockUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **blockUserRequest** | [**BlockUserRequest**](BlockUserRequest.md) |  | 

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

