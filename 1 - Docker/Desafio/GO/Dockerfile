# CONHECENDO DOCKER 

#docker push brendonmascarenhas/fullcycle:latest

FROM golang:1.20-alpine3.18 AS build

WORKDIR /usr/local/bin/app

RUN go mod init myapp
COPY hello.go ./
#RUN go mod tidy
#RUN go build -v -o /usr/src/app/app
RUN go build -v -o go-docker-cli
#RUN go build -o /usr/local/bin/app

FROM scratch

WORKDIR /usr/src/app
COPY --from=build /usr/local/bin/app /usr/local/bin/app
EXPOSE 8080
CMD [ "/usr/local/bin/app/go-docker-cli" ]