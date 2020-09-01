# singleflight-go

- func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error)
    - If f
    n() is panic and is recovered, it may cause key residue and cause magical problems