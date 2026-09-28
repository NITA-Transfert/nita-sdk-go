# PartenaireToCashDtoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nom** | **string** |  | 
**Prenom** | **string** |  | 
**Indicatif** | **string** |  | 
**Numero** | **string** |  | 
**DateNaissance** | Pointer to **time.Time** |  | [optional] 
**LieuNaissance** | Pointer to **string** |  | [optional] 
**SituationMatrimoniale** | Pointer to **string** |  | [optional] 
**Genre** | Pointer to **string** |  | [optional] 
**Profession** | Pointer to **string** |  | [optional] 
**Adresse** | Pointer to **string** |  | [optional] 
**VilleResidence** | Pointer to **string** |  | [optional] 
**PaysResidence** | Pointer to **string** |  | [optional] 
**TypePiece** | Pointer to **string** |  | [optional] 
**NumeroPiece** | Pointer to **string** |  | [optional] 
**DateDelivrance** | Pointer to **time.Time** |  | [optional] 
**LieuDelivrance** | Pointer to **string** |  | [optional] 
**Autorite** | Pointer to **string** |  | [optional] 
**PhotoPieceIdentiteRecto** | Pointer to **string** |  | [optional] 
**PhotoPieceIdentiteVerso** | Pointer to **string** |  | [optional] 
**PhotoIdentite** | Pointer to **string** |  | [optional] 
**RequestId** | **string** |  | 
**FraisInclus** | **bool** |  | 
**Montant** | **float64** |  | 
**VilleDestination** | **string** |  | 
**NomDestinataire** | **string** |  | 
**PrenomDestinataire** | **string** |  | 
**IndicatifDestinataire** | **string** |  | 
**NumeroDestinataire** | **string** |  | 
**MotifTransaction** | Pointer to **string** |  | [optional] 
**AdresseIp** | Pointer to **string** |  | [optional] 
**UrlCallback** | Pointer to **string** |  | [optional] 

## Methods

### NewPartenaireToCashDtoV2

`func NewPartenaireToCashDtoV2(nom string, prenom string, indicatif string, numero string, requestId string, fraisInclus bool, montant float64, villeDestination string, nomDestinataire string, prenomDestinataire string, indicatifDestinataire string, numeroDestinataire string, ) *PartenaireToCashDtoV2`

NewPartenaireToCashDtoV2 instantiates a new PartenaireToCashDtoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartenaireToCashDtoV2WithDefaults

`func NewPartenaireToCashDtoV2WithDefaults() *PartenaireToCashDtoV2`

NewPartenaireToCashDtoV2WithDefaults instantiates a new PartenaireToCashDtoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *PartenaireToCashDtoV2) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *PartenaireToCashDtoV2) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *PartenaireToCashDtoV2) SetNom(v string)`

SetNom sets Nom field to given value.


### GetPrenom

`func (o *PartenaireToCashDtoV2) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *PartenaireToCashDtoV2) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *PartenaireToCashDtoV2) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.


### GetIndicatif

`func (o *PartenaireToCashDtoV2) GetIndicatif() string`

GetIndicatif returns the Indicatif field if non-nil, zero value otherwise.

### GetIndicatifOk

`func (o *PartenaireToCashDtoV2) GetIndicatifOk() (*string, bool)`

GetIndicatifOk returns a tuple with the Indicatif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatif

`func (o *PartenaireToCashDtoV2) SetIndicatif(v string)`

SetIndicatif sets Indicatif field to given value.


### GetNumero

