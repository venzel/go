# Go database

## Docker compose do Mysql e RabbitMQ

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

    rabbitmq:
        image: rabbitmq:3-management
        container_name: rabbitmq
        restart: always
        ports:
            - 5672:5672
            - 15672:15672
        environment:
            RABBITMQ_DEFAULT_USER: guest
            RABBITMQ_DEFAULT_PASS: guest
```

## Subir o container

```bash
docker-container up -d
```

## Acessar o container

```bash
docker exec -it mysql bash
```

## Acessar o Mysql

```bash
mysql -h localhost -u root -proot
```

## Comandos Mysql

```bash
# Exibir tabelas
show databases;

# Criar database
create database orders;

# Dropar uma datrabase
drop database orders;

# Etrar em uma database
use orders;
```
