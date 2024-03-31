CREATE TABLE `users` (
  `uid` int PRIMARY KEY AUTO_INCREMENT,
  `account` int NOT NULL,
  `password` varchar(64) NOT NULL,
  `username` varchar(255) NOT NULL,
  `starttime` timestamp NOT NULL,
  `sex` enum(男,女,保密) NOT NULL,
  `grade` varchar(3) NOT NULL,
  `college` varchar(128),
  `major` varchar(128)
);

CREATE TABLE `forums` (
  `fid` int PRIMARY KEY AUTO_INCREMENT,
  `uid` int,
  `forumname` varchar(255) NOT NULL,
  `membercount` int DEFAULT 1,
  `description` text
);

CREATE TABLE `posts` (
  `pid` bigint PRIMARY KEY AUTO_INCREMENT,
  `fid` int,
  `uid` int,
  `title` varchar(128) NOT NULL,
  `text` text NOT NULL,
  `browse` int DEFAULT 0,
  `comment` int DEFAULT 0,
  `votes` int DEFAULT 0,
  `releasetime` timestamp NOT NULL
);

CREATE TABLE `comments` (
  `cid` bigint PRIMARY KEY AUTO_INCREMENT,
  `uid` int,
  `pid` int,
  `fathercid` bigint DEFAULT -1,
  `votes` int DEFAULT 0,
  `text` text NOT NULL,
  `releasetime` timestamp NOT NULL
);

CREATE TABLE `follow` (
  `followid` int PRIMARY KEY,
  `uid` int,
  `following` int NOT NULL
);

ALTER TABLE `forums` ADD FOREIGN KEY (`uid`) REFERENCES `users` (`uid`);

ALTER TABLE `posts` ADD FOREIGN KEY (`fid`) REFERENCES `forums` (`fid`);

ALTER TABLE `posts` ADD FOREIGN KEY (`uid`) REFERENCES `users` (`uid`);

ALTER TABLE `comments` ADD FOREIGN KEY (`uid`) REFERENCES `users` (`uid`);

ALTER TABLE `comments` ADD FOREIGN KEY (`pid`) REFERENCES `posts` (`pid`);

ALTER TABLE `follow` ADD FOREIGN KEY (`uid`) REFERENCES `users` (`uid`);
