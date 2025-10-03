# URL Shortener

Простой сервис для сокращения URL на Go с возможностью редиректа по короткой ссылке.

---

## Структура проекта
```
.
├── cmd/
│ └── shortener # основной бинарник сервера
├── internal/
│ ├── app/ # инициализация приложения
│ ├── handler/ # HTTP-хендлеры
│ ├── router/ # маршрутизация
│ ├── service/ # бизнес-логика (URLService)
│ ├── repository/ # хранилище URL (MemoryURLRepo)
│ └── model/ # структуры данных
├── README.md
├── shortenertest # бинарник автотестов
```
---

## Установка и сборка

1. Клонируем проект:

```bash
git clone <URL_твоего_репозитория>
cd go-musthave-shortener-tpl/cmd/shortener
```
2. Сборка сервера:
```bash
go build -o shortener
```
3. Проверяем права на выполнение:
```bash
chmod +x shortener
```
Запуск сервера
```bash
./shortener
```
Сервер слушает http://localhost:8080.

## API

1. Сокращение URL (POST /)

Content-Type: text/plain

Тело запроса: исходный URL

Ответ: 201 Created, короткий URL в теле (text/plain)

Пример:
```bash
    curl -i -X POST http://localhost:8080/ \
     -H "Content-Type: text/plain" \
     -d "https://practicum.yandex.ru/"
```
Ответ:

```bash
HTTP/1.1 201 Created
Content-Type: text/plain
http://localhost:8080/Hqn0PQ
```

2. Редирект по короткому URL (GET /{id})

Метод: GET

Путь: /EwHXdJfB — идентификатор короткой ссылки

Ответ: 307 Temporary Redirect

Location: оригинальный URL

Пример:

```bash
curl -i http://localhost:8080/Hqn0PQ
```
Ответ:
```bash
HTTP/1.1 307 Temporary Redirect
Location: https://practicum.yandex.ru/
```

## Структура кода

App — инициализация сервисов и репозитория.

Router — настройка маршрутов и http.Server.

URLService — бизнес-логика, генерация короткого ID и получение оригинального URL.

MemoryURLRepo — хранение URL в памяти (map).

ShortenHandler — обработка POST-запросов.

RedirectUrlHandler — обработка GET-запросов.

## Тестирование

1. Сборка автотестов:
```bash
chmod +x shortenertest
```

2. Запуск теста первой итерации:
```bash
./shortenertest -test.v -test.run='^TestIteration1$' -binary-path=./shortener
```