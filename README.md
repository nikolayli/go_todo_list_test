# go_todo_list_test

клонируем репозиторий:
```git clone https://github.com/nikolayli/go_todo_list_test.git```

открываем папку с проектом:
```cd go_todo_list```

инициализациия базы данных:
```psql -U ваш_логин -h ваш_хост -W -f db/init.sql'```

управление зависимостями:
```go mod tidy```

запуск программы:
```make run```