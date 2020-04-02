package keys

type KeysConfig struct {
	GRPCServiceEnabled      bool
	AllowBadFilePermissions bool
	RemoteAddress           string
	KeysDirectory           string
}

func DefaultKeysConfig() *KeysConfig {
	return &KeysConfig{
		// Default Monax keys port
		GRPCServiceEnabled:      true,
		AllowBadFilePermissions: true, // kingwel, make windows happy
		RemoteAddress:           "",
		KeysDirectory:           DefaultKeysDir,
	}
}
