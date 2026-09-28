# \AchatEnLigneV2API

All URIs are relative to *http://localhost:8584*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAchatEnligne**](AchatEnLigneV2API.md#AddAchatEnligne) | **Post** /api/v2/nitaServices/achatEnLigne/saveAchatEnLigne | Créer un achat en ligne
[**AnnulerAchat**](AchatEnLigneV2API.md#AnnulerAchat) | **Put** /api/v2/nitaServices/achatEnLigne/annulerAchat | Annuler un achat
[**CheckAchatStatus**](AchatEnLigneV2API.md#CheckAchatStatus) | **Post** /api/v2/nitaServices/achatEnLigne/checkAchatStatus | Vérifier le statut d&#39;un achat



## AddAchatEnligne

> ApisResponseV2AddAchatEnligneResponseV2 AddAchatEnligne(ctx).AchatEnLigneModel(achatEnLigneModel).Execute()

Créer un achat en ligne



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
	achatEnLigneModel := *openapiclient.NewAchatEnLigneModel([]string{"DescriptionAchat_example"}, float64(123), "PhoneClient_example", "MotifTransaction_example", "LongTransaction_example", "LatTransaction_example", "RequestId_example", "AdresseIp_example") // AchatEnLigneModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchatEnLigneV2API.AddAchatEnligne(context.Background()).AchatEnLigneModel(achatEnLigneModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchatEnLigneV2API.AddAchatEnligne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAchatEnligne`: ApisResponseV2AddAchatEnligneResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AchatEnLigneV2API.AddAchatEnligne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAchatEnligneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achatEnLigneModel** | [**AchatEnLigneModel**](AchatEnLigneModel.md) |  | 

### Return type

[**ApisResponseV2AddAchatEnligneResponseV2**](ApisResponseV2AddAchatEnligneResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnulerAchat

> ApisResponseV2CheckAchatStatusResponseV2 AnnulerAchat(ctx).AnnulationAchatDto(annulationAchatDto).Execute()

Annuler un achat



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
	annulationAchatDto := *openapiclient.NewAnnulationAchatDto() // AnnulationAchatDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchatEnLigneV2API.AnnulerAchat(context.Background()).AnnulationAchatDto(annulationAchatDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchatEnLigneV2API.AnnulerAchat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnulerAchat`: ApisResponseV2CheckAchatStatusResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AchatEnLigneV2API.AnnulerAchat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnnulerAchatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **annulationAchatDto** | [**AnnulationAchatDto**](AnnulationAchatDto.md) |  | 

### Return type

[**ApisResponseV2CheckAchatStatusResponseV2**](ApisResponseV2CheckAchatStatusResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAchatStatus

> ApisResponseV2CheckAchatStatusResponseV2 CheckAchatStatus(ctx).CheckAchatStatusModel(checkAchatStatusModel).Execute()

Vérifier le statut d'un achat



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
	checkAchatStatusModel := *openapiclient.NewCheckAchatStatusModel() // CheckAchatStatusModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AchatEnLigneV2API.CheckAchatStatus(context.Background()).CheckAchatStatusModel(checkAchatStatusModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AchatEnLigneV2API.CheckAchatStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckAchatStatus`: ApisResponseV2CheckAchatStatusResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AchatEnLigneV2API.CheckAchatStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAchatStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAchatStatusModel** | [**CheckAchatStatusModel**](CheckAchatStatusModel.md) |  | 

### Return type

[**ApisResponseV2CheckAchatStatusResponseV2**](ApisResponseV2CheckAchatStatusResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

