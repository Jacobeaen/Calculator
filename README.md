# HTTP Калькулятор
Перед работой убедитесь, что у вас установлен [golang](https://go.dev/dl/) и [curl](https://curl.se/download.html)  
Также можете использовать [Postman](https://www.postman.com/downloads/)  
После этого:
 1. Перейдите в папку, куда вы хотите сколонировать репозиторий, например `C:/Code/LMS`
 2. Скопируйте репозиторий командой `git clone https://github.com/Jacobeaen/Calculator.git`
```
LMS
└───Calculator/
    ├───pkg/
    │    └───rpn/
    │        ├── calc.go/          
    │        ├── validate.go       
    │        └── rpn_test.go
    ├───internal
    │        └───application/
    │            └───handlers.go
    ├───cmd/
    │    └───main.go
    ├───go.mod
    └───README.md
```
Перейдите в папку _сmd_: `cd Calculator/cmd` и запустите _main.go_: `go run main.go`  
Вам надо будет послать POST запрос на 8080 порт сервера:
1. Пример с помощью _curl_:
  - `curl -L '127.0.01:8080/api/v1/calculate' -H 'Content-Type: application/json' -D '{"expression": "2+2"}'`
2. Пример с помощью _postman_:
  - Откройте postman и создайте POST запрос
  - В URL введите _127.0.01:8080/api/v1/calculate_
  - В Body выберите _raw_ и установите формат _json_
  - В Json вставьте `{"expression": "2+2"}`
