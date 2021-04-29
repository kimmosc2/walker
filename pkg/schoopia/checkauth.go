package schoopia


import (
"encoding/json"
"io"
"net/http"

"github.com/pkg/errors"
"school-walker/walker/serialize"
)

const (
	LoginCheckURL = "http://sso.schoopia.com/login/check"
)

// GlobalUserInfo
type GlobalUserInfo struct {
	RefreshToken string `json:"refresh_token"`

	UserId int `json:"user_id"`
	// UserName is login name for schoopia,school code append student number
	UserName string `json:"user_name"`

	// PhpSessionId current user cookie
	PhpSessionId string `json:"phpSessionId"`

	User `json:"user"`
}

// User More detailed user information
type User struct {
	Id int `json:"id"`
	// UpdateTime unix timestamp
	UpdateTime int `json:"update_time"`
	// Name username
	Name string `json:"name"`

	TypeId int `json:"type_id"`
	// SchoolIdCode school code
	SchoolIdCode string `json:"school_idcode"`
	// Mobile Mobile Number
	Mobile string `json:"mobile"`
	// detailed info
	Student struct {
		StudentClass struct {
			//  class name
			ClassName string `json:"class_name"`
			// grade
			Grade int `json:"grade"`
			Major struct {
				Name       string `json:"name"`
				DegreeName string `json:"degree_name"`
				// department info
				DPM struct {
					Id int `json:"id"`
					// department name
					Name string `json:"name"`
					// SchoolName schoolName
					SchoolName string `json:"school_name"`
					// 	School SchoolIdcode
					SchoolCode string `json:"school_code"`
				} `json:"dpm"`
			} `json:"major"`
		} `json:"student_class"`
	} `json:"student"`
}

// CheckAuth 查看当前cookie是否过期,如果有效,则返回用户
func CheckAuth(cookie string) (serialize.SchoolResponse, error) {
	request, err := http.NewRequest(http.MethodGet, LoginCheckURL, nil)
	if err != nil {
		return serialize.SchoolResponse{}, errors.Wrap(err,"build request fail")
	}
	setSchoolHeader(request)
	request.Header.Set("Cookie", cookie)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return serialize.SchoolResponse{}, errors.Wrap(err,"http request error")
	}
	defer response.Body.Close()
	resp := new(serialize.SchoolResponse)
	global := new(GlobalUserInfo)
	resp.Data = global
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil && err != io.EOF {
		return serialize.SchoolResponse{}, errors.Wrap(err,"json decode error")
	}
	return *resp, nil
}

// setSchoolHeader set http request header look like send from school client
func setSchoolHeader(r *http.Request) {
	r.Header.Set("Accept", `*/*`)
	// r.Header.Set("Accept-Encoding", `gzip, deflate`)
	r.Header.Set("Accept-Language", `zh-Hans-CN;q=1, en-CN;q=0.9`)
	r.Header.Set("Connection", `close`)
	r.Header.Set("Content-Type", `application/x-www-form-urlencoded`)
	r.Header.Set("Host", `b.schoopia.com`)
	r.Header.Set("User-Agent", `fan zhuan xiao yuan/4.5.4 (iPhone; iOS 13.6.1; Scale/2.00)`)
}
