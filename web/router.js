const fs = require('fs');
const http = require('http');
const path = require('path');
// const axios = require('axios');
var qs = require("querystring");

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

//TODO: REWRITE ALL REQUESTS ON AXIOS ???

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
                        ans += `<td>${body[i].id}</td>`;
                        ans += `<td>${body[i].name}</td>`;
                        ans += `<td>${body[i].email}</td>`;
                        ans += `<td id="td${body[i].id}">${body[i].password}<a class="btn-flat" onclick="changePassword(${body[i].id})"><i class="material-icons left">refresh</i></a></td>`;
                        ans += `<td>${body[i].testId}</td>`;
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
    var qs = require("querystring");
    var http = require("http");

    var options = {
        "method": "GET",
        "hostname": "127.0.0.1",
        "port": "8080",
        "path": "/api/scores",
        "headers": {
            "Authorization": parseCookies(request)['token']
        }
    };

    var req = http.request(options, function (res) {
        var chunks = [];

        res.on("data", function (chunk) {
            chunks.push(chunk);
        });

        res.on("end", function () {
            var body = Buffer.concat(chunks);
            // console.log(body.toString());
            let rawString = JSON.parse(body.toString());

            let table = "";


            for (let i = 0; i < rawString.length; i++) {
                table += `<tr>`;
                table += `<td>`;
                table += `${rawString[i].Name}`;
                table += `</td>`;
                table += `<td align="left">`;
                if (rawString[i].Score.SumAmount === 0)
                    table += 'Results not available yet';
                else {
                    table += `<b>Base:</b> ${rawString[i].Score.Base} \\ ${rawString[i].Score.BaseAmount}<br/>`;
                    table += `<b>Reading:</b> ${rawString[i].Score.Reading} \\ ${rawString[i].Score.ReadingAmount}<br/>`;
                    table += `<b>Writing:</b> ${rawString[i].Score.Writing} \\ ${rawString[i].Score.WritingAmount}<br/>`;
                    table += `<b>Speaking:</b> ${rawString[i].Score.Listening} \\ ${rawString[i].Score.ListeningAmount}<br/>`;
                    table += `<b>Summary:</b> ${rawString[i].Score.Sum} \\ ${rawString[i].Score.SumAmount}<br/>`;
                    table += `<b>Grade:</b> ${(rawString[i].Score.Sum / rawString[i].Score.SumAmount * 10).toPrecision(3)}<br/>`;
                    table += `<b>Recommended level:</b> ${rawString[i].Score.RecommendedLevel}`;
                }
                table += `</td>`;
                table += `</tr>`;
            }
            callback(table);
        });
    });

    req.write(qs.stringify({undefined: undefined}));
    req.end();
    // checkToken(request, (valid) => {
    //     if (valid) {
    //         const options = {
    //             hostname: '127.0.0.1',
    //             port: 8080,
    //             path: '/api/students',
    //             method: 'GET',
    //             headers: {
    //                 'Content-Type': 'application/x-www-form-urlencoded',
    //                 'Authorization': parseCookies(request)['token']
    //             }
    //         };
    //
    //         const req = http.request(options, function (res) {
    //             let chunks = [];
    //
    //             res.on("data", function (chunk) {
    //                 chunks.push(chunk);
    //             });
    //
    //             res.on("end", function () {
    //                 let body = Buffer.concat(chunks);
    //                 body = JSON.parse(body.toString());
    //                 let ans = '';
    //                 for (let i = 0; i < body.length; i++) {
    //                     ans += '<tr>';
    //                     ans += '<td>';
    //                     ans += body[i].name;
    //                     ans += '</td>';
    //                     ans += '<td>';
    //                     ans += "10 \\ 10"; //TODO: GET RESULTS FROM SERVER
    //                     ans += '</td>';
    //                     ans += '</tr>';
    //                 }
    //                 callback(ans);
    //             });
    //         });
    //         req.end();
    //     } else
    //         callback('Something is wrong');
    // });
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

