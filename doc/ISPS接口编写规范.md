# ISPS接口编写规范

### 命名规则

```markdown
1、集合命名应该使用复数形式
2、采用驼峰式命名法
3、做到见名知意
```

### 完全的面向对象

> tips:英文小写的情况下，按住shift再按字母就是大写的

```go
//对事物抽象成结构体，赋予其属性和方法，以下是一个简单的小例子
//结构体
type Logger struct{
    name          string         		//日志所属部分，例如[root]，[system]
    appender      loggerAppender 		//输出地 例如os.Stdout, filePath
    timeNow		 time				  //时间
    line		 int16				  //行数
    fileName	 string				  //所处文件位置
}
type loggerAppender struct{
    isStdout	bool			//终端
    isFile		bool			//文件
    filePath	string			//文件地址
}

var gLogger Logger
//登录接口,需要详细的接口描述，包括用途，用法和注意细节等
func Login(){
    inputFile, inputErr := os.Open("input.dat")
    if inputErr != nil{
        gLogger.Error("文件打开失败")
        os.exit(-1)
    }
}

//对于需要跨包的方法需要大写，否则使用小写即可
//结构体Logger的方法
func (ctx Logger) Error(str const string) string {
    var conf Config 				//定义配置器
    path := conf.get("file.path")    //通过配置器获取文件默认输出路径
    if ctx.appender.isFile == true{
		return str.Out(str, path)       
    }else if ctx.appender.isStdout == true{
        return str.Out(str, os.Stdout)
    }
    os.exit(-1)
}


```

### 开发小细节

```markdown
尽量使用const
对于临时变量，尽量使用短声明，即 :=
go语言函数方法是多返回值，所以在编写方法的时候可以在需要的情况下返回err
设计接口前最好先画出接口流程图，画完图理清关系再编写代码
```

### 开发参考资料

> [《Go 入门指南》 | Go 技术论坛 (learnku.com)](https://learnku.com/docs/the-way-to-go)

### 持续更新中......
