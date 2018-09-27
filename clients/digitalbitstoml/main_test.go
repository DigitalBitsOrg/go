package digitalbitstoml

import "log"

// ExampleGetTOML gets the stellar.toml file for coins.asia
func ExampleClient_GetStellarToml() {
	_, err := DefaultClient.GetStellarToml("xdb.io")
	if err != nil {
		log.Fatal(err)
	}
}
