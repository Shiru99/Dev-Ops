# Classroom

## Service 2

Classroom Service + postgresql : Get Classes Details - Endpoint


## DataBase (Postgresql):


***Installation***
```
$ sudo apt-get update
 
$ sudo apt install postgresql postgresql-contrib
```

***Use Default User-psql***
```
$ sudo -u postgres psql postgres
```


***Use New User***
```
$ sudo -u postgres psql postgres
postgres=# CREATE ROLE "user" WITH PASSWORD '1234';
postgres=# ALTER ROLE  "user" WITH LOGIN;
```

***Create New Database***
```
postgres=# CREATE DATABASE classroom;  (dbname - all letters small)
$ psql -U user -h localhost -p 5432 -d classroom;
```