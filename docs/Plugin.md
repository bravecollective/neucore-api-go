# Plugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt32** |  | 
**Name** | **string** |  | 
**ConfigurationDatabase** | Pointer to [**PluginConfigurationDatabase**](PluginConfigurationDatabase.md) |  | [optional] 
**ConfigurationFile** | Pointer to [**PluginConfigurationFile**](PluginConfigurationFile.md) |  | [optional] 

## Methods

### NewPlugin

`func NewPlugin(id NullableInt32, name string, ) *Plugin`

NewPlugin instantiates a new Plugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginWithDefaults

`func NewPluginWithDefaults() *Plugin`

NewPluginWithDefaults instantiates a new Plugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plugin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plugin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plugin) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Plugin) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Plugin) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Plugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plugin) SetName(v string)`

SetName sets Name field to given value.


### GetConfigurationDatabase

`func (o *Plugin) GetConfigurationDatabase() PluginConfigurationDatabase`

GetConfigurationDatabase returns the ConfigurationDatabase field if non-nil, zero value otherwise.

### GetConfigurationDatabaseOk

`func (o *Plugin) GetConfigurationDatabaseOk() (*PluginConfigurationDatabase, bool)`

GetConfigurationDatabaseOk returns a tuple with the ConfigurationDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDatabase

`func (o *Plugin) SetConfigurationDatabase(v PluginConfigurationDatabase)`

SetConfigurationDatabase sets ConfigurationDatabase field to given value.

### HasConfigurationDatabase

`func (o *Plugin) HasConfigurationDatabase() bool`

HasConfigurationDatabase returns a boolean if a field has been set.

### GetConfigurationFile

`func (o *Plugin) GetConfigurationFile() PluginConfigurationFile`

GetConfigurationFile returns the ConfigurationFile field if non-nil, zero value otherwise.

### GetConfigurationFileOk

`func (o *Plugin) GetConfigurationFileOk() (*PluginConfigurationFile, bool)`

GetConfigurationFileOk returns a tuple with the ConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFile

`func (o *Plugin) SetConfigurationFile(v PluginConfigurationFile)`

SetConfigurationFile sets ConfigurationFile field to given value.

### HasConfigurationFile

`func (o *Plugin) HasConfigurationFile() bool`

HasConfigurationFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


