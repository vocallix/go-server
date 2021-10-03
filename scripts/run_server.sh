#!/bin/bash

go build

# 실행에 대한 validation check 존재
if [ $? -ne 0 ]
then
    exit
else
    ./gamedata
fi