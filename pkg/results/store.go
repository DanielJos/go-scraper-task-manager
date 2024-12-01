package results

type Result struct {
	URL string
	// Data SomeDataType
}

type ResultsStore struct {
	Results []Result
}

func New() ResultsStore {
	return ResultsStore{

	}
}

func (r *ResultsStore) Add(result Result) {
	r.Results = append(r.Results, result)
}