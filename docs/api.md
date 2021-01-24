## API

### User

#### POST`/api/user/login`

登录

`application/x-www-form-urlencoded`

| 参数      | 类型 | 备注                |
| --------- | ---- | ------------------- |
| loginName | 必选 | 登录名（手机/邮箱） |
| password  | 必选 | 密码                |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |
| token        | 用户token    |
| refreshToken | refreshToken |



#### POST`/api/user/username`

检验用户名合法性

`application/x-www-form-urlencoded`

| 参数     | 类型 | 备注   |
| -------- | ---- | ------ |
| username | 必选 | 用户名 |



#### POST`/api/user/password`

检验密码合法性

`application/x-www-form-urlencoded`

| 参数     | 类型 | 备注 |
| -------- | ---- | ---- |
| password | 必选 | 密码 |



#### POST`/api/user/register`

`application/x-www-form-urlencoded`

|参数|类型|备注|
|----|----|----|
|username|必选|用户名/账号|
|password|必选|密码|
|phone|必选|手机号|
|verify_code|必选|验证码|

先调用接口发送验证码， 并检查用户名及密码的规范性，确认数据符合规范之后再发送表单

#### POST`/api/user/login`

|参数|类型|备注|
|----|----|----|
|username|必选|用户名/账号|
|password|必选|密码|



### Verify

#### GET `/api/verify/token`

使用refreshToken获取新token

| 参数         | 备注         |
| ------------ | ------------ |
| refreshToken | refreshToken |

| 返回参数 | 备注                                                         |
| -------- | ------------------------------------------------------------ |
| data     | 成功则为新的token，若refreshToken失效则为 "refreshToken失效" |
| status   | 状态码                                                       |



#### POST`/api/verify/phone`

`application/x-www-form-urlencoded`

发送短信验证码

| 参数  | 类型 | 备注   |
| ----- | ---- | ------ |
| phone | 必选 | 手机号 |

## 

## 一般规定

无特殊说明下，返回一个json，含且仅含status(1：成功， 0：失败)， data(成功提示或者错误)