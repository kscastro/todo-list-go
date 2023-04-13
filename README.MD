# Todo List in Golang with mongoDB and Redis for cache

### Para executar o projeto

```
docker-compose up
```
Depois

```
make dev-server
```

## Chamadas para os endpoints


Para pegar as tasks
```
curl --request GET \
  --url http://localhost:3000/ \
  --header 'Content-Type: application/json'
 
```
Retorno da chamada
```
[
	{
		"_id": "6437163b55cbac5501929b80",
		"status": false,
		"task": "teste"
	}
]
```


Para Criar as tasks
```
curl --request POST \
  --url http://localhost:3000/ \
  --header 'Content-Type: application/json' \
  --data '{
	"task":"teste",
	"status": false
}'
 
```
Retorno da chamada
```
{
	"_id": "000000000000000000000000",
	"task": "teste2"
}
```

Para Completar a task
```
curl --request PUT \
  --url http://localhost:3000/task/complete/6437163b55cbac5501929b80 \
  --header 'Content-Type: application/json'
 
```
Retorno da chamada
```
"6437163b55cbac5501929b80"
```

Para Desfazer a task
```
curl --request PUT \
  --url http://localhost:3000/task/undo/6437163b55cbac5501929b80 \
  --header 'Content-Type: application/json'
```

Retorno da chamada
```
"6437163b55cbac5501929b80"
```

Para Deletar uma task
```
curl --request DELETE \
  --url http://localhost:3000/task/delete/6437163b55cbac5501929b80 \
  --header 'Content-Type: application/json'
```

Retorno da chamada
```
"6437163b55cbac5501929b80"
```

Para Deletar uma task
```
curl --request DELETE \
  --url http://localhost:3000/task/deleteAll \
  --header 'Content-Type: application/json'
```

Retorno da chamada
```
1
```