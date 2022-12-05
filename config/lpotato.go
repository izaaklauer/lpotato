package config

type Lpotato struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultLpotatoConfig returns default config values
func DefaultLpotatoConfig() Lpotato {
	return Lpotato{
	HelloWorldMessage:
		"hello from the default config",
	}
}
