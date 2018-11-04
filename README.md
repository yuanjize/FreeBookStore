##API



**主机ip**: 207.148.89.29
**端口**:8080

###注册：
方法:`POST`
路径:`/register`
参数:

| 参数名称| 参数意义| 
| ------ | ------ |
| username | 用户名  | 
| password1 | 登录密码 | 
| nickname | 用户昵称 | 
| email | 用户邮箱 |
| captcha| 验 证 码 |

返回值(json)：
```json
	{
	   "msg":   "email err",
	   "label": "none",
	   "mstatus":"0"
	}
```
| 字段名称| 字段意义| 
| ------ | ------ |
| mstatus | 1代表有错误，0代表注册成功  | 
| msg | 注册成功或者失败的信息 |
| label | 显示在哪个控件上，比如邮箱错误就显示在表单的邮箱控件旁边，但注册成功的时候label是none，这个时候信息显示在页面某个位置 |


###登录
方法:`POST`
路径:`/login`

参数:

| 参数名称| 参数意义| 
| ------ | ------ |
| account | 用户名  | 
| passwd | 登录密码 | 

返回值(json)：
```json
	{
	   "msg":   "Login Success",
	   "mstatus":"0"
	}
```
| 字段名称| 字段意义| 
| ------ | ------ |
| mstatus | 1代表登录失败，0代表登录成功  | 
| msg | 登录成功或者失败的信息 |


###验证码
方法:`GET`
路径:`/vertify`

无参数
返回值(json)：
```json
{
    "vertifypng":base64编码的图片字符串,
}
```
返回的是base64编码的图片


Yuanjize1996!!




 

