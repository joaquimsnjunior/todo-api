package schemes

import "time"

// Tasks representa uma tarefa no sistema
type Tasks struct {
	ID        string    `db:"id" json:"id"`
	Title     string    `db:"title" json:"titulo"`
	Completed bool      `db:"completed" json:"completado"`
	CreatedAt time.Time `db:"created_at" json:"criado_em"`
	UpdatedAt time.Time `db:"updated_at" json:"atualizado_em"`
}
