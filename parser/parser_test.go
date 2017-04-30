package postal_test

import (
	"encoding/json"
	"reflect"
	"testing"

	parser "github.com/openvenues/gopostal/parser"
)

func init() {
	if err := parser.Setup(); err != nil {
		panic(err)
	}
}

func testParse(t *testing.T, address string, expectedOutput []parser.ParsedComponent, expectedJSON string) {
	parsedComponents := parser.ParseAddress(address)

	if len(parsedComponents) != len(expectedOutput) || !reflect.DeepEqual(parsedComponents, expectedOutput) {
		t.Error("parsed != expected: ", parsedComponents, "!=", expectedOutput)
	}

	// Test JSON marshaling.
	marshaledJSON, err := json.Marshal(parsedComponents)
	if err != nil {
		t.Error("JSON.marshal error: " + err.Error())
	}

	if string(marshaledJSON) != expectedJSON {
		t.Error("json != expected: ", string(marshaledJSON), "!=", expectedJSON)
	}

	// Test JSON unmarshaling.
	var unmarshaledComponents []parser.ParsedComponent
	if err := json.Unmarshal(marshaledJSON, &unmarshaledComponents); err != nil {
		t.Error("JSON.unmarshal error: " + err.Error())
	}

	if !reflect.DeepEqual(unmarshaledComponents, expectedOutput) {
		t.Error("unmarshaled != expected: ", unmarshaledComponents, "!=", expectedOutput)
	}
}

func TestParseUSAddress(t *testing.T) {
	t.Log("Testing US address")

	testParse(t, "781 Franklin Ave Crown Heights Brooklyn NYC NY 11216 USA",
		[]parser.ParsedComponent{
			{"house_number", "781"},
			{"road", "franklin ave"},
			{"suburb", "crown heights"},
			{"city_district", "brooklyn"},
			{"city", "nyc"},
			{"state", "ny"},
			{"postcode", "11216"},
			{"country", "usa"},
		},
		`[{"label":"house_number","value":"781"},{"label":"road","value":"franklin ave"},{"label":"suburb","value":"crown heights"},{"label":"city_district","value":"brooklyn"},{"label":"city","value":"nyc"},{"label":"state","value":"ny"},{"label":"postcode","value":"11216"},{"label":"country","value":"usa"}]`,
	)
}
