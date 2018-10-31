package entity

import (
	
)

func Register(username, password, mail, phone string) error {
	
	return nil
}

func Login(user, password string) error {
	
	return nil
}

func Logout() error {
	return nil
}

func ListUsers() error {
	return nil
}

func DelUser() error {
	return nil
}

func CreateMeeting(title string, participators []string, start string, end string) error {
	return nil
}

func AddPar(title string, participators []string) error {
	return nil
}

func RemovePar(title string, participators []string) error {
	return nil
}

func ListMeetings(start, end string) error {
	return nil
}

func CancelMeeting(title string) error {
	return nil
}

func QuitMeeting(title string) error {
	return nil
}

func ClearMeeting() error {
	return nil
}