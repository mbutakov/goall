let  button = document.querySelectorAll(".Box");


let isOpenWindow = false;
for(let e in button){
    this.addEventListener( "click" , onClickBox);
}

function onClickBox(e) {
    if (e.target.classList.contains("Box")) {
    if(isOpenWindow == true){
        alert("Закройте предедущее окно")
    }else{
        if (e.target) {
            var image = document.createElement("img");
            let rootWindow = document.createElement("div");
            let icoDiv = document.createElement("div");
            let gridInfo = document.createElement("div");
            let gridDescription = document.createElement("div");
            let closeButton = document.createElement("div");
            let gridInfoRow = document.createElement("div");
            let gridInfoRow_Id = document.createElement("div");
            let gridInfoRow_status = document.createElement("div");
            rootWindow.className = "WindowAchivment";
            icoDiv.className = "icoAchivmentInGetInfo";
            gridInfo.className = "gridInfo";
            gridInfoRow.className = "gridInfoRow";
            gridDescription.className = "gridDescription";
            closeButton.className = "closeButton";

            let box = e.target;
            let id = box.dataset.idachivment;
            let xhr = new XMLHttpRequest();
            var params = 'id=' + encodeURIComponent(id);
            xhr.open("GET", "/getAchivment?" + params, true);
            xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");
            xhr.onload = function (e) {
                console.log(e.target.response)
                let ach = JSON.parse(e.target.response);
                gridInfo.textContent = "Информация";
                image.id = "id";
                image.className = "class";
                image.src = "assets/img/" +ach.Achivments.image;            // image.src = "IMAGE URL/PATH"
                image.style.width = "100%"
                image.style.height = "100%"


                gridInfoRow_Id.textContent = "id: " + ach.Achivments.id;
                gridDescription.textContent = ach.Achivments.description;
                gridInfoRow.textContent = "name: "+ach.Achivments.name;
                gridInfoRow_status.textContent = "status: received"
                gridInfoRow.append(gridInfoRow_Id,gridInfoRow_status);
                gridInfo.append(gridDescription);
                icoDiv.append(image);
                rootWindow.append(icoDiv);
                icoDiv.append(gridInfoRow);
                rootWindow.append(gridInfo);
                closeButton.onclick = closeWindow;
                rootWindow.append(closeButton);
                document.body.append(rootWindow);
                isOpenWindow = true;
            }
            xhr.send();
        }
    }
    }



}

function closeWindow(event) {
    // Для закрытия окна получаем родительский элемент текущего
    // элемента (т.е. кнопки закрыть) с классом .window
    let thisWindow = this.closest(".WindowAchivment");
    if (thisWindow) {
        // И если его нашли - удаляем
        isOpenWindow = false;
        thisWindow.remove();
    }

}