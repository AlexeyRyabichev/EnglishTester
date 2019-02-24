#!/bin/bash
nohup node web/server.js >>web/out.log 2>&1 &
nohup go run server/main.go >>server/out.log 2>&1 &
