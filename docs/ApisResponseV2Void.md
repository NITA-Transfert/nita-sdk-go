# ApisResponseV2Void

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApisResponseV2Void

`func NewApisResponseV2Void() *ApisResponseV2Void`

NewApisResponseV2Void instantiates a new ApisResponseV2Void object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApisResponseV2VoidWithDefaults

`func NewApisResponseV2VoidWithDefaults() *ApisResponseV2Void`

NewApisResponseV2VoidWithDefaults instantiates a new ApisResponseV2Void object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApisResponseV2Void) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApisResponseV2Void) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApisResponseV2Void) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApisResponseV2Void) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *ApisResponseV2Void) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApisResponseV2Void) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApisResponseV2Void) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApisResponseV2Void) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ApisResponseV2Void) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApisResponseV2Void) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApisResponseV2Void) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApisResponseV2Void) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApisResponseV2Void) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApisResponseV2Void) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApisResponseV2Void) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ApisResponseV2Void) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


