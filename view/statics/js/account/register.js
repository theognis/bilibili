const username_input = document.querySelector('#username input')
const password_input = document.querySelector('#password input')
const username_hint = document.querySelector('#username_hint')
const password_hint = document.querySelector('#password_hint')
const agree_license = document.querySelector('#agree_license')
const phone_number_input = document.querySelector('#phone_number input')
const phone_number_hint = document.querySelector('#phone_number_hint')
const verify_code_button = document.querySelector('#verify_code button')
const verify_code_input = document.querySelector('#verify_code input')
const register_button = document.querySelector('#register button')


username_input.addEventListener('input', checkUsername)
password_input.addEventListener('input', checkPassword)
phone_number_input.addEventListener('input', checkPhone)
agree_license.addEventListener('click', () => {
    if (agree_license.checked) {
        register_button.removeAttribute('disabled')
    } else {
        register_button.setAttribute('disabled', 'disabled')
    }
})
verify_code_button.addEventListener('click', sendMessage)
register_button.addEventListener('click', register)

async function checkUsername(){
    username_hint.innerText = (await checkUsernameReq(username_input.value)).data
}
async function checkPhone(){
    phone_number_hint.innerText = (await checkPhoneReq(phone_number_input.value)).data
}
async function sendMessage(){
    verify_code_button.innerText = '获取中...'
    verify_code_button.setAttribute('disabled', 'disabled')
    let json = await sendMessageReq(phone_number_input.value)
    if(json.status) {
        verify_code_button.innerText = '已获取'
    } else {
        alert('验证码获取失败：' + json.data)
        verify_code_button.innerText = '点击获取'
        verify_code_button.removeAttribute('disabled')
    }
}
async function register(){
    const username = username_input.value;
    const password = password_input.value;
    const phone_number = phone_number_input.value;
    const verify_code = verify_code_input.value;
    let json = await registerReq({
        username: username,
        password: password,
        phone: phone_number,
        verify_code: verify_code
    })
    if(json.status){
        alert("注册成功")
        window.location.href = '/account/login'
    } else {
        alert("注册失败：" + json.data)
    }
}
function checkPassword() {
    if (getStrLength(password_input.value) < 6) {
        password_hint.innerText = '密码不能小于6个字符'
    } else if (getStrLength(password_input.value) > 16) {
        password_hint.innerText = '密码不能大于16个字符'
    } else {
        password_hint.innerText = ''
    }
}

function registerReq(form){
    return fetch('/api/user/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: jsonToQuery(form)
    }).then(data => data.json())
}
function sendMessageReq(phone){
    return fetch('/api/verify/sms/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: jsonToQuery({phone: phone})
    }).then(data => data.json())
}
function checkUsernameReq(username){
    return fetch('/api/check/username?username=' + username, {
        method: 'GET',
    }).then(data => data.json())
}
function checkPhoneReq(phone){
    return fetch('/api/check/phone?phone=' + phone, {
        method: 'GET',
    }).then(data => data.json())
}

function getStrLength(string){
    let length = 0;
    for(let i = 0; i < string.length; i++){
        length += getCharLength(string[i])
    }
    return length;
}
function getCharLength(char){
    if (char.charCodeAt(0) <= 128){
        return 1
    } else if (char.charCodeAt(0) <= 2048){
        return 2
    } else if (char.charCodeAt(0) <= 65536){
        return 3
    } else {
        throw new Error(`chars: ${char}, char: ${char[0]}, code: ${char.charCodeAt(0)}`)
    }
}

