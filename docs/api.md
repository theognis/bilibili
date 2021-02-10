# API

## User

### `/api/user/login/pw` `POST`

* `application/x-www-form-urlencoded`
* 密码登录

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

### `/api/user/login/sms` `POST`

* `application/x-www-form-urlencoded`
* 短信登录

| 请求参数  | 类型 | 说明     |
| ------- | ---- | ------- |
| phone  |  必选  |  手机号  |
| verify_code | 必选 | 验证码  |

| 返回参数     | 说明         |
| ----------- | ------------ |
| status       | 状态码       |
| data         | 返回消息     |
| token        | 用户token    |
| refreshToken | refreshToken |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"手机号不能为空哦"` | `phone` 为空 |
| `false` | `"短信验证码不能为空"` | `verify_code` 为空 |
| `false` | `"验证码错误"` | `phone` 与 `verify_code` 不匹配 |
| `true` | `""` | `phone` 与 `verify_code` 匹配 |

### `/api/user/register` `POST`

* `application/x-www-form-urlencoded`
* 用户注册；先调用接口发送验证码， 并检查用户名及密码的规范性，确认数据符合规范之后再发送表单

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

### `/api/user/info/self` `GET`

* 获取自己的详细信息

| 请求参数 | 类型 | 说明  |
| -------- | ---- | ----- |
| token    | 必选 | token |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token` 为空 |
| `false` | `TOKEN_EXPIRED` | `token`过期 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token` 解析错误 |
| `true` | 参见下述代码 | 参数合法 |

```js
let data = {
    Uid: Number, // int64
    Username: String, // string
    Password: String, // string
    Email: String, // string
    Phone: String, // string
    Salt: String, // string
    RegDate: Date, // time.Time
    Statement: String, // string
    Exp: Number, // int64
    Coins: Number, // float64
}
```

### `/api/user/info/:uid` `GET`

* 获取 UID 为 `:uid` 的用户的个人信息
* 暂未开发

### `/api/user/avatar` `PUT`

* `multipart/form-data` 
* 修改/添加头像；

| 请求参数         | 类型 | 说明                         |
| ---------------- | ---- | ---------------------------- |
| avatar           | 必选 | 头像（二进制文件）             |
| token            | 必选 | token                        |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `false` | `"头像无效"` | `avatar` 无效或为空 |
| `false` | `"头像文件过大"` | 头像大于2MB |
| `false` | `"上传失败"` | 上传失败 |
| `true` | `"上传成功"` | 上传成功 |

### `/api/user/password` `PUT`

* `application/x-www-form-urlencoded` 
* 修改密码；

| 请求参数         | 类型 | 说明                         |
| ---------------- | ---- | ---------------------------- |
| account | 必选 | 账号 手机号/邮箱地址 |
| code    | 必选 | 验证码             |
| new_password  | 必选 | 新密码                   |
| token         | 必选 | token                    |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"账号为空"` | `account`为空 |
| `false` | `"账号不存在"` | `account`未注册 |
| `false` | `"账号无效"` | `account`不合法 |
| `false` | `"验证码为空"` | `code`为空 |
| `false` | `"验证码错误"` | `code`错误 |
| `false` | `"密码不能小于6个字符"` | `new_password` 长度少于 6 个字节 |
| `false` | `"密码不能大于16个字符"` | `new_password` 长度超过 16 个字节 |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `true` | `""` | 修改成功 |

### `/api/user/email` `PUT`

* `application/x-www-form-urlencoded` 
* 修改/添加email；先调用 `/user/info/self` 接口获取用户原先手机/邮箱，然后调用 `/verify/email` 接口发送验证码。

