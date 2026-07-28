# \GroupsAPI

All URIs are relative to *https://api.1msg.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroups**](GroupsAPI.md#CreateGroups) | **Post** /groups | Create Group
[**CreateGroupsGroupId**](GroupsAPI.md#CreateGroupsGroupId) | **Post** /groups/{groupId} | Update Group Info
[**CreateGroupsGroupIdInvitelink**](GroupsAPI.md#CreateGroupsGroupIdInvitelink) | **Post** /groups/{groupId}/inviteLink | Reset Invite Link
[**DeleteGroupsGroupId**](GroupsAPI.md#DeleteGroupsGroupId) | **Delete** /groups/{groupId} | Delete Group
[**GetGroupsGroupId**](GroupsAPI.md#GetGroupsGroupId) | **Get** /groups/{groupId} | Get Group Info
[**GetGroupsGroupIdInvitelink**](GroupsAPI.md#GetGroupsGroupIdInvitelink) | **Get** /groups/{groupId}/inviteLink | Get Invite Link
[**ListGroups**](GroupsAPI.md#ListGroups) | **Get** /groups | Get Groups List



## CreateGroups

> CreateGroups200Response CreateGroups(ctx).Token(token).Execute()

Create Group

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
	resp, r, err := apiClient.GroupsAPI.CreateGroups(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroups`: CreateGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**CreateGroups200Response**](CreateGroups200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupsGroupId

> map[string]interface{} CreateGroupsGroupId(ctx, groupId).Token(token).Execute()

Update Group Info

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
	groupId := "{{createdGroupId}}" // string | 
	token := "token_example" // string | JWT token or API key for authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroupsGroupId(context.Background(), groupId).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroupsGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupsGroupId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroupsGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupsGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 

### Return type

**map[string]interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupsGroupIdInvitelink

> map[string]interface{} CreateGroupsGroupIdInvitelink(ctx, groupId).Token(token).Execute()

Reset Invite Link

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
	groupId := "{{createdGroupId}}" // string | 
	token := "token_example" // string | JWT token or API key for authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroupsGroupIdInvitelink(context.Background(), groupId).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroupsGroupIdInvitelink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupsGroupIdInvitelink`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroupsGroupIdInvitelink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupsGroupIdInvitelinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 

### Return type

**map[string]interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupsGroupId

> DeleteGroupsGroupId200Response DeleteGroupsGroupId(ctx, groupId).Token(token).Execute()

Delete Group

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
	groupId := "{{createdGroupId}}" // string | 
	token := "token_example" // string | JWT token or API key for authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupsGroupId(context.Background(), groupId).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupsGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupsGroupId`: DeleteGroupsGroupId200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupsGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupsGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**DeleteGroupsGroupId200Response**](DeleteGroupsGroupId200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsGroupId

> map[string]interface{} GetGroupsGroupId(ctx, groupId).Token(token).Execute()

Get Group Info

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
	groupId := "{{createdGroupId}}" // string | 
	token := "token_example" // string | JWT token or API key for authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupsGroupId(context.Background(), groupId).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupsGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsGroupId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupsGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 

### Return type

**map[string]interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsGroupIdInvitelink

> map[string]interface{} GetGroupsGroupIdInvitelink(ctx, groupId).Token(token).Execute()

Get Invite Link

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
	groupId := "{{createdGroupId}}" // string | 
	token := "token_example" // string | JWT token or API key for authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupsGroupIdInvitelink(context.Background(), groupId).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupsGroupIdInvitelink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsGroupIdInvitelink`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupsGroupIdInvitelink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsGroupIdInvitelinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 

### Return type

**map[string]interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> map[string]interface{} ListGroups(ctx).Token(token).Execute()

Get Groups List

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
	resp, r, err := apiClient.GroupsAPI.ListGroups(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ListGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

**map[string]interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

