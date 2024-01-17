package main
func cas(val *int32,old,new int32)bool
func semacquire(*int32)
func semrelease(*int32)

type Mutex struct{
	key int32
	sema int32
}

func xadd(val *int32,delta int32)(new int32){
	for{
		v := *val
		if cas(val,v,v+delta){
			return v+delta
		}
		panic("unreached")
	}
}

func (m *Mutex)Lock(){
	if xadd(&m.key,1)==1{
		return
	}
	semacquire(&m.sema)
}

func (m *Mutex)Unlock(){
	if xadd(&m.key,-1)==0{
		return
	}
	semrelease(&m.sema)
}