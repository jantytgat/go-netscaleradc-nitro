package nitro

import (
	"fmt"
	"testing"
	"time"
)

func TestSettings_GetTimeoutDuration(t *testing.T) {
	s := ConnectionSettings{
		Timeout: 10,
	}

	var tests = []struct {
		timeout int
		want    time.Duration
	}{
		{
			timeout: 10,
			want:    10 * time.Second,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.timeout)
		t.Run(testName, func(t *testing.T) {
			s.Timeout = tt.timeout
			result1, _ := s.GetTimeoutDuration()

			if result1 != tt.want {
				t.Errorf("GetTimeoutDuration(%d) = %d, expected %d", tt.timeout, result1, tt.want)
			}
		})
	}
}
