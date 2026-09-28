# TransactionResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicId** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**CodeEnvoi** | Pointer to **string** |  | [optional] 
**Montant** | Pointer to **float64** |  | [optional] 
**FraisEnvoi** | Pointer to **float64** |  | [optional] 
**Reduction** | Pointer to **float64** |  | [optional] 
**FraisInclus** | Pointer to **bool** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**StatusTransaction** | Pointer to **string** |  | [optional] 
**TypeTransaction** | Pointer to **string** |  | [optional] 
**TypeOperation** | Pointer to **string** |  | [optional] 
**NomExpediteur** | Pointer to **string** |  | [optional] 
**PrenomExpediteur** | Pointer to **string** |  | [optional] 
**TelephoneExpediteur** | Pointer to **string** |  | [optional] 
**NomDestinataire** | Pointer to **string** |  | [optional] 
**PrenomDestinataire** | Pointer to **string** |  | [optional] 
**TelephoneDestinataire** | Pointer to **string** |  | [optional] 
**VilleDestinataire** | Pointer to **string** |  | [optional] 
**PaysDestinataire** | Pointer to **string** |  | [optional] 
**DateTransaction** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTransactionResponseV2

`func NewTransactionResponseV2() *TransactionResponseV2`

NewTransactionResponseV2 instantiates a new TransactionResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseV2WithDefaults

`func NewTransactionResponseV2WithDefaults() *TransactionResponseV2`

NewTransactionResponseV2WithDefaults instantiates a new TransactionResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicId

`func (o *TransactionResponseV2) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *TransactionResponseV2) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *TransactionResponseV2) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *TransactionResponseV2) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetRequestId

`func (o *TransactionResponseV2) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionResponseV2) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionResponseV2) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TransactionResponseV2) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCodeEnvoi

`func (o *TransactionResponseV2) GetCodeEnvoi() string`

GetCodeEnvoi returns the CodeEnvoi field if non-nil, zero value otherwise.

### GetCodeEnvoiOk

`func (o *TransactionResponseV2) GetCodeEnvoiOk() (*string, bool)`

GetCodeEnvoiOk returns a tuple with the CodeEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEnvoi

`func (o *TransactionResponseV2) SetCodeEnvoi(v string)`

SetCodeEnvoi sets CodeEnvoi field to given value.

### HasCodeEnvoi

`func (o *TransactionResponseV2) HasCodeEnvoi() bool`

HasCodeEnvoi returns a boolean if a field has been set.

### GetMontant

`func (o *TransactionResponseV2) GetMontant() float64`

GetMontant returns the Montant field if non-nil, zero value otherwise.

### GetMontantOk

`func (o *TransactionResponseV2) GetMontantOk() (*float64, bool)`

GetMontantOk returns a tuple with the Montant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontant

`func (o *TransactionResponseV2) SetMontant(v float64)`

SetMontant sets Montant field to given value.

### HasMontant

`func (o *TransactionResponseV2) HasMontant() bool`

HasMontant returns a boolean if a field has been set.

### GetFraisEnvoi

`func (o *TransactionResponseV2) GetFraisEnvoi() float64`

GetFraisEnvoi returns the FraisEnvoi field if non-nil, zero value otherwise.

### GetFraisEnvoiOk

`func (o *TransactionResponseV2) GetFraisEnvoiOk() (*float64, bool)`

GetFraisEnvoiOk returns a tuple with the FraisEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraisEnvoi

`func (o *TransactionResponseV2) SetFraisEnvoi(v float64)`

SetFraisEnvoi sets FraisEnvoi field to given value.

### HasFraisEnvoi

`func (o *TransactionResponseV2) HasFraisEnvoi() bool`

HasFraisEnvoi returns a boolean if a field has been set.

### GetReduction

`func (o *TransactionResponseV2) GetReduction() float64`

GetReduction returns the Reduction field if non-nil, zero value otherwise.

### GetReductionOk

`func (o *TransactionResponseV2) GetReductionOk() (*float64, bool)`

GetReductionOk returns a tuple with the Reduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduction

`func (o *TransactionResponseV2) SetReduction(v float64)`

SetReduction sets Reduction field to given value.

### HasReduction

`func (o *TransactionResponseV2) HasReduction() bool`

HasReduction returns a boolean if a field has been set.

### GetFraisInclus

`func (o *TransactionResponseV2) GetFraisInclus() bool`

GetFraisInclus returns the FraisInclus field if non-nil, zero value otherwise.

### GetFraisInclusOk

