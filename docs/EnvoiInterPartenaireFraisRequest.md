# EnvoiInterPartenaireFraisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartenaireDestinataireAlias** | **string** | Alias du partenaire destinataire. | 
**CompteExpediteur** | Pointer to **string** | Libelle du compte expediteur (si plusieurs comptes). | [optional] 
**CompteDestinataire** | Pointer to **string** | Libelle du compte destinataire (si plusieurs comptes). | [optional] 
**MontantTransaction** | **float64** | Montant de l&#39;envoi. | 
**TypeFraisEnvoi** | Pointer to **string** | fraisApars (frais en sus, defaut) ou fraisInclus (frais deduits du montant). | [optional] [default to "fraisApars"]

## Methods

### NewEnvoiInterPartenaireFraisRequest

`func NewEnvoiInterPartenaireFraisRequest(partenaireDestinataireAlias string, montantTransaction float64, ) *EnvoiInterPartenaireFraisRequest`

NewEnvoiInterPartenaireFraisRequest instantiates a new EnvoiInterPartenaireFraisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvoiInterPartenaireFraisRequestWithDefaults

`func NewEnvoiInterPartenaireFraisRequestWithDefaults() *EnvoiInterPartenaireFraisRequest`

NewEnvoiInterPartenaireFraisRequestWithDefaults instantiates a new EnvoiInterPartenaireFraisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartenaireDestinataireAlias

`func (o *EnvoiInterPartenaireFraisRequest) GetPartenaireDestinataireAlias() string`

GetPartenaireDestinataireAlias returns the PartenaireDestinataireAlias field if non-nil, zero value otherwise.

### GetPartenaireDestinataireAliasOk

`func (o *EnvoiInterPartenaireFraisRequest) GetPartenaireDestinataireAliasOk() (*string, bool)`

GetPartenaireDestinataireAliasOk returns a tuple with the PartenaireDestinataireAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartenaireDestinataireAlias

`func (o *EnvoiInterPartenaireFraisRequest) SetPartenaireDestinataireAlias(v string)`

SetPartenaireDestinataireAlias sets PartenaireDestinataireAlias field to given value.


### GetCompteExpediteur

`func (o *EnvoiInterPartenaireFraisRequest) GetCompteExpediteur() string`

GetCompteExpediteur returns the CompteExpediteur field if non-nil, zero value otherwise.

### GetCompteExpediteurOk

`func (o *EnvoiInterPartenaireFraisRequest) GetCompteExpediteurOk() (*string, bool)`

GetCompteExpediteurOk returns a tuple with the CompteExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompteExpediteur

`func (o *EnvoiInterPartenaireFraisRequest) SetCompteExpediteur(v string)`

SetCompteExpediteur sets CompteExpediteur field to given value.

### HasCompteExpediteur

`func (o *EnvoiInterPartenaireFraisRequest) HasCompteExpediteur() bool`

HasCompteExpediteur returns a boolean if a field has been set.

### GetCompteDestinataire

`func (o *EnvoiInterPartenaireFraisRequest) GetCompteDestinataire() string`

GetCompteDestinataire returns the CompteDestinataire field if non-nil, zero value otherwise.

### GetCompteDestinataireOk

`func (o *EnvoiInterPartenaireFraisRequest) GetCompteDestinataireOk() (*string, bool)`

GetCompteDestinataireOk returns a tuple with the CompteDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompteDestinataire

`func (o *EnvoiInterPartenaireFraisRequest) SetCompteDestinataire(v string)`

SetCompteDestinataire sets CompteDestinataire field to given value.

### HasCompteDestinataire

`func (o *EnvoiInterPartenaireFraisRequest) HasCompteDestinataire() bool`

HasCompteDestinataire returns a boolean if a field has been set.

### GetMontantTransaction

`func (o *EnvoiInterPartenaireFraisRequest) GetMontantTransaction() float64`

GetMontantTransaction returns the MontantTransaction field if non-nil, zero value otherwise.

### GetMontantTransactionOk

`func (o *EnvoiInterPartenaireFraisRequest) GetMontantTransactionOk() (*float64, bool)`

GetMontantTransactionOk returns a tuple with the MontantTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTransaction

`func (o *EnvoiInterPartenaireFraisRequest) SetMontantTransaction(v float64)`

SetMontantTransaction sets MontantTransaction field to given value.


### GetTypeFraisEnvoi

`func (o *EnvoiInterPartenaireFraisRequest) GetTypeFraisEnvoi() string`

GetTypeFraisEnvoi returns the TypeFraisEnvoi field if non-nil, zero value otherwise.

### GetTypeFraisEnvoiOk

`func (o *EnvoiInterPartenaireFraisRequest) GetTypeFraisEnvoiOk() (*string, bool)`

GetTypeFraisEnvoiOk returns a tuple with the TypeFraisEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFraisEnvoi

`func (o *EnvoiInterPartenaireFraisRequest) SetTypeFraisEnvoi(v string)`

SetTypeFraisEnvoi sets TypeFraisEnvoi field to given value.

### HasTypeFraisEnvoi

`func (o *EnvoiInterPartenaireFraisRequest) HasTypeFraisEnvoi() bool`

HasTypeFraisEnvoi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


