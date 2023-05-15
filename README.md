# VMY Users Crud
Aplicação de CRUD de usuários, desenvolvido em Golang 1.20 e MongoDB.

### Design Tático (organização de pastas)
Foi utilizado conceitos do Domain Driven Design (DDD) para a organização do projeto.  
Segue abaixo um diagrama que construi para representar o design tático utilizado:
  
![diagrama de fluxo](./tactical-design.png)

## Para Executar
Antes de mais nada, crie um arquivo chamado ```.env``` na raiz do projeto com o seguinte formato:
```
PORT=8888
MONGODB_URI=mongodb://mongodb:27017 #Utilize "mongodb://localhost:27017" se optar por rodar a aplicação fora do docker.
MONGODB_DATABASE_NAME=vmytest

JWT_SALT_KEY=sadfkjaçsdj32
```  

Para executar o projeto é necessário ter o docker e docker compose instalados e executar o seguinte comando:
```
docker compose up --build
```

Ao iniciar o MongoDB pela primeira vez, é criado um usuário inicial para utilizar a aplicação.  
As credenciais são:
```
EMAIL = user@initial.com
PASSWORD = 123456
```

Os testes rodam com o seguinte comando:
```
go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
```
Abra o arquivo gerado ```cover.html``` no navegador para checar a cobertura.

# Endpoints
Arquivo para importar os endpoints no Postman [aqui](postman-collection.json).

### POST ```/users/```
