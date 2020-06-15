# Context Manager

`ctxman` makes managing your many environments and contexts a breeze. In the modern DevOps world, developers and operators often have to change contexts or environments many times a day, switching from prod to dev to test or internal tools environments. Each of the changes may involve: changing environment variables, swapping config files, setting a new `kubectl` cluster and namespace, a new vpn connection, etc.

Rather than write a bash script or handle each of these changes manually, simply define your environments in a `ctxman` config file and let it handle the tedium for you!

## Installation

```
$ go get -u github.com/jamesbaker1901/ctxman
```

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
    - PROD_EXAMPLE=thisIsAString
    - PROD_ANOTHER_EXAMPLE=thisIsAlsoAString
    - AWS_DEFAULT_REGION=us-west-2
    - AWS_PROFILE=prod
  openvpn: 
    configFile: /home/user/.ovpn/prod-us.ovpn
  shell: 
    command: /bin/sh someScript.sh
  path:
    include: /home/user/go/bin
    exclude: /home/user/experimental/bin,/opt/testing/bin
  symlink:
    source: /home/user/git/dotfiles/prod/someConfigFile.json
    dest: /home/user/.someConfigFile.json

test:
  kubernetes:
    cluster: test
    namespace: app
  env:
    - TEST_EXAMPLE=thisIsAString
    - TEST_ANOTHER_EXAMPLE=thisIsAlsoAString
    - AWS_DEFAULT_REGION=us-west-2
    - AWS_PROFILE=test
  openvpn: 
    configFile: /home/user/.ovpn/test.ovpn
  shell: 
    command: /bin/sh someScript.sh
  path:
    include: /home/user/go/bin,/opt/testing/bin
  symlink:
    source: /home/user/git/dotfiles/test/someConfigFile.json
    dest: /home/user/.someConfigFile.json
```
### Supported Configuration Blocks
`ctxman` supports the following configuration options.

#### kubernetes
The `kubernetes` block allows you to define a preferred cluster and optionally a preferred namespace. `ctxman` will edit your `~/.kube/config` file to set these values.

#### env
The `env` block should be a list of environment variables and their values. The format should be `ENVIRONMENT_VARIABLE="Desired Value"`, basically exactly what you would put after `export` if setting that value manually.

#### openvpn
The `openvpn` block will establish a vpn connection (requires openvpn be intstalled and accessible via the user's `$PATH`) using the specified pvpn profile. If an openvpn connection already exists, `ctxman` will terminate it first.

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
