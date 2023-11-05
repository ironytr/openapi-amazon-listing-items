# \ListingsAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListingsItem**](ListingsAPI.md#DeleteListingsItem) | **Delete** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**GetListingsItem**](ListingsAPI.md#GetListingsItem) | **Get** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**PatchListingsItem**](ListingsAPI.md#PatchListingsItem) | **Patch** /listings/2021-08-01/items/{sellerId}/{sku} | 
[**PutListingsItem**](ListingsAPI.md#PutListingsItem) | **Put** /listings/2021-08-01/items/{sellerId}/{sku} | 



## DeleteListingsItem

> ListingsItemSubmissionResponse DeleteListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).Execute()





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
    sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
    sku := "sku_example" // string | A selling partner provided identifier for an Amazon listing.
    marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers for the request.
    issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \"en_US\", \"fr_CA\", \"fr_FR\". Localized messages default to \"en_US\" when a localization is not available in the specified locale. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsAPI.DeleteListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.DeleteListingsItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteListingsItem`: ListingsItemSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.DeleteListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsItem

> Item GetListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).IncludedData(includedData).Execute()





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
    sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
    sku := "sku_example" // string | A selling partner provided identifier for an Amazon listing.
    marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers for the request.
    issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \"en_US\", \"fr_CA\", \"fr_FR\". Localized messages default to \"en_US\" when a localization is not available in the specified locale. (optional)
    includedData := TODO // interface{} | A comma-delimited list of data sets to include in the response. Default: summaries. (optional) (default to summaries)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsAPI.GetListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).IssueLocale(issueLocale).IncludedData(includedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.GetListingsItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingsItem`: Item
    fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.GetListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 
 **includedData** | [**interface{}**](interface{}.md) | A comma-delimited list of data sets to include in the response. Default: summaries. | [default to summaries]

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListingsItem

> ListingsItemSubmissionResponse PatchListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IssueLocale(issueLocale).Execute()





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
    sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
    sku := "sku_example" // string | A selling partner provided identifier for an Amazon listing.
    marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers for the request.
    body := *openapiclient.NewListingsItemPatchRequest("ProductType_example", []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("Op_example", "Path_example")}) // ListingsItemPatchRequest | The request body schema for the patchListingsItem operation.
    issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \"en_US\", \"fr_CA\", \"fr_FR\". Localized messages default to \"en_US\" when a localization is not available in the specified locale. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsAPI.PatchListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IssueLocale(issueLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PatchListingsItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchListingsItem`: ListingsItemSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.PatchListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **body** | [**ListingsItemPatchRequest**](ListingsItemPatchRequest.md) | The request body schema for the patchListingsItem operation. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutListingsItem

> ListingsItemSubmissionResponse PutListingsItem(ctx, sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IssueLocale(issueLocale).Execute()





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
    sellerId := "sellerId_example" // string | A selling partner identifier, such as a merchant account or vendor code.
    sku := "sku_example" // string | A selling partner provided identifier for an Amazon listing.
    marketplaceIds := []string{"Inner_example"} // []string | A comma-delimited list of Amazon marketplace identifiers for the request.
    body := *openapiclient.NewListingsItemPutRequest("ProductType_example", map[string]interface{}{"key": interface{}(123)}) // ListingsItemPutRequest | The request body schema for the putListingsItem operation.
    issueLocale := "en_US" // string | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \"en_US\", \"fr_CA\", \"fr_FR\". Localized messages default to \"en_US\" when a localization is not available in the specified locale. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsAPI.PutListingsItem(context.Background(), sellerId, sku).MarketplaceIds(marketplaceIds).Body(body).IssueLocale(issueLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PutListingsItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutListingsItem`: ListingsItemSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsAPI.PutListingsItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sellerId** | **string** | A selling partner identifier, such as a merchant account or vendor code. | 
**sku** | **string** | A selling partner provided identifier for an Amazon listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutListingsItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketplaceIds** | **[]string** | A comma-delimited list of Amazon marketplace identifiers for the request. | 
 **body** | [**ListingsItemPutRequest**](ListingsItemPutRequest.md) | The request body schema for the putListingsItem operation. | 
 **issueLocale** | **string** | A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. Localized messages default to \&quot;en_US\&quot; when a localization is not available in the specified locale. | 

### Return type

[**ListingsItemSubmissionResponse**](ListingsItemSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

