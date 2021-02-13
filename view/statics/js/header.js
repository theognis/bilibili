function renderHeader() {
    const logged = localStorage.getItem('token') || sessionStorage.getItem('token')
    const home = location.pathname === '/'
    document.write(
        `<header>
            <nav class="${logged ? 'logged' : 'not_logged'}">
            <div id="navigator">
                <a id="nav_homepage" href="/">
                    ${home ? 
                        `<span class="bilifont">&#xE72B;</span>` :
                        `<img src="/statics/images/bilibili.svg">`
                    }
                    <span>主站</span>
                </a>
                <a href="javascript:alert('整天就知道看番！')">番剧</a>
                <a href="javascript:alert('整天就知道打游戏！')">游戏中心</a>
                <a href="javascript:alert('整天就知道看直播！')">直播</a>
                <a href="javascript:alert('整天就知道买东西！')">会员购</a>
                <a href="javascript:alert('整天就知道看漫画！')">漫画</a>
                <a href="javascript:alert('整天就知道看比赛！')">赛事</a>
                <a id="nav_download" href="https://github.com/Override0330/WinterExam-Android-2020"
                    target="_blank"><span class="bilifont">&#xE72D;</span>下载APP</a>
            </div>
            <form id="search">
                <input>
                <button class="bilifont">&#xE72C;</button>
            </form>
            <div class="user_operation">
                ${logged ? 
                    `<img class="user">
                    <a>大会员</a>
                    <a>消息</a>
                    <a>动态</a>
                    <a>收藏</a>
                    <a>历史</a>
                    <a>创作中心</a>` :
                    `<img src="/statics/images/akari.jpg">
                    <a href="/account/login">登录</a>
                    <a href="/account/register">注册</a>`
                }
            </div>
            <a id="post" href="/video/upload">投稿</a>
            </nav>
            ${logged ?
                `<div class="hover">
                    <div class="user" style="visibility: hidden;opacity: 0;">
                        <img class="avatar">
                        <span class="username">admin</span>
                        <div class="level_content">
                            <div class="info">
                                <span class="level">等级 0</span>
                                <span class="exp">0 / 0</span>
                            </div>
                            <div class="bar"><div class="progress"></div></div>
                        </div>
                        <div class="money">
                            <a class="coin"><img src="/statics/images/coin.svg"><span>0</span></a>
                            <a class="b-coin"><img src="/statics/images/b-coin.svg"><span>0</span></a>
                        </div>
                        <div class="links">
                            <a class="setting" href="/account/setting">个人中心</a>
                        </div>
                        <div class="logout">
                            <span>登出</span>
                        </div>
                    </div>
                </div>` : ``
            }
        </header>`
    )
}

renderHeader()

const user_button = document.querySelector('body>header>nav>.user_operation>.user')
const user_hover = document.querySelector('body>header>.hover>.user')
const pre_avatar = document.querySelector('body>header>nav>.user_operation>img')
const uh_avatar = document.querySelector('body>header>.hover>.user>.avatar')
const uh_username = document.querySelector('body>header>.hover>.user>.username')
const uh_level = document.querySelector('body>header>.hover>.user>.level_content>.info>.level')
const uh_exp = document.querySelector('body>header>.hover>.user>.level_content>.info>.exp')
const uh_progress = document.querySelector('body>header>.hover>.user>.level_content>.bar>.progress')
const uh_coin = document.querySelector('body>header>.hover>.user>.money>.coin>span')
const uh_bCoin = document.querySelector('body>header>.hover>.user>.money>.b-coin>span')
const logout_button = document.querySelector('body>header>.hover>.user>.logout>span')

const user = { token: '', refreshToken: '', data: {} }

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
        user_hover.style.left = user_button.offsetLeft - 140 + 'px'
        user_button.onmouseover = showUserHover
        user_button.onmouseout = hideUserHover
        user_hover.onmouseover = showUserHover
        user_hover.onmouseout = hideUserHover
        logout_button.onclick = logout
        initUserHover()
    }
}
function initUserHover() {
    if (!(user.data instanceof Object)) {
        return
    }
    const info = user.data
    pre_avatar.src = info.Avatar
    uh_avatar.src = info.Avatar
    uh_username.innerText = info.Username
    uh_level.innerText = '等级 ' + getLevel(info.Exp)
    uh_exp.innerText = info.Exp + ' / ' + getMaxExp(info.Exp)
    uh_progress.style.width = info.Exp / getMaxExp(info.Exp) * 100 + '%'
    uh_coin.innerText = info.Coins
    uh_bCoin.innerText = info.BCoins
}

function logout() {
    localStorage.removeItem('token')
    localStorage.removeItem('refreshToken')
    sessionStorage.removeItem('token')
    sessionStorage.removeItem('refreshToken')
    window.location.reload()
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
function isRemembered() {
    return !!localStorage.getItem('token');
}
function jsonToQuery(json) {
    return Object.entries(json).map(v =>
        v.map(v =>
            v.toString()
                .replace(/=/g, '%3D')
                .replace(/&/g, '%26'))
            .join('=')
    ).join('&')
}

async function initToken() {
    if (localStorage.getItem('token')) {
        user.token = localStorage.getItem('token')
        user.refreshToken = localStorage.getItem('refreshToken')
    } else if (sessionStorage.getItem('token')) {
        user.token = sessionStorage.getItem('token')
        user.refreshToken = sessionStorage.getItem('refreshToken')
    }
    if (user.token !== '') {
        setInterval(refreshToken, 60000)
        await refreshToken()
        await checkIn()
        const res = await getInfo()
        if (!res.status) {
            console.log('Failed to get info: ', res.data)
            return
        }
        user.data = res.data
    }
}
function refreshToken() {
    return fetch('/api/verify/token?refreshToken=' + user.refreshToken, {
        method: 'GET'
    })
        .then(data => data.json())
        .then(RTokenRes => {
            if (RTokenRes.status) {
                user.token = RTokenRes.data
                if (isRemembered())
                    localStorage.setItem('token', RTokenRes.data)
                else
                    sessionStorage.setItem('token', RTokenRes.data)
                console.log('刷新 Token 成功')
            } else {
                console.log('刷新 Token 失败: ', RTokenRes.data)
            }
        })
}
function checkIn() {
    return fetch('/api/user/check-in', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: jsonToQuery({ token: user.token })
    })
        .then(data => data.json())
        .then(checkInRes => {
            if (checkInRes.status) {
                console.log('签到成功')
            } else {
                console.log('签到失败: ', checkInRes.data)
            }
        })
}
function getInfo() {
    return fetch('/api/user/info/self?token=' + user.token, {
        method: 'GET'
    }).then(data => data.json())
}

const tokenInit = initToken()
tokenInit.then(() => initHeader())
