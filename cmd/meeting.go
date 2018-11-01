package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"agenda/entity"
	"agenda/log"
)

var infoLog=log.Info
var errLog=log.Error

var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "create meeting",
	Long: ``,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")	
		participatorStr, _ := comd.Flags().GetString("participators")
		participators := strings.Split(participatorStr, " ")

		startTime, _ := comd.Flags().GetString("start")

		endTime, _ := comd.Flags().GetString("end")

		if err := entity.CreateMeeting(title, participators, startTime, endTime); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("create meeting success")
			fmt.Println("create meeting success")
		}
	},
}

var addParCmd = &cobra.Command{
	Use:   "addPar",
	Short: "Add your own meetings' participators.",
	Long: `You can append some participators from your own meeting
	by specifying the title name.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")

		participatorStr, _ := comd.Flags().GetString("participators")

		participators := strings.Split(participatorStr, ",")

		if err := entity.AddPar(title, participators); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("Add meeting participator successfully!")
			fmt.Println("Add meeting participator successfully!")
		}
	},
}

var removeParCmd = &cobra.Command{
	Use:   "removePar",
	Short: "Remove your own meetings' participators.",
	Long: `You can remove some participators from your own meeting
	by specifying the title name.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")

		participatorStr, _ := comd.Flags().GetString("participators")

		participators := strings.Split(participatorStr, ",")

		if err := entity.RemovePar(title, participators); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("Remove meeting participator successfully!")
			fmt.Println("Remove meeting participator successfully!")
		}
	},
}

var listMeetingsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your own meetings during a time interval.",
	Long: `You can see the detail information of all of meetings,
	which you attended, during a time interval.`,
	Run: func(comd *cobra.Command, args []string) {
		startTime, _ := comd.Flags().GetString("start")

		endTime, _ := comd.Flags().GetString("end")

		if err := entity.ListMeetings(startTime, endTime); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("Listing meeting operation completed successfully!")
			fmt.Println("Listing meeting operation completed successfully!")
		}
	},
}

var cancelMeetingCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel your own meeting by specifying title name.",
	Long:  `Using this command, you are able to cancel the meetings, which are created by you.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")

		if err := entity.CancelMeeting(title); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("The meeting was successfully deleted!")
			fmt.Println("The meeting was successfully deleted!")
		}
	},
}

var quitMeetingCmd = &cobra.Command{
	Use:   "quit",
	Short: "Quit others meeting by specifying title name",
	Long:  `You can quit any meetings you want, which are you attended, not created.`,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")

		if err := entity.QuitMeeting(title); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("You've successfully quit the meeting " + title + "!")
			fmt.Println("You've successfully quit the meeting " + title + "!")
		}
	},
}

var clearMeetingsCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all meetings you attended or created.",
	Long:  `Using this command, you can clear all of the meetings you attended or created.`,
	Run: func(comd *cobra.Command, args []string) {

		if err := entity.ClearMeeting(); err != nil {
			errLog.Println(err)
			fmt.Println(err)
		} else {
			infoLog.Println("You've successfully cleared all the meetings you sponsored!")
			fmt.Println("You've successfully cleared all the meetings you sponsored!")

		}
	},
}

func init() {
	createMeetingCmd.Flags().StringP("title", "t", "", "title")
	createMeetingCmd.Flags().StringP("participators", "p", "", "participator name")
	createMeetingCmd.Flags().StringP("start", "s", "", "start time (format: yyyy-mm-dd/hh:mm)")
	createMeetingCmd.Flags().StringP("end", "e", "", "end time (format: yyyy-mm-dd/hh:mm)")
	rootCmd.AddCommand(createMeetingCmd)

	addParCmd.Flags().StringP("title", "t", "", "meeting title")
	addParCmd.Flags().StringP("participators", "p", "", "meeting participators")
	rootCmd.AddCommand(addParCmd)

	removeParCmd.Flags().StringP("title", "t", "", "meeting title")
	removeParCmd.Flags().StringP("participators", "p", "", "meeting participators name")
	rootCmd.AddCommand(removeParCmd)

	listMeetingsCmd.Flags().StringP("start", "s", "", "start time (format: yyyy-mm-dd/hh:mm)")
	listMeetingsCmd.Flags().StringP("end", "e", "", "end time (format: yyyy-mm-dd/hh:mm)")
	rootCmd.AddCommand(listMeetingsCmd)

	cancelMeetingCmd.Flags().StringP("title", "t", "", "meeting title")
	rootCmd.AddCommand(cancelMeetingCmd)

	quitMeetingCmd.Flags().StringP("title", "t", "", "meeting title")
	rootCmd.AddCommand(quitMeetingCmd)
}
