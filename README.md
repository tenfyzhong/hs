# hs
`hs` a tool to enhance `httpie`/`curl`. 

`httpie`/`curl` is a powerful http client. But it can't save history command and replay it. 
`hs` afford abilities to save history command and can replay it when you need. 

In `hs`, a http message which can be save to a workspace called a session. 
The workspace which to save should be create first. 
A workspace is a directory stored in `--dir` directory. 
A session is a file stored in a workspace. 

`hs` has three subcommands: `workspace`(alias `w`), `session`(alias `s`) and `transfer`(alias `t`).

## command
### `workspace`
The `workspace` command operate a workspace. It has four flags.
- `--create` or `-c` create a workspace.
- `--list` or `-l` list all workspaces.
- `--show-path` or `-P` show the path of the workspace.
- `--remove` or `-R` remove a workspace. 

### `session`
The `session` command operate a session. It has ten flags.  

- `--workspace` or `-w` special the workspace to work.
- `--create` or `-c` create a session into the workspace. Suggest to get the http message by call the `httpie --offline` command. 
- `--list` or `-l` list the sessions in the workspace. 
- `--remove` or `-R` remove a session in the workspace. 
- `--show-path` or `-P` show the path of the session in the workspace. 
- `--replay` or `-r` replay a session in the workspace. The rest args will be add to the `httpie`/`curl` command. 
- `--httpie` or `-H` use httpie to replay the session. Only works with `--replay`.
- `--curl` or `-C` use curl to replay the session. Only works with `--replay`. 
- `--raw` print the raw http message. Only works with `--replay`.
- `--https` or `-S` replay as https. Only works with `--httpie` or `--curl`. 

### `transfer`
The `transfer` command transfer a http message to a `httpie`/`curl` command directly. It has three flags.   
The rest args flag will be add to the `httpie`/`curl` command. 
- `--httpie` or `-H` transfer to httpie command. 
- `--curl` or `-C` transfer to curl command.
- `--https` or `-S` use https.

# Usage
```bash
hs workspace --create test # create a workspace named test
hs w -c test # create a workspace named test
hs w -l # list workspace
hs w --show-path test # show the path of test
hs w -R test # remove test workspace

https --offline google.com -pHB | hs s -w test -c google # create a session named google, the http message is from httpie
https --offline google.com -pHB | hs session --workspace test --create google # create a session named google, the http message is from httpie
hs s -w test -l # list session in the test workspace
hs s -w test --show-path google # show the path of google
hs s -w test -rm google # remote google session
hs s -w test -r google --raw # print the http message of google
hs s -w test -r google -C | sh # use curl to replay the google session
hs s -w test -r google -H -S -- -pHbhb | sh # use httpie to replay the google session

https --offline google.com -pHB | hs t # transfer to curl command
```

# autocomplete
## Bash
```bash
sudo cp autocomplete/hs-complete.bash /etc/bash_completion.d/hs
source /etc/bash_completion/hs # add to your .bashrc
```

## Zsh
```bash
source autocomplete/hs-complete.zsh # add to your .zshrc
```

## Power shell
```bash
& autocomplete/hs.ps1
```

# Debug
Set `ENABLE_HS_LOG` to 1, it will log errors to `hs.log` in the current directory. 
