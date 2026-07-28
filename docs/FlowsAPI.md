# \FlowsAPI

All URIs are relative to *https://api.1msg.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlows**](FlowsAPI.md#CreateFlows) | **Post** /flows | Create Flow
[**CreateFlowsFlowIdDeprecate**](FlowsAPI.md#CreateFlowsFlowIdDeprecate) | **Post** /flows/{flowId}/deprecate | Deprecate Flow
[**CreateFlowsFlowIdPublish**](FlowsAPI.md#CreateFlowsFlowIdPublish) | **Post** /flows/{flowId}/publish | Publish Flow
[**DeleteFlowsFlowId**](FlowsAPI.md#DeleteFlowsFlowId) | **Delete** /flows/{flowId} | Delete Flow
[**GetFlowsFlowId**](FlowsAPI.md#GetFlowsFlowId) | **Get** /flows/{flowId} | Get Flow Details
[**GetFlowsFlowIdPreview**](FlowsAPI.md#GetFlowsFlowIdPreview) | **Get** /flows/{flowId}/preview | Preview Flow
[**GetWhatsappBusinessEncryption**](FlowsAPI.md#GetWhatsappBusinessEncryption) | **Get** /whatsapp_business_encryption | Get business encryption public key
[**ListFlows**](FlowsAPI.md#ListFlows) | **Get** /flows | List Flows
[**PatchFlowsFlowIdAssets**](FlowsAPI.md#PatchFlowsFlowIdAssets) | **Patch** /flows/{flowId}/assets | Update Flow Structure
[**PatchFlowsFlowIdMetadata**](FlowsAPI.md#PatchFlowsFlowIdMetadata) | **Patch** /flows/{flowId}/metadata | Update Flow Metadata
[**SetWhatsappBusinessEncryption**](FlowsAPI.md#SetWhatsappBusinessEncryption) | **Post** /whatsapp_business_encryption | Set business encryption public key



## CreateFlows

> CreateGroups200Response CreateFlows(ctx).Token(token).WabaAccountId(wabaAccountId).Execute()

Create Flow

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
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.CreateFlows(context.Background()).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.CreateFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlows`: CreateGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.CreateFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

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


## CreateFlowsFlowIdDeprecate

> map[string]interface{} CreateFlowsFlowIdDeprecate(ctx, flowId).Token(token).WabaAccountId(wabaAccountId).Execute()

Deprecate Flow

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
	flowId := "flowId_example" // string | Flow ID
	token := "token_example" // string | JWT token or API key for authorization
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.CreateFlowsFlowIdDeprecate(context.Background(), flowId).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.CreateFlowsFlowIdDeprecate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlowsFlowIdDeprecate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.CreateFlowsFlowIdDeprecate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowsFlowIdDeprecateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

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


## CreateFlowsFlowIdPublish

> map[string]interface{} CreateFlowsFlowIdPublish(ctx, flowId).Token(token).WabaAccountId(wabaAccountId).Execute()

Publish Flow

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
	flowId := "flowId_example" // string | Flow ID
	token := "token_example" // string | JWT token or API key for authorization
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.CreateFlowsFlowIdPublish(context.Background(), flowId).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.CreateFlowsFlowIdPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlowsFlowIdPublish`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.CreateFlowsFlowIdPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowsFlowIdPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

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


## DeleteFlowsFlowId

> DeleteFlowsFlowId200Response DeleteFlowsFlowId(ctx, flowId).Token(token).WabaAccountId(wabaAccountId).Execute()

Delete Flow

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
	flowId := "flowId_example" // string | Flow ID
	token := "token_example" // string | JWT token or API key for authorization
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.DeleteFlowsFlowId(context.Background(), flowId).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.DeleteFlowsFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFlowsFlowId`: DeleteFlowsFlowId200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.DeleteFlowsFlowId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlowsFlowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

### Return type

[**DeleteFlowsFlowId200Response**](DeleteFlowsFlowId200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowsFlowId

> CreateGroups200Response GetFlowsFlowId(ctx, flowId).Token(token).WabaAccountId(wabaAccountId).Execute()

Get Flow Details

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
	flowId := "flowId_example" // string | Flow ID
	token := "token_example" // string | JWT token or API key for authorization
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.GetFlowsFlowId(context.Background(), flowId).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlowsFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowsFlowId`: CreateGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlowsFlowId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowsFlowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

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


## GetFlowsFlowIdPreview

> map[string]interface{} GetFlowsFlowIdPreview(ctx, flowId).Token(token).WabaAccountId(wabaAccountId).Execute()

Preview Flow

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
	flowId := "flowId_example" // string | Flow ID
	token := "token_example" // string | JWT token or API key for authorization
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.GetFlowsFlowIdPreview(context.Background(), flowId).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlowsFlowIdPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowsFlowIdPreview`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlowsFlowIdPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowsFlowIdPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

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


## GetWhatsappBusinessEncryption

> GetWhatsappBusinessEncryption200Response GetWhatsappBusinessEncryption(ctx).Token(token).Execute()

Get business encryption public key



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
	resp, r, err := apiClient.FlowsAPI.GetWhatsappBusinessEncryption(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetWhatsappBusinessEncryption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsappBusinessEncryption`: GetWhatsappBusinessEncryption200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetWhatsappBusinessEncryption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsappBusinessEncryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 

### Return type

[**GetWhatsappBusinessEncryption200Response**](GetWhatsappBusinessEncryption200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlows

> ListFlows200Response ListFlows(ctx).Token(token).WabaAccountId(wabaAccountId).Execute()

List Flows

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
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ListFlows(context.Background()).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlows`: ListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ListFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

### Return type

[**ListFlows200Response**](ListFlows200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFlowsFlowIdAssets

> PatchFlowsFlowIdAssets200Response PatchFlowsFlowIdAssets(ctx, flowId).Token(token).WabaAccountId(wabaAccountId).Execute()

Update Flow Structure

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
	flowId := "flowId_example" // string | Flow ID
	token := "token_example" // string | JWT token or API key for authorization
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.PatchFlowsFlowIdAssets(context.Background(), flowId).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.PatchFlowsFlowIdAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFlowsFlowIdAssets`: PatchFlowsFlowIdAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.PatchFlowsFlowIdAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlowsFlowIdAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

### Return type

[**PatchFlowsFlowIdAssets200Response**](PatchFlowsFlowIdAssets200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFlowsFlowIdMetadata

> map[string]interface{} PatchFlowsFlowIdMetadata(ctx, flowId).Token(token).WabaAccountId(wabaAccountId).Execute()

Update Flow Metadata

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
	flowId := "flowId_example" // string | Flow ID
	token := "token_example" // string | JWT token or API key for authorization
	wabaAccountId := "{{wabaAccountId}}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.PatchFlowsFlowIdMetadata(context.Background(), flowId).Token(token).WabaAccountId(wabaAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.PatchFlowsFlowIdMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFlowsFlowIdMetadata`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.PatchFlowsFlowIdMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlowsFlowIdMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | JWT token or API key for authorization | 
 **wabaAccountId** | **string** |  | 

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


## SetWhatsappBusinessEncryption

> CreateCommerce200Response SetWhatsappBusinessEncryption(ctx).Token(token).SetWhatsappBusinessEncryptionRequest(setWhatsappBusinessEncryptionRequest).Execute()

Set business encryption public key



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
	setWhatsappBusinessEncryptionRequest := *openapiclient.NewSetWhatsappBusinessEncryptionRequest("-----BEGIN PUBLIC KEY----- MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA... -----END PUBLIC KEY----- ") // SetWhatsappBusinessEncryptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.SetWhatsappBusinessEncryption(context.Background()).Token(token).SetWhatsappBusinessEncryptionRequest(setWhatsappBusinessEncryptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.SetWhatsappBusinessEncryption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWhatsappBusinessEncryption`: CreateCommerce200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.SetWhatsappBusinessEncryption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetWhatsappBusinessEncryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | JWT token or API key for authorization | 
 **setWhatsappBusinessEncryptionRequest** | [**SetWhatsappBusinessEncryptionRequest**](SetWhatsappBusinessEncryptionRequest.md) |  | 

### Return type

[**CreateCommerce200Response**](CreateCommerce200Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

