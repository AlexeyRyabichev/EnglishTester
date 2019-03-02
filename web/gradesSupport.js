function sendWritingGrade(element){
    alert(element.id);
    // let id = element.id.replace("writing", "");

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "sendWritingGrade", true);
    xhr.setRequestHeader('grade', document.getElementById(element.id + "Grade").value);
    // xhr.setRequestHeader('Content-Type', 'application/json');
    // xhr.send(JSON.stringify({
    //     value: "Hello"
    // }));
}
function sendListeningGrade(element){
    alert(element.id);
    let id = element.id.replace("writing", "");

    // var xhr = new XMLHttpRequest();
    // xhr.open("POST", yourUrl, true);
    // xhr.setRequestHeader('Content-Type', 'application/json');
    // xhr.send(JSON.stringify({
    //     value: value
    // }));
}