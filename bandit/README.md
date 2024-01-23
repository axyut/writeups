# Bandit

# Bandit
How to start? visit [https://overthewire.org/wargames/bandit/bandit0.html](https://overthewire.org/wargames/bandit/bandit0.html)

- `ssh **bandit.labs.overthewire.org -p 2220**`

Report on all the levels from Over The Wire's Bandit documented [here](https://axyut.notion.site/Bandit-3bffccb4fafb4ef2af4696036ba65efd?pvs=4)


## Level 0

Credentials

```markdown
❯ ssh [bandit0@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: bandit0
```

Commands

```markdown
cat readme
`NH2SXQwcBdpmTEzi3bvBHMM9H66vVXjL`
```

## Level 1

Credentials

```markdown
❯ ssh [bandit1@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: NH2SXQwcBdpmTEzi3bvBHMM9H66vVXjL
```

Commands

```markdown
cat ./-
`rRGizSaX8Mk1RTb1CNQoXTcYZWU6lgzi`
```

## Level 2

Credentials

```markdown
❯ ssh [bandit2@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: rRGizSaX8Mk1RTb1CNQoXTcYZWU6lgzi
```

Commands

```markdown
cat "spaces in this filename"
OR
cat spaces\ in\ this\ filename
`aBZ0W5EmUfAf7kHTQeOwd8bauFJ2lAiG`
```

## Level 3

Credentials

```markdown
❯ ssh [bandit3@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:aBZ0W5EmUfAf7kHTQeOwd8bauFJ2lAiG
```

Commands

```markdown
cd inhere
ls -la
cat .hidden
`2EW7BBsr6aMMoJ2HjW067dm8EgX26xNe`
```

## Level 4

Credentials

```markdown
❯ ssh [bandit4@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: 2EW7BBsr6aMMoJ2HjW067dm8EgX26xNe
```

Commands

```markdown
cd inhere
cat ./*
OR 
file ./*
cat ./-file07
`lrIWWI6bB37kxfiCQZqUdOIYfr6eEeqR`
```

## Level 5

Credentials

```markdown
❯ ssh [bandit5@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: lrIWWI6bB37kxfiCQZqUdOIYfr6eEeqR
```

Commands

```markdown
cd inhere
find -type f -size 1033c
cd maybehere07
cat .file2
`P4L4vucdmLnm8I7Vl7jG1ApGSfjYKqJU`
```

## Level 6

Credentials

```markdown
❯ ssh [bandit6@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: P4L4vucdmLnm8I7Vl7jG1ApGSfjYKqJU
```

Commands

```markdown
$ find / -type f -size 33c -group bandit6 -user bandit7 2>&1 | grep -v "Permission denied"
cd /var/lib/dpkg/info/bandit7.password
`z7WtoNQU2XfjmMtWA8u5rN4vzqu4v99S`
```

## Level 7

Credentials

```markdown
❯ ssh [bandit7@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: z7WtoNQU2XfjmMtWA8u5rN4vzqu4v99S
```

Commands

```markdown
cat data.txt | grep millionth
OR
find / -name "data.txt" -exec grep -H 'millionth' {} \; 2>&1 | grep -v "Permission denied"
`TESKZC0XvTetK0S9xNwm25STk5iWrBvP`
```

## Level 8

Credentials

```markdown
❯ ssh [bandit8@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: TESKZC0XvTetK0S9xNwm25STk5iWrBvP
```

Commands

```markdown
sort data.txt | uniq -c | grep "1 "
`EN632PlfYiZbn3PhVK3XOGSlNInNE00t`
```

## Level 9

Credentials

```markdown
❯ ssh [bandit9@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: EN632PlfYiZbn3PhVK3XOGSlNInNE00t
```

Commands

```markdown
strings data.txt | grep "^=="
`G7w8LIi6J3kTb8A7j9LgrywtEUlyyp6s`
```

## Level 10

Credentials

```markdown
❯ ssh [bandit10@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: G7w8LIi6J3kTb8A7j9LgrywtEUlyyp6s
```

Commands

```markdown
cat data.txt | base64 -d
`6zPeziLdR2RKNdNYFNb6nVCKzphlXHBM`
```

## Level 11

Credentials

```markdown
❯ ssh [bandit11@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: 6zPeziLdR2RKNdNYFNb6nVCKzphlXHBM
```

Commands

```markdown
cat data.txt | tr 'A-Za-z' 'N-ZA-Mn-za-m'
`JVNBBFSmZwKKOP0XbFXOoW8chDz5yVRv`
```

## Level 12

Credentials

```markdown
❯ ssh [bandit12@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: JVNBBFSmZwKKOP0XbFXOoW8chDz5yVRv
```

Commands

```markdown
    2  mkdir /tmp/pass
    3  mkdir /tmp/pas
    4  cp data.txt /tmp/pas
    5  cd /tmp/pas
    6  ls
    7  file data.txt 
    8  cat data.txt 
    9  xxd -r data.txt data.out
   10  ls
   11  rm data.out
   12  xxd -r data.txt new
   13  ls
   14  file new
   15  mv new new.gz
   16  gzip -d new.gz 
   17  ls
   18  file new
   19  bzip2 -d new
   20  mv new new.bzip2
   21  ls
   22  bzip2 -d new.out
   23  mv new.out new.bzip2
   24  ls
   25  bzip2 -d new.bzip2
   26  bzip2 -d new
   27  file new.bzip2 
   28  mv new.bzip2 new.gzip
   29  gzip -d new.gzip
   30  mv new.gzip new.gz
   31  gzip -d new.gz
   32  ls
   33  file new
   34  ls
   35  tar -xf new
   36  ls
   37  file data5.bin
   38  tar -xf data5.bin
   39  ls
   40  file data6.bin
   41  bzip2 -d data6.bin
   42  ls
   43  file data6.bin.out 
   44  tar -xf data6.bin.out
   45  ls
   46  file data8.bin
   47  gzip -d data8.bin
   48  mv data8.bin newData.gz
   49  gzip -d newData.gz 
   50  ls
   51  file newData 
   52  cat newData
`wbWdlBxEir4CaE8LaPhauuOo6pwRmrDw`
```

## Level 13

Credentials

```markdown
❯ ssh [bandit13@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: wbWdlBxEir4CaE8LaPhauuOo6pwRmrDw
```

Commands

```markdown
ssh -i sshkey.private bandit14@bandit.labs.overthewire.org -p 2220
```

## Level 14

Credentials

```markdown
❯ ssh [bandit14@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: `either download shh key file or use it directly from level 13`
```

Commands

```markdown
cat /etc/bandit_pass/bandit14 | nc localhost 30000
`jN2kgmIXJ6fShzhT2avhotn4Zcka6tnt`
```

## Level 15

Credentials

```markdown
❯ ssh [bandit15@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: jN2kgmIXJ6fShzhT2avhotn4Zcka6tn
```

Commands

```markdown
cat /etc/bandit_pass/bandit15 | openssl s_client -connect localhost:30001 -quiet
`JQttfApK4SeyHwDlI9SXGR50qclOAil1`
```

## Level 16

Credentials

```markdown
❯ ssh [bandit16@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:JQttfApK4SeyHwDlI9SXGR50qclOAil1
```

Commands

```markdown
cat /etc/bandit_pass/bandit16 | openssl s_client -connect localhost:31790 -quiet
* copy the key
cd /tmp
nano sshkey * paste
chmod 600 sshkey
ssh -i sshkey [bandit17@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
```

## Level 17

Credentials

```markdown
❯ ssh [bandit17@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: sshkey 
```

Commands

```markdown
diff passwords.old passwords.new
`hga5tuuCLF6fFzUpnagiMN8ssu9LFrdg`
```

## Level 18

Credentials

```markdown
❯ ssh [bandit18@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: hga5tuuCLF6fFzUpnagiMN8ssu9LFrdg
```

Commands

```markdown
❯ ssh bandit18@bandit.labs.overthewire.org -p 2220 "cat readme"
`awhqfNnAbc1naukrpqDYcF95h7HoMTrC`
```

## Level 19

Credentials

```markdown
❯ ssh [bandit19@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: awhqfNnAbc1naukrpqDYcF95h7HoMTrC
```

Commands

```markdown
./bandit20-do cat /etc/bandit_pass/bandit20
`VxCazJaVykI6W36BkBU0mJTCM8rR95XT`
```

## Level 20

Credentials

```markdown
❯ ssh [bandit20@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:VxCazJaVykI6W36BkBU0mJTCM8rR95XT
```

Commands

```markdown
nc -lp 31337 < /etc/bandit_pass/bandit20
* another terminal
./suconnect 31337
`NvEJF7oVjkddltPSrdKEFOllh9V1IBcq`
```

## Level 21

Credentials

```markdown
❯ ssh [bandit21@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:NvEJF7oVjkddltPSrdKEFOllh9V1IBcq
```

Commands

```markdown
    1  cat /etc/cron.d/cronjob_bandit22
    2  cat /usr/bin/cronjob_bandit22.sh
    3  cat /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
`WdDozAdTM2z9DiFEQ2mGlwngMfj4EZff`
```

## Level 22

Credentials

```markdown
❯ ssh [bandit22@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:WdDozAdTM2z9DiFEQ2mGlwngMfj4EZff
```

Commands

```markdown
      cat /etc/cron.d/cronjob_bandit23
      cat /usr/bin/cronjob_bandit23.sh
      echo "I am user bandit23" | md5sum
      cat /tmp/8ca319486bfbbc3663ea0fbe81326349
`QYw0Y2aiA672PsMmh9puTQuhoz8SyR2G`
```

## Level 23

Credentials

```markdown
❯ ssh [bandit23@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:QYw0Y2aiA672PsMmh9puTQuhoz8SyR2G
```

Commands

```markdown
cat /etc/cron.d/cronjob_bandit24
cat /usr/bin/cronjob_bandit24.sh
cd /var/spool/bandit24/foo/
vi 1.sh 
	#!/bin/bash
	cat /etc/bandit_pass/bandit24 > /tmp/test/password
* creating script file
chmod 777 script.sh
chmod 777 /tmp/test
cd tmp/test
cat password
`VAfGXJ1PBSsPSnvsjI8p759leLZ9GGar`
```

## Level 24

Credentials

```markdown
❯ ssh [bandit24@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:VAfGXJ1PBSsPSnvsjI8p759leLZ9GGar
```

Commands

```markdown
nc localhost 30002
cd /tmp/test
vi scr.sh
	#!/bin/sh
	for i in {0000..9999}
	do
        echo VAfGXJ1PBSsPSnvsjI8p759leLZ9GGar $i >> possibilities.txt
	done
chmod 777 scr.sh
./scr.sh
cat possibilities.txt | nc localhost 30002 > result.txt
* can take time
sort result.txt | grep -v "Wrong!"
`p7TaowMYrmu23Ol8hiZh9UvD0O9hpx8d`
```

## Level 25

Credentials

```markdown
❯ ssh [bandit25@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: p7TaowMYrmu23Ol8hiZh9UvD0O9hpx8d
```

Commands

```markdown
cat bandit26.sshkey
cp bandit26.sshkey /tmp/test
ssh -i bandit26.sshkey bandit26@localhost
* when More(when terminal size small) 
v 
* go to normal
:set shell=/bin/bash
:e /etc/bandit_pass/bandit26
`c7GvcKlw9mC7aUQaPx7nwFstuAIBw1o1`
```

## Level 26

Credentials

```markdown
❯ ssh [bandit26@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: c7GvcKlw9mC7aUQaPx7nwFstuAIBw1o1
```

Commands

```markdown
./bandit27-do cat /etc/bandit_pass/bandit27
*we can from the level 25
`YnQpBuifNMas1hcUFk70ZmqkhUU2EuaS`
```

## Level 27

Credentials

```markdown
❯ ssh [bandit27@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: YnQpBuifNMas1hcUFk70ZmqkhUU2EuaS
```

Commands

```markdown
git clone ssh://bandit27-git@localhost:2220/home/bandit27-git/repo
cd repo
cat README.md
`AVanL161y9rsbcJIsFHuw35rjaOM19nR`
```

## Level 28

Credentials

```markdown
❯ ssh [bandit28@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: AVanL161y9rsbcJIsFHuw35rjaOM19nR
```

Commands

```markdown
git clone ssh://bandit28-git@localhost:2220/home/bandit27-git/repo
cd repo
cat README.md
git log
git checkout f08b9cc63fa1a4602fb065257633c2dae6e5651b
cat README.md
`tQKvmcwNYcFS6vmPHIUSI3ShmsrQZK8S`

```

## Level 29

Credentials

```markdown
❯ ssh [bandit29@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: tQKvmcwNYcFS6vmPHIUSI3ShmsrQZK8S
```

Commands

```markdown
git clone ssh://bandit29-git@localhost:2220/home/bandit27-git/repo
cd repo
git checkout dev
cat README.md
`xbhV3HpNGlTIdnjUrdAlPzc2L6y9EOnS`
```

## Level 30

Credentials

```markdown
❯ ssh [bandit30@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass:
```

Commands

```markdown
git clone ssh://bandit30-git@localhost:2220/home/bandit27-git/repo
cd repo/
git tag
git show secret
`OoffzGDlzhAlerFJ2cAiz1D41JW1Mhmt`
```

## Level 31

Credentials

```markdown
❯ ssh [bandit31@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: OoffzGDlzhAlerFJ2cAiz1D41JW1Mhmt
```

Commands

```markdown
git clone ssh://bandit31-git@localhost:2220/home/bandit27-git/repo
cd repo
ls
cat README.md
vi key.txt
	May I come in?
git add -f key.txt
git commit -m "done"
git push origin master
`rmCBvG56y58BXzv98yZGdO7ATVL5dW8y`
```

## Level 32

Credentials

```markdown
❯ ssh [bandit32@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: rmCBvG56y58BXzv98yZGdO7ATVL5dW8y
```

Commands

```markdown
$0
vim
  :r /etc/bandit_pass/bandit33
`odHo63fHiFqcWWJG9rLiLDtPm45KzUKy`
```

## Level 33

Credentials

```markdown
❯ ssh [bandit33@bandit.labs.overthewire.org](mailto:bandit0@bandit.labs.overthewire.org) -p 2220
pass: odHo63fHiFqcWWJG9rLiLDtPm45KzUKy
```

Commands

```markdown
ls
cat README.txt
```

![Untitled](Bandit%203bffccb4fafb4ef2af4696036ba65efd/Untitled.png)
