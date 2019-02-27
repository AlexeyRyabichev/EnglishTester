const fs = require('fs');
const http = require('http');
const path = require('path');

//TODO: IF TABLE EMPTY => NOT FAIL

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

function getStudentsTable(request, callback) {
    checkToken(request, (valid) => {
        if (valid) {
            const options = {
                hostname: '127.0.0.1',
                port: 8080,
                path: '/api/students',
                method: 'GET',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    'Authorization': parseCookies(request)['token']
                }
            };

            const req = http.request(options, function (res) {
                let chunks = [];

                res.on("data", function (chunk) {
                    chunks.push(chunk);
                });

                res.on("end", function () {
                    let body = Buffer.concat(chunks);
                    body = JSON.parse(body.toString());
                    let ans = '';
                    for (let i = 0; i < body.length; i++) {
                        ans += '<tr>';
                        ans += '<td>';
                        ans += body[i].name;
                        ans += '</td>';
                        ans += '<td>';
                        // noinspection JSUnresolvedVariable
                        ans += body[i].lastName;
                        ans += '</td>';
                        ans += '<td>';
                        // noinspection JSUnresolvedVariable
                        ans += body[i].fatherName;
                        ans += '</td>';
                        ans += '<td>';
                        ans += body[i].email;
                        ans += '</td>';
                        ans += '</tr>';
                    }
                    callback(ans);
                });
            });
            req.end();
        } else
            callback('Something is wrong');
    });
}

function getStudentsResults(request, callback) {
    checkToken(request, (valid) => {
        if (valid) {
            const options = {
                hostname: '127.0.0.1',
                port: 8080,
                path: '/api/students',
                method: 'GET',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    'Authorization': parseCookies(request)['token']
                }
            };

            const req = http.request(options, function (res) {
                let chunks = [];

                res.on("data", function (chunk) {
                    chunks.push(chunk);
                });

                res.on("end", function () {
                    let body = Buffer.concat(chunks);
                    body = JSON.parse(body.toString());
                    let ans = '';
                    for (let i = 0; i < body.length; i++) {
                        ans += '<tr>';
                        ans += '<td>';
                        ans += body[i].name;
                        ans += '</td>';
                        ans += '<td>';
                        ans += "10 \\ 10"; //TODO: GET RESULTS FROM SERVER
                        ans += '</td>';
                        ans += '</tr>';
                    }
                    callback(ans);
                });
            });
            req.end();
        } else
            callback('Something is wrong');
    });
}

function getUsername(request, callback) {
    checkToken(request, (valid) => {
        if (valid) {
            const options = {
                hostname: '127.0.0.1',
                port: 8080,
                path: '/api/teachers',
                method: 'GET',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    'Authorization': parseCookies(request)['token']
                }
            };

            const req = http.request(options, function (res) {
                let chunks = [];

                res.on("data", function (chunk) {
                    chunks.push(chunk);
                });

                res.on("end", function () {
                    let body = Buffer.concat(chunks);
                    body = JSON.parse(body.toString());
                    let ans = '';
                    for (let i = 0; i < body.length; i++) {
                        // noinspection JSUnresolvedVariable
                        if (parseCookies(request)['email'] === body[i].login) { // noinspection JSUnresolvedVariable
                            ans = body[i].lastName + " " + body[i].name;
                        }
                    }
                    callback(ans);
                });
            });
            req.end();
        } else
            callback('Something is wrong');
    });
}

function openTemplate(request, response) {
    fs.readFile(request.url, "utf-8", function (error, data) {
        if (error) {
            notFound(response);
        } else {
            response.statusCode = 200;
            if (request.url.indexOf("students.html") > -1)
                getStudentsTable(request, (table) => {
                    table = "<div class=\"container\">\n" +
                        "    <table class=\"highlight centered\">\n" +
                        "        <thead>\n" +
                        "        <tr>\n" +
                        "            <th>Name</th>\n" +
                        "            <th>Last Name</th>\n" +
                        "            <th>Father Name</th>\n" +
                        "            <th>Email</th>\n" +
                        "        </tr>\n" +
                        "        </thead>\n" +
                        "        <tbody>\n" +
                        table +
                        "        </tbody>\n" +
                        "    </table>\n" +
                        "</div>";
                    data = data.replace("{TABLE}", table);
                    getUsername(request, (username) => {
                        data = data.replace("{USERNAME}", username);
                        response.end(data);
                    });
                });
            else if (request.url.indexOf("results.html") > -1)
                getStudentsResults(request, (table) => {
                    table = "<div class=\"container\">\n" +
                        "    <table class=\"striped highlight centered\">\n" +
                        "        <thead>\n" +
                        "        <tr>\n" +
                        "            <th>Name</th>\n" +
                        "            <th>Result</th>\n" +
                        "        </tr>\n" +
                        "        </thead>\n" +
                        "        <tbody>\n" +
                        table +
                        "        </tbody>\n" +
                        "    </table>\n" +
                        "</div>";
                    data = data.replace("{TABLE}", table);
                    getUsername(request, (username) => {
                        data = data.replace("{USERNAME}", username);
                        response.end(data);
                    });
                });
            else if (request.url.indexOf("tests.html") > -1) {
                getUsername(request, (username) => {
                    data = data.replace("{USERNAME}", username);
                    response.end(data);
                });
            } else if (request.url.indexOf("settings.html") > -1) {
                getUsername(request, (username) => {
                    data = data.replace("{USERNAME}", username);
                    response.end(data);
                });
            }
        }
    })
}

