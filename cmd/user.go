package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"agenda/entity"
)

func checkEmpty(key, value string) {
	if value == "" {
		printError(key + " can't be empty!\n")
	}
}

func printError(error string) {
	fmt.Fprint(os.Stderr, error)
	os.Exit(1)
}

var registerCmd = &cobra.Command{
	Use: "register",
	Short: "Register user",
	Long: "Register a new user",
	Run: func(com *cobra.Command, args []string){
		username, _ := com.Flags().GetString("user")
		checkEmpty("username", username)

		password, _ := com.Flags().GetString("password")
		checkEmpty("password", password)

		mail, _ := com.Flags().GetString("mail")
		checkEmpty("mail", mail)

		phone, _ := com.Flags().GetString("phone")
		checkEmpty("phone", phone)

		if err := entity.Register(username,password,mail,phone); err !=nil {

		} else{
			fmt.Println("register success")
		}	
	},
}

func init() {
	registerCmd.Flags().StringP("user", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "password")
	registerCmd.Flags().StringP("mail", "m", "", "email")
	registerCmd.Flags().StringP("phone", "t", "", "phone")

	rootCmd.AddCommand(registerCmd);
}
