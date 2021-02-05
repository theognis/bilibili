const video_wrapper = document.querySelector('#video_wrapper')
const video = document.querySelector('video')
const controls_process = document.querySelector('#video_wrapper>.controls>.process')
const controls_process_played = document.querySelector('#video_wrapper>.controls>.process>.played')
const controls_process_icon = document.querySelector('#video_wrapper>.controls>.process>.icon')
const controls_play = document.querySelector('#video_wrapper>.controls>.left>img')
const controls_location = document.querySelector('#video_wrapper>.controls>.left>.location')
const danmaku_area = document.querySelector('#video_wrapper>.danmaku_area')
const danmaku_switch = document.querySelector('#video_wrapper>.bottom>.control>.switch')
const danmaku_font_settings = document.querySelector('#video_wrapper>.hover>.font_settings')
const danmaku_font_switch = document.querySelector('#video_wrapper>.bottom>.control>.send>.edit>div')
const danmaku_send = document.querySelector('#video_wrapper>.bottom>.control>.send button')
const danmaku_input = document.querySelector('#video_wrapper>.bottom>.control>.send input')
const scroll_type = document.querySelector('.font_settings>.type>.scroll')
const top_type = document.querySelector('.font_settings>.type>.top')
const bottom_type = document.querySelector('.font_settings>.type>.bottom')
const recommend = document.querySelector('#recommend')
const color_input = document.querySelector('.font_settings>.color>.edit>input')
const color_preview = document.querySelector('.font_settings>.color>.edit>div')
const color_list = document.querySelector('.font_settings>.color>.list').children
let chosen_danmaku_type = 'scroll'

class Danmaku {
    constructor(value, color, type) {
        let danmaku = document.createElement('div')
        danmaku.classList.add('danmaku')
        danmaku.innerText = value;
        danmaku.style.color = color;
        danmaku.style.top = Danmaku.initTop()
        danmaku.style.left = Danmaku.initLeft()
        danmaku_area.appendChild(danmaku)
        danmaku.onanimationend = () => {
            danmaku_area.removeChild(danmaku)
        }
    }
    static initTop(){
        return Math.floor(Math.random() * 18) * 28 + "px"
    }
    static initLeft(){
        return '0'
    }
}

function loadRecommend(data){
    data.forEach(v => {
        let section = document.createElement('section')
        section.innerHTML =
            `<div><img alt="${v.title}" src="${v.cover}"></div>`+
            `<p class="title">${v.title}</p>`+
            `<p class="author">${v.author}</p>`+
            `<p class="data"><span class="play_number">${v.play_number}</span>`+
            `<span class="danmaku_number">${v.danmaku_number}</span></p>`
        recommend.appendChild(section)
    })
}
function switchVideoPlayStatus(){
    if(video.paused){
        video.play()
        video_wrapper.classList.remove('paused')
        controls_play.setAttribute('src', '/statics/images/video/controls-pause.svg')
    } else {
        video.pause()
        video_wrapper.classList.add('paused')
        controls_play.setAttribute('src', '/statics/images/video/controls-play.svg')
    }
}
function updateVideoLocation(ratio = video.currentTime / video.duration){
    controls_process_played.style.width = ratio * 790 + 'px'
    controls_location.children[0].innerText = formatDuration(video.currentTime)
    controls_process_icon.style.left = ratio * 790 - 5 + 'px' // ratio * 790 - (22/2) + 6 + 'px'
}

function rgbToHex(rgb){
    return '#' + rgb.slice(4, -1).split(', ').map(v => {
        let hex = parseInt(v).toString(16)
        return hex.length === 1 ? '0' + hex : hex
    }).join('').toUpperCase()
}
function checkHex(hex){
    if(hex.length !== 7)
        return false
    if (hex[0] !== '#')
        return false
    return hex.substring(1).every(v => '0123456789ABCDEF'.includes(v))
}
function formatDuration(duration){
    duration = Math.floor(duration)
    let sec = duration % 60
    let min = (duration - sec) / 60
    if (sec < 10) {
        sec = '0' + sec
    }
    if (min < 10) {
        min = '0' + min
    }
    return `${min}:${sec}`
}
function randColor(){
    return '#' + Math.floor(Math.random()*256*256*256).toString(16).toUpperCase()
}

