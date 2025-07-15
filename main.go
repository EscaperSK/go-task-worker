package main

import (
	"flag"
	"tasks/database"
	"tasks/server"
	"tasks/worker"
)

var dbName = flag.String("db-name", "tasks_db", "database name")
var dbUser = flag.String("db-user", "root", "database user name")
var dbPass = flag.String("db-pass", "", "database password")
var dbTasks = flag.Int("db-tasks", 100, "Nomber of tasks to create")

var poolSize = flag.Int("pool-size", 10, "tasks pool size")

var port = flag.String("port", "8080", "server port number")

func main() {
	flag.Parse()

	database.Prepare(*dbName, *dbUser, *dbPass, *dbTasks)

	worker.Run(*poolSize)

	server.Run(*port)
}
