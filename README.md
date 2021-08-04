# algorithm-exercise
## algo
数据结构和算法
### 05_array: 数组
### 06_linkedlist: 链表
- 单向链表的实现
- 链表的场景算法
### 07_linkedlist: 链表常见问题
（代码写在06一起了)
### 08 :栈
- 以链表为基础实现的栈 
- 用链表模拟浏览器前进、后退功能（跳过）
### 09 :队列
- 以链表实现队列
- 以数组实现队列
- 环形队列（未完成，有bug）
### 10 sorts
冒泡排序、插入排序、选择排序
  
## algo-class
算法课练习

### class01
- Code06: 一个无序数组，任意相邻两个元素不相等，找到一个局部最小的值 (需要多练几次)

### class02 异或
- Code01: 不使用第三个变量交换两个数
- Code02: 异或相关题目
- Code03: 一个数组，一种数出现了k次，其他都出现了m次，m>1，m>k,找到k


## go文件运行方式
因为没有用标准的go目录结构，所以如果用vscode就没法正常的使用test运行文件。
goland是可以单独运行文件的。

### goland运行的方式
每个要运行的文件的 pakckge 都是 main。
工具类放入src/utils目录。

1. 配置gopath到根目录
```
go env -w GOPATH=/Users/a/Workspace/my-project/algorithm-exercise
```
如果之前已经配的有了加：分割
```
go env -w GOPATH=/Users/a/go:/Users/a/Workspace/my-project/algorithm-exercise
```

2. cd到运行的文件的目录。执行命令 go run target.go

目前这么改在vscode是可以运行的，
把文件改成用test也能运行。
goland里面还没有试，如果可以，就改成test。




