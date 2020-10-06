/*
    Copyright (C) 2020 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package tfv12

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// list of available cty to golang type converters
var ctyConverterFuncs = []func(cty.Value) (interface{}, error){ctyToStr, ctyToInt, ctyToBool, ctyToSlice}

// ctyToStr tries to convert the given cty.Value into golang string type
func ctyToStr(ctyVal cty.Value) (interface{}, error) {
	var val string
	err := gocty.FromCtyValue(ctyVal, &val)
	return val, err
}

// ctyToInt tries to convert the given cty.Value into golang int type
func ctyToInt(ctyVal cty.Value) (interface{}, error) {
	var val int
	err := gocty.FromCtyValue(ctyVal, &val)
	return val, err
}

// ctyToBool tries to convert the given cty.Value into golang bool type
func ctyToBool(ctyVal cty.Value) (interface{}, error) {
	var val bool
	err := gocty.FromCtyValue(ctyVal, &val)
	return val, err
}

// ctyToSlice tries to convert the given cty.Value into golang slice of
// interfce{}
func ctyToSlice(ctyVal cty.Value) (interface{}, error) {
	var val []interface{}
	err := gocty.FromCtyValue(ctyVal, &val)
	return val, err
}