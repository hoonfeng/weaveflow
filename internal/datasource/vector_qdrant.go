package datasource

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type QdrantVector struct {
    Endpoint string
    APIKey   string
}

func NewQdrantVector(endpoint, apiKey string) *QdrantVector { return &QdrantVector{Endpoint: endpoint, APIKey: apiKey} }

func (q *QdrantVector) Upsert(collection string, id string, vec []float64, meta map[string]any) error {
    url := fmt.Sprintf("%s/collections/%s/points", q.Endpoint, collection)
    payload := map[string]any{
        "points": []map[string]any{{"id": id, "vector": vec, "payload": meta}},
    }
    return q.post(url, payload, nil)
}

func (q *QdrantVector) Search(collection string, vec []float64, topK int) ([]VectorResult, error) {
    url := fmt.Sprintf("%s/collections/%s/points/search", q.Endpoint, collection)
    payload := map[string]any{"vector": vec, "limit": topK}
    var out struct{ Result []struct{ ID any `json:"id"`; Score float64 `json:"score"`; Payload map[string]any `json:"payload"` } `json:"result"` }
    if err := q.post(url, payload, &out); err != nil { return nil, err }
    res := make([]VectorResult, 0, len(out.Result))
    for _, r := range out.Result { res = append(res, VectorResult{ID: fmt.Sprint(r.ID), Score: r.Score, Meta: r.Payload}) }
    return res, nil
}

func (q *QdrantVector) SearchWithOptions(collection string, vec []float64, topK int, options map[string]any) ([]VectorResult, error) {
    url := fmt.Sprintf("%s/collections/%s/points/search", q.Endpoint, collection)
    payload := map[string]any{"vector": vec, "limit": topK}
    if options != nil {
        if f, ok := options["filter"]; ok { payload["filter"] = f }
        if p, ok := options["params"]; ok { payload["params"] = p }
    }
    var out struct{ Result []struct{ ID any `json:"id"`; Score float64 `json:"score"`; Payload map[string]any `json:"payload"` } `json:"result"` }
    if err := q.post(url, payload, &out); err != nil { return nil, err }
    res := make([]VectorResult, 0, len(out.Result))
    for _, r := range out.Result { res = append(res, VectorResult{ID: fmt.Sprint(r.ID), Score: r.Score, Meta: r.Payload}) }
    return res, nil
}

func (q *QdrantVector) Delete(collection string, id string) error {
    url := fmt.Sprintf("%s/collections/%s/points/delete", q.Endpoint, collection)
    payload := map[string]any{"points": []any{id}}
    return q.post(url, payload, nil)
}

func (q *QdrantVector) EnsureCollection(collection string, size int, metric string) error {
    if metric == "" { metric = "Cosine" }
    url := fmt.Sprintf("%s/collections/%s", q.Endpoint, collection)
    payload := map[string]any{
        "vectors": map[string]any{"size": size, "distance": metric},
    }
    // Try create, if already exists ignore
    var out map[string]any
    err := q.put(url, payload, &out)
    if err != nil {
        // Some deployments use POST
        err2 := q.post(url, payload, &out)
        if err2 != nil { return err2 }
    }
    return nil
}

func (q *QdrantVector) post(url string, payload any, dest any) error {
    b, _ := json.Marshal(payload)
    req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    if q.APIKey != "" { req.Header.Set("api-key", q.APIKey) }
    resp, err := http.DefaultClient.Do(req)
    if err != nil { return err }
    defer resp.Body.Close()
    if resp.StatusCode >= 300 { return fmt.Errorf("qdrant http %d", resp.StatusCode) }
    if dest != nil { return json.NewDecoder(resp.Body).Decode(dest) }
    return nil
}

func (q *QdrantVector) put(url string, payload any, dest any) error {
    b, _ := json.Marshal(payload)
    req, _ := http.NewRequest("PUT", url, bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    if q.APIKey != "" { req.Header.Set("api-key", q.APIKey) }
    resp, err := http.DefaultClient.Do(req)
    if err != nil { return err }
    defer resp.Body.Close()
    if resp.StatusCode >= 300 { return fmt.Errorf("qdrant http %d", resp.StatusCode) }
    if dest != nil { return json.NewDecoder(resp.Body).Decode(dest) }
    return nil
}