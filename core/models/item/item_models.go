package item

type Item struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
	Weight   int    `json:"weight" bson:"weight"`
}

type InsertItem struct {
	Name     string `json:"name" bson:"name" validate:"required"`
	Quantity int    `json:"quantity" bson:"quantity" validate:"required"`
	Weight   int    `json:"weight" bson:"weight" validate:"required"`
}
