const username = document.querySelector('#username input')
const password = document.querySelector('#password input')
const username_hint = document.querySelector('#username_hint')
const password_hint = document.querySelector('#password_hint')
const agree_license = document.querySelector('#agree_license')
const register = document.querySelector('#register button')


document.addEventListener('keydown', checkInput)
document.addEventListener('keyup', checkInput)
agree_license.addEventListener('click', () => {
    if (agree_license.checked) {
        register.removeAttribute('disabled')
    } else {
        register.setAttribute('disabled', 'disabled')
    }
})

function checkInput() {
    if (password.value.length === 0) {
        password_hint.innerText = ''
    } else if (password.value.length < 6) {
        password_hint.innerText = '密码不能小于6个字符'
    } else if (password.value.length > 16) {
        password_hint.innerText = '密码不能大于16个字符'
    } else {
        password_hint.innerText = ''
    }
}