Тестовая задача: необходимо создать консольное приложение-сервис, которое принимает HTTP POST запросы:  
1.По пути POST /redis/incr с json вида  
{  
&ensp;&ensp;&ensp;&ensp;“key”: “age”,  
&ensp;&ensp;&ensp;&ensp;“value”: 19  
},  
подключается к redis (хост и порт указываются при запуске в параметрах -host и -port), инкрементирует значение по ключу, указанному в "key" на значение из value", и возвращает его в виде:  
{  
&ensp;&ensp;&ensp;&ensp;“value”: 20  
}  
2.По пути POST /sign/hmacsha512 с json вида  
{  
&ensp;&ensp;&ensp;&ensp;“text”: “test”,  
&ensp;&ensp;&ensp;&ensp;“key”: “test123”  
}  
и возвращает HMAC-SHA512 подпись значения из "text" по ключу "key" в виде hex строки  
3.По пути POST /postgres/users с json вида  
{  
&ensp;&ensp;&ensp;&ensp;“name”: “Alex”,  
&ensp;&ensp;&ensp;&ensp;“age”: 21  
}  
создает в базе данных postgres таблицу users, если она не существует, добавляет в нее строку  
(“Alex”, 21)  
и возвращает идентификатор добавленного пользователя в виде  
{  
&ensp;&ensp;&ensp;&ensp;“id”: 1  
}  
Дополнительные условия:

    Можно использовать любые библиотеки для работы с http, redis и postgres;
    Приложение должно быть протестировано unit-тестами (любой тестовый фреймворк);
    Приложение должно быть протестировано unit-тестами (любой тестовый фреймворк);

На что будем обращать внимание погруженность в тему; процесс поиска решения; процесс поиска решения;

подкючение сторонних библиотек быстро через:
go get github.com/gin-gonic/gin
go get github.com/go-redis/redis
