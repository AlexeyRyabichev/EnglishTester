window.onload=function () {
    var loginButton = document.getElementById("loginButton");
    loginButton.addEventListener("click", loginFunc);
    console.log(window.location.pathname);
};

function loginFunc() {
    // alert("Hello");
    // document.location.href = "../students.html";
    // window.location.replace("students.html");
    window.location.href = "students.html";
}