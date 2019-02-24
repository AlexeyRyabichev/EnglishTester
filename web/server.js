const router = require(('./router.js'));
const http = require('http');

const port = 8000;
const server = http.createServer((request, response) => router.init(request, response));

server.listen(port, function () {
    console.log(`server running at port: ${port}`);
});