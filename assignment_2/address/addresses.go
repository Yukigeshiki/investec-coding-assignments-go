package address

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Addresses []Address

var addrs *Addresses

// PrettyPrintAddresses is the solution to b.
func (addrs *Addresses) PrettyPrintAddresses() {
	for _, addr := range *addrs {
		fmt.Println(addr.GetPrettyPrinting())
	}
}

// ValidateAddresses is the solution to e.
func (addrs *Addresses) ValidateAddresses() []string {
	var errStrings []string

	for _, addr := range *addrs {
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
	var (
		err       error
		jsonFile  *os.File
		byteValue []byte
		absPath   string
	)

	if absPath, err = filepath.Abs("../addresses.json"); err != nil {
		panic(err)
	}
	if jsonFile, err = os.Open(absPath); err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	if byteValue, err = io.ReadAll(jsonFile); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(byteValue, &addrs); err != nil {
		panic(err)
	}
}
