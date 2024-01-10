### 获取代码覆盖率
go test -v  -coverpkg=./... ./... -coverprofile=cover.out.tmp

egrep -v "xxx/pkg/test|mock_.*go|xxx/cmd" cover.out.tmp > cover.out

go tool cover -html=cover.out -o coverage.html

go tool cover -func=cover.out

### 代码组织
https://dev.to/codypotter/layered-architectures-in-go-3cg8
https://segmentfault.com/q/1010000003849810
https://juejin.cn/post/7298160530292703244

### UT
https://www.liwenzhou.com/posts/Go/unit-test-5/
https://juejin.cn/post/7133520098123317256
https://zhangyuyu.github.io/golang-unit-test/#3-monkey