function fillTestsView(request, callback) {
    const options = {
        hostname: '127.0.0.1',
        port: 8080,
        path: '/api/tests',
        method: 'GET',
        headers: {
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
            let obj = JSON.parse(body);
            let table = '<div class="row">' + '<div class="col s12">' + '<ul class="tabs">';

            let size;
            switch (obj.length) {
                case 1:
                    size = "col s12";
                    break;
                case 2:
                    size = "col s6";
                    break;
                case 3:
                    size = "col s4";
                    break;
                case 4:
                    size = "col s3";
                    break;
                case 5:
                    size = "col s2";
                    break;
                case 6:
                    size = "col s2";
                    break;
                case 7:
                    size = "col s1";
                    break;
                case 8:
                    size = "col s1";
                    break;
                case 9:
                    size = "col s1";
                    break;
                case 10:
                    size = "col s1";
                    break;
                case 11:
                    size = "col s1";
                    break;
                case 12:
                    size = "col s1";
                    break;
                default:
                    size = "";
                    break;
            }
            for (let i = 0; i < obj.length; i++) {
                if (i === 0)
                    table += `<li class="tab ${size}"><a class="active" href="#test${obj[i].id}">Test ${obj[i].id}</a></li>`;
                else
                    table += `<li class="tab ${size}"><a href="#test${obj[i].id}">Test ${obj[i].id}</a></li>`;
            }

            table += '</ul>' + '</div>';

            for (let i = 0; i < obj.length; i++) {
                table += `<div id="test${obj[i].id}" class="col s12" style="margin-top: 3%">`;
                // Base
                table += `<h3 align="center">Base questions</h3>`;
                table += `<table class="highlight" id="baseQuestionsTable${obj[i].id}"><thead><tr><th style="width: 2%">№</th><th style="width: 32%">Question</th><th style="width: 32%">Options</th><th>Correct answer</th></tr></thead><tbody>`;

                for (let j = 0; j < obj[i].baseQuestions.length; j++) {
                    table += '<tr>';
                    table += `<td>${obj[i].baseQuestions[j].id}</td>`;
                    table += `<td>${obj[i].baseQuestions[j].question}</td>`;
                    table += `<td>A: ${obj[i].baseQuestions[j].optionA}</br>B: ${obj[i].baseQuestions[j].optionB}</br>C: ${obj[i].baseQuestions[j].optionC}</br>D: ${obj[i].baseQuestions[j].optionD}</td>`;
                    table += `<td>${obj[i].answers.base[j].answer}</td>`;
                    table += `</tr>`;
                }

                table += '</tbody>';
                table += '</table>';

                // Reading
                table += `<h3 align="center" style="margin-top: 5%">Reading</h3>`;
                table += `<strong>Reading task: </strong>${obj[i].reading.question}`;
                table += `<table class="highlight" id="readingQuestionsTable${obj[i].id}"><thead><tr><th style="width: 2%">№</th><th style="width: 32%">Question</th><th style="width: 32%">Options</th><th>Correct answer</th></tr></thead><tbody>`;

                for (let j = 0; j < obj[i].reading.questions.length; j++) {
                    table += '<tr>';
                    table += `<td>${obj[i].reading.questions[j].id}</td>`;
                    table += `<td>${obj[i].reading.questions[j].question}</td>`;
                    table += `<td>A: ${obj[i].reading.questions[j].optionA}</br>B: ${obj[i].reading.questions[j].optionB}</br>C: ${obj[i].reading.questions[j].optionC}</br>D: ${obj[i].reading.questions[j].optionD}</td>`;
                    table += `<td>${obj[i].answers.reading[j].answer}</td>`;
                    table += `</tr>`;
                }

                table += '</tbody>';
                table += '</table>';

                // Writing
                table += `<h3 align="center" style="margin-top: 5%">Writing</h3>`;
                table += `<strong>Writing task: </strong>${obj[i].writing}`;

                table += '</div>';
            }

            callback(table);
        });
    });

    req.end();
}

