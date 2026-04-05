.PHONY: test bench lint run run-bench new readme cover

# Запуск всех тестов
test:
	go test ./...

# Запуск всех бенчмарков
bench:
	go test ./... -bench=. -benchmem

# Линтер
lint:
	golangci-lint run

# Запуск тестов конкретной задачи по имени: make run name=two_sum
run:
	@dir=$$(find problems -type d -name "$(name)" | head -1); \
	if [ -z "$$dir" ]; then echo "Problem '$(name)' not found"; exit 1; fi; \
	echo "Running tests in $$dir"; \
	go test ./$$dir -v

# Запуск бенчмарка конкретной задачи: make run-bench name=two_sum
run-bench:
	@dir=$$(find problems -type d -name "$(name)" | head -1); \
	if [ -z "$$dir" ]; then echo "Problem '$(name)' not found"; exit 1; fi; \
	echo "Running benchmark in $$dir"; \
	go test ./$$dir -bench=. -benchmem

# Создание новой задачи: make new level=easy name=two_sum
new:
	./scripts/create_problem.sh $(level) $(name)

# Автогенерация списка задач в README
readme:
	./scripts/generate_readme.sh

# Покрытие тестами
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
