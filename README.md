# hs
`hs` a tool to enhance `httpie`/`curl`. 

`httpie`/`curl` is a powerful http client. But it can't save history request and replay it. 
`hs` afford abilities to save history request and can replay it when you need. 

`hs` has two subcommands: `workspace`(alias `w`) and `session`(alias `s`).
A request called a session. Every session save to a workspace. You need to create a workspace before save to it. 
A workspace is a directory stored in `--dir` directory. A session is a file stored in a workspace. 

A session is a standard http message. You can write it by yourself. But I suggest you get it from the 
`httpie --offline` output.

# Usage
```bash
hs workspace --create test # create a workspace named test
hs w -c test # create a workspace named test
hs w -l # list workspace
hs w --show-path test # show the path of test
hs w -r test # remove test workspace

https --offline google.com -pHB | hs s -w test -s google # save a session named google, the http message is from httpie
https --offline google.com -pHB | hs session --workspace test --save google # save a session named google, the http message is from httpie
hs s -w test -l # list session in the test workspace
hs s -w test --show-path google # show the path of google
hs s -w test -rm google # remote google session
hs s -w test -r google --raw # print the http message of google
hs s -w test -r google --httpie --https # use httpie to replay the google session
hs s -w test -r google --curl # use curl to replay the google session
```
