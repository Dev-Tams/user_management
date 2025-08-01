package main



type Product struct {
	ID    int
	Name  string
	Price float64
	Tags  []string
}

var products = []Product{
	{
		ID: 1, Name: "Rice", Price: 8.56, Tags: []string{"rice", "fast", "parboiled"},
	},
	{
		ID: 4, Name: "Yam", Price: 4.56, Tags: []string{"white", "sweet"},
	},
	{
		ID: 3, Name: "Sugar", Price: 1.9, Tags: []string{"sweet"},
	},
}



