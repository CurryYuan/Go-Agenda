package entity

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
)

const userPath string = "/users.json"
const meetingPath string = "/meetings.json"
const curUserPath string = "/curUser.json"

var curUser User
var userList []User
var meetingList []Meeting
var dirty bool

type Storage struct{}

var instance *Storage
var lock *sync.Mutex
var dname string

func init() {
	dirty = false
	instance = nil
	lock = &sync.Mutex{}
	GOPATH := os.Getenv("GOPATH")
	dname = GOPATH + "/src/agenda/data"
	readFromFile()
}

/**
 *   read file content into memory
 *   @return if success, true will be returned
 */
func readFromFile() error {
	var errlist []error
	if err := readByJSON(userPath, &userList); err != nil {
		errlist = append(errlist, err)
	}
	if err := readByJSON(meetingPath, &meetingList); err != nil {
		errlist = append(errlist, err)
	}
	if err := readByJSON(curUserPath, &curUser); err != nil {
		errlist = append(errlist, err)
	}
	switch len(errlist) {
	case 1:
		return errlist[0]
	case 2:
		return errors.New(errlist[0].Error() + "\n" + errlist[1].Error())
	case 3:
		return errors.New(errlist[0].Error() + "\n" + errlist[1].Error() + "\n" + errlist[2].Error())
	default:
		return nil
	}
}

func readByJSON(path string, data interface{}) error {
	os.MkdirAll(dname, os.ModeDir|os.ModePerm)
	file, err := os.OpenFile(dname+path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	switch err := json.NewDecoder(file).Decode(data); err {
	case nil, io.EOF:
		return nil
	default:
		return err
	}
}

/**
 *   write file content from memory
 *   @return if success, true will be returned
 */
func writeToFile() error {
	var errlist []error
	if err := writeByJSON(userPath, &userList); err != nil {
		errlist = append(errlist, err)
	}
	if err := writeByJSON(meetingPath, &meetingList); err != nil {
		errlist = append(errlist, err)
	}
	if err := writeByJSON(curUserPath, &curUser); err != nil {
		errlist = append(errlist, err)
	}

	switch len(errlist) {
	case 1:
		return errlist[0]
	case 2:
		return errors.New(errlist[0].Error() + "\n" + errlist[1].Error())
	case 3:
		return errors.New(errlist[0].Error() + "\n" + errlist[1].Error() + "\n" + errlist[2].Error())
	default:
		return nil
	}
}

/**
 *  write file by JSON
 */
func writeByJSON(path string, data interface{}) error {
	os.MkdirAll(dname, os.ModeDir|os.ModePerm)
	file, err := os.OpenFile(dname+path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(data); err != nil {
		return err
	}
	return nil
}

/**
 * get Instance of storage
 * @return the pointer of the instance
 */
func GetInstance() *Storage {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &Storage{}
	}
	return instance
}

/**
 * create a user
 * @param a user object
 */
func (instance *Storage) CreateUser(t_user User) {
	userList = append(userList, t_user)
	dirty = true
}

/**
 * query users
 * @param a lambda function as the filter
 * @return a list of fitted users
 */
func (instance *Storage) QueryUser(filter func(*User) bool) []User {
	var result []User
	for _, v := range userList {
		if filter(&v) {
			result = append(result, v)
		}
	}
	return result
}

/**
 * update users
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the user
 * @return the number of updated users
 */
func (instance *Storage) UpdateUser(filter func(*User) bool, swither func(*User)) int {
	count := 0
	len := len(userList)
	for i := 0; i < len; i++ {
		if filter(&userList[i]) {
			swither(&userList[i])
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

/**
 * delete users
 * @param a lambda function as the filter
 * @return the number of deleted users
 */
func (instance *Storage) DeleteUser(filter func(*User) bool) int {
	count := 0
	len := len(userList)
	for i := 0; i < len; {
		if filter(&userList[i]) {
			count++
			if i+1 < len {
				userList = append(userList[:i], userList[i+1:]...)
			} else {
				userList = userList[:i]
			}
			len--
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

/**
 * create a meeting
 * @param a meeting object
 */
func (instance *Storage) CreateMeeting(t_meeting Meeting) {
	meetingList = append(meetingList, t_meeting)
	dirty = true
}

/**
 * query meetings
 * @param a lambda function as the filter
 * @return a list of fitted meetings
 */
func (instance *Storage) QueryMeeting(filter func(*Meeting) bool) []Meeting {
	var result []Meeting
	for _, v := range meetingList {
		if filter(&v) {
			result = append(result, v)
		}
	}
	return result
}

/**
 * update meetings
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the meeting
 * @return the number of updated meetings
 */
func (instance *Storage) UpdateMeeting(filter func(*Meeting) bool, switcher func(*Meeting)) int {
	count := 0
	len := len(meetingList)
	for i := 0; i < len; i++ {
		if filter(&meetingList[i]) {
			switcher(&meetingList[i])
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

/**
 * delete meetings
 * @param a lambda function as the filter
 * @return the number of deleted meetings
 */

func (instance *Storage) DeleteMeeting(filter func(*Meeting) bool) int {
	count := 0
	len := len(meetingList)
	for i := 0; i < len; {
		if filter(&meetingList[i]) {
			count++
			if i+1 < len {
				meetingList = append(meetingList[:i], meetingList[i+1:]...)
			} else {
				meetingList = meetingList[:i]
			}
			len--
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

/**
 * set current username
 * @param username of current user
 */
func (instance *Storage) SetCurUsername(username string) {
	curUser.SetName(username)
	dirty = true
}

/**
 * get current username
 * @return curUsername
 */
func (instance *Storage) GetCurUsername() string {
	return curUser.GetName()
}

/**
 * sync with the file
 */
func (instance *Storage) Sync() error {
	if dirty {
		if err := writeToFile(); err != nil {
			return err
		} else {
			dirty = false
		}
	}
	return nil
}
