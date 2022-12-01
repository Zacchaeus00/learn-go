# Go和Python用法区别：

1. Go可以直接修改outer scope定义的变量，py需要用nonlocal或者global关键字。（闭包思想）
2. Go for loop和if内定义的变量会在外层失效（for/if即scope）。
3. Go把slice传入函数，定点修改会改变外层slice，这点和py一样。但是append不会影响到外层，而py会。如果需要扩增需要传指针。
4. Go的map如果key不存在则返回0，而py会报错。
5. Py slice是copy而go是reference。
