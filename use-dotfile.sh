#!/bin/bash
if [ ! -d "$HOME/.config/aria2c" ];then
    mkdir $HOME/.config/aria2c
fi
config_file="`pwd`/aria2c.conf.example"
ln -sf ${config_file} $HOME/.config/aria2c/aria2c.conf