| 请求参数    | 类型 | 说明                         |
| ----------- | ---- | ---------------------------- |
| old_account | 必选 | 原有设备账号 手机号/邮箱地址 |
| old_code    | 必选 | 原有设备验证码               |
| new_email   | 必选 | 新email                      |
| new_code    | 必选 | 新email验证码                |
| token       | 必选 | token                        |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"原账号为空"` | `old_account`为空 |
| `false` | `"原账号不存在"` | `old_account`未注册 |
| `false` | `"原账号无效"` | `old_account`不合法 |
| `false` | `"原账号验证码为空"` | `old_code`为空 |
| `false` | `"原账号验证码错误"` | `old_code`错误 |
| `false` | `"新邮箱为空"` | `new_email`为空 |
| `false` | `"新邮箱已存在"` | `new_email`已注册 |
| `false` | `"新邮箱无效"` | `new_email`不合法 |
| `false` | `"新邮箱验证码为空"` | `new_code`为空 |
| `false` | `"新邮箱验证码错误"` | `new_code`错误 |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `true` | `""` | 修改成功 |

### `/api/user/phone` `PUT`

* `application/x-www-form-urlencoded`
* 修改phone；先调用`/user/info`接口获取用户原先手机/邮箱，然后调用`/verify/phone` 接口发送验证码。

| 请求参数    | 类型 | 说明                         |
| ----------- | ---- | ---------------------------- |
| old_account | 必选 | 原有设备账号 手机号/邮箱地址 |
| old_code    | 必选 | 原有设备验证码               |
| new_phone   | 必选 | 新手机号                     |
| new_code    | 必选 | 新手机验证码                 |
| token       | 必选 | token                        |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"原账号为空"` | `old_account`为空 |
| `false` | `"原账号不存在"` | `old_account`未注册 |
| `false` | `"原账号无效"` | `old_account`不合法 |
| `false` | `"原账号验证码为空"` | `old_code`为空 |
| `false` | `"原账号验证码错误"` | `old_code`错误 |
| `false` | `"新手机号为空"` | `new_phone`为空 |
| `false` | `"新手机号已存在"` | `new_phone`已注册 |
| `false` | `"新手机号无效"` | `new_phone`不合法 |
| `false` | `"新手机号验证码为空"` | `new_code`为空 |
| `false` | `"新手机号验证码错误"` | `new_code`错误 |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `true` | `""` | 修改成功 |

### `/api/user/username` `PUT`

* `application/x-www-form-urlencoded`
* 修改用户名，消耗6个硬币

| 请求参数      | 类型 | 说明       |
| ------------- | ---- | ---------- |
| token         | 必选 | token      |
| new_username | 必选 | 新用户名 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `false` | `"昵称不可为空"` | `new_username` 为空 |
| `false` | `"昵称太长了"` | `new_username` 大于 15 个字节 |
| `false` | `"硬币不足"` | 硬币不足 |
| `true` | `""` | 修改成功 |

### `/api/user/statement` `PUT`

* `application/x-www-form-urlencoded`
* 修改签名；如果无new_statement则更改为默认签名

| 请求参数      | 类型 | 说明       |
| ------------- | ---- | ---------- |
| token         | 必选 | token      |
| new_statement | 可选 | 新个性签名 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `false` | `"签名太长了"` | `new_username` 大于 15 个字节 |
| `true` | `""` | 修改成功 |

### `/api/user/gender` `PUT`

* `application/x-www-form-urlencoded`
* 修改性别

| 请求参数   | 类型 | 说明                                    |
| ---------- | ---- | --------------------------------------- |
| token      | 必选 | token                                   |
| new_gender | 必选 | 新性别（感觉好怪）（`F`,`M`, `O`, `N`） |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `false` | `"无效的性别"` | `new_gender` 不合法 |
| `true` | `""` | 修改成功 |

### `/api/user/birth` `PUT`

* `application/x-www-form-urlencoded`
* 修改出生日期

| 请求参数      | 类型 | 说明       |
| ------------- | ---- | ---------- |
| token         | 必选 | token      |
| new_birth    | 必选 | 新的出生日期，格式为`2006-01-02` |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `false` | `"日期格式错误"` | `new_birth` 格式错误 |
| `false` | `"出生日期无效"` | `new_birth` 晚于现在 |
| `true` | `""` | 修改成功 |

### `/api/user/check-in` `PUT`

* `application/x-www-form-urlencoded`
* 日常签到，硬币+1，经验+5

| 请求参数      | 类型 | 说明       |
| ------------- | ---- | ---------- |
| token         | 必选 | token      |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token`失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `false` | `"ALREADY_DONE"` | 已签到 |
| `true` | `"SUCCESS"` | 签到成功 |

## Verify

###  `/api/verify/token` `GET`

* 使用refreshToken获取新token

| 请求参数     | 说明         |
| ------------ | ------------ |
| refreshToken | refreshToken |


| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"refreshToken失效"` | refreshToken失效 |
| `false` | "refreshToken不正确或系统错误" | refreshToken不正确或系统错误 |
| `true` | 新的token | 成功 |

