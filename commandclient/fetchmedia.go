package commandclient

import (
	"errors"
	"fmt"
	"os"

	"github.com/minetest-go/minetest_client/commands"
)

func FetchMedia(cc *CommandClient) error {
	ch := make(chan commands.Command, 1000)
	cc.AddListener(ch)
	defer cc.RemoveListener(ch)
	received_count := 0
	expected_count := 0

	for o := range ch {
		switch cmd := o.(type) {

		case *commands.ServerAnnounceMedia:
			fmt.Printf("Server announces media: %d assets\n", cmd.FileCount)

			_, err := os.Stat("media")
			if os.IsNotExist(err) {
				err := os.Mkdir("media", 0755)
				if err != nil {
					return err
				}
			}

			fmt.Printf("Sending REQUEST_MEDIA len=%d\n", len(cmd.Hashes))
			files := make([]string, 0)
			for name := range cmd.Hashes {
				//fmt.Printf("Name: '%s'\n", name)

				_, err := os.Stat("media/" + name)
				if os.IsNotExist(err) {
					files = append(files, name)
					expected_count++
				}

			}

			if len(files) > 0 {
				reqmedia_cmd := commands.NewClientRequestMedia(files)
				err = cc.SendCommand(reqmedia_cmd)
				if err != nil {
					return err
				}
			} else {
				// nothing to fetch
				return nil
			}

		case *commands.ServerMedia:
			fmt.Printf("Server media: %s\n", cmd)
			received_count += int(cmd.NumFiles)

			for name, data := range cmd.Files {
				_, err := os.Stat("media/" + name)
				if os.IsNotExist(err) {
					err = os.WriteFile("media/"+name, data, 0644)
					if err != nil {
						return err
					}
				}
			}

			fmt.Printf("Media status: received=%d, expected=%d\n", received_count, expected_count)
			if received_count >= expected_count {
				fmt.Println("Received all media files")
				return nil
			}
		}
	}

	return errors.New("channel closed")
}
