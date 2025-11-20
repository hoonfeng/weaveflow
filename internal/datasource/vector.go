package datasource

import (
    "errors"
    "math"
    "sort"
)

type VectorStore interface {
    Upsert(collection string, id string, vec []float64, meta map[string]any) error
    Search(collection string, vec []float64, topK int) ([]VectorResult, error)
    Delete(collection string, id string) error
    EnsureCollection(collection string, size int, metric string) error
}

type VectorResult struct {
    ID    string
    Score float64
    Meta  map[string]any
}

type InMemoryVector struct {
    data map[string]map[string]entry
}

type entry struct { vec []float64; meta map[string]any }

func NewInMemoryVector() *InMemoryVector { return &InMemoryVector{data: map[string]map[string]entry{}} }

func (m *InMemoryVector) Upsert(collection string, id string, vec []float64, meta map[string]any) error {
    if m.data[collection] == nil { m.data[collection] = map[string]entry{} }
    m.data[collection][id] = entry{vec: vec, meta: meta}
    return nil
}

func (m *InMemoryVector) Search(collection string, vec []float64, topK int) ([]VectorResult, error) {
    col := m.data[collection]
    if col == nil { return nil, errors.New("collection not found") }
    res := make([]VectorResult, 0, len(col))
    for id, e := range col {
        score := cosine(vec, e.vec)
        res = append(res, VectorResult{ID: id, Score: score, Meta: e.meta})
    }
    sort.Slice(res, func(i, j int) bool { return res[i].Score > res[j].Score })
    if topK > 0 && topK < len(res) { res = res[:topK] }
    return res, nil
}

func (m *InMemoryVector) Delete(collection string, id string) error {
    col := m.data[collection]
    if col == nil { return nil }
    delete(col, id)
    return nil
}

func (m *InMemoryVector) EnsureCollection(collection string, size int, metric string) error { return nil }

func cosine(a, b []float64) float64 {
    if len(a) == 0 || len(b) == 0 { return 0 }
    n := min(len(a), len(b))
    var dot, na, nb float64
    for i := 0; i < n; i++ { dot += a[i]*b[i]; na += a[i]*a[i]; nb += b[i]*b[i] }
    if na == 0 || nb == 0 { return 0 }
    return dot / (math.Sqrt(na) * math.Sqrt(nb))
}

func min(a, b int) int { if a < b { return a } ; return b }