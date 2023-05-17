package snippet

import (
	"fmt"
)

type NewPost struct {
	Id				int
	Content 	string
	Author 		string
}

var PostById map[int]*NewPost
var PostsByAuthor map[string][]*NewPost

func store(post NewPost) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func InMemoryMain() {
	PostById = make(map[int]*NewPost)
	PostsByAuthor = make(map[string][]*NewPost)

	post1 := NewPost{Id: 1, Content: "Hello World!", Author: "Amos"}
	post2 := NewPost{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := NewPost{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := NewPost{Id: 4, Content: "Greetings", Author: "Amos"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Amos"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}