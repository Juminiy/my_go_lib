Go语言构建算法库

由于Go语言的标准库实现的数据结构少，所以需要自己实现

本项目目标在于从零开始实现编译原理的算法

Go实现的数据结构

栈，队列，链表，二叉树，红黑树，高度平衡树，跳表，图论

先封装底层的基础结构(如不可重复Set,Map,图论算法)，然后向上层提供API支持,作为编译前端的实现前提

2021.09.09start

2021.09.11开始底层数据结构构建（图论算法很🤢debug快吐了）

2021,09,12完成web api

2021.09.13修正图论算法，完成EpsilonClosure构造过程

为了向上层提供各种能力，需要封装底层图论，不可重复Set等API

2021.09.14正则表达式->NFA->DFA->DFA最小化过程

2021.09.15Set运算增强

本proj可部署到物理机或者容器内，作为一个FaaS(云函数)服务

支持三种方式调用

- HTTP RESTful API 格式json
- gRPC 跨语言调用 protobuf协议 尚在开发中
- CLIs 尚在开发中

HTTP REST 形式，GET方法可以用浏览器，POST方法只能用命令行或者Postman等接口测试工具

![image-20210914091100072](https://v.hualingnan.site/typora/image-20210914091100072.png)

POST方法设置HTTP-Headers Content-Type 为 application/json 请求体Body内容为json格式

![image-20210914091140039](https://v.hualingnan.site/typora/image-20210914091140039.png)

![image-20210914091204177](https://v.hualingnan.site/typora/image-20210914091204177.png)

![image-20210914091036833](https://v.hualingnan.site/typora/image-20210914091036833.png)

返回结果

![image-20210914091318761](https://v.hualingnan.site/typora/image-20210914091318761.png)


```json
{
"Edges" :[
{
"NodeIValue": 0,"NodeJValue": 1,"EdgeValue":"epsilon"
},{
"NodeIValue": 0,"NodeJValue": 7,"EdgeValue":"epsilon"
},{
"NodeIValue": 1,"NodeJValue": 2,"EdgeValue":"epsilon"
},{
"NodeIValue": 1,"NodeJValue": 4,"EdgeValue":"epsilon"
},{
"NodeIValue": 2,"NodeJValue": 3,"EdgeValue":"a"
},{
"NodeIValue": 4,"NodeJValue": 5,"EdgeValue":"b"
},{
"NodeIValue": 6,"NodeJValue": 1,"EdgeValue":"epsilon"
},{
"NodeIValue": 6,"NodeJValue": 7,"EdgeValue":"epsilon"
},{
"NodeIValue": 3,"NodeJValue": 6,"EdgeValue":"epsilon"
},{
"NodeIValue": 5,"NodeJValue": 6,"EdgeValue":"epsilon"
},{
"NodeIValue": 7,"NodeJValue": 8,"EdgeValue":"a"
},{
"NodeIValue": 8,"NodeJValue": 9,"EdgeValue":"b"
},{
"NodeIValue": 9,"NodeJValue": 10,"EdgeValue":"b"
}
],
"Nodes":[ 0 ]
}
```
其他REST形式类似
gRPC

跨语言调用
待开发

CLIs

命令行工具
待开发
![img.png](https://v.hualingnan.site/typora/img.png) 快2000line 了 
2021.09.14 13:35 

2021.09.15 
- 不能完全保证算法的正确，需要大量数据去验证
- 在debug过程中内存泄漏，引用非法，dereference指针异常
从这个状态
![](https://v.hualingnan.site/typora/nfa1_1.jpeg) 
转换到
![](https://v.hualingnan.site/typora/dfa1_1.jpeg)