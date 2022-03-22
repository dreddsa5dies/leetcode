#!/bin/bash

echo "Give a directory name to create:"    
read NEW_DIR    
ORIG_DIR=$(pwd)

path_to_files='/home/dreddsa/go/src/github.com/dreddsa5dies/leetcode/scripts/tmp/'

if [ ! -d $path_to_files ]; then
        echo "$path_to_files not found folder";
        exit 1;
fi

mkdir "${ORIG_DIR}/${NEW_DIR}"
for i in $(ls $path_to_files); do
    cp ${path_to_files}${i} "${ORIG_DIR}/${NEW_DIR}"/$i
done