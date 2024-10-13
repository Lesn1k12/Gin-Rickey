package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "NewApp",
	Short: "NewApp is a new app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating a new project...")
		createProject(args[0])
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createProject(name string) {
	fmt.Println("Creating project", name)
	//creating directories
	os.Mkdir(name, 0755)
	os.Mkdir(name+"/cmd", 0755)
	os.Mkdir(name+"/internal", 0755)
	os.Mkdir(name+"/internal/router", 0755)
	os.Mkdir(name+"/internal/controllers", 0755)
	os.Mkdir(name+"/internal/db", 0755)
	os.Mkdir(name+"/tests", 0755)

	//creating files
	createFile(name+"/cmd/main.go", generateMainGo())
	createFile(name+"/.env", "")
	createFile(name+"/Makefile", "")
	createFile(name+"/tests/test.rest", generateTestRest())
	createFile(name+"/internal/router/router.go", generateRouterGo())
	createFile(name+"/internal/controllers/controllers.go", generateControllersGo())
	createFile(name+"/internal/db/db.go", generateDbGo())
}

func createFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(content)
}

func generateMainGo() string {
	return `
package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	gin.Run(":8080")
}
`
}

func generateTestRest() string {
	return `
Get http://localhost:8080/test
`
}

func generateRouterGo() string {
	return `
package router

import (

)

`
}

func generateControllersGo() string {
	return `
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, "Test completed!")
}
`
}

func generateDbGo() string {
	return `
package db

import (

)

func Connect() {

}
`
}
