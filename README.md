# k8s-go-map-binder
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Reference](https://pkg.go.dev/badge/github.com/AbsaOSS/k8s-map-binder.svg)](https://pkg.go.dev/github.com/AbsaOSS/k8s-map-binder?branch=master)
![Build Status](https://github.com/AbsaOSS/k8s-map-binder/actions/workflows/build.yaml/badge.svg?branch=master)
![Linter](https://github.com/AbsaOSS/k8s-map-binder/actions/workflows/lint.yaml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/AbsaOSS/k8s-map-binder)](https://goreportcard.com/report/github.com/AbsaOSS/k8s-map-binder?branch=master)

My Kubernetes operators usually read configurations from two sources. The first is Config values that sit in 
Environment Variables. The second configuration is obtained at runtime by reading annotations and labels.

This library can bind these configurations to GO structures in a uniform way. Besides the actual parsing,
it also offers ways to bind fields such as default values, require, value protection, private fields and 
nested structures.

## Table of Content
- [QuickStart](#quickstart)
  - [K8s binder](#k8s binder)
  - [Environment variables binder](#Environment variables binder)
- [Supported types](#supported types)
- [Supported keywords](#supported keywords)


## QuickStart
Although this library is primarily developed for operators, it can be used anywhere you work with 
Environment Variables or `map[string]string`.

The great advantage of this library is its ease of use and the fact that it allows you to read the configuration from multiple 
sources without having to learn something new. A few keywords will make it much easier for you to perform basic operations like 
setting default values or forcing a value.

### K8s Binder

### Environment variables binder

## Supported types
GO-K8S-OPERATOR-BINDER supports all types listed in the following table.  In addition, it should be noted that in the case
of slices, GO-K8S-OPERATOR-BINDER creates an instance of an empty slice if the value of the environment variable is
declared and its value is empty string. In this case GO-K8S-OPERATOR-BINDER returns an empty slice instead of the vulnerable nil.

| primitive types | slices |
|---|---|
| `int`,`int8`,`int16`,`int32`,`int64` | `[]int`,`[]int8`,`[]int16`,`[]int32`,`[]int64` |
| `float32`,`float64` | `[]float32`,`[]float64` |
| `uint`,`uint8`,`uint16`,`uint32`,`uint64` | `[]uint`,`[]uint8`,`[]uint16`,`[]uint32`,`[]uint64` |
| `bool` | `[]bool` |
| `string` | `[]string` |

## Supported keywords
Besides the fact that GO-K8S-OPERATOR-BINDER works with private fields and can add prefixes to variable names, it
operates with several keywords. The structure in the introductory section works with all types
of these keywords.

- `default` - the value specified in the default tag is used in case env variable does not exist. e.g:
  `env: "SUBNET", default=10.0.1.0/24` or `env: "ENV_SUBNETS", default=[]` which will set an empty slice instead
  of a vulnerable nil value in case `ENV_SUBNETS` does not exist.

- `require` - if `require=true` then env variable must exist otherwise Bind function returns error

- `protected` - if `protected=true` then, in case the field in the structure already has a set value , the
  Bind function will not set it. Otherwise, bind will be applied to it.

You can combine individual tags freely: `env: "ENV_SWITCHER", default=[true, false, true], protected=true`
is a perfectly valid configuration

