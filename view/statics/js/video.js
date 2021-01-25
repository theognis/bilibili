const danmaku_switch = document.querySelector('#danmaku .control .switch')

danmaku_switch.addEventListener('click', () => {
    if (danmaku_switch.classList.contains('on')) {
        danmaku_switch.classList.remove('on')
        danmaku_switch.classList.add('off')
    } else {
        danmaku_switch.classList.remove('off')
        danmaku_switch.classList.add('on')
    }
})