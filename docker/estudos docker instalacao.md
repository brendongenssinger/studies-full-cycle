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



| COMANDOS   | DESCRIÇÃO |
| ------    | ------ |
| `docker ps` |  Quais são os containers rodando|
| `docker run hello-world` |Rodar uma imagem de um container|
| `docker ps -a` | Mostra todos os containers que estão rodando e os que rodaram|
| `docker run -it ubuntu:latest bash` | Rodar uma imagem, cria o container, o comando -i é para attache com o comando bash |
| `docker start <name-container-id>`   | Rodar um container |

### Docker Network

| COMANDOS          | DESCRIÇÃO                          |
| ----------------- | ---------------------------------- |
| `docker run -p`   | Rodar um container expondo portas `-p` **publish** vai redirecionar a porta mapeada para a porta do container |




 

 

