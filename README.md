# test_avito
Тестовое задание в юнит Авито.Недвижимость

## Задача

Нужно сделать простой сервис для управления номерами отелей и бронированиями.

* Методы реализуются в виде HTTP-хендлеров, на выходе возвращают JSON в теле ответа. В случае ошибок нужно вернуть текст ошибки (тоже в JSON).

* Сервис должен иметь хендлеры для работы с каталогом номеров отеля:
    * **добавить номер отеля**. Принимает на вход текстовое описание и цену за ночь. Возвращает ID номера отеля.
    * **удалить номер отеля и все его брони**. Принимает на вход ID номера отеля.
    * **получить список номеров отеля**. Должна быть возможность отсортировать по цене или по дате добавления (по возрастанию и убыванию).

* Сервис должен иметь хендлеры для работы со списком броней номеров:
    * **добавить бронь**. Принимает на вход существующий ID номера отеля, дату начала, дату окончания брони (проверять, свободен ли отель на эти даты, не нужно; даты должны быть в формате `“год-месяц-день”`, например: `“2020-01-30”`; даты должны быть валидными). Возвращает ID брони.
    * **удалить бронь**. Принимает на вход ID брони.
    * **получить список броней номера отеля**. Принимает на вход ID номера отеля. Возвращает список бронирований, каждое бронирование содержит ID, дату начала, дату окончания. Бронирования должны быть отсортированы по дате начала.
    
## Документация

https://app.swaggerhub.com/apis/sergejkoll/test_avito/1.0.0#/
