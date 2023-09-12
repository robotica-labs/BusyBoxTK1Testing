#!/bin/bash
echo "$(pwd)"
cd ./busybox-1.36.1/
read -p "Press enter to continue"
read -p "stop"
make menuconfig
make
make install CONFIG_PREFIX=../Linux_for_Tegra/rootfs/
cd ../Linux_for_Tegra/rootfs/
sudo mkdir dev etc lib proc tmp sys
sudo cp -dpR TK1_Sample_File_System/dev/* ../Linux_for_Tegra/rootfs/dev/
#Optional device setup would go here