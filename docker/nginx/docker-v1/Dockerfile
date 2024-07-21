#UTILIZANDO O COMANDO COPY 

FROM nginx:latest

WORKDIR /app
RUN apt-get update
RUN apt-get install vim -y
#OR RUN apt-get update && \
# apt-get install -y

COPY html/ /usr/share/nginx

