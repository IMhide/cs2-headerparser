package main

import (
	"log"
	"os"

	"parseheader/bitread"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
  anubisPath := "./demo/tn_anibus.dem"
  miragePath := "./demo/lf_mirage.dem"
  nukePath := "./demo/lf_nuke.dem"

  anubisFile, err := os.Open(anubisPath)
  if err != nil {
    log.Panic("failed to open demo file: ", err)
  }

  mirageFile, err := os.Open(miragePath)
  if err != nil {
    log.Panic("failed to open demo file: ", err)
  }

  nukeFile, err := os.Open(nukePath)
  if err != nil {
    log.Panic("failed to open demo file: ", err)
  }

  defer anubisFile.Close()
  defer mirageFile.Close()
  defer nukeFile.Close()

  anubisParser := bitread.CustomLargeBitReader(anubisFile, 280)
  mirageParser := bitread.CustomLargeBitReader(mirageFile, 280)
  nukeParser := bitread.CustomLargeBitReader(nukeFile, 280)

  printTable(anubisParser, mirageParser,nukeParser)
}

//
//
// Utils
//
//

// Note : Some of the variables of the protobuf doesn't show on my example sample so I didn't coded 'em
// https://github.com/SteamDatabase/GameTracking-CS2/blob/master/Protobufs/demo.proto

func printTable(anubisParser *bitread.BitReader, nukeParser *bitread.BitReader, mirageParser *bitread.BitReader){
  headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
  columnFmt := color.New(color.FgYellow).SprintfFunc()

  tbl := table.New("Name", "tn_anibus", "lf_nuke", "lf_mirage")
  tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
  tbl.AddRow("Filestamp", anubisParser.ReadString(), nukeParser.ReadString(), mirageParser.ReadString())  
  tbl.AddRow("Skip 16 bytes", Advance(anubisParser, 16), Advance(nukeParser, 16),Advance(mirageParser, 16))
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Filestamp", anubisParser.NewParseString(), nukeParser.NewParseString(), mirageParser.NewParseString())  
  tbl.AddRow("Network Protocol", anubisParser.ReadSignedInt(32), nukeParser.ReadSignedInt(32), mirageParser.ReadSignedInt(32))
  tbl.AddRow("Server name", anubisParser.NewParseString(), nukeParser.NewParseString(), mirageParser.NewParseString())  
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Client name", anubisParser.NewParseString(), nukeParser.NewParseString(), mirageParser.NewParseString())  
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Map name", anubisParser.NewParseString(), nukeParser.NewParseString(), mirageParser.NewParseString())  
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Game Dir", anubisParser.NewParseString(), nukeParser.NewParseString(), mirageParser.NewParseString())  
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Fullpacket version", anubisParser.ReadSignedInt(32), nukeParser.ReadSignedInt(32), mirageParser.ReadSignedInt(32))
  tbl.AddRow("Allow client-side entities", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Allow client-side particules", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Demo Version Name",anubisParser.NewParseString(), nukeParser.NewParseString(), mirageParser.NewParseString())  
  tbl.AddRow("Skip 1 bytes", Advance(anubisParser, 1), Advance(nukeParser, 1),Advance(mirageParser, 1))
  tbl.AddRow("Demo Version GUID", anubisParser.NewParseString(), nukeParser.NewParseString(), mirageParser.NewParseString())  
  tbl.AddRow("Build num", anubisParser.ReadSignedInt(32), nukeParser.ReadSignedInt(32), mirageParser.ReadSignedInt(32))
  tbl.AddRow("Next 32 bytes", Advance(anubisParser, 32), Advance(nukeParser, 32),Advance(mirageParser, 32))

  tbl.Print()
}

func Advance(parser *bitread.BitReader, bytes int) []byte{
  return parser.ReadBits(bytes*8)
}
