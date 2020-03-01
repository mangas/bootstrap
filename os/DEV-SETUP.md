# DEV Setup

As a Dev to contribute you need to do the following.

## Gmail and Github accounts

- Make a new gmail account called "winwiselyXXX@gmail.com".
	- Pick a number from 100 and upwards to replace the XXX.

- Make a new Chrome Profile and use this account in it. So you do not leak.

- Make a new github account using that email address.

- Tell admin your new gmail and github account, so he can add you to the Team on github.

## SSH local setup

- REF: Setup many accounts in ssh config:  https://medium.com/@xiaolishen/use-multiple-ssh-keys-for-different-github-accounts-on-the-same-computer-7d7103ca8693

- Delete the shit in your global git config. Your leaking ..
	- ``` git config --list ```

- same but shows where the values come from.
	- ``` git config -l ```

- Make a new key
	- ``` ssh-keygen -t rsa -b 4096 -C "winwiselyXXX_github" ```


- Add ssh key 
	- ``` ssh-add ~/.ssh/winwiselyXXX_github  ```

- List added ssh 
	- ``` ssh-add -l  ```


- Add the public key to github on the web site
	- https://github.com/settings/keys

## SSH Config file

Needs this in it:

```

# winwiselyXXX
# https://github.com/winwiselyXXX/dev
# e.g: git clone git@github.com-winwiselyXXX:winwiselyXXX/dev.git
Host github.com-winwiselyXXX
 HostName github.com
 User git
 UseKeychain yes
 AddKeysToAgent yes
 IdentityFile ~/.ssh/winwiselyXXX_github
 
```


## Install OS tools

In this folder are the tools you need.

Its basically flutter and golang. If you already have these then your fine.

Also make sure you install hover in the Tools folder. See make file
