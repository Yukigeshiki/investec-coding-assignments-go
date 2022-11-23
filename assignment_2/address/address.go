package address

import (
	"fmt"
	"regexp"
)

const NotAvailable string = "Not available"

type Address struct {
	ID               string          `json:"id"`
	AddrType         addrType        `json:"type"`
	AddrLineDetail   addrLineDetail  `json:"addressLineDetail,omitempty"`
	ProvinceOrState  provinceOrState `json:"provinceOrState,omitempty"`
	Country          country         `json:"country,omitempty"`
	CityOrTown       string          `json:"cityOrTown,omitempty"`
	PostalCode       string          `json:"postalCode,omitempty"`
	SuburbOrDistrict string          `json:"suburbOrDistrict,omitempty"`
	LastUpdated      string          `json:"lastUpdated"`
}

type addrType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type addrLineDetail struct {
	Line1 string `json:"line1,omitempty"`
	Line2 string `json:"line2,omitempty"`
}

type provinceOrState struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetPrettyPrinting is the solution to a.
func (addr *Address) GetPrettyPrinting() string {
	ct, ps, pc, c := addr.getPrettyPrintingValues()

	return fmt.Sprintf("%s: ", addr.AddrType.Name) +
		fmt.Sprintf("%s ", getPrettyPrintingLineDetail(addr.AddrLineDetail)) +
		fmt.Sprintf("- %s ", ct) +
		fmt.Sprintf("- %s ", ps) +
		fmt.Sprintf("- %s ", pc) +
		fmt.Sprintf("- %s", c)
}

// IsValid is the solution to d.
func (addr *Address) IsValid() bool {
	checkPc, checkC, checkLd, checkP := addr.getAddressChecks()
	return checkPc && checkC && checkLd && checkP
}

// Validate checks whether the needed address fields are valid. If a field is not valid an error message is added
// to a string slice and the slice is returned.
func (addr *Address) Validate() []string {
	var vErrs []string

	checkPc, checkC, checkLd, checkP := addr.getAddressChecks()

	if !checkPc {
		vErrs = append(vErrs, "You must include a valid postal code")
	}
	if !checkC {
		vErrs = append(vErrs, "You must include a country")
	}
	if !checkLd {
		vErrs = append(vErrs, "You must include valid address details (line 1 and/or 2 must be filled in)")
	}
	if !checkP {
		vErrs = append(vErrs, "You must include a province if your country is ZA")
	}

	return vErrs
}

// getAddressChecks returns a tuple of boolean values for address field validation checks.
func (addr *Address) getAddressChecks() (bool, bool, bool, bool) {
	return isValidPostalCode(addr.PostalCode),
		isValidCountry(addr.Country),
		isValidLineDetail(addr.AddrLineDetail),
		addr.hasValidProvince()
}

// hasValidProvince checks whether a province is included when the country is ZA.
func (addr *Address) hasValidProvince() bool {
	if addr.Country.Code == "ZA" {
		return addr.ProvinceOrState.Name != ""
	}
	return true
}

// getPrettyPrintingValues returns a tuple of strings used to pretty print the address.
// If a value does not exist, it is returned as "Not available".
func (addr *Address) getPrettyPrintingValues() (string, string, string, string) {

	ps := getPrettyPrintingString(addr.ProvinceOrState.Name)
	c := getPrettyPrintingString(addr.Country.Name)
	ct := getPrettyPrintingString(addr.CityOrTown)
	pc := getPrettyPrintingString(addr.PostalCode)

	return ct, ps, pc, c
}

// getPrettyPrintingString returns either the given string, or if it's empty, "Not available".
func getPrettyPrintingString(s string) string {
	if s != "" {
		return s
	}
	return NotAvailable
}

// getPrettyPrintingLineDetail returns either the address line details or "Not available".
func getPrettyPrintingLineDetail(ld addrLineDetail) string {
	var (
		ml1 bool
		ml2 bool
	)

	// check if lines are missing
	if ld.Line1 == "" {
		ml1 = true
	}
	if ld.Line2 == "" {
		ml2 = true
	}

	switch {
	case ml1 && ml2:
		return NotAvailable
	case ml1:
		return ld.Line2
	case ml2:
		return ld.Line1
	default:
		return fmt.Sprintf("%s, %s", ld.Line1, ld.Line2)
	}
}

// isValidPostalCode checks whether the address has a valid postal code.
func isValidPostalCode(code string) bool {
	v, _ := regexp.MatchString("[0-9]+?", code)
	return v
}

// isValidLineDetail checks whether the address line details are valid.
func isValidLineDetail(ld addrLineDetail) bool {
	return ld.Line1 != "" || ld.Line2 != ""
}

// isValidCountry checks if the country is valid, i.e. it has a non-empty name.
func isValidCountry(c country) bool {
	return c.Name != ""
}