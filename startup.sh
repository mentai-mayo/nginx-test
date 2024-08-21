#!/usr/bin/sh

# exec nginx
/usr/sbin/nginx -c /etc/nginx/nginx.conf

# exec domain socket server
/home/golang/main
