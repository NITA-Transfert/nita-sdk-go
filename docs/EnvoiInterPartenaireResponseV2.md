# EnvoiInterPartenaireResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicId** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**CodeEnvoi** | Pointer to **string** |  | [optional] 
**MontantTransaction** | Pointer to **float64** |  | [optional] 
**MontantFrais** | Pointer to **float64** |  | [optional] 
**MontantDebite** | Pointer to **float64** |  | [optional] 
**MontantRecu** | Pointer to **float64** |  | [optional] 
**StatusTransaction** | Pointer to **string** |  | [optional] 
**CompteExpediteur** | Pointer to **string** |  | [optional] 
**CompteDestinataire** | Pointer to **string** |  | [optional] 
**DateTransaction** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvoiInterPartenaireResponseV2

`func NewEnvoiInterPartenaireResponseV2() *EnvoiInterPartenaireResponseV2`

NewEnvoiInterPartenaireResponseV2 instantiates a new EnvoiInterPartenaireResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvoiInterPartenaireResponseV2WithDefaults

`func NewEnvoiInterPartenaireResponseV2WithDefaults() *EnvoiInterPartenaireResponseV2`

NewEnvoiInterPartenaireResponseV2WithDefaults instantiates a new EnvoiInterPartenaireResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicId

`func (o *EnvoiInterPartenaireResponseV2) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *EnvoiInterPartenaireResponseV2) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *EnvoiInterPartenaireResponseV2) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *EnvoiInterPartenaireResponseV2) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetRequestId

`func (o *EnvoiInterPartenaireResponseV2) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EnvoiInterPartenaireResponseV2) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EnvoiInterPartenaireResponseV2) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EnvoiInterPartenaireResponseV2) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCodeEnvoi

`func (o *EnvoiInterPartenaireResponseV2) GetCodeEnvoi() string`

GetCodeEnvoi returns the CodeEnvoi field if non-nil, zero value otherwise.

### GetCodeEnvoiOk

`func (o *EnvoiInterPartenaireResponseV2) GetCodeEnvoiOk() (*string, bool)`

GetCodeEnvoiOk returns a tuple with the CodeEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEnvoi

`func (o *EnvoiInterPartenaireResponseV2) SetCodeEnvoi(v string)`

SetCodeEnvoi sets CodeEnvoi field to given value.

### HasCodeEnvoi

`func (o *EnvoiInterPartenaireResponseV2) HasCodeEnvoi() bool`

HasCodeEnvoi returns a boolean if a field has been set.

### GetMontantTransaction

`func (o *EnvoiInterPartenaireResponseV2) GetMontantTransaction() float64`

GetMontantTransaction returns the MontantTransaction field if non-nil, zero value otherwise.

### GetMontantTransactionOk

`func (o *EnvoiInterPartenaireResponseV2) GetMontantTransactionOk() (*float64, bool)`

GetMontantTransactionOk returns a tuple with the MontantTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTransaction

`func (o *EnvoiInterPartenaireResponseV2) SetMontantTransaction(v float64)`

SetMontantTransaction sets MontantTransaction field to given value.

### HasMontantTransaction

`func (o *EnvoiInterPartenaireResponseV2) HasMontantTransaction() bool`

HasMontantTransaction returns a boolean if a field has been set.

### GetMontantFrais

`func (o *EnvoiInterPartenaireResponseV2) GetMontantFrais() float64`

GetMontantFrais returns the MontantFrais field if non-nil, zero value otherwise.

### GetMontantFraisOk

`func (o *EnvoiInterPartenaireResponseV2) GetMontantFraisOk() (*float64, bool)`

GetMontantFraisOk returns a tuple with the MontantFrais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantFrais

`func (o *EnvoiInterPartenaireResponseV2) SetMontantFrais(v float64)`

SetMontantFrais sets MontantFrais field to given value.

### HasMontantFrais

`func (o *EnvoiInterPartenaireResponseV2) HasMontantFrais() bool`

HasMontantFrais returns a boolean if a field has been set.

### GetMontantDebite

`func (o *EnvoiInterPartenaireResponseV2) GetMontantDebite() float64`

GetMontantDebite returns the MontantDebite field if non-nil, zero value otherwise.

### GetMontantDebiteOk

