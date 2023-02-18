#!/bin/sh

case $BINARY_NAME in 
    api) exec /app/api-server ;;

    # defauult error
    *)
       echo "** unrecignzig '$BINARY_NAME' application name ! \n"
       exit 1
       ;;
esac