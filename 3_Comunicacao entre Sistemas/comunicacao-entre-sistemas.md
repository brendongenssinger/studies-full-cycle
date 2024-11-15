# Comunicação Entre Sistemas

## Módulos do Curso

1. [REST](#rest)
2. [GraphQL](#graphql)
3. [Content Negotiation na Pratica](#Content-Negotiation-na-Pratica)
4. [Service Discovery](#service-discovery)
5. [Consul](#consul)

---

## REST

### Introdução
Descreva aqui uma introdução ao REST, seus princípios e como ele é utilizado na comunicação entre sistemas.

### Conceitos Principais
- **Recursos**: Descreva o que são recursos e como eles são representados.
- **Verbos HTTP**: GET, POST, PUT, DELETE, etc.
- **Status Codes**: 200, 404, 500, etc.
- **Headers**: Content-Type, Authorization, etc.

### Exemplos de Uso
```http
GET /api/resource HTTP/1.1
Host: example.com
Authorization: Bearer token
```
- **Nível 3** : No response por exemplo o meu cliente for acessar a conta, eu posso retornar um link, que facilita o desenvolvimento.
```JSON
{
    "account":{
        "account_number" : 123,
        "balance":{
            "currency": "usd",
            "value": 1000.00            
        },
        "links":{
            "deposit":"/accounts/123/deposit",
            "withdraw":"/accounts/123/withdraw",
            "transfer":"/accounts/123/transfer",
            "close":"/accounts/123/close"
        }
    }
}
```
  **REST: HAL, Collection + JSON e SIREN** 
  -  hal + json : sempre vai ter alguns pontos
  ```JSON
  {
  "_links": {
    "self": {
      "href": "https://api.example.com/books"
    },
    "next": {
      "href": "https://api.example.com/books?page=2"
    }
  },
  "books": [
    {
      "title": "Book Title",
      "author": "Author Name",
      "_links": {
        "self": {
          "href": "https://api.example.com/books/123"
        }
      }
    }
  ]
}
```

**REST: HTTP Method Negotiation**
Method OPTIONS : Ele nos permite quais métodos, são permitidos quais recursos.

- ``OPTION /api/product HTTP/1.1``

- ``Host : fullcycle.com.br``
 
**REST: Content Negotiation**

``Accept: application/json``

``Content-type: application/json`` / **Error 415** (**Unsupported Media type**)


---

## GraphQL

### Introdução
Descreva aqui uma introdução ao REST, seus princípios e como ele é utilizado na comunicação entre sistemas.

### Conceitos Principais
- **Recursos**: Descreva o que são recursos e como eles são representados.
- **Verbos HTTP**: GET, POST, PUT, DELETE, etc.
- **Status Codes**: 200, 404, 500, etc.
- **Headers**: Content-Type, Authorization, etc.

### Exemplos de Uso
___
```http
GET /api/resource HTTP/1.1
Host: example.com
Authorization: Bearer token
```

## Histórico para rodar o laminas_API_TOOLS do que fez
```
docker-desktop@Desktop-Brendon:~/api-tools-skeleton$ `chmod 777 data`

docker-desktop@Desktop-Brendon:~/api-tools-skeleton$ `cd data`

docker-desktop@Desktop-Brendon:~/api-tools-skeleton/data$ `chmod 777 cache`

docker-desktop@Desktop-Brendon:~/api-tools-skeleton$ `code Dockerfile`
```
> `RUN apt-get update && apt-get install -y sqlite3`
 // Adicionei esse comando no dockerFiles
 ____
> 
 `docker-compose run apigility sh -c "sqlite3 test.sqlite '.databases'"`

## Content-Negotiation-na-Pratica
- Accept whiteList :  Qual é o tipo de cara que você vai trabalhar
  - > application/vnd `(uso do vnd é versions)`
- Content-Type whitelist : Qual é o tipo de conteudo que eu de retorno 
- Nível de Maturidade : 
  ````JSON
  {
    "_links": {
        "self": {
            "href": "http://localhost:8080/users?page=1"
        },
        "first": {
            "href": "http://localhost:8080/users"
        },
        "last": {
            "href": "http://localhost:8080/users?page=1"
        }
    },
    "_embedded": {
        "users": [
            {
                "id": 1,
                "name": "Brendon",
                "email": "brendon.genssinger@gmail.com",
                "_links": {
                    "self": {
                        "href": "http://localhost:8080/users/1"
                    }
                }
            }
        ]
    },
    "page_count": 1,
    "page_size": 25,
    "total_items": 1,
    "page": 1
}
   ```

   