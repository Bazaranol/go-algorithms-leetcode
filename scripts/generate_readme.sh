#!/bin/bash

START="<!-- START_PROBLEMS -->"
END="<!-- END_PROBLEMS -->"

TEMP_BLOCK=$(mktemp)

echo "$START" > $TEMP_BLOCK
echo "" >> $TEMP_BLOCK

for level in easy medium hard
do
  LEVEL_UPPER=$(echo $level | awk '{print toupper(substr($0,1,1)) substr($0,2)}')

  echo "### $LEVEL_UPPER" >> $TEMP_BLOCK
  echo "" >> $TEMP_BLOCK

  for dir in problems/$level/*
  do
    [ -d "$dir" ] || continue

    name=$(basename $dir)

    title=$(echo $name | awk -F_ '{
      for(i=1;i<=NF;i++) {
        printf toupper(substr($i,1,1)) substr($i,2) " "
      }
    }')

    echo "- [$title](./problems/$level/$name)" >> $TEMP_BLOCK
  done

  echo "" >> $TEMP_BLOCK
done

echo "$END" >> $TEMP_BLOCK

# аккуратно обновляем README

TEMP_README=$(mktemp)

inside=0

while IFS= read -r line
do
  if [[ "$line" == "$START" ]]; then
    inside=1
    cat $TEMP_BLOCK >> $TEMP_README
    continue
  fi

  if [[ "$line" == "$END" ]]; then
    inside=0
    continue
  fi

  if [[ $inside -eq 0 ]]; then
    echo "$line" >> $TEMP_README
  fi

done < README.md

mv $TEMP_README README.md
rm $TEMP_BLOCK

echo "README updated"