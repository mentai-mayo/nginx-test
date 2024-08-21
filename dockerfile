
# ----- prod -----

FROM nginx:latest

# update package manager
RUN apt update
RUN apt upgrade -y

# copy nginx configs
RUN rm -f /etc/nginx/confi.d/*.conf
COPY nginx/nginx.conf /etc/nginx/nginx.conf

# copy startup script
COPY startup.sh ./
RUN chmod 777 startup.sh

# set modulation authority to /var/cache/nginx/*
# - /var/cache/nginx/*   chown nginx
# - /var/run/nginx.pid/*
RUN chown -R nginx /var/cache/nginx

USER nginx

# make log file
RUN mkdir -p /tmp/nginx
RUN cat > /tmp/nginx/error.log
RUN cat > /tmp/nginx/access.log

# execute go output as user:nginx
CMD [ "./startup.sh" ]
