const to_by_password = document.querySelector('#to_by_password')
const to_by_sms = document.querySelector('#to_by_sms')
const by_password = document.querySelector('#by_password')
const by_sms = document.querySelector('#by_sms')
const username_input = document.querySelector('#username')
const password_input = document.querySelector('#password')
const phone_number_input = document.querySelector('#phone_number input')
const verify_code_input = document.querySelector('#verify_code input')
const username_status = document.querySelector('#username_status')
const password_status = document.querySelector('#password_status')
const phone_number_status = document.querySelector('#phone_number_status')
const verify_code_status = document.querySelector('#verify_code_status')
const remember_me = document.querySelector('#remember_me')
const send_message_button = document.querySelector('#verify_code>button')
const login_button = document.querySelector('#login')
const register_button = document.querySelector('#register')


function login(){
    let chosen_tab = document.querySelector('.chosen_tab').getAttribute('id')
    chosen_tab === 'to_by_password' ? loginPw() : loginSms()
}
async function loginPw(){
    let ok = true
    if (username_input.value.length === 0) {
        username_status.innerText =  '请输入注册时用的邮箱或者手机号呀'
        ok = false
    }
    if (password_input.value.length === 0) {
        password_status.innerText =  '喵，你没输入密码么？'
        ok = false
    }
    if (!ok) {
        return
    }
    const form = {
        loginName: username_input.value,
        password: password_input.value
    }
    const json = await loginReq('pw', form)
    handleLoginRes(json)
}
async function loginSms(){
    let ok = true
    if (phone_number_input.value.length === 0) {
        phone_number_status.innerText =  '手机号不能为空哦'
        ok = false
    }
    if (verify_code_input.value.length === 0) {
        verify_code_status.innerText =  '短信验证码不能为空'
        ok = false
    }
    if (!ok) {
        return
    }
    const form = {
        phone: phone_number_input.value,
        verify_code: verify_code_input.value
    }
    const json = await loginReq('sms', form)
    handleLoginRes(json)
}
async function sendMessage(){
    if (phone_number_input.value === '') {
        phone_number_status.innerText =  '手机号不能为空哦'
        return
    }
    send_message_button.innerText = '获取中...'
    send_message_button.setAttribute('disabled', 'disabled')
    let json = await sendMessageReq(phone_number_input.value)
    if(json.status) {
        send_message_button.innerText = '已获取'
    } else {
        alert('验证码获取失败：' + json.data)
        send_message_button.innerText = '点击获取'
        send_message_button.removeAttribute('disabled')
    }
}
function loginReq(type, form){
    return fetch('/api/user/login/' + type, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: jsonToQuery(form)
    }).then(data => data.json())
}
function sendMessageReq(phone){
    return fetch('/api/verify/sms/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: jsonToQuery({phone: phone})
    }).then(data => data.json())
}
function handleLoginRes(json) {
    if (json.status) {
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
    username_input.oninput = () => {
        username_status.innerText = username_input.value.length === 0 ? '请输入注册时用的邮箱或者手机号呀' : ''
    }
    password_input.oninput = () => {
        password_status.innerText = password_input.value.length === 0 ? '喵，你没输入密码么？' : ''
    }
    phone_number_input.oninput = () => {
        phone_number_status.innerText = phone_number_input.value.length === 0 ? '手机号格式错误' : ''
    }
    verify_code_input.oninput = () => {
        verify_code_status.innerText = verify_code_input.value.length === 0 ? '短信验证码不能为空' : ''
    }
    send_message_button.onclick = sendMessage
    login_button.onclick = login
    register_button.onclick = () => window.location.href = '/account/register'
}

init()
