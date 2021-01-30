const to_by_password = document.querySelector('#to_by_password')
const to_by_sms = document.querySelector('#to_by_sms')
const by_password = document.querySelector('#by_password')
const by_sms = document.querySelector('#by_sms')
const username = document.querySelector('#username')
const password = document.querySelector('#password')
const phone_number = document.querySelector('#phone_number input')
const verify_code = document.querySelector('#verify_code input')
const username_status = document.querySelector('#username_status')
const password_status = document.querySelector('#password_status')
const phone_number_status = document.querySelector('#phone_number_status')
const verify_code_status = document.querySelector('#verify_code_status')
const remember_me = document.querySelector('#remember_me')
const login_button = document.querySelector('#login')
const register_button = document.querySelector('#register')

async function login(){
    let ok = true
    if (username.value.length === 0) {
        username_status.innerText =  '请输入注册时用的邮箱或者手机号呀'
        ok = false
    }
    if (password.value.length === 0) {
        password_status.innerText =  '喵，你没输入密码么？'
        ok = false
    }
    if (!ok) {
        return
    }
    const form = {
        loginName: username.value,
        password: password.value
    }
    const json = await loginReq(form)
    if (json.status) {
        alert("登录成功")
        if (remember_me.checked) {
            localStorage.setItem('token', json.token)
            localStorage.setItem('refreshToken', json.refreshToken)
        } else {
            sessionStorage.setItem('token', json.token)
            sessionStorage.setItem('refreshToken', json.refreshToken)
        }
        window.location.href = '/'
    } else {
        alert("登录失败：" + json.data)
    }
}
function loginReq(form){
    return fetch('/api/user/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: encodeJson(form)
    }).then(data => data.json())
}

function encodeJson(json){
    return Object.entries(json).map(v =>
        v.map(v =>
            v.toString()
                .replace(/=/g,'%3D')
                .replace(/&/g,'%26'))
            .join('=')
    ).join('&')
}

function init(){
    to_by_password.onclick = () => {
        to_by_password.setAttribute('class', 'chosen_tab')
        to_by_sms.removeAttribute('class')
        by_password.style.display = 'block'
        by_sms.style.display = 'none'
    }
    to_by_sms.onclick = () => {
        to_by_sms.setAttribute('class', 'chosen_tab')
        to_by_password.removeAttribute('class')
        by_password.style.display = 'none'
        by_sms.style.display = 'block'
    }
    username.oninput = () => {
        username_status.innerText = username.value.length === 0 ? '请输入注册时用的邮箱或者手机号呀' : ''
    }
    password.oninput = () => {
        password_status.innerText = password.value.length === 0 ? '喵，你没输入密码么？' : ''
    }
    phone_number.oninput = () => {
        phone_number_status.innerText = phone_number.value.length === 0 ? '手机号格式错误' : ''
    }
    verify_code.oninput = () => {
        verify_code_status.innerText = verify_code.value.length === 0 ? '短信验证码不能为空' : ''
    }
    login_button.onclick = login
    register_button.onclick = () => window.location.href = '/passport/register'
}

init()
