# golang-lms-backend
Серверная часть обучающей платформы

// ToDo add description here

# Ссылки
- Workflow: [Trello](https://trello.com/b/6s2dfFx6/lms-backend)
- Documentation: [Notion](https://www.notion.so/Backend-docs-20a877f80e824dad962fecd948f20a2f)

# Запуск
```bash
go run cmd/lms/main.go --config ./configs/config.yaml
```

```shell
docker build --no-cache -t test -f ./build/Dockerfile .
docker run -it test
```