### `/api/verify/sms/general` `POST`

* `application/x-www-form-urlencoded`
* 向 `phone` 发送短信验证码

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| phone    | 必选 | 手机号 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"手机号不可为空"` | `phone` 为空 |
| `false` | `"手机号不合法"` | `phone` 不合法 |
| `true` | `""` | 发送验证码成功 |

### `/api/verify/sms/register` `POST`

* `application/x-www-form-urlencoded`
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

### `/api/verify/sms/login` `POST`

* `application/x-www-form-urlencoded`
* 登录时向 `phone` 发送短信验证码

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| phone    | 必选 | 手机号 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"手机号不可为空"` | `phone` 为空 |
| `false` | `"手机号未被注册"` | `phone` 未注册 |
| `false` | `"手机号不合法"` | `phone` 不合法 |
| `true` | `""` | 发送验证码成功 |

### `/api/verify/email` `POST`

* `application/x-www-form-urlencoded`
* 发送邮箱验证码

| 请求参数 | 类型 | 说明  |
| -------- | ---- | ----- |
| email    | 必选 | email |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"邮箱不可为空"` | `email` 为空 |
| `false` | `"邮箱格式不合法"` | `email` 不合法 |
| `true` | `""` | 发送验证码成功 |

## Check

### `/api/check/username` `GET`

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

### `/api/check/phone` `GET`

* 检验手机号是否合法

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| phone | 必选 | 手机号 |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"请告诉我你的手机号吧"` | `phone` 参数为空 |
| `false` | `"手机号已被使用"` | 手机号已被使用 |
| `true` | `""` | 手机号合法 |

## Video

### `/api/video/video` `GET`

* 请求视频；

| 请求参数      | 类型 | 说明         |
| ------------ | ---- | ----------- |
| id           | 必选 | 视频 ID      |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `ID 不存在` | `id` 不存在 |
| `true` | 参见下述代码 | 参数合法 |

```js
let data = {
	video: String, // string, 视频地址
	cover: String, // string, 封面地址
	title: String, // string, 视频标题
	channel: String, // string, 分区，字符串编号，参见`channel.md`
	label: Object, // []string, 标签，字符串切片
	description: String, // string, 简介
    author: Number, // int64, 作者 UID
    time: String, // Time, 上传时间
    views: Number, // int64, 播放次数
    likes: Number, // int64, 点赞数量
    coins: Number, // int64, 投币数量
    saves: Number, // int64, 收藏数量
    shares: Number, // int64, 分享数量
    danmakus: Object, // []Danmaku, 弹幕，弹幕切片
    comments: Object, // []Comment, 评论，评论切片
}
```

### `/api/video/video` `POST`

* `multipart/form-data` 
* 视频投稿；

| 请求参数       | 类型 | 说明                        |
| ------------ | ---- | --------------------------- |
| video        | 必选 | 视频（视频格式，二进制文件）      |
| cover        | 必选 | 封面（图片格式，二进制文件）       |
| title        | 必选 | 标题，至多 80 字              |
| channel      | 必选 | 分区，字符串编号，参见`channel.md` |
| label        | 必选 | 标签，至多 10 个，数组转 json 字符串 |
| description  | 必选 | 简介，至多 250 字             |
| token        | 必选 | token                       |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"NO_TOKEN_PROVIDED"` | `token`为空 |
| `false` | `"TOKEN_EXPIRED"` | `token` 失效 |
| `false` | `"PRASE_TOKEN_ERROR"` | `token`解析失败 |
| `false` | `"视频不可为空"` | `video`为空 |
| `false` | `"视频体积不可大于 2GB"` | `video`体积大于 2GB |
| `false` | `"封面不可为空"` | `cover`为空 |
| `false` | `"封面体积不可大于 2MB"` | `cover`体积大于 2MB |
| `false` | `"标题不可为空"` | `title`为空 |
| `false` | `"标题过长"` | `title`长度大于 80 |
| `false` | `"分区无效"` | `channel`为空、无效或过长 |
| `false` | `"标签无效"` | `label`为空或无效 |
| `false` | `"简介不可为空"` | `description`为空 |
| `false` | `"简介过长"` | `description`长度大于 250 |
| `true` | `""` | 上传成功 |

# 一般规定

如无特殊说明，则返回一个以下格式的 json：

```javascript
{
    status: true, // true：成功， false：失败
    data: "" // 提示信息
}
```