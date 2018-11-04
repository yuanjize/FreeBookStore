#!/usr/bin/expect

spawn go install github.com/yuanjize/FreeBookStore

spawn scp /home/yuanjize/golangs/bin/FreeBookStore yuanjize@207.148.89.29:/home/yuanjize/FreeBookStore
expect "*password:"
send "yuanjize\r"
interact
spawn scp /home/yuanjize/golangs/src/github.com/yuanjize/FreeBookStore/config.yaml  yuanjize@207.148.89.29:/home/yuanjize/FreeBookStore
expect "*password:"
send "yuanjize\r"
interact