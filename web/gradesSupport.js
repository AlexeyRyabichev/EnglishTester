function sendWritingGrade(element) {
    // alert(element.id);
    let id = element.id.replace("writing", "");

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "sendWritingGrade", true);
    xhr.setRequestHeader('grade', document.getElementById(element.id + "Grade").value);
    xhr.setRequestHeader('studentid', id);
    xhr.send();
    M.toast({html: `Writing grade sent!`})
}

function sendListeningGrade(element) {
    let id = element.id.replace("listening", "");

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "sendListeningGrade", true);
    xhr.setRequestHeader('grade', document.getElementById(element.id + "Grade").value);
    xhr.setRequestHeader('studentid', id);
    xhr.send();
    M.toast({html: `Listening grade sent!`})
}