# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Configuration** | Pointer to [**ServiceConfiguration**](ServiceConfiguration.md) |  | [optional] 

## Methods

### NewService

`func NewService(id int32, name string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *Service) GetConfiguration() ServiceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Service) GetConfigurationOk() (*ServiceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Service) SetConfiguration(v ServiceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Service) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


