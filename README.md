# ~~ptp2022-8-todo-list~~ Slavatidika
## TODO-лист c дополнительной мотивацией, выполненной в игровой форме  
Kапитан: **Александр Старовойтов**


Последняя версия проекта доступна [здесь](https://ptp.starovoytovai.ru)


# Содержание
## Потом перегенерирую
- [ptp2022-8-todo-list](#ptp2022-8-todo-list)
  - [TODO-лист c дополнительной мотивацией, выполненной в игровой форме](#todo-лист-c-дополнительной-мотивацией-выполненной-в-игровой-форме)
- [О проекте:](#о-проекте)
  - [Подробнее о бонусах:](#подробнее-о-бонусах)
  - [Кратко о герое:](#кратко-о-герое)
  - [Документация](#документация)
  - [Что касается стека технологий:](#что-касается-стека-технологий)
    - [Фронтенд](#фронтенд)
    - [Бэкенд](#бэкенд)
    - [Запуск с помощью Yarn](#запуск-с-помощью-yarn)
    - [Запуск с помощью Docker](#запуск-с-помощью-docker)
  - [Наша команда](#наша-команда)

# О проекте:

Все мы хотим быть эффективными и всё успевать, решать несколько задач одновременно и ничего не забывать, поэтому наша команда поставила перед собой задачу предложить решение для повышения личной эффективности в условиях многозадачности.

Одним из вариантов, прорабатываемых нашей командой, **стал TODO-лист с элементами геймификации**.

Todo-лист представляет собой web-приложение для любой платформы. Среди достоинств можно выделить:
- Не требуется установка на устройство
- Все ваши данные хранятся на удаленном сервере(облаке)
- Бесплатный и не содержит рекламы
- Адаптирован под любой размер и ориентацию устройста
- Не требует высоких вычислительных мощностей и может быть открыт на любом чайнике с доступом в интернет
- Имеет интуитивно понятный интерфейс
- Комплексная система создания и сортировки задач
- Несмотря на все достоинства выше, todo-лист не перегружен лишними функциями
- Получил высокую оценку от нашей команды тестирования (им никто не угрожал)
- Имеет открытый исходный код, что позволяет любому человеку вносить свой вклад в улучшение проекта.

## Функциональность списка дел

Включает в себя следующие возможности: 

- Создавать задачи (**ВАУ**)
- Выполнять/Архивировать/Удалять задачи
- Присваивать задачам **свои** категории и сортировать по ним
- За выполнение задач и ежедневный вход получать **Slav~~e~~a коины**
- Отслеживать статистику по завершенным задачам и наблюдать личный прогресс (TODO)
- Связывать задачи между собой (TODO)
- Отслеживать сроки выполнения задач
- Прикреплять к задачам файлы разного формата для упрощения работы над задачей

## Система мотивации

В ходе ее разработки применены элементы геймификации. Это было обусловлено тем, что целевая аудитория продукта - молодые люди в возрасте 12-25 лет,
испытывающие сложности с организацией своего времени и соблюдения распорядка дня и проводящие большую часть своего времени в виртуальных мирах компьютерных игр.

У каждого пользователя есть возможность выбрать своего **персонажа** и прокачивать его, путем выполнения задач, повышения уровня и покупки новых предметов в магазине.

В **магазине** можно приобрести предметы для персонажа за внутриигровую валюту.

**Внутриигровую валюту** можно получить за выполнение задач в срок, ежедневный вход и повышение уровня персонажа.

# Технические составляющие проекта

## Техническое задание

Является немного устаревшим и проект местами не соответствует описанным там идеям и концептам

- [Техническое заданеие](/docs/technical-requirements.md)

## Документация

- [Контракт API](https://ptp.starovoytovai.ru/api/docs)

## Что касается стека технологий:

### Фронтенд

Разрабатывая клиентскую часть приложения, мы постарались на максимум задействовать все доступные нам по условию практики технологии. Поэтому в нашем проекте используются:

- `html` и `scss` для общей вертски страниц.
- `bootstrap` для создания красивых стилей и отзывчивых модальных окон.
- Так же активно применялся `bootstrap grid` для создания адаптивных интерфейсов, которыми можно пользоваться как с десктопных устройств, так и с мобильных телефонов
- `typescript` для создания качественного, безопасного и масштабируемого кода, засчет повсеместного использования типов и классов.
- `parcel` для простой в настройке, но в то же время крайне быстрой сборки проекта. Это позволило быстро верстать страницы с поддержкой `Hot-Reload`(Обновление страницы в браузере сразу же при изменении кода страницы), а так же собирать оптимизированные файлы для деплоя `production` версии.
- `nodeJS` для разработки и создания `production` версии
- `eslint` и `tsc` интегрированные в `Github CI` для автоматической проверки кода, написанного командой разработчиков
- И конечно же `Docker` для удобной разработки и деплоя на любом устройстве в любом месте.

Для хостинга `static` файлов фронтенда написан самодельный `0-dependencies` сервер на `NodeJS`. Сервер умеет поставлять различные страницы по одному адресу в зависимости от значения `ENV_MODE`, для всех страниц реализованы красивые адреса без `.html` в ссылке. Так же реализовано полноценное логирование запросов на сервер вместе с кодом ответа. 

    > TODO: Написать про client-server side рендеринг

### Бэкенд

Говоря о серверной части приложения было принято решение о
реализации архитектуры RESTful API на языке Golang. Т.е весь код по каждому блоку
разделён на 3 части:
- `api` - получение запросов с сервера
- `service` - бизнес-логика
- `repository` - взаимодействие с базой данных 

Почти каждая из частей кода в нашем проекте покрыта `unit-тестами`, что обеспечивает гарантию стабильности кода.  

Мы старались использовать только встроенные средства языка Go и минимизировали
включение сторонних библиотек. В нашем проекте используется:

- язык `Golang` версии `1.17`
- база данных `Postgres`
- библиотека [HttpRouter](https://github.com/julienschmidt/httprouter) для более удобной маршрутизации HTTP-запросов
- библиотека-драйвер [pq](https://github.com/lib/pq) для работы с `Postgres` 
- библиотека [jwt](https://github.com/golang-jwt/jwt) для работы с `JWT-Токенами` при авторизации пользователя 
- библиотека [uuid](github.com/google/uuid) для генерации `uuid` в `ActivationLink` при подтверждении аккаунта по почте


## Запуск проекта

Для запуска `production` версии проекта необходимо:

Скачать [docker](https://docs.docker.com/engine/install/) и `make`.

Склонировать себе репозиторий проекта:

```sh
git clone  git@github.com:bmstu-iu9/ptp2022-8-todo-list.git
cd ptp2022-8-todo-list
```

Запустить `docker-composer`:

```sh
make
```

или

```sh
docker compose up
```

После чего клиент будет доступен на [localhost:8000](http://localhost:8000), а api на [localhost:8080](http://localhost:8080).

Для остановки нажать `CTRL + C` или написать `docker compose down`

## Работа с проектом

### Разработка с помощью Yarn

Удобно использовать при разработке клиентской части приложения

Потребуется `NodeJS`(желательно версии 12.x) и `yarn`:

В корне проекта:

| Команда      | Результат                                                     |
|--------------|---------------------------------------------------------------|
| yarn install | Устанавливает все необходимые локальные пакеты                |
| yarn dev     | Запускает фронтенд в режиме разработки с поддержкой HotReload |
| yarn lint    | Запускает линтер для frontend                                 |
| yarn build   | Упаковывает файлы frontend для деплоя                         |
| yarn deploy  | = `yarn build` + ... + запуск на `localhost:3000`             |

Рекомендуется использовать `yarn lint` или `make lint` перед любым запуском или коммитом вашего кода на JS/TS

Более подробная справка по работе с фронтендом находится в папке `frontend/`, то есть [тут](/frontend/readme.md) 

### Разработка с помощью Docker

Потребуется [docker](https://docs.docker.com/engine/install/) и `make`.

В корне проекта:

| Команда         | Результат                                                                                |
|-----------------|------------------------------------------------------------------------------------------|
| `make`          | Запускает frontend, api и базу данных в режиме `prod`                                    |
| `make dev`      | Запускает frontend, api и базу данных в режиме `dev`                                     |
| `make frontend` | Запускает frontend с realtime обновлениями                                               |
| `make frontend-prod` | Запускает `prod` frontend с подменой стартовой страницы и с работающим `Single Page Application|
| `make lint`     | Запускает линтер для frontend                                                            |

Frontend доступен на [localhost:8000](http://localhost:8000), а api на [localhost:8080](http://localhost:8080).

В папке `backend/`:

| Команда     | Результат                        |
|-------------|----------------------------------|
| `make`      | Запускает локальный сервер api   |
| `make test` | Запускает линтер и тесты для api |
| `make db`   | Запускает базу данных            |

## Наша команда

 - Александр **[@stewkk](https://github.com/stewkk)** Старовойтов — LEAD Backend  
 - Денис **[@OkDenAl](https://github.com/OkDenAl)** Окутин — Backend  
 - Арсений **[@uma-op](https://github.com/uma-op)** Банников — Backend  
 - Георгий **[@geogreck](https://github.com/geogreck)** Гречко — LEAD frontend  (А ещё я написал этот ужасающий README)
 - Вячеслав **[@VyacheslavIsWorkingNow](https://github.com/VyacheslavIsWorkingNow)** Локшин — frontend  
 - Кирилл **[@t1d333](https://github.com/t1d333)** Киселёв —  frontend  
 - Татьяна **[@Tanya-g99](https://github.com/Tanya-g99)** Гнатенко — frontend  
