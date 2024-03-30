CREATE TABLE `relation` (
  `uid` int PRIMARY KEY,
  `pid` int UNIQUE,
  `fid` int UNIQUE,
  `cid` bigint UNIQUE,
  `followee` varchar(65535),
  `following` varchar(65535)
);

CREATE TABLE `users` (
  `uid` integer PRIMARY KEY,
  `account` int NOT NULL,
  `password` varchar(64) NOT NULL,
  `username` varchar(255) NOT NULL,
  `starttime` timestamp NOT NULL,
  `sex` enum(男,女,保密) NOT NULL,
  `grade` varchar(3) NOT NULL,
  `college` varchar(128),
  `major` varchar(128)
);

