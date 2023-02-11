package hclient

import "testing"

func Test_getQueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   map[string]string
		expected string
	}{
		{
			name:     "Test empty params",
			params:   map[string]string{},
			expected: "",
		},
		{
			name: "Test one-element map",
			params: map[string]string{
				"key": "myKey",
			},
			expected: "key=myKey",
		},
		{
			name: "Test 2-element map",
			params: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expected: "key1=value1&key2=value2",
		},
		{
			name: "Test empty value gets discarded",
			params: map[string]string{
				"key1": "value1",
				"key2": "",
			},
			expected: "key1=value1",
		},
		{
			name: "Test all values are empty",
			params: map[string]string{
				"key1": "",
				"key2": "",
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getQueryString(test.params)
			if actual != test.expected {
				t.Errorf("failed to get querystring, expected: %s, current: %s ", test.expected, actual)
			}
		})
	}
}
