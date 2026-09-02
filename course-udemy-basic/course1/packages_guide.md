# Пакеты в Go: как создавать свои и делать их доступными другим

Гайд по работе с пакетами на примере проекта `course1` (модуль `example.com/bank`).

---

## Часть 1. Что такое пакет

В Go **пакет — это просто папка с `.go`-файлами**, у которых одинаковая строка
`package <имя>` вверху файла.

Главные правила:

- **Папка = пакет.** Все `.go`-файлы в одной папке должны объявлять один и тот же
  `package`.
- **Имя пакета** обычно совпадает с именем папки (но не обязано). Пакет `main` —
  особый: это исполняемая программа с `func main()`.
- **Экспорт по заглавной букве.** Идентификатор виден снаружи пакета, только если
  начинается с **большой** буквы:
  - `func Deposit()` — экспортируется, виден другим пакетам;
  - `func deposit()` — приватный, доступен только внутри своего пакета.
- **Импорт по пути модуля**, а не по имени папки: `module` из `go.mod` + путь к папке.

---

## Часть 2. Создаём свой пакет

Сейчас всё лежит в `package main`. Вынесем работу с балансом в отдельный пакет.

**Структура:**
```
course1/
├── go.mod              // module example.com/bank
├── main.go             // package main
└── fileops/            // новая папка = новый пакет
    └── fileops.go      // package fileops
```

**`course1/fileops/fileops.go`**
```go
package fileops   // имя пакета = имя папки

import (
	"errors"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

// GetBalance читает баланс из файла balance.txt.
// С большой буквы — значит виден снаружи пакета.
func GetBalance() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	return strconv.ParseFloat(string(data), 64)
}

// WriteBalance сохраняет баланс в файл. Тоже экспортируется.
func WriteBalance(balance float64) {
	text := strconv.FormatFloat(balance, 'f', -1, 64)
	_ = os.WriteFile(balanceFile, []byte(text), 0644)
}
```

**`course1/main.go`**
```go
package main

import (
	"fmt"

	"example.com/bank/fileops" // путь = module из go.mod + папка
)

func main() {
	balance, err := fileops.GetBalance() // обращение через имя пакета
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Баланс:", balance)
	fileops.WriteBalance(balance + 100)
}
```

Ключевое:
- путь импорта `example.com/bank/fileops` = `module` из `go.mod` + относительный путь к папке;
- снаружи функция вызывается как `имяпакета.ИмяФункции` → `fileops.GetBalance()`.

### Запуск

`go run .` сам подтянет все импортированные пакеты — отдельно собирать `fileops`
не нужно.

### Частые ошибки при создании пакетов

| Ошибка | Причина |
|---|---|
| `found packages main (a.go) and fileops (b.go) in ...` | в одной папке файлы с **разными** `package` |
| `undefined: fileops.getBalance` | функция с **маленькой** буквы — не экспортируется |
| `package example.com/bank/fileops is not in std` | неверный путь импорта (не совпадает с `module` в `go.mod`) |

---

## Часть 3. Как сделать пакет доступным другим

В Go нет центрального реестра как npm — пакет скачивается **прямо из
git-репозитория**, а путь модуля = адрес репозитория.

### 1. Путь модуля = адрес репозитория

`example.com` — заглушка, по ней Go ничего не скачает. Для публикации путь должен
указывать на реальный репозиторий:

```bash
cd course1
go mod edit -module github.com/yuriy/bank
```

После этого внутренние импорты тоже меняются:
`example.com/bank/fileops` → `github.com/yuriy/bank/fileops`.

### 2. Подготовить пакет

- **Экспортируй** нужное снаружи — функции/типы с **большой буквы**.
- Добавь **doc-комментарии** (начинаются с имени идентификатора).
- Положи **`LICENSE`** (без лицензии другие юридически не могут использовать код)
  и `README.md`.
- Проверь, что `go build ./...` и `go vet ./...` проходят.

> Пакет `package main` импортировать **нельзя** — это исполняемая программа. Для
> библиотеки нужен обычный пакет (`package bank`, `package fileops` и т.п.).

### 3. Запушить в публичный репозиторий
xxxxxччxxxxxxччxxxxxчччxxчxxч
```bash
git init
git add .
git commit -m "bank package"
git remote add origin https://github.com/yuriy/bank.git
git push -u origin main
```

### 4. Поставить версию (git-тег)

Go использует **семантическое версионирование** через git-теги:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Без тега Go подтянет псевдо-версию по последнему коммиту
(`v0.0.0-2026...-abcdef`), но тег — правильный способ.

### 5. Как это используют другие

В другом проекте:
```bash
go get github.com/yuriy/bank@v1.0.0
```
```go
import "github.com/yuriy/bank/fileops"

balance, _ := fileops.GetBalance()
```

Модуль кешируется через прокси `proxy.golang.org`, а документация автоматически
появляется на **pkg.go.dev/github.com/yuriy/bank**.

### 6. Если репозиторий приватный

```bash
go env -w GOPRIVATE=github.com/yuriy/*
```
(и нужен git-доступ к репо — по SSH или токену).

### 7. Поделиться локально, без публикации

Подключить пакет из соседней папки на своём компьютере — через `replace` в
`go.mod` потребителя:
```
require github.com/yuriy/bank v0.0.0
replace github.com/yuriy/bank => ../bank
```

---

## Кратко

1. **Пакет = папка** с файлами, у которых одинаковый `package`.
2. **Экспорт — с большой буквы.**
3. **Импорт** по пути `module + папка`.
4. Чтобы поделиться: **переименуй модуль под адрес репозитория → запушь на
   GitHub → поставь тег `vX.Y.Z`** → другие делают `go get`.
