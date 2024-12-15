package compositions

import (
	"os"
	"strings"
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestGeneration(t *testing.T) {
	comp := Composition{
		Name:     "Test",
		Author:   "Test Author",
		Category: "Test Category",
		Init:     "Test",
	}
	t.Run("generate header", func(t *testing.T) {
		writer := strings.Builder{}

		err := RenderHeader(&writer, comp)
		if err != nil {
			t.Errorf("RenderComposition() failed: %v", err)
		}

		autogold.ExpectFile(t, autogold.Raw(writer.String()))
	})

	t.Run("generate composition", func(t *testing.T) {
		writer := strings.Builder{}

		err := RenderComposition(&writer, comp)
		if err != nil {
			t.Errorf("RenderComposition() failed: %v", err)
		}

		autogold.ExpectFile(t, autogold.Raw(writer.String()))
	})
}

func TestCleanSQF(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		dirty := "Test = 1;"
		want := "Test = 1;"
		got := CleanSQF(dirty)

		if got != want {
			t.Fatalf("CleanSQF() = %v, want %v", got, want)
		}
	})

	t.Run("multiline comments", func(t *testing.T) {
		dirty := `comment "
	TODO Expung3d:
 - EZM Eventhandlers
 - Add Dead Soldier compositions to all factions 
 - NATO+ 
 - Better Looters 
 - Paradrop Reinforcements 
 - Airdrop selected object 
 - Disable/Enable Thermals 
 - More waypoints 
 - Composition wrecks do not attach objects correctly 
 - Other seasonal modules 
 - Add more building interiors (Malden)
 - Advanced Difficulty Settings

 - Airstrike Helicopter
 - Cinematics
 - Play video module
";`

		got := CleanSQF(dirty)

		autogold.ExpectFile(t, autogold.Raw(got))
	})

	t.Run("evil discord thing", func(t *testing.T) {
		dirty := `[
					MAZ_zeusModulesTree,
					MAZ_EZMLabelTree,
					format ["ZAM Edition - %1",missionNamespace getVariable ['MAZ_EZM_Version','']],
					"Framework originally created by: M9-SD & GamesByChris.\nExpanded and made public by: Expung3d to enhance Public Zeus.\n\nNeed help? Found a bug? Join our Discord:\nhttps://discord.gg/W4ew5HP",
					"MAZ_EZM_fnc_hiddenEasterEggModule"
				] call MAZ_EZM_fnc_zeusAddModule;`
		got := CleanSQF(dirty)
		autogold.ExpectFile(t, autogold.Raw(got))
	})

	t.Run("entirety of ezm", func(t *testing.T) {
		dirty, err := os.ReadFile("testdata/TestCleanSQF/entirety_of_ezm.sqf")
		if err != nil {
			t.Fatalf("os.ReadFile() failed: %v", err)
		}

		got := CleanSQF(string(dirty))

		autogold.ExpectFile(t, autogold.Raw(got))
	})
}
