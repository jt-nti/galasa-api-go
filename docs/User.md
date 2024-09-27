# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginId** | Pointer to **string** | The ID representing the user&#39;s username. Used for authenticating and token creation. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginId

`func (o *User) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *User) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *User) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.

### HasLoginId

`func (o *User) HasLoginId() bool`

HasLoginId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


