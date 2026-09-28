# FraisPreviewResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Montant** | Pointer to **float64** | Montant de la transaction demande. | [optional] 
**Frais** | Pointer to **float64** | Frais NITA applicables. | [optional] 
**Timbre** | Pointer to **float64** | Timbre fiscal (0 si non applicable). | [optional] 
**MontantReduction** | Pointer to **float64** | Reduction appliquee (0 si aucune). | [optional] 
**MontantDebite** | Pointer to **float64** | Total debite du compte partenaire. | [optional] 
**MontantRecu** | Pointer to **float64** | Montant net recu par le beneficiaire. | [optional] 
**FraisInclus** | Pointer to **bool** | true &#x3D; frais deduits du montant, false &#x3D; frais en sus. | [optional] 
**Devise** | Pointer to **string** | Devise des montants (XOF). | [optional] 

## Methods

### NewFraisPreviewResponseV2

`func NewFraisPreviewResponseV2() *FraisPreviewResponseV2`

NewFraisPreviewResponseV2 instantiates a new FraisPreviewResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFraisPreviewResponseV2WithDefaults

`func NewFraisPreviewResponseV2WithDefaults() *FraisPreviewResponseV2`

NewFraisPreviewResponseV2WithDefaults instantiates a new FraisPreviewResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMontant

`func (o *FraisPreviewResponseV2) GetMontant() float64`

GetMontant returns the Montant field if non-nil, zero value otherwise.

### GetMontantOk

`func (o *FraisPreviewResponseV2) GetMontantOk() (*float64, bool)`

GetMontantOk returns a tuple with the Montant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontant

`func (o *FraisPreviewResponseV2) SetMontant(v float64)`

SetMontant sets Montant field to given value.

### HasMontant

`func (o *FraisPreviewResponseV2) HasMontant() bool`

HasMontant returns a boolean if a field has been set.

### GetFrais

`func (o *FraisPreviewResponseV2) GetFrais() float64`

GetFrais returns the Frais field if non-nil, zero value otherwise.

### GetFraisOk

`func (o *FraisPreviewResponseV2) GetFraisOk() (*float64, bool)`

GetFraisOk returns a tuple with the Frais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrais

`func (o *FraisPreviewResponseV2) SetFrais(v float64)`

SetFrais sets Frais field to given value.

### HasFrais

`func (o *FraisPreviewResponseV2) HasFrais() bool`

HasFrais returns a boolean if a field has been set.

### GetTimbre

`func (o *FraisPreviewResponseV2) GetTimbre() float64`

GetTimbre returns the Timbre field if non-nil, zero value otherwise.

### GetTimbreOk

`func (o *FraisPreviewResponseV2) GetTimbreOk() (*float64, bool)`

GetTimbreOk returns a tuple with the Timbre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimbre

`func (o *FraisPreviewResponseV2) SetTimbre(v float64)`

SetTimbre sets Timbre field to given value.

### HasTimbre

`func (o *FraisPreviewResponseV2) HasTimbre() bool`

HasTimbre returns a boolean if a field has been set.

### GetMontantReduction

`func (o *FraisPreviewResponseV2) GetMontantReduction() float64`

GetMontantReduction returns the MontantReduction field if non-nil, zero value otherwise.

### GetMontantReductionOk

`func (o *FraisPreviewResponseV2) GetMontantReductionOk() (*float64, bool)`

GetMontantReductionOk returns a tuple with the MontantReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantReduction

`func (o *FraisPreviewResponseV2) SetMontantReduction(v float64)`

SetMontantReduction sets MontantReduction field to given value.

### HasMontantReduction

`func (o *FraisPreviewResponseV2) HasMontantReduction() bool`

HasMontantReduction returns a boolean if a field has been set.

### GetMontantDebite

`func (o *FraisPreviewResponseV2) GetMontantDebite() float64`

GetMontantDebite returns the MontantDebite field if non-nil, zero value otherwise.

### GetMontantDebiteOk

`func (o *FraisPreviewResponseV2) GetMontantDebiteOk() (*float64, bool)`

GetMontantDebiteOk returns a tuple with the MontantDebite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantDebite

`func (o *FraisPreviewResponseV2) SetMontantDebite(v float64)`

SetMontantDebite sets MontantDebite field to given value.

### HasMontantDebite

`func (o *FraisPreviewResponseV2) HasMontantDebite() bool`

HasMontantDebite returns a boolean if a field has been set.

### GetMontantRecu

`func (o *FraisPreviewResponseV2) GetMontantRecu() float64`

GetMontantRecu returns the MontantRecu field if non-nil, zero value otherwise.

### GetMontantRecuOk

`func (o *FraisPreviewResponseV2) GetMontantRecuOk() (*float64, bool)`

GetMontantRecuOk returns a tuple with the MontantRecu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantRecu

`func (o *FraisPreviewResponseV2) SetMontantRecu(v float64)`

SetMontantRecu sets MontantRecu field to given value.

### HasMontantRecu

`func (o *FraisPreviewResponseV2) HasMontantRecu() bool`

HasMontantRecu returns a boolean if a field has been set.

### GetFraisInclus

`func (o *FraisPreviewResponseV2) GetFraisInclus() bool`

GetFraisInclus returns the FraisInclus field if non-nil, zero value otherwise.

### GetFraisInclusOk

`func (o *FraisPreviewResponseV2) GetFraisInclusOk() (*bool, bool)`

GetFraisInclusOk returns a tuple with the FraisInclus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraisInclus

`func (o *FraisPreviewResponseV2) SetFraisInclus(v bool)`

SetFraisInclus sets FraisInclus field to given value.

### HasFraisInclus

`func (o *FraisPreviewResponseV2) HasFraisInclus() bool`

HasFraisInclus returns a boolean if a field has been set.

### GetDevise

`func (o *FraisPreviewResponseV2) GetDevise() string`

GetDevise returns the Devise field if non-nil, zero value otherwise.

### GetDeviseOk

`func (o *FraisPreviewResponseV2) GetDeviseOk() (*string, bool)`

GetDeviseOk returns a tuple with the Devise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevise

`func (o *FraisPreviewResponseV2) SetDevise(v string)`

SetDevise sets Devise field to given value.

### HasDevise

`func (o *FraisPreviewResponseV2) HasDevise() bool`

HasDevise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


