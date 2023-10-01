package database

import "time"

//@ Бэкап БД
//? Аргументы: d - duration, задержка в часах
func RunDatabaseBackup(d time.Duration) error {
	for range time.Tick(d * time.Hour) {
		BackupDB()
		SaveStorage()
	}
	return nil
}
