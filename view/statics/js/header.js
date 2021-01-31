const nav = document.querySelector('body>header>nav')
const user_button = document.querySelector('body>header>nav>.user_operation.logged>.user')
const user_hover = document.querySelector('body>header>.hover>.user')

let token, refreshToken

function showUserHover() {
    user_hover.style.visibility = 'visible'
    user_hover.style.opacity = '1'
    user_button.style.opacity = '0'
}
function hideUserHover() {
    user_hover.style.opacity = '0'
    user_button.style.opacity = '1'
    user_hover.style.visibility = 'hidden'
}

function initToken() {
    if (localStorage.getItem('token')){
        token = localStorage.getItem('token')
        refreshToken = localStorage.getItem('refreshToken')
    } else if (sessionStorage.getItem('token')){
        token = sessionStorage.getItem('token')
        refreshToken = sessionStorage.getItem('refreshToken')
    }
}
function initHeader() {
    if (token) {
        nav.setAttribute('class', 'logged')
        user_button.onmouseover = showUserHover
        user_button.onmouseout = hideUserHover
        user_hover.onmouseover = showUserHover
        user_hover.onmouseout = hideUserHover
    } else {
        nav.setAttribute('class', 'not_logged')
    }
}
function locateHover() {
    user_hover.style.left = user_button.offsetLeft - 140 + 'px'
}

/*async function info(){
    let json = await fetch('/api/user/info/self?token=' + refreshToken, {
        method: 'GET'
    }).then(data => data.json())
    console.log(json)
}*/

initToken()
initHeader()
locateHover()