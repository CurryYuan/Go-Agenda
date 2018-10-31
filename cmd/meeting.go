package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"agenda/entity"
)

var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "create meeting",
	Long: ``,
	Run: func(comd *cobra.Command, args []string) {
		title, _ := comd.Flags().GetString("title")
		checkEmpty("title", title)

		participatorStr, _ := comd.Flags().GetString("participators")
		checkEmpty("participators", participatorStr)
		participators := strings.Split(participatorStr, " ")

		startTime, _ := comd.Flags().GetString("start")
		checkEmpty("Start Time", startTime)

		endTime, _ := comd.Flags().GetString("end")
		checkEmpty("End Time", endTime)

		if err := entity.CreateMeeting(title, participators, startTime, endTime); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("create meeting success")
		}
	},
}

func init() {
	createMeetingCmd.Flags().StringP("title", "t", "", "title")
	createMeetingCmd.Flags().StringP("participators", "p", "", "participator name")
	createMeetingCmd.Flags().StringP("start", "s", "", "start time (yyyy-mm-dd)")
	createMeetingCmd.Flags().StringP("end", "e", "", "end time (yyyy-mm-dd)")

	rootCmd.AddCommand(createMeetingCmd)
}
