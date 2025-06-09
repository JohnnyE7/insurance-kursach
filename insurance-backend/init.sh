#!/bin/bash

echo "Выполняем инициализацию..."

# Пример: дождаться PostgreSQL
until pg_isready -h db -p 5432 -U johnny; do
  echo "⏳ Ожидаем базу данных..."
  sleep 1
done

echo "База готова, запускаем приложение..."
./app
