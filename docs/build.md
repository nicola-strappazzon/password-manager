# Development

## Test completion for bash

```bash
go build -o pm .
alias pm="./pm"
pm completion bash > /tmp/pm_completion
source /tmp/pm_completion
```
