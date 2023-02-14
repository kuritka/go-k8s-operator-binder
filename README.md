# k8s-go-map-binder
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Reference](https://pkg.go.dev/badge/github.com/kuritka/go-k8s-operator-binder.svg)](https://pkg.go.dev/github.com/kuritka/go-k8s-operator-binder?branch=main)
![Build Status](https://github.com/kuritka/go-k8s-operator-binder/actions/workflows/test.yaml/badge.svg?branch=main)
![Linter](https://github.com/kuritka/go-k8s-operator-binder/actions/workflows/lint.yaml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kuritka/go-k8s-operator-binder)](https://goreportcard.com/report/github.com/kuritka/go-k8s-operator-binder?branch=main)

My Kubernetes operators usually read configurations from two sources. The first is Config values that sit in 
Environment Variables. The second configuration is obtained at runtime by reading annotations and labels.

This library can bind these configurations to GO structures in a uniform way. Besides the actual parsing,
it also offers ways to bind fields such as default values, require, value protection, private fields and 
nested structures.

## Table of Content
- [QuickStart](#quickstart)
  - [K8s binder](#k8s-binder)
  - [Environment variables binder](#environment-variables-binder)
- [Supported types](#supported-types)
- [Supported keywords](#supported-keywords)


## QuickStart
Although this library is primarily developed for operators, it can be used anywhere you work with 
Environment Variables or `map[string]string`.

The great advantage of this library is its ease of use and the fact that it allows you to read the configuration from multiple 
sources without having to learn other libs. A few keywords will make it much easier for you to perform basic operations like 
setting default values or forcing a value.

### K8s Binder
The **ENV-GO-K8S-OPERATOR-BINDER/k8s** package is used to easily bind kubernetes annotations and labels into GO structures.
Effectively it is about `map[string]string` configurations.


Imagine we have a resource that contains the following annotations (it works exactly the same for Labels).
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    company.example.io/strategy: failover
    company.example.io/primary-geotags: us,za
    company.example.io/ttl-seconds: "30"
    company.example.io/autoscale: "true"
    company.example.io/uid: "eyJhbGciOiJ"
  name: ing
```

```golang
import "github.com/kuritka/go-k8s-operator-binder/k8smap"

var settings = struct {
	// binding required value
    Type          string   `k8smap"company.example.io/strategy, require=true"`
	// binding slice of strings
    PrimaryGeoTag []string `k8smap:"company.example.io/primary-geotags"`
	// if the configuration is missing, then set the default int value to 30 
    TTLSeconds int         `k8smap:"k8gb.io/dns-ttl-seconds, default=30"`
    // binding bool value
    Autoscale bool         `k8smap:"k8gb.io/dns-ttl-seconds"`
	// nested structure
    Credentials struct{
        // public nested protected with default test 
        UID       `k8smap:"k8gb.io/uid, protected=true, default=test`
        // private nested protected with default test
        base64    `k8smap:"k8gb.io/uid, protected=true, default=MMaBcDeFgH...`
    }
}{}

// Fetch resource instance
ing := &netv1.Ingress{}
err := r.Get(ctx, req.NamespacedName, ing)
if err != nil {
	// do Something
}
annotations = ing.GetAnnotations()
// binding 
err = k8smap.Bind(annotations, &settings)
if err != nil {
	// do something
}
// fmt.Println(settings)
```
This is all you need, the `settings` will contain the correct values loaded from the Annotations.
Keywords like `protected` or `default` can be configured in various ways as well as various data-types, 
see below in this documentation.

### Environment variables binder
The **ENV-GO-K8S-OPERATOR-BINDER/env** package is used to easily bind environment variables to GO structures. The package 
is designed to be usable in the widest possible range of scenarios. Among other things, it supports variable prefixes 
and bindings to unexported arrays. Take a look at the following usage example:
```golang
import "github.com/kuritka/go-k8s-operator-binder/env"

type Endpoint struct {
	URL string `env:"ENDPOINT_URL, require=true"`
}

type Config struct {

	// reading string value from NAME
	Name string `env:"NAME"`

	// reuse an already bound env variable NAME
	Description string `env:"NAME"`

	// reuse an already bound variable NAME, but replace only when name 
	// was not set before
	AlternativeName string `env:"NAME, protected=true"`

	// reading int with 8080 as default value
	DefaultPort uint16 `env:"PORT, default=8080"`

	// reading slice of strings with default values
	Regions []string `env:"REGIONS, default=[us-east-1,us-east-2,us-west-1]"`

	// reading slice of strings from env var
	Subnets []string `env:"SUBNETS, default=[10.0.0.0/24,192.168.1.0/24]"`
	
	// default=[] ensures that if INTERVALS does not exist, 
	// []int8{} is set instead of []int8{nil}
	Interval []uint8 `env:"INTERVALS, default=[]"`

	// nested structure
	Credentials struct {

		// binding required value
		KeyID string `env:"ACCESS_KEY_ID, require=true"`

		// binding to private field
		secretKey string `env:"SECRET_ACCESS_KEY, require=true"`
	}

	// expected PRIMARY_ prefix in nested environment variables
	Endpoint1 Endpoint `env:"PRIMARY"`

	// expected FAILOVER_ prefix in nested environment variables
	Endpoint2 Endpoint `env:"FAILOVER"`

	// the field does not have a bind tag set, 
	// so it will be ignored during bind
	Args []string
}


func main() {
	defer clean()
	os.Setenv("PRIMARY_ENDPOINT_URL", "https://ep1.cloud.example.com")
	os.Setenv("FAILOVER_ENDPOINT_URL", "https://ep2.cloud.example.com")
	os.Setenv("ACCESS_KEY_ID", "AKIAIOSFODNN7EXAMPLE")
	os.Setenv("SECRET_ACCESS_KEY", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")
	os.Setenv("NAME", "Hello from 12-factor")
	os.Setenv("PORT", "9000")
	os.Setenv("SUBNETS", "10.0.0.0/24,10.0.1.0/24, 10.1.0.0/24,  10.1.1.0/24")

	c := &Config{}
	c.AlternativeName = "protected name"
	c.Description = "hello from ENV-BINDER"
	if err := env.Bind(c); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(JSONize(c))
}

```
function main generates the following output:
```json
{
  "Name": "Hello from 12-factor",
  "Description": "Hello from 12-factor",
  "AlternativeName": "protected name",
  "DefaultPort": 9000,
  "Regions": [
    "us-east-1",
    "us-east-2",
    "us-west-1"
  ],
  "Subnets": [
    "10.0.0.0/24",
    "10.0.1.0/24",
    "10.1.0.0/24",
    "10.1.1.0/24"
  ],
  "Interval": [],
  "Credentials": {
    "KeyID": "AKIAIOSFODNN7EXAMPLE"
  },
  "Endpoint1": {
    "URL": "https://ep1.cloud.example.com"
  },
  "Endpoint2": {
    "URL": "https://ep2.cloud.example.com"
  },
  "Args": null
}
```

## Supported types
**GO-K8S-OPERATOR-BINDER** supports all types listed in the following table.  In addition, it should be noted that in the case
of slices, **GO-K8S-OPERATOR-BINDER** creates an instance of an empty slice if the value of the environment variable is
declared and its value is empty string. In this case **GO-K8S-OPERATOR-BINDER** returns an empty slice instead of the vulnerable nil.

| primitive types | slices |
|---|---|
| `int`,`int8`,`int16`,`int32`,`int64` | `[]int`,`[]int8`,`[]int16`,`[]int32`,`[]int64` |
| `float32`,`float64` | `[]float32`,`[]float64` |
| `uint`,`uint8`,`uint16`,`uint32`,`uint64` | `[]uint`,`[]uint8`,`[]uint16`,`[]uint32`,`[]uint64` |
| `bool` | `[]bool` |
| `string` | `[]string` |

## Supported keywords
Besides the fact that **GO-K8S-OPERATOR-BINDER** works with private fields and can add prefixes to variable names, it
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

