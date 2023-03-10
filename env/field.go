package env

/*
Copyright 2023 Absa Group Limited

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func boolean(env env) (b bool, err error) {
	var d bool
	if env.def.exists {
		d, err = strconv.ParseBool(env.def.value)
		if err != nil {
			err = fmt.Errorf("can't convert default value %s of '%s' to bool", env.name, env.def.value)
			return
		}
	}
	b, err = GetEnvAsBoolOrFallback(env.name, d)
	return
}

func float(env env) (f float64, err error) {
	var d float64
	if env.def.exists {
		d, err = strconv.ParseFloat(env.def.value, 64)
		if err != nil {
			err = fmt.Errorf("can't convert default value %s of '%s' to float64", env.name, env.def.value)
			return
		}
	}
	f, err = GetEnvAsFloat64OrFallback(env.name, d)
	if err != nil {
		err = fmt.Errorf("can't read %s and parse value '%s' to float64", env.name, env.value)
	}
	return
}

func floatSlice(env env) (fs []float64, err error) {
	var d []float64
	if env.def.asStringSlice() != nil {
		d = make([]float64, 0)
		for _, s := range env.def.asStringSlice() {
			var fl float64
			fl, err = strconv.ParseFloat(strings.Trim(s, " "), 64)
			if err != nil {
				err = fmt.Errorf("can't convert default %s to slice of float64", env.def.asStringSlice())
				return
			}
			d = append(d, fl)
		}
	}
	fs, err = GetEnvAsArrayOfFloat64OrFallback(env.name, d)
	if err != nil {
		err = fmt.Errorf("can't parse %s as slice of float64 '%s'", env.name, env.value)
	}
	return
}

func boolSlice(env env) (bs []bool, err error) {
	var d []bool
	if env.def.asStringSlice() != nil {
		d = make([]bool, 0)
		for _, s := range env.def.asStringSlice() {
			var b bool
			b, err = strconv.ParseBool(strings.Trim(s, " "))
			if err != nil {
				err = fmt.Errorf("can't convert default %s to slice of bool", env.def.asStringSlice())
				return
			}
			d = append(d, b)
		}
	}
	bs, err = GetEnvAsArrayOfBoolOrFallback(env.name, d)
	if err != nil {
		err = fmt.Errorf("can't parse %s as slice of bool '%s'", env.name, env.value)
	}
	return
}
