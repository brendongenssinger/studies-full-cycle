# Ciclo de um Projeto.

## Pontos Importantes sobre Arquitetura : 
- Crescimento Sustentável
- Software precisa se pagar ao passar do tempo

## Reflexões Importantes
- Visão de futuro
- Limites bem definidos
- Troca e adição de componentes
- Escala
- Otimização frequentes
- Preparação para mudanças bruscas

## Arquitetura vs Design de software

# Arquitetura Hexagonal ou "Ports and Adapters"

Arquitetura Hexagonal está muito mais ligado com decisões de design de software do 
que necessariamente de arquitetura.

- Defini limites e proteção nas regras da aplicação 
- Componentização e desacoplamento
  - Logs
  - Cache
    - Upload storage
  - banco de dados
  - Filas
  - Http / APIs / GraphQL
- Facilidade na quebra para microsserviços

## Por que "Hexagonal"?

O nome "Hexagonal" foi escolhido por Alistair Cockburn por razões principalmente visuais e conceituais:
- O hexágono é apenas simbólico, poderia ter qualquer número de lados
- Permite visualizar "portas" e "adaptadores" em cada lado do hexágono
- Elimina a ideia de camadas (layers) em favor de "portas e adaptadores"
- Facilita a visualização de que a aplicação pode ter múltiplos "lados" de interação
- Enfatiza que não há diferença hierárquica entre os lados do hexágono

## Arquitetura Hexagonal vs DDD (Domain-Driven Design)

### Principais Diferenças:
1. **Escopo**:
   - DDD: Metodologia completa de design focada em modelagem de domínio complexo
   - Arquitetura Hexagonal: Padrão arquitetural focado em organização e isolamento de componentes

2. **Foco Principal**:
   - DDD: Modelagem do domínio e resolução de complexidade de negócio
   - Arquitetura Hexagonal: Isolamento do core da aplicação e gerenciamento de dependências

3. **Conceitos**:
   - DDD: Trabalha com conceitos como Agregados, Entidades, Value Objects, Bounded Contexts
   - Arquitetura Hexagonal: Trabalha com Ports, Adapters e Domain

### Como se Complementam:
- DDD pode ser implementado usando Arquitetura Hexagonal
- Arquitetura Hexagonal fornece uma excelente estrutura para implementar DDD
- Podem ser usados em conjunto para criar sistemas mais robustos
- DDD ajuda a modelar o domínio dentro do hexágono
- Arquitetura Hexagonal ajuda a proteger o domínio modelado com DDD

## Conceitos Fundamentais da Arquitetura Hexagonal

### Domain (Core)
- Contém as regras de negócio da aplicação
- Completamente isolado de detalhes de implementação
- Não depende de frameworks ou bibliotecas externas
- Representa o coração da aplicação

### Ports (Interfaces)
- Definem contratos de como a aplicação se comunica com o mundo exterior
- Portas de Entrada (Input/Driving Ports): Como a aplicação é acessada
- Portas de Saída (Output/Driven Ports): Como a aplicação acessa recursos externos

### Adapters
- Implementam as interfaces definidas pelas Ports
- Primary/Driving Adapters: Consumem a aplicação (REST, CLI, gRPC, etc)
- Secondary/Driven Adapters: Implementam acesso a recursos externos (BD, Cache, etc)

### Benefícios Adicionais
- Testabilidade aprimorada
- Manutenibilidade facilitada
- Evolução independente de componentes
- Proteção do domínio da aplicação
- Flexibilidade para mudanças tecnológicas

## SqlLite 

> touch `db.sqlite3`
> 
> sqlite3 db.sqlite3
> 
> `create table products(id string, name string, price float, status string);
.tables`
> 
> 
>


## Implementation CLI with Cobra
cobra-cli init  
go run main.go  
go mod tidy

## Estrutura de Pastas .NET para uma API Simples

MyApi.Solution/
├── src/
│   ├── MyApi.Domain/                           # Núcleo do domínio (não depende de nada)
│   │   ├── Entities/
│   │   ├── ValueObjects/
│   │   ├── DomainEvents/
│   │   ├── Services/                           # Domain services (puro domínio)
│   │   └── Exceptions/
│   │
│   ├── MyApi.Application/                      # Casos de uso + Ports (interfaces)
│   │   ├── Abstractions/                       # Contratos/Ports (ex.: ICustomerRepository, IEmailSender)
│   │   ├── UseCases/
│   │   │   ├── Customers/
│   │   │   │   ├── CreateCustomer/
│   │   │   │   │   ├── CreateCustomerCommand.cs
│   │   │   │   │   ├── CreateCustomerHandler.cs
│   │   │   │   │   └── CreateCustomerResult.cs
│   │   │   │   └── GetCustomer/...
│   │   ├── DTOs/
│   │   ├── Behaviors/                          # Pipeline behaviors (validação, logging, tx)
│   │   └── Common/                             # Result, errors, mapeadores básicos
│   │
│   ├── MyApi.Adapters.Inbound.WebApi/          # Adapter de entrada (HTTP)
│   │   ├── Controllers/
│   │   ├── Filters/
│   │   ├── Middleware/
│   │   ├── Contracts/                          # Requests/Responses HTTP (não vazar DTO da App)
│   │   ├── Mapping/                            # Mapas WebApi <-> Application DTOs/Commands
│   │   └── Program.cs
│   │
│   ├── MyApi.Adapters.Outbound.Persistence/    # Adapter de saída: persistência (ex.: EF Core)
│   │   ├── Ef/                                 # DbContext, Configs, Migrations
│   │   ├── Repositories/                       # Implementações dos ports da Application
│   │   └── Mappers/                            # Mapeamento Domínio <-> Models persistência
│   │
│   ├── MyApi.Adapters.Outbound.External/       # Outros adapters de saída (e-mail, filas, APIs)
│   │   ├── Email/
│   │   ├── Messaging/                          # SQS/Rabbit/Kafka
│   │   └── HttpClients/
│   │
│   └── MyApi.CompositionRoot/                  # Bootstrap/DI (referencia App + Adapters)
│       └── DependencyInjection.cs
│
└── tests/
├── MyApi.UnitTests/                        # Testes de domínio e application (puros)
├── MyApi.ComponentTests/                   # Testes dos adapters (ex.: repos + db em memória)
└── MyApi.IntegrationTests/                 # Fluxos ponta a ponta via WebApi

