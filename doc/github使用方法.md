# github使用方法

> tips:命令行按tap键可以快速补全，按两次可以显示当前路径下与名字匹配的所有文件/文件夹
>
> 以下操作在vscode也有对应的图形化界面，如有问题及时问我

## 常用操作

### 提交代码

```sh
先更新本地代码
git checkout master
git pull		    //拉取
git checkout dev
git merge master	//合并
提交代码
git add (文件1 文件2) || git add .(添加所有文件)
git commit -m '提交信息'
git push
```

### 提交信息规范

```sh
提交信息格式如下
增:'add:增加的内容'
删:'del:删除的内容'
改:'mod:修改的内容'
修复:'fix:修复的内容'
```

### 自己的远程分支落后主分支

```sh
git checkout master
git pull		    //拉取
git checkout dev
git push			//提交
```



## 使用前准备工作

### 获取仓库代码

```SH
git clone + 仓库链接
如：
git clone git@github.com:aloser6/ISP.git
```

> tips:git clone后默认有maste/main分支

### 创建本地分支并绑定远端分支

```sh
git branch + 分支名      									//分支名为英文本人姓名
git checkout + 分支名	   									//切换分支，这里要切换到刚创建的本地分支
git push origin + 分支名   								//在远端创建分支，这里最好和刚刚创建的本地分支同名
git branch --set-upstream-to=origin/远端分支名 本地分支名       //本地分支与远端分支做关联
```

### 开发流程

```sh
按照以上步骤此时你本地有两个分支，一个是master/main，另一个是刚刚创建的分支，这里以dev代指刚刚所创建的分支名
首先切换到dev分支，在该分支上进行开发，开发完成后切换回到master/main分支，在该分支上拉去远端的master/main分支，此时master/main为最新的，随后切换到dev分支，将dev和master进行合并并解决出现的分支冲突(不一定会出现)，然后在dev分支上提交到远端dev分支，可以结合下图理解
```

### 开发流程指令实现

```sh
将设本地创建的为dev分支，主分支为master
git checkout dev
本地开发完成后
git checkout master
git pull		    //拉取
git checkout dev
git merge master	//合并
git push		   //提交到远端
```



<img src="C:\Users\高扬\Desktop\图片\Snipaste_2023-01-23_16-40-26.png" alt="Snipaste_2023-01-23_16-40-26" style="zoom: 33%;" />





### 本地git未与github进行连接

```SH
输入指令
ssh-keygen -t rsa
然后敲三次回车键(可以设置密码，这里就以无密码版举例)，在~/.ssh文件内有两个文件，分别是id_rsa和公钥id_rsa.pub，
将公钥id_rsa.pub里的内容复制到github设置中的密钥(SSH)一项即可将本地git和远端github进行绑定
可用指令ssh -T git@github.com进行验证是否连接成功

操作简述
ssh-keygen -t rsa
cd ~/.ssh
vim id_rsa.pub
复制所打开文件里的内容
```





