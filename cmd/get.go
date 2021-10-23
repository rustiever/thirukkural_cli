package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// getCmd fetches the kural for a given number, otherwise random kural will be returned
var getCmd = &cobra.Command{
	Use:   "get [ number from 1-1330 ]",
	Short: "Gets Kural for  specified number",
	Long: `You can specify either number ranging from 1-1330.
			If you don't specify then you get a random kural`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("only one valid number is expected")
		} else if len(args) == 0 {
			return nil
		} else if n, ok := strconv.Atoi(args[0]); ok != nil {
			return errors.New("valid number is expected")
		} else if (n-0)*(1331-n) <= 0 { // check whether number is between 1-1330
			return errors.New("number must be in range from 1-1330")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			kural Kural
			err   error
		)
		if len(args) == 0 {
			kural, err = fetchKural("rnd")
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			kural, err = fetchKural(args[0])
			if err != nil {
				log.Fatalln(err)
			}
		}

		printKural(kural)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

// prints Kural Data in tabular format
func printKural(kural Kural) {
	t := table.NewWriter()

	t.AppendRow(table.Row{"Number", kural.Number})
	t.AppendRow(table.Row{"Paal", kural.Paal})
	t.AppendRow(table.Row{"Athigaram", kural.Athigaram})
	t.AppendRow(table.Row{"Iyal", kural.Iyal})
	t.AppendRow(table.Row{"Line1", kural.Line1})
	t.AppendRow(table.Row{"Line2", kural.Line2})
	t.AppendRow(table.Row{"Translation", kural.Translation})
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 2, WidthMax: 150},
	})

	t.SetStyle(table.StyleRounded)

	fmt.Println(t.Render())
}

// calls API to fetch the Kural for given valid number
func fetchKural(number string) (kural Kural, err error) {
	var apiUrl string = "http://getthirukural.appspot.com/api/3.0/kural/" + number + "?appid=" + apiKey + "&format=json"
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &kural)
	if err != nil {
		return kural, err
	}
	return kural, nil
}

// Kural struct to decode the json payload from API
type Kural struct {
	Paal        string `json:"paal"`
	Urai3       string `json:"urai3"`
	Iyal        string `json:"iyal"`
	Urai1       string `json:"urai1"`
	Urai3Author string `json:"urai3Author"`
	Urai2       string `json:"urai2"`
	Number      int    `json:"number"`
	Urai1Author string `json:"urai1Author"`
	Athigaram   string `json:"athigaram"`
	Urai2Author string `json:"urai2Author"`
	Translation string `json:"translation"`
	Line2       string `json:"line2"`
	Line1       string `json:"line1"`
}
