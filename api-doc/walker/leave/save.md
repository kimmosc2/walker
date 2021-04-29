信息保存
===

#### 简要描述
 - 接受用户保存的数据,持久化到数据库中

#### 请求url
 - `/api/v1/leave/save`

#### 请求方式
 - POST

#### 参数
| 参数名 | 必选 | 类型 | 说明 |
| --- | --- | --- | --- |
| code | 是 | string | 学号 |
| teacher | 否 | string | 辅导员 |
| contact | 否 | string | 紧急联系人 |
| contact_tel | 否 | string | 紧急联系人电话 |
| direction | 否 | string | 去向 |
| reason | 否 | string | 请假事由 |
| start_time | 否 | unsigned int 8 | 开始时间(24小时制) |
| end_time | 否 | unsigned int 8 | 结束时间(24小时制) |

#### 返回示例
```
{
}
```
#### 返回参数说明
|参数名  |类型 |说明|  
| ---  | --- | ---|  
| code | int | 状态码 |
| desc | string | 描述 |  
| data | objection | 内容 |
#### 备注
