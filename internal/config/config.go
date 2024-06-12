package config

type (
	Config struct {
		App        App
		ClientURLs ClientURLs
	}

	App struct {
		ServerPort string
	}

	ClientURLs struct {
		EthRpcClientURL string
	}
)

// Cfg is a configuration instance.
var Cfg Config

// Load will load hardcoded default config into memory
// For production use, we should have custom config file
// with support for reading from env values (modules like viper can be good choice)
func Load() {
	Cfg = Config{
		App: App{
			ServerPort: ":8080",
		},
		ClientURLs: ClientURLs{
			EthRpcClientURL: "https://cloudflare-eth.com",
		},
	}
}
