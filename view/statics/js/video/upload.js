let channel_data = [
    { 'id': '0101', 'main': '生活', 'sub': '日常', 'description': '一般日常向的生活类视频' },
    { 'id': '0102', 'main': '生活', 'sub': '搞笑', 'description': '搞笑挑战、剪辑、表演、配音以及各类日常沙雕视频' },
    { 'id': '0103', 'main': '生活', 'sub': '运动', 'description': '一般向运动项目以及惊险刺激的户外极限运动' },
    { 'id': '0104', 'main': '生活', 'sub': '汽车', 'description': '汽车、摩托车、公交车以及竞赛用车等轮式交通工具相关' },
    { 'id': '0105', 'main': '生活', 'sub': '绘画', 'description': '绘画爱好者们关于绘画技巧、绘图过程的分享交流场所' },
    { 'id': '0106', 'main': '生活', 'sub': '手工', 'description': '建议手工艺品的diy制作视频，例如：折纸、手账、橡皮章等' },
    { 'id': '0107', 'main': '生活', 'sub': '其他', 'description': '无法判断分区所属的视频' },
    { 'id': '0201', 'main': '游戏', 'sub': '单机游戏', 'description': '以单机或其联机模式为主要内容的相关视频' },
    { 'id': '0202', 'main': '游戏', 'sub': '网络游戏', 'description': '多人在线游戏为主要内容的相关视频' },
    { 'id': '0203', 'main': '游戏', 'sub': '手机游戏', 'description': '手机及平板设备平台上的游戏相关视频' },
    { 'id': '0204', 'main': '游戏', 'sub': '电子竞技', 'description': '电子竞技游戏项目为主要内容的相关视频' },
    { 'id': '0205', 'main': '游戏', 'sub': '桌游棋牌', 'description': '桌游、棋牌、卡牌、聚会游戏等相关视频' },
    { 'id': '0206', 'main': '游戏', 'sub': '音游', 'description': '通过配合音乐与节奏而进行的音乐类游戏视频' },
    { 'id': '0207', 'main': '游戏', 'sub': 'GMV', 'description': '使用游戏内容或CG为素材制作的MV类型的视频' },
    { 'id': '0208', 'main': '游戏', 'sub': 'Mugen', 'description': '使用Mugen引擎制作或与Mugen相关的游戏视频' },
    { 'id': '0301', 'main': '娱乐', 'sub': '综艺', 'description': '国内外电视或网络自制的综艺正片及精彩剪辑视频' },
    { 'id': '0302', 'main': '娱乐', 'sub': '明星', 'description': '三次元明星相关的动态、资讯、粉丝向自制剪辑视频' },
    { 'id': '0401', 'main': '知识', 'sub': '科学科普', 'description': '以自然科学或基于科学思维展开的科普视频' },
    { 'id': '0402', 'main': '知识', 'sub': '社科人文', 'description': '社会/历史/法律/心里/文化/趣闻/艺术/人物/互联网' },
    { 'id': '0403', 'main': '知识', 'sub': '财经', 'description': '财政金融经济相关视频' },
    { 'id': '0404', 'main': '知识', 'sub': '校园学习', 'description': '学校里学习的相关内容，分享学习经验或是课程干货' },
    { 'id': '0405', 'main': '知识', 'sub': '职业职场', 'description': '职场经验、行业知识相关内容' },
    { 'id': '0406', 'main': '知识', 'sub': '野生技术协会', 'description': '技术型技能展示或是技能经验分享视频' },
    { 'id': '0501', 'main': '影视', 'sub': '短片', 'description': '具有一定故事线的剧情短片或者追求印象诗意化的实验短片' },
    { 'id': '0502', 'main': '影视', 'sub': '影视杂谈', 'description': '影视评论、解说、吐槽、科普、配音等' },
    { 'id': '0503', 'main': '影视', 'sub': '影视剪辑', 'description': '对影视素材进行剪辑再创作的视频' },
    { 'id': '0504', 'main': '影视', 'sub': '预告·资讯', 'description': '影视类相关资讯，预告，花絮等视频' },
    { 'id': '0601', 'main': '音乐', 'sub': '音乐综合', 'description': '收录无法定义到其他音乐子分区的音乐视频' },
    { 'id': '0602', 'main': '音乐', 'sub': '音乐现场', 'description': '音乐实况表演视频' },
    { 'id': '0603', 'main': '音乐', 'sub': '演奏', 'description': '使用各种传统及非传统乐器进行演奏创作' },
    { 'id': '0604', 'main': '音乐', 'sub': '翻唱', 'description': '非官方性质的人声歌曲演绎' },
    { 'id': '0605', 'main': '音乐', 'sub': 'MV', 'description': '音乐录像带，官方为搭配音乐而拍摄的短片。' },
    { 'id': '0607', 'main': '音乐', 'sub': 'VOCALOID·UTAU', 'description': '以Vocaloid及UTAU引擎环境为主，运用各种音源进行创作' },
    { 'id': '0608', 'main': '音乐', 'sub': '电音', 'description': '以电子合成器、音乐软体等产生的电子声响制作的音乐' },
    { 'id': '0609', 'main': '音乐', 'sub': '原创音乐', 'description': '有上传者本人参与制作或原作者授权代投的歌曲及纯音乐' },
    { 'id': '0701', 'main': '动画', 'sub': 'MAD·AMV', 'description': '具有一定创作度的动/静画二次创作视频' },
    { 'id': '0702', 'main': '动画', 'sub': 'MMD·3D', 'description': '使用MMD和其他3D软件制作的视频' },
    { 'id': '0703', 'main': '动画', 'sub': '综合', 'description': '以动画及相关内容的素材的创作' },
    { 'id': '0704', 'main': '动画', 'sub': '短片·手书·配音', 'description': '颇具特色的短片/手书及ACG相关配音' },
    { 'id': '0705', 'main': '动画', 'sub': '手办·模玩', 'description': '手办模玩的测评、改造或其他衍生内容' },
    { 'id': '0706', 'main': '动画', 'sub': '特摄', 'description': '特摄相关衍生视频' },
    { 'id': '0801', 'main': '时尚', 'sub': '美妆', 'description': '涵盖妆容、发型、美甲等教程，彩妆、护肤相关产品测评、分享等' },
    { 'id': '0802', 'main': '时尚', 'sub': '服饰', 'description': '服饰风格、穿搭技巧的教程与演示' },
    { 'id': '0803', 'main': '时尚', 'sub': 'T台', 'description': '发布会走秀现场及模特相关时尚片、采访、后台花絮' },
    { 'id': '0804', 'main': '时尚', 'sub': '健身', 'description': '器械、有氧、拉伸运动等，以达到强身健体、减肥瘦身、形体塑造目的' },
    { 'id': '0805', 'main': '时尚', 'sub': '风尚标', 'description': '时尚明星专访、街拍、时尚购物相关知识科普' },
    { 'id': '0901', 'main': '美食', 'sub': '美食制作', 'description': '包括但不限于料理制作教程，各种菜系、甜点、速食、饮料、小吃制作等' },
    { 'id': '0902', 'main': '美食', 'sub': '美食侦探', 'description': '包括但不限于探店、街边美食、饮食文化、发现特色地域美食、路边摊与热门网红食物等' },
    { 'id': '0903', 'main': '美食', 'sub': '美食评测', 'description': '包括但不限于边吃边聊、测评推荐或吐槽各种美食等' },
    { 'id': '0904', 'main': '美食', 'sub': '田园美食', 'description': '包括但不限于乡野美食、三农采摘、钓鱼赶海等' },
    { 'id': '0905', 'main': '美食', 'sub': '美食记录', 'description': '记录一日三餐，美食vlog、料理、便当、饮品合集、美食小剧场等' },
    { 'id': '1001', 'main': '数码', 'sub': '手机平板', 'description': '手机平板、app和产品教程等相关视频' },
    { 'id': '1002', 'main': '数码', 'sub': '电脑装机', 'description': '电脑笔记本、装机配件、外设和软件教程等相关视频' },
    { 'id': '1003', 'main': '数码', 'sub': '摄影摄像', 'description': '摄影摄像器材、拍摄剪辑技巧等相关视频' },
    { 'id': '1004', 'main': '数码', 'sub': '影音智能', 'description': '影音设备、智能硬件、生活家电等相关视频' },
    { 'id': '1101', 'main': '动物圈', 'sub': '喵星人', 'description': '与猫相关的视频，包括但不限于猫咪日常、猫咪喂养、猫咪知识、猫咪剧场、猫咪救助、猫咪娱乐相关的内容' },
    { 'id': '1102', 'main': '动物圈', 'sub': '汪星人', 'description': '与狗相关的视频，包括但不限于狗狗日常、狗狗喂养、狗狗知识、狗狗剧场、狗狗救助、狗狗娱乐相关的内容' },
    { 'id': '1103', 'main': '动物圈', 'sub': '野生动物', 'description': '与野生动物相关的视频，包括但不限于狮子、老虎、狼等动物内容' },
    { 'id': '1104', 'main': '动物圈', 'sub': '爬宠', 'description': '与爬宠相关的视频，包括但不限于龟、蜘蛛、蜥蜴等以宠物方式养育的爬行类动物内容' },
    { 'id': '1105', 'main': '动物圈', 'sub': '大熊猫', 'description': '与大熊猫相关的视频，包括但不限于熊猫日常、熊猫喂养、熊猫知识、熊猫剧场、熊猫救助、熊猫娱乐相关的内容' },
    { 'id': '1106', 'main': '动物圈', 'sub': '动物综合', 'description': '除喵星人、汪星人、大熊猫、爬宠、野生动物之外的其余动物相关视频以及动物相关的延伸内容' },
    { 'id': '1201', 'main': '舞蹈', 'sub': '宅舞', 'description': '与ACG相关的翻跳、原创舞蹈' },
    { 'id': '1202', 'main': '舞蹈', 'sub': '舞蹈综合', 'description': '收录无法定义到其他舞蹈子分区的舞蹈视频' },
    { 'id': '1203', 'main': '舞蹈', 'sub': '舞蹈教程', 'description': '动作分解，基础教程等具有教学意义的舞蹈视频' },
    { 'id': '1204', 'main': '舞蹈', 'sub': '街舞', 'description': '收录街舞相关内容，包括赛事现场，舞室作品、个人翻跳、FREESTYLE等' },
    { 'id': '1205', 'main': '舞蹈', 'sub': '明星舞蹈', 'description': '国内外明星发布的官方舞蹈及翻跳内容' },
    { 'id': '1206', 'main': '舞蹈', 'sub': '中国舞', 'description': '传承中国艺术文化的舞蹈内容，包括古典舞、民族民间舞、汉唐舞、古风舞' },
    { 'id': '1301', 'main': '国创', 'sub': '国产动画', 'description': '国产连载动画，国产完结动画' },
    { 'id': '1302', 'main': '国创', 'sub': '国产原创相关', 'description': '以国产动画、漫画、小说为素材的二次创作' },
    { 'id': '1303', 'main': '国创', 'sub': '布袋戏', 'description': '布袋戏以及相关剪辑节目' },
    { 'id': '1304', 'main': '国创', 'sub': '资讯', 'description': '原创国产动画、漫画的相关资讯、宣传节目等' },
    { 'id': '1305', 'main': '国创', 'sub': '动态漫·广播剧', 'description': '国产动态漫画、有声漫画、广播剧' },
    { 'id': '1401', 'main': '鬼畜', 'sub': '鬼畜调教', 'description': '使用素材在音频、画面上做一定处理，达到与BGM具有一定同步感的视频' },
    { 'id': '1402', 'main': '鬼畜', 'sub': '音MAD', 'description': '使用素材音频进行一定的二次创作来达到还原原曲的非商业性质稿件' },
    { 'id': '1403', 'main': '鬼畜', 'sub': '人力VOCALOID', 'description': '将人物或者角色的无伴奏素材进行人工调音，使其就像VOCALOID一样歌唱的技术' },
    { 'id': '1404', 'main': '鬼畜', 'sub': '鬼畜剧场', 'description': '使用素材进行人工剪辑编排的有剧情的视频' },
    { 'id': '1405', 'main': '鬼畜', 'sub': '教程演示', 'description': '使用鬼畜相关的科普和教程演示' },
    { 'id': '1501', 'main': '纪录片', 'sub': '人文·历史', 'description': '除宣传片、影视剪辑外的，人文艺术历史纪录剧集或电影、预告、花絮、二创、5分钟以上纪录短片' },
    { 'id': '1502', 'main': '纪录片', 'sub': '科学·探索·自然', 'description': '除演讲、网课、教程外的，科学探索自然纪录剧集或电影、预告、花絮、二创、5分钟以上纪录短片' },
    { 'id': '1503', 'main': '纪录片', 'sub': '军事', 'description': '除时政军事新闻外，军事纪录剧集或电影、预告、花絮、二创、5分钟以上纪录短片' },
    { 'id': '1504', 'main': '纪录片', 'sub': '社会·美食·旅行', 'description': '除VLOG、风光摄影外的，社会美食旅行纪录剧集或电影、预告、花絮、二创、5分钟以上纪录短片' },
    { 'id': '1601', 'main': '番剧', 'sub': '资讯', 'description': '以动画/轻小说/漫画/杂志为主的资讯内容，PV/CM/特报/冒头/映像/预告' },
    { 'id': '1602', 'main': '番剧', 'sub': '官方延伸', 'description': '以动画番剧及声优为主的EVENT/生放送/DRAMA/RADIO/LIVE/特典/冒头等' },
    { 'id': '1701', 'main': '电视剧', 'sub': '国产剧', 'description': '' },
    { 'id': '1702', 'main': '电视剧', 'sub': '海外剧', 'description': '' },
    { 'id': '1801', 'main': '电影', 'sub': '其他国家', 'description': '' },
    { 'id': '1802', 'main': '电影', 'sub': '欧美电影', 'description': '' },
    { 'id': '1803', 'main': '电影', 'sub': '日本电影', 'description': '' },
    { 'id': '1804', 'main': '电影', 'sub': '国产电影', 'description': '' },
]

