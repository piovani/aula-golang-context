# Aula Golang Context

This project was used to exemplify Context in the YouTube video:

## Types

### Background

Retorna um contexto vazio, onde pode ser atribuidos valores e estendidos posteriormente pelas outras funções

```go
ctx := context.Background()
```

### TODO

Tambem retorna um contexto vazio, é mais usado para casos genericos onde não se sabe que contexto deve ser usado

```go
ctx := context.TODO()
```

### WithCancel

retorna um contexto igual ao que foi passado como parametro, no atributo Done do cancel ele possui um canal justamente para controlar 

```go
ctx, cancelFunc := context.WithCancel(context.Background())
```

### WithDeadline

### WithTimeout

### WithValue