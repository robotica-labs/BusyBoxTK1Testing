#!/bin/bash

# Copyright (c) 2011-2015, NVIDIA CORPORATION.  All rights reserved.
#
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions
# are met:
#  * Redistributions of source code must retain the above copyright
#    notice, this list of conditions and the following disclaimer.
#  * Redistributions in binary form must reproduce the above copyright
#    notice, this list of conditions and the following disclaimer in the
#    documentation and/or other materials provided with the distribution.
#  * Neither the name of NVIDIA CORPORATION nor the names of its
#    contributors may be used to endorse or promote products derived
#    from this software without specific prior written permission.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS ``AS IS'' AND ANY
# EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
# PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
# CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
# EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
# PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
# PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY
# OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
# (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


#
# flash.sh: Flash the target board.
#	    flash.sh performs the best in LDK release environment.
#
# Usage: Place the board in recovery mode and run:
#
#	flash.sh [options] <target board> <root_device>
#
#	for more detail enter 'flash.sh -h'
#
# Examples:
# ./flash.sh <target> mmcblk0p1			- boot <target> from eMMC
# ./flash.sh <target> mmcblk1p1			- boot <target> from SDCARD
# ./flash.sh <target> sda1			- boot <target> from USB device
# ./flash.sh -N <IPaddr>:/nfsroot <target> eth0	- boot <target> from NFS
# ./flash.sh -k LNX <target> mmcblk1p1		- update <target> kernel
# ./flash.sh -k EBT <target> mmcblk1p1		- update <target> BCT
#
# Optional Environment Variables:
# BCTFILE ---------------- Boot control table configuration file to be used.
# BOARDID ---------------- Pass boardid to override EEPROM value
# BOOTLOADER ------------- Bootloader binary to be flashed
# BOOTPARTLIMIT ---------- GPT data limit. (== Max BCT size + PPT size)
# BOOTPARTSIZE ----------- Total eMMC HW boot partition size.
# CFGFILE ---------------- Partition table configuration file to be used.
# CMDLINE ---------------- Target cmdline. See help for more information.
# DEVSECTSIZE ------------ Device Sector size. (default = 512Byte).
# DTBFILE ---------------- Device Tree file to be used.
# EMMCSIZE --------------- Size of target device eMMC (boot0+boot1+user).
# FLASHAPP --------------- Flash application running in host machine.
# FLASHER ---------------- Flash server running in target machine.
# IGNOREFASTBOOTCMDLINE -- Block fastboot from filling unspecified kernel
#                          cmdline parameters with its defaults.
# INITRD ----------------- Initrd image file to be flashed.
# ITSFILE ---------------- Multi image u-boot package template file.
# KERNEL_IMAGE ----------- Linux kernel zImage file to be flashed.
# MTS -------------------- MTS file name such as mts_si.
# MTSPREBOOT ------------- MTS preboot file name such as mts_preboot_si.
# NFSARGS ---------------- Static Network assignments.
#			   <C-ipa>:<S-ipa>:<G-ipa>:<netmask>
# NFSROOT ---------------- NFSROOT i.e. <my IP addr>:/exported/rootfs_dir.
# ODMDATA ---------------- Odmdata to be used.
# ROOTFSSIZE ------------- Linux RootFS size (internal emmc/nand only).
# ROOTFS_DIR ------------- Linux RootFS directory name.
# TEGRABOOT -------------- lowerlayer bootloader such as nvtboot.bin.
# UBOOTSCRIPT ------------ U-boot HUSH boot script file.
# UBOOT_TEXT_BASE -------- U-boot Image Load Address.
# UIMAGE_LABEL ----------- Kernel version for U-boot image header.
# UIMAGE_NAME ------------ uImage file name.
# WB0BOOT ---------------- Warmboot code such as nvtbootwb0.bin
#
chkerr ()
{
	if [ $? -ne 0 ]; then
		if [ "$1" != "" ]; then
			echo "$1";
		else
			echo "failed.";
		fi;
		exit 1;
        fi;
	if [ "$1" = "" ]; then
		echo "done.";
	fi;
}

pr_conf()
{
	if [ "${zflag}" != "true" ]; then
		return 0;
	fi;
	echo "target_board=${target_board}";
	echo "target_rootdev=${target_rootdev}";
	echo "rootdev_type=${rootdev_type}";
	echo "rootfssize=${rootfssize}";
	echo "odmdata=${odmdata}";
	echo "flashapp=${flashapp}";
	echo "flasher=${flasher}";
	echo "bootloader=${bootloader}";
	echo "tegraboot=${tegraboot}";
	echo "wb0boot=${wb0boot}";
	echo "mtspreboot=${mtspreboot}";
	echo "mts=${mts}";
	echo "ubootscript=${ubootscript}";
	echo "bctfile=${bctfile}";
	echo "cfgfile=${cfgfile}";
	echo "kernel_image=${kernel_image}";
	echo "rootfs_dir=${rootfs_dir}";
	echo "nfsroot=${nfsroot}";
	echo "nfsargs=${nfsargs}";
	echo "kernelinitrd=${kernelinitrd}";
	echo "cmdline=${cmdline}";
	echo "boardid=${boardid}";
	exit 0;
}

