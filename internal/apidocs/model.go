package apidocs

type DocParam struct {
    Source      string            `json:"source"`
    Name        string            `json:"name"`
    Type        string            `json:"type"`
    Required    bool              `json:"required"`
    Description string            `json:"description"`
    Constraints map[string]string `json:"constraints"`
}

type DocResponse struct {
    Status         int         `json:"status"`
    Body           string      `json:"body"`
    BodyDescription string     `json:"body_description"`
    Headers        []DocHeader `json:"headers"`
    Schema         []DocSchema `json:"schema"`
}

type DocHeader struct {
    Name        string `json:"name"`
    Value       string `json:"value"`
    Description string `json:"description"`
}

type DocSchema struct {
    Name        string      `json:"name"`
    Type        string      `json:"type"`
    Description string      `json:"description"`
    Children    []DocSchema `json:"children"`
}

type DocEndpoint struct {
    Module      string      `json:"module"`
    Endpoint    string      `json:"endpoint"`
    Method      string      `json:"method"`
    Path        string      `json:"path"`
    Title       string      `json:"title"`
    Description string      `json:"description"`
    Auth        string      `json:"auth"`
    Roles       []string    `json:"roles"`
    Params      []DocParam  `json:"params"`
    Files       []DocParam  `json:"files"`
    Response    DocResponse `json:"response"`
    Errors      []DocError  `json:"errors"`
}

type DocError struct {
    Code        string `json:"code"`
    Message     string `json:"message"`
    Description string `json:"description"`
}