#!/bin/bash

dir_names=$1
path_to_files='/home/dreddsa/go/src/github.com/dreddsa5dies/leetcode/scripts/tmp/'

if [ ! -d $path_to_files ]; then
        echo "$path_to_files not found folder";
        exit 1;
fi

if [ -d $dir_names ]; then
        echo "$dir_names is exists";
        exit 1;
fi


echo "Creating $i and copying over files..."
mkdir $dir_names
for i in $(ls $path_to_files); do
    cp -rf ${path_to_files}${i} ${dir_names}/${dir_names}-$i
done