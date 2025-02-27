## Как запускать: выполнить команду "docker compose --env-file ./config.env up".

## Open API документация формируется автоматически из комментариев в коде Go во время сборки образа, описанного в Dockerfile.

## Приложение доступно по адресу "http://localhost:80". Документация - по адресу "http://localhost:80/swagger".

## Эндпоинт "[PUT] /tasks/:id<int>" переводит задачу в следующее состояние и НЕ изменяет значения полей Title и Description. Причина этого не в том, что я не могу реализовать это изменение, а в том, что в задании не было явно сказано, что такое изменение должно происходить, и в том, что я вижу этот эндпоинт именно таким, каким я его реализовал.
