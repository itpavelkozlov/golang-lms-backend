# golang-lms-backend
Серверная часть обучающей платформы

// ToDo add description here

# Ссылки
- Workflow: [Trello](https://trello.com/b/6s2dfFx6/lms-backend)
- Documentation: [Notion](https://www.notion.so/Backend-docs-20a877f80e824dad962fecd948f20a2f)

# Запуск

### Локально
```bash
go run cmd/lms/main.go --config ./configs/config.yaml
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
