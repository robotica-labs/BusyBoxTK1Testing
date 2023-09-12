#!/bin/bash

cd ../busybox-1.36.1
echo "Script will now open BusyBox Menu Config"
echo "See readme for configuration settings"
read -p "Press enter to continue"
make menuconfig
make
make install CONFIG_PREFIX=../Linux_for_Tegra/rootfs/
cd ../Linux_for_Tegra/rootfs/
mkdir dev etc lib proc tmp sys
cp -dpR TK1_Sample_File_System/dev/* ../Linux_for_Tegra/rootfs/dev/
#Optional device setup would go here