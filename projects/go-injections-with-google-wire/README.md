# Go Injections With Google Wire

## Instalação

[Documentação](https://github.com/google/wire)

```bash
go install github.com/google/wire/cmd/wire@latest
```

## Verificar onde o pacote foi instalado

```bash
go env | grep GOPATH
```

## Pacotes

```bash
# Driver do Mysql
go get -u github.com/go-sql-driver/mysql
```

## Docker compose do Mysql

```yaml
# docker-compose.yaml
version: '3'

services:
    mysql:
        image: mysql:5.7
        container_name: mysql
        restart: always
        environment:
            MYSQL_ROOT_PASSWORD: root
            MYSQL_DATABASE: orders
            MYSQL_PASSWORD: root
        ports:
            - 3306:3306
        volumes:
            - .docker/mysql:/var/lib/mysql
```

<div>
  <img align="left" src="https://imgur.com/k8HFd0F.png" width=35 alt="Profile"/>
  <sub>Made with 💙 by <a href="https://github.com/venzel">Enéas Almeida</a></sub>
</div>
