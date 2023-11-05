# Go API client for openapi

The Selling Partner API for Listings Items (Listings Items API) provides programmatic access to selling partner listings on Amazon. Use this API in collaboration with the Selling Partner API for Product Type Definitions, which you use to retrieve the information about Amazon product types needed to use the Listings Items API.

For more information, see the [Listings Items API Use Case Guide](doc:listings-items-api-v2021-08-01-use-case-guide).

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2021-08-01
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://sellercentral.amazon.com/gp/mws/contactus.html](https://sellercentral.amazon.com/gp/mws/contactus.html)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ListingsAPI* | [**DeleteListingsItem**](docs/ListingsAPI.md#deletelistingsitem) | **Delete** /listings/2021-08-01/items/{sellerId}/{sku} | 
*ListingsAPI* | [**GetListingsItem**](docs/ListingsAPI.md#getlistingsitem) | **Get** /listings/2021-08-01/items/{sellerId}/{sku} | 
*ListingsAPI* | [**PatchListingsItem**](docs/ListingsAPI.md#patchlistingsitem) | **Patch** /listings/2021-08-01/items/{sellerId}/{sku} | 
*ListingsAPI* | [**PutListingsItem**](docs/ListingsAPI.md#putlistingsitem) | **Put** /listings/2021-08-01/items/{sellerId}/{sku} | 


## Documentation For Models

 - [Error](docs/Error.md)
 - [ErrorList](docs/ErrorList.md)
 - [FulfillmentAvailability](docs/FulfillmentAvailability.md)
 - [Issue](docs/Issue.md)
 - [Item](docs/Item.md)
 - [ItemImage](docs/ItemImage.md)
 - [ItemOfferByMarketplace](docs/ItemOfferByMarketplace.md)
 - [ItemProcurement](docs/ItemProcurement.md)
 - [ItemSummaryByMarketplace](docs/ItemSummaryByMarketplace.md)
 - [ListingsItemPatchRequest](docs/ListingsItemPatchRequest.md)
 - [ListingsItemPutRequest](docs/ListingsItemPutRequest.md)
 - [ListingsItemSubmissionResponse](docs/ListingsItemSubmissionResponse.md)
 - [Money](docs/Money.md)
 - [PatchOperation](docs/PatchOperation.md)
 - [Points](docs/Points.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



