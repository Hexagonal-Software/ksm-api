package secrets

import (
	"errors"

	ksm "github.com/keeper-security/secrets-manager-go/core"
	"hexagonal.software/ksm-api/internal/config"
)

var (
	KsmEngine        *ksm.SecretsManager
	ErrInitKsmEngine = errors.New("failed to initialize KSM engine")
)

func BootstrapKsmEngine(c *config.KeeperVault) error {
	eng, err := NewKsmEngine(c.KsmConfig)

	if err != nil {
		return err
	}

	KsmEngine = eng

	return nil
}

func NewKsmEngine(ksmConfig string) (*ksm.SecretsManager, error) {
	clientOptions := &ksm.ClientOptions{
		Config: ksm.NewMemoryKeyValueStorage(ksmConfig),
	}
	sm := ksm.NewSecretsManager(clientOptions)

	if sm == nil {
		return nil, ErrInitKsmEngine
	}

	return sm, nil
}

func GetKsmEngine(c ...string) *ksm.SecretsManager {
	if len(c) > 0 {
		eng, err := NewKsmEngine(c[0])

		if err != nil {
			return nil
		}

		return eng
	}

	return KsmEngine
}
