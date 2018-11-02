package entity

import (
	"errors"
	"fmt"
)

/**
 * check if the username match password
 * @param userName the username want to login
 * @param password the password user enter
 * @return if success, nil will be returned
 */
func Register(username, password, email, phone string) error {
	if username == "" || password == "" || email == "" || phone == "" {
		return errors.New("Expect 4 parameters")
	}
	storage := GetInstance()
	if userList := storage.QueryUser(func(user *User) bool {
		return user.GetName() == username
	}); len(userList) != 0 {
		return errors.New("Already exist " + username)
	} else {
		storage.CreateUser(User{username, password, email, phone})
		return storage.Sync()
	}
}

/**
 * check if the username match password
 * @param userName the username want to login
 * @param password the password user enter
 * @return if success, true will be returned
 */
func Login(username, password string) error {
	if username == "" || password == "" {
		return errors.New("Expect 2 parameters")
	}
	storage := GetInstance()
	if userList := storage.QueryUser(func(user *User) bool {
		return user.GetName() == username
	}); len(userList) == 0 {
		return errors.New(username + " doesn't exist")
	} else {
		storage.SetCurUsername(username)
		return storage.Sync()
	}
}

/**
 * quit the Agenda
 */
func Logout() error {
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	storage.SetCurUsername("")
	return storage.Sync()
}

/**
 * list all users from storage
 */
func ListUsers() error {
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	userList := storage.QueryUser(func(user *User) bool {
		return true
	})
	fmt.Println("[list all users]")
	fmt.Println("---------------------------------------------------")
	for _, user := range userList {
		fmt.Println("[username]	" + user.GetName())
		fmt.Println("[email]		" + user.GetEmail())
		fmt.Println("[phone]		" + user.GetPhone())
		fmt.Println("---------------------------------------------------")
	}
	return nil
}

func DelUser() error {
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	curUsername := storage.GetCurUsername()
	if result := storage.DeleteUser(func(user *User) bool {
		return user.GetName() == curUsername
	}); result == 0 {
		return errors.New("fail to delete " + curUsername)
	}
	storage.UpdateMeeting(func(meeting *Meeting) bool {
		return meeting.IsParticipator(curUsername)
	}, func(meeting *Meeting) {
		meeting.RemoveParticipator(curUsername)
	})
	storage.DeleteMeeting(func(meeting *Meeting) bool {
		return meeting.GetSponsor() == curUsername || len(meeting.GetParticipators()) == 0
	})
	storage.SetCurUsername("")
	return storage.Sync()
}

func CreateMeeting(title string, participators []string, start string, end string) error {
	if title == "" || len(participators) == 0 || start == "" || end == "" {
		return errors.New("Expect 4 parameters")
	}
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	var startDate, endDate Date
	var err error
	if startDate, err = StringToDate(start); err != nil {
		return err
	}
	if endDate, err = StringToDate(end); err != nil {
		return err
	}
	if startDate.Less(endDate) == false {
		return errors.New("End date is earlier than( equal to) start date!")
	}
	if meetingList := storage.QueryMeeting(func(meeting *Meeting) bool {
		return meeting.GetTitle() == title
	}); len(meetingList) != 0 {
		return errors.New("exist meeting " + title)
	}
	curUsername := storage.GetCurUsername()
	for _, par := range participators {
		if par == curUsername {
			return errors.New("Sponsor can't be participator")
		}
		if userList := storage.QueryUser(func(user *User) bool {
			return user.GetName() == par
		}); len(userList) == 0 {
			return errors.New(par + " isn't registered")
		}
		count := 0
		for _, par2 := range participators {
			if par == par2 {
				count++
				if count > 1 {
					return errors.New(par + " repeatedly appear")
				}
			}
		}
	}
	meetingList := storage.QueryMeeting(func(meeting *Meeting) bool {
		m_startDate, m_endDate := meeting.GetStartDate(), meeting.GetEndDate()
		return startDate.LessOrEqual(m_startDate) && m_startDate.LessOrEqual(endDate) ||
			startDate.LessOrEqual(m_endDate) && m_endDate.LessOrEqual(endDate) ||
			m_startDate.LessOrEqual(startDate) && endDate.LessOrEqual(m_endDate)
	})
	for _, meeting := range meetingList {
		if meeting.GetSponsor() == curUsername || meeting.IsParticipator(curUsername) {
			return errors.New(curUsername + " is busy for other meeting!")
		}
		for _, par := range participators {
			if meeting.GetSponsor() == par || meeting.IsParticipator(par) {
				return errors.New(par + " is busy for other meeting!")
			}
		}
	}
	storage.CreateMeeting(Meeting{curUsername, participators, startDate, endDate, title})
	return storage.Sync()
}

