#!/bin/bash
if [ ! -d "${HOME}/.aria2" ];then
    mkdir ${HOME}/.aria2
    touch ${HOME}/.aria2/aria2.session
fi
sed -i "s|/home/---|${HOME}|g" aria2.conf.example
config_file="`pwd`/aria2.conf.example"
ln -sf ${config_file} ${HOME}/.aria2/aria2.conf
