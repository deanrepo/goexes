package newsfeed

type (
	Getter interface {
		GetAll() []Item
	}

	Adder interface {
		Add(item Item)
	}

	Item struct {
		Title string `json:"title"`
		Post  string `json:"post"`
	}

	Repo struct {
		Items []Item
	}
)

func New() *Repo {
	return &Repo{Items: []Item{}}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
