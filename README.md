# Go-CleanArchitecture-Sample

- echo
- gorm

### db
```
docker container run -d --rm --name some-mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=hoge -p 3306:3306 mysql:5.7
```


### request

```
curl -i -X POST -H "Content-Type: application/json" -d '{"name":"yurakawa","email":"example@example.com"}' http://127.0.0.1:8000/users
```
