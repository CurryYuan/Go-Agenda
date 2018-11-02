package entity

type Meeting struct {
	Sponsor       string   `json:"sponsor"`
	Participators []string `json:"participators"`
	StartDate     Date     `json:"startDate"`
	EndDate       Date     `json:"endDate"`
	Title         string   `json:"title"`
}

/*
 *	get and set
 */
func (m_meeting Meeting) GetSponsor() string {
	return m_meeting.Sponsor
}

func (m_meeting *Meeting) SetSponsor(t_sponsor string) {
	m_meeting.Sponsor = t_sponsor
}

func (m_meeting Meeting) GetParticipators() []string {
	return m_meeting.Participators
}

func (m_meeting *Meeting) SetParticipators(t_participators []string) {
	for _, v := range t_participators {
		m_meeting.Participators = append(m_meeting.Participators, v)
	}
}

func (m_meeting Meeting) GetStartDate() Date {
	return m_meeting.StartDate
}

func (m_meeting *Meeting) SetStartDate(t_startDate Date) {
	m_meeting.StartDate = t_startDate
}

func (m_meeting Meeting) GetEndDate() Date {
	return m_meeting.EndDate
}

func (m_meeting *Meeting) SetEndDate(t_endDate Date) {
	m_meeting.EndDate = t_endDate
}

func (m_meeting Meeting) GetTitle() string {
	return m_meeting.Title
}

func (m_meeting *Meeting) SetTitle(t_title string) {
	m_meeting.Title = t_title
}

/**
 * @brief add new participators to the meeting
 * @param the new participator
 */
func (m_meeting *Meeting) AddParticipator(t_usernames []string) {
	m_meeting.Participators = append(m_meeting.Participators, t_usernames...)
}

/**
 * @brief remove a participator of the meeting
 * @param the participator to be removed
 */
func (m_meeting *Meeting) RemoveParticipator(t_username string) {
	len := len(m_meeting.Participators)
	for i := 0; i < len; i++ {
		if m_meeting.Participators[i] == t_username {
			if i+1 < len {
				m_meeting.Participators = append(m_meeting.Participators[:i], m_meeting.Participators[i+1:]...)
			} else {
				m_meeting.Participators = m_meeting.Participators[:i]
			}
			break
		}
	}
}

/**
 * @brief check if the user take part in this meeting
 * @param t_username the source username
 * @return if the user take part in this meeting
 */
func (m_meeting *Meeting) IsParticipator(t_username string) bool {
	for _, username := range m_meeting.Participators {
		if username == t_username {
			return true
		}
	}
	return false
}
