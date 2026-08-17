# Bug 复现说明

## Bug 是什么

`Reverse` 原地翻转并返回原切片，修改了调用方入参。

## 如何触发

```bash
go test ./internal/rev/ -run TestReverseCopy -count=1 -v
```

## 错误信息（修前）

```text
input mutated: ...
```

修后应 PASS；`go test ./...` 全绿。
