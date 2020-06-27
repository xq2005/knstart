# knstart

基于K8S的Operator开发，95%甚至99%的用户都会选择使用[Operator-SDK](https://sdk.operatorframework.io/)自动生成K8S Client的框架代码，因为Operator-SDK拥有完整的文档，活跃的社区，丰富的语言种类，再加上多多的参考实例，实现一个新的Operator成为一件简单的事情。因为实在是太流行了，随便在网上搜索一下就可以找到详细的步骤，这里也不多做介绍。

最近正在阅读[Knative Operator](https://github.com/knative/operator)源代码，发现Knative Operator并没有按照常理出牌，这个operator不是基于Operator-SDK。它是基于[Knative Common Package](https://github.com/knative/pkg)生成的框架代码。 这个Knative Common Package是在[code-generator](https://github.com/kubernetes/code-generator)的基础之上，创建了一种新的generator： injection。 injection对code-generator生成的clientset, informers和listers进行了第二次封装，提供给用户的interface变成了两个**ReconcileKind**（cr创建或者更新的时候这个interface会被调用）和 **FinalizeKind**（当cr被删除的这个interface会被调用），后面打算详细描述一下这两个接口的工作方式。 先介绍如何使用Knative Common Package完成一个operator的详细步骤。

## 如何创建一个基于Knative的工程

### 1.准备工作

+ Linux环境（Knative Common Package生成框架代码的脚本只提供了shell版本）
+ golang环境1.13或者以上，GOROOT和GOPATH都要设置
+ Kubernetes环境，1.17.6或以上 （MiniKube也是可以的）
+ 编辑器，我用的是vscode

### 2. 要先将自己要创建的CRD的名字和相关的定义确定

```yaml
API Group: xq2005.com
Version: v1
Resource Name: KNLearing
```

### 3. 准备用于自动生成框架的代码

下面是详细的步骤，可以先不用关注代码的内容
1). $GOPATH/src下创建目录knstart/pkg/apis/operator/v1

```bash
mkdir -p $GOPATH/src/knstart/pkg/apis/operator/v1
```

2). 创建$GOPATH/src/knstart/pkg/apis/operator/[register.go](pkg/apis/operator/register.go)

3). 创建$GOPATH/src/knstart/pkg/apis/operator/v1/[types.go](pkg/apis/operator/v1/types.go)

4). 创建$GOPATH/src/knstart/pkg/apis/operator/v1/[doc.go](pkg/apis/operator/v1/doc.go)

5). 创建$GOPATH/src/knstart/pkg/apis/operator/v1/[register.go](pkg/apis/operator/v1/register.go)

6). 创建$GOPATH/src/knstart/pkg/apis/operator/v1/[lifecycle.go](pkg/apis/operator/v1/lifecycle.go)

7). 创建$GOPATH/src/knstart/hack目录

```bash
mkdir -p $GOPATH/src/knstart/hack
```

8). 下载[sample-source/tree/master/hack](https://github.com/knative-sandbox/sample-source/tree/master/hack)的所有文件到$GOPATH/src/knstart/hack

9). 修改文件$GOPATH/src/knstart/hack/[update-codegen.sh](hack/update-codegen.sh), 34行,35行，42行和43行按照当前的项目情况进行修改。

10). 编辑$GOPATH/src/knstart/[go.mod](go.mod)，把最重要的k8s.io/code-generator和knative.dev/pkg这两个依赖库加入。

11). 进入当前工程的目录

```bash
cd $GOPATH/src/knstart/
```

12). 因为$GOPATH/src/knstart/hack/下的shell脚本会使用vendor目录下面的code-generator和Knative Common Package的binary生成代码。所以要把这两个project放到vendor目录下

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct  #国内的小伙伴们一定要加上这一行
go mod vendor  
```

13). 自动生成框架

```bash
chmod a+x hack/*.sh
hack/update-codegen.sh
```

14). 生成的代码都放在了$GOPATH/src/knstart/pkg/client下面，共有四个目录

+ clientset
+ informers
+ injection   ---- 这个是knative所特有的
+ listers

15). 书写controller的代码，因为基于knative的api与operator-sdk的有一些差别。参考本项目中的代码就可以(后面可以详细描述一下与operator-sdk的区别有哪些)

16). 调试，编译，创建image，部署的过程与operator与operator-sdk完全是相同的，在本项目中也没有添加这部分内容。
