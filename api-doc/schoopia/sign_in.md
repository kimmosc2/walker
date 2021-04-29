签到接口
===

#### 简要描述  
 - 学生上课点名签到接口
 
#### 请求url
 - `http://b.schoopia.com/bapp/bcourse/course/signin`
 
#### 请求方式
 - POST
 
#### 参数
| 参数名 | 必选 | 类型 | 说明 |  
| --- | --- | --- | --- |  
| code | 是 | int | 课堂点名码   |
| latitude | 否 | double | 纬度   |
| longitude | 否 | double | 经度   |

#### 请求头示例
```
Accept: */*
Accept-Encoding: gzip, deflate
Accept-Language: zh-Hans-CN;q=1, en-CN;q=0.9
Connection: close
Content-Length: 55
Content-Type: application/x-www-form-urlencoded
Cookie: PHPSESSID=xxx-xxx-xxx-xxx; UM_distinctid=xxxxxxxxxxxxx
Host: b.schoopia.com
User-Agent: fan zhuan xiao yuan/4.5.6 (Android; Android 12.6; Scale/2.00)
```
#### 请求参数示例
```
code=111111&latitude=22.222222222222&longitude=33.33333333333
```
#### 返回示例
```
{
  "code": -4,
  "data": {},
  "desc": "点名码不存在或已结束",
  "errCode": "500",
  "errMsg": "点名码不存在或已结束",
  "exception": "com.yuhong.core.BusinessException",
  "rpcMsg": "",
  "success": false,
  "throwable": "{\"cause\":null,\"stackTrace\":[{\"methodName\":\"signIn\",\"fileName\":\"TRollCallService.java\",\"lineNumber\":209,\"className\":\"com.yuhong.bcourse.service.TRollCallService\",\"nativeMethod\":false}],\"message\":\"点名码不存在或已结束\",\"localizedMessage\":\"点名码不存在或已结束\",\"suppressed\":[]}",
  "time": xxxxxxxxxx
}
```

#### 返回参数说明
|参数名  |类型 |说明|  
| ---  | --- | ---|  
| code | int | 状态码 |
| desc | string | 描述 |  
#### 备注
