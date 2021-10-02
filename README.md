# dt_go_stack
#### Simple Implementation of a Stack in go

init a new stack:
``` go 
import ( 
  st github.com/Bet20/dt_go_stack
)

func main() {
	s := st.stack{nodes: nil, size: 64}
}
```

functions :
- Push()
- Peek()
- Pop()
- IsEmpty()
- IsFull()
