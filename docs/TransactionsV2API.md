# \TransactionsV2API

All URIs are relative to *http://localhost:8584*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnulePartenaireToWallet**](TransactionsV2API.md#AnnulePartenaireToWallet) | **Post** /api/v2/nitaServices/transaction/AnnulePartenaireToWallet | Annulation Partenaire To Wallet (v2)
[**AnnulerEnvoiMyNita**](TransactionsV2API.md#AnnulerEnvoiMyNita) | **Put** /api/v2/nitaServices/transaction/annulerEnvoi | Annuler un envoi MyNita (v2)
[**CalculerFraisPartenaireToCash**](TransactionsV2API.md#CalculerFraisPartenaireToCash) | **Post** /api/v2/nitaServices/transaction/partenaireToCash/frais | Calcul des frais Partenaire To Cash (P2C)
[**CalculerFraisPartenaireToWallet**](TransactionsV2API.md#CalculerFraisPartenaireToWallet) | **Post** /api/v2/nitaServices/transaction/partenaireToWallet/frais | Calcul des frais Partenaire To Wallet (P2W)
[**GetEnvoiToEdit**](TransactionsV2API.md#GetEnvoiToEdit) | **Post** /api/v2/nitaServices/transaction/getEnvoiToEdit | Récupérer les informations d&#39;un envoi à éditer (v2)
[**GetPartenaireToWallet**](TransactionsV2API.md#GetPartenaireToWallet) | **Post** /api/v2/nitaServices/transaction/GetPartenaireToWallet | Récupération d&#39;une opération Partenaire To Wallet (v2)
[**PartenaireToCash**](TransactionsV2API.md#PartenaireToCash) | **Post** /api/v2/nitaServices/transaction/partenaireToCash | Envoi Partenaire To Cash (v2)
[**PartenaireToWallet**](TransactionsV2API.md#PartenaireToWallet) | **Post** /api/v2/nitaServices/transaction/partenaireToWallet | Partenaire To Wallet (v2)
[**UpdateEnvoiMyNita**](TransactionsV2API.md#UpdateEnvoiMyNita) | **Put** /api/v2/nitaServices/transaction/updateEnvoi | Modifier un envoi MyNita (v2)



## AnnulePartenaireToWallet

> ApisResponseV2AnnulationPartenaireToWalletResponseV2 AnnulePartenaireToWallet(ctx).AnnulePartenaireToWalletModel(annulePartenaireToWalletModel).Execute()

Annulation Partenaire To Wallet (v2)



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
	annulePartenaireToWalletModel := *openapiclient.NewAnnulePartenaireToWalletModel("RequestId_example", "CodeRecharge_example", "AdresseIp_example", "MotifAnnulation_example") // AnnulePartenaireToWalletModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.AnnulePartenaireToWallet(context.Background()).AnnulePartenaireToWalletModel(annulePartenaireToWalletModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.AnnulePartenaireToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnulePartenaireToWallet`: ApisResponseV2AnnulationPartenaireToWalletResponseV2
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.AnnulePartenaireToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnnulePartenaireToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **annulePartenaireToWalletModel** | [**AnnulePartenaireToWalletModel**](AnnulePartenaireToWalletModel.md) |  | 

### Return type

[**ApisResponseV2AnnulationPartenaireToWalletResponseV2**](ApisResponseV2AnnulationPartenaireToWalletResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnulerEnvoiMyNita

> ApisResponseV2TransactionResponseV2 AnnulerEnvoiMyNita(ctx).AnnulationDto(annulationDto).Execute()

Annuler un envoi MyNita (v2)



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
	annulationDto := *openapiclient.NewAnnulationDto("CodeEnvoi_example", "RequestId_example", "NumeroExpediteur_example", "IndicatifExpediteur_example") // AnnulationDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.AnnulerEnvoiMyNita(context.Background()).AnnulationDto(annulationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.AnnulerEnvoiMyNita``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnulerEnvoiMyNita`: ApisResponseV2TransactionResponseV2
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.AnnulerEnvoiMyNita`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnnulerEnvoiMyNitaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **annulationDto** | [**AnnulationDto**](AnnulationDto.md) |  | 

### Return type

[**ApisResponseV2TransactionResponseV2**](ApisResponseV2TransactionResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalculerFraisPartenaireToCash

> ApisResponseV2FraisPreviewResponseV2 CalculerFraisPartenaireToCash(ctx).CalculeFraisOperationRequest(calculeFraisOperationRequest).Execute()

Calcul des frais Partenaire To Cash (P2C)



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
	calculeFraisOperationRequest := *openapiclient.NewCalculeFraisOperationRequest("IndicatifExpediteur_example", "NumeroExpediteur_example", float64(123), false, "VilleExpedition_example", "VilleDestination_example") // CalculeFraisOperationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.CalculerFraisPartenaireToCash(context.Background()).CalculeFraisOperationRequest(calculeFraisOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.CalculerFraisPartenaireToCash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculerFraisPartenaireToCash`: ApisResponseV2FraisPreviewResponseV2
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.CalculerFraisPartenaireToCash`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculerFraisPartenaireToCashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculeFraisOperationRequest** | [**CalculeFraisOperationRequest**](CalculeFraisOperationRequest.md) |  | 

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


## CalculerFraisPartenaireToWallet

> ApisResponseV2FraisPreviewResponseV2 CalculerFraisPartenaireToWallet(ctx).PartenaireToWalletFraisRequest(partenaireToWalletFraisRequest).Execute()

Calcul des frais Partenaire To Wallet (P2W)



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
	partenaireToWalletFraisRequest := *openapiclient.NewPartenaireToWalletFraisRequest("PhoneClient_example", float64(123)) // PartenaireToWalletFraisRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.CalculerFraisPartenaireToWallet(context.Background()).PartenaireToWalletFraisRequest(partenaireToWalletFraisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.CalculerFraisPartenaireToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculerFraisPartenaireToWallet`: ApisResponseV2FraisPreviewResponseV2
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.CalculerFraisPartenaireToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculerFraisPartenaireToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partenaireToWalletFraisRequest** | [**PartenaireToWalletFraisRequest**](PartenaireToWalletFraisRequest.md) |  | 

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


## GetEnvoiToEdit

> ApisResponseV2ModelEditEnvoieRequest GetEnvoiToEdit(ctx).EnvoiToEditDto(envoiToEditDto).Execute()

Récupérer les informations d'un envoi à éditer (v2)



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
	envoiToEditDto := *openapiclient.NewEnvoiToEditDto("CodeEnvoi_example", "RequestId_example", "IndicatifExpediteur_example", "NumeroExpediteur_example") // EnvoiToEditDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.GetEnvoiToEdit(context.Background()).EnvoiToEditDto(envoiToEditDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.GetEnvoiToEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvoiToEdit`: ApisResponseV2ModelEditEnvoieRequest
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.GetEnvoiToEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvoiToEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envoiToEditDto** | [**EnvoiToEditDto**](EnvoiToEditDto.md) |  | 

### Return type

[**ApisResponseV2ModelEditEnvoieRequest**](ApisResponseV2ModelEditEnvoieRequest.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartenaireToWallet

> ApisResponseV2GetPartenaireToWalletResponseV2 GetPartenaireToWallet(ctx).GetPartenaireToWalletModel(getPartenaireToWalletModel).Execute()

Récupération d'une opération Partenaire To Wallet (v2)



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
	getPartenaireToWalletModel := *openapiclient.NewGetPartenaireToWalletModel("RequestId_example") // GetPartenaireToWalletModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.GetPartenaireToWallet(context.Background()).GetPartenaireToWalletModel(getPartenaireToWalletModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.GetPartenaireToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartenaireToWallet`: ApisResponseV2GetPartenaireToWalletResponseV2
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.GetPartenaireToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPartenaireToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPartenaireToWalletModel** | [**GetPartenaireToWalletModel**](GetPartenaireToWalletModel.md) |  | 

### Return type

[**ApisResponseV2GetPartenaireToWalletResponseV2**](ApisResponseV2GetPartenaireToWalletResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartenaireToCash

> ApisResponseV2TransactionResponseV2 PartenaireToCash(ctx).PartenaireToCashDtoV2(partenaireToCashDtoV2).Execute()

Envoi Partenaire To Cash (v2)



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
	partenaireToCashDtoV2 := *openapiclient.NewPartenaireToCashDtoV2("Nom_example", "Prenom_example", "Indicatif_example", "Numero_example", "RequestId_example", false, float64(123), "VilleDestination_example", "NomDestinataire_example", "PrenomDestinataire_example", "IndicatifDestinataire_example", "NumeroDestinataire_example") // PartenaireToCashDtoV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.PartenaireToCash(context.Background()).PartenaireToCashDtoV2(partenaireToCashDtoV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.PartenaireToCash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartenaireToCash`: ApisResponseV2TransactionResponseV2
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.PartenaireToCash`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartenaireToCashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partenaireToCashDtoV2** | [**PartenaireToCashDtoV2**](PartenaireToCashDtoV2.md) |  | 

### Return type

[**ApisResponseV2TransactionResponseV2**](ApisResponseV2TransactionResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartenaireToWallet

> ApisResponseV2PartenaireToWalletResponseV2 PartenaireToWallet(ctx).PartenaireToWalletModel(partenaireToWalletModel).Execute()

Partenaire To Wallet (v2)



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
	partenaireToWalletModel := *openapiclient.NewPartenaireToWalletModel() // PartenaireToWalletModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.PartenaireToWallet(context.Background()).PartenaireToWalletModel(partenaireToWalletModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.PartenaireToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartenaireToWallet`: ApisResponseV2PartenaireToWalletResponseV2
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.PartenaireToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartenaireToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partenaireToWalletModel** | [**PartenaireToWalletModel**](PartenaireToWalletModel.md) |  | 

### Return type

[**ApisResponseV2PartenaireToWalletResponseV2**](ApisResponseV2PartenaireToWalletResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvoiMyNita

> ApisResponseV2Void UpdateEnvoiMyNita(ctx).UpdateEnvoiDto(updateEnvoiDto).Execute()

Modifier un envoi MyNita (v2)



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
	updateEnvoiDto := *openapiclient.NewUpdateEnvoiDto("CodeEnvoi_example", "RequestId_example", "IndicatifExpediteur_example", "NumeroExpediteur_example", "NomDestinataire_example", "PrenomDestinataire_example", "IndicatifDestinataire_example", "NumeroDestinataire_example", "VilleDestination_example") // UpdateEnvoiDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsV2API.UpdateEnvoiMyNita(context.Background()).UpdateEnvoiDto(updateEnvoiDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsV2API.UpdateEnvoiMyNita``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnvoiMyNita`: ApisResponseV2Void
	fmt.Fprintf(os.Stdout, "Response from `TransactionsV2API.UpdateEnvoiMyNita`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvoiMyNitaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEnvoiDto** | [**UpdateEnvoiDto**](UpdateEnvoiDto.md) |  | 

### Return type

[**ApisResponseV2Void**](ApisResponseV2Void.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

