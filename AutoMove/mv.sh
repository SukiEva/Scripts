#!/data/data/com.termux/files/usr/bin/bash

# 工作目录
workDir="/storage/emulated/0/Documents/AutoMove"
sdCardList=("/sdcard/" "/storage/emulated/0/" "/data/media/0/")


logFile="${workDir}/run.log"
configFile="${workDir}/config.prop"

logd(){
   echo "[$(date '+%g/%m/%d %H:%M')] | $@" >> $logFile
}

hasPrefix(){ # 检查配置文件前缀
    local has="false"
    for sdcard in ${sdCardList[@]}
    do  
        if [[ $1 == ${sdcard}* ]]; then
            has="true"
        fi
    done
    echo $has
    return $?
}

addSuffix(){ # 添加文件夹后缀
    local path=$1
    if [[ $path != */ ]]; then
        path="$path/"
    fi
    echo $path
    return $?
}


readConfig(){ # 读取配置文件
    if [[ ! -e "$configFile" ]]; then
        logd "配置文件不存在"
    else
        local IFS=$'\n'
        local idx=0
        for line in `cat $configFile | grep -v '#'`
        do  
            if [[ $(hasPrefix $line) == "true" ]]; then
                clearList[$idx]=$line
                ((idx++))
            fi
        done
    fi
}

move(){ # 移动文件
    for list in ${clearList[@]}
    do  
        if [[ $list =~ "&" ]]; then
            local srcPath=$(addSuffix ${list%&*})
            local dstPath=$(addSuffix ${list#*&})
            for file in `ls $srcPath`
            do
                mv -f "${srcPath}${file}" $dstPath
                logd "$file ==> $dstPath"
            done
        fi
    done
}

readConfig
move