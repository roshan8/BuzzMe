package config

import (
	"fmt"
	"os"

	psh "github.com/platformsh/config-reader-go/v2"
	libpq "github.com/platformsh/config-reader-go/v2/libpq"
)

const (
	// EnvDev const represents dev environment
	EnvDev = "dev"
	// EnvStaging const represents staging environment
	EnvStaging = "staging"
	// EnvProduction const represents production environment
	EnvProduction = "production"
)

// Env holds the current environment
var (
	Env          string
	Port         string
	DBDriver     string
	DBDataSource string
)

// Initialize ...
func Initialize() {

	// GetAllPlatformshEnv()
	//TODO: Uncomment the following once we stop using platformsh
	GetAllEnv()
}

func GetAllPlatformshEnv() {

	// Create a NewRuntimeConfig object to ease reading the Platform.sh environment variables.
	config, _ := psh.NewRuntimeConfig()
	credentials, _ := config.Credentials("postgresql")
	formatted, _ := libpq.FormattedCredentials(credentials)

	fmt.Println(formatted)

}

// GetAllEnv should get all the env configs required for the app.
func GetAllEnv() {
	// API Configs
	mustEnv("ENV", &Env, EnvDev)
	mustEnv("PORT", &Port, "8080")
	mustEnv("DB_DRIVER", &DBDriver, "postgres")
	mustEnv("DB_DATASOURCE", &DBDataSource,
		"user=main password=main dbname=main sslmode=disable host=postgresdatabase.internal")
}

// mustEnv get the env variable with the name 'key' and store it in 'value'
func mustEnv(key string, value *string, defaultVal string) {
	if *value = os.Getenv(key); *value == "" {
		*value = defaultVal
		fmt.Printf("%s env variable not set, using default value.\n", key)
	}
}
