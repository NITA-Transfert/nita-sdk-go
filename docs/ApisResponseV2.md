# ApisResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApisResponseV2

`func NewApisResponseV2() *ApisResponseV2`

NewApisResponseV2 instantiates a new ApisResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApisResponseV2WithDefaults

`func NewApisResponseV2WithDefaults() *ApisResponseV2`

NewApisResponseV2WithDefaults instantiates a new ApisResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApisResponseV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApisResponseV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApisResponseV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApisResponseV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *ApisResponseV2) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApisResponseV2) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApisResponseV2) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApisResponseV2) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ApisResponseV2) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApisResponseV2) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApisResponseV2) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApisResponseV2) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApisResponseV2) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApisResponseV2) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApisResponseV2) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ApisResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


