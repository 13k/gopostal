package postal_test

import (
	"testing"

	expand "github.com/openvenues/gopostal/expand"
)

func init() {
	if err := expand.Setup(); err != nil {
		panic(err)
	}
}

func testExpansionInOutput(t *testing.T, address string, output string, expansions []string) {
	for i := 0; i < len(expansions); i++ {
		if expansions[i] == output {
			return
		}
	}

	t.Error("expansion", output, "not found in expansions for address", address)
}

func testExpansion(t *testing.T, address string, output string) {
	expansions := expand.ExpandAddress(address)
	testExpansionInOutput(t, address, output, expansions)
}

func testExpansionWithOptions(t *testing.T, address string, output string, options expand.ExpandOptions) {
	expansions := expand.ExpandAddressOptions(address, options)
	testExpansionInOutput(t, address, output, expansions)
}

func TestEnglishExpansions(t *testing.T) {
	testExpansion(t, "123 Main St", "123 main street")

	englishOptions := expand.DefaultExpansionOptions
	englishOptions.Languages = []string{"en"}

	testExpansionWithOptions(t, "30 West Twenty-sixth St Fl No. 7", "30 west 26th street floor number 7", englishOptions)
	testExpansionWithOptions(t, "Thirty W 26th St Fl #7", "30 west 26th street floor number 7", englishOptions)
}

func TestMultilingualExpansions(t *testing.T) {
	multilingualOptions := expand.DefaultExpansionOptions
	multilingualOptions.Languages = []string{"en", "fr", "de"}

	testExpansionWithOptions(t, "st", "sankt", multilingualOptions)
	testExpansionWithOptions(t, "st", "saint", multilingualOptions)
}

func TestNonASCIIExpansions(t *testing.T) {
	testExpansion(t, "Friedrichstraße 128, Berlin, Germany", "friedrich straße 128 berlin germany")
}
