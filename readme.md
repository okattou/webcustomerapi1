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


commande importante

https://editor.swagger.io/

swagger generate server -f api.json -A api
go mod tidy
cmd/main.go

aller a la ligne 27
	server := restapi.NewServer(api)
	server.Port =8080
server.Port=8080

sudo docker build -tag "web" .


export DBHOST=192.168.1.21
export DBPORT=9006
export DBUSER=postgres
export DBPASS=postgres
export DBNAME=postgres
export HOST=192.168.1.21

go run cmd/api-server/main.go

sudo docker build -t web:v1 .
 sudo docker run -it --rm web:v1 /bin/sh
 sudo docker run -it --rm web:v1
 sudo docker run -it -p 8080:9007 --rm web:v1
 docker run --name mysql-container -e MYSQL ROOT PASSWORD=thisIsAPassword -d mysql:latest
sudo docker run -p 192.168.1.21:9007:8080 -p 192.168.1.21:9008:8081 -it web:v1 /bin/sh

DROP TABLE IF EXISTS connection;
DROP TABLE IF EXISTS connection_customers;
DROP TABLE IF EXISTS contacts;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS produits;

CREATE TABLE IF NOT EXISTS connection (
  connect_id INT NOT NULL,
  ip varchar(250) NOT NULL,
  date_connect timestamp  NOT NULL,
  PRIMARY KEY (connect_id)
);


CREATE TABLE IF NOT EXISTS connection_customer (
  connection_customer_id  INT NOT NULL,
  connect_id INT NOT NULL,
  customer_id INT NOT NULL,
  PRIMARY KEY (connection_customer_id)
);


CREATE TABLE customers(
   customer_id INT GENERATED ALWAYS AS IDENTITY,
   customer_name VARCHAR(255) NOT NULL,
   PRIMARY KEY(customer_id)
);

CREATE TABLE contacts(
   contact_id INT GENERATED ALWAYS AS IDENTITY,
   customer_id INT,
   contact_name VARCHAR(255) NOT NULL,
   phone VARCHAR(15),
   email VARCHAR(100),
   PRIMARY KEY(contact_id),
   CONSTRAINT fk_customer
      FOREIGN KEY(customer_id) 
	  REFERENCES customers(customer_id)
	  ON DELETE SET NULL
);


CREATE TABLE produits(
   produits_id INT GENERATED ALWAYS AS IDENTITY,
   nom VARCHAR(255) NOT NULL,
   prix  INT,
   link_photo_source VARCHAR(255) ,
   link_photo_emplacement VARCHAR(255),
   PRIMARY KEY(produits_id)
);

INSERT INTO connection(connect_ip,ip,date_connect)
VALUES(1,'192.168.1.21',now()),
(2,'192.168.1.21',now());

INSERT INTO customers(customer_id,customer_name)
VALUES(1,'John Doe'),
(2,'Jane Doe');


INSERT INTO contacts(customer_id, contact_name, phone, email)
VALUES(1,'John Doe','(408)-111-1234','john.doe@bluebird.dev'),
      (1,'Jane Doe','(408)-111-1235','jane.doe@bluebird.dev'),
      (2,'David Wright','(408)-222-1234','david.wright@dolphin.dev');
 sudo docker run -p 8080:8080 -p 8081:8081 -it web:v1