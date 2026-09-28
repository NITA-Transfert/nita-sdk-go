# \LocalitsV2API

All URIs are relative to *http://localhost:8584*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllActiveVillesListe**](LocalitsV2API.md#GetAllActiveVillesListe) | **Get** /api/v2/nitaServices/localite/ville | Obtenir la liste des villes actives
[**GetAllIndicatifs**](LocalitsV2API.md#GetAllIndicatifs) | **Get** /api/v2/nitaServices/localite/pays | Obtenir la liste des indicatifs téléphoniques de tous les pays actifs



## GetAllActiveVillesListe

> ApisResponseV2ListVilleResponse GetAllActiveVillesListe(ctx).Execute()

Obtenir la liste des villes actives

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NITA-Transfert/nita-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalitsV2API.GetAllActiveVillesListe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalitsV2API.GetAllActiveVillesListe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllActiveVillesListe`: ApisResponseV2ListVilleResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalitsV2API.GetAllActiveVillesListe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllActiveVillesListeRequest struct via the builder pattern


### Return type

[**ApisResponseV2ListVilleResponse**](ApisResponseV2ListVilleResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIndicatifs

> ApisResponseV2IndicatifsResponseV2 GetAllIndicatifs(ctx).Execute()

Obtenir la liste des indicatifs téléphoniques de tous les pays actifs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NITA-Transfert/nita-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalitsV2API.GetAllIndicatifs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalitsV2API.GetAllIndicatifs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllIndicatifs`: ApisResponseV2IndicatifsResponseV2
	fmt.Fprintf(os.Stdout, "Response from `LocalitsV2API.GetAllIndicatifs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIndicatifsRequest struct via the builder pattern


### Return type

[**ApisResponseV2IndicatifsResponseV2**](ApisResponseV2IndicatifsResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

