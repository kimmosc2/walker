package leave

import (
	"github.com/pkg/errors"
	"school-walker/walker/model"
	"school-walker/walker/pkg/schoopia"
	"school-walker/walker/serialize"
)

// CheckCookie check user's authority
func CheckCookie(cookie string) (serialize.SchoolResponse, error) {
	return schoopia.CheckAuth(cookie)
}

// CheckAuthExists verify user exists
func CheckAuthExists(number string) (model.ResponseUser, error) {
	user := new(model.WalkerAuth)
	if err := model.DB.Where("school_number = ?", number).First(&user).Error; err != nil {
		return model.ResponseUser{}, errors.Wrap(err, "can not find user")
	}
	return model.BuildUserResponse(*user), nil
}

// CreateAuth create record to database
func CreateAuth(auth model.WalkerAuth) int {
	return int(model.DB.Create(&auth).RowsAffected)
}

// SaveAuth save authority info
func SaveAuth(auth model.WalkerAuth) error {
	return errors.Wrapf(model.DB.Model(&auth).Where("school_number = ?", auth.SchoolNumber).Updates(auth).Error, "update auth failed: Auth:%#v", auth)
}
