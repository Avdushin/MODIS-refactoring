package logger

/* Пакет Logger
* Содержит логику работы с логгером
* Пример использования:
* logger.INFO("Инфа 100", args...)
* logger.WARNING("Юрааааа, мы всё про...теряли", args...)
* logger.ERROR("Не получилось восстановить нервную систему", args...)
* logger.DB("Создана таблица пользователей", args...)
 */

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func InitLogger() {
	/* Уровни логгера:
	* INFO - информационные сообщения
	* WARNING - ошибки, требующие внимания
	* ERROR - критические ошибки
	* DB - лог базы данных
	 */
	infoLogger = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
}

func logMessage(prefix string, message ...interface{}) {
	//? Формируем строку с датой, временем и префиксом
	logString := fmt.Sprintf("%s %s ", time.Now().Format("2006-01-02 15:04:05"), prefix)
	//? Переданные сообщения
	logString += fmt.Sprintln(message...)
	//? Выводим лог
	infoLogger.Print(logString)
}

// @ Префикс "INFO"
// ? информационные сообщения
func INFO(message ...interface{}) {
	logMessage("[INFO]", message...)
}

// @ Префикс "WARNING"
// ? ошибки, требующие внимания
func WARNING(message ...interface{}) {
	logMessage("[WARNING]", message...)
}

// @ Префикс "ERROR"
// ? критические ошибки
func ERROR(message ...interface{}) {
	logMessage("[ERROR]", message...)
}

// @ Префикс "DB"
// ? лог базы данных
func DB(message ...interface{}) {
	logMessage("[DB]", message...)
}

// @ Префикс "FATAL" - фатальные ошибки
// ? завершают работу приложения при возникновении ошибки
func FATAL(message ...interface{}) {
	errorLogger.Fatal(message...)
}
