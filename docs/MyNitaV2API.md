# \MyNitaV2API

All URIs are relative to *http://localhost:8584*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCompteExistence**](MyNitaV2API.md#CheckCompteExistence) | **Post** /api/v2/nitaServices/compteMynita/checkCompteExistence | Vérifier l&#39;existence d&#39;un compte MyNita



## CheckCompteExistence

> ApisResponseV2String CheckCompteExistence(ctx).VerificationCompteDtoV2(verificationCompteDtoV2).Execute()

Vérifier l'existence d'un compte MyNita



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
	verificationCompteDtoV2 := *openapiclient.NewVerificationCompteDtoV2("IndicatifPhoneClient_example", "PhoneClient_example") // VerificationCompteDtoV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyNitaV2API.CheckCompteExistence(context.Background()).VerificationCompteDtoV2(verificationCompteDtoV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyNitaV2API.CheckCompteExistence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCompteExistence`: ApisResponseV2String
	fmt.Fprintf(os.Stdout, "Response from `MyNitaV2API.CheckCompteExistence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCompteExistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verificationCompteDtoV2** | [**VerificationCompteDtoV2**](VerificationCompteDtoV2.md) |  | 

### Return type

[**ApisResponseV2String**](ApisResponseV2String.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

