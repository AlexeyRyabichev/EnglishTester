const router = require(('./router.js'));
const http = require('http');
const https = require('https');
const fs = require('fs');

const port = 8000;

const options = {
  cert: fs.readFileSync('/etc/letsencrypt/live/entest.tk-0001/fullchain.pem'),
  key: fs.readFileSync('/etc/letsencrypt/live/entest.tk-0001/privkey.pem')
};

const server = https.createServer(options,(request, response) => router.init(request, response));
const lowServer = http.createServer((request, response) => {
    response.writeHead(302, {
        'Location' : 'https://entest.tk:' + port + request.url
    });
    response.end();
});

server.listen(port, function () {
    console.log(`server running at port: ${port}`);
});

lowServer.listen(80, function () {
    console.log(`Redirecting all requests from http to https`);
});