# CalculeFraisOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndicatifExpediteur** | **string** |  | 
**NumeroExpediteur** | **string** |  | 
**Montant** | **float64** |  | 
**FraisInclus** | **bool** |  | 
**VilleExpedition** | **string** |  | 
**VilleDestination** | **string** |  | 
**Profil** | Pointer to **string** | Grille tarifaire du reseau de retrait cash : CCP (Niger Poste, defaut) ou BIN. | [optional] [default to "CCP"]

## Methods

### NewCalculeFraisOperationRequest

`func NewCalculeFraisOperationRequest(indicatifExpediteur string, numeroExpediteur string, montant float64, fraisInclus bool, villeExpedition string, villeDestination string, ) *CalculeFraisOperationRequest`

NewCalculeFraisOperationRequest instantiates a new CalculeFraisOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculeFraisOperationRequestWithDefaults

`func NewCalculeFraisOperationRequestWithDefaults() *CalculeFraisOperationRequest`

NewCalculeFraisOperationRequestWithDefaults instantiates a new CalculeFraisOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicatifExpediteur

`func (o *CalculeFraisOperationRequest) GetIndicatifExpediteur() string`

GetIndicatifExpediteur returns the IndicatifExpediteur field if non-nil, zero value otherwise.

### GetIndicatifExpediteurOk

`func (o *CalculeFraisOperationRequest) GetIndicatifExpediteurOk() (*string, bool)`

GetIndicatifExpediteurOk returns a tuple with the IndicatifExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatifExpediteur

`func (o *CalculeFraisOperationRequest) SetIndicatifExpediteur(v string)`

SetIndicatifExpediteur sets IndicatifExpediteur field to given value.


### GetNumeroExpediteur

`func (o *CalculeFraisOperationRequest) GetNumeroExpediteur() string`

GetNumeroExpediteur returns the NumeroExpediteur field if non-nil, zero value otherwise.

### GetNumeroExpediteurOk

`func (o *CalculeFraisOperationRequest) GetNumeroExpediteurOk() (*string, bool)`

GetNumeroExpediteurOk returns a tuple with the NumeroExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroExpediteur

`func (o *CalculeFraisOperationRequest) SetNumeroExpediteur(v string)`

SetNumeroExpediteur sets NumeroExpediteur field to given value.


### GetMontant

`func (o *CalculeFraisOperationRequest) GetMontant() float64`

GetMontant returns the Montant field if non-nil, zero value otherwise.

### GetMontantOk

`func (o *CalculeFraisOperationRequest) GetMontantOk() (*float64, bool)`

GetMontantOk returns a tuple with the Montant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontant

`func (o *CalculeFraisOperationRequest) SetMontant(v float64)`

SetMontant sets Montant field to given value.


### GetFraisInclus

`func (o *CalculeFraisOperationRequest) GetFraisInclus() bool`

GetFraisInclus returns the FraisInclus field if non-nil, zero value otherwise.

### GetFraisInclusOk

`func (o *CalculeFraisOperationRequest) GetFraisInclusOk() (*bool, bool)`

GetFraisInclusOk returns a tuple with the FraisInclus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraisInclus

`func (o *CalculeFraisOperationRequest) SetFraisInclus(v bool)`

SetFraisInclus sets FraisInclus field to given value.


### GetVilleExpedition

`func (o *CalculeFraisOperationRequest) GetVilleExpedition() string`

GetVilleExpedition returns the VilleExpedition field if non-nil, zero value otherwise.

### GetVilleExpeditionOk

`func (o *CalculeFraisOperationRequest) GetVilleExpeditionOk() (*string, bool)`

GetVilleExpeditionOk returns a tuple with the VilleExpedition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleExpedition

`func (o *CalculeFraisOperationRequest) SetVilleExpedition(v string)`

SetVilleExpedition sets VilleExpedition field to given value.


### GetVilleDestination

`func (o *CalculeFraisOperationRequest) GetVilleDestination() string`

GetVilleDestination returns the VilleDestination field if non-nil, zero value otherwise.

### GetVilleDestinationOk

`func (o *CalculeFraisOperationRequest) GetVilleDestinationOk() (*string, bool)`

GetVilleDestinationOk returns a tuple with the VilleDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDestination

`func (o *CalculeFraisOperationRequest) SetVilleDestination(v string)`

SetVilleDestination sets VilleDestination field to given value.


### GetProfil

`func (o *CalculeFraisOperationRequest) GetProfil() string`

GetProfil returns the Profil field if non-nil, zero value otherwise.

### GetProfilOk

`func (o *CalculeFraisOperationRequest) GetProfilOk() (*string, bool)`

GetProfilOk returns a tuple with the Profil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfil

`func (o *CalculeFraisOperationRequest) SetProfil(v string)`

SetProfil sets Profil field to given value.

### HasProfil

`func (o *CalculeFraisOperationRequest) HasProfil() bool`

HasProfil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


