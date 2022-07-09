# ptp2022-8-todo-list
## TODO-лист c дополнительной мотивацией, выполненной в игровой форме  
Kапитан: **Александр Старовойтов**

[Последняя версия проекта](https://ptp.starovoytovai.ru)

- [ptp2022-8-todo-list](#ptp2022-8-todo-list)
  - [TODO-лист c дополнительной мотивацией, выполненной в игровой форме](#todo-лист-c-дополнительной-мотивацией-выполненной-в-игровой-форме)
- [О проекте:](#о-проекте)
  - [Подробнее о бонусах:](#подробнее-о-бонусах)
  - [Кратко о герое:](#кратко-о-герое)
  - [Что касается стека технологий:](#что-касается-стека-технологий)
    - [Фронтенд](#фронтенд)
    - [Запуск с помощью Yarn](#запуск-с-помощью-yarn)
    - [Бэкенд](#бэкенд)
    - [Запуск с помощью Docker](#запуск-с-помощью-docker)
  - [Наша команда](#наша-команда)

О проекте:
===

Все мы хотим быть продуктивными, поэтому наша команда
>(в процессе ###______________: 10%)
>
разработала TODO-приложение, в котором вы сможете:

+ Создавать задачи
+ Соотносить задачи с датами
+ Группировать таски в проекты
+ ПОЛУЧАТЬ бонусы (бесплатно!)
+ За бонусы прокачивать своего персонажа
+
+ ... (в разработке)
+

Подробнее о бонусах:
---
Они характеризуются двумя группами — **единовременные и накопительные**.  
**Единовременные** пользователь получает за выполнение задачи и за выполнение всех задач за день
при этом будет введен дополнительный коэффициент - если планирование задачи попадает в диапазон от дня до месяца,
то он домножается на число >1, тем самым поощряя пользователя планировать свои дела, а потом выполнять.  
**Накопительные** бонусы будут даваться за каждодневный вход.

>(Задача MVP+ а также за выполнение ачивок).

Кратко о герое:
---

Каждый пользователь может кастомизировать своего персонажа внешне;
Добавлять новые элементы одежды, менять черты лица и фигуру,
открывать новые аксессуары и украшения.
Сам персонаж будет неразрывно связан с пользователем.

>(Задача MVP+ добавить возможность улучшать характеристики персонажа)

---
## Что касается стека технологий:

### Фронтенд

- `html`
- `scss`
- `typescript + CommonJs`
- `parcel`
- `nodeJS` для разработки

### Бэкенд

RESTful API на Go.

### Запуск с помощью Yarn

Потребуется `NodeJS`(желательно версии 12.x), `npm`, `yarn` и `parcel`:
```sh
    sudo npm install -g yarn
    sudo npm install -g parcel
```
В корне проекта:

| Команда      | Результат                                                     |
|--------------|---------------------------------------------------------------|
| yarn install | устанавливает все необходимые локальные пакеты                |
| yarn dev     | Запускает фронтенд в режиме разработки с поддержкой HotReload |
| yarn lint    | Запускает линтер для frontend                                 |
| yarn start   | Собирает файлы frontend для деплоя                            |

Рекомендуется использовать `yarn lint` или `make lint` перед любым запуском или коммитом вашего кода на JS/TS

Более подробная справка по работе с фронтендом находится в папке `frontend/`

### Запуск с помощью Docker

Потребуется [docker](https://docs.docker.com/engine/install/) и `make`.

В корне проекта:

| Команда         | Результат                                                                                |
|-----------------|------------------------------------------------------------------------------------------|
| `make`          | Запускает frontend, api и базу данных, при этом в realtime обновляется *только frontend* |
| `make frontend` | Запускает frontend с realtime обновлениями                                               |
| `make lint`     | Запускает линтер для frontend                                                            |

Frontend доступен на [localhost:8000](http://localhost:8000), а api на [localhost:8080](http://localhost:8080).

В папке `backend/`:

| Команда     | Результат                        |
|-------------|----------------------------------|
| `make`      | Запускает локальный сервер api   |
| `make test` | Запускает линтер и тесты для api |

## Наша команда

Александр **[@stewkk](https://github.com/stewkk)** Старовойтов — LEAD Backend  
Денис **[@OkDenAl](https://github.com/OkDenAl)** Окутин — Backend  
Арсений **[@uma-op](https://github.com/uma-op)** Банников — Backend  
Георгий **[@geogreck](https://github.com/geogreck)** Гречко — LEAD frontend  
Вячеслав **[@VyacheslavIsWorkingNow](https://github.com/VyacheslavIsWorkingNow)** Локшин — frontend  
Кирилл **[@t1d333](https://github.com/t1d333)** Киселёв —  frontend  
Татьяна **[@Tanya-g99](https://github.com/Tanya-g99)** Гнатенко — frontend  