`func (o *PartenaireToCashDtoV2) GetNumero() string`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *PartenaireToCashDtoV2) GetNumeroOk() (*string, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *PartenaireToCashDtoV2) SetNumero(v string)`

SetNumero sets Numero field to given value.


### GetDateNaissance

`func (o *PartenaireToCashDtoV2) GetDateNaissance() time.Time`

GetDateNaissance returns the DateNaissance field if non-nil, zero value otherwise.

### GetDateNaissanceOk

`func (o *PartenaireToCashDtoV2) GetDateNaissanceOk() (*time.Time, bool)`

GetDateNaissanceOk returns a tuple with the DateNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateNaissance

`func (o *PartenaireToCashDtoV2) SetDateNaissance(v time.Time)`

SetDateNaissance sets DateNaissance field to given value.

### HasDateNaissance

`func (o *PartenaireToCashDtoV2) HasDateNaissance() bool`

HasDateNaissance returns a boolean if a field has been set.

### GetLieuNaissance

`func (o *PartenaireToCashDtoV2) GetLieuNaissance() string`

GetLieuNaissance returns the LieuNaissance field if non-nil, zero value otherwise.

### GetLieuNaissanceOk

`func (o *PartenaireToCashDtoV2) GetLieuNaissanceOk() (*string, bool)`

GetLieuNaissanceOk returns a tuple with the LieuNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLieuNaissance

`func (o *PartenaireToCashDtoV2) SetLieuNaissance(v string)`

SetLieuNaissance sets LieuNaissance field to given value.

### HasLieuNaissance

`func (o *PartenaireToCashDtoV2) HasLieuNaissance() bool`

HasLieuNaissance returns a boolean if a field has been set.

### GetSituationMatrimoniale

`func (o *PartenaireToCashDtoV2) GetSituationMatrimoniale() string`

GetSituationMatrimoniale returns the SituationMatrimoniale field if non-nil, zero value otherwise.

### GetSituationMatrimonialeOk

`func (o *PartenaireToCashDtoV2) GetSituationMatrimonialeOk() (*string, bool)`

GetSituationMatrimonialeOk returns a tuple with the SituationMatrimoniale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSituationMatrimoniale

`func (o *PartenaireToCashDtoV2) SetSituationMatrimoniale(v string)`

SetSituationMatrimoniale sets SituationMatrimoniale field to given value.

### HasSituationMatrimoniale

`func (o *PartenaireToCashDtoV2) HasSituationMatrimoniale() bool`

HasSituationMatrimoniale returns a boolean if a field has been set.

### GetGenre

`func (o *PartenaireToCashDtoV2) GetGenre() string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *PartenaireToCashDtoV2) GetGenreOk() (*string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *PartenaireToCashDtoV2) SetGenre(v string)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *PartenaireToCashDtoV2) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetProfession

