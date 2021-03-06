// Copyright (c) 2017 Intel Corporation
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

package virtcontainers

import (
	"testing"
)

func testIsSandbox(t *testing.T, cType ContainerType, expected bool) {
	if result := cType.IsSandbox(); result != expected {
		t.Fatalf("Got %t, Expecting %t", result, expected)
	}
}

func TestIsPodSandboxTrue(t *testing.T) {
	testIsSandbox(t, PodSandbox, true)
}

func TestIsPodContainerFalse(t *testing.T) {
	testIsSandbox(t, PodContainer, false)
}

func TestIsSandboxUnknownContainerTypeFalse(t *testing.T) {
	testIsSandbox(t, UnknownContainerType, false)
}
