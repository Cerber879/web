var title = {
    fieldname: "title",
    value: null
};

var subTitle = {
    fieldname: "subTitle",
    value: null
};

var authorName = {
    fieldname: "authorName",
    value: null
};
var authorPhoto = {
    fieldname: "authorPhoto",
    value: null
};
var bigImage = {
    fieldname: "bigImage",
    value: null
};
var smallImage = {
    fieldname: "miniImage",
    value: null
};

var Data = {
    fieldname: "Data",
    value: null
};

var postContent = {
    fieldname: "postContent",
    value: null
};

function Click() {
    ChangePreview();
}

function ChangePreview() {
    getDataFromForms();
    ChangePostAuthorPhoto();
    ChangePostBigImage();
    ChangePostSmallImage();
    validateForm();
}

function getDataFromForms() {


    let titleArticle = document.getElementById('preview-title-article');
    let titlePost = document.getElementById('preview-title');
    let title = document.getElementById('Title');
    let subTitle = document.getElementById('Subtitle');
    let subTitleArticle = document.getElementById('preview-subtitle-article');
    let subTitlePost = document.getElementById('preview-subtitle');
    let authorName = document.getElementById('Author-name');
    let authorNamePost = document.getElementById('preview-author-name');
    let Data = document.getElementById('data');
    let dataPost = document.getElementById('preview-data');


    titleArticle.innerHTML = title.value;
    titlePost.innerHTML = title.value;
    subTitleArticle.innerHTML = subTitle.value;
    subTitlePost.innerHTML = subTitle.value;
    authorNamePost.innerHTML = authorName.value;
    authorPhoto.src = authorPhoto.value;
    dataPost.innerHTML = Data.value;
}

function ChangeAuthorPhoto() {
    let fileInput = document.getElementById("author-photo");
    let file = fileInput.files[0];

    if (file) {

        const reader = new FileReader();
        reader.addEventListener("load", () => {
            const authorPhoto = document.getElementById("author-photo-image");
            authorPhoto.src = reader.result;
        }, false);

        reader.readAsDataURL(file);
    }
}
function ChangeBigImage() {
    let fileInput = document.getElementById("big-hero-image");
    let file = fileInput.files[0];

    if (file) {

        const reader = new FileReader();
        reader.addEventListener("load", () => {
            const authorPhoto = document.getElementById("big-image");
            authorPhoto.src = reader.result;
        }, false);

        reader.readAsDataURL(file);
    }
}
function ChangeSmallImage() {
    let fileInput = document.getElementById("small-hero-image");
    let file = fileInput.files[0];

    if (file) {

        const reader = new FileReader();
        reader.addEventListener("load", () => {
            const authorPhoto = document.getElementById("small-image");
            authorPhoto.src = reader.result;
        }, false);

        reader.readAsDataURL(file);
    }
}
function ChangePostAuthorPhoto() {
    let fileInput = document.getElementById("author-photo");
    let file = fileInput.files[0];

    if (file) {

        const reader = new FileReader();
        reader.addEventListener("load", () => {
            const authorPhoto = document.getElementById("preview-author-photo");
            authorPhoto.src = reader.result;
        }, false);

        reader.readAsDataURL(file);
    }
}
function ChangePostBigImage() {
    let fileInput = document.getElementById("big-hero-image");
    let file = fileInput.files[0];

    if (file) {

        const reader = new FileReader();
        reader.addEventListener("load", () => {
            const authorPhoto = document.getElementById("preview-big-image");
            authorPhoto.src = reader.result;
        }, false);

        reader.readAsDataURL(file);
    }
}
function ChangePostSmallImage() {
    let fileInput = document.getElementById("small-hero-image");
    let file = fileInput.files[0];

    if (file) {

        const reader = new FileReader();
        reader.addEventListener("load", () => {
            const authorPhoto = document.getElementById("preview-small-image");
            authorPhoto.src = reader.result;
        }, false);

        reader.readAsDataURL(file);
    }
}

function validateForm() {
    var inputs = document.getElementsByTagName("input");
    for (var i = 0; i < inputs.length; i++) {
        if (inputs[i].value == "") {
            var error = document.querySelector(".main-top__message_complited");
            if (!error) {
                error = document.createElement("div");
                error.classList.add("main-top__message_complited");

                var errorIcon = document.createElement("img");
                errorIcon.src = "../static/images/alert_circle.svg";
                error.appendChild(errorIcon);

                var errorMessage = document.createElement("span");
                errorMessage.innerHTML = "Whoops! Some fields need your attention";
                error.insertBefore(errorIcon, error.firstChild);
                error.appendChild(errorMessage);
                document.getElementById("validation-error").appendChild(error);
            }
            break;
        }
    }
    if (!error && inputs.length > 0) {
        error = document.querySelector(".main-top__massage_complited");
        if (error) {
            error.remove();
        }
        
    }
}
