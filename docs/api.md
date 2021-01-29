## API

### User

#### `/api/user/login` `POST`

* Content-Type:  `application/x-www-form-urlencoded`
* 登录

| 请求参数  | 类型 | 说明                |
| --------- | ---- | ------------------- |
| loginName | 必选 | 登录名（手机/邮箱） |
| password  | 必选 | 密码                |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |
| token        | 用户token    |
| refreshToken | refreshToken |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"请输入注册时用的邮箱或者手机号呀"` | `loginName` 为空 |
| `false` | `"喵，你没输入密码么？"` | `password` 为空 |
| `false` | `"用户名或密码错误"` | `loginName` 不存在 |
| `false` | `"用户名或密码错误"` | `loginName` 与 `password` 不匹配 |
| `true` | `""` | `loginName` 与 `password` 匹配 |

#### `/api/user/register` `POST`

* Content-Type: `application/x-www-form-urlencoded`
* 先调用接口发送验证码， 并检查用户名及密码的规范性，确认数据符合规范之后再发送表单

| 请求参数    | 类型 | 说明        |
| ----------- | ---- | ----------- |
| username    | 必选 | 用户名/账号 |
| password    | 必选 | 密码        |
| phone       | 必选 | 手机号      |
| verify_code | 必选 | 验证码      |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"用户名不能为空"` | `username` 为空 |
| `false` | `"用户名太长了"` | `username` 长度超过 15 个字节 |
| `false` | `"密码不能小于6个字符"` | `password` 长度少于 6 个字节 |
| `false` | `"密码不能大于16个字符"` | `password` 长度超过 16 个字节 |
| `false` | `"手机号不可为空"` | `phone` 为空 |
| `false` | `"该手机号已经被注册"` | `phone` 已被注册 |
| `false` | `"请输入验证码"` | `verify_code` 为空 |
| `false` | `"未发送验证码"` | `verify_code` 无对应验证码 |
| `false` | `"验证码错误"` | `verify_code` 与对应验证码不符 |
| `true` | `"注册成功！"` | 参数合法 |

#### `/api/user/info/self` `GET`

* 获取自己的详细信息

| 请求参数 | 类型 | 说明  |
| -------- | ---- | ----- |
| token    | 必选 | token |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"token 不可为空"` | `token` 为空 |
| `false` | `"token 不合法"` | `token` 不合法 |
| `true` | 以下 json | 参数合法 |

```js
{
    Uid: Number, // int64
    Username: String, // string
    Password: String, // string
    Email: String, // string
    Phone: String, // string
    Salt: String, // string
    RegDate: Date, // time.Time
    Statement: String, // string
}
```

#### `/api/user/info/:uid` `GET`

* 获取 UID 为 `:uid` 的用户的个人信息
* 暂未开发

#### `/api/user/email` `PUT`

* Content-Type: `application/x-www-form-urlencoded` 
* 修改/添加email；先调用 `/user/info/self` 接口获取用户原先手机/邮箱，然后调用 `/verify/email` 接口发送验证码。

| 请求参数         | 类型 | 说明                         |
| ---------------- | ---- | ---------------------------- |
| original_address | 必选 | 原有设备账号 手机号/邮箱地址 |
| original_code    | 必选 | 原有设备验证码               |
| new_email        | 必选 | 新email                      |
| new_code         | 必选 | 新email验证码                |
| token            | 必选 | token                        |

| status | data | 说明   |
| -------- | ---- | ------ |
|  |  |  |

#### `/api/user/phone` `PUT`

* Content-Type: `application/x-www-form-urlencoded`
* 修改phone；先调用`/user/info`接口获取用户原先手机/邮箱，然后调用`/verify/phone` 接口发送验证码。

| 请求参数         | 类型 | 说明                         |
| ---------------- | ---- | ---------------------------- |
| original_address | 必选 | 原有设备账号 手机号/邮箱地址 |
| original_code    | 必选 | 原有设备验证码               |
| new_phone        | 必选 | 新手机号                     |
| new_code         | 必选 | 新手机验证码                 |
| token            | 必选 | token                        |

| status | data | 说明   |
| -------- | ---- | ------ |
|  |  |  |

#### `/api/user/statement` `PUT`

* Content-Type `application/x-www-form-urlencoded`
* 修改个性签名；如果无new_statement则更改为默认签名

| 请求参数      | 类型 | 说明       |
| ------------- | ---- | ---------- |
| token         | 必选 | token      |
| new_statement | 可选 | 新个性签名 |

| status | data | 说明   |
| -------- | ---- | ------ |
|  |  |  |

### Verify

####  `/api/verify/token` `GET`

* 使用refreshToken获取新token

| 请求参数     | 说明         |
| ------------ | ------------ |
| refreshToken | refreshToken |


| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"refreshToken失效"` | refreshToken失效 |
| `true` | 新的token | 成功 |

#### `/api/verify/sms/general` `POST`

* Content-Type: `application/x-www-form-urlencoded`
* 向 `phone` 发送短信验证码

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| phone    | 必选 | 手机号 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"手机号不可为空"` | `phone` 为空 |
| `false` | `"手机号不合法"` | `phone` 不合法 |
| `true` | `""` | 发送验证码成功 |

#### `/api/verify/sms/register` `POST`

* Content-Type: `application/x-www-form-urlencoded`
* 注册时向 `phone` 发送短信验证码

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| phone    | 必选 | 手机号 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"手机号不可为空"` | `phone` 为空 |
| `false` | `"手机号已被使用"` | `phone` 已被使用 |
| `false` | `"手机号不合法"` | `phone` 不合法 |
| `true` | `""` | 发送验证码成功 |

#### `/api/verify/email` `POST`

* Content-Type: `application/x-www-form-urlencoded`
* 发送邮箱验证码

| 请求参数 | 类型 | 说明  |
| -------- | ---- | ----- |
| email    | 必选 | email |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"邮箱不可为空"` | `email` 为空 |
| `false` | `"邮箱格式不合法"` | `email` 不合法 |
| `true` | `""` | 发送验证码成功 |

### Check

#### `/api/check/username` `GET`

* 检验用户名是否合法

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| username | 必选 | 用户名 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"请告诉我你的昵称吧"` | `username` 为空 |
| `false` | `"昵称过长"` | `username` 长度超过 14 个字节 |
| `false` | `"昵称已存在"` | `username` 已被使用 |
| `true` | `""` | `username` 合法 |

#### `/api/check/phone` `GET`

* 检验手机号是否合法

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| phone | 必选 | 手机号 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"请告诉我你的手机号吧"` | `phone` 参数为空 |
| `false` | `"手机号已被使用"` | 手机号已被使用 |
| `true` | `""` | 手机号合法 |

## 一般规定

如无特殊说明，则返回一个以下格式的 json：

```javascript
{
    status: true, // true：成功， false：失败
    data: "" // 提示信息
}
```