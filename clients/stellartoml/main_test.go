package stellartoml

import "log"

// ExampleGetTOML gets the stellar.toml file for coins.asia
func ExampleClient_GetStellarToml() {
	_, err := DefaultClient.GetStellarToml("digitalbits.io")
	if err != nil {
		log.Fatal(err)
	}
}
