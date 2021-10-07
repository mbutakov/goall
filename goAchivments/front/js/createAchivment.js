let button = document.querySelector(".form > button");
if (button) {
    button.onclick = function (e) {
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/create");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("Достижение добавлено");
                } else {
                    console.log(response.Error);
                }
            } else {
                console.log("Некорректные данные");
            }
        };
        xhr.send(JSON.stringify(data));
    }
}

/**
 * @type {HTMLInputElement}
 */
let inputFile = document.querySelector("input[type=\"file\"]");
if (inputFile) {
    inputFile.onchange = function (e) {
        let data = new FormData();

        data.set("File", inputFile.files[0], inputFile.files[0].name);

        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/upload");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("Достижение добавлено");
                } else {
                    console.log(response.Error);
                }
            } else {
                console.log("Некорректные данные");
            }
        };
        xhr.send(data);
    }
}