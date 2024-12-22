# Простой калькулятор на го
Маленький веб сервер, который позволяет производить математические расчеты

# Функционал
поддержка операций +, -, /, * 

# Запуск проекта 
- клонируем репозиторий 
- переходим в консоли в папку проекта
- запускаем go run ./cmd/main.go

# Примеры использования 
curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{ "expression": "2+2*2" }'
Вернет 6.000
curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{ "expression": "2+2*21" }'
error: Expression is not valid
