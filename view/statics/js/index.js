let list = document.querySelector('section>.list')
let rank = document.querySelector('section>.rank-list')

function loadList(list, data){
    for(let i = 0; i < 8; i++) {
        let article = document.createElement('article')
        article.innerHTML =
            `<a class="cover" href="/video/114514" target="_blank"><img src="${data[i].cover}"></a>` +
            `<a class="title" href="/video/114514" target="_blank">${data[i].title}</a>` +
            `<a class="author" href="/space/114514" target="_blank">${data[i].author}</a>`
        list.appendChild(article)
    }
}
function loadRank(rank, data){
    for(let i = 0; i < 10; i++) {
        let a = document.createElement('a')
        a.setAttribute('href', '/video/114514')
        a.setAttribute('target', '_blank')
        a.innerText = data[i].title
        rank.appendChild(a)
    }
}
function test(){
    let card_data = [
        {
            cover: '/statics/test/rcmd7.webp',
            title: '【8848】跌跟头手机',
            author: '古月浪子'
        },{
            cover: '/statics/test/rcmd2.webp',
            title: '[CS:GO]经典差点干掉队友拿五杀',
            author: 'ほしの雲しょう'
        },{
            cover: '/statics/test/rcmd3.webp',
            title: 'CSS进阶',
            author: 'kying-star'
        },{
            cover: '/statics/test/rcmd4.webp',
            title: 'Web后端第四节课-go杂谈&常用包',
            author: 'sarail'
        },{
            cover: '/statics/test/rcmd5.png',
            title: '我是#鹿乃#的NO.008757号真爱粉，靓号在手，走路带风，解锁专属粉丝卡片，使用专属粉丝装扮，你也来生成你的专属秀起来吧！',
            author: '辇道增柒'
        },{
            cover: '/statics/test/rcmd6.webp',
            title: '打爆灯塔！快乐的Sword Art Online: Fatal Bullet',
            author: 'ほしの雲しょう'
        },{
            cover: '/statics/test/rcmd1.webp',
            title: 'Dota2主播日记226期：翔哥NB，zardNB，肚皇NB（都破音）',
            author: '抽卡素材库'
        },{
            cover: '/statics/test/rcmd8.webp',
            title: '【原神钢琴】晚安，璃月 | Good Night, Liyue',
            author: 'jerritaaa'
        }
    ]
    let rank_data = [
        {
            cover: '/statics/test/rcmd7.webp',
            title: '【8848】跌跟头手机',
            author: '古月浪子',
        },{
            title: '[CS:GO]经典差点干掉队友拿五杀',
            author: 'ほしの雲しょう',
        },{
            title: 'CSS进阶',
            author: 'kying-star',
        },{
            title: 'Web后端第四节课-go杂谈&常用包',
            author: 'sarail',
        },{
            title: '我是#鹿乃#的NO.008757号真爱粉，靓号在手，走路带风，解锁专属粉丝卡片，使用专属粉丝装扮，你也来生成你的专属秀起来吧！',
            author: '辇道增柒',
        },{
            title: '打爆灯塔！快乐的Sword Art Online: Fatal Bullet',
            author: 'ほしの雲しょう',
        },{
            title: 'Dota2主播日记226期：翔哥NB，zardNB，肚皇NB（都破音）',
            author: '抽卡素材库',
        },{
            title: '【原神钢琴】晚安，璃月 | Good Night, Liyue',
            author: 'jerritaaa',
        },{
            title: '【原神钢琴】晚安，璃月 | Good Night, Liyue',
            author: 'jerritaaa',
        },{
            title: '【原神钢琴】晚安，璃月 | Good Night, Liyue',
            author: 'jerritaaa',
        }
    ]

    loadList(list, card_data)
    loadRank(rank, rank_data)
}

test()