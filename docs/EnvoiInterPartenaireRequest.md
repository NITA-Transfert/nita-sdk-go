# EnvoiInterPartenaireRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | Votre référence d&#39;opération (idempotence). | [optional] 
**PartenaireDestinataireAlias** | Pointer to **string** | Alias du partenaire destinataire (même organisation). | [optional] 
**CompteDestinataire** | Pointer to **string** | Libellé du compte destinataire (optionnel si unique). | [optional] 
**MontantTransaction** | Pointer to **float64** |  | [optional] 
**TypeFraisEnvoi** | Pointer to **string** | fraisApars (défaut) ou fraisInclus. | [optional] 
**MotifTransaction** | Pointer to **string** |  | [optional] 
**CompteExpediteur** | Pointer to **string** | Libellé du compte expéditeur (optionnel si unique). | [optional] 

## Methods

### NewEnvoiInterPartenaireRequest

`func NewEnvoiInterPartenaireRequest() *EnvoiInterPartenaireRequest`

NewEnvoiInterPartenaireRequest instantiates a new EnvoiInterPartenaireRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvoiInterPartenaireRequestWithDefaults

`func NewEnvoiInterPartenaireRequestWithDefaults() *EnvoiInterPartenaireRequest`

NewEnvoiInterPartenaireRequestWithDefaults instantiates a new EnvoiInterPartenaireRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EnvoiInterPartenaireRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EnvoiInterPartenaireRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EnvoiInterPartenaireRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EnvoiInterPartenaireRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetPartenaireDestinataireAlias

`func (o *EnvoiInterPartenaireRequest) GetPartenaireDestinataireAlias() string`

GetPartenaireDestinataireAlias returns the PartenaireDestinataireAlias field if non-nil, zero value otherwise.

### GetPartenaireDestinataireAliasOk

`func (o *EnvoiInterPartenaireRequest) GetPartenaireDestinataireAliasOk() (*string, bool)`

GetPartenaireDestinataireAliasOk returns a tuple with the PartenaireDestinataireAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartenaireDestinataireAlias

`func (o *EnvoiInterPartenaireRequest) SetPartenaireDestinataireAlias(v string)`

SetPartenaireDestinataireAlias sets PartenaireDestinataireAlias field to given value.

### HasPartenaireDestinataireAlias

`func (o *EnvoiInterPartenaireRequest) HasPartenaireDestinataireAlias() bool`

HasPartenaireDestinataireAlias returns a boolean if a field has been set.

### GetCompteDestinataire

`func (o *EnvoiInterPartenaireRequest) GetCompteDestinataire() string`

GetCompteDestinataire returns the CompteDestinataire field if non-nil, zero value otherwise.

### GetCompteDestinataireOk

`func (o *EnvoiInterPartenaireRequest) GetCompteDestinataireOk() (*string, bool)`

GetCompteDestinataireOk returns a tuple with the CompteDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompteDestinataire

`func (o *EnvoiInterPartenaireRequest) SetCompteDestinataire(v string)`

SetCompteDestinataire sets CompteDestinataire field to given value.

### HasCompteDestinataire

`func (o *EnvoiInterPartenaireRequest) HasCompteDestinataire() bool`

HasCompteDestinataire returns a boolean if a field has been set.

### GetMontantTransaction

`func (o *EnvoiInterPartenaireRequest) GetMontantTransaction() float64`

GetMontantTransaction returns the MontantTransaction field if non-nil, zero value otherwise.

### GetMontantTransactionOk

`func (o *EnvoiInterPartenaireRequest) GetMontantTransactionOk() (*float64, bool)`

GetMontantTransactionOk returns a tuple with the MontantTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTransaction

`func (o *EnvoiInterPartenaireRequest) SetMontantTransaction(v float64)`

SetMontantTransaction sets MontantTransaction field to given value.

### HasMontantTransaction

`func (o *EnvoiInterPartenaireRequest) HasMontantTransaction() bool`

HasMontantTransaction returns a boolean if a field has been set.

### GetTypeFraisEnvoi

`func (o *EnvoiInterPartenaireRequest) GetTypeFraisEnvoi() string`

GetTypeFraisEnvoi returns the TypeFraisEnvoi field if non-nil, zero value otherwise.

### GetTypeFraisEnvoiOk

`func (o *EnvoiInterPartenaireRequest) GetTypeFraisEnvoiOk() (*string, bool)`

GetTypeFraisEnvoiOk returns a tuple with the TypeFraisEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFraisEnvoi

`func (o *EnvoiInterPartenaireRequest) SetTypeFraisEnvoi(v string)`

SetTypeFraisEnvoi sets TypeFraisEnvoi field to given value.

### HasTypeFraisEnvoi

`func (o *EnvoiInterPartenaireRequest) HasTypeFraisEnvoi() bool`

HasTypeFraisEnvoi returns a boolean if a field has been set.

### GetMotifTransaction

`func (o *EnvoiInterPartenaireRequest) GetMotifTransaction() string`

GetMotifTransaction returns the MotifTransaction field if non-nil, zero value otherwise.

### GetMotifTransactionOk

`func (o *EnvoiInterPartenaireRequest) GetMotifTransactionOk() (*string, bool)`

GetMotifTransactionOk returns a tuple with the MotifTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifTransaction

`func (o *EnvoiInterPartenaireRequest) SetMotifTransaction(v string)`

SetMotifTransaction sets MotifTransaction field to given value.

### HasMotifTransaction

`func (o *EnvoiInterPartenaireRequest) HasMotifTransaction() bool`

HasMotifTransaction returns a boolean if a field has been set.

### GetCompteExpediteur

`func (o *EnvoiInterPartenaireRequest) GetCompteExpediteur() string`

GetCompteExpediteur returns the CompteExpediteur field if non-nil, zero value otherwise.

### GetCompteExpediteurOk

`func (o *EnvoiInterPartenaireRequest) GetCompteExpediteurOk() (*string, bool)`

GetCompteExpediteurOk returns a tuple with the CompteExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompteExpediteur

`func (o *EnvoiInterPartenaireRequest) SetCompteExpediteur(v string)`

SetCompteExpediteur sets CompteExpediteur field to given value.

### HasCompteExpediteur

`func (o *EnvoiInterPartenaireRequest) HasCompteExpediteur() bool`

HasCompteExpediteur returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


