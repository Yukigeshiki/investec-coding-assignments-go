package address

import (
	"fmt"
	"regexp"
)

const notAvailable string = "Not available"

type CodeAndName struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Type CodeAndName

type ProvinceOrState CodeAndName

type Country CodeAndName

type LineDetail struct {
	Line1 string `json:"line1,omitempty"`
	Line2 string `json:"line2,omitempty"`
}

type Address struct {
	ID               string `json:"id"`
	Type             `json:"type"`
	LineDetail       `json:"addressLineDetail,omitempty"`
	ProvinceOrState  `json:"provinceOrState,omitempty"`
	Country          `json:"country,omitempty"`
	CityOrTown       string `json:"cityOrTown,omitempty"`
	PostalCode       string `json:"postalCode,omitempty"`
	SuburbOrDistrict string `json:"suburbOrDistrict,omitempty"`
	LastUpdated      string `json:"lastUpdated"`
}

// GetPrettyPrinting is the solution to a.
func (a *Address) GetPrettyPrinting() string {
	return fmt.Sprintf("%s: ", a.Type.Name) +
		fmt.Sprintf("%s ", getPrettyPrintingLineDetail(a.LineDetail)) +
		fmt.Sprintf("- %s ", getPrettyPrintingString(a.CityOrTown)) +
		fmt.Sprintf("- %s ", getPrettyPrintingString(a.ProvinceOrState.Name)) +
		fmt.Sprintf("- %s ", getPrettyPrintingString(a.PostalCode)) +
		fmt.Sprintf("- %s", getPrettyPrintingString(a.Country.Name))
}

// IsValid is the solution to d.
func (a *Address) IsValid() bool {
	return isValidPostalCode(a.PostalCode) &&
		isValidCountry(a.Country) &&
		isValidLineDetail(a.LineDetail) &&
		a.hasValidProvince()
}

// Validate checks whether the needed address fields are valid. If a field is not valid an error message is added
// to a string slice and the slice is returned.
func (a *Address) Validate() (vErrs []string) {

	if !isValidPostalCode(a.PostalCode) {
		vErrs = append(vErrs, "You must include a valid postal code")
	}
	if !isValidCountry(a.Country) {
		vErrs = append(vErrs, "You must include a country")
	}
	if !isValidLineDetail(a.LineDetail) {
		vErrs = append(vErrs, "You must include valid address details (line 1 and/or 2 must be filled in)")
	}
	if !a.hasValidProvince() {
		vErrs = append(vErrs, "You must include a province if your country is ZA")
	}

	return
}

// hasValidProvince checks whether a province is included when the country is ZA.
func (a *Address) hasValidProvince() bool {
	if a.Country.Code == "ZA" {
		return a.ProvinceOrState.Name != ""
	}
	return true
}

// getPrettyPrintingString returns either the given string, or if it's empty, "Not available".
func getPrettyPrintingString(s string) string {
	if s != "" {
		return s
	}
	return notAvailable
}

// getPrettyPrintingLineDetail returns either the address line details properly formatted or "Not available".
func getPrettyPrintingLineDetail(ld LineDetail) string {
	// check if lines are missing
	ml1 := ld.Line1 == ""
	ml2 := ld.Line2 == ""

	switch {
	case ml1 && ml2:
		return notAvailable
	case ml1:
		return ld.Line2
	case ml2:
		return ld.Line1
	default:
		return fmt.Sprintf("%s, %s", ld.Line1, ld.Line2)
	}
}

// isValidPostalCode checks whether the address has a valid postal code.
func isValidPostalCode(code string) (v bool) {
	v, _ = regexp.MatchString("[0-9]+", code)
	return
}

// isValidLineDetail checks whether the address line details are valid.
func isValidLineDetail(ld LineDetail) bool {
	return ld.Line1 != "" || ld.Line2 != ""
}

// isValidCountry checks if the country is valid, i.e. it has a non-empty name.
func isValidCountry(c Country) bool {
	return c.Name != ""
}
