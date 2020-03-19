package transecode

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Transecodes struct {
	Items []Item
}

func New() *Transecodes {
	return &Transecodes{}
}

func (r *Transecodes) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Transecodes) GetAll() []Item {
	return r.Items
}
