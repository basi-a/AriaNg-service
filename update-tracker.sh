#!/bin/bash
if [ -f "best_aria2.txt" ];then
    rm best_aria2.txt
fi
wget https://raw.githubusercontent.com/XIU2/TrackersListCollection/master/best_aria2.txt -q --show-progress
list="`cat best_aria2.txt`"
sed -i "s\bt-tracker=\bt-tracker=${list}\g" aria2.conf.example
