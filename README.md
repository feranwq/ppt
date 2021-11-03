# golang 注释自动翻译器

### 功能说明
阅读 golang 源码时有大段的注释使用 `//` 分隔, 直接复制到翻译网站里语句不太通顺

本程序从剪切板里自动读取文本内容, 并处理掉 `//` 字符后通过百度翻译 API 自动返回翻译结果

### 支持平台

- osx

- windows

- linux/unix/wsl (依赖`xclip` 或者 `xsel`)

### 使用方法

```bash
make build
export BD_APP_ID=xxx
export BD_SECRET_KEY=xxx

./_output/ppt
```