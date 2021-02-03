const avatar_label = document.querySelector('#main .avatar label')
const avatar_input = document.querySelector('#main .avatar input')
const avatar_img = document.querySelector('#main .avatar img')
const avatar_reads= new FileReader();

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
};
