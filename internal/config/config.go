package config

// Config is a main config structure.
type Config struct {
	TelegramBot telegramBot `toml:"telegram_bot"`
}

type telegramBot struct {
	Token string `toml:"token"`
}
