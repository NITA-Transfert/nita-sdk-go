# UpdateEnvoiDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeEnvoi** | **string** |  | 
**RequestId** | **string** |  | 
**IndicatifExpediteur** | **string** |  | 
**NumeroExpediteur** | **string** |  | 
**NomDestinataire** | **string** |  | 
**PrenomDestinataire** | **string** |  | 
**IndicatifDestinataire** | **string** |  | 
**NumeroDestinataire** | **string** |  | 
**VilleDestination** | **string** |  | 

## Methods

### NewUpdateEnvoiDto

`func NewUpdateEnvoiDto(codeEnvoi string, requestId string, indicatifExpediteur string, numeroExpediteur string, nomDestinataire string, prenomDestinataire string, indicatifDestinataire string, numeroDestinataire string, villeDestination string, ) *UpdateEnvoiDto`

NewUpdateEnvoiDto instantiates a new UpdateEnvoiDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnvoiDtoWithDefaults

`func NewUpdateEnvoiDtoWithDefaults() *UpdateEnvoiDto`

NewUpdateEnvoiDtoWithDefaults instantiates a new UpdateEnvoiDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeEnvoi

`func (o *UpdateEnvoiDto) GetCodeEnvoi() string`

GetCodeEnvoi returns the CodeEnvoi field if non-nil, zero value otherwise.

### GetCodeEnvoiOk

`func (o *UpdateEnvoiDto) GetCodeEnvoiOk() (*string, bool)`

GetCodeEnvoiOk returns a tuple with the CodeEnvoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEnvoi

`func (o *UpdateEnvoiDto) SetCodeEnvoi(v string)`

SetCodeEnvoi sets CodeEnvoi field to given value.


### GetRequestId

`func (o *UpdateEnvoiDto) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateEnvoiDto) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateEnvoiDto) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetIndicatifExpediteur

`func (o *UpdateEnvoiDto) GetIndicatifExpediteur() string`

GetIndicatifExpediteur returns the IndicatifExpediteur field if non-nil, zero value otherwise.

### GetIndicatifExpediteurOk

`func (o *UpdateEnvoiDto) GetIndicatifExpediteurOk() (*string, bool)`

GetIndicatifExpediteurOk returns a tuple with the IndicatifExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatifExpediteur

`func (o *UpdateEnvoiDto) SetIndicatifExpediteur(v string)`

SetIndicatifExpediteur sets IndicatifExpediteur field to given value.


### GetNumeroExpediteur

`func (o *UpdateEnvoiDto) GetNumeroExpediteur() string`

GetNumeroExpediteur returns the NumeroExpediteur field if non-nil, zero value otherwise.

### GetNumeroExpediteurOk

`func (o *UpdateEnvoiDto) GetNumeroExpediteurOk() (*string, bool)`

GetNumeroExpediteurOk returns a tuple with the NumeroExpediteur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroExpediteur

`func (o *UpdateEnvoiDto) SetNumeroExpediteur(v string)`

SetNumeroExpediteur sets NumeroExpediteur field to given value.


### GetNomDestinataire

`func (o *UpdateEnvoiDto) GetNomDestinataire() string`

GetNomDestinataire returns the NomDestinataire field if non-nil, zero value otherwise.

### GetNomDestinataireOk

`func (o *UpdateEnvoiDto) GetNomDestinataireOk() (*string, bool)`

GetNomDestinataireOk returns a tuple with the NomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomDestinataire

`func (o *UpdateEnvoiDto) SetNomDestinataire(v string)`

SetNomDestinataire sets NomDestinataire field to given value.


### GetPrenomDestinataire

`func (o *UpdateEnvoiDto) GetPrenomDestinataire() string`

GetPrenomDestinataire returns the PrenomDestinataire field if non-nil, zero value otherwise.

### GetPrenomDestinataireOk

`func (o *UpdateEnvoiDto) GetPrenomDestinataireOk() (*string, bool)`

GetPrenomDestinataireOk returns a tuple with the PrenomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomDestinataire

`func (o *UpdateEnvoiDto) SetPrenomDestinataire(v string)`

SetPrenomDestinataire sets PrenomDestinataire field to given value.


### GetIndicatifDestinataire

`func (o *UpdateEnvoiDto) GetIndicatifDestinataire() string`

GetIndicatifDestinataire returns the IndicatifDestinataire field if non-nil, zero value otherwise.

### GetIndicatifDestinataireOk

`func (o *UpdateEnvoiDto) GetIndicatifDestinataireOk() (*string, bool)`

GetIndicatifDestinataireOk returns a tuple with the IndicatifDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatifDestinataire

`func (o *UpdateEnvoiDto) SetIndicatifDestinataire(v string)`

SetIndicatifDestinataire sets IndicatifDestinataire field to given value.


### GetNumeroDestinataire

`func (o *UpdateEnvoiDto) GetNumeroDestinataire() string`

GetNumeroDestinataire returns the NumeroDestinataire field if non-nil, zero value otherwise.

### GetNumeroDestinataireOk

`func (o *UpdateEnvoiDto) GetNumeroDestinataireOk() (*string, bool)`

GetNumeroDestinataireOk returns a tuple with the NumeroDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroDestinataire

`func (o *UpdateEnvoiDto) SetNumeroDestinataire(v string)`

SetNumeroDestinataire sets NumeroDestinataire field to given value.


### GetVilleDestination

`func (o *UpdateEnvoiDto) GetVilleDestination() string`

GetVilleDestination returns the VilleDestination field if non-nil, zero value otherwise.

### GetVilleDestinationOk

`func (o *UpdateEnvoiDto) GetVilleDestinationOk() (*string, bool)`

GetVilleDestinationOk returns a tuple with the VilleDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDestination

`func (o *UpdateEnvoiDto) SetVilleDestination(v string)`

SetVilleDestination sets VilleDestination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


