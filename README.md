# safe-encoding
Custom base64 encoding written in Go. U can decode text only with this library

# Install
```
go get github.com/dinaxu-attack/go-safe-encoding
```

# Usage
```
package main

import (
	"github.com/dinaxu-attack/go-safe-encoding"
)

func main() {

	println("Encoded:")
	enc := safe.Encode("Hello World")
	println(enc)

	println("Decoded:")
	println(safe.Decode(enc))
}
```

Output
```
Encoded:
!+=|Q/G|b/y/9#2/V#g|8-G|b$s|V$G-S#!+
Decoded:
Hello World
```

# Proof
![image](https://user-images.githubusercontent.com/102496559/205692691-a3a180fc-e7dc-4a1a-9a26-1d8716aa4902.png)

