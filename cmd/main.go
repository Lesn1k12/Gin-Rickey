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
	os.Mkdir(name+"/tests", 0755)

	//creating files
	createFile(name+"/main.go", generateMainGo())
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
