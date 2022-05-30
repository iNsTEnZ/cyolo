package service

import (
	"cyolo-exercise/dao"
	"cyolo-exercise/model"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
)

type Service struct {
	storage dao.DAO
	sorted  []model.Pair // Contains the already sorted data according to histogram
	index   sync.Map     // Contains the mapping of the sorted data so that we won't have to access "database" every time
	lock    sync.RWMutex
}

func NewService(storage dao.DAO) *Service {
	return &Service{
		storage: storage,
	}
}

func (srv *Service) Histogram(top int) []model.Pair {
	if top >= len(srv.sorted) {
		return srv.sorted
	}

	return srv.sorted[:top]
}

func (srv *Service) Process(data string) []error {
	var wg sync.WaitGroup
	errors := make([]error, 0)
	words := strings.Split(data, ",")

	if len(words) == 0 {
		log.Printf("no words to process")
		return nil
	}

	for _, word := range words {
		if word == "" {
			continue
		}

		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			err := srv.handleWord(word)

			if err != nil {
				errors = append(errors, err)
			}
		}(word)
	}

	wg.Wait()
	return errors
}

func (srv *Service) handleWord(word string) error {
	value, err := srv.storage.Get(word, 0)

	if err != nil {
		return fmt.Errorf("error while trying to get value of '%s'", word)
	}

	if intValue, ok := value.(int); ok {
		err = srv.storage.Put(word, intValue+1)

		if err != nil {
			return fmt.Errorf("error while trying to put value of '%s'", word)
		}

		location, exists := srv.index.Load(word)

		if !exists {
			srv.lock.Lock()
			srv.sorted = append(srv.sorted, model.Pair{Key: word, Value: 1})
			srv.lock.Unlock()
			srv.index.Store(word, len(srv.sorted)-1)
		} else {
			srv.lock.Lock()
			srv.sorted[location.(int)].Value = intValue + 1
			srv.lock.Unlock()
		}

		srv.sortData()
	}

	return nil
}

func (srv *Service) sortData() {
	srv.lock.RLock()
	sort.Slice(srv.sorted, func(i, j int) bool {

		result := srv.sorted[i].Value > srv.sorted[j].Value

		if result {
			srv.index.Store(srv.sorted[i].Key, j)
			srv.index.Store(srv.sorted[j].Key, i)
		}

		return result
	})
	srv.lock.RUnlock()
}
