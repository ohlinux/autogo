autogo
======

Go语言是静态语言，修改源代码总是需要编译、运行，如果用Go做Web开发，修改一点就要编译、运行，然后才能看结果，很痛苦。
autogo就是为了让Go开发更方便。在开发阶段，做到修改之后，立马看到修改的效果，如果编译出错，能够及时显示错误信息！

使用说明
======

1、下载
将源代码git clone到任意一个位置

2、修改config/project.json文件
  该文件的作用是配置需要被autogo管理的项目，一个项目是一个json对象（{}），包括name、root和depends，
  其中depends是可选的，name是项目最终生成可执行文件的文件名（也就是main包所在的目录）；root是项目的根目录。

3、执行install.sh(linux)/install.bat(windows)，编译autogo

5、在autogo的config/project.json中将该项目加进去
  
 [
    {
        "name": "test",
        "root": "../test",
        "go_way": "install",
        "deamon": true,
        "main": "test1/test.go",
        "depends": []
    }
]
    root可以是相对路径或决定路径.


4、运行autogo：autogo 直接运行不需要切换目录

5、因为基本流程的修改，被监控程序要可以接收 syscall.SIGINT, syscall.SIGUSR1 至少这两个信号具体的sample可以先查看https://github.com/ohlinux/to-do-list (还不够完善)

版本更新历史
=====

2012-12-26 autogo 2.1 发布
( Ajian modify )
    ```
    1、去掉原来的error输出到html的方式
    2、重新定义工作方式:
    autogo负责send signal到被监控的服务，被监控服务有一个接受signal处理的机制（关闭连接和重新启动自己）.
       基本流程: autogo --> Watch到文件更新-->进行compile，编译是否出问题，并且报告到终端显示-->如果没有问题进行send restart signal,让被监控程序进行重启
    ;如果编译有问题，退出并报告不进行重启。
       这里存在一个需要解决的问题，重启后的服务是一个fork出来的程序，被监控的程序本身的日志是不会出现在标准输入的，因为默认go已经关闭了。
    这种方式解决了在autogo进行killall时造成的defunct程序。
      接下来要完善的:
      1，被监控程序接收信息处理的lib 
      2、统一的日志lib，将被监控程序的日志输出到文件
    3、简化了程序，只对linux进行了丰富，没有测试windows。
    4、增加了一个logger package,对于调试和输出都相当的方便。
    5、解决了autogo对于./bin/autogo启动方式的依赖，并且优化了shell脚本，不需要export oldpath 因为在子程序的export不会影响system。
    ```

2012-12-20  autogo 2.0发布
```
1、优化编译、运行过程（只会执行一次）
2、支持多种goway方式：go run、build、install，这样对于测试项目也支持了
3、修复了 被监控项目如果有问题 autogo启动不了的情况
4、调整了代码结构
```

2012-12-18  autogo 1.0发布

使用的第三方库
======

为了方便，autogo中直接包含了第三方库，不需要另外下载。

1、[fsnotify](https://github.com/howeyc/fsnotify)，File system notifications

2、[simplejson](https://github.com/bitly/go-simplejson)，解析JSON，我做了一些改动

感谢
=====

johntech

[ohlinux](https://github.com/ohlinux)

LICENCE
======

The MIT [License](https://github.com/polaris1119/autogo/master/LICENSE)
