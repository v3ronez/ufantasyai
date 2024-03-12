package sb

import (
	"os"

	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func InitSB() error {
	Client = supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SECRET"))
	return nil
}
