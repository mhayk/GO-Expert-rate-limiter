<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/6295/6295417.png" width="100" />
</p>
<p align="center">
    <h1 align="center">GO-EXPERT RATE LIMITER</h1>
</p>
<p align="center">
    <em>Desafio da pós em GO Expert</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/last-commit/mhayk/GO-Expert-rate-limiter?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/mhayk/GO-Expert-rate-limiter?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/mhayk/GO-Expert-rate-limiter?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
    <em>Developed with ❤️ by Mhayk Whandson</em>
</p>
<p align="center">
		<em>Developed with the language, software and tools below.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/V8-4B8BF5.svg?style=flat&logo=V8&logoColor=white" alt="V8">
	<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=flat&logo=Docker&logoColor=white" alt="Docker">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
</p>
<hr>


# Desafio GO: Rate Limiter

O Rate Limiter é uma funcionalidade que limita a quantidade de requisições que um usuário pode fazer em um determinado período de tempo. O objetivo é evitar que um único usuário sobrecarregue o sistema com um número muito grande de requisições.

## Funcionamento
O projeto implementa um rate limiter que utiliza o Redis para armazenar e gerenciar as requisições, oferecendo suporte tanto para limitação baseada em endereço IP quanto para tokens de acesso exclusivos. Aqui está como ele funciona:

- **Limitação por IP**: Restringe o número de requisições permitidas por segundo para cada endereço IP.
- **Limitação por Token**: Permite a configuração de limites distintos para diferentes tokens de acesso, cada um com seu próprio tempo de expiração e contagem de requisições.

Este mecanismo garante um controle eficaz sobre o fluxo de tráfego e ajuda a prevenir abusos e sobrecargas no sistema.

## Configuração

### Pré-requisitos

Verifique se o Docker está instalado e configurado no seu ambiente. Você pode baixar e instalar o Docker a partir do [site oficial](https://www.docker.com/get-started).


### Execução

Para iniciar o projeto, siga os passos abaixo:

1. Na raiz do projeto, execute o comando `docker-compose up --build`.
2. A aplicação estará acessível em `http://localhost:8080`.

### Uso

- **API_KEY**: Para autenticar suas requisições, adicione um cabeçalho `API_KEY` com o token apropriado. Por exemplo: `curl -H "API_KEY: token1" http://localhost:8080/`.

Entendido! Aqui está a seção para o seu `README.md` explicando as configurações do arquivo `.env`:

### Configurações de Limite de Taxa

As configurações de limite de taxa são definidas no arquivo `.env`, localizado na pasta raiz do projeto. Estas configurações ajudam a controlar e limitar o número de requisições à API, prevenindo abusos e garantindo a estabilidade do serviço. Abaixo está uma descrição rápida de cada configuração:

- `MAX_REQUESTS_PER_SECOND=100`: Define o limite global de 100 requisições por segundo para a API. Se esse limite for excedido, novas requisições serão bloqueadas.

- `BLOCK_DURATION_SECONDS=60`: Especifica a duração do bloqueio em segundos quando o limite global de requisições por segundo é excedido. Neste caso, o bloqueio dura 60 segundos (1 minuto).

- `IP_MAX_REQUESTS_PER_SECOND=5`: Define um limite de 5 requisições por segundo por endereço IP. Se um IP específico exceder esse limite, novas requisições desse IP serão bloqueadas.

- `IP_BLOCK_DURATION_SECONDS=300`: Especifica a duração do bloqueio em segundos quando o limite de requisições por IP é excedido. Neste caso, o bloqueio dura 300 segundos (5 minutos).

- `TOKEN_MAX_REQUESTS_PER_SECOND=10`: Define um limite de 10 requisições por segundo para cada token de autenticação. Se um token específico exceder esse limite, novas requisições usando esse token serão bloqueadas.

- `TOKEN_BLOCK_DURATION_SECONDS=300`: Especifica a duração do bloqueio em segundos quando o limite de requisições por token é excedido. Neste caso, o bloqueio dura 300 segundos (5 minutos).

Estas configurações garantem que a API possa gerenciar a carga de maneira eficiente, evitando sobrecarga e garantindo que os recursos estejam disponíveis de forma justa para todos os usuários.

### Exemplo de Configuração no Arquivo `.env`

```plaintext
MAX_REQUESTS_PER_SECOND=100         # 100 requests per second
BLOCK_DURATION_SECONDS=60           # 1 minute
IP_MAX_REQUESTS_PER_SECOND=5        # 5 requests per second
IP_BLOCK_DURATION_SECONDS=300       # 5 minutes
TOKEN_MAX_REQUESTS_PER_SECOND=10    # 10 requests per second
TOKEN_BLOCK_DURATION_SECONDS=300    # 5 minutes
```

Adicione essas configurações ao seu arquivo `.env` para proteger a API contra abusos e manter a integridade do serviço. 🤓

## Teste

- Com o servidor em execução, utilize ferramentas como cURL para testar o rate limiter e verificar se as requisições estão sendo limitadas conforme configurado.

### Comando para teste de carga com Apache Bench

```sh
$ ab -n 9 -c 1 -H "API_KEY: token6" http://localhost:8080/
```

### Saída esperada

```plaintext
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
...
Concurrency Level:      1
Time taken for tests:   0.004 seconds
Complete requests:      1
Failed requests:        0
Total transferred:      130 bytes
HTML transferred:       13 bytes
Requests per second:    280.66 [#/sec] (mean)
Time per request:       3.563 [ms] (mean)
Time per request:       3.563 [ms] (mean, across all concurrent requests)
Transfer rate:          35.63 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     3    3   0.0      3       3
Waiting:        3    3   0.0      3       3
Total:          3    3   0.0      3       3
...
```

### Explicação dos parâmetros

- `-n 9`: Total de requisições a serem realizadas.
- `-c 1`: Número de requisições simultâneas.
- `-H "API_KEY: token1"`: Adiciona o cabeçalho de autenticação com o token especificado.

### Detalhamento do relatório

- **Concurrency Level**: Nível de concorrência (número de requisições simultâneas).
- **Time taken for tests**: Tempo total para a execução dos testes.
- **Complete requests**: Total de requisições completadas.
- **Failed requests**: Total de requisições falhadas.
- **Write errors**: Total de erros de escrita.
- **Total transferred**: Total de dados transferidos.
- **HTML transferred**: Total de dados HTML transferidos.
- **Requests per second**: Requisições por segundo (média).
- **Time per request**: Tempo por requisição (média).
- **Time per request**: Tempo por requisição (média, entre todas as requisições simultâneas).
- **Transfer rate**: Taxa de transferência.

Copie e cole esse conteúdo no seu arquivo Markdown para documentar o processo e os resultados do teste de carga na sua API.

### Rodar os testes isolados

```bash
$ make test/config
```

### Dependências de desenvolvimento

- [Mockery](https://vektra.github.io/mockery/latest/)

```bash
$ go install github.com/vektra/mockery/v2@v2.20.0
$ mockery
```