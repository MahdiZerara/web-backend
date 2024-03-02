package config

// Shared configs
type SharedConfig struct {
	// This config can be extended when we need configs that
	// are shared accross multiple sections of the code
}

func newSharedConfig() (*SharedConfig, error) {
	return &SharedConfig{}, nil
}
