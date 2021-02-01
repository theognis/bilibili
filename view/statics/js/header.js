const nav = document.querySelector('body>header>nav')
const user_button = document.querySelector('body>header>nav>.user_operation.logged>.user')
const user_hover = document.querySelector('body>header>.hover>.user')
const uh_username = document.querySelector('body>header>.hover>.user>.username')
const uh_level = document.querySelector('body>header>.hover>.user>.level_content>.info>.level')
const uh_exp = document.querySelector('body>header>.hover>.user>.level_content>.info>.exp')
const uh_progress = document.querySelector('body>header>.hover>.user>.level_content>.bar>.progress')
const uh_coin = document.querySelector('body>header>.hover>.user>.money>.coin>span')
const logout_button = document.querySelector('body>header>.hover>.user>.logout>span')

const user = { token: '', refreshToken: '' }

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
function initHeader() {
    if (user.token) {
        nav.setAttribute('class', 'logged')
        user_button.onmouseover = showUserHover
        user_button.onmouseout = hideUserHover
        user_hover.onmouseover = showUserHover
        user_hover.onmouseout = hideUserHover
        logout_button.onclick = logout
        initUserHover()
    } else {
        nav.setAttribute('class', 'not_logged')
    }
}
async function initUserHover() {
    const info = (await getInfo()).data
    uh_username.innerText = info.Username
    uh_level.innerText = '等级 ' + getLevel(info.Exp)
    uh_exp.innerText = info.Exp + ' / ' + getMaxExp(info.Exp)
    uh_progress.style.width = info.Exp / getMaxExp(info.Exp) * 100 + '%'
    uh_coin.innerText = info.Coins
}
function locateHover() {
    user_hover.style.left = user_button.offsetLeft - 140 + 'px'
}

function getLevel(exp) {
    if (exp < 200) {
        return 1
    } else if (exp < 1500) {
        return 2
    } else if (exp < 4500) {
        return 3
    } else if (exp < 10800) {
        return 4
    } else if (exp < 28800) {
        return 5
    } else {
        return 6
    }
}
function getMaxExp(exp) {
    if (exp < 200) {
        return 200
    } else if (exp < 1500) {
        return 1500
    } else if (exp < 4500) {
        return 4500
    } else if (exp < 10800) {
        return 10800
    } else if (exp < 28800) {
        return 28800
    } else {
        return 114514
    }
}
function logout() {
    localStorage.removeItem('token')
    localStorage.removeItem('refreshToken')
    sessionStorage.removeItem('token')
    sessionStorage.removeItem('refreshToken')
    window.location.reload()
}

function initToken() {
    if (localStorage.getItem('token')){
        user.token = localStorage.getItem('token')
        user.refreshToken = localStorage.getItem('refreshToken')
    } else if (sessionStorage.getItem('token')){
        user.token = sessionStorage.getItem('token')
        user.refreshToken = sessionStorage.getItem('refreshToken')
    }
}
function isRemembered() {
    return !!localStorage.getItem('token');
}
function refreshToken(refreshToken = user.refreshToken) {
    return fetch('/api/verify/token?refreshToken=' + user.refreshToken, {
        method: 'GET'
    })
        .then(data => data.json())
        .then(json => {
            if(json.status) {
                user.token = json.data
                if (isRemembered()) {
                    localStorage.setItem('token', json.data)
                } else {
                    sessionStorage.setItem('token', json.data)
                }
            } else {
                console.log('Failed to refresh token: ', json.data)
            }
        })
}

async function getInfo(){
    const res = (await fetch('/api/user/info/self?token=' + user.token, {
        method: 'GET'
    }).then(data => data.json()))
    if (res.data === 'TOKEN_EXPIRED') {
        await refreshToken()
        return (await fetch('/api/user/info/self?token=' + user.token, {
            method: 'GET'
        }).then(data => data.json()))
    }
    return res
}

initToken()
initHeader()
locateHover()