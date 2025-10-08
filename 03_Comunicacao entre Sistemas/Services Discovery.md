# Anotações de Videoaulas

## Services Discovery
- **Data:** 20/01/2024 
- **Duração:** 
- **Instrutor:** 
- **Link:** 

## Resumo
- Cenário comum em aplicações distribuídas
- Utilização do Consul
  - Consul DNS Server 

## Visão Geral do Consul
### Health Check :
  A partir do momento que para de responder, ele avisa que o serviço está fora do ar
Caso esteja fora do ar, tira ele do server Consul

Consul é multi-cloud
Os consul, conseguem se comunicar.


## Agent, Client e Server : 
Agent : Processo que fica sendo executado em todos nós do cluster.

Client : Registra os serviços localmente, Health check.

Server : Mantém o estado do cluster, registra os serviços, Membership (quem é client e quem é server)

### DEV Mode : 
- Nunca utilize em produção
- Teste de features, exemplos
- Roda como servidor
- Registra tudo em memória


> docker-compose up -d
> docker exec -it consult01
> consult agent -dev
> apk -U add bind-tools 
> dig @localhost -p 8600
Ele retorna o ip 
> consul agent -server -bootstrap-expect=3 -node=consulserver01 -bind=172.21.0.4 -data-dir=var/lib/consul -config-dir=etc/consul.d
>
#### COnsul Cliente
> mkdir /libconsul
> consul agent -bind=<ip_server> -data-dir=/var/lib/consul -config=dir=/etc/consul
> 

> consul catalog nodes -service nginx

> ifconfig
# 172.21.0.4

## Utilizar a UI

> consul agent -config-dir=/etc/consul.d
