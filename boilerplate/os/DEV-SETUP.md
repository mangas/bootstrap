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


## when you fork

add this to .git/config:
````
[user]
	email = winwisely99@gmail.com
````

then:

````
make git-fork-setup
````



## Install golang and then tools

Its basically flutter and golang. If you already have these then your fine.

Dont forget to put in the right env variables. !!

For mac. https://github.com/getcouragenow/bootstrap/blob/master/os/mac/.bashrc#L15
For other OS's see folder for your OS.

Then in Tools, you need hover and i18n.
See: https://github.com/getcouragenow/bootstrap/tree/master/tool
- run the make file for hover and i18n folders.

