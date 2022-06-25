package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {

	rock := playlist{genre: "Rock", songs: []song{
		{title: "The Sound of Silence", artist: "Beethoven"},
		{title: "The Sound of Silence2", artist: "Beethoven2"},
	}}

	song := rock.songs[0]
	song.title = "The Sound of Silence3"

	rock.songs[0].title = "Live forever"

	fmt.Printf("%+v\n%+v\n\n", song, rock.songs[0])

	for _, s := range rock.songs {
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	fmt.Printf("\n%-20s %20s\n", "Title", "Artist")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
