#! /bin/bash

OutType=$1
ROOT=$PWD
DEFAULT_CONF=$ROOT/default.conf

if [[ $OutType == "" ]];then
    echo "useage:   ./Nginx [setup,use] "
fi

if  [[ $OutType == "setup" ]];then
    apt install -y nginx
    cp -r $ROOT/cert /root
fi

if  [[ $OutType == "use" ]];then
    cp $DEFAULT_CONF /etc/nginx/conf.d
    service nginx restart
    systemctl status nginx
fi