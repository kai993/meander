package meander

// Facade : https://ja.wikipedia.org/wiki/Facade_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3
type Facade interface {
	Public() interface{}
}

// Public : 構造体を表す外部向けのビューを返す。
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}

	return o
}
