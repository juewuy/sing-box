# DNS

### 结构

```json
{
  "dns": {
    "servers": [],
    "rules": [],
    "final": [
      "local"
    ],
    "strategy": "",
    "disable_cache": false,
    "disable_expire": false,
    "independent_cache": false,
    "reverse_mapping": false,
    "fakeip": {}
  }
}

```

### 字段

| 键        | 格式                      |
|----------|-------------------------|
| `server` | 一组 [DNS 服务器](./server/) |
| `rules`  | 一组 [DNS 规则](./rule/)    |

#### final

!!! note ""

    当内容只有一项时，可以忽略 JSON 数组 [] 标签

默认 DNS 服务器的标签列表。

当数量大于一时并发请求所有目标 DNS 服务器，取最快非空响应。

默认使用第一个服务器。

#### strategy

默认解析域名策略。

可选值: `prefer_ipv4` `prefer_ipv6` `ipv4_only` `ipv6_only`。

如果设置了 `server.strategy`，则不生效。

#### disable_cache

禁用 DNS 缓存。

#### disable_expire

禁用 DNS 缓存过期。

#### independent_cache

使每个 DNS 服务器的缓存独立，以满足特殊目的。如果启用，将轻微降低性能。

#### reverse_mapping

在响应 DNS 查询后存储 IP 地址的反向映射以为路由目的提供域名。

由于此过程依赖于应用程序在发出请求之前解析域名的行为，因此在 macOS 等 DNS 由系统代理和缓存的环境中可能会出现问题。

#### fakeip

[FakeIP](./fakeip/) 设置。
