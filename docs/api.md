## API

### User

#### `/api/user/login`
| Key          | Value                               |
| ------------ | ----------------------------------- |
| url          | `/api/user/login`                   |
| content-type | `application/x-www-form-urlencoded` |
| method       | POST                                |
| description  | 登录                                |

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

#### `/api/user/register`

| Key          | Value                                                        |
| ------------ | ------------------------------------------------------------ |
| url          | `/api/user/register`                                         |
| content-type | `application/x-www-form-urlencoded`                          |
| method       | POST                                                         |
| description  | 先调用接口发送验证码， 并检查用户名及密码的规范性，确认数据符合规范之后再发送表单 |

| 请求参数    | 类型 | 说明        |
| ----------- | ---- | ----------- |
| username    | 必选 | 用户名/账号 |
| password    | 必选 | 密码        |
| phone       | 必选 | 手机号      |
| verify_code | 必选 | 验证码      |

#### `/api/user/hasUsername`

| Key          | Value                               |
| ------------ | ----------------------------------- |
| url          | `/api/user/hasUsername`            |
| content-type | `application/x-www-form-urlencoded` |
| method       | POST                                |
| description  | 检验用户名是否存在                     |

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| username | 必选 | 用户名 |

#### `/api/user/info`

| key         | value              |
| ----------- | ------------------ |
| url         | /api/user/info     |
| method      | GET                |
| description | 获取自己的详细信息 |

| 请求参数 | 类型 | 说明  |
| -------- | ---- | ----- |
| token    | 必选 | token |

返回参数: 有且仅有以下结构体userinfo中的字段 

https://github.com/theognis/bilibili/blob/master/model/User.go 

#### `/api/user/email`

| key          | value                               |
| ------------ | ----------------------------------- |
| url          | /api/user/email                     |
| content-type | `application/x-www-form-urlencoded` |
| method       | PUT                                 |
| description  | 修改/添加email                      |

##### PS

先调用/user/info接口获取用户原先手机/邮箱，然后调用/verify/phone or /verify/email 接口发送验证码。

| 请求参数         | 类型 | 说明                         |
| ---------------- | ---- | ---------------------------- |
| original_address | 必选 | 原有设备账号 手机号/邮箱地址 |
| original_code    | 必选 | 原有设备验证码               |
| new_email        | 必选 | 新email                      |
| new_code         | 必选 | 新email验证码                |
| token            | 必选 | token                        |

#### `/api/user/phone`

修改电话，信息同上

| 请求参数         | 类型 | 说明                         |
| ---------------- | ---- | ---------------------------- |
| original_address | 必选 | 原有设备账号 手机号/邮箱地址 |
| original_code    | 必选 | 原有设备验证码               |
| new_phone        | 必选 | 新手机号                     |
| new_code         | 必选 | 新手机验证码                 |
| token            | 必选 | token                        |



#### `/api/user/statement`

| key          | value                               |
| ------------ | ----------------------------------- |
| url          | /api/user/statement                 |
| content-type | `application/x-www-form-urlencoded` |
| method       | PUT                                 |
| description  | 修改个性签名                        |

| 请求参数      | 类型 | 说明       |
| ------------- | ---- | ---------- |
| token         | 必选 | token      |
| new_statement | 可选 | 新个性签名 |

##### PS

如果无new_statement则更改为默认签名

### Verify

####  `/api/verify/token`

| Key          | Value                       |
| ------------ | --------------------------- |
| url          | `/api/verify/token`         |
| method       | GET                         |
| description  | 使用refreshToken获取新token |

| 请求参数     | 说明         |
| ------------ | ------------ |
| refreshToken | refreshToken |

| 返回参数 | 说明                                                         |
| -------- | ------------------------------------------------------------ |
| data     | 成功则为新的token，若refreshToken失效则为 "refreshToken失效" |
| status   | 状态码                                                       |

#### `/api/verify/phone`

| Key          | Value                               |
| ------------ | ----------------------------------- |
| url          | `/api/verify/phone`                 |
| content-type | `application/x-www-form-urlencoded` |
| method       | POST                                |
| description  | 发送短信验证码                      |

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| phone    | 必选 | 手机号 |

#### `/api/verify/email`

| key          | value                               |
| ------------ | ----------------------------------- |
| url          | /api/verify/email                   |
| content-type | `application/x-www-form-urlencoded` |
| method       | POST                                |
| description  | 发送邮箱验证码                      |

| 请求参数 | 类型 | 说明  |
| -------- | ---- | ----- |
| email    | 必选 | email |





## 一般规定

如无特殊说明，则返回一个以下格式的 json：

```javascript
{
    status: true, // true：成功， false：失败
    data: "" // 提示信息
}
```