function init(){
    video.addEventListener('click', switchVideoPlayStatus)
    controls_play.addEventListener('click', switchVideoPlayStatus)
    danmaku_switch.addEventListener('click', () => {
        if (danmaku_switch.classList.contains('on')) {
            danmaku_switch.classList.remove('on')
            danmaku_switch.classList.add('off')
            danmaku_area.style.visibility = 'hidden'
        } else {
            danmaku_switch.classList.remove('off')
            danmaku_switch.classList.add('on')
            danmaku_area.style.visibility = 'visible'
        }
    })
    danmaku_font_switch.addEventListener('mouseover', () => danmaku_font_settings.style.display = 'block')
    danmaku_font_switch.addEventListener('mouseout', () => danmaku_font_settings.style.display = 'none')
    danmaku_font_settings.addEventListener('mouseover', () => danmaku_font_settings.style.display = 'block')
    danmaku_font_settings.addEventListener('mouseout', () => danmaku_font_settings.style.display = 'none')
    danmaku_send.addEventListener('click', () => {
        new Danmaku(danmaku_input.value, color_input.value, chosen_danmaku_type)
    })
    controls_process.addEventListener('click', e => {
        let player_ratio = e.offsetX / controls_process.offsetWidth
        video.currentTime = player_ratio * video.duration
        updateVideoLocation(player_ratio)
    })
    scroll_type.addEventListener('click', () => {
        scroll_type.classList.add('chosen')
        top_type.classList.remove('chosen')
        bottom_type.classList.remove('chosen')
        chosen_danmaku_type = 'scroll'
    })
    top_type.addEventListener('click', () => {
        scroll_type.classList.remove('chosen')
        top_type.classList.add('chosen')
        bottom_type.classList.remove('chosen')
        chosen_danmaku_type = 'top'
    })
    bottom_type.addEventListener('click', () => {
        scroll_type.classList.remove('chosen')
        top_type.classList.remove('chosen')
        bottom_type.classList.add('chosen')
        chosen_danmaku_type = 'bottom'
    })
    color_input.addEventListener('input', () => {
        color_input.value = color_input.value.toUpperCase()
        if(checkHex(color_input.value))
            color_preview.style.backgroundColor = color_input.value
        else
            color_preview.style.backgroundColor = 'rgba(0,0,0,0)'
    })
    for (let i = 0; i < color_list.length; i++) {
        color_list[i].addEventListener('click', () => {
            color_input.value = rgbToHex(color_list[i].style.backgroundColor)
            color_preview.style.backgroundColor = color_list[i].style.backgroundColor
        })
    }
    setInterval(updateVideoLocation, 500)
    video.addEventListener('canplay', () => {
        controls_location.children[2].innerText = formatDuration(video.duration)
    })
}
function test(){
    let recommend_data = [
        {
            cover: '/statics/test/rcmd7.webp',
            title: '【8848】跌跟头手机',
            author: '古月浪子',
            play_number: '162',
            danmaku_number: '10',
        },{
            cover: '/statics/test/rcmd2.webp',
            title: '[CS:GO]经典差点干掉队友拿五杀',
            author: 'ほしの雲しょう',
            play_number: '6',
            danmaku_number: '0',
        },{
            cover: '/statics/test/rcmd3.webp',
            title: 'CSS进阶',
            author: 'kying-star',
            play_number: '789',
            danmaku_number: '2',
        },{
            cover: '/statics/test/rcmd4.webp',
            title: 'Web后端第四节课-go杂谈&常用包',
            author: 'sarail',
            play_number: '48',
            danmaku_number: '0',
        },{
            cover: '/statics/test/rcmd5.png',
            title: '我是#鹿乃#的NO.008757号真爱粉，靓号在手，走路带风，解锁专属粉丝卡片，使用专属粉丝装扮，你也来生成你的专属秀起来吧！',
            author: '辇道增柒',
            play_number: '40',
            danmaku_number: '0',
        },{
            cover: '/statics/test/rcmd6.webp',
            title: '打爆灯塔！快乐的Sword Art Online: Fatal Bullet',
            author: 'ほしの雲しょう',
            play_number: '74',
            danmaku_number: '0',
        },{
            cover: '/statics/test/rcmd1.webp',
            title: 'Dota2主播日记226期：翔哥NB，zardNB，肚皇NB（都破音）',
            author: '抽卡素材库',
            play_number: '4183',
            danmaku_number: '29',
        },{
            cover: '/statics/test/rcmd8.webp',
            title: '【原神钢琴】晚安，璃月 | Good Night, Liyue',
            author: 'jerritaaa',
            play_number: '33',
            danmaku_number: '0',
        }
    ]

    loadRecommend(recommend_data)
    setInterval(() => {
        let danmaku = new Danmaku("翔哥NB", randColor(),"scroll")
        console.log(danmaku_area.children.length)
    }, 50)
}

init()
test()