# Music Library API

REST API сервис для управления музыкальной библиотекой, написанный на Go с использованием PostgreSQL в качестве базы данных и Redis для кэширования.

## Технологии

- Go
- Fiber (веб-фреймворк)
- PostgreSQL (GORM)
- Redis
- Docker & Docker Compose
- JWT (аутентификация)
- Swagger (документация API)
- Viper (конфигурация)
- Logrus (логирование)

## Предварительные требования

- Docker
- Docker Compose
- Go 1.21 или выше

## API Documentation

Полная API документация доступна через Swagger UI по адресу: `http://localhost:8080/swagger/`

### Доступные эндпоинты
#### Аутентификация

| Метод | Эндпоинт      | Описание                    | Требует авторизации |
|-------|---------------|----------------------------|-------------------|
| POST  | /auth/sign-up | Регистрация нового пользователя | Нет |
| POST  | /auth/sign-in | Вход в систему (получение JWT токена) | Нет |

#### Музыкальные треки (Songs)

| Метод | Эндпоинт      | Описание                    | Требует авторизации |
|-------|---------------|----------------------------|-------------------|
| GET   | /api/songs    | Получить все треки пользователя | Да |
| POST  | /api/songs    | Добавить новый трек        | Да |
| GET   | /api/songs/:id | Получить трек по ID        | Да |
| PUT   | /api/songs/:id | Обновить информацию о треке | Да |
| DELETE| /api/songs/:id | Удалить трек               | Да |

## Установка и запуск

### Использование Docker Compose

1. Клонируйте репозиторий:
```bash
git clone https://github.com/MDmitryM/music-lib-go.git
cd music-lib-go
```

2. Создайте файл `.env` в корневой директории проекта:
```env
# Database
POSTGRES_USER=dbuser
POSTGRES_DB=db
DB_PASSWORD=your_password

# Redis
REDIS_DB_PASSWORD=your_redis_password

# Application
SIGNING_KEY=your_jwt_signing_key
ENV=production
```

3. Запустите приложение:
```bash
docker-compose up --build
```

Приложение будет доступно по адресу `http://localhost:8080`

### Переменные окружения

| Переменная    | Описание                           |
|---------------|------------------------------------|
| POSTGRES_USER | Имя пользователя PostgreSQL        |
| POSTGRES_DB   | Название базы данных               |
| DB_PASSWORD   | Пароль для базы данных             |
| REDIS_DB_PASSWORD | Пароль для Redis               |
| SIGNING_KEY   | Ключ для подписи JWT токенов       |
| ENV           | Окружение (development/production) |

### Порты

- 8080: API сервер
- 5439: PostgreSQL (доступен на хосте)
- 6379: Redis (доступен на хосте)

## Структура проекта

- `/cmd` - точка входа в приложение
- `/pkg` - основной код приложения
  - `/handler` - обработчики HTTP запросов
  - `/repository` - работа с базами данных
  - `/service` - бизнес-логика
- `/configs` - конфигурационные файлы
- `/docs` - документация API (Swagger)

## База данных

PostgreSQL база данных автоматически инициализируется при первом запуске. Данные сохраняются в директории `.database/postgres/data`.

### Конфигурация

Основные настройки приложения находятся в:
- `.env` - переменные окружения (учетные данные БД, ключи)
- `configs/config.yml` - конфигурация приложения (порты, хосты)
