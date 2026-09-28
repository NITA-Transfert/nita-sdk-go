# PartenaireToWalletFraisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneClient** | **string** | Numero du client MyNITA (wallet cible). | 
**MontantTransaction** | **float64** | Montant de la recharge. | 

## Methods

### NewPartenaireToWalletFraisRequest

`func NewPartenaireToWalletFraisRequest(phoneClient string, montantTransaction float64, ) *PartenaireToWalletFraisRequest`

NewPartenaireToWalletFraisRequest instantiates a new PartenaireToWalletFraisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartenaireToWalletFraisRequestWithDefaults

`func NewPartenaireToWalletFraisRequestWithDefaults() *PartenaireToWalletFraisRequest`

NewPartenaireToWalletFraisRequestWithDefaults instantiates a new PartenaireToWalletFraisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneClient

`func (o *PartenaireToWalletFraisRequest) GetPhoneClient() string`

GetPhoneClient returns the PhoneClient field if non-nil, zero value otherwise.

### GetPhoneClientOk

`func (o *PartenaireToWalletFraisRequest) GetPhoneClientOk() (*string, bool)`

GetPhoneClientOk returns a tuple with the PhoneClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneClient

`func (o *PartenaireToWalletFraisRequest) SetPhoneClient(v string)`

SetPhoneClient sets PhoneClient field to given value.


### GetMontantTransaction

`func (o *PartenaireToWalletFraisRequest) GetMontantTransaction() float64`

GetMontantTransaction returns the MontantTransaction field if non-nil, zero value otherwise.

### GetMontantTransactionOk

`func (o *PartenaireToWalletFraisRequest) GetMontantTransactionOk() (*float64, bool)`

GetMontantTransactionOk returns a tuple with the MontantTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTransaction

`func (o *PartenaireToWalletFraisRequest) SetMontantTransaction(v float64)`

SetMontantTransaction sets MontantTransaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


