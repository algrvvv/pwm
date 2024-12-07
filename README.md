# PWM by algrvvv

PWM - simple terminal password and something more manager

## Preparation

```shell
gpg --full-generate-key
```

## Install

```shell
go install github.com/algrvvv/pwm@latest
```

```shell
pwm init
```

## Auto completion

```shell
pwm completion zsh > ~/tmp/pwm
# It's better to put this command in your shell configuration, for example ~/.zshrc
source ~/tmp/pwm
```

## See usage

To see available commands use:
```shell
pwm help
```

## Examples

Store new password or note:

```shell
pwm store remote_server_ip "127.0.0.1"
# or
pwm store github-token "my_github_token"
```

Get password or note:
```shell
pwm get remote_server_ip
# or (if you need see and copy)
pwm get github-token --clip
```

Only copy:
```shell
pwm copy github-token
```

Remove password or note:
```shell
pwm rm remote_server_ip
```

Get all passwords or notes:
```shell
pwm list
```

Generate new password:
```shell
pwm generate --len 32 --clip --save password_name
```

All flags for generate password:
```text
Flags:
  -c, --clip               save to clipboard
  -h, --help               help for generate
  -l, --len int            password len. (default 12)
  -s, --save string        save note by name
  -D, --without-digits     dont use digits
  -S, --without-specials   dont use special symbols
  -U, --without-uppers     dont use upper case symbols
```









