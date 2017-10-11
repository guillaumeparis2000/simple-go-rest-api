# Simple Rest Api test with GO

## DB installation
To install the database you just need to start a docker container like this one:
```sh
docker run -d --name mysql-container -p 3306:3306 -e MYSQL_ROOT_PASSWORD=MyPassword -d mysql:latest
```

Create the todo database:
```sql
CREATE SCHEMA `todos` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
CREATE TABLE `todo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `completed` tinyint(1) DEFAULT '0',
  `due` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```

## Build executable
```sh
go build
```

## Start rest Api
```sh
./rest-api
```

## Available methos
| Path  | Method  | Action  |
|---|---|---|---|---|
| /todos  | GET  | List all todos  |
| /todos  | POST  | Create a new Todo  |
| /todos/{id}  | GET  | Show a todo detail  | 
| /todos/{id}  | PUT  | Update a Todo  |
