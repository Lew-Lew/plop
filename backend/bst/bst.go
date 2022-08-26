package bst

import "strings"

// déclaration de la classe Node qui prend 3 variables d'instance
type Node struct {
	Letter   string
	End      bool
	Children map[string]*Node
}

// création d'une instance de la classe Node
func NewNode() *Node {
	return &Node{
		Letter:   "",
		End:      false,
		Children: make(map[string]*Node),
	}
}

// méthode de la classe Node. Comme c'est un setter on met *Node pour signifier qu'on modifie et non qu'on lit.
func (n *Node) InsertWord(word string) {
	word = strings.ToUpper(word)
	current := n // le noeud actuel
	for i := 0; i < len(word); i++ {
		next, ok := current.Children[string(word[i])] // next: le noeud où on veut aller, ok: si la valeur exist ou pas
		// si ça n'existe pas:
		if !ok {
			next = NewNode()                         // on crée un nouveau Noeud
			next.Letter = string(word[i])            // on lui assigne la lettre sur laquelle on est
			current.Children[string(word[i])] = next // dans la map des children, à la clé de la lettre on lui assigne notre nouveau noeud comme valeur
		}
		current = next // on réassigne la valeur de current qui devient notre "nouveau" noeud
	}
	current.End = true
}

// méthode de classe (getter) => vérifie si un mot existe. Retour boolean.
func (n *Node) WordExist(word string) bool {
	word = strings.ToUpper(word)
	current := n
	for i := 0; i < len(word); i++ {
		next, ok := current.Children[string(word[i])]
		if !ok {
			return false
		}
		current = next
	}
	return current.End
}
