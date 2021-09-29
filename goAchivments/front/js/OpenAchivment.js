let  button = document.querySelectorAll(".Box");

for(let e in button){
    this.addEventListener( "click" , onClickBox);
}

function onClickBox(e) {
    if (e.target.classList.contains("Box")) {
        if (e.target) {
            console.log("asdasd");
            let rootWindow = document.createElement("div");
            let globalWindow = document.createElement("div");
            rootWindow.className = "FrameDescription";






            let box = e.target;
            let id = box.dataset.idachivment;
            let xhr = new XMLHttpRequest();
            var params = 'id=' + encodeURIComponent(id);
            xhr.open("GET", "/getAchivment?" + params, true);
            xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");
            xhr.onload = function (e){
                console.log(e.target.response)
                let ach = JSON.parse(e.target.response);
                rootWindow.textContent = ach.Achivments.name;
                document.body.append(rootWindow);
            }
            xhr.send();
        }

    }

}