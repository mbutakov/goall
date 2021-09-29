let  button = document.querySelectorAll(".Box");

    for(let e in button){
        this.addEventListener( "click" , onClickBox);
    }

function onClickBox(e) {
        if(e.target.classList.contains("Box")) {
            if (e.target) {
                let box = e.target;
                let id = box.dataset.idachivment;
                let xhr = new XMLHttpRequest();
                var params = 'id=' + encodeURIComponent(id);
                xhr.open("GET", "/r?" + params, true);
                xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");
                e.target.remove();
                xhr.send();
            }

        }

}
