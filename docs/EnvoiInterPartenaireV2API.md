# \EnvoiInterPartenaireV2API

All URIs are relative to *http://localhost:8584*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculerFraisEnvoiInterPartenaire**](EnvoiInterPartenaireV2API.md#CalculerFraisEnvoiInterPartenaire) | **Post** /api/v2/nitaServices/envoi-inter-partenaire/frais | Calcul des frais d&#39;un envoi inter-partenaire (P2P)
[**CreerEnvoiInterPartenaire**](EnvoiInterPartenaireV2API.md#CreerEnvoiInterPartenaire) | **Post** /api/v2/nitaServices/envoi-inter-partenaire/creation | Créer un envoi inter-partenaire (P2P)



## CalculerFraisEnvoiInterPartenaire

> ApisResponseV2FraisPreviewResponseV2 CalculerFraisEnvoiInterPartenaire(ctx).EnvoiInterPartenaireFraisRequest(envoiInterPartenaireFraisRequest).Execute()

Calcul des frais d'un envoi inter-partenaire (P2P)



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
	envoiInterPartenaireFraisRequest := *openapiclient.NewEnvoiInterPartenaireFraisRequest("PartenaireDestinataireAlias_example", float64(123)) // EnvoiInterPartenaireFraisRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvoiInterPartenaireV2API.CalculerFraisEnvoiInterPartenaire(context.Background()).EnvoiInterPartenaireFraisRequest(envoiInterPartenaireFraisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvoiInterPartenaireV2API.CalculerFraisEnvoiInterPartenaire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculerFraisEnvoiInterPartenaire`: ApisResponseV2FraisPreviewResponseV2
	fmt.Fprintf(os.Stdout, "Response from `EnvoiInterPartenaireV2API.CalculerFraisEnvoiInterPartenaire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculerFraisEnvoiInterPartenaireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envoiInterPartenaireFraisRequest** | [**EnvoiInterPartenaireFraisRequest**](EnvoiInterPartenaireFraisRequest.md) |  | 

### Return type

[**ApisResponseV2FraisPreviewResponseV2**](ApisResponseV2FraisPreviewResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreerEnvoiInterPartenaire

> ApisResponseV2EnvoiInterPartenaireResponseV2 CreerEnvoiInterPartenaire(ctx).EnvoiInterPartenaireRequest(envoiInterPartenaireRequest).Execute()

Créer un envoi inter-partenaire (P2P)



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
	envoiInterPartenaireRequest := *openapiclient.NewEnvoiInterPartenaireRequest() // EnvoiInterPartenaireRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvoiInterPartenaireV2API.CreerEnvoiInterPartenaire(context.Background()).EnvoiInterPartenaireRequest(envoiInterPartenaireRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvoiInterPartenaireV2API.CreerEnvoiInterPartenaire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreerEnvoiInterPartenaire`: ApisResponseV2EnvoiInterPartenaireResponseV2
	fmt.Fprintf(os.Stdout, "Response from `EnvoiInterPartenaireV2API.CreerEnvoiInterPartenaire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreerEnvoiInterPartenaireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envoiInterPartenaireRequest** | [**EnvoiInterPartenaireRequest**](EnvoiInterPartenaireRequest.md) |  | 

### Return type

[**ApisResponseV2EnvoiInterPartenaireResponseV2**](ApisResponseV2EnvoiInterPartenaireResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

