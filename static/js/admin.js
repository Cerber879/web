//document.addEventListener('DOMContentLoaded', function () {
const form = document.getElementById('form');
form.addEventListener('sumbit', formSend);

async function formSend() {
    let error = formValidate(form);

    if (error === 0) {
        //в сервак
        Preview();
        publish();
        alert('корректный ввод данных');
    } else {
        alert("не все данные верны, перепроверьте пожалуйста")
    }
}

function formValidate(form) {
    let error = 0;
    let formReq = document.querySelectorAll('._req');
    let formReq1 = document.querySelectorAll('._req1');
    for (let index = 0; index < formReq.length; index++) {
        const input = formReq[index];
        formRemoveError(input)
        if (input.value === "") {
            formAddError(input);
            error++;
        }
    };

    for (let index = 0; index < formReq1.length; index++) {
        const input = formReq1[index];
        formRemoveErrorFile(input)
        if (input.value === "") {
            formAddErrorFile(input);
            error++;
        }
    };
    return error;
}

function formAddError(input) {
    input.classList.add('_error');
}

function formRemoveError(input) {
    input.classList.remove('_error');
}

function formAddErrorFile(input) {
    input.parentElement.classList.add('_error');
}

function formRemoveErrorFile(input) {
    input.parentElement.classList.remove('_error');
}

//take image
const author_url = document.getElementById("author-image");
const author_preview = document.getElementById("previewA");
const big_image = document.getElementById("big-image");
const small_image = document.getElementById("small-image");
const parent_big = document.getElementsByClassName("input__big-image");
const first = document.getElementById("preview-post-one");
const second = document.getElementById("preview-post-two");
//take text
const title = document.getElementsByName("titleInput");
const subtitle = document.getElementsByName("subtitleInput");
const date = document.getElementsByName("dateInput");
const author_name = document.getElementsByName("AnameInput");
const text_area_content = document.getElementById("text_area");
let
    preview_title = document.getElementsByName("preview_title"),
    preview_subtitle = document.getElementsByName("preview_subtitle"),
    preview_date = document.getElementsByName("preview_date"),
    preview_author_name = document.getElementsByName("preview_author_name");
//take blank
const block3 = document.getElementById("change_au-image");


let post = {
    "title_g": null,
    "subtitle_g": null,
    "author_name_g": null,
    "date_g": null,
    "author_url_name": null,
    "author_url_name_base64": null,
    "big_image_name": null,
    "big_image_name_base64": null,
    "small_image_name": null,
    "small_image_name_base64": null,
    "text_area_content_g": null,
}


function Preview() {
    preview_title[0].textContent = title[0].value;
    preview_title[1].textContent = title[0].value;
    preview_author_name[0].textContent = author_name[0].value;
    preview_date[0].textContent = date[0].value;
    preview_subtitle[0].textContent = subtitle[0].value;
    preview_subtitle[1].textContent = subtitle[0].value;
};

author_url.addEventListener("change", () => {
    uploadFile(author_url.files[0], author_preview);
});


big_image.addEventListener("change", () => {
    uploadBig(big_image.files[0], first);
    block1.children[0].remove();
});

small_image.addEventListener("change", () => {
    uploadSmall(small_image.files[0], second);
    block2.children[0].remove();
});

function uploadFile(file, filewrite) {
    var reader = new FileReader();
    reader.addEventListener("load", () => {
        filewrite.src = reader.result;
        if (reader.result != "") {
            post.author_url_name_base64 = reader.result;
            block3.src = reader.result;
        };
    },
        false
    );
    post.author_url_name = document.getElementById("author-image").files[0].name;
    reader.readAsDataURL(file);
};


function uploadBig(image) {
    var reader = new FileReader();
    const block1 = document.getElementById("block1");
    reader.addEventListener("load", function (e) {
        block1.insertAdjacentHTML(
            "beforeend",
            `
            <div>
                <img class="big-image__size position__bs-image" src = ` + reader.result + `>
                <div class="position-for-url andpadding">
                    <div class="position_text_camera">
                        <img class="" src="../static/images/camera.png">
                        <span>Upload New </span>
                    </div>
                </div>
            </div>
            `
        );
        if (document.getElementById("big_trash-remove") == null) {
            block1.insertAdjacentHTML(
                "afterend",
                `<div id="big_trash-remove" onclick="mega_remove()">
                <img src="../static/images/trash-2.png">
                <span>Remove </span>
            </div>`
            )
        };
        post.big_image_name_base64 = reader.result;
        first.src = reader.result;
    },
        false
    );
    post.big_image_name = document.getElementById("big-image").files[0].name;
    reader.readAsDataURL(image);
};

function uploadSmall(image) {
    var reader = new FileReader();
    const block2 = document.getElementById("block2");
    reader.addEventListener("load", function (e) {
        block2.insertAdjacentHTML(
            "beforeend",
            `
            <div>
                <img class="small-image__size position__bs-image" src = ` + reader.result + `>
                <div class="position-for-url andpadding">
                    <div>
                        <img class="" src="../static/images/camera.png">
                        <span>Upload New </span>
                    </div>
                </div>
            </div>
            `
        );
        if (document.getElementById("small_trash-remove") == null) {
            block2.insertAdjacentHTML("afterend",
                `<div id="small_trash-remove" onclick="small_remove()">
                    <img src="../static/images/trash-2.png">
                    <span>Remove </span>
                </div>`
            );
        };
        post.small_image_name_base64 = reader.result;
        second.src = reader.result;
    },
        false
    );
    post.small_image_name = document.getElementById("small-image").files[0].name;
    reader.readAsDataURL(image);
}
function mega_remove() {
    let trash_big = document.getElementById("big_trash-remove");
    trash_big.addEventListener("click",
        function (e) {
            trash_big.parentElement.removeChild(block1);
            trash_big.parentElement.insertAdjacentHTML("beforeend",
                `<label id="block1" for="big-image">
                <div class="big-image__size position__bs-image" name="Bblank">
                    <img src="../static/images/camera.png" alt="">
                    <p class="text-upload">Upload</p>
                </div>
            </label>`);
            trash_big.parentElement.removeChild(trash_big);
        },);
    first.src = "../static/images/kek.jpg";
}

function small_remove() {
    let trash_big = document.getElementById("small_trash-remove");
    trash_big.addEventListener("click",
        function (e) {
            trash_big.parentElement.removeChild(block2);
            trash_big.parentElement.insertAdjacentHTML("beforeend", `
            <label id="block2" for="small-image">
                <div class="small-image__size position__bs-image" name="Sblank">
                    <img src="../static/images/camera.png" alt="">
                    <p class="text-upload">Upload</p>
                </div>
            </label>`);
            trash_big.parentElement.removeChild(trash_big);
        },);
    second.src = "../static/images/kek.jpg";
}



function publish() {
    post.title_g = title[0].value;
    post.subtitle_g = subtitle[0].value;
    post.author_name_g = author_name[0].value;
    post.date_g = date[0].value;
    post.text_area_content_g = text_area_content.value;
    let XHR = new XMLHttpRequest();
    XHR.open('POST', '/api/post');
    console.log(JSON.stringify(post));
    XHR.send(JSON.stringify(post));
}