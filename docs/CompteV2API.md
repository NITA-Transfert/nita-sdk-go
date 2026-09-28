# \CompteV2API

All URIs are relative to *http://localhost:8584*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsulterSoldeCompte**](CompteV2API.md#ConsulterSoldeCompte) | **Get** /api/v2/nitaServices/account/balance | Consulter le solde du compte



## ConsulterSoldeCompte

> ApisResponseV2Double ConsulterSoldeCompte(ctx).Execute()

Consulter le solde du compte



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
	resp, r, err := apiClient.CompteV2API.ConsulterSoldeCompte(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompteV2API.ConsulterSoldeCompte``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsulterSoldeCompte`: ApisResponseV2Double
	fmt.Fprintf(os.Stdout, "Response from `CompteV2API.ConsulterSoldeCompte`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsulterSoldeCompteRequest struct via the builder pattern


### Return type

[**ApisResponseV2Double**](ApisResponseV2Double.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

