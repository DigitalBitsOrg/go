package stellar

import (
	"sync"

	"github.com/digitalbitsorg/go/clients/frontier"
	"github.com/digitalbitsorg/go/support/log"
)

// AccountConfigurator is responsible for configuring new DigitalBits accounts that
// participate in ICO.
type AccountConfigurator struct {
	Horizon           frontier.ClientInterface `inject:""`
	NetworkPassphrase string
	IssuerPublicKey   string
	SignerSecretKey   string
	NeedsAuthorize    bool
	TokenAssetCode    string
	StartingBalance   string
	OnAccountCreated  func(destination string)
	OnAccountCredited func(destination string, assetCode string, amount string)

	signerPublicKey      string
	sequence             uint64
	sequenceMutex        sync.Mutex
	processingCount      int
	processingCountMutex sync.Mutex
	log                  *log.Entry
}
