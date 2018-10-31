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

		if err := entity.Register(username,password,mail,phone); err ==nil {
			fmt.Println("register success")
		}
	},
}

var loginCmd = &cobra.Command{
	Use: "login",
	Short: "user login",
	Long: ``,
	Run: func(com *cobra.Command, args []string){
		username, _ := com.Flags().GetString("user")
		checkEmpty("username", username)

		password, _ := com.Flags().GetString("password")
		checkEmpty("password", password)

		if err := entity.Login(username,password); err ==nil {
			fmt.Println("login success")
		}
	},
}

var logoutCmd = &cobra.Command{
	Use: "logout",
	Short: "user logout",
	Long: ``,
	Run: func(com *cobra.Command, args []string){
		if err := entity.Logout(); err ==nil {
			fmt.Println("logout success")
		}
	},
}

var listUsersCmd = &cobra.Command{
	Use: "listUsers",
	Short: "list all users",
	Long: ``,
	Run: func(com *cobra.Command, args []string){
		if err := entity.ListUsers(); err ==nil {
			fmt.Println("list user success")
		}
	},
}

var delUserCmd = &cobra.Command{
	Use: "delUser",
	Short: "delete user",
	Long: ``,
	Run: func(com *cobra.Command, args []string){
		if err := entity.DelUser(); err ==nil {
			fmt.Println("del user success")
		}
	},
}

func init() {
	registerCmd.Flags().StringP("user", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "password")
	registerCmd.Flags().StringP("mail", "m", "", "email")
	registerCmd.Flags().StringP("phone", "t", "", "phone")

	loginCmd.Flags().StringP("user", "u", "", "username")
	loginCmd.Flags().StringP("password", "p", "", "password")

	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(listUsersCmd)
	rootCmd.AddCommand(delUserCmd)
}
