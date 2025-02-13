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
