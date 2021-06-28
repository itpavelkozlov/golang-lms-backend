# golang-lms-backend
Серверная часть обучающей платформы

// ToDo add description here

# Ссылки
- Workflow: [Trello](https://trello.com/b/6s2dfFx6/lms-backend)
- Documentation: [Notion](https://www.notion.so/Backend-docs-20a877f80e824dad962fecd948f20a2f)

# Запуск

### Локально
- генерация зависимостей
```bash
make wire
```
- запуск
```shell
make run
```

### Docker
```shell
docker compose up -d
```

## Миграции
- создать новую
```bash
make migrations-new NAME=<MIGRATION NAME>
```
- накатить текущие
```bash
make migrations-up
```
- откатить текущие
```bash
make migrations-down
```
