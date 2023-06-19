package entity

type OrderQueue struct {
	Orders []*Order
}

// Less - Comparar o preço entre duas ordens
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Swap - Inverte a ordem das ordens
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Len - Retorna o tamanho da fila de nossas ordens
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Push - Adicionar uma ordem na fila
func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))

}

// Pop - Remove uma ordem
func (oq *OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
