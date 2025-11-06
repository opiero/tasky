package database

type DatabaseInterface interface {
	AddTask() error
	ListTasks() ([][]string, error)
	CompleteTask() error
	DeleteTask() error
	PrintDatabase(int) error
}
