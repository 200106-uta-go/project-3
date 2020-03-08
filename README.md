<!--
*** Thanks for checking out this README Template. If you have a suggestion that would
*** make this better, please fork the repo and create a pull request or simply open
*** an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->





<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTAiMpuDegbDyd4bjoWMCi8MbKyo2epjq9rrkyDx6dQEP9PwRcc" alt="Logo" width="160" height="80">
  </a>

  <h3 align="center">Revature</h3>

  <p align="center">
    An awesome README template to jumpstart your projects!
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template">View Demo</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Report Bug</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Request Feature</a>
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
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

There are many great README templates available on GitHub, however, I didn't find one that really suit my needs so I created this enhanced one. I want to create a README template so amazing that it'll be the last one you ever need.

Here's why:
* Your time should be focused on creating something amazing. A project that solves a problem and helps others
* You shouldn't be doing the same tasks over and over like creating a README from scratch
* You should element DRY principles to the rest of your life :smile:

Of course, no one template will serve all projects since your needs may be different. So I'll be adding more in the near future. You may also suggest changes by forking this repo and creating a pull request or opening an issue.

A list of commonly used resources that I find helpful are listed in the acknowledgements.

### Built With
This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Bootstrap](https://getbootstrap.com)
* [JQuery](https://jquery.com)
* [Laravel](https://laravel.com)



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

<<<<<<< HEAD
### Chart Sub-Command
=======
### Creating a new Profile

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

### kreate.CreateChart
>>>>>>> efd3bab437210986d5e3992c96f9e5bdcc6cea5d
```
kreate chart my_profile
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

### Run Command
```
kreate run my_profile
```

RunProfile is a function which utilizes helm to deploy a profile directly to the Kubernetes Cluster. Given a profile name, RunProfile will attempt the following\:
1. Determine if Helm is properly initialized. The tiller must be installed to the cluster prior to running a profile
2. Build a custom Helm chart for the specified profile using kreate.CreateChart()
3. Create the custom ingress configmap and install the portal custom resource to the cluster
4. Install the Helm chart, or if the profile was previously deployed, Upgrade the existing installation.

RunProfile anticipates that Kreate.InitializeEnvironment() has been completed successfully. **Thus, the user is required to run kreate Init prior to kreate Run.**

### Edit Command
```
kreate edit my_profile
```

A function that receives a profile struct and allows edits based on set flags to create a new profile struct that is used to `update` the yaml file with the same name corresponding to the input profile struct.

### kreate remove [PROFILE_NAME / --all / -a]

The `remove` command removes a specified profile from `/etc/kreate/` directory. When using `--all` (or the shorthand `-a`) inplace of a profile name, all profile will be removed.

#### Getting Started

Our implementation will allow changes to Cluster specific entries by setting values to each flag, the `FlagName` will map to the value within the file pointed to by the YamlFileName to modify their values.

*Note* that the `AppName` flag must be set to an existing app's name within the yaml file, otherwise the function will not change any values and log to the user that a name was not correctly specified.

#### Configuration flags specific to Cluster

This section below list all the configuration flags for cluster related settings.

> `Name` - *Name for config*

> `ClusterName` - *ClusterName for config*

> `ClusterIP` - *ClusterIp for config*

> `ClusterPort` - *ClusterPort for config*

>
#### Configuration flags specific to individual app

This section below list all configuration flags for app specific setting.

> `AppName` - *Under App, the Name value*

> `AppImageURL` - *Under App, the ImageURL*

> `AppServiceName` - *Under App, the ServiceName value*

> `AppServicePort` - *Under App, the ServicePort value*

> `AppPort` - *Under App, Port Value*

> `AppEndpoint` - *Under App, Endpoint Value*

#### Prerequisites

This function requires the name of the yaml in the form of `"defaultName"`.

Before starting, initialization must be executed with:

1. `kreate.Initialization()`

Create a profile yaml named *defaultName.yaml* and store it under */etc/kreate/*, which is define as *kreate.PROFILES*.

2. `kreate.CreateProfile("defaultName")`

Call `EditProfile()` to profile to change the values of this yaml to reflect how the configuration of the cluster and apps within the cluster should be.

`kreate.EditProfile("defaultName")`

- Note that edit calls GetProfile to create an instance of a *kreate.Profile* from the name of the yaml file, and unmarshal the values into that instance:

`profileInstance := kreate.GetProfile("defaultName")`

### Help Command

```
kreate help
```

This sub-command will display a brief helpt text to familiarize the user with various commands associated with the tool.

<!-- ROADMAP -->
## Roadmap

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

Your Name - [@your_twitter](https://twitter.com/your_username) - email@example.com

Project Link: [https://github.com/your_username/repo_name](https://github.com/your_username/repo_name)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Img Shields](https://shields.io)
* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Pages](https://pages.github.com)
* [Animate.css](https://daneden.github.io/animate.css)
* [Loaders.css](https://connoratherton.com/loaders)
* [Slick Carousel](https://kenwheeler.github.io/slick)
* [Smooth Scroll](https://github.com/cferdinandi/smooth-scroll)
* [Sticky Kit](http://leafo.net/sticky-kit)
* [JVectorMap](http://jvectormap.com)
* [Font Awesome](https://fontawesome.com)





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=flat-square
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=flat-square
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=flat-square
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=flat-square
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=flat-square
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
