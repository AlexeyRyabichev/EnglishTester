window.onload=function () {
    const loginButton = document.getElementById("loginButton");
    loginButton.addEventListener("click", loginFunc);
    console.log(window.location.pathname);
};

function loginFunc() {
    var login = document.getElementById("login").value;
    var password = document.getElementById("password").value;
    console.log(login);
    console.log(password);
    if (login === "123" && password === "123")
        window.location.href = "/teacher/students.html";
    else
        alert("Wrong login or\\and password");
}