`func (o *PartenaireToCashDtoV2) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *PartenaireToCashDtoV2) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *PartenaireToCashDtoV2) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *PartenaireToCashDtoV2) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetAdresse

`func (o *PartenaireToCashDtoV2) GetAdresse() string`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *PartenaireToCashDtoV2) GetAdresseOk() (*string, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *PartenaireToCashDtoV2) SetAdresse(v string)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *PartenaireToCashDtoV2) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetVilleResidence

`func (o *PartenaireToCashDtoV2) GetVilleResidence() string`

GetVilleResidence returns the VilleResidence field if non-nil, zero value otherwise.

### GetVilleResidenceOk

`func (o *PartenaireToCashDtoV2) GetVilleResidenceOk() (*string, bool)`

GetVilleResidenceOk returns a tuple with the VilleResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleResidence

`func (o *PartenaireToCashDtoV2) SetVilleResidence(v string)`

SetVilleResidence sets VilleResidence field to given value.

### HasVilleResidence

`func (o *PartenaireToCashDtoV2) HasVilleResidence() bool`

HasVilleResidence returns a boolean if a field has been set.

### GetPaysResidence

`func (o *PartenaireToCashDtoV2) GetPaysResidence() string`

GetPaysResidence returns the PaysResidence field if non-nil, zero value otherwise.

### GetPaysResidenceOk

`func (o *PartenaireToCashDtoV2) GetPaysResidenceOk() (*string, bool)`

GetPaysResidenceOk returns a tuple with the PaysResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysResidence

`func (o *PartenaireToCashDtoV2) SetPaysResidence(v string)`

SetPaysResidence sets PaysResidence field to given value.

### HasPaysResidence

`func (o *PartenaireToCashDtoV2) HasPaysResidence() bool`

HasPaysResidence returns a boolean if a field has been set.

### GetTypePiece

`func (o *PartenaireToCashDtoV2) GetTypePiece() string`

GetTypePiece returns the TypePiece field if non-nil, zero value otherwise.

### GetTypePieceOk

`func (o *PartenaireToCashDtoV2) GetTypePieceOk() (*string, bool)`

GetTypePieceOk returns a tuple with the TypePiece field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypePiece

`func (o *PartenaireToCashDtoV2) SetTypePiece(v string)`

SetTypePiece sets TypePiece field to given value.

### HasTypePiece

`func (o *PartenaireToCashDtoV2) HasTypePiece() bool`

HasTypePiece returns a boolean if a field has been set.

### GetNumeroPiece

`func (o *PartenaireToCashDtoV2) GetNumeroPiece() string`

GetNumeroPiece returns the NumeroPiece field if non-nil, zero value otherwise.

### GetNumeroPieceOk

`func (o *PartenaireToCashDtoV2) GetNumeroPieceOk() (*string, bool)`

GetNumeroPieceOk returns a tuple with the NumeroPiece field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroPiece

`func (o *PartenaireToCashDtoV2) SetNumeroPiece(v string)`

SetNumeroPiece sets NumeroPiece field to given value.

### HasNumeroPiece

`func (o *PartenaireToCashDtoV2) HasNumeroPiece() bool`

HasNumeroPiece returns a boolean if a field has been set.

### GetDateDelivrance

`func (o *PartenaireToCashDtoV2) GetDateDelivrance() time.Time`

GetDateDelivrance returns the DateDelivrance field if non-nil, zero value otherwise.

### GetDateDelivranceOk

`func (o *PartenaireToCashDtoV2) GetDateDelivranceOk() (*time.Time, bool)`

GetDateDelivranceOk returns a tuple with the DateDelivrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDelivrance

`func (o *PartenaireToCashDtoV2) SetDateDelivrance(v time.Time)`

SetDateDelivrance sets DateDelivrance field to given value.

### HasDateDelivrance

`func (o *PartenaireToCashDtoV2) HasDateDelivrance() bool`

HasDateDelivrance returns a boolean if a field has been set.

### GetLieuDelivrance

`func (o *PartenaireToCashDtoV2) GetLieuDelivrance() string`

GetLieuDelivrance returns the LieuDelivrance field if non-nil, zero value otherwise.

### GetLieuDelivranceOk

`func (o *PartenaireToCashDtoV2) GetLieuDelivranceOk() (*string, bool)`

GetLieuDelivranceOk returns a tuple with the LieuDelivrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLieuDelivrance

`func (o *PartenaireToCashDtoV2) SetLieuDelivrance(v string)`

SetLieuDelivrance sets LieuDelivrance field to given value.

### HasLieuDelivrance

`func (o *PartenaireToCashDtoV2) HasLieuDelivrance() bool`

HasLieuDelivrance returns a boolean if a field has been set.

### GetAutorite

`func (o *PartenaireToCashDtoV2) GetAutorite() string`

GetAutorite returns the Autorite field if non-nil, zero value otherwise.

### GetAutoriteOk

`func (o *PartenaireToCashDtoV2) GetAutoriteOk() (*string, bool)`

GetAutoriteOk returns a tuple with the Autorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorite

`func (o *PartenaireToCashDtoV2) SetAutorite(v string)`

SetAutorite sets Autorite field to given value.

### HasAutorite

`func (o *PartenaireToCashDtoV2) HasAutorite() bool`

HasAutorite returns a boolean if a field has been set.

### GetPhotoPieceIdentiteRecto

`func (o *PartenaireToCashDtoV2) GetPhotoPieceIdentiteRecto() string`

GetPhotoPieceIdentiteRecto returns the PhotoPieceIdentiteRecto field if non-nil, zero value otherwise.

### GetPhotoPieceIdentiteRectoOk

`func (o *PartenaireToCashDtoV2) GetPhotoPieceIdentiteRectoOk() (*string, bool)`

GetPhotoPieceIdentiteRectoOk returns a tuple with the PhotoPieceIdentiteRecto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoPieceIdentiteRecto

`func (o *PartenaireToCashDtoV2) SetPhotoPieceIdentiteRecto(v string)`

SetPhotoPieceIdentiteRecto sets PhotoPieceIdentiteRecto field to given value.

### HasPhotoPieceIdentiteRecto

`func (o *PartenaireToCashDtoV2) HasPhotoPieceIdentiteRecto() bool`

HasPhotoPieceIdentiteRecto returns a boolean if a field has been set.

### GetPhotoPieceIdentiteVerso

`func (o *PartenaireToCashDtoV2) GetPhotoPieceIdentiteVerso() string`

GetPhotoPieceIdentiteVerso returns the PhotoPieceIdentiteVerso field if non-nil, zero value otherwise.

### GetPhotoPieceIdentiteVersoOk

`func (o *PartenaireToCashDtoV2) GetPhotoPieceIdentiteVersoOk() (*string, bool)`

GetPhotoPieceIdentiteVersoOk returns a tuple with the PhotoPieceIdentiteVerso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoPieceIdentiteVerso

`func (o *PartenaireToCashDtoV2) SetPhotoPieceIdentiteVerso(v string)`

SetPhotoPieceIdentiteVerso sets PhotoPieceIdentiteVerso field to given value.

### HasPhotoPieceIdentiteVerso

`func (o *PartenaireToCashDtoV2) HasPhotoPieceIdentiteVerso() bool`

HasPhotoPieceIdentiteVerso returns a boolean if a field has been set.

### GetPhotoIdentite

`func (o *PartenaireToCashDtoV2) GetPhotoIdentite() string`

GetPhotoIdentite returns the PhotoIdentite field if non-nil, zero value otherwise.

### GetPhotoIdentiteOk

`func (o *PartenaireToCashDtoV2) GetPhotoIdentiteOk() (*string, bool)`

GetPhotoIdentiteOk returns a tuple with the PhotoIdentite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoIdentite

`func (o *PartenaireToCashDtoV2) SetPhotoIdentite(v string)`

SetPhotoIdentite sets PhotoIdentite field to given value.

### HasPhotoIdentite

`func (o *PartenaireToCashDtoV2) HasPhotoIdentite() bool`

HasPhotoIdentite returns a boolean if a field has been set.

### GetRequestId

`func (o *PartenaireToCashDtoV2) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PartenaireToCashDtoV2) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PartenaireToCashDtoV2) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetFraisInclus

