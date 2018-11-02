package cmd

import (
	"fmt"
	//"os"

	"agenda/entity"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register user",
	Long:  "Register a new user",
	Run: func(com *cobra.Command, args []string) {
		username, _ := com.Flags().GetString("user")

		password, _ := com.Flags().GetString("password")

		mail, _ := com.Flags().GetString("mail")

		phone, _ := com.Flags().GetString("phone")

		if err := entity.Register(username, password, mail, phone); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			fmt.Println("register success")
			infoLog.Println("user " + username + " register success")
		}
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "user login",
	Long:  ``,
	Run: func(com *cobra.Command, args []string) {
		username, _ := com.Flags().GetString("user")

		password, _ := com.Flags().GetString("password")

		if err := entity.Login(username, password); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			fmt.Println("login success")
			infoLog.Println("user " + username + " login success")
		}
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "user logout",
	Long:  ``,
	Run: func(com *cobra.Command, args []string) {
		if err := entity.Logout(); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			fmt.Println("logout success")
			infoLog.Println("user logout success")
		}
	},
}

var listUsersCmd = &cobra.Command{
	Use:   "listUsers",
	Short: "list all users",
	Long:  ``,
	Run: func(com *cobra.Command, args []string) {
		if err := entity.ListUsers(); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("list all users")
		}
	},
}

var delUserCmd = &cobra.Command{
	Use:   "delUser",
	Short: "delete user",
	Long:  ``,
	Run: func(com *cobra.Command, args []string) {
		if err := entity.DelUser(); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			fmt.Println("delete user success")
			infoLog.Println("delete user success")
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
