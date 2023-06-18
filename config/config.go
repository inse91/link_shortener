package config

import "os"

const (
	grpcPortKey     = "GRPC_PORT"
	defaultGrpcPort = "10000"

	httpPortKey     = "HTTP_PORT"
	defaultHttpPort = "10010"

	dbPortKey     = "DB_CONNECTION"
	defaultDbConn = ""
)

type Config struct {
	GrpcPort string
	HttpPort string
	DataBase string
}

func Get() Config {
	return Config{
		GrpcPort: getEnvOrSetToDefault(grpcPortKey, defaultGrpcPort),
		HttpPort: getEnvOrSetToDefault(httpPortKey, defaultHttpPort),
		DataBase: getEnvOrSetToDefault(dbPortKey, defaultDbConn),
	}
}

func getEnvOrSetToDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
