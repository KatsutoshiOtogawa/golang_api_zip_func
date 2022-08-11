# golang_api_function

```bash
# export FUNCTION_TARGET=HelloWorld
go run cmd/main.go
```

```bash
pack build --builder gcr.io/buildpacks/builder:v1 golang_api_function

docker run -p8080:8080 --env-file=.devcontainer/.env golang_api_function
```

[refer](https://github.com/GoogleCloudPlatform/functions-framework-go)

cloud buildが動かない場合は
リージョンを変えたら動くことがある。

制限されているため。

no concurrent builds quota available to create builds

<!-- # golang_api

```bash
$ go mod tidy -go=1.17

```

htpasswd -bnBC 10 "" password

htpasswd -bnBC 10 "" password | tr -d ':\n'


url=https://localhost/nologin/user/signup 
curl -X POST -H "Content-Type: application/json" -d '' $url

https://localhost/

```py
import requests

url = "https://localhost/nologin/user/signup "
payload = {"username": "myname", "password": "aaaaa", "confirm_password": "aaaaa"}

r = requests.get(url, params=payload)
print(r)

``` -->

golang_api_function
