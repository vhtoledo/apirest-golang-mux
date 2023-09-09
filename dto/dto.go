package dto

type CategoriaDto struct {
	Nombre string `json:"nombre"`
}
type ProductoDto struct {
	Nombre      string `json:"nombre"`
	Precio      int    `json:"precio"`
	Stock       int    `json:"stock"`
	Descripcion string `json:"descripcion"`
	CategoriaID uint   `json:"categoria_id"`
}