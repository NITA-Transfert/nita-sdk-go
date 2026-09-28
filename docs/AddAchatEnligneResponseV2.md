# AddAchatEnligneResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicId** | Pointer to **string** |  | [optional] 
**PhoneClient** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**AdresseIP** | Pointer to **string** |  | [optional] 
**CodeAchat** | Pointer to **string** |  | [optional] 
**DescriptionAchat** | Pointer to **string** |  | [optional] 
**Montant** | Pointer to **float64** |  | [optional] 
**StatusTransaction** | Pointer to **string** |  | [optional] 
**TypeTransaction** | Pointer to **string** |  | [optional] 
**ComptePartenaire** | Pointer to **string** |  | [optional] 
**NomPartenaire** | Pointer to **string** |  | [optional] 
**PrenomPartenaire** | Pointer to **string** |  | [optional] 
**DateTransaction** | Pointer to **time.Time** |  | [optional] 
**UrlPaiement** | Pointer to **string** | Lien de règlement à transmettre à l&#39;acheteur, dérivé du codeAchat. Nul tant que la page de paiement publique n&#39;est pas en service : un champ absent vaut mieux qu&#39;un lien qui ne s&#39;ouvre pas. Ne pas le reconstruire à la main. | [optional] 

## Methods

### NewAddAchatEnligneResponseV2

`func NewAddAchatEnligneResponseV2() *AddAchatEnligneResponseV2`

NewAddAchatEnligneResponseV2 instantiates a new AddAchatEnligneResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAchatEnligneResponseV2WithDefaults

`func NewAddAchatEnligneResponseV2WithDefaults() *AddAchatEnligneResponseV2`

NewAddAchatEnligneResponseV2WithDefaults instantiates a new AddAchatEnligneResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicId

`func (o *AddAchatEnligneResponseV2) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *AddAchatEnligneResponseV2) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *AddAchatEnligneResponseV2) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *AddAchatEnligneResponseV2) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetPhoneClient

`func (o *AddAchatEnligneResponseV2) GetPhoneClient() string`

GetPhoneClient returns the PhoneClient field if non-nil, zero value otherwise.

### GetPhoneClientOk

`func (o *AddAchatEnligneResponseV2) GetPhoneClientOk() (*string, bool)`

GetPhoneClientOk returns a tuple with the PhoneClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneClient

`func (o *AddAchatEnligneResponseV2) SetPhoneClient(v string)`

SetPhoneClient sets PhoneClient field to given value.

### HasPhoneClient

`func (o *AddAchatEnligneResponseV2) HasPhoneClient() bool`

HasPhoneClient returns a boolean if a field has been set.

### GetRequestId

`func (o *AddAchatEnligneResponseV2) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AddAchatEnligneResponseV2) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AddAchatEnligneResponseV2) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AddAchatEnligneResponseV2) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetAdresseIP

`func (o *AddAchatEnligneResponseV2) GetAdresseIP() string`

GetAdresseIP returns the AdresseIP field if non-nil, zero value otherwise.

### GetAdresseIPOk

`func (o *AddAchatEnligneResponseV2) GetAdresseIPOk() (*string, bool)`

GetAdresseIPOk returns a tuple with the AdresseIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseIP

`func (o *AddAchatEnligneResponseV2) SetAdresseIP(v string)`

SetAdresseIP sets AdresseIP field to given value.

### HasAdresseIP

`func (o *AddAchatEnligneResponseV2) HasAdresseIP() bool`

HasAdresseIP returns a boolean if a field has been set.

### GetCodeAchat

`func (o *AddAchatEnligneResponseV2) GetCodeAchat() string`

GetCodeAchat returns the CodeAchat field if non-nil, zero value otherwise.

### GetCodeAchatOk

`func (o *AddAchatEnligneResponseV2) GetCodeAchatOk() (*string, bool)`

GetCodeAchatOk returns a tuple with the CodeAchat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeAchat

`func (o *AddAchatEnligneResponseV2) SetCodeAchat(v string)`

SetCodeAchat sets CodeAchat field to given value.

### HasCodeAchat

`func (o *AddAchatEnligneResponseV2) HasCodeAchat() bool`

HasCodeAchat returns a boolean if a field has been set.

### GetDescriptionAchat

`func (o *AddAchatEnligneResponseV2) GetDescriptionAchat() string`

GetDescriptionAchat returns the DescriptionAchat field if non-nil, zero value otherwise.

### GetDescriptionAchatOk

`func (o *AddAchatEnligneResponseV2) GetDescriptionAchatOk() (*string, bool)`

GetDescriptionAchatOk returns a tuple with the DescriptionAchat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionAchat

`func (o *AddAchatEnligneResponseV2) SetDescriptionAchat(v string)`

SetDescriptionAchat sets DescriptionAchat field to given value.

### HasDescriptionAchat

`func (o *AddAchatEnligneResponseV2) HasDescriptionAchat() bool`

HasDescriptionAchat returns a boolean if a field has been set.

### GetMontant

`func (o *AddAchatEnligneResponseV2) GetMontant() float64`

GetMontant returns the Montant field if non-nil, zero value otherwise.

### GetMontantOk

`func (o *AddAchatEnligneResponseV2) GetMontantOk() (*float64, bool)`

GetMontantOk returns a tuple with the Montant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontant

`func (o *AddAchatEnligneResponseV2) SetMontant(v float64)`