validateIP ()
{
	local ip=$1;
	local ret=1;

	if [[ $ip =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
		OIFS=${IFS};
		IFS='.';
		ip=($ip);
		IFS=${OIFS};
		[[ ${ip[0]} -le 255 && ${ip[1]} -le 255 && \
		   ${ip[2]} -le 255 && ${ip[3]} -le 255 ]];
		ret=$?;
	fi;
	if [ ${ret} -ne 0 ]; then
		echo "Invalid IP address: $ip";
		exit 1;
	fi;
}

netmasktbl=(\
	"255.255.255.252" \
	"255.255.255.248" \
	"255.255.255.240" \
	"255.255.255.224" \
	"255.255.255.192" \
	"255.255.255.128" \
	"255.255.255.0" \
	"255.255.254.0" \
	"255.255.252.0" \
	"255.255.248.0" \
	"255.255.240.0" \
	"255.255.224.0" \
	"255.255.192.0" \
	"255.255.128.0" \
	"255.255.0.0" \
	"255.254.0.0" \
	"255.252.0.0" \
	"255.248.0.0" \
	"255.240.0.0" \
	"255.224.0.0" \
	"255.192.0.0" \
	"255.128.0.0" \
	"255.0.0.0" \
);

validateNETMASK ()
{
	local i;
	local nm=$1;
	for (( i=0; i<${#netmasktbl[@]}; i++ )); do
		if [ "${nm}" = ${netmasktbl[$i]} ]; then
			return 0;
		fi;
	done;
	echo "Error: Invalid netmask($1)";
	exit 1;
}

validateNFSargs ()
{
	local a=$2;

	OIFS=${IFS};
	IFS=':';
	a=($a);
	IFS=${OIFS};

	if [ ${#a[@]} -ne 4 ]; then
		echo "Error: Invalid nfsargs($2)";
		exit 1;
	fi;
	validateIP ${a[0]};
	ipaddr=${a[0]};
	if [ "${serverip}" = "" ]; then
		validateIP ${a[1]};
		serverip=${a[1]};
	fi;
	validateIP ${a[2]};
	gatewayip=${a[2]};
	validateNETMASK ${a[3]};
	netmask=${a[3]};
	if [ "$1" != "" ]; then
		eval "$1=$2";
	fi;
	return 0;
}

validateNFSroot ()
{
	if [ "$2" = "" ]; then
		return 1;
	fi;
	OIFS=${IFS};
	IFS=':';
	local var=$1;
	local a=($2);
	IFS=${OIFS};
	if [ ${#a[@]} -ne 2 ]; then
		echo "Error: Invalid nfsroot($2)";
		exit 1;
	fi;
	validateIP ${a[0]};
	if [[ "${a[1]}" != /* ]]; then
		echo "Error: Invalid nfsroot($2)";
		exit 1;
	fi;
	tftppath=${a[0]}:/tftpboot/${uimage_name};
	tftpfdtpath=${a[0]}:/tftpboot/${dtbfilename};
	if [ "${serverip}" = "" ]; then
		serverip=${a[0]};
	fi;
	eval "${var}=$2";
	return 0;
}

usage ()
{
	state=$1
	retval=$2

	if [[ $state == allunknown ]]; then
		echo -e "
Usage: sudo ./flash.sh [options] <target board> <rootdev>
  Where,
	target board: Valid target board name.
	rootdev: Proper root device."

	elif [[ $state == rootdevunknown ]]; then
		echo -e "
Usage: sudo ./flash.sh [options] ${target_board} <rootdev>
  Where,
    rootdev for ${target_board}:
	${ROOT_DEV}"

	else
		echo "
Usage: sudo ./flash.sh [options] ${target_board} ${target_rootdev}"
	fi;

	cat << EOF
    options:
        -b <bctfile> --------- nvflash boot control table config file.
        -c <cfgfile> --------- nvflash partition table config file.
        -d <dtbfile> --------- device tree file.
        -e <emmc size> ------- Target device's eMMC size.
        -f <flashapp> -------- Path to flash application: nvflash or tegra-rcm.
        -h ------------------- print this message.
        -i ------------------- pass user kernel commandline as-is to kernel.
        -k <partition id> ---- partition name or number specified in flash.cfg.
        -m <mts preboot> ----- MTS preboot such as mts_preboot_si.
        -n <nfs args> -------- Static nfs network assignments
                               <Client IP>:<Server IP>:<Gateway IP>:<Netmask>
        -o <odmdata> --------- ODM data.
        -p <bp size> --------- Total eMMC HW boot partition size.
        -r ------------------- skip building and reuse existing system.img.
        -s <ubootscript> ----- HUSH bootscript file for U-Boot.
        -t <tegraboot> ------- tegraboot binary such as nvtboot.bin
        -u <dbmaster> -------- PKC server in <user>@<IP address> format.
        -w <wb0boot> --------- warm boot binary such as nvtbootwb0.bin
        -x <tegraid> --------- 0x40 for jetson-tk1.
        -y <fusetype> -------- PKC for secureboot, NS for non-secureboot.
        -z <sn> -------------- Serial Number of target board.
        -C <cmdline> --------- Kernel commandline arguments.
                               WARNING:
                               Each option in this kernel commandline gets
                               higher preference over the same option from
                               fastboot. In case of NFS booting, this script
                               adds NFS booting related arguments, if -i option
                               is omitted.
        -F <flasher> --------- Flash server such as fastboot.bin.
        -I <initrd> ---------- initrd file. Null initrd is default.
        -K <kernel> ---------- Kernel image file such as zImage or Image.
        -L <bootloader> ------ Bootloader such as fastboot.bin or u-boot.bin.
        -M <mts boot> -------- MTS boot file such as mts_si.
        -N <nfsroot> --------- i.e. <my IP addr>:/my/exported/nfs/rootfs.
        -P <end of PPT + 1> -- Primary GPT start address + size of PPT + 1.
        -R <rootfs dir> ------ Sample rootfs directory.
        -S <size> ------------ Rootfs size in bytes. Valid only for internal
                               rootdev. KiB, MiB, GiB short hands are allowed,
                               for example, 1GiB means 1024 * 1024 * 1024 bytes.
        -T <its file> -------- ITS file name. Valid only for u-boot.
	-B <boardid> --------- BoardId.
EOF
	exit $retval;
}

setdflt ()
{
	local var="$1";
	if [ "${!var}" = "" ]; then
		eval "${var}=$2";
	fi;
}

setval ()
{
	local var="$1";
	local val="$2";
	if [ "${!val}" = "" ]; then
		echo "Error: missing $val not defined.";
		exit 1;
	fi;
	eval "${var}=${!val}";
}

mkfilesoft ()
{
	local var="$1";
	local varname="$1name";

	eval "${var}=$2";
	if [ "${!var}" = "" -a "$3" != "" -a -f "$3" ]; then
		eval "${var}=$3";
	fi;
	if [ "${!var}" != "" ]; then
		if [ ! -f ${!var} ]; then
			echo "Warning: missing $var (${!var}), continue... ";
			return 1;
		fi;
		eval "${var}=`readlink -f ${!var}`";
		eval "${varname}=`basename ${!var}`";
	fi;
}

mkfilepath ()
{
	local var="$1";
	local varname="$1name";

	eval "${var}=$2";
	setdflt "${var}" "$3";
	if [ "${!var}" != "" ]; then
		eval "${var}=`readlink -f ${!var}`";
		if [ ! -f "${!var}" ]; then
			echo "Error: missing $var (${!var}).";
			usage allknown 1;
		fi;
		eval "${varname}=`basename ${!var}`";
	fi;
}

mkdirpath ()
{
	local var="$1";
	eval "${var}=$2";
	setdflt "$1" "$3";
	if [ "${!var}" != "" ]; then
		eval "${var}=`readlink -f ${!var}`";
		if [ ! -d "${!var}" ]; then
			echo "Error: missing $var (${!var}).";
			usage allknown 1;
		fi;
	fi;
}

getsize ()
{
	local var="$1";
	local val="$2";
	if [[ ${!val} != *[!0-9]* ]]; then
		eval "${var}=${!val}";
	elif [[ (${!val} == *KiB) && (${!val} != *[!0-9]*KiB) ]]; then
		eval "${var}=$(( ${!val%KiB} * 1024 ))";
	elif [[ (${!val} == *MiB) && (${!val} != *[!0-9]*MiB) ]]; then
		eval "${var}=$(( ${!val%MiB} * 1024 * 1024 ))";
	elif [[ (${!val} == *GiB) && (${!val} != *[!0-9]*GiB) ]]; then
		eval "${var}=$(( ${!val%GiB} * 1024 * 1024 * 1024))";
	else
		echo "Error: Invalid $1: ${!val}";
		exit 1;
	fi;
}

validatePartID ()
{
	local idx=0;
	declare -A cf;

	while read aline; do
		if [ "$aline" != "" ]; then
			arr=( $(echo $aline | tr '=' ' ') );
			if [ ${arr[0]} = "name" ]; then
				cf[$idx,1]=${arr[1]};
			elif [ ${arr[0]} = "id" ]; then
				cf[$idx,0]=${arr[1]};
				idx=$((idx+1));
			fi;
	fi;
	done < $4;

	for ((i = 0; i < idx; i++)) do
		if [ "$3" = ${cf[$i,0]} -o "$3" = ${cf[$i,1]} ]; then
			eval "$1=${cf[$i,0]}";
			eval "$2=${cf[$i,1]}";
			return 0;
		fi;
	done;
	echo "Error: invalid partition id ($3)";
	exit 1;
}

cp2local ()
{
	local src=$1;
	if [ "${!src}" = "" ]; then return 1; fi;
	if [ ! -f "${!src}" ]; then return 1; fi;
	if [ "$2" = "" ];      then return 1; fi;
	if [ -f $2 -a ${!src} = $2 ]; then
		local sum1=`sum ${!src}`;
		local sum2=`sum $2`;
		if [ "$sum1" = "$sum2" ]; then
			echo "Existing ${src}($2) reused.";
			return 0;
		fi;
	fi;
	echo -n "copying ${src}(${!src})... ";
	cp -f ${!src} $2;
	chkerr;
	return 0;
}

build_fsimg ()
{
	local loop_dev="/dev/loop0"
	echo "Making $1... ";
	# On some linux distributions, loop device is not created by default
	# In such cases, create one manually
	if [ ! -b ${loop_dev} ]; then
		loop_dev=`losetup --find`
		if [ "$?" != "0" ]; then
			echo "Cannot find loop device. Terminating..";
			exit 1;
		fi
	fi
	umount ${loop_dev} > /dev/null 2>&1;
	losetup -d ${loop_dev} > /dev/null 2>&1;
	rm -f $1;	chkerr "clearing $1 failed.";
	rm -rf mnt;	chkerr "clearing $4 mount point failed.";

	local bcnt=$(( $3 / 512 ));
	local bcntdiv=$(( $3 % 512 ));
	if [ $bcnt -eq 0 -o $bcntdiv -ne 0 ]; then
		echo "Error: $4 file system size has to be 512 bytes allign.";
		exit 1;
	fi
	if [ "$2" != "" -a "$2" != "0" ]; then
		local fc=`printf '%d' $2`;
		local fillc=`printf \\\\$(printf '%02o' $fc)`;
		< /dev/zero head -c $3 | tr '\000' ${fillc} > $1;
		chkerr "making $1 with fillpattern($fillc}) failed.";
	else
		truncate --size $3 $1;
		chkerr "making $1 with zero fillpattern failed.";
	fi;
	losetup ${loop_dev} $1 > /dev/null 2>&1;
	chkerr "mapping $1 to loop device failed.";
	if [ "$4" = "FAT32" ]; then
		mkfs.msdos -I -F 32 ${loop_dev} > /dev/null 2>&1;
	else
		mkfs -t $4 ${loop_dev} > /dev/null 2>&1;
	fi;
	chkerr "formating $4 filesystem on $1 failed.";
	mkdir -p mnt;		chkerr "make $4 mount point failed.";
	mount ${loop_dev} mnt;	chkerr "mount $1 failed.";
	mkdir -p mnt/boot/dtb;	chkerr "make $1/boot/dtb failed.";
	cp -f ${kernel_image} mnt/boot;
	chkerr "Copying ${kernel_image} failed.";
	if [ -f "${dtbfile}" ]; then
		cp -f ${dtbfile} mnt/boot/dtb;
		chkerr "populating ${dtbfile} to $1 failed.";
		cp -f ${dtbfile} mnt/boot/dtb/tegra.dtb;
		chkerr "populating ${dtbfile} to $1/boot/dtb/tegra.dtb failed.";
	fi;
	if [ "$4" = "FAT32" ]; then
		touch -f mnt/boot/cmdline.txt > /dev/null 2&>1;
		chkerr "Creating cmdline.txt failed.";
	fi;
	if [ "$5" != "" ]; then
		pushd mnt > /dev/null 2>&1;
		echo -n -e "\tpopulating rootfs from $5 ... ";
		(cd $5; tar cf - *) | tar xf - ; chkerr;
		popd > /dev/null 2>&1;
	fi;
	echo -e -n "\tSync'ing $1 ... ";
	sync; sync; sleep 5;	# Give FileBrowser time to terminate gracefully.
	echo "done.";
	umount ${loop_dev} > /dev/null 2>&1;
	losetup -d ${loop_dev} > /dev/null 2>&1;
	rmdir mnt > /dev/null 2>&1;
	if [ "$2" != "" -a -x mksparse ]; then
		echo -e "\tConverting RAW image to Sparse image... ";
		mv -f $1 $1.raw;
		./mksparse -v --fillpattern=$2 $1.raw $1; chkerr;
	fi;
	echo "$1 built successfully. ";
}

if [ $# -lt 2 ]; then
	usage allunknown 1;
fi;

# if the user is not root, there is not point in going forward
THISUSER=`whoami`
if [ "$THISUSER" != "root" ]; then
    echo "flash.sh requires root privilege"
    exit 1
fi
argv=($@);
target_board=${argv[$#-2]};
target_rootdev=${argv[$#-1]};

if [ ! -r ${target_board}.conf ]; then
	echo "Error: Invalid target board - ${target_board}."
	usage allunknown 1;
fi
source ${target_board}.conf

# Determine rootdev_type
#
rootdev_type="external";
if [[ "${target_rootdev}" == mmcblk0p* ]]; then
	rootdev_type="internal";
elif [ "${target_rootdev}" = "eth0" -o "${target_rootdev}" = "eth1" ]; then
	rootdev_type="network";
elif [[ "${target_rootdev}" != mmcblk1p* && \
	"${target_rootdev}" != sd* ]]; then
	echo "Error: Invalid target rootdev($target_rootdev).";
	usage rootdevunknown 1;
fi;

opstr+="b:c:d:e:f:hik:m:n:o:p:rs:t:u:w:x:y:z:B:C:F:I:K:L:M:N:P:R:S:T:Z:";
while getopts "${opstr}" OPTION
do
	case $OPTION in
	b) BCTFILE=${OPTARG}; ;;
	c) CFGFILE=${OPTARG}; ;;
	d) DTBFILE=${OPTARG}; ;;
	e) EMMCSIZE=${OPTARG}; ;;
	f) FLASHAPP=${OPTARG}; ;;
	h) usage allunknown 0; ;;
	i) IGNOREFASTBOOTCMDLINE="ignorefastboot"; ;;
	k) target_partname=${OPTARG}; ;;	# cmdline only
	m) MTSPREBOOT=${OPTARG}; ;;
	n) NFSARGS=${OPTARG}; ;;
	o) ODMDATA=${OPTARG}; ;;
	p) BOOTPARTSIZE=${OPTARG}; ;;
	r) reuse_systemimg="true"; ;;		# cmdline only
	s) UBOOTSCRIPT=${OPTARG}; ;;
	t) TEGRABOOT=${OPTARG}; ;;
	u) dbmaster="${OPTARG}"; ;;
	w) WB0BOOT=${OPTARG}; ;;
	x) tegraid=${OPTARG}; ;;
	y) fusetype=${OPTARG}; ;;
	z) sn=${OPTARG}; ;;
	B) BOARDID=${OPTARG}; ;;
	C) CMDLINE="${OPTARG}"; ;;
	F) FLASHER=${OPTARG}; ;;
	I) INITRD=${OPTARG}; ;;
	K) KERNEL_IMAGE=${OPTARG}; ;;
	L) BOOTLOADER=${OPTARG}; ;;
	M) MTS=${OPTARG}; ;;
	N) NFSROOT=${OPTARG}; ;;
	P) BOOTPARTLIMIT=${OPTARG}; ;;
	R) ROOTFS_DIR=${OPTARG}; ;;
	S) ROOTFSSIZE=${OPTARG}; ;;
	T) ITSFILE=${OPTARG}; ;;
	Z) zflag="true"; ;;			# cmdline only
	*) usage allunknown 1; ;;
	esac
done

###########################################################################
# System default values: should be defined AFTER target_board value.
#
LDK_DIR=$(cd `dirname $0` && pwd)
LDK_DIR=`readlink -f "${LDK_DIR}"`
BL_DIR="${LDK_DIR}/bootloader";
TARGET_DIR="${BL_DIR}/${target_board}"
KERNEL_DIR="${LDK_DIR}/kernel";
export PATH="${KERNEL_DIR}:${PATH}";		# preferrence on our DTC
DTB_DIR="${KERNEL_DIR}/dtb";
ROOTFS_TYPE="${ROOTFS_TYPE:-ext4}";
DEVSECTSIZE="${DEVSECTSIZE:-512}";		# default sector size = 512
BOOTPARTLIMIT="${BOOTPARTLIMIT:-10485760}";	# 1MiB limit
fillpat="${FSFILLPATTERN:-0}";			# no cmdline: default=0
cmdline="${IGNOREFASTBOOTCMDLINE} ${CMDLINE_ADD} ${CMDLINE} ";
boardid="${BOARDID}";
###########################################################################
# System mandatoriy vars:
#
setval     odmdata	ODMDATA;	# .conf mandatory
setval     devsectsize	DEVSECTSIZE;
setval     rootfs_type	ROOTFS_TYPE;
getsize    rootfssize	ROOTFSSIZE;	# .conf mandatory
getsize    emmcsize	EMMCSIZE;	# .conf mandatory
getsize    bootpartsize	BOOTPARTSIZE;	# .conf mandatory
getsize    bootpartlim	BOOTPARTLIMIT;
mkfilepath flashapp	"${FLASHAPP}"	"${BL_DIR}/nvflash";
mkfilepath flasher	"${FLASHER}"	"${TARGET_DIR}/fastboot.bin";
mkfilepath bootloader	"${BOOTLOADER}"	"${TARGET_DIR}/fastboot.bin";
mkdirpath  rootfs_dir	"${ROOTFS_DIR}"	"${LDK_DIR}/rootfs";
mkfilepath kernel_image	"$KERNEL_IMAGE" "${KERNEL_DIR}/${DFLT_KERNEL:-zImage}";
mkfilepath bctfile	"${BCTFILE}"	"${TARGET_DIR}/BCT/${EMMC_BCT}";
mkfilepath cfgfile	"${CFGFILE}"	"${TARGET_DIR}/cfg/${EMMC_CFG}";
mkfilepath dtbfile	"${DTBFILE}"	"${DTB_DIR}/${DTB_FILE}";

mkfilesoft kernelinitrd	"${INITRD}"	"";
mkfilesoft tegraboot	"${TEGRABOOT}"	"${TARGET_DIR}/nvtboot.bin";
mkfilesoft wb0boot	"${WB0BOOT}"	"${TARGET_DIR}/nvtbootwb0.bin";
mkfilesoft mtspreboot	"${MTSPREBOOT}"	"${BL_DIR}/mts_preboot_si";
mkfilesoft mts		"${MTS}"	"${BL_DIR}/mts_si";

if [ "${rootdev_type}" = "network" ]; then
	if [ "${NFSROOT}" = "" -a "${NFSARGS}" = "" ]; then
		echo "Error: network argument(s) missing.";
		usage allknown 1;
	fi;
	if [ "${NFSROOT}" != "" ]; then
		validateNFSroot nfsroot "${NFSROOT}";
	fi;
	if [ "${NFSARGS}" != "" ]; then
		validateNFSargs nfsargs "${NFSARGS}";
	fi;
	if [ "${nfsroot}" != "" ]; then
		nfsdargs="root=/dev/nfs rw netdevwait";
		cmdline+="${nfsdargs} ";
		if [ "${nfsargs}" != "" ]; then
			nfsiargs="ip=${nfsargs}";
			nfsiargs+="::${target_rootdev}:off";
		else
			nfsiargs="ip=:::::${target_rootdev}:on";
		fi;
		cmdline+="${nfsiargs} ";
		cmdline+="nfsroot=${nfsroot} ";
	fi;
fi;

if [ "${bootloadername}" = "u-boot.bin" ]; then
	if [ "${rootdev_type}" = "network" ]; then
		UBOOTBSF="${TARGET_DIR}/${NET_BSF}";
	else
		UBOOTBSF="${TARGET_DIR}/${EMMC_BSF}";
	fi;
	DFLT_ITSFILE="${TARGET_DIR}/${target_board}_kernel_fdt.its";
	if [ "${SYSBOOTFILE}" != "" ]; then
		if [ "${rootdev_type}" = "network" ]; then
			SYSBOOTFILE="${TARGET_DIR}/${SYSBOOTFILE}.nfs";
		elif [[ "${target_rootdev}" == mmcblk1p* ]]; then
			SYSBOOTFILE="${TARGET_DIR}/${SYSBOOTFILE}.sdcard";
		elif [[ "${target_rootdev}" == sd* ]]; then
			SYSBOOTFILE=${TARGET_DIR}/"${SYSBOOTFILE}.usb";
		else
			SYSBOOTFILE="${TARGET_DIR}/${SYSBOOTFILE}.emmc";
		fi;
	fi;
	mkfilesoft sysbootfile	"${SYSBOOTFILE}";
	mkfilesoft ubootscript	"${UBOOTSCRIPT}" "${UBOOTBSF}";
	mkfilesoft itsfile	"${ITSFILE}"	 "${DFLT_ITSFILE}";
	setdflt uboot_text_base "${UBOOT_TEXT_BASE}";
	setdflt uimage_label	"${UIMAGE_LABEL:-Linux-tegra}";
	setdflt uimage_name	"${UIMAGE_NAME:-vmlinux.uimg}";
fi;

##########################################################################
pr_conf;	# print config and terminate if requested.
##########################################################################

pushd $BL_DIR > /dev/null 2>&1;

### Localize files and build TAGS ########################################
# BCT_TAG:::
#
cp2local bctfile "${BL_DIR}/${bctfilename}";

# EBT_TAG:
#
cp2local bootloader "${BL_DIR}/${bootloadername}";
EBT_TAG+="-e s/filename=fastboot.bin/filename=${bootloadername}/ ";
EBT_TAG+="-e s/\#filename=fastboot.bin/filename=${bootloadername}/ ";

# LNX_TAG:
#
localbootfile=boot.img;
if [ "${bootloadername}" = "u-boot.bin" ]; then
	if [ "${sysbootfile}" != "" -a -f "${sysbootfile}" ]; then
		mkdir -p ${rootfs_dir}/boot > /dev/null 2>&1;
		echo -e -n "\tpopulating kernel to rootfs... ";
		cp -f ${kernel_image} ${root_dir}/boot; chkerr;
		echo -e -n "\tpopulating ${sysbootfilename} to rootfs... ";
		mkdir -p ${rootfs_dir}/boot/extlinux; chkerr;
		cp -f ${sysbootfile} ${rootfs_dir}/boot/extlinux/extlinux.conf; chkerr;
		LNX_TAG+="-e /filename=${localbootfile}/d ";
	else
		echo -e -n "\tpopulating kernel to rootfs... ";
		if [ "${itsfile}" != "" -a -f "${itsfile}" -a \
		     "${dtbfile}" != "" -a -f "${dtbfile}" ]; then
			ITSCONV="";
			ITSCONV+="-e s/gzip/none/ ";
			ITSCONV+="-e s?vmlinux.bin.gz?${kernel_image}? ";
			ITSCONV+="-e s/LOADADDR/${uboot_text_base}/ ";
			ITSCONV+="-e s/ENTRYPOINT/${uboot_text_base}/ ";
			ITSCONV+="-e s/LinuxKernel/${uimage_label}/ ";
			ITSCONV+="-e s/target.dtb/${dtbfilename}/ ";
			cat ${itsfile} | sed ${ITSCONV} > ${itsfilename};
			MKUARG="-f ${itsfilename}";
		else
			MKUARG+="-A arm ";
			MKUARG+="-O linux ";
			MKUARG+="-T kernel ";
			MKUARG+="-C none ";
			MKUARG+="-a ${uboot_text_base} ";
			MKUARG+="-e ${uboot_text_base} ";
			MKUARG+="-n ${uimage_label} ";
			MKUARG+="-d ${kernel_image} ";
		fi;
		mkdir -p ${rootfs_dir}/boot;
		rm -f ${rootfs_dir}/boot/${uimage_name};
		mkimageapp=${LDK_DIR}/bootloader/mkimage;
		${mkimageapp} ${MKUARG} ${rootfs_dir}/boot/${uimage_name};
		chkerr;
		echo -n "generating boot script (${ubootscript}) ... ";
		if [ "${ubootscript}" != "" -a -f ${ubootscript} ]; then
			NFSCONV=" -e s/bootdelay=3/bootdelay=2/ ";
			NFSCONV+="-e s/IPADDR/${ipaddr}/ ";
			NFSCONV+="-e s/SERVERIP/${serverip}/ ";
			NFSCONV+="-e s/GATEWAYIP/${gatewayip}/ ";
			NFSCONV+="-e s/NETMASK/${netmask}/ ";
			NFSCONV+="-e s%TFTPPATH%${tftppath}% ";
			NFSCONV+="-e s%TFTPFDTPATH%${tftpfdtpath}% ";
			NFSCONV+="-e s/NFSARGS/${nfsiargs}/ ";
			NFSCONV+="-e s%NFSROOT%${nfsroot}% ";
			if [ "$dtbfile" != "" ]; then
				NFSCONV+="-e s%DTBFILENAME%${dtbfilename}% ";
			fi;
			cat ${ubootscript} | \
				sed ${NFSCONV} > ${ubootscriptname};
			./mkubootscript -i ${ubootscriptname} \
					-o ${localbootfile};
			chkerr;
		else
			LNX_TAG+="-e /filename=${localbootfile}/d ";
			echo "Missing. Using embedded bootscript... ";
		fi;
	fi;
else
	if [ "$kernelinitrd" != "" -a -f "$kernelinitrd" ]; then
		echo -n "copying initrd(${kernelinitrd})... ";
		cp -f ${kernelinitrd} initrd;
	else
		echo "making zero initrd... ";
		rm -f initrd; touch initrd;
	fi;
	chkerr;
	echo -n "Making Boot image... "
	MKBOOTARG=" --kernel ${kernel_image} ";
	MKBOOTARG+="--ramdisk initrd ";
	MKBOOTARG+="--board ${target_rootdev} ";
	MKBOOTARG+="--output ${localbootfile}";
	./mkbootimg ${MKBOOTARG} --cmdline "${cmdline}" > /dev/null 2>&1;
	chkerr;
fi;

# SOS_TAG:
#
localsosfile=recovery.img;
#
# XXX: recovery is yet to be implemented.
#
if [ -f ${localsosfile} ]; then
	SOS_TAG="-e s/#filename=recovery.img/filename=${localsosfile}/";
fi;

# NVC_TAG:
#
if [ "${tegraboot}" != "" ]; then
	cp2local tegraboot "${BL_DIR}/${tegrabootname}";
	NVC_TAG="-e s/#filename=${tegrabootname}/filename=${tegrabootname}/ ";
	NVC_TAG+="-e s/type=data\s\+#TEGRABOOT/type=bootloader/ ";
fi;

# MPB_TAG:
#
if [ "${mtspreboot}" != "" ]; then
	cp2local mtspreboot "${BL_DIR}/${mtsprebootname}";
	MPB_TAG="-e s/#filename=mts_preboot_si/filename=${mtsprebootname}/ ";
	MPB_TAG+="-e s/type=data\s\+#MTSPREBOOT/type=mts_preboot/ ";
	MTSARGS+="--preboot ${mtsprebootname} ";
fi;

# MBP_TAG:
#
if [ "${mts}" != "" ]; then
	cp2local mts "${BL_DIR}/${mtsname}";
	MBP_TAG="-e s/#filename=mts_si/filename=${mtsname}/ ";
	MBP_TAG+="-e s/type=data\s\+#MTSBOOTPACK/type=mts_bootpack/ ";
	MTSARGS+="--bootpack ${mtsname} ";
fi;

# APP_TAG:
#
localsysfile=system.img;
APP_TAG+="-e s/size=1073741824/size=${rootfssize}/ ";
if [ "${reuse_systemimg}" = "true" ]; then
	echo "Reusing existing ${localsysfile}... ";
	if [ ! -e "${localsysfile}" ]; then
		echo "file does not exist.";
		exit 1;
	fi;
	echo "done.";
elif [ "${rootdev_type}" = "internal" -o "${target_partname}" = "APP" ]; then
	build_fsimg "$localsysfile" "$fillpat" \
		    "$rootfssize" "$rootfs_type" "$rootfs_dir";
elif [ "${rootdev_type}" = "network" -a \
       "${bootloadername}" = "u-boot.bin" -a \
       "${sysbootfile}" != "" -a -f "${sysbootfile}" ]; then
	echo -n "generating /boot/extlinux/extlinux.conf files... ";
	NFSCONV="-e s/NFSARGS/${nfsiargs}/ ";
	NFSCONV+="-e s%NFSROOT%${nfsroot}% ";
	sed ${NFSCONV} < ${rootfs_dir}/boot/extlinux/extlinux.conf > ./extlinux.conf;
	mv ./extlinux.conf ${rootfs_dir}/boot/extlinux/extlinux.conf;
	echo "done.";

	echo "generating system.img for network booting... ";
	tmpdir=`mktemp -d`;
	mkdir -p ${tmpdir}/boot/extlinux > /dev/null 2>&1;
	cp -f ${rootfs_dir}/boot/extlinux/extlinux.conf ${tmpdir}/boot/extlinux > /dev/null 2>&1;
	cp -f ${kernel_image} ${tmpdir}/boot > /dev/null 2>&1;
	cp -f ${dtbfile} ${tmpdir}/boot > /dev/null 2>&1;
	build_fsimg "$localsysfile" "$fillpat" \
		    "$rootfssize" "$rootfs_type" "$tmpdir";
else
	APP_TAG+="-e /filename=system.img/d ";
fi;

# DTB_TAG: localize
#
if [ "${dtbfile}" != "" ]; then
	cp2local dtbfile "${BL_DIR}/${dtbfilename}";
	DTB_TAG="-e s/#filename=tegra.dtb/filename=${dtbfilename}/ ";
else
	DTB_TAG="-e s/name=DTB/name=DTX/ ";
fi;

# EFI_TAG: Minimum FAT32 partition size is 64MiB (== 1 FAT cluster)
#
localefifile=efi.img;
efifs_size=$(( 64 * 1024 * 1024 ));
EFI_TAG="-e s/size=67108864\s\+#EFISIZE/size=${efifs_size}/ ";
if [ "${bootloadername}" = "uefi.bin" ]; then
	build_fsimg $localefifile "" $efifs_size "FAT32" "";
	EFI_TAG+="-e s/#filename=efi.img/filename=${localefifile}/";
fi;

# WB0_TAG:
#
if [ "${wb0boot}" != "" ]; then
	cp2local wb0boot "${BL_DIR}/${wb0bootname}";
	WB0_TAG="-e s/#filename=nvtbootwb0.bin/filename=${wb0bootname}/ ";
	WB0_TAG+="-e s/type=data\s\+#WB0BOOT/type=WB0/ ";
fi;

# GPT_TAG: tag should created before cfg and actual img should be
#	   created after cfg.
#
localpptfile=ppt.img;
localsptfile=gpt.img;
if [ "${bootpartsize}" != "" ]; then
	bplmod=$(( ${bootpartlim} % ${devsectsize} ));
	if [ ${bplmod} -ne 0 ]; then
		echo "Error: Boot partition limit is not modulo ${devsectsize}";
		exit 1;
	fi;
	bpsmod=$(( ${bootpartsize} % ${devsectsize} ));
	if [ ${bpsmod} -ne 0 ]; then
		echo "Error: Boot partition size is not modulo ${devsectsize}";
		exit 1;
	fi;
	gptsize=$(( ${bootpartlim} - ${bootpartsize} ));
	if [ ${gptsize} -lt ${devsectsize} ]; then
		echo "Error: No space for primary GPT.";
		exit 1;
	fi;
	GPT_TAG="-e s/size=2097152\s\+#BCTSIZE/size=${bootpartsize}/ ";
	GPT_TAG+="-e s/size=8388608\s\+#PPTSIZE/size=${gptsize}/ ";
	GPT_TAG+="-e s/#filename=ppt.img/filename=${localpptfile}/ ";
	GPT_TAG+="-e s/#filename=spt.img/filename=${localsptfile}/ ";
fi;

# CFG:
#
localcfgfile=flash.cfg;
echo -n "copying cfgfile(${cfgfile}) to ${localcfgfile}... ";
CFGCONV+="${EBT_TAG} ";
CFGCONV+="${LNX_TAG} ";
CFGCONV+="${SOS_TAG} ";
CFGCONV+="${NVC_TAG} ";
CFGCONV+="${MPB_TAG} ";
CFGCONV+="${MBP_TAG} ";
CFGCONV+="${WB0_TAG} ";
CFGCONV+="${APP_TAG} ";
CFGCONV+="${EFI_TAG} ";
CFGCONV+="${DTB_TAG} ";
CFGCONV+="${GPT_TAG} ";
cat ${cfgfile} | sed ${CFGCONV} > ${localcfgfile}; chkerr;

# GPT:
if [ ! -z "${bootpartsize}" -a ! -z "${emmcsize}" ]; then
	echo "creating gpt(${localpptfile})... ";
	MKGPTOPTS="-c ${localcfgfile} -P ${localpptfile} ";
	MKGPTOPTS+="-t ${emmcsize} -b ${bootpartsize} -s 4KiB ";
	MKGPTOPTS+="-a GPT -v GP1 ";
	MKGPTOPTS+="-V ${MKGPTCMD} ";
	./mkgpt ${MKGPTOPTS};
	chkerr "creating gpt(${localpptfile}) failed.";
fi;

# FLASH:
#
cp2local flasher	"${BL_DIR}/${flashername}";
cp2local flashapp	"${BL_DIR}/${flashappname}";
if [ "${target_partname}" != "" ]; then
	validatePartID target_partid target_partname $target_partname $cfgfile;
	if [ "${flashername}" = "u-boot.bin" ]; then
		# XXX: when external MSD scheme is available, implement it.
		echo "Error: Update not supported for u-boot yet.";
		exit 1;
	fi;
	tmp_updateid="${target_partid}:${target_partname}";
	case ${target_partname} in
	BCT) target_partfile="${bctfilename}";
	     FLASHARGS="--bct ${target_partfile} --updatebct SDRAM "; ;;
	PPT) target_partfile="${localpptfile}"; ;;
	EBT) target_partfile="${bootloadername}"; ;;
	LNX) target_partfile="${localbootfile}"; ;;
	SOS) target_partfile="${localsosfile}"; ;;
	NVC) target_partfile="${tegrabootname}"; ;;
	MPB) target_partfile="${mtsprebootname}"; ;;
	MBP) target_partfile="${mtsname}"; ;;
	APP) target_partfile="${localsysfile}"; ;;
	DTB) target_partfile="${dtbfilename}"; ;;
	EFI) target_partfile="${localefifile}"; ;;
	WB0) target_partfile="${wb0bootname}"; ;;
	GPT) target_partfile="${localsptfile}"; ;;
	*)   echo "*** Update ${tmp_updateid} is not supported. ***";
	     exit 1; ;;
	esac;
	echo "*** Updating ${tmp_updateid} with ${target_partfile} ***";
	if [ "${FLASHARGS}" = "" ]; then
		FLASHARGS="--download ${target_partid} ${target_partfile} ";
	fi;
	FLASHARGS+="${MTSARGS} --bl ${flashername} --go";
	echo "./${flashappname} ${FLASHARGS}";
	./${flashappname} ${FLASHARGS};
	chkerr "Failed to flash ${target_board}."
	echo "*** The ${tmp_updateid} has been updated successfully. ***"
	exit 0;
fi;

if [ -f odmsign.func ]; then
	source odmsign.func;
	odmsign;
	if [ $? -ne 0 ]; then
		exit 1;
	fi;
fi;

echo "*** Flashing target device started. ***"
if [ "${boardid}" != "" ]; then
	FLASHARGS="--boardid $boardid";
fi;
FLASHARGS+=" --bct ${bctfilename} --setbct --configfile ${localcfgfile} ";
FLASHARGS+="${MTSARGS} --create --bl ${flashername} --odmdata $odmdata --go";
echo "./${flashappname} ${FLASHARGS}";
./${flashappname} ${FLASHARGS};
chkerr "Failed flashing ${target_board}.";
echo "*** The target ${target_board} has been flashed successfully. ***"
if [ "${rootdev_type}" = "internal" ]; then
	echo "Reset the board to boot from internal eMMC."
elif [ "${rootdev_type}" = "network" ]; then
	if [ "${nfsroot}" != "" ]; then
		echo -n "Make target nfsroot(${nfsroot}) exported ";
		echo "on the network and reset the board to boot";
	else
		echo -n "Make the target nfsroot exported on the network, ";
		echo -n "configure your own DHCP server with ";
		echo -n "\"option-root=<nfsroot export path>;\" ";
		echo "properly and reset the board to boot";
	fi;
else
	echo -n "Make the target filesystem available to the device ";
	echo "and reset the board to boot from external ${target_rootdev}."
fi;
echo
exit 0;
