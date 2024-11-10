# Fundamentos da Arquitetura de Software

## Informações do Curso
- **Nome do Curso:** Fundamentos da Arquitetura de Software
- **Instrutor:** [Wesley]
- **Plataforma:** [Nome da Plataforma]
- **Data de Início:** [04/10/2024]
- **Data de Conclusão:** [Data de Conclusão]

## Módulo 1: [Fundamentos]
### Tópicos Abordados
- #### Arquitetura Tecnológica
  - Especialidade em tecnologias específicas de mercado
  - Diversidade de processionais especialistas 
- #### Arquitetura Corporativa
  - Avaliaçaõ de custos de toda a area de desenvolvimento
  - Padronização de Tecnologias
  - Planejamento de grandes implantações
- #### Arquitetura de Soluções:
  - Transformar requisito de negócio em soluções de software
  - Desenhos arquiteturais da solução de um software, para reproduzir como ele irá funcionar
  - Pode participar do processo comercial de pré-venda e venda
- #### Arquitetura de Software :
  - É uma disciplina da engenharia de software
  - Diretamente ligada ao processo de desenvolvimento de software
  - Afeta diretamente na estrutura organizacional da empresa
### Por que aprender arquitetura de software
  - Poder navegar da visão macro para a visão micro de um ou mais softwares
  - Entender quais são as diversas opções que temos para desenvolver a mesma coisa !
  - Pensar a longo prazo no projeto e sua sustentabilidade
  - Mergulhado em padrões de projeto e de desenvolvimento e de boas práticas !
  - Ter mais clareza do impacto que o software possui na organização como um todo
- #### Arquitetura ou Design : 
  - Atividades relacionadas a arquitetura de software são sempre de designer. Entretando, nem todas atividade de design são sobre arquitetura.
- #### Sustentabilidade
  - Software precisa se pagar ao longo do tempo !
- ### RA's (Requisitos Arquiteturais ):
  - Performance
  - Armazenamento de dados
  - escalabilidade
  - Segurança
  - Legal
  - Audit (Auditoria)
  

## Módulo 2: [Características Arquiteturais]
### Tópicos Abordados:
- ### Operacionais:
  - Nível de disponibilidadconteue
    - Observabilidade, criar um nível de créditos de indisponibilidade
    - Quanto tempo disponivel que quero deixar o sistema 
  - Recuperação de desastres
    - Como vou recuperar quando meu sistema estiver fora do ar.
    - Perfomance, desenvolver software para 50 requisições e 1mil requisições
    - Recuperação (Backup)
    - Confiabilidade e Segurança
    - Robustez
    - Escalabilidade (vertical e horizontal) 
        - Twelve Factors na hora de criar uma aplicação para ajudar com que ela seja escalável
- ### Estruturais:
  - Software precisa ser configurável.
  - Extensibilidade : 
    - Tem que estar pensada, que ela consiga crescer
    - Plugadas nela
  - Fácil Instalação, padronização do ambiente ( container )
  - Internacionalização : Não é recomendado no back-end
  - Internacionalização  : Não é recomendado no back-end
- ### Características cross cutting:
  - Acessibilidade
  - Processo de Retenção e recuperação de dados (quanto tempo os dados serão mantidos)
  - Autenticação e Autorização 
  - Privacidade
  - Segurança
    - Backup em uma outra rede
  - Usabilidade não é só front-end
    - Sua api tem documentação
    - Está utilizando documentação 
    - Tem um readme
## Módulo 3: [Perfomance] 
- ### Performance
  - É o desempenho que um software possui para completar um determinado workload
  - As unidades de medida para avaliarmos a performance de um software são :
    - > Latência ou response time
    - >Throughput
  - Ter um software `performático` é diferente de ter um `software escalável`
- Métricas para medir a performance
  - Diminuindo a latência
    - Normalmente medida em miliseconds
  - 
- ### Escalabilidade
  - ### Escalando Banco de dados:
    - Aumentando recursos computacionais
    - Distribuindo Responsabilidades (`escrita` vs `leitura`)
    - Shards de forma horizontal :
    - Serveless ( Não se preocupa do lado do servidor )
    - Otimização de Queries e índices, tenha um APM
    - Explain nas queries
    - CQRS ( Command Query Responsability Segregation)
  - ### Proxy Reverso
    - Um proxy reverso é um servidor que fica na frente dos servidores web e encaminha as solicitações do cliente ( por exemplo, navegador web) para esses servidores web
    - NGINX
    - HAProxy (HA= HIGH AVAILABILITY)
    - Traefik
  


## Módulo 4: [Resiliência]
### Resiliência
- ### O que é ? :
  - > Resiliência é um conjunto de estratégias adotadas intencionalmente para a adaptação de um sistema quando uma falha ocorre.
  - Ter estratégias de resiliência nos possibilita minimizar os riscos de perda de dados e transações importante para o negócio
- ### Proteger e Ser Protegido : 
  - Em um microserviço que existe 3 serviços, ter um ecosistema que funciona perfeitamente.
  - Um sistema precisa não pode ser `egoísta` ao ponto de realizar mais requisições em um sistema que está falhando.
- ### Health Check
  - Sempre conferir antes a saúde !
- ### Rate limiting
  - Protege o sistema baseado no que ele foi projetado para suportar
  - Você não saber o limite que o sistema não suporta, isso pode da ruim
  - Exemplo, se a aplicação suporta até 100 requisições, se passar disso, ele pode retornar um erro 500
  - Dê preferencias do uso do rate limite
- ### Circuit Break
  - Protege o sistema fazendo com que as requisições feitas para ele, sejam negadas. Ex : `error 500`
  - 
