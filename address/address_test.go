package address

import "testing"

type testCase struct {
	addressToTest  string
	expectedResult string
}

func TestAddressType(t *testing.T) {
	testCases := []testCase{
		{"Sheppards Run", "Run"},
		{"SHEPPARDS RUN", "Run"},
		{"something else run", "Run"},
		{"Washmill Lake Drive", "Drive"},
		{"another drive", "Drive"},
		{"UPPERCASE DRIVE", "Drive"},
		{"Sam Crescent", "Crescent"},
		{"another crescent", "Crescent"},
		{"Main Street", "Street"},
		{"main street", "Street"},
		{"Xumes Lane", "Lane"},
		{"another lane", "Lane"},
		{"Bedford Highway", "Highway"},
		{"another highway", "Highway"},
		{"no name", "Invalid address type"},
		{"45 Watter Level", "Invalid address type"},
		{"", "Invalid address type"},
	}

	for _, scenario := range testCases {
		currentAddressType := AddressType(scenario.addressToTest)
		if currentAddressType != scenario.expectedResult {
			t.Errorf("Expected %s but received %s", scenario.expectedResult, currentAddressType)
		}
	}
}
