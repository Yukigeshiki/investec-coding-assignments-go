package address

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const jsonFilePath = "../addresses.json"

type Addresses []Address

var addrs *Addresses

// PrettyPrintAddresses is the solution to b.
func (a *Addresses) PrettyPrintAddresses() {
	for _, addr := range *a {
		fmt.Println(addr.GetPrettyPrinting())
	}
}

// ValidateAddresses is the solution to e.
func (a *Addresses) ValidateAddresses() []string {
	var errStrings []string

	for _, addr := range *a {
		vErrs := addr.Validate()
		if len(vErrs) != 0 {
			errStrings = append(
				errStrings,
				fmt.Sprintf("Address for ID: %s is invalid. Validation errors: %v", addr.ID, vErrs),
			)
		}
	}

	return errStrings
}

// init imports a json file of addresses.
func init() {
	absPath, err := filepath.Abs(jsonFilePath)
	if err != nil {
		panic(err)
	}
	jsonFile, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(byteValue, &addrs); err != nil {
		panic(err)
	}
}
