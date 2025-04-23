# Build

## Test completion for bash

```bash
go build -o pm .
alias pm="./pm"
pm completion bash > /tmp/pm_completion
source /tmp/pm_completion
```


source <(pm completion bash)


complete -F __start_pm pm
type __start_pm



complete -r pm
unset -f __start_pm

bash --noprofile --norc
source ~/.bash_profile
