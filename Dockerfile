FROM nginx
ADD my-nginx.conf /etc/nginx/conf.d/default.conf
RUN mkdir /etc/nginx/html
ADD root/healthcheck /etc/nginx/html/healthcheck
