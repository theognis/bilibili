const nav = document.querySelector('body>header>nav')

let token

if (localStorage.getItem('token')){
    token = localStorage.getItem('token')
} else if (sessionStorage.getItem('token')){
    token = sessionStorage.getItem('token')
}

nav.setAttribute('class', token ? 'logged' : 'not_logged')