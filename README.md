
# app Folder
```
Have admin ,middleware ,models  transformers
```
## admin
```
have all modules eash module have files:
controller : have apis 
helpers :have helper functions like filters
model : have all function related to database 
route : have api routes 
validation : have validation functions 
```

## middleware

```
have all middlewares
```
## models

```
all models and thier migrations
```

## transformers

```
all transformers models to response map and can remove some 
attribute not return to response
```

# Run Test Cases
```
  cmd  cd test 
  cmd go test 
```
# Run project 
```
    cmd go mod download 
    cmd go mod vednor 
    copy env to .env 
    cmd go build 
    cmd ./be-task
```
# Apis
```
  /customers/all for get all customers data 
  /swagger/index.html  for api documention 
```

# Run Docker for BE 
```
cmd docker build -t be-task -f scrpits/be.dockerfile .
cmd docker run  -p 8000:8000  be-test 
open localhost:8000/swagger/index.html
```

# Run Docker for FE
```
cmd docker build -t fe-task -f scrpits/fe.dockerfile .
cmd docker run  -p 3000:3000  be-test 
open localhost:3000 

```

# Run Docker for task 
```
cmd docker compose build 
cmd codker compose run be-task
cmd codker compose run fe-task
open localhost:3000 
```