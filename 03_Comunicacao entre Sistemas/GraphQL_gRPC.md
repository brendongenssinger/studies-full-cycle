## GraphQL

[gqlgen](https://gqlgen.com/)


O que é o GraphQL
GraphQL é uma linguagem de consulta para APIs e um runtime para executar essas consultas em seus dados existentes. Ele fornece uma descrição completa e compreensível dos dados em sua API, dá aos clientes o poder de pedir exatamente o que precisam e nada mais, facilita a agregação de dados de múltiplas fontes e usa um sistema de tipos para descrever os dados.

GraphQL foi desenvolvido internamente pelo Facebook em 2012 antes de ser lançado publicamente em 2015. Ele é uma alternativa ao REST e adere a uma abordagem diferente para a comunicação entre cliente e servidor.
___
<br>

# gRPC : [Link do site](https://grpc.io/) | Protocolo Buffers (https://protobuf.dev/)
### O que é :
- gRPC é um framework desenvolvido pela google que tem o objetivo facilitar o processo de comunicação entre sistemas de uma forma extremamente rápida, leve, independente de linguagem ( Google )
- Faz parte da CNCF 
### Em quais casos podemos utilizar :
- Ideal para microsserviços

[link gRPC Dotnet](https://learn.microsoft.com/pt-br/aspnet/core/grpc/basics?view=aspnetcore-9.0)

### RPC - Remote Procedure Call
Client : server.soma(a,b) => Server func soma(int a, int b)

### Protocol Buffers
o gRPC tráfega do HTTP/2, os dados são tráfegados em dados binários. 
Com o protocol buffers, precisa de um contrato (Schema), serialização e deserialização.

**Definir sintax**
>  syntax = "proto3"
sempre colocar o tipo e o posicionamento:
```
    message SearchRequest{ 
        string query = 1
        int32 page_number = 2
        int32 result_per_page = 3
    }
```
---

## HTTP/2

- Nome original criado pela Google era SPDY
- Lançado em 2015
- Dados trafegados são binários e não texto como no HTTP 1.1
- Utiliza a mesma conexão TCP para enviar e receber dados do cliente e do servidor (Multiplex)
- Server PUSH
- Headers são comprmimidos
- Processo é mais veloz

### Formatos de Comunicaçao : 
gRPC - API "unary"
gRPC - API "Server streaming"
gRPC - API "Client streaming"
gRPC - API "Bi directional streaming"

# Diferença entre Rest e gRPC :
No rest o contrato é pré definido, no gRPC não !
Rest tem uma latência maior, devido aguardar o processamento do servidor e a resposta, no gRPC ele é assincrono. 

___

Rodar o comando 