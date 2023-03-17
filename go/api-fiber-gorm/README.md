# Fiber with Auth JWT and Gorm

### Url pattern FLAT

[PEP 20 – The Zen of Python](https://peps.python.org/pep-0020/#id3)  

Flat is better than nested. 

### !Important

When first deploy uncomment in database/connect.go  

    // DB.Create(&model.UserType{Role: "normal"}) // init database !import!

### Development

Fist clone the project, then run the next commands inside del folder project

    docker compose up -d

if you use docker compose app (go project in container) but if you just run the container for database then run the go server  
    
    go run main.go
    
consume the api, check uris in file *router/router.go*
