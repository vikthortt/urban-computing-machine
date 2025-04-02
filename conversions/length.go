package conversions

import (
	"encoding/json"
	"os"
)

type LengthConversion struct {
	FromUnit string  `json:"from_unit"`
	ToUnit   string  `json:"to_unit"`
	Factor   float64 `json:"factor"`
}

// LengthConversions is a list of length conversions loaded from a JSON file.
var LengthConversions []LengthConversion

// LoadLengthConversions loads length conversions from a JSON file.
// filePath is the path to the JSON file containing the conversions.
// Returns an error if there is any issue opening or decoding the file.
func LoadLengthConversions(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var conversions []LengthConversion
	if err := json.NewDecoder(file).Decode(&conversions); err != nil {
		return err
	}

	LengthConversions = conversions

	return nil
}
