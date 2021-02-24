# choco

## 上传到 coco 教程
1. 把新编译文件,改名为 ferebook_0.0.1_win64_x64.zip 格式
2. 上传到 release 上
3. 修改 ferebookd.nuspec 的版本信息
4. 修改 /tools/chocolateyInstall.ps1 的版本信息
5. 打包 `choco pack`
6. 测试 `choco install ferebookd -s .`
7. 提交 `choco  push  -api-key your_api_key`
7. 审核
