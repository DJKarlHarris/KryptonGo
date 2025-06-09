#!/bin/bash
# 重建数据库脚本（适用MySQL 8.0+）
# 执行权限：chmod +x rebuild_db.sh
# 用法：./rebuild_db.sh [数据库名] [SQL文件路径]

DB_NAME="${1:-test_db}"  # 默认数据库名
SQL_FILE_NAME="${2:-db_full.sql}" # SQL文件名
BACKUP_DIR="../../db/backup"
DATE=$(date +%Y%m%d%H%M%S)
PASSWD="123456"
USER="root"
IP="172.29.224.1"

SQL_FILE_PATH="../../db/${SQL_FILE_NAME}" # SQL文件路径

#全量构建
full_build_sql() {
  echo "⚙️ 全量构建，重建数据库 $DB_NAME..."
  mysql -h $IP -u $USER -p$PASSWD<<EOF
DROP DATABASE IF EXISTS $DB_NAME;
CREATE DATABASE $DB_NAME 
  CHARACTER SET utf8mb4 
  COLLATE utf8mb4_general_ci;
  USE $DB_NAME;
  SOURCE $SQL_FILE_PATH;
EOF
}

full_build() {
  full_build_sql   
}

#增量构建
alter_build_sql() {
    echo "⚙️ 增量构建，重建数据库 $DB_NAME..."
    mysql -h $IP -u $USER -p$PASSWD<<EOF
USE $DB_NAME;
SOURCE $SQL_FILE_PATH;
EOF
}

alter_build() {
    read -p "⚠️当前为Alter build，是否继续执行？(y/n): " choice
    case "$choice" in
        [yY]|[yY][eE][sS])  # 匹配 y/Y/yes/YES
            alter_build_sql
            ;;
        [nN]|[nN][oO])      # 匹配 n/N/no/NO
            echo "操作取消。"
            exit 1
            ;;
        *)                  # 其他输入
            echo "无效输入，操作终止。"
            exit 1
            ;;
    esac
}

TYPE=$(echo ${SQL_FILE_NAME} | awk -F '[_]' '{print $2}' | awk -F '[.]' '{print $1}') #变更类型

if [[ "alter" = "${TYPE}" ]]; then
  alter_build    
elif [[ "full" = "${TYPE}" ]]; then
  full_build
else
    echo "unkown type ${TYPE}"
    exit 1
fi

if [ $? -ne 0 ]; then
    echo "❌ MySQL 执行失败！"
    exit 1
fi

#todo备份现有数据库
#echo "▶ 备份数据库 $DB_NAME..."
#touch $BACKUP_DIR/${DB_NAME}_${DATE}.sql
#mysqlDump -h $IP -u $USER -p $PASSWD --single-transaction --routines $DB_NAME > $BACKUP_DIR/${DB_NAME}_${DATE}.sql

# 导入表结构及数据

echo "✅ 构建完成！备份文件：$BACKUP_DIR/${DB_NAME}_${DATE}.sql"