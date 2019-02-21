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

module.exports = {
    init: function (request, response) {
        console.log(`Requested: ${request.url}`);

        switch (request.url) {
            case "/students.html":
                let data = [];
                request.on('data', chunk => {
                    data.push(chunk);
                });
                request.on('end', () => {
                    console.log(fs.readAsText(data));
                });
                break;
            case "/":
                request.url = "./teacher/index.html";
                break;
            default:
                request.url = "." + request.url;
                // request.url = path.join(__dirname, request.url);
                break;
        }
        open(request.url, response);
    }
};