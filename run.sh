#!/bin/bash

# Run your process in background
app_name=$1
launchlog=$2
timeout=$3

if [[ "$#" != "3" ]]; then
  echo "Usage ./run.sh APP_NAME OUTPUT_FILENAME TIMEOUT"
  exit 1
fi
cpid=`cat run.pid`;
if [[ "$cpid" > "0" ]]; then
  echo "Stopping previous running app with PID $cpid";
  kill $cpid;
fi
if [ ! -f $launchlog ]; then
    echo "$launchlog not found!, creating"
    touch $launchlog
fi
./$app_name > $launchlog 2>&1 </dev/null & echo $! > run.pid
# Check if the services started successfully
while read line; do
        echo $line
        if [[ "$line" == *"msg=Started"* ]]; then
                pid=`cat run.pid`
                sleep 3
                if [[ `ps -p $pid -o comm=` ==  "$app_name" ]]; then
                  echo "Started! PID=$pid"
                  exit 0
                else
                  echo "PID $pid is not detected. Set PID to 0"
                  echo "0" > run.pid
                  exit 1
                fi
        fi
done < <(timeout $timeout tail -f $launchlog)
echo "Not started after $timeout seconds, Exit"
exit 1