const video_preview = document.querySelector('#video video')
const cover_preview = document.querySelector('#cover img')
const upload_video = document.querySelector('#upload_video')
const upload_cover = document.querySelector('#upload_cover')
const main_channel = document.querySelector('#main_channel')
const sub_channel = document.querySelector('#sub_channel')
const channel_description = document.querySelector('#channel_description')
const label_list = document.querySelector('#label>div')
const label_input = document.querySelector('#label>input')
const label_button = document.querySelector('#label>button')
const video_reader = new FileReader()
const cover_reader = new FileReader()

function createLabel(){
    const newLabel = document.createElement('span')
    newLabel.textContent = label_input.value
    newLabel.onclick = function () {
        label_list.removeChild(newLabel)
        updateLabelButton()
    }
    return newLabel
}
function updateLabelButton () {
    if (label_list.children.length === 10) {
        label_button.setAttribute('disabled', 'disabled')
        return
    }
    if (label_input.value === '') {
        label_button.setAttribute('disabled', 'disabled')
        return
    }
    if([...label_list.children].map(v => v.textContent).filter(v => v === label_input.value).length) {
        label_button.setAttribute('disabled', 'disabled')
        return
    }
    label_button.removeAttribute('disabled')
}

main_channel.onchange = function () {
    sub_channel.innerHTML =
        channel_data
            .filter(v => v.main === main_channel.options[main_channel.selectedIndex].value)
            .map(v => `<option value="${v.sub}">${v.sub}</option>`)
            .join('')
    channel_description.textContent =
        channel_data
            .filter(v =>
                v.main === main_channel.options[main_channel.selectedIndex].value &&
                v.sub === sub_channel.options[sub_channel.selectedIndex].value
            )[0].description
}
sub_channel.onchange = function () {
    channel_description.textContent =
        channel_data
            .filter(v =>
                v.main === main_channel.options[main_channel.selectedIndex].value &&
                v.sub === sub_channel.options[sub_channel.selectedIndex].value
            )[0].description
}
label_input.oninput = updateLabelButton
label_button.onclick = function () {
    let newLabel = createLabel()
    label_list.appendChild(newLabel)
    updateLabelButton()
}
upload_video.onchange = function () {
    let file = upload_video.files[0]
    video_reader.readAsDataURL(file)
}
upload_cover.onchange = function () {
    let file = upload_cover.files[0]
    cover_reader.readAsDataURL(file)
}
video_reader.onload = function () {
    video_preview.src = video_reader.result
}
cover_reader.onload = function () {
    cover_preview.src = cover_reader.result
}