# Go

Repositório destinado ao Go.

## Projetos

-   [Go ViaCEP](./go-viacep/)
-   [Go Clean Architecture Básico](./go-clean-architecture-basic/)
-   [Go Agoritimos](./go-algo/)
-   [Go Encoder](./go-encoder/)

## Download, instalação e configuração

-   **Download oficial:** https://go.dev/dl/

```bash
# Remove o diretorio se existir, e extrai o arquivo diretório /usr/local
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz

# Cria a pasta onde irá ficar os pacotes
mkdir $HOME/go

# Edita o zshrc
nano ~/.zshrc

# Insere
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin

# Atualiza o zshrc
source ~/.zshrc

# Verifica a versão
go version
```

## Comandos

```bash
# Exibe a versão do go
go version

# Exibe as variáveis do go
go env

# Inicializa um módulo
go mod init <nome_do_modulo>

# Baixa as dependências de um projeto
go mod tidy

# Cria o build da aplicação de acordo com a variável GOOS do go env
go build

# Cria um build para windows
GOOS=windows go build
```

## Pacotes

```bash
# UUID
go get github.com/google/uuid

# Crypto
go get golang.org/x/crypto
```

## Ferramentas úteis

-   [Converte JSON to struct](https://mholt.github.io/json-to-go/)

## Apis externas

-   [ViaCEP](https://viacep.com.br/)

## Forks de terceiros

-   [Go concorrência](https://github.com/venzel/concorrencia-go)
-   [60 dias com Go](https://github.com/venzel/60-days-of-go)

<hr>

<div>
  <img align="left" src="https://imgur.com/k8HFd0F.png" width=35 alt="Profile"/>
  <sub>Made with 💙 by <a href="https://github.com/venzel">Enéas Almeida</a></sub>
</div>
