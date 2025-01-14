#Docker Desktop 
<img src="image-docker.png" width=250 height=200>&nbsp;

- Roda com Hyper-V (Máquina Virtual)
- Precisa de licença PRO
- Exige mais recursos da máquina
- Desempenho bem superior ao Toolbox


#Surge o WSL  - 2016 ( Windows Subsystem for Linux)
Jogar o linux embarcado dentro do linux 
- Acesso a "quase" todos os comandos Linux
- Mas não tem o Kernel completo do Linux
  
 
 #WSL 2 
 - Link : https://docs.docker.com/desktop/wsl/
 - Link : https://github.com/codeedu/wsl2-docker-quickstart

### Fazer backup do wsl2

Copy :
>C:\Users\brend\AppData\Local\Packages\CanonicalGroupLimited.Ubuntu_79rhkp1fndgsc\LocalState\ext4.vhdx
___

### Comandos para configurar o wsl2

Ver a lista de distribuições que tenho instalado do Linux
>wsl --list --verbos


Listar distribuições
>wsl --list --verbose

Definir "Ubuntu-20.04" como padrão
>wsl --set-default Ubuntu-20.04

Verificar se a distribuição padrão foi alterada
>wsl

Entrar no ubuntu
> wsl -d "docker-desktop"


> [!NOTE]
> `docker-desktop` faz o compartilhamento dentro do `Ubuntu` ou outras distribuições **instaladas** na máquina 
>
> ![alt text](image.png)

## Comandos simples para ver os containers, iniciar um novo container

| COMANDOS                            | DESCRIÇÃO                                                                          |
| ----------------------------------- | ---------------------------------------------------------------------------------- |
| `docker ps`                         | Quais são os containers rodando                                                    |
| `docker run hello-world`            | Rodar uma imagem de um container                                                   |
| `docker ps -a`                      | Mostra todos os containers que estão rodando e os que rodaram                      |
| `docker run -it ubuntu:latest bash` | Rodar uma imagem, cria o container, o comando -i é para attache com o comando bash |
| `docker start <name-container-id>`  | Rodar um container                                                                 |

## Comandos para rodar o container, porta e acessar o bash

| COMANDOS                     | DESCRIÇÃO                                                                                                     |
| ---------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `docker run -p`              | Rodar um container expondo portas `-p` **publish** vai redirecionar a porta mapeada para a porta do container |
| `docker rm nginx -f`         | Vou remover um container que está rodando ` -f` é para **forçar**                                             |
| `docker rm nginx -f`         | Vou remover um container que está rodando ` -f` é para **forçar**                                             |
| `docker exec nginx ls`       | Executa um comando ls dentro do container                                                                     |
| `docker exec -it nginx bash` | Modo interativo com o bash no docker                                                                          |
 
## Bind Mounts : Se o container morre, os arquivos estão ainda dentro do seu computador 
| COMANDOS                                                                                                     | DESCRIÇÃO                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `docker run -d --name nginx -p 8080:80 -v ~/Projects/fullcycle2/docker/html/: /usr/share/nginx/html nginx`   | Rodar o comando para criar o volume na raiz do sistema operacional <span style="color:red;">Comando Simples.</span>                                  |
| <span style="color:yellow;font-size:17px">-v</span> é um comando antigo `--mount type=bind,source="$(pwd)"/` | É um comando bem antigo                                                                                                                              |
| `echo $(pwd)` <br>/home/docker-desktop                                                                       | É um comando que eu pegue o diretorio, sem ter a necessidade de ficar digitando o caminho<br> `/home/docker-desktop/Projects/fullcycle2/docker/html` |
| ` docker run -d --name nginx -p 8080:80 --mount type=bind,source="$(pwd)",target=/usr/share/nginx/html`      | comando para usar com o mount                                                                                                                        |
| `docker run -d --name nginx -p 8080:80 --mount type=bind,source="$(pwd)"/`                                   | É um comando para criar o mount                                                                                                                      |

<a style="text-align:center" href="https://github.com/brendongenssinger/studies-full-cycle/blob/master/docker/explicacao-sobre-volumes-docker.md"><h3 style="color:red;text-align:center">Link => Diferença entre VOLUMES e BIND AMOUNT </h3></a>
<br><br>
<h2 style="color:green;text-align:center">Criação de Volumes </h2>


| COMANDOS                                                                             | DESCRIÇÃO                                                        |
| ------------------------------------------------------------------------------------ | ---------------------------------------------------------------- |
| `docker volume create meuvolume`                                                     | É um comando para criar o volume                                 |
| ` docker volume ls`                                                                  | Listo os volumes                                                 |
| `docker volume inspect meuvolume`                                                    | Inspeciono um volume especifico                                  |
| `docker run --name nginx -d --mount type=volume,source=meuvolume,target=/app nginx`  | Rodar um novo container para compartilhar o mesmo volume         |
| `docker run --name nginx2 -d --mount type=volume,source=meuvolume,target=/app nginx` | Rodar um novo container para compartilhar o mesmo volume         |
| `docker exec -it nginx2 bash`                                                        | Entro no bash e tenho a possibilidade de abrir o mesmo diretorio |
| `docker run --name nginx3 -d -v meuvolume:/app nginx`                                | Comando para compartilhar o meuvolume dentro do container        |
| `docker volume prune`                                                                | Vai matar tudo que está dentro do volume                         |

 ## Sobre as imagens
 ### Como criar uma imagem com DockerFile
 É uma receita de bolo !

