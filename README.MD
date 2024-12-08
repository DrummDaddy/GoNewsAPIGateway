# Проект сервисов

## API Gateway

API Gateway служит для маршрутизации запросов от клиентов к различным внутренним сервисам. Он поддерживает фильтрацию новостей по названию и пагинацию.

- **Путь к исходному коду:** `./api-gateway`
- **Порт:** 8080
- **Dockerfile:** Включён в директорию службы.

### Как запустить

1. Перейдите в папку, содержащую Dockerfile API Gateway:
   cd api-gateway 

2. Соберите DOCKER образ: 
docker build -t api-gateway 

3. Запустите контейнер: 
docker run -p 8080:8080 api-gateway