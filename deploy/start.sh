#!/bin/bash

printUsage()
{
	echo "Usage: ./start.sh [param]"
	echo "  serve: to start the deployment"
	echo "    second param: server ip"
	echo "    example: ./start.sh serve http://192.168.44.100"
	echo "  down: to stop and remove the deployment"
	echo "  build: just build the project"
}

if [ -z $1 ];then
	printUsage
elif [ $1 = 'serve' ];then
	if [ -z $2 ];then
		exit 1
	fi
	
	sed -i 's!^.*imageBaseUrl.*$!  imageBaseUrl: '$2'!' ./backend/conf/application.yaml
	mkdir -p ~/blog/nginx/conf
	mkdir -p ~/blog/nginx/html
	cp -r ./frontend/blog/* ~/blog/nginx/html/
	gzip ~/blog/nginx/html/js/*
	gzip ~/blog/nginx/html/css/*
	gzip ~/blog/nginx/html/fonts/*
	cp -r ./backend/golang_web/images ~/blog
	cp -r ./frontend/conf/* ~/blog/nginx/conf
	docker-compose up -d
elif [ $1 = 'down' ];then
	docker-compose down
	rm -rf ~/blog
elif [ $1 = 'build' ];then
	sed -i 's!^.*imageBaseUrl.*$!  imageBaseUrl: '$2'!' ./backend/conf/application.yaml
	docker-compose build
fi
