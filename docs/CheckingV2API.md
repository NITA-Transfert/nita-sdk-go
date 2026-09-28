# \CheckingV2API

All URIs are relative to *http://localhost:8584*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRequestOperation**](CheckingV2API.md#CheckRequestOperation) | **Post** /api/v2/nitaServices/checkStatus/transaction | Vérification de l&#39;état d&#39;un envoi



## CheckRequestOperation

> ApisResponseV2CheckStatusResponseV2 CheckRequestOperation(ctx).CheckTransactionRequest(checkTransactionRequest).Execute()

Vérification de l'état d'un envoi



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
	checkTransactionRequest := *openapiclient.NewCheckTransactionRequest("RequestId_example") // CheckTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckingV2API.CheckRequestOperation(context.Background()).CheckTransactionRequest(checkTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckingV2API.CheckRequestOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckRequestOperation`: ApisResponseV2CheckStatusResponseV2
	fmt.Fprintf(os.Stdout, "Response from `CheckingV2API.CheckRequestOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckRequestOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkTransactionRequest** | [**CheckTransactionRequest**](CheckTransactionRequest.md) |  | 

### Return type

[**ApisResponseV2CheckStatusResponseV2**](ApisResponseV2CheckStatusResponseV2.md)

### Authorization

[ApiKey](../README.md#ApiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

