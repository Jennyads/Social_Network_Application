package controllers

import "net/http"

//CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Criando usuário!"))
}
//BuscarUsuarios busca todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Buscando todos os usuários!"))
}
//BuscarUsuario busca um usuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Buscando um usuário!"))
}
//AtualiarUsuario altera as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Atualizando suário!"))
}
//DeletarUsuario exclui as informações de um usuário no banco
func DeletarUsuario(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Deletando um usuário!"))
}