function createStudentsList(request, callback) {
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
            String.prototype.replaceAll = function (search, replacement) {
                var target = this;
                return target.replace(new RegExp(search, 'g'), replacement);
            };
            let body = Buffer.concat(chunks);
            body = JSON.parse(body.toString());
            let ans = '<ul class="collapsible popout" id="studentsList"">';
            for (let i = 0; i < body.length; i++) {
                ans += `<li class="collapsibleStudent">`;
                ans += `<div class="collapsible-header"><i class="material-icons">person</i>${body[i].name}</div>`;
                ans += `<div class="collapsible-body"><b>Writing:</b><br/><p>${(body[i].answers.writing).replaceAll("\n", "<br/>")}</p><br/>`;
                ans += `<a class="waves-effect waves-light btn" id="writing${body[i].id}" onclick="sendWritingGrade(this)" style="margin-right: 2%"><i class="material-icons left">done</i>Set writing grade</a><div class="input-field inline"><input id="writing${body[i].id}Grade"></div><br/>`;
                ans += `<a class="waves-effect waves-light btn" id="speaking${body[i].id}" onclick="sendSpeakingGrade(this)" style="margin-right: 2%"><i class="material-icons left">done</i>Set speaking grade</a><div class="input-field inline"><input id="speaking${body[i].id}Grade"></div>`;
                ans += `</div>`;
                ans += `</li>`;
            }
            ans += '</ul>';
            callback(ans);
        });
    });
    req.end();
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
                        "            <th>ID</th>\n" +
                        "            <th>Name</th>\n" +
                        "            <th>Email</th>\n" +
                        "            <th>Password</th>\n" +
                        "            <th>Test №</th>\n" +
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
                    data = data.replace("{TABLE}", table);
                    getUsername(request, (username) => {
                        data = data.replace("{USERNAME}", username);
                        response.end(data);
                    });
                });
            else if (request.url.indexOf("viewtests.html") > -1)
                fillTestsView(request, (table) => {
                    data = data.replace("{TABLE}", table);
                    getUsername(request, (username) => {
                        data = data.replace("{USERNAME}", username);
                        response.end(data);
                    });
                });
            else if (request.url.indexOf("checkwork.html") > -1)
                createStudentsList(request, (list) => {
                    data = data.replace("{TABLE}", list);
                    getUsername(request, (username) => {
                        data = data.replace("{USERNAME}", username);
                        response.end(data);
                    });
                });
            else {
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
                'Content-Type': 'application/x-www-form-urlencoded',
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

        req.write(body);

        req.on("error", (e) => {
            console.log(e);
        });

        req.end();
    });
}

function sendWritingGradeToDB(request, callback) {
    let auth = parseCookies(request)['token'];
    let body = '';
    request.on('data', chunk => {
        body += chunk.toString();
    });
    request.on('end', (flag) => {
        console.log("GRADE: " + request.headers.grade);
        console.log("STUDENT ID: " + request.headers.studentid);

        if (!auth || auth === '' || auth === 'undefined') {
            callback(false);
            return;
        }

        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/sendWritingGrade',
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                'Authorization': auth,
                'StudentId': request.headers.studentid.toString(),
                'Grade': request.headers.grade
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
    });
}

function sendSpeakingGradeToDB(request, callback) {
    let auth = parseCookies(request)['token'];
    let body = '';
    request.on('data', chunk => {
        body += chunk.toString();
    });
    request.on('end', (flag) => {
        console.log("DATA: " + request.headers.grade);
        console.log("STUDENT ID: " + request.headers.studentid);

        if (!auth || auth === '' || auth === 'undefined') {
            callback(false);
            return;
        }

        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/sendListeningGrade',
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                'Authorization': auth,
                'StudentId': request.headers.studentid.toString(),
                'Grade': request.headers.grade
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
    });
}

function getStudentsFile(request, callback) {
    let options = {
        "method": "GET",
        "hostname": '127.0.0.1',
        "port": "8080",
        "path": "/api/studentsExport",
        "headers": {
            "Authorization": parseCookies(request)['token']
        }
    };

    let req = http.request(options, function (res) {
        let chunks = [];

        res.on("data", function (chunk) {
            chunks.push(chunk);
        });

        res.on("end", function () {
            let body = Buffer.concat(chunks);
            callback(body);
        });
    });

    req.end();
}