SetMontant sets Montant field to given value.

### HasMontant

`func (o *AddAchatEnligneResponseV2) HasMontant() bool`

HasMontant returns a boolean if a field has been set.

### GetStatusTransaction

`func (o *AddAchatEnligneResponseV2) GetStatusTransaction() string`

GetStatusTransaction returns the StatusTransaction field if non-nil, zero value otherwise.

### GetStatusTransactionOk

`func (o *AddAchatEnligneResponseV2) GetStatusTransactionOk() (*string, bool)`

GetStatusTransactionOk returns a tuple with the StatusTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTransaction

`func (o *AddAchatEnligneResponseV2) SetStatusTransaction(v string)`

SetStatusTransaction sets StatusTransaction field to given value.

### HasStatusTransaction

`func (o *AddAchatEnligneResponseV2) HasStatusTransaction() bool`

HasStatusTransaction returns a boolean if a field has been set.

### GetTypeTransaction

`func (o *AddAchatEnligneResponseV2) GetTypeTransaction() string`

GetTypeTransaction returns the TypeTransaction field if non-nil, zero value otherwise.

### GetTypeTransactionOk

`func (o *AddAchatEnligneResponseV2) GetTypeTransactionOk() (*string, bool)`

GetTypeTransactionOk returns a tuple with the TypeTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeTransaction

`func (o *AddAchatEnligneResponseV2) SetTypeTransaction(v string)`

SetTypeTransaction sets TypeTransaction field to given value.

### HasTypeTransaction

`func (o *AddAchatEnligneResponseV2) HasTypeTransaction() bool`

HasTypeTransaction returns a boolean if a field has been set.

### GetComptePartenaire

`func (o *AddAchatEnligneResponseV2) GetComptePartenaire() string`

GetComptePartenaire returns the ComptePartenaire field if non-nil, zero value otherwise.

### GetComptePartenaireOk

`func (o *AddAchatEnligneResponseV2) GetComptePartenaireOk() (*string, bool)`

GetComptePartenaireOk returns a tuple with the ComptePartenaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComptePartenaire

`func (o *AddAchatEnligneResponseV2) SetComptePartenaire(v string)`

SetComptePartenaire sets ComptePartenaire field to given value.

### HasComptePartenaire

`func (o *AddAchatEnligneResponseV2) HasComptePartenaire() bool`

HasComptePartenaire returns a boolean if a field has been set.

### GetNomPartenaire

`func (o *AddAchatEnligneResponseV2) GetNomPartenaire() string`

GetNomPartenaire returns the NomPartenaire field if non-nil, zero value otherwise.

### GetNomPartenaireOk

`func (o *AddAchatEnligneResponseV2) GetNomPartenaireOk() (*string, bool)`

GetNomPartenaireOk returns a tuple with the NomPartenaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomPartenaire

`func (o *AddAchatEnligneResponseV2) SetNomPartenaire(v string)`

SetNomPartenaire sets NomPartenaire field to given value.

### HasNomPartenaire

`func (o *AddAchatEnligneResponseV2) HasNomPartenaire() bool`

HasNomPartenaire returns a boolean if a field has been set.

### GetPrenomPartenaire

`func (o *AddAchatEnligneResponseV2) GetPrenomPartenaire() string`

GetPrenomPartenaire returns the PrenomPartenaire field if non-nil, zero value otherwise.

### GetPrenomPartenaireOk

`func (o *AddAchatEnligneResponseV2) GetPrenomPartenaireOk() (*string, bool)`

GetPrenomPartenaireOk returns a tuple with the PrenomPartenaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomPartenaire

`func (o *AddAchatEnligneResponseV2) SetPrenomPartenaire(v string)`

SetPrenomPartenaire sets PrenomPartenaire field to given value.

### HasPrenomPartenaire

`func (o *AddAchatEnligneResponseV2) HasPrenomPartenaire() bool`

HasPrenomPartenaire returns a boolean if a field has been set.

### GetDateTransaction

`func (o *AddAchatEnligneResponseV2) GetDateTransaction() time.Time`

GetDateTransaction returns the DateTransaction field if non-nil, zero value otherwise.

### GetDateTransactionOk

`func (o *AddAchatEnligneResponseV2) GetDateTransactionOk() (*time.Time, bool)`

GetDateTransactionOk returns a tuple with the DateTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTransaction

`func (o *AddAchatEnligneResponseV2) SetDateTransaction(v time.Time)`

SetDateTransaction sets DateTransaction field to given value.

### HasDateTransaction

`func (o *AddAchatEnligneResponseV2) HasDateTransaction() bool`

HasDateTransaction returns a boolean if a field has been set.

### GetUrlPaiement

`func (o *AddAchatEnligneResponseV2) GetUrlPaiement() string`

GetUrlPaiement returns the UrlPaiement field if non-nil, zero value otherwise.

### GetUrlPaiementOk

`func (o *AddAchatEnligneResponseV2) GetUrlPaiementOk() (*string, bool)`

GetUrlPaiementOk returns a tuple with the UrlPaiement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPaiement

`func (o *AddAchatEnligneResponseV2) SetUrlPaiement(v string)`

SetUrlPaiement sets UrlPaiement field to given value.

### HasUrlPaiement

`func (o *AddAchatEnligneResponseV2) HasUrlPaiement() bool`

HasUrlPaiement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


