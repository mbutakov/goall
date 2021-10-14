
let button = document.querySelector(".form > button");
if (button) {
    button.onclick = function (e) {
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/loginPost");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                    setTimeout(() => {
                            window.location.href = '/';
                        },500)
                        console.log("Вы авторизовались");
                } else {
                    console.log(response.Error);
                }
            } else {
                console.log("Ошибка");
            }
        };
        xhr.send(JSON.stringify(data));
    }
}