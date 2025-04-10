# zerodb
本视频是一个用 Go 语言构建分布式文件存储的教程。从创建主文件和 makefile 开始，逐步构建 peer-to-peer 的 TCP 库，涵盖消息编解码、RPC 实现、连接处理、存储功能构建等。包括生成文件路径、存储/读取/删除文件、添加存储根目录、创建服务器、实现对等节点功能、广播文件等操作，还处理了连接、消息、流传输、读取阻塞等问题，并不断优化，实现加密功能、测试优化以及解决服务器启动问题、防止代码重复等，最后还有添加删除文件功能的评估作业。
[00:00] 开始构建
	[00:00] 介绍用 Go 构建分布式文件存储项目
	[04:33] 创建 peer-to-peer 文件夹及文件
	[08:12] 定义 TCP transport 类型
	[11:37] 创建新 TCP transport
	[16:53] 用 TR 演示 accept 或 start 函数
[20:03] 处理连接
	[21:10] 创建接受循环，处理连接并调用 handleCon
	[24:56] 在传输接口中添加 listen 和 accept
	[27:45] 拨号连接为 outbound P，接受连接为 inbound P
	[34:49] 监听和接受连接成功，有新的对等节点
	[36:16] 创建握手函数，用于处理连接前的握手
[40:05] 设计握手函数
	[40:43] 设计握手函数并讨论其实现方式
	[43:48] 兴奋地讨论程序相关内容
	[47:37] 继续讨论解码器及相关操作
	[51:01] 讨论与连接握手相关问题
	[55:48] 强调先握手再解码的重要性
[01:00:07] 消息编解码与测试
	[01:00:10] 视频介绍用 Go 构建分布式文件存储教程
	[01:03:22] 介绍消息编解码与测试内容
[01:20:08] RPC 实现与测试
	[01:20:57] 实现分布式存储中 RPC 消息解码及测试
	[01:25:31] 定义 peer 接口首个方法 close
	[01:32:12] 服务器负责维护 peers 列表而非 transport
	[01:36:00] 在 main 中测试 onp 函数效果
	[01:39:24] 测试 peer close 功能效果
[01:40:15] 存储功能初步构建
	[01:45:09] 实现网络连接关闭时丢弃对等节点
	[01:45:58] 创建新文件构建存储功能
	[01:49:26] 创建存储测试文件
	[01:52:46] 确定文件名方式用于存储功能
	[01:58:13] 进行文件系统根目录配置
[02:00:19] 生成文件路径及存储
	[02:04:11] 生成文件路径，实现路径拼接并打印
	[02:07:43] 指定填充函数，扩展自定义转换函数
	[02:10:07] 文件名哈希后转换为十六进制字符串
	[02:13:35] 返回填充键，包含填充名和原始键
	[02:18:19] 获取文件名，测试存储并准备读取
[02:20:21] 读取文件及问题处理
	[02:20:21] 讨论文件名设置及读取文件的问题
	[02:24:10] 讨论读取文件的多种实现方式
	[02:28:08] 强调读取文件时返回 reader 的重要性
	[02:32:10] 考虑存储相同文件问题及创建删除功能
	[02:35:32] 测试删除功能及考虑更多函数需求
[02:40:24] 添加存储根目录及完善删除功能
	[02:48:31] 删除功能测试成功，开始考虑添加存储根目录
	[02:51:31] 添加存储根目录，考虑系统设计和实现方式
	[02:56:01] 存储根目录设置成功，考虑优化使其更简洁
	[02:59:17] 功能正常运行，准备制作删除函数并完善测试
	[03:00:19] 完善读取功能，为制作删除函数做准备
[03:00:25] 文件存储测试与清理函数
	[03:00:25] 视频介绍用 Go 构建分布式文件存储教程
	[03:02:46] 文件存储测试，探讨删除功能是否可行
	[03:08:00] 为测试构建辅助函数及清理函数
	[03:12:16] 运行测试，确保无文件夹残留
	[03:15:44] 测试新存储，发现部分测试失败
[03:20:26] 创建文件服务器
	[03:22:52] 开始创建文件服务器
	[03:31:32] 配置文件服务器选项
	[03:32:37] 创建对等 TCP 传输
	[03:35:16] 检查传输错误
	[03:36:17] 在传输中添加日志
[03:40:29] 文件服务器的循环与连接功能
	[03:40:29] Go 语言构建分布式文件存储，介绍文件服务器循环与连接
	[03:45:17] 讨论文件服务器的循环与连接中的错误处理
	[03:49:13] 考虑用户退出动作对文件的影响及清理操作
	[03:53:07] 检查错误并处理文件服务停止后的清理工作
	[03:56:40] 文件服务器的引导网络及连接操作
[04:00:30] 初步构建与连接问题
	[04:00:30] 视频介绍构建分布式文件存储，探讨问题
	[04:03:47] 继续讨论构建中的问题及尝试连接
	[04:09:27] 困惑于代码问题，考虑是否重录视频
	[04:13:16] 连接被拒，创建服务器函数测试
	[04:16:44] 在函数中处理问题，创建服务器二
[04:20:31] 对等节点功能实现
	[04:20:31] 讨论对等节点功能实现，包括添加和删除对等节点
	[04:24:36] 继续讨论对等节点功能兼容性
	[04:28:44] 阐述远程地址实现及对等节点接口功能需求
	[04:33:29] 讨论 TCP 发送功能及对等节点后续操作
	[04:37:02] 实现文件存储及向网络广播文件功能
