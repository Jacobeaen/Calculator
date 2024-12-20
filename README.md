# HTTP Калькулятор
Перед работой убедитесь, что у вас установлен [golang](https://go.dev/dl/) и [curl](https://curl.se/download.html)  
Также для POST запросов можно использовать [Postman](https://www.postman.com/downloads/)  
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
Либо сразу запустите _main.go_: `go run Calclator/cmd/main.go`  
Вам надо будет послать POST запрос на 8080 порт сервера:
1. Пример с помощью _curl_:

   - `curl -L '127.0.01:8080/api/v1/calculate' -H 'Content-Type: application/json' -d '{"expression": "2+2"}'`
2. Пример с помощью _postman_:

    - Откройте postman и создайте POST запрос
    - В **URL** введите _127.0.01:8080/api/v1/calculate_
    - В **Body** выберите _raw_ и установите формат _json_
    - В **Json** вставьте `{"expression": "2+2"}`

В зависимости от запроса сервер может вернуть 3 HTTP ответа с телами:
1. `"result": "результат выражения"` и код 200, если выражение вычислено успешно
2. `"error": "Expression is not valid"` и код 422, если в выражении не правильно
3. `"error": "Internal server error"` и код 500, если была зарегестрирована ошибка, не учитываемая в пункте 2

Примеры запросов для сервера и ожидаемые ответы:
| ✅Правильные запросы | Ответы сервера|
|--------------------  |-------|
| Это случайный текст |Ответ №1|
| Пример случайного текста |Ответ №2|
| Еще один случайный текст |Ответ №3|