`func (o *EnvoiInterPartenaireResponseV2) GetMontantDebiteOk() (*float64, bool)`

GetMontantDebiteOk returns a tuple with the MontantDebite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantDebite

`func (o *EnvoiInterPartenaireResponseV2) SetMontantDebite(v float64)`

SetMontantDebite sets MontantDebite field to given value.

### HasMontantDebite

`func (o *EnvoiInterPartenaireResponseV2) HasMontantDebite() bool`

HasMontantDebite returns a boolean if a field has been set.

### GetMontantRecu

`func (o *EnvoiInterPartenaireResponseV2) GetMontantRecu() float64`

GetMontantRecu returns the MontantRecu field if non-nil, zero value otherwise.

### GetMontantRecuOk

`func (o *EnvoiInterPartenaireResponseV2) GetMontantRecuOk() (*float64, bool)`

GetMontantRecuOk returns a tuple with the MontantRecu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantRecu

`func (o *EnvoiInterPartenaireResponseV2) SetMontantRecu(v float64)`

SetMontantRecu sets MontantRecu field to given value.

### HasMontantRecu

`func (o *EnvoiInterPartenaireResponseV2) HasMontantRecu() bool`

HasMontantRecu returns a boolean if a field has been set.

### GetStatusTransaction

`func (o *EnvoiInterPartenaireResponseV2) GetStatusTransaction() string`

GetStatusTransaction returns the StatusTransaction field if non-nil, zero value otherwise.

### GetStatusTransactionOk

`func (o *EnvoiInterPartenaireResponseV2) GetStatusTransactionOk() (*string, bool)`

GetStatusTransactionOk returns a tuple with the StatusTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTransaction

`func (o *EnvoiInterPartenaireResponseV2) SetStatusTransaction(v string)`

SetStatusTransaction sets StatusTransaction field to given value.

### HasStatusTransaction

`func (o *EnvoiInterPartenaireResponseV2) HasStatusTransaction() bool`

HasStatusTransaction returns a boolean if a field has been set.

### GetCompteExpediteur

`func (o *EnvoiInterPartenaireResponseV2) GetCompteExpediteur() string`

GetCompteExpediteur returns the CompteExpediteur field if non-nil, zero value otherwise.

### GetCompteExpediteurOk

`func (o *EnvoiInterPartenaireResponseV2) GetCompteExpediteurOk() (*string, bool)`

GetCompteExpediteurOk returns a tuple with the CompteExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompteExpediteur

`func (o *EnvoiInterPartenaireResponseV2) SetCompteExpediteur(v string)`

SetCompteExpediteur sets CompteExpediteur field to given value.

### HasCompteExpediteur

`func (o *EnvoiInterPartenaireResponseV2) HasCompteExpediteur() bool`

HasCompteExpediteur returns a boolean if a field has been set.

### GetCompteDestinataire

`func (o *EnvoiInterPartenaireResponseV2) GetCompteDestinataire() string`

GetCompteDestinataire returns the CompteDestinataire field if non-nil, zero value otherwise.

### GetCompteDestinataireOk

`func (o *EnvoiInterPartenaireResponseV2) GetCompteDestinataireOk() (*string, bool)`

GetCompteDestinataireOk returns a tuple with the CompteDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompteDestinataire

`func (o *EnvoiInterPartenaireResponseV2) SetCompteDestinataire(v string)`

SetCompteDestinataire sets CompteDestinataire field to given value.

### HasCompteDestinataire

`func (o *EnvoiInterPartenaireResponseV2) HasCompteDestinataire() bool`

HasCompteDestinataire returns a boolean if a field has been set.

### GetDateTransaction

`func (o *EnvoiInterPartenaireResponseV2) GetDateTransaction() time.Time`

GetDateTransaction returns the DateTransaction field if non-nil, zero value otherwise.

### GetDateTransactionOk

`func (o *EnvoiInterPartenaireResponseV2) GetDateTransactionOk() (*time.Time, bool)`

GetDateTransactionOk returns a tuple with the DateTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTransaction

`func (o *EnvoiInterPartenaireResponseV2) SetDateTransaction(v time.Time)`

SetDateTransaction sets DateTransaction field to given value.

### HasDateTransaction

`func (o *EnvoiInterPartenaireResponseV2) HasDateTransaction() bool`

HasDateTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