function parseCookies(request) {
    const list = {}, rc = request.headers.cookie;

    rc && rc.split(';').forEach(function (cookie) {
        const parts = cookie.split('=');
        list[parts.shift().trim()] = decodeURI(parts.join('='));
    });

    return list;
}

function handleAuth(request, callback) {
    let body = '', email = '';
    request.on('data', chunk => {
        body += chunk.toString();
        const list = {}, rc = body;

        rc && rc.split('&').forEach(function (cookie) {
            const parts = cookie.split('=');
            list[parts.shift().trim()] = decodeURI(parts.join('='));
        });

        email = list['email'];
        email = email.replace("%40", "@");
    });
    request.on('end', () => {
        console.log("CREDENTIALS: " + body);
        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/login',
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
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
                callback(res.headers.authorization, email);
            });
        });

        req.on("error", (e) => {
            console.log(e);
        });

        req.write(body);
        req.end();
    });
}

function checkToken(request, callback) {
    let auth = parseCookies(request)['token'];

    if (!auth || auth === '' || auth === 'undefined') {
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
        if (res.statusCode === 200)
            callback(true);
        else
            callback(false);
    });

    req.on("error", (e) => {
        console.log(e);
    });

    req.end();
}

function sendTest(request, callback) {
    let auth = parseCookies(request)['token'];
    let body = '';
    request.on('data', chunk => {
        body += chunk.toString();
    });
    request.on('end', (flag) => {
        body = body.replace("testText%3A=", "");
        console.log("JSON: " + body);

        if (!auth || auth === '' || auth === 'undefined') {
            callback(false);
            return;
        }

        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/test',
            method: 'POST',
            headers: {
                'Authorization': auth
            },
            body: body
        };

        const req = http.request(options, (res) => {
            console.log(`status: ${res.statusCode}`);
            if (res.statusCode === 200)
                callback(true);
            else
                callback(false);
        });

        req.on("error", (e) => {
            console.log(e);
        });

        req.end();
    });
}

module.exports = {
    init: function (request, response) {
        console.log("---------------------------------------------------");
        const date = new Date();
        console.log(`Requested: ${request.url} at ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`);
        if (request.url === '/' || request.url === '/index.html') {
            open(path.join(__dirname, '/teacher/index.html'), response);
        } else if (request.url === '/res/logo.png' || request.url === '/sass/materialize.css' || request.url === '/js/bin/materialize.min.js') {
            request.url = path.join(__dirname, request.url);
            open(request.url, response);
        } else if (request.url === '/sendTest') {
            sendTest(request, (valid) => {
                if (valid) {
                    request.url = "/teacher/tests.html";
                    response.writeHead(302, {
                        'Location': "/tests.html"
                    });
                    request.url = path.join(__dirname, request.url);
                    openTemplate(request, response);
                } else {
                    request.url = "/teacher/index.html";
                    response.writeHead(302, {
                        'Location': "/index.html"
                    });
                    request.url = path.join(__dirname, request.url);
                    openTemplate(request, response);
                }

            });
        } else if (request.url === '/auth') {
            handleAuth(request, (token, email) => {
                if (token) {
                    request.url = "/teacher/students.html";
                    console.log("TOKEN: " + token);
                    response.writeHead(302, {
                        'Set-Cookie': ["token=" + token, "email=" + email],
                        'Location': "/students.html"
                    });
                    request.headers.authorization = token;
                    request.url = path.join(__dirname, request.url);
                    openTemplate(request, response);
                } else {
                    response.writeHead(302, {
                        'Location': "/index.html"
                    });
                    open(path.join(__dirname, '/teacher/index.html'), response);
                }
            });
        } else {
            checkToken(request, (valid) => {
                let pathToGo = request.url;
                if (!valid) {
                    response.writeHead(302, {
                        'Location': "/index.html"
                    });
                    open(path.join(__dirname, '/teacher/index.html'), response);
                }
                switch (pathToGo) {
                    case "/":
                        pathToGo = "/teacher/index.html";
                        break;
                    case "/index.html":
                        pathToGo = "/teacher/index.html";
                        break;
                    case "/students.html":
                        pathToGo = "/teacher/students.html";
                        break;
                    case "/tests.html":
                        pathToGo = "/teacher/tests.html";
                        break;
                    case "/results.html":
                        pathToGo = "/teacher/results.html";
                        break;
                    case "/settings.html":
                        pathToGo = "/teacher/settings.html";
                        break;
                    default:
                        pathToGo = path.join(__dirname, pathToGo);
                        break;
                }
                console.log("GONNA CHECK GOT THIS REQUEST" + pathToGo);
                if (pathToGo.indexOf('results.html') > -1 || pathToGo.indexOf('settings.html') > -1 || pathToGo.indexOf('students.html') > -1 || pathToGo.indexOf('tests.html') > -1) {
                    request.url = path.join(__dirname, pathToGo);
                    console.log(request.url);
                    openTemplate(request, response);
                } else open(pathToGo, response);
            });
        }
    }
};