/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [ number from 1-1330 ]",
	Short: "Gets Kural for  specified number",
	Long: `You can specify either number ranging from 1-1330.
			If you don't specify then you get a random kural`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("pass only one valid number, otherwise don't pass")
		} else if len(args) == 0 {
			return nil
		} else if n, ok := strconv.Atoi(args[0]); ok != nil {
			return errors.New("pass valid number ")
		} else if (n-0)*(1330-n) <= 0 { // check whether number is between 1-1330
			return errors.New("number must range from 1-1330")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var kural Kural
		var err error
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

		// printKural(kural)
		printk(kural)
	},
}

func printk(kural Kural) {
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
	// t.AppendRow(table.Row{"Urai-1", kural.Urai1})
	// t.AppendRow(table.Row{"Urai-2", kural.Urai2, "Urai Author", kural.Urai2Author})
	// t.AppendRow(table.Row{"Urai-3", kural.Urai3, "Urai Author", kural.Urai3Author})

	t.SetStyle(table.StyleRounded)

	fmt.Println(t.Render())
}

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

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func printKural(kural Kural) {
// 	fmt.Printf("Number: %v\n", kural.Number)
// 	fmt.Printf("Paal: %v\n", kural.Paal)
// 	fmt.Printf("Athigaram: %v\n", kural.Athigaram)
// 	fmt.Printf("Iyal: %v\n", kural.Iyal)
// 	fmt.Printf("Line1: %v\n", kural.Line1)
// 	fmt.Printf("Line2: %v\n", kural.Line2)
// 	fmt.Printf("Translation: %v\n", kural.Translation)
// 	fmt.Printf("Athigaram: %v\n", kural.Urai1)
// 	fmt.Printf("Athigaram: %v\n", kural.Urai1Author)
// 	fmt.Printf("Athigaram: %v\n", kural.Urai2)
// 	fmt.Printf("Athigaram: %v\n", kural.Urai2Author)
// 	fmt.Printf("Athigaram: %v\n", kural.Urai3)
// 	fmt.Printf("Athigaram: %v\n", kural.Urai3Author)
// }
