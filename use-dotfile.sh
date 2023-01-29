#!/bin/bash
if [ ! -d "$HOME/.aria2c" ];then
    mkdir $HOME/.aria2c
fi

ln -sf aria2c.conf.example $HOME/.aria2c/aria2c.conf