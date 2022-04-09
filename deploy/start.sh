#!/bin/bash

if [ -z $1 ];then
        docker-compose up -d
	cp -r ./frontend/blog/* ~/blog/nginx/html/
	cp ./frontend/nginx.conf ~/blog/nginx/conf/
elif [ $1 = 'down' ];then 
        docker-compose down
elif [ $1 = 'build' ];then
	docker-compose build	
fi
