package gonlp

import "sync"

type Dictionary interface {
	AddDocuments(docs [][]string)
	Token2ID(token string) int
	Doc2Bow(doc []string) [][2]int // Sparse vector (id, number of occurrences)
}

type DefaultDictionary struct {
	m    map[string]int
	lock *sync.RWMutex
}

func NewDefaultDictionary() *DefaultDictionary {
	return &DefaultDictionary{
		m:    make(map[string]int),
		lock: &sync.RWMutex{},
	}
}

// AddDocuments adds
func (dd *DefaultDictionary) AddDocuments(docs [][]string) {
	dd.lock.Lock()
	for _, doc := range docs {
		for _, word := range doc {
			if _, ok := dd.m[word]; !ok {
				dd.m[word] = len(dd.m)
			}
		}
	}
	dd.lock.Unlock()
}

// Token2ID returns the id of the given token in the dictionary, otherwise -1.
func (dd *DefaultDictionary) Token2ID(token string) int {
	dd.lock.RLock()
	defer dd.lock.RUnlock()
	if id, ok := dd.m[token]; ok {
		return id
	}
	return -1
}

// Doc2Bow converts a document (list of words) into a bag-of-words format (list of tuple [token_id, frequency in document]).
func (dd *DefaultDictionary) Doc2Bow(doc []string) [][2]int {
	dd.lock.RLock()
	m := make(map[int]int)
	for _, token := range doc {
		if id, ok := dd.m[token]; ok {
			// if c, ok := m[id]; !ok {
			// 	m[id]=0
			// }
			m[id]++
		}
	}
	dd.lock.RUnlock()

	vec := make([][2]int, len(m))
	i := 0
	for id, c := range m {
		vec[i] = [2]int{id, c}
		i++
	}
	return vec
}
