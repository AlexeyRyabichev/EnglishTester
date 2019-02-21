const http = require('http');
const fs = require('fs');

const port = 8000;
const hostname = '127.0.0.1';
// const hostname = '138.68.78.205';


const server = http.createServer(function (requset, response) {
    console.log(`Requested: ${requset.url}`);
    if (requset.url.startsWith("/teacher/")) {
        // requset.url = "C:/Code/EnglishTester/web" + requset.url;
        requset.url = __dirname + requset.url;

        if (requset.url.includes("main.html")) {
            fs.readFile(requset.url, "utf-8", function (error, data) {
                data = data.replace("{Message}", "Kinda message");
                response.end(data);
            });
        } else {
            fs.readFile(requset.url, "utf-8", function (error, data) {
                if (error) {
                    console.log(error);
                    response.statusCode = 404;
                    response.end("Page not found");
                } else {
                    response.statusCode = 200;
                    response.setHeader("Content-Type", "text/html");
                    response.end(data);
                }
            })
        }
    } else if (requset.url.includes("res")){
        // var GETIMAGE
    }
});

server.listen(port, function () {
    console.log(`server running at port: ${port}`);
});