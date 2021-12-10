package main

import (
	"log"

	"github.com/michaeltelford/swoop"
	"github.com/michaeltelford/swoop/component"
	"github.com/michaeltelford/swoop/layout"
	"github.com/michaeltelford/swoop/page"
	"github.com/michaeltelford/swoop/server"
)

type (
	// Post represents a blog post model (grabbed from a Datastore somewhere etc).
	Post struct {
		Title string
		Body  string
	}
)

func main() {
	pages := buildPages()
	srv := server.NewServer(":8080", pages)

	log.Printf("Starting server (swoop v%s)...\n", swoop.Version())
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func buildPages() []page.IPage {
	base := layout.NewLayout("main.tmpl", map[string]interface{}{
		"title":       "Swoop",
		"description": "A simple web development library for producing server rendered HTML",
		"author":      "Michael Telford",
	})

	return []page.IPage{
		// Home page
		page.NewPageFromString("GET", "/", base, "Welcome!"),

		// About page
		page.NewPageFromString("GET", "/about", layout.NewLayoutFromLayout(base, map[string]interface{}{
			"title": "About",
		}), "Demo website for <a href=\"https://github.com/michaeltelford/swoop\" target=\"_blank\">Swoop</a>."),

		// Posts page
		page.NewPageFromComponents("GET", "/posts", layout.NewLayoutFromLayout(base, map[string]interface{}{
			"title": "Posts",
		}), buildPosts()),
	}
}

func buildPosts() []component.IComponent {
	ctx := map[string]interface{}{"posts": getPosts()}

	return []component.IComponent{
		component.NewComponent("Posts Summary", "posts_summary.tmpl", ctx),
		component.NewComponent("Posts", "posts.tmpl", ctx),
	}
}

func getPosts() []Post {
	return []Post{
		{
			"First Post",
			"This is the first post.",
		},
		{
			"Second Post",
			"This is the second post.",
		},
	}
}