func AddPar(title string, participators []string) error {
	if title == "" || len(participators) == 0 {
		return errors.New("Expect 2 parameters")
	}
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	curUsername := storage.GetCurUsername()
	var startDate, endDate Date
	//check curUsername sponsor the meeting title
	if meetingList := storage.QueryMeeting(func(meeting *Meeting) bool {
		return meeting.GetSponsor() == curUsername && meeting.GetTitle() == title
	}); len(meetingList) == 0 {
		return errors.New(curUsername + " doesn't sponsor the meeting " + title)
	} else {
		startDate, endDate = meetingList[0].GetStartDate(), meetingList[0].GetEndDate()
	}
	//check participators are valid
	for _, par := range participators {
		if par == curUsername {
			return errors.New("Sponsor can't be participator")
		}
		if userList := storage.QueryUser(func(user *User) bool {
			return user.GetName() == par
		}); len(userList) == 0 {
			return errors.New(par + " isn't registered")
		}
		count := 0
		for _, par2 := range participators {
			if par == par2 {
				count++
				if count > 1 {
					return errors.New(par + " repeatedly appear")
				}
			}
		}
	}
	//check if participators are fre during the meeting title
	meetingList := storage.QueryMeeting(func(meeting *Meeting) bool {
		m_startDate, m_endDate := meeting.GetStartDate(), meeting.GetEndDate()
		return startDate.LessOrEqual(m_startDate) && m_startDate.LessOrEqual(endDate) ||
			startDate.LessOrEqual(m_endDate) && m_endDate.LessOrEqual(endDate) ||
			m_startDate.LessOrEqual(startDate) && endDate.LessOrEqual(m_endDate)
	})
	for _, meeting := range meetingList {
		for _, par := range participators {
			if meeting.GetSponsor() == par || meeting.IsParticipator(par) {
				return errors.New(par + " is busy for other meeting!")
			}
		}
	}
	//add participators
	storage.UpdateMeeting(func(meeting *Meeting) bool {
		return meeting.GetTitle() == title
	}, func(meeting *Meeting) {
		meeting.AddParticipator(participators)
	})
	return storage.Sync()
}

func RemovePar(title string, participators []string) error {
	if title == "" || len(participators) == 0 {
		return errors.New("Expect 2 parameters")
	}
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	curUsername := storage.GetCurUsername()
	var thisMeeting Meeting
	if meetingList := storage.QueryMeeting(func(meeting *Meeting) bool {
		return meeting.GetSponsor() == curUsername && meeting.GetTitle() == title
	}); len(meetingList) == 0 {
		return errors.New(curUsername + " doesn't sponsor the meeting " + title)
	} else {
		thisMeeting = meetingList[0]
	}
	var notParticipators []string
	for i := 0; i < len(participators); {
		if thisMeeting.IsParticipator(participators[i]) {
			notParticipators = append(notParticipators, participators[i])
			i++
		} else {
			if i+1 < len(participators) {
				participators = append(participators[:i], participators[i+1:]...)
			} else {
				participators = participators[:i]
			}
		}
	}
	if len(participators) > 0 {
		storage.UpdateMeeting(func(meeting *Meeting) bool {
			return meeting.GetSponsor() == curUsername && meeting.GetTitle() == title
		}, func(meeting *Meeting) {
			for _, par := range participators {
				meeting.RemoveParticipator(par)
			}
		})
		storage.DeleteMeeting(func(meeting *Meeting) bool {
			return len(meeting.GetParticipators()) == 0
		})
	}
	if len(notParticipators) > 0 {
		str := notParticipators[0]
		for i := 1; i < len(notParticipators); i++ {
			str += "," + notParticipators[i]
		}
		err := storage.Sync()
		if err != nil {
			return errors.New("Remove meeting participator successfully!\nBut" +
				str + " not the participator(s) of the meeting")
		}
	}
	return storage.Sync()
}

