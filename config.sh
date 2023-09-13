#!/bin/bash
echo "$(pwd)"
cd ./busybox-1.36.1/
make menuconfig
make
make install CONFIG_PREFIX=../Linux_for_Tegra/rootfs/
cd ../Linux_for_Tegra/rootfs/
mkdir dev etc lib proc tmp sys
cp -dpR  TK1_Sample_File_System/dev/ Linux_for_Tegra/rootfs/dev/

Linux_for_Tegra/rootfs/dev/
TK
"proc  /proc proc  defaults  0 0"
"none  /tmp  ramfs defaults  0 0"
"mdev  /dev  ramfs defaults  0 0"
"sysfs /sys  sysfs defaults  0 0"

#Optional device setup would go here