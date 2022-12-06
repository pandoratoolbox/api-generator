# Golang REST API + TypeScript Client Generator
 Generates a REST API from a Postgresql database as Golang source code along with a TypeScript client

 Free to use and open-source, feel free to contribute.
 
 Managed version with a SaaS panel for deployment support, user management, RBAC configuration, module access configuration., etc coming soon.
 
How to use (Requires Go & NPM):

1. Create .env file
e.g:
```
APP_NAME=audienceviral-api
POSTGRES_HOST=xxx.xxx.xxx.xxx
POSTGRES_PORT=5432
POSTGRES_DB=audienceviral
POSTGRES_USER=postgres
POSTGRES_PASSWORD=xxxxxxx
PORT=3333
```

2. Run the code
3. Your new Golang API will be in a folder in the same directory (named from the APP_NAME env variable), the TypeScript client will be in the /client folder

The /example folder contains an example of a generated Golang API + TypeScript client.
