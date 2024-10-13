<div align="center">
<h1>Go AliMail</h1>

[![Auth](https://img.shields.io/badge/Auth-eryajf-ff69b4)](https://github.com/eryajf)
[![GitHub contributors](https://img.shields.io/github/contributors/eryajf/go-alimail)](https://github.com/eryajf/go-alimail/graphs/contributors)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/eryajf/go-alimail)](https://github.com/eryajf/go-alimail/pulls)
[![GitHub Pull Requests](https://img.shields.io/github/stars/eryajf/go-alimail)](https://github.com/eryajf/go-alimail/stargazers)
[![HitCount](https://views.whatilearened.today/views/github/eryajf/go-alimail.svg)](https://github.com/eryajf/go-alimail)
[![GitHub license](https://img.shields.io/github/license/eryajf/go-alimail)](https://github.com/eryajf/go-alimail/blob/main/LICENSE)
[![](https://img.shields.io/badge/Awesome-MyStarList-c780fa?logo=Awesome-Lists)](https://github.com/eryajf/awesome-stars-eryajf#readme)

<p> 🧰 阿里企业邮箱 GO 语言 SDK 🧰 </p>

<img src="https://cdn.jsdelivr.net/gh/eryajf/tu@main/img/image_20240420_214408.gif" width="800"  height="3">
</div><br>

![go-alimail](https://socialify.git.ci/eryajf/go-alimail/image?description=1&descriptionEditable=%E9%80%90%E6%AD%A5%E8%BF%88%E5%90%91%E8%BF%90%E7%BB%B4%E7%9A%84%E5%9B%9B%E4%B8%AA%E7%8E%B0%E4%BB%A3%E5%8C%96%EF%BC%9A%E8%A7%84%E8%8C%83%E5%8C%96%EF%BC%8C%E6%A0%87%E5%87%86%E5%8C%96%EF%BC%8C%E9%AB%98%E6%95%88%E5%8C%96%EF%BC%8C%E4%BC%98%E9%9B%85%E5%8C%96&font=Bitter&forks=1&issues=1&language=1&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Light)

</div>

## 项目简介

[阿里企业邮箱](https://wanwang.aliyun.com/mail)GO语言SDK。

接口文档地址：[https://mailhelp.aliyun.com/openapi/index.html](https://mailhelp.aliyun.com/openapi/index.html)

**目前只写了部分接口，其余接口空了补充，也欢迎有需要的同仁PR。**

- [x] 组织
- [x] 域名
- [x] 用户
- [ ] 部门
- [ ] 邮件组
- [ ] 公共联系人
	- [ ] 联系人
	- [ ] 分组
- [ ] 邮件
	- [ ] 邮件
	- [ ] 邮件文件夹
- [ ] 日历
	- [ ] 日历
	- [ ] 日历文件夹
- [ ] 登录&登出
- [ ] 审计日志
- [ ] 文件流

## 快速开始

**安装依赖**

```
go get github.com/eryajf/go-alimail
```

**简单示例**

```go
package main

import (
	"context"
	"fmt"

	"github.com/eryajf/go-alimail/alimail"
)

func main() {
	client := alimail.NewClient("appID", "appSecret")
	on, err := client.Organization.Get(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Organization: %+v\n", on)
	domains, err := client.Domain.List(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Domains: %+v\n", domains)
}
```


## 其他说明

- 如果觉得项目不错，麻烦动动小手点个 ⭐️star⭐️!
- 如果你还有其他想法或者需求，欢迎在 issue 中交流！

## 捐赠打赏

如果你觉得这个项目对你有帮助，你可以请作者喝杯咖啡 ☕️

| 支付宝|微信|
|:--------: |:--------: |
|![](https://t.eryajf.net/imgs/2023/01/fc21022aadd292ca.png)| ![](https://t.eryajf.net/imgs/2023/01/834f12107ebc432a.png) |
