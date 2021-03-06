/*
 * Copyright 2020 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	"reflect"
	"testing"
)

func TestConfigMapPropagation_GetGroupVersionKind(t *testing.T) {
	cmp := ConfigMapPropagation{}
	gvk := cmp.GetGroupVersionKind()
	if gvk.Kind != "ConfigMapPropagation" {
		t.Errorf("Should be ConfigMapPropagation.")
	}
}
func TestConfigMapPropagation_GetUntypedSpec(t *testing.T) {
	cmp := ConfigMapPropagation{}
	in := cmp.GetUntypedSpec()
	if !reflect.DeepEqual(in, cmp.Spec) {
		t.Errorf("Should be ConfigMapPropagationSpec.")
	}
}
