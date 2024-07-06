# Текстовое задание для Effective Mobile

В данном проекте я реализовал api для task-tracker.

## Запуск приложения
Запуск миграции зависимостей

```console
make dep
```

Запуск проекта локально

```console
make run-tracker
```

Запуск тестов

```console
make test
```

Сборка проекта

```console
make build-tacker
```

Запуск с hot reload

```console
air
```

## Core library

| Library    | Usage             |
| ---------- | ----------------- |
| gin        | Base framework    |
| database/sql | SQL library       |
| postgres   | Database          |
| logrus     | Logger library    |
| viper      | Config library    |