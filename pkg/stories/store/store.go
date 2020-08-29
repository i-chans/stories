package store

type Store struct {
	ss StoriesStore
}

func (s *Store) StoriesStore() StoriesStore {
	return s.ss
}

func NewStore(ss StoriesStore) *Store {
	return &Store{
		ss: ss,
	}
}
