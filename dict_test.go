package gonlp

import "testing"

func TestDictionary(t *testing.T) {
	token := []string{"dog", "cat", "horse"}
	doc1 := []string{"dog", "cat"}
	doc2 := []string{"dog", "horse"}
	doc3 := []string{"dog", "horse", "elephant"}

	dict := NewDefaultDictionary()
	dict.AddDocuments([][]string{doc1, doc2})
	for _, token := range token {
		if id := dict.Token2ID(token); id < 0 {
			t.Error("No id for token")
		}
	}

	// res := dict.Doc2Bow(doc1)
	// reflect.DeepEquals(res, [
	// TODO: check doc2bow

}
