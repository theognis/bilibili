const tab_home = document.querySelector('#tab_home')
const tab_moments = document.querySelector('#tab_moments')
const tab_post = document.querySelector('#tab_post')
const tab_underline = document.querySelector('#tab_underline')
const uid = location.pathname.split('/')[2]
let tab_underline_left = tab_underline.style.left

tab_home.addEventListener('mouseover', () => moveUnderline('15px'))
tab_moments.addEventListener('mouseover', () => moveUnderline('96px'))
tab_post.addEventListener('mouseover', () => moveUnderline('177px'))
tab_home.addEventListener('mouseout', () => moveUnderline())
tab_moments.addEventListener('mouseout', () => moveUnderline())
tab_post.addEventListener('mouseout', () => moveUnderline())

tab_home.addEventListener('click', function() {
    history.pushState(null ,null ,`/space/${uid}`)
    tab_home.setAttribute('class', 'now_tab')
    tab_moments.removeAttribute('class')
    tab_post.removeAttribute('class')
    tab_underline_left = '15px'
})

tab_moments.addEventListener('click', function() {
    history.pushState(null ,null ,`/space/${uid}/moments`)
    tab_home.removeAttribute('class')
    tab_moments.setAttribute('class', 'now_tab')
    tab_post.removeAttribute('class')
    tab_underline_left = '96px'
})

tab_post.addEventListener('click', function() {
    history.pushState(null ,null ,`/space/${uid}/post`)
    tab_home.removeAttribute('class')
    tab_moments.removeAttribute('class')
    tab_post.setAttribute('class', 'now_tab')
    tab_underline_left = '177px'
})

function moveUnderline(left = tab_underline_left) {
    tab_underline.style.left = left
}