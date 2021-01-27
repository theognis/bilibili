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

## 一般规定

如无特殊说明，则返回一个以下格式的 json：

```javascript
{
    status: 1, // 1：成功， 0：失败
    data: "" // 成功提示或者错误
}
```



