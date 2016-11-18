package slack

import (
	"testing"
)

func TestJSON(t *testing.T) {
	attachments := NewAttachments(2)

	attachments[0].SetFallback("a")
	attachments[1].SetFallback("b")

	expected := `[{"fallback":"a"},{"fallback":"b"}]`

	json, err := JSON(attachments)
	if err != nil {
		t.Error(err)
	}
	if json != expected {
		t.Errorf("json : got %s, expected %s", json, expected)
	}

}

func TestJSONFullParam(t *testing.T) {
	fields := NewFields(1)
	fields[0].SetTitle("i").SetValue("j").SetShort(true)

	attachments := NewAttachments(1)
	attachments[0].SetFallback("a").SetColor(COLOR_GOOD).SetPretext(
		"b",
	).SetAuthorName("c").SetAuthorLink("d").SetAuthorIcon("e").SetTitle(
		"f",
	).SetTitleLink("g").SetText("h").SetFields(fields).SetImageURL(
		"k",
	).SetThumbURL("l").SetFooterText("m").SetFooterIcon("n").SetFooterTs(123456789)

	expected := `[{"fallback":"a","color":"good","pretext":"b","author_name":"c","author_link":"d","author_icon":"e","title":"f","title_link":"g","text":"h","fields":[{"title":"i","value":"j","short":true}],"image_url":"k","thumb_url":"l","footer":"m","footer_icon":"n","ts":123456789}]`

	json, err := JSON(attachments)
	if err != nil {
		t.Error(err)
	}
	if json != expected {
		t.Errorf("json : got %s, expected %s", json, expected)
	}

}
