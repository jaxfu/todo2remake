package main

import (
	"fmt"
	"furrj/todo_2_remake/internal/config"
	"furrj/todo_2_remake/internal/routing/consts/routeURL"
	"furrj/todo_2_remake/internal/routing/routes"
	"furrj/todo_2_remake/internal/routing/routes/todos"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/mandrigin/gin-spa/spa"

	"github.com/gin-gonic/gin"
)

func main() {
	envVars, err := config.GetEnv(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting environment vars: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", envVars)

	// DB
	// db := dbHandler.InitDBHandler(os.Getenv("DB_URL"))
	// defer db.Conn.Close()

	// ROUTING
	router := gin.Default()
	// enable CORS to make requests from localhost (DEV ONLY)
	if os.Getenv("MODE") == "DEV" {
		fmt.Println("**DEV MODE DETECTED, ENABLING CORS**")
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
		// config.AllowHeaders = []string{consts.HeaderTypeAuthorization, "Content-Type"}
		router.Use(cors.New(config))
	}

	// register routes
	router.POST(routeURL.Register, routes.Register())
	router.POST(routeURL.Login, routes.Login())
	router.POST(routeURL.GetTodos, todos.GetTodos())
	router.POST(routeURL.AddTodo, todos.AddTodo())
	router.PUT(routeURL.UpdateTodo, todos.UpdateTodo())
	router.DELETE(routeURL.DeleteTodo, todos.DeleteTodo())

	// server react bundle
	router.Use(spa.Middleware("/", "client"))

	// listen on port
	log.Panic(router.Run(fmt.Sprintf(":%s", envVars.PORT)))
}
