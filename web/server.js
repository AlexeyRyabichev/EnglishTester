const router = require(('./router.js'));

const http = require('http');
// const https = require('https');
const fs = require('fs');
const crypto = require('crypto');

const port = 8000;

// const options = {
//     // key: fs.readFile('server.key'),
//     // cert: fs.readFile('server.crt')
//     key: fs.readFile('key.pem'),
//     cert: fs.readFile('cert.pem')
// };

const server = http.createServer((request, response) => router.init(request, response));

server.listen(port, function () {
    console.log(`server running at port: ${port}`);
});