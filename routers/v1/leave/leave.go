package leave

import (
	"github.com/gin-gonic/gin"
	"school-walker/walker/model"
	"school-walker/walker/pkg/schoopia"
	"school-walker/walker/serialize"
	"school-walker/walker/serialize/ecode"
	leaveV1 "school-walker/walker/service/v1/leave"
	"school-walker/walker/service/v1/syncdata"
)

type FormData struct {
	SchoolIdcode string `form:"school_idcode" json:"school_idcode" binding:"required"`
	Teacher      string `form:"teacher" json:"teacher"`
	Contact      string `form:"contact" json:"contact"`
	ContactTel   string `form:"contact_tel" json:"contact_tel"`
	Direction    string `form:"direction" json:"direction"`
	Reason       string `form:"reason" json:"reason"`
	StartTime    uint8  `form:"start_time" json:"start_time"`
	EndTime      uint8  `form:"end_time" json:"end_time"`
}

func (f FormData) ToAuth() model.WalkerAuth {
	return model.WalkerAuth{
		Teacher:      f.Teacher,
		Contact:      f.Contact,
		ContactTel:   f.ContactTel,
		Direction:    f.Direction,
		Reason:       f.Reason,
		SchoolNumber: f.SchoolIdcode,
		StartTime:    f.StartTime,
		EndTime:      f.EndTime,
	}
}

// AuthCheck check user's authority
func AuthCheck(c *gin.Context) {
	cookie := c.Query("cookie")
	if cookie == "" {
		c.JSON(200, serialize.BuildErrorResponse(ecode.StatusErrorParams))
		return
	}
	resp, err := leaveV1.CheckCookie(cookie)
	if err != nil {
		c.Set("err", err)
		c.JSON(200, serialize.BuildErrorResponse(ecode.StatusInternalError))
		return
	}
	switch resp.Code {
	case 0:
		user := resp.Data.(*schoopia.GlobalUserInfo)
		// if database exists this auth
		auth, err := leaveV1.CheckAuthExists(user.User.SchoolIdCode)
		if err != nil {

			auth = model.ResponseUser{}
			auth.Name = user.Name
			auth.SchoolIdCode = user.SchoolIdCode

			t := model.BuildWalkAuth(*user)
			t.State = 1
			t.ID = 0

			leaveV1.CreateAuth(t)
			// TODO: reformat this json
			c.JSON(200, serialize.BuildCustomResponse(auth, ecode.GetMsg(ecode.StatusFirstUseLeave), ecode.StatusFirstUseLeave))
			return
		}
		user.Id = int(auth.ID)
		if err = syncdata.SyncCookie(model.BuildWalkAuth(*user)); err != nil {
			c.Set("err", err)
			c.JSON(200, serialize.BuildErrorResponse(ecode.StatusInternalError))
			return
		}
		c.JSON(200, serialize.BuildCustomResponse(auth, ecode.GetMsg(ecode.StatusOK), ecode.StatusOK))
		return
	// 	login failure
	case -4:
		c.JSON(200, serialize.BuildErrorResponse(ecode.StatusLoginFailure))
		return
	default:
		c.JSON(200, serialize.BuildCustomResponse(resp, ecode.GetMsg(ecode.StatusUnknown), ecode.StatusUnknown))
		return
	}
}

// SaveAuth save user info
func SaveAuth(c *gin.Context) {
	var d FormData
	if err := c.ShouldBind(&d); err != nil {
		c.JSON(200, serialize.BuildErrorResponse(ecode.StatusErrorParams))
		return
	}
	leaveV1.SaveAuth(d.ToAuth())
	c.JSON(200, serialize.BuildErrorResponse(ecode.StatusOK))
}
