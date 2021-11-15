# PoC Go PubSub

Esse projeto consiste em uma prova de conceito e exemplificação da aplicabilidade de PubSub.

## Rodando a PoC

Configurações necessárias para rodar a POC:

- Ter um Redis rodando para a aplicação, recomendo utilizar um container com o docker.

```sh
docker run --name redis -d -p 6379:6379 redis redis-server
```

- Criar um arquivo `.env` e configurar a URL do Redis nele de acordo com o `.env.example`

Após as configurações, basta rodar o `worker`:

```sh
go run worker/worker.go
```

E logo em seguida rodar o `client`:

```sh
go run client/client.go
```

## Testando a PoC

```sh
curl --request POST \
  --url http://localhost:8080/client \
  --header 'Content-Type: application/json' \
  --data '{
	"test": "Any message"
}'
```

Após ser feito uma requisição `POST` para a rota `/client`, será mostrado no console do `worker` toda essa mensagem.

## Referencias:

- [PubSub]("https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern")
- [Redis]("https://redis.io/")
