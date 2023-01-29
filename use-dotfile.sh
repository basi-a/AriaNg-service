#!/bin/bash
if [ ! -d "${HOME}/.aria2" ];then
    mkdir ${HOME}/.aria2
    touch ${HOME}/.aria2/aria2.session
fi
config_file="`pwd`/aria2.conf.example"

sed -i "s|/home/---|${HOME}|g" ${config_file}

ln -sf ${config_file} ${HOME}/.aria2/aria2.conf
