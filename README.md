<h1>Калькулятор с HTTP</h1>

<p>Перед работой убедитесь, что у вас установлен компилятор языка - <a href="https://go.dev/dl/">Golang</a> и утилита <a href="https://curl.se/download.html">curl</a></p>
<ol>
    <li>Перейдите в папку, куда вы хотите сколонировать репозиторий, например C:/Code/LMS</li>
    <li>Скопируйте репозиторий командой git clone https://github.com/Jacobeaen/Calculator.git</li>
</ol>
<p>После клонирования репозитория у вас будет вот такая картина:</p>
```
LMS
├───main.go (1881b)
├───main_test.go (1318b)
└───Calculator
	├───pkg
	│	├───rpn (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)
```
<p>Перейдите в сmd: cd Calculator/cmd и запустите main.go: go run main.go</main></p>                 
<p>После того, как вы запустили главный файл проекта вам нужно будет передать запрос серверу с помощью утилиты curl или postman.</p>