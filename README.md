git-rebase-merge
============

A bash script to rebase and merge a Git topic branch to current branch.

### Usage

```sh
# Assuming current branch is master,
# it reabses a list of topic branches on top of master and then merge into master
$ git-rebase-merge topic-branch1 topic-branch2 ...
```

### Install

```sh
$ curl https://raw.github.com/jingweno/git-rebase-merge/master/install | sh
```

### Uninstall

```sh
$ curl https://raw.github.com/jingweno/git-rebase-merge/master/uninstall | sh
```

### License

See LICENSE
