FROM node:16

WORKDIR /usr/src/app

# Adicione opções para apt-get para evitar prompts interativos e limpar o cache
RUN apt-get update && apt-get install -y --no-install-recommends wget \
    && rm -rf /var/lib/apt/lists/*

ENV DOCKERIZE_VERSION v0.8.0

RUN apt-get update \
    && apt-get install -y --no-install-recommends wget \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apt-get autoremove -yqq --purge wget && rm -rf /var/lib/apt/lists/*

COPY package*.json ./

RUN npm install

EXPOSE 3000

# Copia o script de entrada e garante permissões de execução
COPY docker-entrypoint.sh /usr/src/app/
RUN chmod +x /usr/src/app/docker-entrypoint.sh

#ENTRYPOINT ["/usr/src/app/docker-entrypoint.sh"]

#CMD ["node", "index.js"]