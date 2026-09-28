# ModelEditEnvoieRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VilleDestination** | Pointer to **string** |  | [optional] 
**NomDestinateur** | Pointer to **string** |  | [optional] 
**PrenomDestinateur** | Pointer to **string** |  | [optional] 
**NomExpediteur** | Pointer to **string** |  | [optional] 
**PrenomExpediteur** | Pointer to **string** |  | [optional] 
**PhoneExpediteur** | Pointer to **string** |  | [optional] 
**PhoneDestinateur** | Pointer to **string** |  | [optional] 
**MotifEnvoi** | Pointer to **string** |  | [optional] 
**CodeEnvoi** | Pointer to **string** |  | [optional] 
**MontantEnvoi** | Pointer to **float64** |  | [optional] 
**FraisEnvoi** | Pointer to **float64** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**FraisInclus** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelEditEnvoieRequest

`func NewModelEditEnvoieRequest() *ModelEditEnvoieRequest`

NewModelEditEnvoieRequest instantiates a new ModelEditEnvoieRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEditEnvoieRequestWithDefaults

`func NewModelEditEnvoieRequestWithDefaults() *ModelEditEnvoieRequest`

NewModelEditEnvoieRequestWithDefaults instantiates a new ModelEditEnvoieRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVilleDestination

`func (o *ModelEditEnvoieRequest) GetVilleDestination() string`

GetVilleDestination returns the VilleDestination field if non-nil, zero value otherwise.

### GetVilleDestinationOk

`func (o *ModelEditEnvoieRequest) GetVilleDestinationOk() (*string, bool)`

GetVilleDestinationOk returns a tuple with the VilleDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDestination

`func (o *ModelEditEnvoieRequest) SetVilleDestination(v string)`

SetVilleDestination sets VilleDestination field to given value.

### HasVilleDestination

`func (o *ModelEditEnvoieRequest) HasVilleDestination() bool`

HasVilleDestination returns a boolean if a field has been set.

### GetNomDestinateur

`func (o *ModelEditEnvoieRequest) GetNomDestinateur() string`

GetNomDestinateur returns the NomDestinateur field if non-nil, zero value otherwise.

### GetNomDestinateurOk

`func (o *ModelEditEnvoieRequest) GetNomDestinateurOk() (*string, bool)`

GetNomDestinateurOk returns a tuple with the NomDestinateur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomDestinateur

`func (o *ModelEditEnvoieRequest) SetNomDestinateur(v string)`

SetNomDestinateur sets NomDestinateur field to given value.

### HasNomDestinateur

`func (o *ModelEditEnvoieRequest) HasNomDestinateur() bool`

HasNomDestinateur returns a boolean if a field has been set.

### GetPrenomDestinateur

`func (o *ModelEditEnvoieRequest) GetPrenomDestinateur() string`

GetPrenomDestinateur returns the PrenomDestinateur field if non-nil, zero value otherwise.

### GetPrenomDestinateurOk

`func (o *ModelEditEnvoieRequest) GetPrenomDestinateurOk() (*string, bool)`

GetPrenomDestinateurOk returns a tuple with the PrenomDestinateur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomDestinateur

`func (o *ModelEditEnvoieRequest) SetPrenomDestinateur(v string)`

SetPrenomDestinateur sets PrenomDestinateur field to given value.

### HasPrenomDestinateur

`func (o *ModelEditEnvoieRequest) HasPrenomDestinateur() bool`

HasPrenomDestinateur returns a boolean if a field has been set.

### GetNomExpediteur

`func (o *ModelEditEnvoieRequest) GetNomExpediteur() string`

GetNomExpediteur returns the NomExpediteur field if non-nil, zero value otherwise.

### GetNomExpediteurOk

`func (o *ModelEditEnvoieRequest) GetNomExpediteurOk() (*string, bool)`

GetNomExpediteurOk returns a tuple with the NomExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomExpediteur

`func (o *ModelEditEnvoieRequest) SetNomExpediteur(v string)`

SetNomExpediteur sets NomExpediteur field to given value.

### HasNomExpediteur

`func (o *ModelEditEnvoieRequest) HasNomExpediteur() bool`

HasNomExpediteur returns a boolean if a field has been set.

### GetPrenomExpediteur

`func (o *ModelEditEnvoieRequest) GetPrenomExpediteur() string`

GetPrenomExpediteur returns the PrenomExpediteur field if non-nil, zero value otherwise.

### GetPrenomExpediteurOk

`func (o *ModelEditEnvoieRequest) GetPrenomExpediteurOk() (*string, bool)`

GetPrenomExpediteurOk returns a tuple with the PrenomExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomExpediteur

`func (o *ModelEditEnvoieRequest) SetPrenomExpediteur(v string)`

SetPrenomExpediteur sets PrenomExpediteur field to given value.

### HasPrenomExpediteur

`func (o *ModelEditEnvoieRequest) HasPrenomExpediteur() bool`

HasPrenomExpediteur returns a boolean if a field has been set.

### GetPhoneExpediteur

`func (o *ModelEditEnvoieRequest) GetPhoneExpediteur() string`

GetPhoneExpediteur returns the PhoneExpediteur field if non-nil, zero value otherwise.

### GetPhoneExpediteurOk

