# k8s-map-binder
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Reference](https://pkg.go.dev/badge/github.com/AbsaOSS/k8s-map-binder.svg)](https://pkg.go.dev/github.com/AbsaOSS/k8s-map-binder?branch=master)
![Build Status](https://github.com/AbsaOSS/k8s-map-binder/actions/workflows/build.yaml/badge.svg?branch=master)
![Linter](https://github.com/AbsaOSS/k8s-map-binder/actions/workflows/lint.yaml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/AbsaOSS/k8s-map-binder)](https://goreportcard.com/report/github.com/AbsaOSS/k8s-map-binder?branch=master)

My kubernetes operators usually read the annotations of operated resourcers.
This GO library provides simple way to parse annotations into custom structures. 
It supports slices in addition to primitive types. Besides the actual parsing, 
it also offers ways to bind fields such as default values, require, 
value protection, private fields and nested structures.

### QuickStart
