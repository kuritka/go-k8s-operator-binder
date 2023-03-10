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

func convertToFloat32(ar []float64) (newar []float32) {
	if ar == nil {
		return
	}
	newar = make([]float32, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = float32(v)
	}
	return
}

func convertToInt(ar []float64) (newar []int) {
	if ar == nil {
		return
	}
	newar = make([]int, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int(v)
	}
	return
}

func convertToInt8(ar []float64) (newar []int8) {
	if ar == nil {
		return
	}
	newar = make([]int8, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int8(v)
	}
	return
}

func convertToInt16(ar []float64) (newar []int16) {
	if ar == nil {
		return
	}
	newar = make([]int16, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int16(v)
	}
	return
}

func convertToInt32(ar []float64) (newar []int32) {
	if ar == nil {
		return
	}
	newar = make([]int32, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int32(v)
	}
	return
}

func convertToInt64(ar []float64) (newar []int64) {
	if ar == nil {
		return
	}
	newar = make([]int64, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int64(v)
	}
	return
}

func convertToUInt(ar []float64) []uint {
	if ar == nil {
		return nil
	}
	newar := make([]uint, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = uint(v)
	}
	return newar
}

func convertToUInt8(ar []float64) (newar []uint8) {
	if ar == nil {
		return
	}
	newar = make([]uint8, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = uint8(v)
	}
	return
}

func convertToUInt16(ar []float64) (newar []uint16) {
	if ar == nil {
		return
	}
	newar = make([]uint16, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = uint16(v)
	}
	return
}

func convertToUInt32(ar []float64) (newar []uint32) {
	if ar == nil {
		return nil
	}
	newar = make([]uint32, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = uint32(v)
	}
	return
}

func convertToUInt64(ar []float64) (newar []uint64) {
	if ar == nil {
		return nil
	}
	newar = make([]uint64, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = uint64(v)
	}
	return
}
