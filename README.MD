Based on https://elinux.org/Jetson/Busybox_RootFS

Requirements
Ubuntu OS, tested on Ubuntu 20.04.6 LTS 64-bit
Jetson Developer Board. Examples are TK1, TX1, and TX2. This has only been tested with Nvidia Jetson Orin Nano Development Kit with ARMv8 Processor.

Ncurses Library
apt-get install libncurses5-dev libncursesw5-dev

Required Libraries:
Ncurses Library
sudo apt-get install libncurses5-dev libncursesw5-dev
ARM GCC compiler
sudo apt-get install gcc-arm*

BusyBox menuconfig
After when running the config.sh script the user is required to change the following setting in BusyBox Configurator "make menuconfig" select Settings --> Build Options, then "Build static binary"  Next, set "Cross Compiler prefix" to arm-linux-gnueabihf- Exit the menu, and be sure to save your changes.