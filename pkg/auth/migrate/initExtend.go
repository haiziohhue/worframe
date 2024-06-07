package migrate

func (m *DBMigrate) initTend() error {
	return m.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
}