function sendStudentsFile(request, callback) {
    let auth = parseCookies(request)['token'];
    let chunks = [];
    let tmpHeader = request.headers;
    request.on('data', chunk => {
        chunks.push(chunk);
    });
    request.on('end', (flag) => {
        // console.log("BODY" + body);
        if (!auth || auth === '' || auth === 'undefined') {
            callback(false);
            return;
        }

        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/student/createWithExcel',
            method: 'POST',
            headers: tmpHeader
            // headers: {
            //     // 'Content-Type': tmpHeader,
            //     'Authorization': auth,
            //     // 'Accept-Encoding': 'gzip, deflate'
            //     // 'Content-Length' : body.length
            // }
        };

        const req = http.request(options, (res) => {
            console.log(`status: ${res.statusCode}`);
            if (res.statusCode === 200)
                callback(true);
            else
                callback(false);
        });

        req.setHeader('Authorization', auth);
        // req.setHeader('Content-Length', "");
        // req.setHeader('cache-control', "no-control");

        req.on("error", (e) => {
            console.log(e);
        });

        req.end(Buffer.concat(chunks));
    });
}

function getTestsFiles(request, callback) {
    let options = {
        "method": "GET",
        "hostname": '127.0.0.1',
        "port": "8080",
        "path": "/api/testsZip",
        "headers": {
            "Authorization": parseCookies(request)['token']
        }
    };

    let req = http.request(options, function (res) {
        let chunks = [];

        res.on("data", function (chunk) {
            chunks.push(chunk);
        });

        res.on("end", function () {
            let body = Buffer.concat(chunks);
            callback(body);
        });
    });

    req.end();
}

function sendTestFile(request, callback) {
    let auth = parseCookies(request)['token'];
    let chunks = [];
    let tmpHeader = request.headers;
    request.on('data', chunk => {
        chunks.push(chunk);
    });
    request.on('end', (flag) => {
        // console.log("BODY" + body);
        if (!auth || auth === '' || auth === 'undefined') {
            callback(false);
            return;
        }

        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/testDoc',
            method: 'POST',
            headers: tmpHeader
            // headers: {
            //     // 'Content-Type': tmpHeader,
            //     'Authorization': auth,
            //     // 'Accept-Encoding': 'gzip, deflate'
            //     // 'Content-Length' : body.length
            // }
        };

        const req = http.request(options, (res) => {
            console.log(`status: ${res.statusCode}`);
            if (res.statusCode === 200)
                callback(true);
            else
                callback(false);
        });

        req.setHeader('Authorization', auth);
        // req.setHeader('Content-Length', "");
        // req.setHeader('cache-control', "no-control");

        req.on("error", (e) => {
            console.log(e);
        });

        req.end(Buffer.concat(chunks));
    });
}

function sendNewPassword(request, callback) {
    let auth = parseCookies(request)['token'];
    let chunks = [];
    let tmpHeader = request.headers;
    request.on('data', chunk => {
        chunks.push(chunk)
    });
    request.on('end', (flag) => {
        if (!auth || auth === '' || auth === 'undefined') {
            callback(false);
            return;
        }

        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/teacher/password',
            method: 'POST',
            headers: tmpHeader
        };

        const req = http.request(options, (res) => {
            console.log(`status: ${res.statusCode}`);
            if (res.statusCode === 200)
                callback(true);
            else
                callback(false);
        });

        req.setHeader('Authorization', auth);

        req.on("error", (e) => {
            console.log(e);
        });

        req.end(Buffer.concat(chunks));
    });
}

function generateNewPassword(request, callback) {
    let auth = parseCookies(request)['token'];
    let chunks = [];
    let tmpHeader = request.headers;
    request.on('data', chunk => {
        chunks.push(chunk)
    });
    request.on('end', (flag) => {
        if (!auth || auth === '' || auth === 'undefined') {
            callback(false);
            return;
        }

        const options = {
            hostname: '127.0.0.1',
            port: 8080,
            path: '/api/student/changePassword',
            method: 'POST',
            headers: tmpHeader
        };

        const req = http.request(options, (res) => {
            console.log(`status: ${res.statusCode}`);
            let body = '';
            res.on('data', chunk => body += chunk.toString());
            res.on('end', (flag) => {
                if (res.statusCode === 200)
                    callback(body);
                else
                    callback(false);
            })
        });

        req.setHeader('Authorization', auth);

        req.on("error", (e) => {
            console.log(e);
        });

        req.end(Buffer.concat(chunks));
    });
}

