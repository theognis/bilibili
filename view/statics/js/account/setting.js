const avatar_label = document.querySelector('#main .avatar label')
const avatar_input = document.querySelector('#main .avatar input')
const avatar_img = document.querySelector('#main .avatar img')
const avatar_reads= new FileReader();
const username_span = document.querySelector('#main span.username')
const statement_span = document.querySelector('#main span.statement')
const gender_span = document.querySelector('#main span.gender')
const birthday_span = document.querySelector('#main span.birthday')

avatar_input.onchange = () => {
    if (avatar_input.files.length === 0) {
        return
    }
    avatar_label.innerText = '上传中...'
    const file = avatar_input.files[0]
    avatar_reads.readAsDataURL(file);
    const form = new FormData();
    form.append("avatar", file);
    form.append("token", user.token);
    fetch('/api/user/avatar', {
        method: 'PUT',
        body: form
    })
        .then(data => data.json())
        .then(json => {
            if (json.status) {
                alert("上传成功")
            } else {
                alert("上传失败：" + json.data)
            }
            avatar_label.innerText = '上传头像'
        })
}
avatar_reads.onload = () => {
    avatar_img.src = avatar_reads.result;
}

async function init() {
    avatar_img.src = user.data.Avatar
    username_span.innerText = user.data.Username
    statement_span.innerText = user.data.Statement
    switch (user.data.Gender) {
        case 'F': gender_span.innerText = '女'; break;
        case 'M': gender_span.innerText = '男'; break;
        case 'O': gender_span.innerText = '其他'; break;
        case 'N': gender_span.innerText = '保密'; break;
    }
    if (user.data.Birthday === '9999-12-12T00:00:00Z') {
        birthday_span.innerText = '未填写'
    } else {
        birthday_span.innerText = user.data.Birthday.substring(0,10)
    }
}

tokenInit.then(() => init())