# GetPartenaireToWalletModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** |  | 
**CodeRecharge** | Pointer to **string** | Facultatif : absent, la recharge est retrouvée par le seul requestId (reprise après un timeout à la création). | [optional] 

## Methods

### NewGetPartenaireToWalletModel

`func NewGetPartenaireToWalletModel(requestId string, ) *GetPartenaireToWalletModel`

NewGetPartenaireToWalletModel instantiates a new GetPartenaireToWalletModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPartenaireToWalletModelWithDefaults

`func NewGetPartenaireToWalletModelWithDefaults() *GetPartenaireToWalletModel`

NewGetPartenaireToWalletModelWithDefaults instantiates a new GetPartenaireToWalletModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetPartenaireToWalletModel) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetPartenaireToWalletModel) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetPartenaireToWalletModel) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetCodeRecharge

`func (o *GetPartenaireToWalletModel) GetCodeRecharge() string`

GetCodeRecharge returns the CodeRecharge field if non-nil, zero value otherwise.

### GetCodeRechargeOk

`func (o *GetPartenaireToWalletModel) GetCodeRechargeOk() (*string, bool)`

GetCodeRechargeOk returns a tuple with the CodeRecharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRecharge

`func (o *GetPartenaireToWalletModel) SetCodeRecharge(v string)`

SetCodeRecharge sets CodeRecharge field to given value.

### HasCodeRecharge

`func (o *GetPartenaireToWalletModel) HasCodeRecharge() bool`

HasCodeRecharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


