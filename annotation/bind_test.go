/*
Copyright 2021 The k8gb Contributors.

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

package annotation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testee public test structure

type arnold struct {
	// int value
	id int `annotation:"controller.example.com/id, require=true"` //nolint:unused
	// float value
	Budget float64
	// string value
	Name string `annotation:"controller.example.com/Name"`
	// private field
	surname string `annotation:"controller.example.com/surname, default=Schwarzenegger"`
	// slice value
	Ranks []string `annotation:"controller.example.com/ranks"`
	// bool value
	Armed bool `annotation:"controller.example.com/armed,protected=true"`
	// nested structure
	Inventory struct {
		//
		Rope string `annotation:"controller.example.com/rope, require=true"`
		//
		paperWithSecretKey string `annotation:"controller.example.com/secret, require=true"`
	}
}

func (a *arnold) String() string {
	if a.Armed {
		return fmt.Sprintf("Machine Gun %s %s, ranks:%v, inventory: [%s, %s]", a.Name, a.surname, a.Ranks, a.Inventory.Rope, a.Inventory.paperWithSecretKey)
	}
	return fmt.Sprintf("%s %s in full retreat, ranks:%v, inventory: [%s, %s]", a.Name, a.surname, a.Ranks, a.Inventory.Rope, a.Inventory.paperWithSecretKey)
}

func TestBind(t *testing.T) {
	// arrange
	getPredefinedInput := func() (map[string]string, *arnold) {
		return map[string]string{
			"controller.example.com/id":      "1500",
			"controller.example.com/Name":    "Arnold",
			"controller.example.com/surname": "Rimmer",
			"controller.example.com/ranks":   "officer1,capitan,general,major",
			"controller.example.com/armed":   "false",
			"controller.example.com/rope":    "dipper rope",
			"controller.example.com/secret":  "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		}, new(arnold)
	}

	tests := []struct {
		name           string
		expectedResult string
		expectedError  bool
		modifier       func(map[string]string, *arnold)
	}{
		{
			name:           "private field",
			expectedResult: "Arnold Rimmer in full retreat, ranks:[officer1 capitan general major], inventory: [dipper rope, wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY]",
			expectedError:  false,
			modifier:       func(m map[string]string, a *arnold) {},
		},
		{
			name:           "protected field",
			expectedResult: "Machine Gun Arnold Rimmer, ranks:[officer1 capitan general major], inventory: [dipper rope, wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY]",
			expectedError:  false,
			modifier: func(m map[string]string, a *arnold) {
				a.Armed = true
			},
		},
		{
			name:          "missing required field",
			expectedError: true,
			modifier: func(m map[string]string, a *arnold) {
				delete(m, "controller.example.com/id")
			},
		},
		{
			name:          "empty required field",
			expectedError: true,
			modifier: func(m map[string]string, a *arnold) {
				m["controller.example.com/id"] = ""
			},
		},
		{
			name:           "empty non required field",
			expectedError:  false,
			expectedResult: "Arnold  in full retreat, ranks:[officer1 capitan general major], inventory: [dipper rope, wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY]",
			modifier: func(m map[string]string, a *arnold) {
				m["controller.example.com/surname"] = ""
			},
		},
		{
			name:          "empty map",
			expectedError: true,
			modifier: func(m map[string]string, a *arnold) {
				for k := range m {
					delete(m, k)
				}
			},
		},
	}

	// act
	// assert
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, a := getPredefinedInput()
			test.modifier(m, a)
			err := Bind(m, a)
			assert.Equal(t, test.expectedError, err != nil)
			if !test.expectedError {
				assert.Equal(t, test.expectedResult, a.String())
			}
		})
	}
}

func TestBreaking(t *testing.T) {
	tests := []struct {
		name          string
		m             map[string]string
		s             interface{}
		expectedError bool
	}{
		{name: "nil map", m: nil, s: &arnold{}, expectedError: false},
		{name: "nil structure", m: map[string]string{}, s: nil, expectedError: true},
	}
	// act
	// assert
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Bind(test.m, test.s)
			assert.Equal(t, err != nil, test.expectedError)
		})
	}
}
