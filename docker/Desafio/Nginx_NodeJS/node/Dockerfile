FROM node:15-alpine

WORKDIR /usr/src/app
COPY package*.json ./

RUN npm install --only=production

COPY . .
CMD [ "node","index.js" ]