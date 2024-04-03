#!/bin/bash

# 设置 MySQL 登录凭据
MYSQL_USER="root"
MYSQL_PASSWORD="123456"

# 创建数据库
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD -e "CREATE DATABASE IF NOT EXISTS ahutforum;"

# 切换到 ahutforum 数据库
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD -e "USE ahutforum;"

# 创建 users 表
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD ahutforum << EOF
CREATE TABLE IF NOT EXISTS \`users\` (
  \`uid\` int PRIMARY KEY AUTO_INCREMENT,
  \`account\` int NOT NULL,
  \`password\` varchar(64) NOT NULL,
  \`username\` varchar(255) NOT NULL,
  \`starttime\` timestamp NOT NULL,
  \`sex\` enum('男','女','保密') NOT NULL,
  \`grade\` varchar(3) NOT NULL,
  \`college\` varchar(128),
  \`major\` varchar(128)
);
EOF

# 创建 forums 表
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD ahutforum << EOF
CREATE TABLE IF NOT EXISTS \`forums\` (
  \`fid\` int PRIMARY KEY AUTO_INCREMENT,
  \`uid\` int,
  \`forumname\` varchar(255) NOT NULL,
  \`membercount\` int DEFAULT 1,
  \`description\` text,
  FOREIGN KEY (\`uid\`) REFERENCES \`users\` (\`uid\`)
);
EOF

# 创建 posts 表
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD ahutforum << EOF
CREATE TABLE IF NOT EXISTS \`posts\` (
  \`pid\` bigint PRIMARY KEY AUTO_INCREMENT,
  \`fid\` int,
  \`uid\` int,
  \`title\` varchar(128) NOT NULL,
  \`text\` text NOT NULL,
  \`browse\` int DEFAULT 0,
  \`comment\` int DEFAULT 0,
  \`votes\` int DEFAULT 0,
  \`releasetime\` timestamp NOT NULL,
  FOREIGN KEY (\`fid\`) REFERENCES \`forums\` (\`fid\`),
  FOREIGN KEY (\`uid\`) REFERENCES \`users\` (\`uid\`)
);
EOF

# 创建 comments 表
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD ahutforum << EOF
CREATE TABLE IF NOT EXISTS \`comments\` (
  \`cid\` bigint PRIMARY KEY AUTO_INCREMENT,
  \`uid\` int,
  \`pid\` bigint,
  \`fathercid\` bigint DEFAULT -1,
  \`votes\` int DEFAULT 0,
  \`text\` text NOT NULL,
  \`releasetime\` timestamp NOT NULL,
  FOREIGN KEY (\`uid\`) REFERENCES \`users\` (\`uid\`),
  FOREIGN KEY (\`pid\`) REFERENCES \`posts\` (\`pid\`)
);
EOF

# 创建 follow 表
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD ahutforum << EOF
CREATE TABLE IF NOT EXISTS \`follow\` (
  \`followid\` int PRIMARY KEY,
  \`uid\` int,
  \`following\` int NOT NULL,
  FOREIGN KEY (\`uid\`) REFERENCES \`users\` (\`uid\`)
);
EOF

echo "数据库 ahutforum 和表已成功创建。"
