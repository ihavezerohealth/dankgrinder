// Copyright (C) 2021 The Dank Grinder authors.
//
// This source code has been released under the GNU Affero General Public
// License v3.0. A copy of this license is available at
// https://www.gnu.org/licenses/agpl-3.0.en.html

package discord

import (
	"encoding/json"
	"time"
)

const (
	EmbedTypeRich     = "rich"
	EmbedTypeImage    = "image"
	EmbedTypeVideo    = "video"
	EmbedTypeGifVideo = "gifv"
	EmbedTypeArticle  = "article"
	EmbedTypeLink     = "link"
)

const (
	MessageTypeDefault = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTier1
	MessageTypeUserPremiumGuildSubscriptionTier2
	MessageTypeUserPremiumGuildSubscriptionTier3
	MessageTypeChannelFollowAdd
	MessageTypeGuildDiscoveryDisqualified = iota + 1
	MessageTypeGuildDiscoveryRequalified
	MessageTypeReply = iota + 4
	MessageTypeApplicationCommand
)

type Message struct {
	// The ID of the message.
	ID string `json:"id,omitempty"`

	// The ID of the channel the message was sent in.
	ChannelID string `json:"channel_id,omitempty"`

	// The ID of the guild the message was sent in.
	GuildID string `json:"guild_id,omitempty"`

	// The author of the message. Not guaranteed to be a valid user.
	//
	// The author object follows the structure of the User object, but is only a
	// valid user in the case where the message is generated by a user or bot
	// user. If the message is generated by a webhook, the author object
	// corresponds to the webhook's ID, username, and avatar. You can tell if a
	// message is generated by a webhook by checking for the WebhookID on the
	// message object.
	Author User `json:"author,omitempty"`

	// The contents of the message.
	Content string `json:"content,omitempty"`

	// When the message was sent.
	Time time.Time `json:"timestamp,omitempty"`

	// When this message was edited (default value if never).
	EditedTime time.Time `json:"edited_timestamp,omitempty"`

	// Whether this is a TTS message.
	TTS bool `json:"tts,omitempty"`

	// Whether this message mentions everyone.
	MentionEveryone bool `json:"mention_everyone,omitempty"`

	// Users specifically mentions in the message.
	Mentions []User `json:"mentions"`

	// The type of message.
	Type int `json:"type,omitempty"`

	// Whether the message is pinned.
	Pinned bool `json:"pinned,omitempty"`

	// Any embedded content.
	Embeds []Embed `json:"embeds,omitempty"`

	// If the message is generated by a webhook, this is the webhook's ID.
	WebhookID string `json:"webhook_id,omitempty"`

	// A list of components attached to the message.
	Components []MessageComponent `json:"-"`

	// The message which this message references. This field is only set for
	// messages with type MessageTypeReply.
	//
	// If the message is a reply but this field is not set, either the Discord
	// backend did not attempt to fetch the message that was being replied to,
	// or the referenced message was deleted.
	ReferencedMessage *Message `json:"referenced_message,omitempty"`
}

// UnmarshalJSON is a helper function to unmarshal the Message.
func (m *Message) UnmarshalJSON(data []byte) error {
	type message Message
	var v struct {
		message
		RawComponents []unmarshalableMessageComponent `json:"components"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*m = Message(v.message)
	m.Components = make([]MessageComponent, len(v.RawComponents))
	for i, v := range v.RawComponents {
		m.Components[i] = v.MessageComponent
	}
	return err
}

type Embed struct {
	Title string `json:"title,omitempty"`

	// The type of embed. Always EmbedTypeRich for webhook embeds.
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`

	// The time at which the embed was sent.
	Time time.Time `json:"timestamp,omitempty"`

	// The color code of the embed.
	Color    int           `json:"color,omitempty"`
	Footer   EmbedFooter   `json:"footer,omitempty"`
	Provider EmbedProvider `json:"provider,omitempty"`
	Author   EmbedAuthor   `json:"author,omitempty"`
	Fields   []EmbedField  `json:"fields,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
