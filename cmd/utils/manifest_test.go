// Copyright 2022 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWorkdirFromManifest(t *testing.T) {
	var tests = []struct {
		name         string
		path         string
		expectedPath string
	}{
		{
			name:         "inside .okteto folder",
			path:         ".okteto/okteto.yml",
			expectedPath: ".",
		},
		{
			name:         "one path ahead",
			path:         "test/okteto.yml",
			expectedPath: "test",
		},
		{
			name:         "same path",
			path:         "okteto.yml",
			expectedPath: ".",
		},
		{
			name:         "full path",
			path:         "/usr/okteto.yml",
			expectedPath: "/usr",
		},
		{
			name:         "full path on .okteto",
			path:         "/usr/.okteto/okteto.yml",
			expectedPath: "/usr",
		},
		{
			name:         "relative path with more than two paths ahead",
			path:         "~/app/.okteto/okteto.yml",
			expectedPath: "~/app",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetWorkdirFromManifestPath(tt.path)
			assert.Equal(t, tt.expectedPath, result)
		})
	}
}
