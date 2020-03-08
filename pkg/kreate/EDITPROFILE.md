# kreate.EditProfile

A function that receives a profile struct and allows edits based on set flags to create a new profile struct that is used to `update` the yaml file with the same name corresponding to the input profile struct.

## Getting Started

Our implementation will allow changes to Cluster specific entries by setting values to each flag, the `FlagName` will map to the value within the file pointed to by the YamlFileName to modify their values.

*Note* that the `AppName` flag must be set to an existing app's name within the yaml file, otherwise the function will not change any values and log to the user that a name was not correctly specified.

### Configuration flags specific to Cluster

This section below list all the configuration flags for cluster related settings.

> `Name` - *Name for config*

> `ClusterName` - *ClusterName for config*

> `ClusterIP` - *ClusterIp for config*

> `ClusterPort` - *ClusterPort for config*

>
### Configuration flags specific to individual app

This section below list all configuration flags for app specific setting.

> `AppName` - *Under App, the Name value*

> `AppImageURL` - *Under App, the ImageURL*

> `AppServiceName` - *Under App, the ServiceName value*

> `AppServicePort` - *Under App, the ServicePort value*

> `AppPort` - *Under App, Port Value*

> `AppEndpoint` - *Under App, Endpoint Value*

### Prerequisites

This function requires the name of the yaml in the form of `"defaultName"`.

Before starting, initialization must be executed with:

1. `kreate.Initialization()`

Create a profile yaml named *defaultName.yaml* and store it under */etc/kreate/*, which is define as *kreate.PROFILES*.

2. `kreate.CreateProfile("defaultName")`

Create an isntance of a *kreate.Profile* from the name of the yaml file, and unmarshal the values into that isntance.

3. `profileInstance := kreate.GetProfile("defaultName")`

Call `editprofile()` to profile to change the values of this yaml to reflect how the configuration of the cluster and apps within the cluster should be.

`kreate.editprofile("defaultName")`

## Deployment

// Add additional notes about how to deploy this on a live system

## Built With

* [Yaml]("gopkg.in/yaml.v2") - The configuration mark up language

## Authors

* **Joshua Nguyen** - *Co-Author* - [CodeZipline](https://github.com/CodeZipline)

* **Hector Moreno** - *Co-Author* - [higgyhiggy](https://github.com/higgyhiggy)

See also the list of [contributors](https://github.com/200106-uta-go/project-3/graphs/contributors) who participated in this project.
