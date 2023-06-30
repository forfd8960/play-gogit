## Git Init

```sh

sh-3.2$ go build -o goinit main.go
sh-3.2$ ./goinit -h
Usage of ./goinit:
  -gitpath string
        the git repo path (default ".")
sh-3.2$ ./goinit -gitpath testrepo
```

```
play-gogit/gitinit  cd testrepo                                                    ✔  10857  23:32:30
play-gogit/gitinit/testrepo   master  git status                                 ✔  10858  23:32:33
On branch master

No commits yet

nothing to commit (create/copy files and use "git add" to track)
gitinit/testrepo   master  cd .git

play-gogit/gitinit/testrepo/.git   master  ll                                    ✔  10865  23:33:08
total 8
  staff    23B Jun 30 23:30 HEAD
  staff   128B Jun 30 23:30 objects
  staff   128B Jun 30 23:30 refs
```
