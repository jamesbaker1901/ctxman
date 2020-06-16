# Context Manager

NOTE: `ctxman` is under active development! Some features are merely placeholder code and others are largely untested. Proceed at your own risk!

`ctxman` makes managing your many environments and contexts a breeze. In the modern DevOps world, developers and operators often have to change contexts or environments many times a day, switching from prod to dev to test or internal tools environments. Each of the changes may involve: changing environment variables, swapping config files, setting a new `kubectl` cluster and namespace, a new vpn connection, etc.

Rather than write a bash script or handle each of these changes manually, simply define your environments in a `ctxman` config file and let it handle the tedium for you!

## Installation
If you have Go installed:

```
$ go get -u github.com/jamesbaker1901/ctxman
```

Otherwise, download the binary for your OS and architecture from the releases page.

NOTE: `ctxman` is for Linux and OSX. Some features may work for Windows, but `ctxman` is completely untested and not targeted for Windows. Pull-requests are welcome though!

Ensure `ctxman` is in your `$PATH` with `which ctxman`. If not, you may need to add `~/go/bin` to your `$PATH`.

If you intend to use `ctxman` to handle environment variables, you will also need to add the following to your `~/.bashrc` or `~/.zshrc`:

#### Bash
`~/.bashrc` or `~/bash_profile` for OSX
```
alias cc="ctxman && source ~/.config/ctxman/env"
source ~/.config/ctxman/env
PROMPT_COMMAND='source ~/.config/ctxman/env'
```
#### zsh
`~/.zshrc`
```
alias cc="ctxman && source ~/.config/ctxman/env"
source ~/.config/ctxman/env
PROMPT_COMMAND='source ~/.config/ctxman/env'
precmd() { eval "$PROMPT_COMMAND" }
```


`ctxman` can't set environment variables directly for your shell. It currently works by adding the the variables you define in the configuration file to a simple shell script at `~/.config/ctxman/env` which needs to be sourced for every shell. To avoid having to do so manually, we can make use of the `$PROMPT_COMMAND` variable, which defines commands to be run before the `PS1` is drawn to the screen, i.e., every time you press enter in the terminal. If this behavior is undesirable, simply remove `PROMPT_COMMAND` lines from your shell's startup script.

We also see in this example a suggested alias for `ctxman` of `cc` (change context). This is of course optional, but recommended as you will likely call upon `ctxman` many times throughout the day; may as well save a few keystrokes!

## Configuration

`ctxman` is configured via a yaml file located at `~/.config/ctxman/config.yaml`.

Sample config:
```
---
prod:
  kubernetes:
    cluster: prod
    namespace: app
  env:
    - variable:
        key: PROD_EXAMPLE
        value: thisIsAString
    - variable:
        key: PROD_ANOTHER_EXAMPLE
        value: thisIsAlsoAString
    - variable:
        key: AWS_DEFAULT_REGION
        value: us-west-2
    - variable:
        key: AWS_PROFILE
        value: prod
  openvpn: 
    configFile: /home/user/.ovpn/prod-us.ovpn
  shell: 
    command: /bin/sh /home/user/someScript.sh
  path:
    include: /home/user/go/bin
    exclude: /home/user/experimental/bin
  symlink:
    source: /home/user/git/dotfiles/prod/someConfigFile.json
    dest: /home/user/.someConfigFile.json

test:
  kubernetes:
    cluster: test
    namespace: app
  env:
    - variable:
        key: TEST_EXAMPLE
        value: thisIsAString
    - variable:
        key: TEST_ANOTHER_EXAMPLE
        value: thisIsAlsoAString
    - variable:
        key: AWS_DEFAULT_REGION
        value: us-west-2
    - variable:
        key: AWS_PROFILE
        value: test
  openvpn: 
    configFile: /home/user/.ovpn/test.ovpn
  shell: 
    command: /bin/sh /home/user/someScript.sh
  path:
    include: /home/user/go/bin
    exclude: /home/user/experimental/bin
  symlink:
    source: /home/user/git/dotfiles/test/someConfigFile.json
    dest: /home/user/.someConfigFile.json


```
### Supported Configuration Blocks
`ctxman` supports the following configuration options.

#### kubernetes
The `kubernetes` block allows you to define a preferred cluster and optionally a preferred namespace. `ctxman` will edit your `~/.kube/config` file to set these values.

#### env
The `env` block should be a list of environment variables and their values. 

#### openvpn
The `openvpn` block will establish a vpn connection (requires openvpn be intstalled and accessible via the user's `$PATH`) using the specified ovpn profile. If an openvpn connection already exists, `ctxman` will terminate it first.

#### shell
The `shell` block will execute the command provided. If multiple commands are needed, simply place them in a script and call the script in this block.

#### path
The `path` block allows the user to specify directories to be included or excluded in the user's `$PATH`. Multiple comma-separated directories can be specified for both `include` and `exclude`.

#### symlink
The `symlink` will create a symlink from the `source` to the `dest`.

## Usage
```
$ ctxman <environment> <namespace>
```

Simply specify the desired environment, and `ctxman` will evaluate the blocks provided by the user and make any chagnes necessary. The `<namespace>` argument is completely optional, if you don't have a preferred namespace for a particular environment, simply omit it.
