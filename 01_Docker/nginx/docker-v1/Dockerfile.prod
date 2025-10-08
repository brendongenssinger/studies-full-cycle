#UTILIZANDO O COMANDO COPY 

FROM nginx:1.15-alpine

RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d

## - p é para criar os diretorios e subdiretorios
RUN mkdir /var/www/html -p && touch /var/www/html/index.php

