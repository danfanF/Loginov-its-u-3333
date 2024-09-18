package postgresqlstorage

import "github.com/akayo16/jorik_loginov_first_task/config"

func InitDatabase(storageConfig config.PostgresqlConfig) {
	initPostgresql(storageConfig)
	initAllBootstrapper()
}
