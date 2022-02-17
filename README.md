# go-egsys

# PROCESSO SELETIVO :: Vehicle Tracking System

Nesse Processo de Seleção o candidato desenvolveu uma API básica para o rastreamento de veículos

## REQUISITOS

* Docker Compose

```
https://docs.docker.com/compose/install
```

* Golang (version 1.7)

```
https://go.dev/
```

## INSTALAÇÃO
### Clone o projeto:
```
git clone https://github.com/fabiom2211/go-egsys.git
```

### Banco de dados:
1. Com o docker-compose instalado basta digitar o seguinte comando, na raiz do projeto:
```
docker-compose up -d 
```

### Colocando a aplicação no ar :
1. Com Goland instalado basta digitar o seguinte comando, na raiz do projeto:
```
go run main.go
```

#### Acesso:
* EGSYS: http://localhost:8077/

##### Modelo:
![alt text](https://github.com/fabiorosa/go-egsys/doc/Diagram.png?raw=true)

## Sobre API
1. Listar todas as Frotas cadastradas (GET)
```
http://localhost:8077/fleets
```
2. Cadastrar Frota (POST)
```
http://localhost:8077/fleets
```
