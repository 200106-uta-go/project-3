<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][https://github.com/200106-uta-go/project-3/graphs/contributors]
[![Forks][forks-shield]][https://github.com/200106-uta-go/project-3/network/members]
[![Stargazers][stars-shield]][https://github.com/200106-uta-go/project-3/stargazers]
[![Issues][issues-shield]][https://github.com/200106-uta-go/project-3/issues]


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/200106-uta-go/project-3">
    <img src="Revature.png" alt="Logo" width="480" height="160">
  </a>

  <h3 align="center">Revature</h3>

  <p align="center">
    A custom ingress controller to route failed requests between Kubernetes clusters. A custom CLI for the deployment of Helm, Istio, Jaeger, Prometheus, and Grafana.
    <br />
    <a href="https://github.com/200106-uta-go/project-3"><strong>Explore the docs »</strong></a>
    <br />
    <br />
     <a href="https://github.com/200106-uta-go/project-3/issues">Report Bug</a>
    ·
    <a href="https://github.com/200106-uta-go/project-3/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
  * [Creating a New Profile](#Creating-a-New-Profile)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

### Revature Hybrid Ingress Controller Tasks
* Create a CLI tool that functions similar to helm create, which creates an empty Helm chart scaffold set up with the following dependencies fully configured
    * Istio
    * Jaeger
    * Grafana
* Implement a custom Kubernetes resource Cluster, which represents necessary details of some Revature Kubernetes Cluster, Cluster B
* Create a Custom Ingress Controller that can be deployed in some Kubernetes Cluster Cluster A such that
    * If a request is made to Cluster A and fails for any reason, the request will be retried against the same Service in Cluster B
    * If the retried request is made to Cluster B, then returning a failed response is acceptable

### Built With

* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)
* [Kubernetes](https://kubernetes.io/)
* [Terraform](https://www.terraform.io/)
* [Istio](https://istio.io/)
* [Jaeger](https://www.jaegertracing.io/)
* [Grafana](https://grafana.com)
* [Helm](https://helm.sh/)
* [Prometheus](https://prometheus.io/)



<!-- GETTING STARTED -->
## Getting Started

This repository contains the source code for the Kreate cli, and Custom Ingress controller.

### Prerequisites

- Kubernetes cluster must be already active for Kreate and Custom Ingress to function.

- Uses Helm 2.10+ and not Helm 3, currently installs newest version of Helm 2 durring initialization.

### Installation

1. Build Kreate, navagate to project root and run build command below
```sh
go build ./cmd/kreate
```
2. Initilize Kreate
```sh
Kreate init
```

<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

```
kreate <sub-command> [PROFILE_NAME]
```

_For more examples, please refer to the [Documentation](https://example.com)_

### Creating a New Profile

To create a new profile use the following command.

```bash
kreate profile [PROFILE_NAME]
```

This command creates a new folder named `/kreate` under `/etc/` directory and places a new `.yaml` file with the provided named.

*Example:*

```bash
kreate profile myprofile
```

*Output:*

```view
/etc
└── /kreate
    └── myprofile.yaml
```

### Outputting a Chart
```
Kreate chart [PROFILE_NAME]
```

CreateChart is a function that generates a values.yaml, Chart.yaml, yaml templates for use with helm, and already-templated yamls ready for deployment in a Kubernetes cluster. 

When this command is used, a charts folder will be added to your current working directory with the following structure.
```
.
└── charts
    └── example
        ├── Chart.yaml
        ├── deploy
        │   ├── deployment.yaml
        │   ├── ingress.yaml
        │   └── service.yaml
        ├── templates
        │   ├── deployment.yaml
        │   ├── ingress.yaml
        │   └── service.yaml
        └── values.yaml
```
The `charts` directory is where all charts generated using CreateChart will be located. Each folder underneath `charts`, will be a separate chart based on a unique kreate profile. If the program is run multiple times without editing the `name` value in `Chart.yaml`, the new deployment will overwrite any existing chart with the same name.

Within each unique chart folder, the `deploy` folder will hold already-templated .yaml files ready for deployment using [kubectl apply](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#apply).

The `templates` folder will hold a copy of the templates stored in `/var/local/kreate` that are used to generate the templated yaml files in the `deploy` folder. These templates are for use with [Go text templating](https://golang.org/pkg/text/template/), and can be used directly with [Helm](https://v2.helm.sh/docs/) or expanded with more templated values.

### Using Run Command to apply a Profile
```
kreate run [PROFILE_NAME]
```

RunProfile is a function which utilizes helm to deploy a profile directly to the Kubernetes Cluster. Given a profile name, RunProfile will attempt the following\:
1. Determine if Helm is properly initialized. The tiller must be installed to the cluster prior to running a profile
2. Build a custom Helm chart for the specified profile using kreate.CreateChart()
3. Create the custom ingress configmap and install the portal custom resource to the cluster
4. Install the Helm chart, or if the profile was previously deployed, Upgrade the existing installation.

RunProfile anticipates that Kreate.InitializeEnvironment() has been completed successfully. **Thus, the user is required to run kreate Init prior to kreate Run.**

### Using Remove Command to remove a Profile
```
kreate remove [PROFILE_NAME | --all | -a]
```

The remove command removes a specified profile from /etc/kreate/ directory. When using --all (or the shorthand -a) inplace of a profile name, all profiles will be removed.

### Using the edit Command to change a Profile
```
kreate edit [PROFILE_NAME]
```

A function that receives a profile struct and allows edits based on set flags to create a new profile struct that is used to `update` the yaml file with the same name corresponding to the input profile struct.

#### Getting Started

Our implementation will allow changes to Cluster specific entries by setting values to each flag, each flag is checked to see if the value is different from the default flag values and if so, an instance of a *Profile* will be created and its values will be set to reflect the values set in the flags. After all values have been set, the *profile* instance wil be used to write to a new file with the same file name and delete the old instance of the yaml. 

To view all the configuration flags for edit will be located in: 
`./kreate usage`

*Note* that the `name` flag must be set to an existing app's name within the yaml file, otherwise the function will not change any values corresponding to any app, and a message will be logged to the user that a name was not correctly specified.

#### Configuration flags specific to Cluster

This section below list all the configuration flags for cluster related settings.

> `name` - *Sets the name of profile*

> `clustername` - *Sets the clustername of the profile*

> `clusterip` - *Sets the clusterip of the profile*

> `clusterport` - *Append a clusterport to the profile*

>
#### Configuration flags specific to individual app

This section below list all configuration flags for app specific setting.

> `NameOfApp` - *Specifies the name of the app which will be modified by the app-related input flags*

> `imageurl` - *An App-related flag. Sets the imageurl of the App specified by the NameOfApp flag*

> `servicename` - *An App-related flag. Sets the servicename of the App specified by the NameOfApp flag*

> `serviceport` - *An App-related flag. Sets the serviceport of the App specified by the NameOfApp flag*

> `port` - *An App-related flag. Appends a port to the App specified by the NameOfApp flag*

> `endpoint` - *An App-related flag. Appends an endpoint to the App specified by the NameOfApp flag*

#### Prerequisites

This function requires the name of the yaml in the form of `"defaultName"`.

Before starting, within the cmd folder, initialization must be executed with:

1. `./kreate init`

Create a profile yaml named *defaultName.yaml* and store it under */etc/kreate/*, which is define as *kreate.PROFILES*.

2. `./kreate profile defaultName`

Call `Edit` to profile to change the values of this yaml to reflect how the configuration of the cluster and apps within the cluster should be.

`./kreate edit defaultName`
*Note* Without passing any flag values nano text editor will open the yaml file and allow manual edits from that program, otherwise after defining the filename, flags can be passed to this function in the form of *-flag FlagValue*.

### How to view the Help Text


```
kreate help
```

This sub-command will display a brief helpt text to familiarize the user with various commands associated with the tool.

<!-- ROADMAP -->
## Roadmap
### Features 
* Allow the creation of multiple profiles
* Add a command to reset the cluster to a pre-init configuration
* Add a command to uninstall a profile chart from the cluster
* Add portals that query additional times
* Remove helm dependencies
* Configure metrics applications outside of Istio profile
* Add a command to add a prometheus exporter to the cluster


See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

See also the list of [contributors](https://github.com/200106-uta-go/project-3/graphs/contributors) who participated in this project.



<!-- LICENSE -->



<!-- CONTACT -->
## Contact

-Last Name | First Name | Responsibilities | Github User
-----------|------------|------------------|------------
-Ackard | Matt | CRD,CLI | [mattackard](https://github.com/mattackard)
-Bland | Jessey | CLI | [JesseyBland](https://github.com/JesseyBland)
-Campbell | Nehemiah | CLI | [NehemiahG7](https://github.com/NehemiahG7)
-Estrada | Dania | Ingress | 
-Feliciano | Emilio | CLI | [FelicianoEJ](https://github.com/FelicianoEJ)
-Kim | Aaron | CLI | [ajkim19](https://github.com/ajkim19)
-Locker | Brandon | CLI | [Gamemastertwig](https://github.com/Gamemastertwig)
-McDole | Ken | CRD,CLI | [ken343](https://github.com/ken343)
-Moreno | Hector | CLI | 
-Nguyen | Josh | CLI | [CodeZipline](https://github.com/CodeZipline)
-Oh | Jaeik | CLI,Visuals | [flyerjayden](https://github.com/flyerjayden)
-Theiss | Joseph | Ingress | [jtheiss19](https://github.com/jtheiss19)
-Thomas | Zach | Ingress | [zachthomas823](https://github.com/zachthomas823)
-Zoeller | Joseph | CLI | [JosephZoeller](https://github.com/JosephZoeller)
