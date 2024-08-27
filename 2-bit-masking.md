 > The process of modifying and utilizing binary representations of numbers or any other data is known as bitmasking

Quetion: 

find ith bit ? 
```
0010101 , find 5th bit ? 

take 1 and shift(<<) it right by ith bit 1 << i 
perform and(&) 
if answer > 0 then 1 else 0
```

set ith bit ? 
```
take 1 and shift(<<) it right by ith bit 1 <<i
perform or(|)
```

clear ith bit ?
```
take 1 and shit(<<) it right by ith bit 1 << i 
invert it ~(num)
peform and (&)
```