func ListMeetings(start, end string) error {
	if start == "" || end == "" {
		return errors.New("Expect 2 parameters")
	}
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	curUsername := storage.GetCurUsername()
	var startDate, endDate Date
	var err error
	if startDate, err = StringToDate(start); err != nil {
		return err
	}
	if endDate, err = StringToDate(end); err != nil {
		return err
	}
	if startDate.Less(endDate) == false {
		return errors.New("End date is earlier than( equal to) start date!")
	}
	meetingList := storage.QueryMeeting(func(meeting *Meeting) bool {
		if meeting.GetSponsor() != curUsername && meeting.IsParticipator(curUsername) {
			return false
		}
		m_startDate, m_endDate := meeting.GetStartDate(), meeting.GetEndDate()
		return startDate.LessOrEqual(m_startDate) && m_startDate.LessOrEqual(endDate) ||
			startDate.LessOrEqual(m_endDate) && m_endDate.LessOrEqual(endDate) ||
			m_startDate.LessOrEqual(startDate) && endDate.LessOrEqual(m_endDate)
	})
	fmt.Println("[list meetings in " + start + " —— " + end + "]")
	fmt.Println("---------------------------------------------------")
	var startTime, endTime string
	for _, meeting := range meetingList {
		fmt.Println("[sponsor]	" + meeting.GetSponsor())
		fmt.Println("[title]		" + meeting.GetTitle())
		startTime, err = DateToString(meeting.GetStartDate())
		fmt.Println("[start]   	" + startTime)
		endTime, err = DateToString(meeting.GetEndDate())
		fmt.Println("[end]		" + endTime)
		fmt.Println("[participator]\n	")
		for _, par := range meeting.GetParticipators() {
			fmt.Print("  " + par)
		}
		fmt.Println("---------------------------------------------------")
	}
	return nil
}

func CancelMeeting(title string) error {
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	curUsername := storage.GetCurUsername()
	if result := storage.DeleteMeeting(func(meeting *Meeting) bool {
		return meeting.GetSponsor() == curUsername && meeting.GetTitle() == title
	}); result == 0 {
		return errors.New(curUsername + " doesn't sponsor the meeting " + title)
	}
	return storage.Sync()
}

func QuitMeeting(title string) error {
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	curUsername := storage.GetCurUsername()
	if result := storage.UpdateMeeting(
		func(meeting *Meeting) bool {
			return meeting.GetTitle() == title && meeting.IsParticipator(curUsername)
		},
		func(meeting *Meeting) {
			meeting.RemoveParticipator(curUsername)
		}); result == 0 {
		return errors.New(curUsername + " isn't a participator of the meeting " + title)
	}
	storage.DeleteMeeting(func(meeting *Meeting) bool {
		return len(meeting.GetParticipators()) == 0
	})
	return storage.Sync()
}

func ClearMeeting() error {
	storage := GetInstance()
	if storage.GetCurUsername() == "" {
		return errors.New("Not logged in!")
	}
	curUsername := storage.GetCurUsername()
	result := storage.DeleteMeeting(func(meeting *Meeting) bool {
		return meeting.GetSponsor() == curUsername
	})
	if result == 0 {
		return errors.New(curUsername + " doesn't sponsor any meeting")
	}
	return storage.Sync()
}
