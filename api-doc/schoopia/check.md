登录信息验证
===
#### 简要描述
 - 当前登录用户信息
#### 请求url
 - `http://sso.schoopia.com/login/check`
#### 请求方式
 - GET
#### 参数
| 参数名 | 必选 | 类型 | 说明 |
| --- | --- | --- | --- |
| key  | 否 | int | 暂时不清楚什么用,一串长度为10的数字,例如4055801924|

#### 返回示例
```
{
  "data": {
    "refresh_token": "xxxxxxxxxxxxxxxxxxxxxx",
    "changed_password": 1,
    "user_id": xxxxxxxxxxxxxxxxxxxxxx,
    "user_name": "xxxxxxxxxxxxxxxxxxxxxx",
    "phpSessionId": "xxxxxxxxxxxxxxxxxxxxxx",
    "teacher_mobile": 1,
    "bind_device": 0,
    "student_mobile": 1,
    "user": {
      "id": xxxxxxxxxxxxxxxxxxxxxx,
      "update_time": xxxxxxxxxxxxxxxxxxxxxx,
      "user_name": "xxxxxxxxxxxxxxxxxxxxxx",
      "name": "xxxxxxxxxxxxxxxxxxxxxx",
      "type_id": xxxxxxxxxxxxxxxxxxxxxx,
      "status_id": 1,
      "school_idcode": "xxxxxxxxxxxxxxxxxxxxxx",
      "nick": "xxxxxxxxxxxxxxxxxxxxxx",
      "mobile": "xxxxxxxxxxxxxxxxxxxxxx",
      "email": null,
      "sex": 1,
      "qq": null,
      "birthday": null,
      "avatar_id": xxxxxxxxxxxxxxxxxxxxxx,
      "student": {
        "student_class": {
          "id": xxxxxxxxxxxxxxxxxxxxxx,
          "class_id": xxxxxxxxxxxxxxxxxxxxxx,
          "index": xxxxxxxxxxxxxxxxxxxxxx,
          "code": "xxxxxxxxxxxxxxxxxxxxxx",
          "class_name": "xxxxxxxxxxxxxxxxxxxxxx",
          "grade": 2017,
          "major": {
            "code": "xxxxxxxxxxxxxxxxxxxxxx",
            "name": "xxxxxxxxxxxxxxxxxxxxxx",
            "abb": null,
            "degree_name": "xxxxxxxxxxxxxxxxxxxxxx",
            "dpm": {
              "id": xxxxxxxxxxxxxxxxxxxxxx,
              "code": "01",
              "name": "xx部",
              "abb": "xx部",
              "school_code": "xxxxx",
              "school_name": "xxxx学院",
              "school": {
                "code": "xxxxx",
                "name": "xxxx学院"
              }
            }
          }
        },
        "major": {
          "code": "xxxx",
          "name": "xxxx",
          "abb": null,
          "degree_name": "xxxx",
          "dpm": {
            "id": xxxx,
            "code": "01",
            "name": "xxxx部",
            "abb": "xxxx部",
            "school_code": "xxxx",
            "school_name": "xxxx学院",
            "school": {
              "code": "xxxx",
              "name": "xxxx学院"
            }
          }
        }
      },
      "type": {
        "name": "学生",
        "id": xxxx,
        "tag": "student"
      },
      "avatar_image": {
        "id": xxxx,
        "update_time": xxxx,
        "width": 240,
        "height": 240,
        "format": "jpg",
        "url": "xxxx",
        "local": false
      },
      "SPORT_USER_SCHOOL": false,
      "school_code": "xxxx"
    },
    "device": null
  },
  "code": 0,
  "time": xxxx,
  "desc": "success"
}
```
#### 返回参数说明
 - 无
#### 备注
- 敏感信息做了处理