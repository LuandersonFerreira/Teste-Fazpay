# Case FazPay

> Olá nesse projeto temos como objetivo demonstrar conhecimentos de desenvolvimento em Golang, 
  o sistema consiste em um CRUD de usuários.

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:

- Utilizamos a versão `<v1.22.0>` do Go.
- Para fazer o setup do banco de dados certifique-se de que você tenha o docker desktop instalado em sua máquina;
- Rode os seguintes comandos:
- `docker pull postgres:14.11-alpine`
- `docker run -d --name fazpay-api -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:14.11-alpine`
- `docker exec -it fazpay psql -U postgres`

## 🚀 Instalando <Projeto>

Para instalar o <sistema> basta ter os requisitos acima e fazer o clone do projeto na sua máquina.


## ☕ Usando <Sistema>

Para usar <Sistema>, siga estas etapas:

dentro da raiz, rode o seguinte comando

```
go run main.go
```
Aqui você já estará com o BackEnd rodando na sua máquina na porta 9000.

## ☕ Documentação do <Sistema>
[Você pode acessar toda a documentação da API através desse link](https://testev3.postman.co/workspace/My-Workspace~c800ff9d-ea5e-47d3-8797-9c92680eea97/documentation/14355244-79440ed2-2693-4e22-8de1-5a587550340d)

## 🚀 decisões do projeto

Utilizei alguns conceitos para tornar o projeto mais escalável: 

- UUID ao invés de Id serial, como identificador primário decidi utilizar uuid baseado no artigo a seguir: https://espigah.medium.com/id-vs-uuid-vs-publickey-5f7e19455b90
  com o id universal e unico(uuid) adicionamos uma camada a mais de proteção para a nossa Api evitando que usuários possam varrer nossa base de dados e pegar alguns dados que podem ser sensíveis, além disso tornamos a validação por chave primária bem mais robusta e sólida;
- Montagem das configs de acordo com um arquivo toml, para processos onde utilizamos nuvem e dockerização da aplicação sempre é importante poder controlar isso e maneira externa ao código, para isso decidi construir as configs do projeto a partir de um arquivo toml;
- Delete logic, para manter um histórico de usuários na plataforma utilizei o delete logic para não descartar dados que possam ser relevantes no futuro, posteriormente seria interessante adicionar o datetime no lugar do bool.
