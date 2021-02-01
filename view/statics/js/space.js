const tab_home = document.querySelector('#tab_home')
const tab_moments = document.querySelector('#tab_moments')
const tab_post = document.querySelector('#tab_post')
const tab_underline = document.querySelector('#tab_underline')
const queries = queryToJson()

function moveUnderline(left = getInitTabLeft()) {
    tab_underline.style.left = left
}
function getInitTabLeft() {
    switch (queries.tab) {
        case 'home': return '15px';
        case 'moments': return '96px';
        case 'post': return '177px';
        default: return '15px';
    }
}

function queryToJson() {
    const queryStr = window.location.search.substring(1)
    const json = {}
    if (queryStr !== '') {
        queryStr.split('&').map(v => v.split('=')).forEach(v => json[v[0]] = v[1])
    }
    return json
}

function init() {
    tab_home.addEventListener('mouseover', () => moveUnderline('15px'))
    tab_moments.addEventListener('mouseover', () => moveUnderline('96px'))
    tab_post.addEventListener('mouseover', () => moveUnderline('177px'))
    tab_home.addEventListener('mouseout', () => moveUnderline())
    tab_moments.addEventListener('mouseout', () => moveUnderline())
    tab_post.addEventListener('mouseout', () => moveUnderline())

    tab_home.addEventListener('click', function() {
        queries.tab = 'home'
        history.pushState(null ,null ,`/space/?${jsonToQuery(queries)}`)
        tab_home.classList.add('now_tab')
        tab_moments.classList.remove('now_tab')
        tab_post.classList.remove('now_tab')
    })
    tab_moments.addEventListener('click', function() {
        queries.tab = 'moments'
        history.pushState(null ,null ,`/space/?${jsonToQuery(queries)}`)
        tab_home.classList.remove('now_tab')
        tab_moments.classList.add('now_tab')
        tab_post.classList.remove('now_tab')
    })
    tab_post.addEventListener('click', function() {
        queries.tab = 'post'
        history.pushState(null ,null ,`/space/?${jsonToQuery(queries)}`)
        tab_home.classList.remove('now_tab')
        tab_moments.classList.remove('now_tab')
        tab_post.classList.add('now_tab')
    })

    moveUnderline()
}

init()