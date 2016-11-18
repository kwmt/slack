package slack

import (
	"encoding/json"
	// "fmt"
)

const (
	COLOR_GOOD    = "good"
	COLOR_WARNING = "warning"
	COLOR_DANGER  = "danger"
)

type Attachment struct {
	Fallback string `json:"fallback"`
	Color    string `json:"color,omitempty"`
	Pretext  string `json:"pretext,omitempty"`
	Author
	Title
	Text     string  `json:"text,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
	ImageUrl string  `json:"image_url,omitempty"`
	ThumbUrl string  `json:"thumb_url,omitempty"`
	Footer
}

type Author struct {
	Name string `json:"author_name,omitempty"`
	Link string `json:"author_link,omitempty"`
	Icon string `json:"author_icon,omitempty"`
}

type Title struct {
	Title string `json:"title,omitempty"`
	Link  string `json:"title_link,omitempty"`
}

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

type Footer struct {
	// brief text
	Footer string `json:"footer,omitempty"`
	Icon   string `json:"footer_icon,omitempty"`
	Ts     int64  `json:"ts,omitempty"`
}

func NewAttachments(count int) []Attachment {
	return make([]Attachment, count)
}

func JSON(attachments []Attachment) (string, error) {
	b, err := json.Marshal(attachments)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (attachment *Attachment) SetFallback(fallback string) *Attachment {
	attachment.Fallback = fallback
	return attachment
}
func (attachment *Attachment) SetText(text string) *Attachment {
	attachment.Text = text
	return attachment
}

// This is able to set value either be one of good, warning, danger, or any hex color code (eg. #439FE0).
func (attachment *Attachment) SetColor(color string) *Attachment {
	attachment.Color = color
	return attachment
}
func (attachment *Attachment) SetPretext(pretext string) *Attachment {
	attachment.Pretext = pretext
	return attachment
}
func (attachment *Attachment) SetAuthorName(authorName string) *Attachment {
	attachment.Author.Name = authorName
	return attachment
}
func (attachment *Attachment) SetAuthorLink(authorLink string) *Attachment {
	attachment.Author.Link = authorLink
	return attachment
}
func (attachment *Attachment) SetAuthorIcon(authorIcon string) *Attachment {
	attachment.Author.Icon = authorIcon
	return attachment
}
func (attachment *Attachment) SetTitle(title string) *Attachment {
	attachment.Title.Title = title
	return attachment
}
func (attachment *Attachment) SetTitleLink(titleLink string) *Attachment {
	attachment.Title.Link = titleLink
	return attachment
}
func (attachment *Attachment) SetImageURL(imageUrl string) *Attachment {
	attachment.ImageUrl = imageUrl
	return attachment
}
func (attachment *Attachment) SetThumbURL(thumbUrl string) *Attachment {
	attachment.ThumbUrl = thumbUrl
	return attachment
}
func (attachment *Attachment) SetFooterText(text string) *Attachment {
	attachment.Footer.Footer = text
	return attachment
}
func (attachment *Attachment) SetFooterIcon(icon string) *Attachment {
	attachment.Footer.Icon = icon
	return attachment
}

// timestamp
func (attachment *Attachment) SetFooterTs(ts int64) *Attachment {
	attachment.Footer.Ts = ts
	return attachment
}
func (attachment *Attachment) SetFields(fields []Field) *Attachment {
	attachment.Fields = fields
	return attachment
}

func NewFields(count int) []Field {
	return make([]Field, count)
}

func (field *Field) SetTitle(title string) *Field {
	field.Title = title
	return field
}
func (field *Field) SetValue(value string) *Field {
	field.Value = value
	return field
}
func (field *Field) SetShort(short bool) *Field {
	field.Short = short
	return field
}
