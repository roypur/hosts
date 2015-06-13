#!/system/xbin/bash
mount -o remount,rw /system
sed -i 's/127.0.0.1/::1/g' /system/etc/hosts >> /system/etc/hosts
