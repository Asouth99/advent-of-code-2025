#!/usr/bin/env bash

dayNumber=$1

if [[ -z $dayNumber ]]; then
    echo "Day Number required" >&2
    exit 1
fi

# Make new folder
cp -r template/dayXX ./day$dayNumber

# Update file names
mv day${dayNumber}/dayXX_test.go day${dayNumber}/day${dayNumber}_test.go
mv day${dayNumber}/dayXX.go day${dayNumber}/day${dayNumber}.go

# Update references to dayXX within the files
sed -i -e 's/dayXX/day'${dayNumber}'/g' ./day${dayNumber}/day${dayNumber}.go
sed -i -e 's/dayXX/day'${dayNumber}'/g' ./day${dayNumber}/day${dayNumber}_test.go
sed -i -e 's/template\/day'${dayNumber}'/day'${dayNumber}'/g' ./day${dayNumber}/day${dayNumber}_test.go