package entity

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

/*
 *	get and set
 */
func (m_User User) GetName() string {
	return m_User.Name
}

func (m_User *User) SetName(t_name string) {
	m_User.Name = t_name
}

func (m_User User) GetPassword() string {
	return m_User.Password
}

func (m_User *User) SetPassword(t_password string) {
	m_User.Password = t_password
}

func (m_User User) GetEmail() string {
	return m_User.Email
}

func (m_User *User) SetEmail(t_email string) {
	m_User.Email = t_email
}

func (m_User User) GetPhone() string {
	return m_User.Phone
}

func (m_User *User) SetPhone(t_phone string) {
	m_User.Phone = t_phone
}
