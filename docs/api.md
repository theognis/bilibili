## API

### User

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

先调用接口发送验证码， 检查用户名及密码的规范性。



#### POST`/api/user/login`

|参数|类型|备注|
|----|----|----|
|username|必选|用户名/账号|
|password|必选|密码|



### Verify

#### POST`/api/verify/phone`

`application/x-www-form-urlencoded`

| 参数  | 类型 | 备注   |
| ----- | ---- | ------ |
| phone | 必选 | 手机号 |

## 

## 一般规定

无特殊说明下，返回一个json，含且仅含status(1：成功， 0：失败)， data(成功提示或者错误)