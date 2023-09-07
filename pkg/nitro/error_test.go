/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package nitro

import (
	"fmt"
	"testing"
)

func TestError_Error(t *testing.T) {
	var tests = []struct {
		err  Error
		want string
	}{
		{
			err:  NSERR_SSL_NOCERT,
			want: "Nitro API specific error: Certificate does not exist",
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.err.code)
		t.Run(testName, func(t *testing.T) {
			if tt.err.Error() != tt.want {
				t.Errorf("Error() = %s, expected %s", tt.err.Error(), tt.want)
			}
		})
	}
}