`func (o *PartenaireToCashDtoV2) GetFraisInclus() bool`

GetFraisInclus returns the FraisInclus field if non-nil, zero value otherwise.

### GetFraisInclusOk

`func (o *PartenaireToCashDtoV2) GetFraisInclusOk() (*bool, bool)`

GetFraisInclusOk returns a tuple with the FraisInclus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraisInclus

`func (o *PartenaireToCashDtoV2) SetFraisInclus(v bool)`

SetFraisInclus sets FraisInclus field to given value.


### GetMontant

`func (o *PartenaireToCashDtoV2) GetMontant() float64`

GetMontant returns the Montant field if non-nil, zero value otherwise.

### GetMontantOk

`func (o *PartenaireToCashDtoV2) GetMontantOk() (*float64, bool)`

GetMontantOk returns a tuple with the Montant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontant

`func (o *PartenaireToCashDtoV2) SetMontant(v float64)`

SetMontant sets Montant field to given value.


### GetVilleDestination

`func (o *PartenaireToCashDtoV2) GetVilleDestination() string`

GetVilleDestination returns the VilleDestination field if non-nil, zero value otherwise.

### GetVilleDestinationOk

`func (o *PartenaireToCashDtoV2) GetVilleDestinationOk() (*string, bool)`

GetVilleDestinationOk returns a tuple with the VilleDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDestination

`func (o *PartenaireToCashDtoV2) SetVilleDestination(v string)`

SetVilleDestination sets VilleDestination field to given value.


### GetNomDestinataire

`func (o *PartenaireToCashDtoV2) GetNomDestinataire() string`

GetNomDestinataire returns the NomDestinataire field if non-nil, zero value otherwise.

### GetNomDestinataireOk

`func (o *PartenaireToCashDtoV2) GetNomDestinataireOk() (*string, bool)`

