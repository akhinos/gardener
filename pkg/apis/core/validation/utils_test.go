// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation_test

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func makeDurationPointer(d time.Duration) *metav1.Duration {
	return &metav1.Duration{Duration: d}
}

func makeFloat64Pointer(f float64) *float64 {
	ptr := f
	return &ptr
}

func makeIntPointer(i int) *int {
	ptr := i
	return &ptr
}

func makeInt32Pointer(i int32) *int32 {
	ptr := i
	return &ptr
}

func makeBoolPointer(i bool) *bool {
	ptr := i
	return &ptr
}

func makeStringPointer(s string) *string {
	ptr := s
	return &ptr
}
