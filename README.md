# HTTP Калькулятор
Перед работой убедитесь, что у вас установлен [golang](https://go.dev/dl/) и [curl](https://curl.se/download.html)  
Также можете использовать [Postman](https://www.postman.com/downloads/)
После этого:
  1. Перейдите в папку, куда вы хотите сколонировать репозиторий, например ```C:/Code/LMS```
  2. Скопируйте репозиторий командой ```git clone https://github.com/Jacobeaen/Calculator.git```
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
Перейдите в папку сmd: ```cd Calculator/cmd``` и запустите main.go: ```go run main.go``` 

