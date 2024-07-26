#docker build -t brendonmascarenhas/laravel:latest

#docker run --rm -d --name laravel 
#-p 8001:8001 brendonmascarenhas/laravel --host=0.0.0.0 --port=8001

#Imagem Multi stage
FROM php:7.4-cli as builder

WORKDIR /var/www

#Comando para instalar o docker php
RUN apt-get update && \
    apt-get install libzip-dev -y && \
    docker-php-ext-install zip

#Instalar o composer do php
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');"

#php composer
RUN php composer.phar create-project --prefer-dist laravel/laravel laravel

#Nova imagem
FROM php:7.4-fpm-alpine

WORKDIR /var/www/
# Removendo a pasta html
RUN rm -rf /var/www/html
## Copiando o build da imagem e colocando nesta pasta
COPY --from=builder /var/www/laravel .

#Permissão para o usuário para ser dono desses arquivos
RUN chown -R www-data:www-data /var/www
EXPOSE 9000
CMD [ "php-fpm" ]