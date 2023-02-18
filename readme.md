![Build Status](https://gitlab.com/pages/plain-html/badges/master/build.svg)

---

API Web

---

Test Golang 

in linux terminal

```
export DBHOST=xxxxxxxx
export DBPORT=xxxx
export DBUSER=xxxxx
export DBPASS=xxxxx
export DBNAME=xxxxx

```

``` go run main.go ```


```
swagger generate server -f api.json -A api 

```

in swagger 
cmd/api-server/main.go

trouver la ligne ou il y a 
```
api := operations.NewAPIAPI(swaggerSpec) 
server := restapi.NewServer(api)

```

et ajouter 
```
server.Port=8080

```

aller sur le lien http://127.0.0.1:8080/docs#/ apres avoir fait la commande ci-dessous

```
go run cmd/api-server/main.go
```

![alt text](https://github.com/Bossuser1/webcustomerapi/blob/swagger/documentation/img/add_image_2.JPG?raw=true)