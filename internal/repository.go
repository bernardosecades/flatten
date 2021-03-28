package flatten

type Repository interface {
	SaveHistory(req []byte, res []byte, depth int) error
	GetHistoryByLimit(l uint) ([]History, error)
}