`func (o *TransactionResponseV2) GetFraisInclusOk() (*bool, bool)`

GetFraisInclusOk returns a tuple with the FraisInclus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraisInclus

`func (o *TransactionResponseV2) SetFraisInclus(v bool)`

SetFraisInclus sets FraisInclus field to given value.

### HasFraisInclus

`func (o *TransactionResponseV2) HasFraisInclus() bool`

HasFraisInclus returns a boolean if a field has been set.

### GetTotal

`func (o *TransactionResponseV2) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TransactionResponseV2) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TransactionResponseV2) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TransactionResponseV2) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetStatusTransaction

`func (o *TransactionResponseV2) GetStatusTransaction() string`

GetStatusTransaction returns the StatusTransaction field if non-nil, zero value otherwise.

### GetStatusTransactionOk

`func (o *TransactionResponseV2) GetStatusTransactionOk() (*string, bool)`

GetStatusTransactionOk returns a tuple with the StatusTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTransaction

`func (o *TransactionResponseV2) SetStatusTransaction(v string)`

SetStatusTransaction sets StatusTransaction field to given value.

### HasStatusTransaction

`func (o *TransactionResponseV2) HasStatusTransaction() bool`

HasStatusTransaction returns a boolean if a field has been set.

### GetTypeTransaction

`func (o *TransactionResponseV2) GetTypeTransaction() string`

GetTypeTransaction returns the TypeTransaction field if non-nil, zero value otherwise.

### GetTypeTransactionOk

`func (o *TransactionResponseV2) GetTypeTransactionOk() (*string, bool)`

GetTypeTransactionOk returns a tuple with the TypeTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeTransaction

`func (o *TransactionResponseV2) SetTypeTransaction(v string)`

SetTypeTransaction sets TypeTransaction field to given value.

### HasTypeTransaction

`func (o *TransactionResponseV2) HasTypeTransaction() bool`

HasTypeTransaction returns a boolean if a field has been set.

### GetTypeOperation

`func (o *TransactionResponseV2) GetTypeOperation() string`

GetTypeOperation returns the TypeOperation field if non-nil, zero value otherwise.

### GetTypeOperationOk

`func (o *TransactionResponseV2) GetTypeOperationOk() (*string, bool)`

GetTypeOperationOk returns a tuple with the TypeOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOperation

`func (o *TransactionResponseV2) SetTypeOperation(v string)`

SetTypeOperation sets TypeOperation field to given value.

### HasTypeOperation

`func (o *TransactionResponseV2) HasTypeOperation() bool`

HasTypeOperation returns a boolean if a field has been set.

### GetNomExpediteur

`func (o *TransactionResponseV2) GetNomExpediteur() string`

GetNomExpediteur returns the NomExpediteur field if non-nil, zero value otherwise.

### GetNomExpediteurOk

`func (o *TransactionResponseV2) GetNomExpediteurOk() (*string, bool)`

GetNomExpediteurOk returns a tuple with the NomExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomExpediteur

`func (o *TransactionResponseV2) SetNomExpediteur(v string)`

SetNomExpediteur sets NomExpediteur field to given value.

### HasNomExpediteur

`func (o *TransactionResponseV2) HasNomExpediteur() bool`

HasNomExpediteur returns a boolean if a field has been set.

### GetPrenomExpediteur

`func (o *TransactionResponseV2) GetPrenomExpediteur() string`

GetPrenomExpediteur returns the PrenomExpediteur field if non-nil, zero value otherwise.

### GetPrenomExpediteurOk

`func (o *TransactionResponseV2) GetPrenomExpediteurOk() (*string, bool)`

GetPrenomExpediteurOk returns a tuple with the PrenomExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomExpediteur

`func (o *TransactionResponseV2) SetPrenomExpediteur(v string)`

SetPrenomExpediteur sets PrenomExpediteur field to given value.

### HasPrenomExpediteur

`func (o *TransactionResponseV2) HasPrenomExpediteur() bool`

HasPrenomExpediteur returns a boolean if a field has been set.

### GetTelephoneExpediteur

`func (o *TransactionResponseV2) GetTelephoneExpediteur() string`

GetTelephoneExpediteur returns the TelephoneExpediteur field if non-nil, zero value otherwise.

### GetTelephoneExpediteurOk

`func (o *TransactionResponseV2) GetTelephoneExpediteurOk() (*string, bool)`

GetTelephoneExpediteurOk returns a tuple with the TelephoneExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneExpediteur

`func (o *TransactionResponseV2) SetTelephoneExpediteur(v string)`

SetTelephoneExpediteur sets TelephoneExpediteur field to given value.

### HasTelephoneExpediteur

`func (o *TransactionResponseV2) HasTelephoneExpediteur() bool`

HasTelephoneExpediteur returns a boolean if a field has been set.

### GetNomDestinataire

`func (o *TransactionResponseV2) GetNomDestinataire() string`

GetNomDestinataire returns the NomDestinataire field if non-nil, zero value otherwise.

### GetNomDestinataireOk

`func (o *TransactionResponseV2) GetNomDestinataireOk() (*string, bool)`

GetNomDestinataireOk returns a tuple with the NomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomDestinataire

`func (o *TransactionResponseV2) SetNomDestinataire(v string)`

SetNomDestinataire sets NomDestinataire field to given value.

### HasNomDestinataire

`func (o *TransactionResponseV2) HasNomDestinataire() bool`

HasNomDestinataire returns a boolean if a field has been set.

### GetPrenomDestinataire

`func (o *TransactionResponseV2) GetPrenomDestinataire() string`

GetPrenomDestinataire returns the PrenomDestinataire field if non-nil, zero value otherwise.

### GetPrenomDestinataireOk

`func (o *TransactionResponseV2) GetPrenomDestinataireOk() (*string, bool)`

GetPrenomDestinataireOk returns a tuple with the PrenomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomDestinataire

`func (o *TransactionResponseV2) SetPrenomDestinataire(v string)`

SetPrenomDestinataire sets PrenomDestinataire field to given value.

### HasPrenomDestinataire

`func (o *TransactionResponseV2) HasPrenomDestinataire() bool`

HasPrenomDestinataire returns a boolean if a field has been set.

### GetTelephoneDestinataire

`func (o *TransactionResponseV2) GetTelephoneDestinataire() string`

GetTelephoneDestinataire returns the TelephoneDestinataire field if non-nil, zero value otherwise.

### GetTelephoneDestinataireOk

`func (o *TransactionResponseV2) GetTelephoneDestinataireOk() (*string, bool)`

GetTelephoneDestinataireOk returns a tuple with the TelephoneDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneDestinataire

`func (o *TransactionResponseV2) SetTelephoneDestinataire(v string)`

SetTelephoneDestinataire sets TelephoneDestinataire field to given value.

### HasTelephoneDestinataire

`func (o *TransactionResponseV2) HasTelephoneDestinataire() bool`

HasTelephoneDestinataire returns a boolean if a field has been set.

### GetVilleDestinataire

`func (o *TransactionResponseV2) GetVilleDestinataire() string`

GetVilleDestinataire returns the VilleDestinataire field if non-nil, zero value otherwise.

### GetVilleDestinataireOk

`func (o *TransactionResponseV2) GetVilleDestinataireOk() (*string, bool)`

GetVilleDestinataireOk returns a tuple with the VilleDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDestinataire

`func (o *TransactionResponseV2) SetVilleDestinataire(v string)`

SetVilleDestinataire sets VilleDestinataire field to given value.

### HasVilleDestinataire

`func (o *TransactionResponseV2) HasVilleDestinataire() bool`

HasVilleDestinataire returns a boolean if a field has been set.

### GetPaysDestinataire

`func (o *TransactionResponseV2) GetPaysDestinataire() string`

GetPaysDestinataire returns the PaysDestinataire field if non-nil, zero value otherwise.

### GetPaysDestinataireOk

`func (o *TransactionResponseV2) GetPaysDestinataireOk() (*string, bool)`

GetPaysDestinataireOk returns a tuple with the PaysDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDestinataire

`func (o *TransactionResponseV2) SetPaysDestinataire(v string)`

SetPaysDestinataire sets PaysDestinataire field to given value.

### HasPaysDestinataire

`func (o *TransactionResponseV2) HasPaysDestinataire() bool`

HasPaysDestinataire returns a boolean if a field has been set.

### GetDateTransaction

`func (o *TransactionResponseV2) GetDateTransaction() time.Time`

GetDateTransaction returns the DateTransaction field if non-nil, zero value otherwise.

### GetDateTransactionOk

`func (o *TransactionResponseV2) GetDateTransactionOk() (*time.Time, bool)`

GetDateTransactionOk returns a tuple with the DateTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTransaction

`func (o *TransactionResponseV2) SetDateTransaction(v time.Time)`

SetDateTransaction sets DateTransaction field to given value.

### HasDateTransaction

`func (o *TransactionResponseV2) HasDateTransaction() bool`

HasDateTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