`func (o *ModelEditEnvoieRequest) GetPhoneExpediteurOk() (*string, bool)`

GetPhoneExpediteurOk returns a tuple with the PhoneExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneExpediteur

`func (o *ModelEditEnvoieRequest) SetPhoneExpediteur(v string)`

SetPhoneExpediteur sets PhoneExpediteur field to given value.

### HasPhoneExpediteur

`func (o *ModelEditEnvoieRequest) HasPhoneExpediteur() bool`

HasPhoneExpediteur returns a boolean if a field has been set.

### GetPhoneDestinateur

`func (o *ModelEditEnvoieRequest) GetPhoneDestinateur() string`

GetPhoneDestinateur returns the PhoneDestinateur field if non-nil, zero value otherwise.

### GetPhoneDestinateurOk

`func (o *ModelEditEnvoieRequest) GetPhoneDestinateurOk() (*string, bool)`

GetPhoneDestinateurOk returns a tuple with the PhoneDestinateur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneDestinateur

`func (o *ModelEditEnvoieRequest) SetPhoneDestinateur(v string)`

SetPhoneDestinateur sets PhoneDestinateur field to given value.

### HasPhoneDestinateur

`func (o *ModelEditEnvoieRequest) HasPhoneDestinateur() bool`

HasPhoneDestinateur returns a boolean if a field has been set.

### GetMotifEnvoi

`func (o *ModelEditEnvoieRequest) GetMotifEnvoi() string`

GetMotifEnvoi returns the MotifEnvoi field if non-nil, zero value otherwise.

### GetMotifEnvoiOk

`func (o *ModelEditEnvoieRequest) GetMotifEnvoiOk() (*string, bool)`

GetMotifEnvoiOk returns a tuple with the MotifEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifEnvoi

`func (o *ModelEditEnvoieRequest) SetMotifEnvoi(v string)`

SetMotifEnvoi sets MotifEnvoi field to given value.

### HasMotifEnvoi

`func (o *ModelEditEnvoieRequest) HasMotifEnvoi() bool`

HasMotifEnvoi returns a boolean if a field has been set.

### GetCodeEnvoi

`func (o *ModelEditEnvoieRequest) GetCodeEnvoi() string`

GetCodeEnvoi returns the CodeEnvoi field if non-nil, zero value otherwise.

### GetCodeEnvoiOk

`func (o *ModelEditEnvoieRequest) GetCodeEnvoiOk() (*string, bool)`

GetCodeEnvoiOk returns a tuple with the CodeEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEnvoi

`func (o *ModelEditEnvoieRequest) SetCodeEnvoi(v string)`

SetCodeEnvoi sets CodeEnvoi field to given value.

### HasCodeEnvoi

`func (o *ModelEditEnvoieRequest) HasCodeEnvoi() bool`

HasCodeEnvoi returns a boolean if a field has been set.

### GetMontantEnvoi

`func (o *ModelEditEnvoieRequest) GetMontantEnvoi() float64`

GetMontantEnvoi returns the MontantEnvoi field if non-nil, zero value otherwise.

### GetMontantEnvoiOk

`func (o *ModelEditEnvoieRequest) GetMontantEnvoiOk() (*float64, bool)`

GetMontantEnvoiOk returns a tuple with the MontantEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantEnvoi

`func (o *ModelEditEnvoieRequest) SetMontantEnvoi(v float64)`

SetMontantEnvoi sets MontantEnvoi field to given value.

### HasMontantEnvoi

`func (o *ModelEditEnvoieRequest) HasMontantEnvoi() bool`

HasMontantEnvoi returns a boolean if a field has been set.

### GetFraisEnvoi

`func (o *ModelEditEnvoieRequest) GetFraisEnvoi() float64`

GetFraisEnvoi returns the FraisEnvoi field if non-nil, zero value otherwise.

### GetFraisEnvoiOk

`func (o *ModelEditEnvoieRequest) GetFraisEnvoiOk() (*float64, bool)`

GetFraisEnvoiOk returns a tuple with the FraisEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraisEnvoi

`func (o *ModelEditEnvoieRequest) SetFraisEnvoi(v float64)`

SetFraisEnvoi sets FraisEnvoi field to given value.

### HasFraisEnvoi

`func (o *ModelEditEnvoieRequest) HasFraisEnvoi() bool`

HasFraisEnvoi returns a boolean if a field has been set.

### GetTotal

`func (o *ModelEditEnvoieRequest) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ModelEditEnvoieRequest) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ModelEditEnvoieRequest) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ModelEditEnvoieRequest) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetFraisInclus

`func (o *ModelEditEnvoieRequest) GetFraisInclus() bool`

GetFraisInclus returns the FraisInclus field if non-nil, zero value otherwise.

### GetFraisInclusOk

`func (o *ModelEditEnvoieRequest) GetFraisInclusOk() (*bool, bool)`

GetFraisInclusOk returns a tuple with the FraisInclus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraisInclus

`func (o *ModelEditEnvoieRequest) SetFraisInclus(v bool)`

SetFraisInclus sets FraisInclus field to given value.

### HasFraisInclus

`func (o *ModelEditEnvoieRequest) HasFraisInclus() bool`

HasFraisInclus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


