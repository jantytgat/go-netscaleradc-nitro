package nitro

import (
	"fmt"
	"testing"
)

// func TestNitroRequestParams_GetResourceName(t *testing.T) {
//	var tests = []struct {
//		r    NitroResourceType
//		want string
//	}{
//		{UnknownResource, ""},
//		{SystemBackup, "systembackup"},
//	}
//
//	for _, tt := range tests {
//		testName := fmt.Sprintf("%s", tt.r.GetTypeName())
//		t.Run(testName, func(t *testing.T) {
//			p := Request{ResourceType: tt.r}
//			result := p.GetTypeName()
//			if result != tt.want {
//				t.Errorf("result: %s, expected: %s", result, tt.want)
//			}
//		})
//	}
// }

// func TestNitroRequestParams_GetMethod(t *testing.T) {
//	var tests = []struct {
//		params Request
//		want   string
//	}{
//		{Request{Method: NitroUnknownMethod}, ""},
//		{Request{Method: NitroGetMethod}, "GET"},
//		{Request{Method: NitroPostMethod}, "POST"},
//		{Request{Method: 100}, ""},
//	}
//
//	for _, tt := range tests {
//		testName := fmt.Sprintf("%s", tt.params.Method.String())
//		t.Run(testName, func(t *testing.T) {
//			result := tt.params.GetMethod()
//			if result != tt.want {
//				t.Errorf("result: %s, expected: %s", result, tt.want)
//			}
//		})
//	}
// }

// func TestNitroRequestParams_GetUrl(t *testing.T) {
//	var tests = []struct {
//		params Request
//		want   string
//	}{
//		{Request{ResourceType: UnknownResource, RequestType: NitroUnknownRequest}, ""},
//		{Request{ResourceType: UnknownResource, RequestType: 300}, ""},
//		{Request{ResourceType: UnknownResource, RequestType: NitroConfigRequest}, "/nitro/v1/config/"},
//		{Request{ResourceType: SystemBackup, RequestType: NitroConfigRequest}, "/nitro/v1/config/systembackup"},
//		{Request{ResourceType: UnknownResource, RequestType: NitroStatRequest}, "/nitro/v1/stat/"},
//		{Request{ResourceType: SystemBackup, RequestType: NitroStatRequest}, "/nitro/v1/stat/systembackup"},
//		{Request{
//			ResourceType: SystemBackup,
//			RequestType:  NitroConfigRequest,
//			Arguments:    map[string]string{"key1": "value1"}},
//			"/nitro/v1/config/systembackup?args=key1:value1"}, {Request{
//			ResourceType: SystemBackup,
//			RequestType:  NitroConfigRequest,
//			Arguments:    map[string]string{"fileLocation": "/var/ns_sys_backup"}},
//			"/nitro/v1/config/systembackup?args=fileLocation:%2Fvar%2Fns_sys_backup"},
//		{Request{
//			ResourceType: SystemBackup,
//			RequestType:  NitroConfigRequest,
//			Arguments:    map[string]string{"key1": "value1", "key2": "value2"}},
//			"/nitro/v1/config/systembackup?args=key1:value1,key2:value2"},
//		{Request{
//			ResourceType: SystemBackup,
//			RequestType:  NitroConfigRequest,
//			Arguments:    map[string]string{"key3": "value3", "key1": "value1", "key2": "value2"}},
//			"/nitro/v1/config/systembackup?args=key1:value1,key2:value2,key3:value3"},
//		{Request{
//			ResourceType: SystemBackup,
//			RequestType:  NitroConfigRequest,
//			Arguments:    map[string]string{"key3": "value3", "key1": "value1", "key2": "value2"},
//			Filter:       map[string]string{"key1": "value1"}},
//			"/nitro/v1/config/systembackup?args=key1:value1,key2:value2,key3:value3&filter=key1:value1"},
//	}
//
//	for _, tt := range tests {
//		testName := fmt.Sprintf("%s", tt.params.GetTypeName())
//		t.Run(testName, func(t *testing.T) {
//			result := tt.params.GetUrlPathAndQuery()
//			if result != tt.want {
//				t.Errorf("result: %s, expected: %s", result, tt.want)
//			}
//		})
//	}
// }

// func TestGetNitroUrlPath(t *testing.T) {
//	var tests = []struct {
//		params Request
//		want   string
//	}{
//		{Request{ResourceType: UnknownResource, RequestType: NitroUnknownRequest}, ""},
//		{Request{ResourceType: UnknownResource, RequestType: NitroConfigRequest}, "/nitro/v1/config/"},
//		{Request{ResourceType: UnknownResource, RequestType: NitroStatRequest}, "/nitro/v1/stat/"},
//		{Request{ResourceType: UnknownResource, RequestType: 300}, ""},
//		{Request{ResourceType: SystemBackup, RequestType: NitroUnknownRequest}, ""},
//		{Request{ResourceType: SystemBackup, RequestType: NitroConfigRequest}, "/nitro/v1/config/systembackup"},
//		{Request{ResourceType: SystemBackup, RequestType: NitroStatRequest}, "/nitro/v1/stat/systembackup"},
//		{Request{ResourceType: SystemBackup, RequestType: 300}, ""},
//	}
//
//	for _, tt := range tests {
//		testName := fmt.Sprintf("%s", tt.params.RequestType.String())
//		t.Run(testName, func(t *testing.T) {
//			result := tt.params.GetUrlPath()
//			if result != tt.want {
//				t.Errorf("result: %s, expected: %s", result, tt.want)
//			}
//		})
//	}
// }

func TestGetUrlQueryStringSeparator(t *testing.T) {
	var tests = []struct {
		length int
		want   string
	}{
		{0, "?"},
		{1, "&"},
		{100, "&"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.length)
		t.Run(testName, func(t *testing.T) {
			result := getUrlQueryStringSeparator(tt.length)
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestGetUrlQueryMapStringSeparator(t *testing.T) {
	var tests = []struct {
		index     int
		lastIndex int
		want      string
	}{
		{0, 10, ","},
		{10, 10, ""},
		{100, 10, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.index, tt.lastIndex)
		t.Run(testName, func(t *testing.T) {
			result := getUrlQueryMapStringSeparator(tt.index, tt.lastIndex)
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestGetUrlQueryMapEntriesAsString(t *testing.T) {
	var tests = []struct {
		entries       map[string]string
		wrapCharacter string
		want          string
	}{
		{entries: map[string]string{"key1": "value1"}, wrapCharacter: "", want: "key1:value1"},
		{entries: map[string]string{"key1": "value1", "key2": "value2"}, wrapCharacter: "", want: "key1:value1,key2:value2"},
		{entries: map[string]string{"key3": "value3", "key1": "value1", "key2": "value2"}, wrapCharacter: "", want: "key1:value1,key2:value2,key3:value3"},
		{entries: map[string]string{"key1": "value1"}, wrapCharacter: "/", want: "key1:%2Fvalue1%2F"},
		{entries: map[string]string{"key1": "value1", "key2": "value2"}, wrapCharacter: "/", want: "key1:%2Fvalue1%2F,key2:%2Fvalue2%2F"},
		{entries: map[string]string{"key3": "value3", "key1": "value1", "key2": "value2"}, wrapCharacter: "/", want: "key1:%2Fvalue1%2F,key2:%2Fvalue2%2F,key3:%2Fvalue3%2F"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.want)
		t.Run(testName, func(t *testing.T) {
			result := getUrlQueryMapEntriesAsString(tt.entries, tt.wrapCharacter)
			if result != tt.want {
				t.Errorf("result: %s - expected: %s", result, tt.want)
			}
		})
	}
}
