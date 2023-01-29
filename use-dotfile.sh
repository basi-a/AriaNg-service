#!/bin/bash
if [ ! -d "$HOME/.aria2c" ];then
    mkdir $HOME/.aria2c
fi
config_file="`pwd`/aria2c.conf.example"
ln -sf ${config_file} $HOME/.aria2c/aria2c.conf
