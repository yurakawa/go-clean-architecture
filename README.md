# Go-CleanArchitecture-Sample

- echo
- gorm

### db

```
docker container run -d --rm --name article \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_USER=user \
    -e MYSQL_PASSWORD=password \
    -e MYSQL_DATABASE=article \
    -p 3306:3306 mysql:5.7
    
```


### request

```
curl -i -X POST -H "Content-Type: application/json" -d '{"name":"yurakawa","email":"example@example.com"}' http://127.0.0.1:8000/users
```
