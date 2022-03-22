package address

import "testing"

func TestAddressType(t *testing.T) {
	mockAddress := "Sheppards run"

	expectedAddressType := "Run"
	currentAddressType := AddressType(mockAddress)

	if currentAddressType != expectedAddressType {
		// t.Error("Current type is different from expected")
		t.Errorf("Received %s, Expected %s", currentAddressType, expectedAddressType)
	}
}
