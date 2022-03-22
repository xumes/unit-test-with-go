package address

import "strings"

// AddressType checks if the address has a valid type and returns it
func AddressType(address string) string {
	validTypes := []string{"street", "avenue", "lane", "highway", "run", "drive", "crescent"}

	addressInLowerCase := strings.ToLower(address)
	addressSlice := strings.Split(addressInLowerCase, " ")

	lastAddressWord := addressSlice[len(addressSlice)-1]

	hasValidType := false

	for _, addressType := range validTypes {
		if addressType == lastAddressWord {
			hasValidType = true
		}
	}

	if hasValidType {
		return strings.Title(lastAddressWord)
	}

	return "Invalid address type"
}
