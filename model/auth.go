package model

import (
	"school-walker/walker/pkg/schoopia"
)

// database struct
type WalkerAuth struct {
	ID uint `gorm:"primarykey;autoIncrement"`
	// username
	Name string
	// telephone number
	Tel string
	// teacher
	Teacher string
	// contacts
	Contact string
	// contacts telephone
	ContactTel string
	// where u go
	Direction string
	// why u leave
	Reason string
	// school number
	SchoolNumber string
	// last cookie
	LastCookie string
	// start time
	StartTime uint8
	// end time
	EndTime      uint8
	PayStartTime int64
	PayEndTime   int64
	// state,0 as disabled, 1 as enabled
	State      int8
	CreatedAt  int64 `gorm:"autoCreateTime"`
	ModifiedAt int64 `gorm:"autoUpdateTime"`
}

func (w WalkerAuth) TableName() string {
	return "walker_auths"
}

// response user
type ResponseUser struct {
	ID           uint   `json:"id"`
	SchoolIdCode string `json:"school_idcode"`
	// name
	Name       string `json:"name"`
	Teacher    string `json:"teacher"`
	Contact    string `json:"contact"`
	ContactTel string `json:"contact_tel"`
	Direction  string `json:"direction"`
	Reason     string `json:"reason"`
	// hour
	StartTime uint8 `json:"start_time"`
	// hour
	EndTime uint8 `json:"end_time"`
	// 0 as disabled, 1 as enabled
	State int8 `json:"state"`
}

// make database model to response model
func BuildUserResponse(auth WalkerAuth) ResponseUser {
	return ResponseUser{
		ID:           auth.ID,
		SchoolIdCode: auth.SchoolNumber,
		Name:         auth.Name,
		Teacher:      auth.Teacher,
		Contact:      auth.Contact,
		ContactTel:   auth.ContactTel,
		Direction:    auth.Direction,
		Reason:       auth.Reason,
		StartTime:    auth.StartTime,
		EndTime:      auth.EndTime,
		State:        auth.State,
	}
}

func BuildWalkAuth(user schoopia.GlobalUserInfo) WalkerAuth {
	return WalkerAuth{
		ID:           uint(user.Id),
		Name:         user.User.Name,
		Tel:          user.User.Mobile,
		Contact:      "",
		ContactTel:   "",
		Direction:    "",
		Reason:       "",
		SchoolNumber: user.SchoolIdCode,
		LastCookie:   user.PhpSessionId,
		StartTime:    8,
		EndTime:      20,
		State:        0,
	}
}
