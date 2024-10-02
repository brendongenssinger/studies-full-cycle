FROM node:15-alpine

WORKDIR /usr/src/app
COPY package*.json ./

RUN npm install --only=production

EXPOSE 3000

COPY . .
CMD [ "node","index.js" ]