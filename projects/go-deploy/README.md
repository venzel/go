# Go deploy with k8s

## Makefile

```bash
# Sobe o container
make up

# Mata o container
make down

# Executa o programa
make run

# Entra no container
make exec

# Cria o binário com nome server
make build
```

## .Dockerfile

```bash
# Cria uma imagem
FROM golang:latest

# Aponta o diretório app como raiz da aplicação
WORKDIR /app

# Executa o comando apenas para manter o processo funcionando
# e poder entrar dentro do container
CMD ["tail", "-f", "/dev/null"]
```

## .Dockerfile.prod

```bash
# as build: cria um nome de etapa chamada builder, poderia ser outro nome
FROM golang:latest as builder
WORKDIR /app
# Copia todos os arquivos para o workdir
COPY . .
# GOOS: seta o sistema operacional linux
# GO_ENABLED: desabilita a utilização de pacotes do C
# -o: nomeia o binário
# -ldflags="-w -s": Remove o DWARF (Opções de debugging e etc.)
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server .
# scratch: cria uma imagem docker sem nada dentro
FROM scratch
# Copia o binário server da etapa builder para workdir
COPY --from=builder /app/server .
# Executa o binário
CMD ["./server"]
```

## Criar a imagem e subir container

```bash
# Cria
docker build -t venzel/godeploy:latest -f Dockerfile.prod .

# Verifica o tamanho da imagem gerada
docker images | grep godeploy

# Roda o container na porta 3000
# --rm: mata o container quando para de rodar
docker run --rm -p 3000:3000 venzel/godeploy
```

## Subir o container no docker hub

```bash
# Vai pedir o user e o password
docker login

docker push venzel/godeploy:latest
```

## K8s

## Download e instalação

[Documentação](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

## Kind

[Documentação](https://kind.sigs.k8s.io/)

### Download e instalação do kind

```bash
# Versão corrente v0.20.0
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64

chmod +x ./kind

sudo mv ./kind /usr/local/bin/kind

kind --version
```

### Criação de um cluster

```bash
kind create cluster --name=godeploy
```

### Comandos

```bash
# Informações do cluster
kubectl cluster-info --context kind-godeploy

# Verificar os nodes
kubectl get nodes

# Aplica
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Verifica os pods
kubectl get pods

# Verifica os services
kubectl get svc

# Reedirecionamento de portas
kubectl port-forward svc/serversvc 3000:3000

# Testa a API
curl localhost:3000
```

<hr />

<div>
  <img align="left" src="https://imgur.com/k8HFd0F.png" width=35 alt="Profile"/>
  <sub>Made with 💙 by <a href="https://github.com/venzel">Enéas Almeida</a></sub>
</div>
```
