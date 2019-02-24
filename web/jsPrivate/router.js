const fs = require('fs');
const http = require('http');
const path = require('path');

function notFound(response) {
    console.log("notFound");
    fs.readFile(path.join("./teacher/notFound.html"), function (error, data) {
        console.log("---------------------------------------------------\n\n");
        response.end(data);
    })
}

function open(path, response) {
    fs.readFile(path, function (error, data) {
        if (error) {
            notFound(response);
        } else {
            console.log("---------------------------------------------------\n\n");
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

function handleAuth(request, callback) {
    let body = '';
    request.on('data', chunk => {
        body += chunk.toString();
    });
    request.on('end', () => {
        // body = body.replace("%40", "@");
        console.log("CREDENTIALS: " + body);
        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/login',
            method: 'POST',
            headers: {
                // 'Content-Type': 'multipart/form-data',
                'Content-Type': 'application/x-www-form-urlencoded',
                // 'Content-Length': Buffer.byteLength(body),
            }
        };

        const req = http.request(options, (res) => {
            console.log(`status: ${res.statusCode}`);
            let tmp = '';
            res.on('data', chunk => {
                tmp += chunk.toString();
            });
            res.on('end', () => {
                console.log("TOKEN: " + res.headers.authorization);
                callback(res.headers.authorization);
            });
        });

        req.on("error", (e) => {
            console.log(e);
        });

        req.write(body);
        // console.log("BODY2: " + body);
        req.end();
    });
}

function getAuthToken(request) {
    // console.log('AUTH: ' + request.cookie.authorization);
    let rc = request.headers.cookie, list = {};
    rc && rc.split(';').forEach(function (cookie) {
        var parts = cookie.split('=');
        list[parts.shift().trim()] = decodeURI(parts.join('='));
    });
    // console.log("COOKIE: " + list);
    // return request.headers.authorization ? request.headers.authorization : "nope";
}

function checkToken(request, callback) {
    let auth = parseCookies(request)['token'];

    if (!auth || auth === '' || auth === 'undefined'){
        callback(false);
        return;
    }

    const options = {
        hostname: '127.0.0.1',
        port: 8080,
        path: '/api/students',
        method: 'GET',
        headers: {
            'Authorization': auth
        }
    };

    const req = http.request(options, (res) => {
        console.log(`status: ${res.statusCode}`);
        if (res.statusCode === 200){
            console.log("I'm HERE");
            callback(true);
        }
        else
            callback(false);
        return;
    });

    req.on("error", (e) => {
        console.log(e);
    });

    req.end();
}

module.exports = {
    init: function (request, response) {
        console.log("---------------------------------------------------");
        console.log(`Requested: ${request.url}`);
        if (request.url === '/') {
            open('./teacher/index.html', response);
        } else if (request.url === '/res/HseLogo.png' || request.url === '/sass/materialize.css' || request.url === '/js/bin/materialize.min.js') {
            request.url = path.join(__dirname, "../", request.url);
            open(request.url, response);
        } else if (request.url === '/auth') {
            handleAuth(request, (token) => {
                let path;
                if (token) {
                    path = "./teacher/students.html";
                    console.log("TOKEN: " + token);
                    response.writeHead(200, {
                        'Set-Cookie': 'token=' + token,
                    });
                    request.headers.authorization = token;
                } else
                    path = "./teacher/index.html";
                console.log("GOTO: " + request.url);
                open(path, response);
            });
        } else {
            checkToken(request, (valid) => {
                let pathtoGo = request.url;
                if (!valid)
                    pathtoGo = "/";
                console.log("VALID: " + valid);
                console.log("VALID FOR REQUEST: " + pathtoGo);
                switch (pathtoGo) {
                    case "/":
                        pathtoGo = "./teacher/index.html";
                        break;
                    case "/index.html":
                        pathtoGo = "./teacher/index.html";
                        break;
                    case "/students.html":
                        pathtoGo = "./teacher/students.html";
                        break;
                    case "/tests.html":
                        pathtoGo = "./teacher/tests.html";
                        break;
                    case "/results.html":
                        pathtoGo = "./teacher/results.html";
                        break;
                    case "/settings.html":
                        pathtoGo = "./teacher/settings.html";
                        break;
                    default:
                        pathtoGo = path.join(__dirname, "../", pathtoGo);
                        console.log("IN DEFAULT WITH URL: " + pathtoGo);
                        break;
                }
                open(pathtoGo, response);
            });
        }
    }
};