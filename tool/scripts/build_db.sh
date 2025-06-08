#!/bin/bash
# 重建数据库脚本（适用MySQL 8.0+）
# 执行权限：chmod +x rebuild_db.sh
# 用法：./rebuild_db.sh [数据库名] [SQL文件路径]

DB_NAME="${1:-test_db}"  # 默认数据库名
SQL_FILE="${2:-../../db/db.sql}" # 默认SQL文件路径
BACKUP_DIR="../../db/backup"
DATE=$(date +%Y%m%d%H%M%S)
PASSWD="123456"
USER="root"
IP="192.168.31.70"

# 1. 备份现有数据库（关键安全步骤）[1,6](@ref)
#echo "▶ 备份数据库 $DB_NAME..."
#touch $BACKUP_DIR/${DB_NAME}_${DATE}.sql
#mysqlDump -h $IP -u $USER -p $PASSWD --single-transaction --routines $DB_NAME > $BACKUP_DIR/${DB_NAME}_${DATE}.sql

# 3. 重建数据库（幂等操作）[5,8](@ref)
echo "⚙️ 重建数据库 $DB_NAME..."
mysql -h $IP -u $USER -p$PASSWD <<EOF
DROP DATABASE IF EXISTS $DB_NAME;
CREATE DATABASE $DB_NAME 
  CHARACTER SET utf8mb4 
  COLLATE utf8mb4_general_ci;

# 4. 导入表结构及数据 [3](@ref)
USE $DB_NAME;
SOURCE $SQL_FILE;
EOF

echo "✅ 重建完成！备份位置：$BACKUP_DIR/${DB_NAME}_${DATE}.sql"