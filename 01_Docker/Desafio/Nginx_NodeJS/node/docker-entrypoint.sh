#!/bin/sh

# Espera pelo banco de dados
dockerize -wait tcp://db:3306 -timeout 20s

# Executa npm install para garantir que todas as dependências estão instaladas
npm install

# Inicia o aplicativo Node.js
exec node index.js