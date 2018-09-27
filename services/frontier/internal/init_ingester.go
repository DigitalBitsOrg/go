package frontier

import (
	"log"

	"github.com/digitalbitsorg/go/services/frontier/internal/ingest"
)

func initIngester(app *App) {
	if !app.config.Ingest {
		return
	}

	if app.networkPassphrase == "" {
		log.Fatal("Cannot start ingestion without network passphrase.  Please confirm connectivity with digitalbits-core.")
	}

	app.ingester = ingest.New(
		app.networkPassphrase,
		app.config.StellarCoreURL,
		app.CoreSession(nil),
		app.HorizonSession(nil),
	)

	app.ingester.SkipCursorUpdate = app.config.SkipCursorUpdate
	app.ingester.HistoryRetentionCount = app.config.HistoryRetentionCount
}

func init() {
	appInit.Add("ingester", initIngester, "app-context", "log", "horizon-db", "core-db", "stellarCoreInfo")
}
