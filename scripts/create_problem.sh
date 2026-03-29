#!/bin/bash

LEVEL=$1
NAME=$2

DIR="problems/$LEVEL/$NAME"

mkdir -p $DIR

cat <<EOF > $DIR/solution.go
package $NAME

func Solution() {

}
EOF

cat <<EOF > $DIR/solution_test.go
package $NAME

import "testing"

func TestSolution(t *testing.T) {

}
EOF

cat <<EOF > $DIR/README.md
# $NAME

## 🧩 Условие

TODO

## 🏷️ Tags

TODO
EOF

echo "Created $DIR"