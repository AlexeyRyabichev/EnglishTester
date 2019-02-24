#/bin/bash  
nohup node web/server.js >>out.log 2>&1 &
nohup go run server/main.go >>out.log 2>&1 &

# TO STOP PROCESSES RUN COMMAND
# kill ${PID}