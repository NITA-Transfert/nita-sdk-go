# AchatEnLigneModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptionAchat** | **[]string** |  | 
**MontantTransaction** | **float64** |  | 
**PhoneClient** | **string** |  | 
**MotifTransaction** | **string** |  | 
**LongTransaction** | Pointer to **string** |  | [optional] 
**LatTransaction** | Pointer to **string** |  | [optional] 
**RequestId** | **string** |  | 
**UrlCallback** | Pointer to **string** |  | [optional] 
**UrlRetour** | Pointer to **string** |  | [optional] 
**AdresseIp** | **string** |  | 

## Methods

### NewAchatEnLigneModel

`func NewAchatEnLigneModel(descriptionAchat []string, montantTransaction float64, phoneClient string, motifTransaction string, requestId string, adresseIp string, ) *AchatEnLigneModel`

NewAchatEnLigneModel instantiates a new AchatEnLigneModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchatEnLigneModelWithDefaults

`func NewAchatEnLigneModelWithDefaults() *AchatEnLigneModel`

NewAchatEnLigneModelWithDefaults instantiates a new AchatEnLigneModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptionAchat

`func (o *AchatEnLigneModel) GetDescriptionAchat() []string`

GetDescriptionAchat returns the DescriptionAchat field if non-nil, zero value otherwise.

### GetDescriptionAchatOk

`func (o *AchatEnLigneModel) GetDescriptionAchatOk() (*[]string, bool)`

GetDescriptionAchatOk returns a tuple with the DescriptionAchat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionAchat

`func (o *AchatEnLigneModel) SetDescriptionAchat(v []string)`

SetDescriptionAchat sets DescriptionAchat field to given value.


### GetMontantTransaction

`func (o *AchatEnLigneModel) GetMontantTransaction() float64`

GetMontantTransaction returns the MontantTransaction field if non-nil, zero value otherwise.

### GetMontantTransactionOk

`func (o *AchatEnLigneModel) GetMontantTransactionOk() (*float64, bool)`

GetMontantTransactionOk returns a tuple with the MontantTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontantTransaction

`func (o *AchatEnLigneModel) SetMontantTransaction(v float64)`

SetMontantTransaction sets MontantTransaction field to given value.


### GetPhoneClient

`func (o *AchatEnLigneModel) GetPhoneClient() string`

GetPhoneClient returns the PhoneClient field if non-nil, zero value otherwise.

### GetPhoneClientOk

`func (o *AchatEnLigneModel) GetPhoneClientOk() (*string, bool)`

GetPhoneClientOk returns a tuple with the PhoneClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneClient

`func (o *AchatEnLigneModel) SetPhoneClient(v string)`

SetPhoneClient sets PhoneClient field to given value.


### GetMotifTransaction

`func (o *AchatEnLigneModel) GetMotifTransaction() string`

GetMotifTransaction returns the MotifTransaction field if non-nil, zero value otherwise.

### GetMotifTransactionOk

`func (o *AchatEnLigneModel) GetMotifTransactionOk() (*string, bool)`

GetMotifTransactionOk returns a tuple with the MotifTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotifTransaction

`func (o *AchatEnLigneModel) SetMotifTransaction(v string)`

SetMotifTransaction sets MotifTransaction field to given value.


### GetLongTransaction

`func (o *AchatEnLigneModel) GetLongTransaction() string`

GetLongTransaction returns the LongTransaction field if non-nil, zero value otherwise.

### GetLongTransactionOk

`func (o *AchatEnLigneModel) GetLongTransactionOk() (*string, bool)`

GetLongTransactionOk returns a tuple with the LongTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTransaction

`func (o *AchatEnLigneModel) SetLongTransaction(v string)`

SetLongTransaction sets LongTransaction field to given value.

### HasLongTransaction

`func (o *AchatEnLigneModel) HasLongTransaction() bool`

HasLongTransaction returns a boolean if a field has been set.

### GetLatTransaction

`func (o *AchatEnLigneModel) GetLatTransaction() string`

GetLatTransaction returns the LatTransaction field if non-nil, zero value otherwise.

### GetLatTransactionOk

`func (o *AchatEnLigneModel) GetLatTransactionOk() (*string, bool)`

GetLatTransactionOk returns a tuple with the LatTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatTransaction

`func (o *AchatEnLigneModel) SetLatTransaction(v string)`

SetLatTransaction sets LatTransaction field to given value.

### HasLatTransaction

`func (o *AchatEnLigneModel) HasLatTransaction() bool`

HasLatTransaction returns a boolean if a field has been set.

### GetRequestId

`func (o *AchatEnLigneModel) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AchatEnLigneModel) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AchatEnLigneModel) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetUrlCallback

`func (o *AchatEnLigneModel) GetUrlCallback() string`

GetUrlCallback returns the UrlCallback field if non-nil, zero value otherwise.

### GetUrlCallbackOk

`func (o *AchatEnLigneModel) GetUrlCallbackOk() (*string, bool)`

GetUrlCallbackOk returns a tuple with the UrlCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCallback

`func (o *AchatEnLigneModel) SetUrlCallback(v string)`

SetUrlCallback sets UrlCallback field to given value.

### HasUrlCallback

`func (o *AchatEnLigneModel) HasUrlCallback() bool`

HasUrlCallback returns a boolean if a field has been set.

### GetUrlRetour

`func (o *AchatEnLigneModel) GetUrlRetour() string`

GetUrlRetour returns the UrlRetour field if non-nil, zero value otherwise.

### GetUrlRetourOk

`func (o *AchatEnLigneModel) GetUrlRetourOk() (*string, bool)`

GetUrlRetourOk returns a tuple with the UrlRetour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRetour

`func (o *AchatEnLigneModel) SetUrlRetour(v string)`

SetUrlRetour sets UrlRetour field to given value.

### HasUrlRetour

`func (o *AchatEnLigneModel) HasUrlRetour() bool`

HasUrlRetour returns a boolean if a field has been set.

### GetAdresseIp

`func (o *AchatEnLigneModel) GetAdresseIp() string`

GetAdresseIp returns the AdresseIp field if non-nil, zero value otherwise.

### GetAdresseIpOk

`func (o *AchatEnLigneModel) GetAdresseIpOk() (*string, bool)`

GetAdresseIpOk returns a tuple with the AdresseIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseIp

`func (o *AchatEnLigneModel) SetAdresseIp(v string)`

SetAdresseIp sets AdresseIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


