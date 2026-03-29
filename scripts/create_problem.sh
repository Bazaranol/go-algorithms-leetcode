#!/bin/bash

LEVEL=$1
NAME=$2

DIR="problems/$LEVEL/$NAME"

mkdir -p $DIR

# snake_case → PascalCase
FUNC_NAME=$(echo $NAME | awk -F_ '{for(i=1;i<=NF;i++) printf toupper(substr($i,1,1)) substr($i,2)}')

cat <<EOF > $DIR/solution.go
package $NAME

func $FUNC_NAME() {

}
EOF

cat <<EOF > $DIR/solution_test.go
package $NAME

import "testing"

func Test$FUNC_NAME(t *testing.T) {

}
EOF

cat <<EOF > $DIR/benchmark_test.go
package $NAME

import "testing"

func Benchmark$FUNC_NAME(b *testing.B) {
    for i := 0; i < b.N; i++ {
        $FUNC_NAME()
    }
}
EOF

cat <<EOF > $DIR/README.md
# $FUNC_NAME

## Условие

TODO

---

## Примеры

TODO

---

## Подход

TODO

---

## Сложность

* Time:
* Space:

---

## Идея

TODO

---

## Оригинал условия

[Ссылка](#)
EOF

echo "Created $DIR with function $FUNC_NAME"