` FROM: Define a imagem base a ser utilizada.`

`RUN: Executa comandos durante a construção da imagem.`

`COPY: Copia arquivos e diretórios para dentro da imagem.`

`WORKDIR: Define o diretório de trabalho para os comandos seguintes.`

`CMD ou ENTRYPOINT: Define o comando a ser executado quando o contêiner for iniciado.`

`docker ps -a -q ` : Lista apenas os ids dos container
  
`docker rm $(docker ps -a -q) -f` Encerra todos os containers 
## Diferença entre CMD x EntryPoint
<b>O comando e entryPoint segura o comando dentro do docker </b>s

CMD:

Define o comando padrão que será executado no contêiner.
Pode ser substituído por argumentos passados ao docker run.
Exemplo:
Dockerfile
Copy code
> `FROM ubuntu`
> 
> `CMD ["echo", "Hello, World!"]`

###Comandos:
> `docker run myimage`: Executa "echo Hello, World!".
> 
> `docker run myimage Hi there!`: 
Substitui CMD, executa "echo Hi there!".
 
### Publicar a imagem no DOCKER HUB

> `docker push` brendonmascarenhas/nginx-fullcycle

###Tipos de Network DOCKER
Rede internas dentro 

1 - Um container comunicar com outro

###Tipos de network :

|Tipos de redes | Explicação|
|----- |----- |
|----- |----- |
|Bridge (PONTE) | É uma rede privada interna criada pelo Docker no host. Cada contêiner conectado a essa rede e se comunicar com outros contêineres na mesma rede, mas é isolado do tráfego de rede externa por padrão. São isolados dos outros contêineres em outras redes ou no host . É uma boa prática de segurança, evitando interferências entre contêiners. Permite a comunicação entre contêineres usando seus nomes em vez de endereços IP.|
| Host (Network) | Principal motivo de não utilizar, no MACOS este serviço não funciona. O docker foi feito para rodar no linux, juntando a rede do container docker e não com a do MACOS, agora quando está utilizando um Windows, mas com WSL Linux, o cenário muda. ` docker run --rm -d --name nginx-host --network host nginx`
| Do seu container precisa acessar alguma porta ou recurso do docker-host, como que faz ? | |

> docker network create -d bridge network-node
___

### Docker logs 
`docker logs <name-docker>`

`docker logs <name-docker> --follow`

`docker logs <name-docker> --details`

____

### Remove todos os container 
`docker rm $(docker ps -a -q ) -f`
___


### Docker-Compose : 

### Remova containers órfãos `Docker-Compose`
`docker-compose -f docker-compose.node.yaml down --remove-orphans`

### Forçar reconstrução da imagem `Docker-Compose` 
`docker-compose -f docker-compose.node.yaml build --no-cache`

### Subir os container `Docker-Compose`
`docker-compose -f docker-compose.node.yaml up -d`



# Docker Compose `depends_on` Documentation

## Overview
The `depends_on` option in Docker Compose is used to define dependencies between services. It specifies the order in which the services should be started. This is particularly useful when one service relies on another to be fully operational before it can start.

## Usage
In a `docker-compose.yml` file, you can use the `depends_on` option within a service definition to indicate that the service depends on one or more other services. This ensures that Docker Compose starts the dependent services in the specified order.

## Example


### Uso do `DOCKERIZE`

`Dockerize` é uma ferramenta útil para garantir que as dependências de um serviço estejam disponíveis antes de iniciar o serviço em si. Isso é especialmente útil em ambientes de contêineres onde a ordem de inicialização pode ser importante.

#### Exemplo de uso do `Dockerize` no caminho `docker/Desafio/Nginx_NodeJS`

1. **Instalar o Dockerize**:
    Adicione o Dockerize ao seu Dockerfile para garantir que ele esteja disponível no contêiner.

    ```Dockerfile
    FROM nginx:alpine

    # Instalar o Dockerize
    RUN apk add --no-cache wget \
         && wget -qO- https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz | tar -C /usr/local/bin -xz

    # Copiar arquivos de configuração
    COPY . /app

    WORKDIR /app

    CMD ["dockerize", "-wait", "tcp://db:5432", "-timeout", "20s", "nginx", "-g", "daemon off;"]
    ```

2. **Configurar o `docker-compose.yml`**:
    Utilize o `docker-compose` para definir os serviços e suas dependências.

    ```yaml
    version: '3.8'

    services:
      nginx:
         build: .
         ports:
            - "80:80"
         depends_on:
            - app
            - db

      app:
         image: node:14
         working_dir: /app
         volumes:
            - .:/app
         command: ["npm", "start"]
        entrypoint: dockerize -wait tcp://db:3306 -timeout 20s docker-entrypoint.sh

      db:
         image: postgres:13
         environment:
            POSTGRES_USER: user
            POSTGRES_PASSWORD: password
            POSTGRES_DB: mydb
    ```

Neste exemplo, o `Dockerize` é usado para garantir que o serviço `nginx` só inicie após o serviço `db` estar disponível na porta 5432. Isso ajuda a evitar erros de conexão e garante que todas as dependências estejam prontas antes de iniciar o serviço principal.

