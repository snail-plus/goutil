package errutil

type Block struct {
	Try     func()
	Catch   func(interface{})
	Finally func()
}

func (b *Block) Do() {
	if b.Finally != nil {
		defer b.Finally()
	}

	if b.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				b.Catch(r)
			}
		}()
	}

	b.Try()
}
