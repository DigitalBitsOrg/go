package frontier

import (
	"encoding/json"
	"testing"

	"github.com/digitalbitsorg/go/services/frontier/internal/resource"
	"github.com/digitalbitsorg/go/services/frontier/internal/test"
)

func TestRootAction(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	server := test.NewStaticMockServer(`{
			"info": {
				"network": "test",
				"build": "test-core",
				"protocol_version": 4
			}
		}`)
	defer server.Close()

	ht.App.horizonVersion = "test-frontier"
	ht.App.config.StellarCoreURL = server.URL

	w := ht.Get("/")
	if ht.Assert.Equal(200, w.Code) {
		var actual resource.Root
		err := json.Unmarshal(w.Body.Bytes(), &actual)
		ht.Require.NoError(err)
		ht.Assert.Equal("test-frontier", actual.HorizonVersion)
		ht.Assert.Equal("test-core", actual.StellarCoreVersion)
	}
}
