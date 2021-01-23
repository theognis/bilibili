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
const login = document.querySelector('#login')
const register = document.querySelector('#register')

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

username.addEventListener('change', () => {
    username_status.innerText = username.value.length === 0 ? '请输入注册时用的邮箱或者手机号呀' : ''
})
password.addEventListener('change', () => {
    password_status.innerText = password.value.length === 0 ? '喵，你没输入密码么？' : ''
})
phone_number.addEventListener('change', () => {
    phone_number_status.innerText = phone_number.value.length === 0 ? '手机号格式错误' : ''
})
verify_code.addEventListener('change', () => {
    verify_code_status.innerText = verify_code.value.length === 0 ? '短信验证码不能为空' : ''
})
register.addEventListener('click', () => {
    console.log("???")
    window.location.href = '/passport/register'
    console.log("???")
})