function getResultsFile(request, callback) {
    let options = {
        "method": "GET",
        "hostname": '127.0.0.1',
        "port": "8080",
        "path": "/api/scoreExcel",
        "headers": {
            "Authorization": parseCookies(request)['token']
        }
    };

    let req = http.request(options, function (res) {
        let chunks = [];

        res.on("data", function (chunk) {
            chunks.push(chunk);
        });

        res.on("end", function () {
            let body = Buffer.concat(chunks);
            callback(body);
        });
    });

    req.end();
}

module.exports = {
    init: function (request, response) {
        console.log("---------------------------------------------------");
        const date = new Date();
        console.log(`Requested: ${request.url} at ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`);
        if (request.url === '/' || request.url === '/index.html') {
            open(path.join(__dirname, '/teacher/index.html'), response);
        } else if (request.url === '/res/logo.png' || request.url === '/sass/materialize.css' || request.url === '/js/bin/materialize.min.js' || request.url === '/favicon.ico') {
            request.url = path.join(__dirname, request.url);
            open(request.url, response);
        } else if (request.url === '/sendWritingGrade') {
            sendWritingGradeToDB(request, callback => {
                if (callback) {
                    response.statusCode = 200;
                    response.end();
                }
            });
        } else if (request.url === '/sendSpeakingGrade') {
            sendSpeakingGradeToDB(request, callback => {
                if (callback) {
                    response.statusCode = 200;
                    response.end();
                }
            });
        } else if (request.url === '/sendTest') {
            sendTest(request, (valid) => {
                if (valid) {
                    request.url = "/teacher/createtest.html";
                    response.writeHead(302, {
                        'Location': "/createtest.html"
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
        } else if (request.url === '/getResultsFile') {
            getResultsFile(request, (valid) => {
                if (valid !== false)
                    response.end(valid);
                response.end()
            })
        } else if (request.url === '/getStudentsFile') {
            getStudentsFile(request, (valid) => {
                if (valid) {
                    response.writeHead(200, {
                        'Content-Type': 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
                        'Content-Disposition': 'attachment; filename=StudentsList.xlsx'
                    });
                    response.end(valid);
                }
            })
        } else if (request.url === '/sendStudentsFile') {
            sendStudentsFile(request, (valid) => {
                response.end()
            })
        } else if (request.url === '/getTestsFiles') {
            getTestsFiles(request, (valid) => {
                if (valid) {
                    response.writeHead(200, {
                        'Content-Type': 'application/zip',
                        'Content-Disposition': 'attachment; filename=Test.zip'
                    });
                    response.end(valid);
                }
            })
        } else if (request.url === '/sendTestFile') {
            sendTestFile(request, (valid) => {
                response.end()
            })
        } else if (request.url === '/sendNewPassword') {
            sendNewPassword(request, (valid) => {
                response.end();
            })
        } else if (request.url === '/generateNewPassword') {
            generateNewPassword(request, (valid) => {
                response.end(valid);
            })
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
                if (pathToGo.endsWith(".html"))
                    switch (pathToGo) {
                        case "/":
                            pathToGo = "/teacher/index.html";
                            break;
                        case "/index.html":
                            pathToGo = "/teacher/index.html";
                            break;
                        // case "/students.html":
                        //     pathToGo = "/teacher/students.html";
                        //     break;
                        // case "/createtest.html":
                        //     pathToGo = "/teacher/createtest.html";
                        //     break;
                        // case "/viewtests.html":
                        //     pathToGo = "/teacher/viewtests.html";
                        //     break;
                        // case "/results.html":
                        //     pathToGo = "/teacher/results.html";
                        //     break;
                        // case "/settings.html":
                        //     pathToGo = "/teacher/settings.html";
                        //     break;
                        default:
                            pathToGo = path.join("teacher", pathToGo);
                            break;
                    }
                else
                    pathToGo = path.join(__dirname, pathToGo);
                console.log("GONNA CHECK THIS REQUEST" + pathToGo);
                if (pathToGo.endsWith(".html") && pathToGo.indexOf("index.html") === -1) {
                    request.url = path.join(__dirname, pathToGo);
                    console.log(request.url);
                    openTemplate(request, response);
                } else open(pathToGo, response);
            });
        }
    }
};