[04:40:32] 广播文件及问题解决
	[04:40:32] 用 Go 语言构建分布式文件存储之广播文件及问题解决
	[04:44:34] 解决连接私有变量问题，实现广播文件功能
	[04:48:37] 解释多写器原理及实现广播文件的重要性
	[04:52:18] 处理广播文件时的读取阻塞及数据为空问题
	[04:55:48] 解码接收到的消息并发现问题
[05:00:33] 初步构建与问题排查
	[05:00:33] 创建文件服务器及处理消息的思路探讨
	[05:04:43] 构建 Constructor 及处理传输问题
	[05:09:19] 处理错误及探讨文件结束问题
	[05:13:51] 大文件处理问题及流式传输思路
	[05:17:11] 发送消息给所有对等节点及测试
[05:20:34] 解决读取阻塞问题
	[05:20:34] 教程讲解发送大文件及遇到的问题
	[05:24:10] 继续处理问题，鼠标挡路小插曲
	[05:27:57] 检查消息，测试解决方案是否可行
	[05:32:21] 打印消息，接收大文件并处理
	[05:35:45] 定义消息类型，处理存储文件问题
[05:40:36] 完善文件存储功能
	[05:44:43] 代码调整后文件存储功能正常
	[05:45:46] 准备将文件存储到文件并清理
	[05:50:26] 清理代码准备打开连接存储文件
	[05:56:11] 解释文件存储中出现阻塞原因
	[06:00:08] 发送文件大小信息以便正确存储
[06:00:41] 初步构建与问题发现
	[06:00:41] Go 语言构建分布式文件存储教程开篇
	[06:04:25] 测试中出现意外文件，查找问题
	[06:11:17] 调整 T reader，发现错误并改正
	[06:14:06] 完成存储分布，但代码仍需完善
	[06:18:58] 完善存储数据功能，广播消息
[06:20:42] 存储与读取功能完善
	[06:20:42] 用 Go 构建分布式文件存储教程
	[06:21:38] 重命名并定义存储和获取函数
	[06:24:42] 本地无文件则在网络中搜索
	[06:29:02] 确保新消息可在网络中发送
	[06:36:54] 处理解码错误并继续讲解
[06:40:48] 处理接收文件问题与优化
	[06:40:48] 用 Go 语言构建分布式文件存储教程，涵盖多个操作
	[06:40:48] 处理接收文件问题，发现多处问题待解决
	[06:40:48] 分析接收文件问题产生原因及影响
	[06:40:48] 提出解决接收文件问题的思路
	[06:40:48] 在库中添加对文件流的支持方案
[07:00:50] 处理流与广播消息
	[07:00:50] 处理流与广播消息，等待流并发送消息
	[07:04:17] 确定流发送位置，广播消息并接收存储
	[07:08:00] 测试并调整睡眠时间，使程序更优
	[07:12:40] 处理广播消息及流，检查存储并解锁
	[07:17:00] 发现问题，调整睡眠时间使程序正常
[07:20:53] 优化与扩展功能
	[07:23:03] 检查设置，考虑如何优化传输功能
	[07:25:30] 下集将处理获取文件及相关问题
	[07:28:26] 从本地或网络获取文件并广播
	[07:32:35] 循环处理消息，获取文件
	[07:36:27] 返回存储结果，考虑读取问题
[07:40:55] 发送文件大小并优化读取流
	[07:41:26] 发送文件大小，先硬编码再逐步优化
	[07:44:36] 删除硬编码内容并继续讨论优化方案
	[07:49:01] 提出新方案在 read stream 中获取文件大小
	[07:52:59] 讨论读取流方案及后续优化待办事项
	[07:56:08] 优化读取流，考虑返回读取流而非复制
[08:00:58] 加密功能实现
	[08:03:09] 考虑实现加密，避免文件在网络存储时被窥探
	[08:04:37] 加密不存入内存，支持流式传输，每次读取部分字节
	[08:07:36] 填充随机字节创建 IV，用于加密文件
	[08:11:55] 将 IV 前置到文件，用于解密
	[08:17:50] 下一步实现文件解密功能
[08:20:59] 加密文件的解密与测试
	[08:24:23] 加密文件后可解密并测试成功
	[08:28:14] 测试加密文件，IV 为 16 字节
	[08:31:29] 虽方式不佳但加密功能正常
	[08:35:50] 进行加密文件存储测试
	[08:38:41] 应正确输出加密后结果
[08:41:00] 加密文件的检索与优化
	[08:41:00] 视频介绍构建 Go 语言分布式文件存储教程
	[08:42:41] 寻找方法检索加密文件并优化存储
	[08:49:06] 解决加密文件每次启动生成新密钥问题
	[08:54:55] 修复待办事项，优化加密文件处理
	[09:00:11] 添加新服务器并讨论自动连接问题
[09:01:01] 解决服务器启动问题
	[09:01:01] Go 语言构建分布式文件存储，解决服务器启动问题
	[09:05:22] S3 启动问题，部分连接被拒
	[09:09:47] 服务器仍有问题，不断尝试
	[09:14:48] 服务器全部启动，连接成功
	[09:17:45] 对加密部分代码进行重构
[09:21:03] 防止代码重复与加密文件名
	[09:21:49] 防止代码重复，提供一处调整代码的地方
	[09:22:53] 加密文件名，防止文件名被轻易看到
	[09:25:22] 解决不知文件名无法同步的问题
	[09:28:26] 生成 ID 增强文件服务器安全性
	[09:38:26] 发现存储文件位置错误需重构
[09:41:06] 添加唯一标识符与评估作业
	[09:41:06] 考虑重构代码，引入唯一标识符
	[09:47:20] 存储出现问题，需添加唯一标识符
	[09:51:13] 处理存储和消息中的唯一标识符
	[09:54:32] 服务器可按唯一标识符存储文件
	[09:55:57] 给出添加删除功能的评估作业