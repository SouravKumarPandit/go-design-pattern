package _0Temp
type OrderStatusType string

var OrderStatus = struct {
	APPROVED         OrderStatusType
	APPROVAL_PENDING OrderStatusType
	REJECTED         OrderStatusType
	REVISION_PENDING OrderStatusType
}{
	APPROVED:         "approved",
	APPROVAL_PENDING: "approval pending",
	REJECTED:         "rejected",
	REVISION_PENDING: "revision pending",
}
//--------------------------------------skip value
const (
	C1 = iota + 1
	_
	C3
	C4
)
//-------------------------------------- string type
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}