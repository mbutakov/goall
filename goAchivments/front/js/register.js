let button = document.querySelector(".form > button");
if (button) {
    button.onclick = function (e) {
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/registerPost");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            let text = document.querySelector(".feedback");
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("Вы создались");
                    text.innerHTML = "Успешно";
                    setTimeout(() => {
                        window.location.href = '/login';
                        }, 2000);
                } else {
                    console.log(response.Error);
                    text.innerHTML = "не успешно";
                }
            } else {
                console.log("Некорректные данные");
            }
        };
        xhr.send(JSON.stringify(data));
    }
}