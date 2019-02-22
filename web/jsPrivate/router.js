const fs = require('fs');
const http = require('http');
const path = require('path');
// const JSON = require('JSON');

// let request, response;

function notFound(response){
    console.log("notFound");
    fs.readFile(path.join("./teacher/notFound.html"), function (error, data) {
        response.end(data);
    })
}

function open(path, response){
    fs.readFile(path, function (error, data) {
        if (error){
            notFound(response);
        } else{
            response.statusCode = 200;
            response.end(data);
        }
    })
}

function parseCookies(request) {
    var list = {}, rc = request.headers.cookie;

    rc && rc.split(';').forEach(function (cookie) {
        var parts = cookie.split('=');
        list[parts.shift().trim()] = decodeURI(parts.join('='));
    });

    return list;
}

function handleAuth() {
    return true;
}

function getSessionID() {

}

module.exports = {
    init: function (request, response) {
        console.log(`Requested: ${request.url}`);
        console.log(parseCookies(request));
        switch (request.url) {
            case "/auth":
                if (handleAuth()){
                    request.url = "./teacher/students.html";
                    response.writeHead(200, {
                        'Set-Cookie': 'id=' + new Date().getMilliseconds()
                    });
                }
                else
                    request.url = "./teacher/index.html";
                break;
            case "/":
                request.url = "./teacher/index.html";
                break;
            case "/index.html":
                request.url = "./teacher/index.html";
                break;
            case "/students.html":
                request.url = "./teacher/students.html";
                break;
            case "/tests.html":
                request.url = "./teacher/tests.html";
                break;
            case "/results.html":
                request.url = "./teacher/results.html";
                break;
            case "/settings.html":
                request.url = "./teacher/settings.html";
                break;
            default:
                request.url = path.join(__dirname, "../", request.url);
                break;
        }
        open(request.url, response);
    }
};