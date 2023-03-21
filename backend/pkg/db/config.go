package db

type Config struct {
	Addr                         string `env:"DATABASE_ADDR" envDefault:"localhost:26257"`
	User                         string `env:"DATABASE_USER" envDefault:"root"`
	Password                     string `env:"DATABASE_PASSWORD" envDefault:"root"`
	Database                     string `env:"DATABASE_NAME" envDefault:"banking"`
	MaxOpenConnections           int    `env:"MAX_OPEN_CONNECTIONS" envDefault:"20"`
	MaxIdleConnections           int    `env:"MAX_IDLE_CONNECTIONS" envDefault:"20"`
	MaxConnectionLifetimeMinutes int    `env:"MAX_CONNECTION_LIFETIME" envDefault:"30"`
	DriverName                   string `env:"DRIVER_NAME" envDefault:"postgres"`
}
