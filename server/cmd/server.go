package main

import (
	"fmt"
	"furrj/todo_2_remake/internal/routing/consts"
	"furrj/todo_2_remake/internal/routing/routes"
	"furrj/todo_2_remake/internal/routing/routes/todos"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/mandrigin/gin-spa/spa"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// ENV CONFIG
	if os.Getenv("MODE") != "PROD" {
		if err := godotenv.Load(".env"); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("No env variable PORT")
		os.Exit(1)
	}

	// DB
	// db := dbHandler.InitDBHandler(os.Getenv("DB_URL"))
	// defer db.Conn.Close()

	// ROUTING
	router := gin.Default()
	if os.Getenv("MODE") == "DEV" {
		fmt.Println("**DEV MODE DETECTED, ENABLING CORS**")
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
		// config.AllowHeaders = []string{consts.HeaderTypeAuthorization, "Content-Type"}
		router.Use(cors.New(config))
	}

	router.POST(consts.RouteUrlRegister, routes.Register())
	router.POST(consts.RouteUrlLogin, routes.Login())
	router.POST(consts.RouteUrlGetTodos, todos.GetTodos())
	router.POST(consts.RouteUrlAddTodo, todos.AddTodo())
	router.PUT(consts.RouteUrlUpdateTodo, todos.UpdateTodo())
	router.DELETE(consts.RouteUrlDeleteTodo, todos.DeleteTodo())

	router.Use(spa.Middleware("/", "client"))

	log.Panic(router.Run(PORT))
}