GetNomDestinataireOk returns a tuple with the NomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomDestinataire

`func (o *PartenaireToCashDtoV2) SetNomDestinataire(v string)`

SetNomDestinataire sets NomDestinataire field to given value.


### GetPrenomDestinataire

`func (o *PartenaireToCashDtoV2) GetPrenomDestinataire() string`

GetPrenomDestinataire returns the PrenomDestinataire field if non-nil, zero value otherwise.

### GetPrenomDestinataireOk

`func (o *PartenaireToCashDtoV2) GetPrenomDestinataireOk() (*string, bool)`

GetPrenomDestinataireOk returns a tuple with the PrenomDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomDestinataire

`func (o *PartenaireToCashDtoV2) SetPrenomDestinataire(v string)`

SetPrenomDestinataire sets PrenomDestinataire field to given value.


### GetIndicatifDestinataire

`func (o *PartenaireToCashDtoV2) GetIndicatifDestinataire() string`

GetIndicatifDestinataire returns the IndicatifDestinataire field if non-nil, zero value otherwise.

### GetIndicatifDestinataireOk

`func (o *PartenaireToCashDtoV2) GetIndicatifDestinataireOk() (*string, bool)`

GetIndicatifDestinataireOk returns a tuple with the IndicatifDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatifDestinataire

`func (o *PartenaireToCashDtoV2) SetIndicatifDestinataire(v string)`

SetIndicatifDestinataire sets IndicatifDestinataire field to given value.


### GetNumeroDestinataire

`func (o *PartenaireToCashDtoV2) GetNumeroDestinataire() string`

GetNumeroDestinataire returns the NumeroDestinataire field if non-nil, zero value otherwise.

### GetNumeroDestinataireOk

`func (o *PartenaireToCashDtoV2) GetNumeroDestinataireOk() (*string, bool)`

GetNumeroDestinataireOk returns a tuple with the NumeroDestinataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroDestinataire

`func (o *PartenaireToCashDtoV2) SetNumeroDestinataire(v string)`

SetNumeroDestinataire sets NumeroDestinataire field to given value.


### GetMotifTransaction

`func (o *PartenaireToCashDtoV2) GetMotifTransaction() string`

GetMotifTransaction returns the MotifTransaction field if non-nil, zero value otherwise.

### GetMotifTransactionOk

`func (o *PartenaireToCashDtoV2) GetMotifTransactionOk() (*string, bool)`

GetMotifTransactionOk returns a tuple with the MotifTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifTransaction

`func (o *PartenaireToCashDtoV2) SetMotifTransaction(v string)`

SetMotifTransaction sets MotifTransaction field to given value.

### HasMotifTransaction

`func (o *PartenaireToCashDtoV2) HasMotifTransaction() bool`

HasMotifTransaction returns a boolean if a field has been set.

### GetAdresseIp

`func (o *PartenaireToCashDtoV2) GetAdresseIp() string`

GetAdresseIp returns the AdresseIp field if non-nil, zero value otherwise.

### GetAdresseIpOk

`func (o *PartenaireToCashDtoV2) GetAdresseIpOk() (*string, bool)`

GetAdresseIpOk returns a tuple with the AdresseIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseIp

`func (o *PartenaireToCashDtoV2) SetAdresseIp(v string)`

SetAdresseIp sets AdresseIp field to given value.

### HasAdresseIp

`func (o *PartenaireToCashDtoV2) HasAdresseIp() bool`

HasAdresseIp returns a boolean if a field has been set.

### GetUrlCallback

`func (o *PartenaireToCashDtoV2) GetUrlCallback() string`

GetUrlCallback returns the UrlCallback field if non-nil, zero value otherwise.

### GetUrlCallbackOk

`func (o *PartenaireToCashDtoV2) GetUrlCallbackOk() (*string, bool)`

GetUrlCallbackOk returns a tuple with the UrlCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCallback

`func (o *PartenaireToCashDtoV2) SetUrlCallback(v string)`

SetUrlCallback sets UrlCallback field to given value.

### HasUrlCallback

`func (o *PartenaireToCashDtoV2) HasUrlCallback() bool`

HasUrlCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


