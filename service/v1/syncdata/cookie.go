package syncdata

import (
	"github.com/pkg/errors"
	"school-walker/walker/model"
)

func SyncCookie(auth model.WalkerAuth) error {
	return errors.Wrap(model.DB.Model(&auth).Update("last_cookie", auth.LastCookie).Error,"synchronize cookie fail")
}


