# douyin-server

## 安装依赖
```shell
go mod tidy
```

## 配置数据库
复制`config`文件夹中的`configs_example.go`文件，命名为`configs.go`，解除注释，填写相关信息，运行`model`文件夹下的`mysql_init_test.go`的`TestMySQLInit`函数测试无误即可。
切勿把任何敏感数据提交到git仓库中！

## git commit 规范
前缀规范：
- [feat] 新增功能
- [fix] 修复bug
- [refactor] 重构代码(同一功能不同逻辑)
- [style] 修改代码格式，包含拼写等错误的修正
- [perf] 性能优化
- [test] 新增测试用例
- [docs] 文档说明

请勿提交没有意义的代码到git仓库中！